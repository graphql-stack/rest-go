package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/libgo/gin/ginerr"
	"github.com/zcong1993/libgo/gin/ginhelper"
	"github.com/zcong1993/libgo/validator"
	"github.com/zcong1993/rest-go/common"
	"github.com/zcong1993/rest-go/db"
	"github.com/zcong1993/rest-go/middleware"
	"github.com/zcong1993/rest-go/model"
	"net/http"
)

type PostView struct {
	ginhelper.Rest
	ginhelper.RestView
}

func (b *PostView) GetQuerySet() *gorm.DB {
	return db.ORM.Model(new(model.Post))
}

func (b *PostView) GetModel(isMany bool) interface{} {
	if isMany {
		var book []model.Post
		return &book
	}
	return &model.Post{}
}

// Posts query all for /posts
// @Summary Posts pagination query all
// @Description
// @Accept  json
// @Produce  json
// @Param   default limit query int false "limit" default(100)
// @Param   default offset query int false "offset" default(0)
// @Success 200 {array} model.Post
// @Failure 500 "StatusInternalServerError"
// @Router /posts [get]
func DocsPostList() {

}

// Posts query by id for /posts/:id
// @Summary Posts query by id
// @Description
// @Accept  json
// @Produce  json
// @Param   id  path  string  true  "Post id"
// @Success 200 {object} model.Post
// @Failure 500 "StatusInternalServerError"
// @Router /posts/{id} [get]
func DocsPostRetrieve() {

}

// CreatePost is handler for /posts
// @Summary Register a new user
// @Description Register a new user
// @Accept  json
// @Produce  json
// @Param   post     body    common.PostInput     true        "Post content"
// @Success 201 {object} model.Post
// @Failure 400 {object} common.ErrResp
// @Failure 500 "StatusInternalServerError"
// @Security ApiKeyAuth
// @Router /posts [post]
func CreatePost(ctx *gin.Context) ginerr.ApiError {
	var f common.PostInput
	if err := ctx.ShouldBindJSON(&f); err != nil {
		return common.CreateInvalidErr(validator.NormalizeErr(err))
	}

	user := middleware.GetCurrentUser(ctx)

	post := &model.Post{
		Title:   f.Title,
		Content: f.Content,
		Author:  user,
	}

	err := db.ORM.Create(post).Error

	if err != nil {
		return common.INTERVAL_ERROR
	}

	ctx.JSON(http.StatusCreated, post)

	return nil
}
