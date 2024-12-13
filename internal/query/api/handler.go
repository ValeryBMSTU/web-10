package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get возвращает приветствие пользователю
func (srv *Server) GetQuery(e echo.Context) error {
	name := e.QueryParam("name")
	msg, err := srv.uc.SelectNameQuery(name)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, msg)
}

// Post Помещает новый вариант приветствия в БД
func (srv *Server) PostQuery(e echo.Context) error {
	name := e.QueryParam("name")

	err := srv.uc.InsertNameQuery(name)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.NoContent(http.StatusOK)
}
