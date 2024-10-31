package handler

import (
	// "fmt"

	"go-clean-architecture-by-ahr/domain"
	helper_http "go-clean-architecture-by-ahr/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type ProductHandler struct {
	usecase domain.ProductUsecase
}

func ProductRoute(e *echo.Echo, uc domain.ProductUsecase) {
	handler := ProductHandler {
		usecase: uc,
	}
	e.POST("/product", handler.Create)
	e.GET("/product", handler.GetAllHandler)
	e.GET("/product/:id", handler.GetByIDHandler)
// 	e.PUT("/product/:id", handler.Update)
// 	e.DELETE("/product/:id", handler.Delete)
}     

func (h *ProductHandler) GetAllHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil{
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all product")

	return resp
}

func (h *ProductHandler) GetByIDHandler(c echo.Context) error {
	id := c.Param("id")
	// id = fmt.Sprintf("%s")
	// num, err := strconv.Atoi(id)

	// if err != nil {
	// 	panic(err)
	// }

	data, err := h.usecase.GetByID(id)

	if err != nil{
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by id")
	return resp
}

func (h *ProductHandler) Create(c echo.Context) error {
	var data domain.Product

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Create(&data) 
	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, data, "success create product")
}

// func (h *ProductHandler) Update(c echo.Context) error {
// 	var data domain.Product
	
// 	err := c.Bind(&data)
// 	if err != nil {
// 		return c.JSON(http.StatusUnprocessableEntity, err.Error())
// 	}

// 	id := c.Param("id")
	
// 	affect, err := h.usecase.Update(id, &data) 
	
// 	if affect == 0 {
// 		return helper_http.NotFoundResponse(c, err.Error())
// 	}


// 	if err != nil {
// 		return helper_http.ErrorResponse(c, err)
// 	}
// 	return helper_http.SuccessResponse(c, data, "success update product")

// }

// func (h *ProductHandler) Delete(c echo.Context) error {
// 	id := c.Param("id")
// 	affect, err := h.usecase.Delete(id) 

// 	if affect == 0 {
// 		return helper_http.NotFoundResponse(c, err.Error())
// 	}

// 	if err != nil {
// 		return helper_http.ErrorResponse(c, err)
// 	}
// 	return c.NoContent(http.StatusNoContent)

// }

