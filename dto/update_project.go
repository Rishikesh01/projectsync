package dto

import "github.com/google/uuid"

type UpdateProject struct {
	ID      uuid.UUID  `json:"id,omitempty"`
	Project NewProject `json:"project"`
}
