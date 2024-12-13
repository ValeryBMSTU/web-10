package api

import (
	"errors"
	"net/http"

	"github.com/ValeryBMSTU/web-10/pkg/vars"

	"github.com/labstack/echo/v4"
)

func (srv *Server) Handler(e echo.Context) error {
	name := e.QueryParam("name")
	if name != "" {

		if len([]rune(name)) > srv.maxSize {
			return e.String(http.StatusBadRequest, "name is too large")
		}

		err := srv.uc.SetHelloMessage(name)
		if err != nil {
			if !errors.Is(err, vars.ErrAlreadyExist) {
				return e.String(http.StatusInternalServerError, err.Error())
			}
		}

		return e.String(http.StatusOK, "Hello, "+name+"!")
	} else {
		msg, err := srv.uc.FetchHelloMessage()
		if err != nil {
			return e.String(http.StatusInternalServerError, err.Error())
		}
		return e.String(http.StatusOK, "Hello, "+msg+"!")
	}
}
