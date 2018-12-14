package model

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"github.com/zcong1993/rest-go/utils"
)

var (
	UserFactory = factory.NewFactory(&User{}).Attr("Username", func(args factory.Args) (i interface{}, e error) {
		return randomdata.FullName(randomdata.RandomGender), nil
	}).Attr("Email", func(args factory.Args) (i interface{}, e error) {
		return randomdata.Email(), nil
	}).Attr("Password", func(args factory.Args) (i interface{}, e error) {
		return randomdata.RandStringRunes(10), nil
	})

	TokenFactory = factory.NewFactory(&Token{}).Attr("Token", func(args factory.Args) (i interface{}, e error) {
		return utils.GenerateToken(), nil
	}).SubFactory("User", UserFactory)
)
