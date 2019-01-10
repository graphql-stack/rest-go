package model

import (
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/libgo/utils"
	utils2 "github.com/zcong1993/rest-go/utils"
	"time"
)

var (
	// 3 days
	TOKEN_EXPIRE = time.Hour * 24 * 3
)

type Model struct {
	ID        string     `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

type User struct {
	Model
	Name     string `json:"name" gorm:"type:varchar(150);unique_index;not null"`
	Password string `json:"-" gorm:"type:varchar(150);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index;not null"`
	Avatar   string `json:"avator" gorm:"type:varchar(100)"`
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(u).Update("password", utils.HashPassword(u.Password))
	return
}

type Token struct {
	Model
	User   *User  `json:"-"`
	UserID string `json:"user_id"`
	Token  string `json:"token" gorm:"type:varchar(100);index;not null"`
}

func (t *Token) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(t).Update("token", utils2.GenerateToken())
	return
}

func (t *Token) IsExpired() bool {
	return time.Now().Sub(t.UpdatedAt) > TOKEN_EXPIRE
}

type Post struct {
	Model
	Title    string    `json:"title" gorm:"type:varchar(150);index;not null"`
	Content  string    `json:"content" gorm:"type:text"`
	Author   *User     `json:"-" gorm:"foreignkey:AuthorID"`
	AuthorID string    `json:"author_id"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Model
	Content  string `json:"content" gorm:"type:text"`
	Author   *User  `json:"-" gorm:"foreignkey:AuthorID"`
	AuthorID string `json:"author_id"`
	Post     *Post  `json:"-"`
	PostID   string `json:"post_id"`
}
