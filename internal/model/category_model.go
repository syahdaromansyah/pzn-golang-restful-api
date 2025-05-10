package model

type (
	CategoryResponse struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	CreateCategoryRequest struct {
		Name string `json:"name" validate:"required,min=3,max=128"`
	}

	UpdateCategoryRequest struct {
		Name string `json:"name" validate:"required,min=3,max=128"`
	}
)
