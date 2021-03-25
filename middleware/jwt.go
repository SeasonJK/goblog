package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goblog/utils"
	"goblog/utils/errmsg"
	"net/http"
	"strings"
)

type JWT struct {
	JwtKey []byte
}

func NewJwt() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

type MyClaims struct {
	username string `json:"username"`
	jwt.StandardClaims
}

var (
	TokenExpied 		error = errors.New("token已经过期，请重新登录")
	TokenNotValidYet 	error = errors.New("token无效，请重新登录")
	TokenMaformed	 	error = errors.New("token不正确，请重新登录")
	TokenInvalid	 	error = errors.New("非token，请重新登录")
)

// 生成token
func (j *JWT)CreateToken(claims MyClaims)(string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// 解析token
func (j *JWT)ParseToken(tokenString string)(*MyClaims, error){
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	if err != nil{
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors & jwt.ValidationErrorMalformed != 0{
				return nil, TokenMaformed
			}else if ve.Errors & jwt.ValidationErrorExpired != 0{
				return nil, TokenExpied
			}else if ve.Errors & jwt.ValidationErrorNotValidYet != 0{
				return nil, TokenNotValidYet
			}else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil{
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}

// jwt 中间件
func JwtToken() gin.HandlerFunc{
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == ""{
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0{
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if len(checkToken) != 2 || checkToken[0] == "Bearer"{
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		j := NewJwt()
		// 解析token
		claims, err := j.ParseToken(checkToken[1])
		if err != nil{
			if err == TokenExpied {
				c.JSON(http.StatusOK, gin.H{
					"status": code,
					"message":  "token已过期，请重新登录",
					"data": nil,
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": errmsg.ERROR,
				"message": errmsg.GetErrMsg(errmsg.ERROR),
				"data": nil,
			})
			c.Abort()
			return
		}
		c.Set("username", claims)
		c.Next()
	}
}
























