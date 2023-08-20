package endpoint

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var SECRET = []byte("super-secret-auth-key") //TODO вынести в env

func (e *Endpoint) CreateJWT(idx string) (string, string, error) { //бизнес логика

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix() //дата
	claims["idx"] = idx                              //добавление ника пользователя в jwt

	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		fmt.Println(err.Error())
		return "", " ", err
	}
	refreshClaims := jwt.MapClaims{
		"idx": idx,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString([]byte(SECRET))

	// Store refresh token in database
	_, err = e.s.UpdateRefToken(idx, refreshString)
	if err != nil {
		return "", "", err
	}
	return tokenStr, refreshString, nil
}

func ExtracIdxFromJWT(tokenString string) (string, error) {
	var idx string
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		idx = fmt.Sprint(claims["idx"])
	}

	if idx == "" {
		return "", fmt.Errorf("invalid token payload")
	}
	return idx, nil
}
