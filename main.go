package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"projectsync/handler"
	"projectsync/repo"
	"projectsync/service"
)

type env struct {
	Host     string `required:"true"`
	Port     int    `required:"true"`
	Sslmode  string `required:"true" split_words:"true"`
	TimeZone string `required:"true" split_words:"true"`
	User     string `required:"true"`
	DbName   string `required:"true"`
	Password string `required:"true"`
}

func dbconnconfig(env *env) string {
	dsn := fmt.Sprintf("host=%suser=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", env.Host, env.User, env.Password, env.DbName, env.Port, env.Sslmode, env.TimeZone)
	return dsn
}

func checkAuth(service service.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		err := service.ValidateToken(tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}

func main() {
	var e env
	err := envconfig.Process("", &e)
	if err != nil {
		log.Fatal(err)
		return
	}
	r := gin.Default()
	db, err := gorm.Open(postgres.Open(dbconnconfig(&e)), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	projectRepo := repo.NewProjectRepo(db)
	userdetailssRepo := repo.NewUserdetailsRepo(db)
	svcProject := service.NewProjectService(projectRepo)
	svcUserDetails := service.NewUserService(userdetailssRepo)
	svcAuth := service.NewAuthService(userdetailssRepo)
	projectHandler := handler.NewProjectHandler(svcProject)
	userHandler := handler.NewUserHandler(svcUserDetails, svcAuth)

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	v1 := r.RouterGroup.Group("/v1")
	v1.Use(checkAuth(svcAuth))
	v1.GET("/project/list/:id", projectHandler.GetAllProjects)
	v1.POST("/projects/add", projectHandler.AddProjects)
	v1.PATCH("/projects/update", projectHandler.UpdateProject)
	v1.DELETE("/projects/delete/:id", projectHandler.DeleteProject)

}
