package service

import (
	"github.com/zcong1993/rest-go/common"
	"github.com/zcong1993/rest-go/model"
	"github.com/zcong1993/rest-go/mysql"
	"github.com/zcong1993/rest-go/utils"
)

func GetOrCreateToken(user model.User) (*model.Token, error) {
	var token model.Token
	err := mysql.DB.Where(model.Token{UserID: user.ID}).Attrs(model.Token{Token: utils.GenerateToken()}).FirstOrCreate(&token).Error
	if token.IsExpired() {
		err = token.Refresh()
	}
	return &token, err
}

func GetUserByToken(tk string) (*model.User, error) {
	var token model.Token
	err := mysql.DB.Preload("User").Where("token = ?", tk).First(&token).Error
	if err != nil {
		return nil, err
	}
	if token.IsExpired() {
		return nil, common.ErrExpired
	}
	return token.User, err
}
