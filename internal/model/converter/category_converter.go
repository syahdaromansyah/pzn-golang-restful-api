package converter

import (
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/entity"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/model"
)

func CategoryToResponse(category *entity.Category) *model.CategoryResponse {
	return &model.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func CategoriesToResponse(categories []entity.Category) []model.CategoryResponse {
	categoriesResponse := []model.CategoryResponse{}

	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, *CategoryToResponse(&category))
	}

	return categoriesResponse
}
