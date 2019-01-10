package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zcong1993/libgo/gin/ginerr"
	"github.com/zcong1993/libgo/gin/ginhelper"
	"github.com/zcong1993/libgo/utils"
	"github.com/zcong1993/libgo/validator"
	"github.com/zcong1993/rest-go/common"
	"github.com/zcong1993/rest-go/db"
	"github.com/zcong1993/rest-go/model"
	"github.com/zcong1993/rest-go/service"
	utils2 "github.com/zcong1993/rest-go/utils"
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

	u := &model.User{}
	ginhelper.MustCopy(u, &f)

	if u.Avatar == "" {
		u.Avatar = utils2.GenerateAvatar(u.Name)
	}

	err = db.ORM.Create(u).Error
	if err != nil {
		if utils2.IsDuplicateError(err) {
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

// Login is handler for /login
// @Summary Login a user
// @Description Login a user
// @Accept  json
// @Produce  json
// @Param   user     body    common.LoginForm     true        "Login user"
// @Success 201 {object} model.Token
// @Failure 400 {object} common.ErrResp
// @Failure 500 "StatusInternalServerError"
// @Router /login [post]
func Login(c *gin.Context) ginerr.ApiError {
	var f common.LoginForm
	err := c.ShouldBindJSON(&f)
	if err != nil {
		return common.CreateInvalidErr(validator.NormalizeErr(err))
	}
	var u model.User
	err = db.ORM.First(&u, "email=?", f.Email).Error
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

func userToDataLoaderResp(keys []string, data []model.User) []*model.User {
	l := len(keys)

	tmpMap := make(map[string]model.User, l)
	resp := make([]*model.User, l)

	for _, v := range data {
		tmpMap[v.ID] = v
	}

	for i, key := range keys {
		d, ok := tmpMap[key]
		if ok {
			resp[i] = &d
		} else {
			resp[i] = nil
		}
	}

	return resp
}

func UsersBatch(ctx *gin.Context) {
	batchIds := ctx.Query("ids")
	if batchIds == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}
	ids, ok := ginhelper.NormalizeBatchQuery(batchIds)
	if !ok {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var users []model.User
	err := db.ORM.Where("id in (?)", ids).Find(&users).Error
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, userToDataLoaderResp(ids, users))
}
