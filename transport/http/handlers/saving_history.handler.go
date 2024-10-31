package handler

import (
	// "fmt"

	"go-clean-architecture-by-ahr/domain"
	helper_http "go-clean-architecture-by-ahr/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type SavingHistoryHandler struct {
	usecase domain.SavingHistoryUsecase
}

func SavingHistoryRoute(e *echo.Echo, uc domain.SavingHistoryUsecase) {
	handler := SavingHistoryHandler {
		usecase: uc,
	}
	e.POST("/saving_history", handler.Create)
	e.GET("/saving_history", handler.GetAllHandler)
	e.GET("/saving_history/:id", handler.GetByIDHandler)
// 	e.PUT("/saving_history/:id", handler.Update)
// 	e.DELETE("/saving_history/:id", handler.Delete)
}     

func (h *SavingHistoryHandler) GetAllHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil{
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all saving_history")

	return resp
}

func (h *SavingHistoryHandler) GetByIDHandler(c echo.Context) error {
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

func (h *SavingHistoryHandler) Create(c echo.Context) error {
	var data domain.SavingHistory

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Create(&data) 
	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, data, "success create saving_history")
}

// func (h *SavingHistoryHandler) Update(c echo.Context) error {
// 	var data domain.SavingHistory
	
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
// 	return helper_http.SuccessResponse(c, data, "success update saving_history")

// }

// func (h *SavingHistoryHandler) Delete(c echo.Context) error {
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

