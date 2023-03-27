package dto

type Project struct {
	ProjectName string `json:"project_name,omitempty"`
	GithubLink  string `json:"github_link,omitempty"`
	IsActive    bool   `json:"is_active,omitempty"`
}
