package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/libgo/gin/ginerr"
	"github.com/zcong1993/libgo/validator"
	"github.com/zcong1993/rest-go/common"
	"github.com/zcong1993/rest-go/db"
	"github.com/zcong1993/rest-go/middleware"
	"github.com/zcong1993/rest-go/model"
	"github.com/zcong1993/rest-go/utils"
	"net/http"
)

// CreateComment is handler for /posts
// @Summary Create a new comment
// @Description Create a new comment
// @Accept  json
// @Produce  json
// @Param   post     body    common.CommentInput     true        "Post comment"
// @Success 201 {object} model.Comment
// @Failure 400 {object} common.ErrResp
// @Failure 500 "StatusInternalServerError"
// @Security ApiKeyAuth
// @Router /comments [post]
func CreateComment(ctx *gin.Context) ginerr.ApiError {
	var f common.CommentInput
	if err := ctx.ShouldBindJSON(&f); err != nil {
		return common.CreateInvalidErr(validator.NormalizeErr(err))
	}

	user := middleware.GetCurrentUser(ctx)

	var post model.Post
	err := db.ORM.Where("id=?", f.PostID).First(&post).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return common.NOT_FOUND_ERROR
		}
		return common.INTERVAL_ERROR
	}

	comment := &model.Comment{
		Content: f.Content,
		Author:  user,
		Post:    &post,
	}

	err = db.ORM.Create(comment).Error
	if err != nil {
		return common.INTERVAL_ERROR
	}
	ctx.JSON(http.StatusCreated, comment)
	return nil
}

// GetPostComments is handler for /posts
// @Summary Get a post's comments
// @Description Get a post's comments
// @Accept  json
// @Produce  json
// @Param   id  path  string  true  "Post id"
// @Success 200 {array} model.Comment
// @Failure 400 {object} common.ErrResp
// @Failure 500 "StatusInternalServerError"
// @Security ApiKeyAuth
// @Router /posts/{id}/comments [get]
func GetPostComments(ctx *gin.Context) ginerr.ApiError {
	id := ctx.Param("id")
	if !utils.IsUuid4(id) {
		return common.NOT_UUID
	}

	var post model.Post
	err := db.ORM.Where("id=?", id).First(&post).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return common.NOT_FOUND_ERROR
		}
		return common.INTERVAL_ERROR
	}

	var comments []model.Comment
	err = db.ORM.Model(&post).Related(&comments).Error
	if err != nil {
		return common.INTERVAL_ERROR
	}

	ctx.JSON(http.StatusOK, comments)
	return nil
}
