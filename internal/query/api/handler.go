package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (srv *Server) GetQuery(e echo.Context) error {
	// get query name
	name := e.QueryParam("name")
	msg, err := srv.uc.FetchQuery(name)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, msg)
}

func (srv *Server) PostQuery(e echo.Context) error {
	// post query name
	name := e.QueryParam("name")
	err := srv.uc.InsertQuery(name)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.NoContent(http.StatusOK)
}
