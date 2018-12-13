package controller

import (
	"github.com/VividCortex/mysqlerr"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/zcong1993/libgo/gin"
	"github.com/zcong1993/libgo/utils"
	"github.com/zcong1993/rest-go/common"
	"github.com/zcong1993/rest-go/model"
	mysql2 "github.com/zcong1993/rest-go/mysql"
	"github.com/zcong1993/rest-go/service"
	"net/http"
)

func Register(c *gin.Context) ginerr.ApiError {
	var f common.RegisterForm
	err := c.ShouldBindJSON(&f)
	if err != nil {
		return common.INVALID_PARAMS
	}

	u := &model.User{
		Username: f.Username,
		Email:    f.Email,
		Password: f.Password,
	}

	err = u.Save()
	if err != nil {
		if err, ok := err.(*mysql.MySQLError); ok && err.Number == mysqlerr.ER_DUP_ENTRY {
			return common.DUPLICATE_USER
		}
		return common.INTERVAL_ERROR
	}

	tk, err := service.GetOrCreateToken(*u)

	if err != nil {
		return common.INTERVAL_ERROR
	}

	c.JSON(http.StatusCreated, tk)
	return nil
}

func Login(c *gin.Context) ginerr.ApiError {
	var f common.LoginForm
	err := c.ShouldBindJSON(&f)
	if err != nil {
		return common.INVALID_PARAMS
	}
	var u model.User
	err = mysql2.DB.First(&u, "email=?", f.Email).Error
	if err != nil {
		if err.Error() == "record not found" {
			return common.INVALID_USERNAME_OR_PASSWORD
		}

		return common.INTERVAL_ERROR
	}

	if !utils.ComparePassword(f.Password, u.Password) {
		return common.INVALID_USERNAME_OR_PASSWORD
	}

	tk, err := service.GetOrCreateToken(u)

	if err != nil {
		return common.INTERVAL_ERROR
	}

	c.JSON(http.StatusOK, tk)

	return nil
}

func Me(c *gin.Context) {
	v, _ := c.Get(common.AUTH_CONTEXT_KEY)
	vv, ok := v.(*model.User)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, vv)
}
