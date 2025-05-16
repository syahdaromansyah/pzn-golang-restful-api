package http

import (
	"net/http"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/exception"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/model"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/usecase"

	"github.com/julienschmidt/httprouter"
)

type categoryControllerImpl struct {
	UseCase usecase.CategoryUseCase
}

func NewCategoryControllerImpl(useCase usecase.CategoryUseCase) CategoryController {
	return &categoryControllerImpl{
		UseCase: useCase,
	}
}

func (c *categoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := new(model.CreateCategoryRequest)

	err := helper.ReadFromRequestBody(r, categoryCreateRequest)
	helper.ClientPanicIfError(err, exception.NewErrorClientRequest(err, http.StatusBadRequest, "malformed request body"))

	categoryResponse := c.UseCase.Create(r.Context(), categoryCreateRequest)

	webResponse := &model.WebResponse[*model.CategoryResponse]{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   categoryResponse,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = helper.WriteToResponseBody(w, webResponse)
	helper.InternalServerPanicIfError(err, "category > http/controller > Create")
}

func (c *categoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryUpdateRequest := new(model.UpdateCategoryRequest)

	err := helper.ReadFromRequestBody(r, categoryUpdateRequest)
	helper.ClientPanicIfError(err, exception.NewErrorClientRequest(err, http.StatusBadRequest, "malformed request body"))

	categoryId := params.ByName("categoryId")

	categoryResponse := c.UseCase.Update(r.Context(), categoryId, categoryUpdateRequest)

	webResponse := &model.WebResponse[*model.CategoryResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = helper.WriteToResponseBody(w, webResponse)
	helper.InternalServerPanicIfError(err, "category > http/controller > Update")
}

func (c *categoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")

	c.UseCase.Delete(r.Context(), categoryId)

	webResponse := &model.WebResponseMessage{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "category is successfully deleted",
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := helper.WriteToResponseBody(w, webResponse)
	helper.InternalServerPanicIfError(err, "category > http/controller > Delete")
}

func (c *categoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")

	categoryResponse := c.UseCase.FindById(r.Context(), categoryId)

	webResponse := &model.WebResponse[*model.CategoryResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := helper.WriteToResponseBody(w, webResponse)
	helper.InternalServerPanicIfError(err, "category > http/controller > FindById")
}

func (c *categoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoriesResponse := c.UseCase.FindAll(r.Context())

	webResponse := &model.WebResponse[[]model.CategoryResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoriesResponse,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := helper.WriteToResponseBody(w, webResponse)
	helper.InternalServerPanicIfError(err, "category > http/controller > FindAll")
}
