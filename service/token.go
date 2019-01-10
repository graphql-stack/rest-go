package service

import (
	"github.com/zcong1993/rest-go/common"
	"github.com/zcong1993/rest-go/db"
	"github.com/zcong1993/rest-go/model"
	"github.com/zcong1993/rest-go/utils"
)

func RefreshToken(token model.Token) {
	db.ORM.Model(&token).Update("token", utils.GenerateToken())
}

func GetOrCreateToken(user model.User) (*model.Token, error) {
	var token model.Token
	err := db.ORM.Where(model.Token{UserID: user.ID}).Attrs(model.Token{Token: utils.GenerateToken()}).FirstOrCreate(&token).Error
	if token.IsExpired() {
		RefreshToken(token)
	}
	return &token, err
}

func GetUserByToken(tk string) (*model.User, error) {
	var token model.Token
	err := db.ORM.Preload("User").Where("token = ?", tk).First(&token).Error
	if err != nil {
		return nil, err
	}
	if token.IsExpired() {
		return nil, common.ErrExpired
	}
	return token.User, err
}
