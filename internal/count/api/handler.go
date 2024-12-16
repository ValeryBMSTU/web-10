package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetHello возвращает случайное приветствие пользователю
func (srv *Server) GetCount(e echo.Context) error {
	msg, err := srv.uc.FetchCount()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, msg)
}

// PostHello Помещает новый вариант приветствия в БД
func (srv *Server) PostCount(e echo.Context) error {
	input := struct {
		Msg *int `json:"count"`
	}{}

	err := e.Bind(&input)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	if input.Msg == nil {
		return e.String(http.StatusBadRequest, "msg is empty")
	}

	res, err := srv.uc.IncrementCount(*input.Msg)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]int{"new_count": res})

}
