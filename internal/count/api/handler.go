package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (srv *Server) GetCounter(e echo.Context) error {
	value, err := srv.uc.FetchCount()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.String(http.StatusOK, strconv.Itoa(value))
}

func (srv *Server) PostCounter(e echo.Context) error {
	a, err := strconv.Atoi(e.FormValue("count"))
	if err != nil {
		e.Logger().Error(err)
		return e.String(http.StatusBadRequest, "это не число")
	}

	if a > srv.maxSize {
		return e.String(http.StatusBadRequest, "число слишком большое")
	}

	err = srv.uc.IncreaseCount(a)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusOK, "OK!")
}
