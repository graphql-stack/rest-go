package model

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"github.com/zcong1993/rest-go/utils"
)

var UserFactory = factory.NewFactory(&User{}).Attr("Name", func(args factory.Args) (i interface{}, e error) {
	return randomdata.FullName(randomdata.RandomGender), nil
}).Attr("Email", func(args factory.Args) (i interface{}, e error) {
	return randomdata.Email(), nil
}).Attr("Password", func(args factory.Args) (i interface{}, e error) {
	return randomdata.RandStringRunes(10), nil
}).Attr("Avatar", func(args factory.Args) (i interface{}, e error) {
	u := args.Instance().(*User)
	return utils.GenerateAvatar(u.Name), nil
})

var PostFactory = factory.NewFactory(new(Post)).Attr("Title", func(args factory.Args) (i interface{}, e error) {
	return randomdata.Title(randomdata.RandomGender), nil
}).Attr("Content", func(args factory.Args) (i interface{}, e error) {
	return randomdata.Paragraph(), nil
}).SubFactory("Author", UserFactory)

var CommentFactory = factory.NewFactory(new(Comment)).Attr("Content", func(args factory.Args) (i interface{}, e error) {
	return randomdata.Paragraph(), nil
}).SubFactory("Author", UserFactory).SubFactory("Post", PostFactory)
