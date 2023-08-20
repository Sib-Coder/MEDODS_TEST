package endpoint

//import (
//	"awesomeProject13/internal/model"
//	"fmt"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/labstack/echo/v4"
//	"net/http"
//	"time"
//)
//
//// var SECRET = []byte("super-secret-auth-key") //вынести в env
//
//func (e *Endpoint) ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//
//		if ctx.Request().Header["Token"] != nil {
//			token, _ := jwt.Parse(ctx.Request().Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
//				_, ok := t.Method.(*jwt.SigningMethodHMAC)
//				if !ok {
//					ctx.String(http.StatusUnauthorized, "not authorized: ")
//
//				}
//				return SECRET, nil
//			})
//
//			if token.Valid {
//				next(ctx)
//			}
//
//			tokentime := token.Claims.(jwt.MapClaims)["exp"].(float64)
//			tokentime2 := time.Unix(int64(tokentime), 0)
//
//			//fmt.Println("Token", tokentime, "\n Now", floatTime)
//			if time.Now().After(tokentime2) {
//				reftoken, err := jwt.Parse(ctx.Request().Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
//					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//					}
//					return []byte(SECRET), nil
//				})
//				if err != nil {
//					ctx.String(http.StatusUnauthorized, "not valid token")
//				}
//				idx := reftoken.Claims.(jwt.MapClaims)["idx"].(string)
//				//добавить проверку токена из бд
//				dbtoken, err := e.s.GetRefToken(idx) //получение токена из базы для сравнения
//				if err != nil {
//					//return err
//					ctx.String(http.StatusUnauthorized, "not authorized: ")
//				}
//				rt, err := reftoken.SignedString([]byte(SECRET))
//				if dbtoken == rt {
//					token, refresh, err := e.CreateJWT(idx)
//					if err != nil {
//						ctx.String(http.StatusUnauthorized, "not authorized: ")
//					}
//					tockens := model.Tokens{
//						Token:        token,
//						Refreshtoken: refresh,
//					}
//
//					ctx.JSON(http.StatusUnauthorized, tockens)
//				}
//
//			}
//
//		} else {
//			ctx.String(http.StatusUnauthorized, "not authorized: ")
//
//		}
//		return nil
//	}
//}
