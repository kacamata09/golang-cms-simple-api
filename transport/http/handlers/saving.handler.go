package handler

import (
	// "fmt"

	"go-clean-architecture-by-ahr/domain"
	helper_http "go-clean-architecture-by-ahr/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type SavingHandler struct {
	usecase domain.SavingUsecase
}

func SavingRoute(e *echo.Echo, uc domain.SavingUsecase) {
	handler := SavingHandler {
		usecase: uc,
	}
	e.POST("/saving", handler.Create)
	e.GET("/saving", handler.GetAllHandler)
	e.GET("/saving/:id", handler.GetByIDHandler)
// 	e.PUT("/saving/:id", handler.Update)
// 	e.DELETE("/saving/:id", handler.Delete)
}     

func (h *SavingHandler) GetAllHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil{
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all saving")

	return resp
}

func (h *SavingHandler) GetByIDHandler(c echo.Context) error {
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

func (h *SavingHandler) Create(c echo.Context) error {
	var data domain.Saving

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Create(&data) 
	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, data, "success create saving")
}

// func (h *SavingHandler) Update(c echo.Context) error {
// 	var data domain.Saving
	
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
// 	return helper_http.SuccessResponse(c, data, "success update saving")

// }

// func (h *SavingHandler) Delete(c echo.Context) error {
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

