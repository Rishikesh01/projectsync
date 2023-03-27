package service

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"log"
	"projectsync/dto"
	"projectsync/repo"
	"projectsync/utils"
	"time"
)

type AuthService interface {
	AuthenticateUser(in dto.SignIn) (string, error)
}

type userClaimsToken struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}

type JwtMiddleware struct {
	Realm          string
	Key            []byte
	TimeOut        time.Duration
	RefreshTimeOut time.Duration
	TimeFunc       func() time.Time
}

type jwtAuthService struct {
	userRepo repo.UserdetailsRepo
	*JwtMiddleware
}

func (j *jwtAuthService) AuthenticateUser(in dto.SignIn) (string, error) {
	model, err := j.userRepo.FindByEmail(in.Email)
	if err != nil {
		return "", err
	}
	err = utils.BcryptUtil{}.CheckPasswordHash(in.Password, model.Password)
	if err != nil {
		return "", err
	}
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	claim := userClaimsToken{
		model.ID.String(),
		model.Email,
		model.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(j.TimeFunc().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(j.TimeFunc()),
			NotBefore: jwt.NewNumericDate(j.TimeFunc()),
			Issuer:    j.Realm,
			Subject:   "user",
			ID:        uid.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)

	signedToken, err := token.SignedString([]byte("demo"))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return signedToken, nil
}

func NewAuthService(userRepo repo.UserdetailsRepo) AuthService {
	return &jwtAuthService{userRepo: userRepo}
}
