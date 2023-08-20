package endpoint

import (
	"awesomeProject13/internal/model"
	"fmt"
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
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(user.ID)

	idx := user.ID
	if idx != " " {
		fmt.Println(idx)
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
func (e *Endpoint) Secret(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(user.ID)

	idx := user.ID
	res, err := e.s.SelectSecret(idx)
	if err != nil {
		return c.String(http.StatusOK, "Error")
	}
	return c.String(http.StatusOK, res.Secter)
}
