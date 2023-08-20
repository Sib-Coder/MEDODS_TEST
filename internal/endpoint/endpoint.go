package endpoint

import (
	"awesomeProject13/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	UpdateRefToken(idx string, reft string) (bool, error)
	SelectSecret(idx string) (model.Traning, error)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) GetJwt(c echo.Context) error {
	idx := c.Param("id")
	if idx != " " {

		token, refresh, err := e.CreateJWT(idx)
		if err != nil {
			return err
		}
		tokens := model.Tokens{
			Token:        token,
			Refreshtoken: refresh,
		}
		return c.JSON(http.StatusOK, tokens)
	}
	return c.String(http.StatusUnauthorized, "Param id nil")
}
