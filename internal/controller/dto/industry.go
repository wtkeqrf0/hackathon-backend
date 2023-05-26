package dto

type IndustryBranch struct {
	Branch string `json:"branch,omitempty" validate:"required,lte=150" example:"Авиационная промышленность"`
}
