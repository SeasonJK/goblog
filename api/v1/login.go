package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goblog/middleware"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"time"
)

// 后台登录
func Login(c *gin.Context){
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int

	formData, code = model.CheckLogin(formData.Username, formData.Password)

	if code == errmsg.SUCCESS{
		setToken(c, formData)
	}else {
		c.JSON(http.StatusOK, gin.H{
			"status": 		code,
			"username":		formData.Username,
			"id": 			formData.ID,
			"message": 		errmsg.GetErrMsg(code),
			"token":		token,
		})
	}
}


//前台登录
func LoginFront(c *gin.Context){
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var code int

	formData, code = model.CheckLogin(formData.Username, formData.Password)

	c.JSON(http.StatusOK, gin.H{
		"status": 		code,
		"data":			formData.Username,
		"id": 			formData.ID,
		"message": 		errmsg.GetErrMsg(code),
	})
}

// 生成token
func setToken(c *gin.Context, user model.User){
	j := middleware.NewJwt()
	claims := middleware.MyClaims{
		Username: user.Username,
		StandardClaims:jwt.StandardClaims{
			NotBefore: time.Now().Unix()-100,
			ExpiresAt: time.Now().Unix()+7200,
			Issuer: "goblog",
		},
	}
	token, err := j.CreateToken(claims)

	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token": token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": errmsg.SUCCESS,
		"data": user.Username,
		"id": user.ID,
		"message": errmsg.GetErrMsg(errmsg.SUCCESS),
		"token": token,
	})
	return
}










