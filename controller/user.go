package controller

import (
	"github.com/VividCortex/mysqlerr"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/libgo/gin/ginerr"
	"github.com/zcong1993/libgo/gin/ginhelper"
	"github.com/zcong1993/libgo/utils"
	"github.com/zcong1993/libgo/validator"
	"github.com/zcong1993/rest-go/common"
	"github.com/zcong1993/rest-go/model"
	mysql2 "github.com/zcong1993/rest-go/mysql"
	"github.com/zcong1993/rest-go/service"
	"net/http"
)

// Register is handler for /register
// @Summary Register a new user
// @Description Register a new user
// @Accept  json
// @Produce  json
// @Param   user     body    common.RegisterForm     true        "Register user"
// @Success 201 {object} model.Token
// @Failure 400 {object} common.ErrResp
// @Failure 500 "StatusInternalServerError"
// @Router /register [post]
func Register(c *gin.Context) ginerr.ApiError {
	var f common.RegisterForm
	err := c.ShouldBindJSON(&f)
	if err != nil {
		return common.CreateInvalidErr(validator.NormalizeErr(err))
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
		return common.CreateInvalidErr(validator.NormalizeErr(err))
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

// Me get login user message
// @Summary Me message
// @Description
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User
// @Failure 401 {object} common.ErrResp
// @Failure 500 "StatusInternalServerError"
// @Security ApiKeyAuth
// @Router /me [get]
func Me(c *gin.Context) {
	v, _ := c.Get(common.AUTH_CONTEXT_KEY)
	vv, ok := v.(*model.User)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, vv)
}

// Users query by id for /users/:id
// @Summary Users query by id
// @Description
// @Accept  json
// @Produce  json
// @Param   id  path  string  true  "User id"
// @Success 200 {object} model.User
// @Failure 500 "StatusInternalServerError"
// @Router /users/{id} [get]
func UsersGet(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	err := mysql2.DB.First(&user, "id = ?", id).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}
	ginhelper.WithCacheControl(c, common.LONG_CACHE_DURATION)
	c.JSON(http.StatusOK, user)
}
