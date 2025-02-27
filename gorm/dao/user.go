package dao

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID         int64
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime"`
	Admin      bool   `gorm:"-"`
}

type UserV2 struct {
	gorm.Model
	Name string
}

// 等效于
type UserV3 struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

// static table name
func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}

// dynamic table name
func UserTable(user *User) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if user.Admin {
			return tx.Table("admin_users")
		}

		return tx.Table("users")
	}
}

func Save(user *User) {
	//err := DB.Create(user).Error
	user.Admin = true
	//err := DB.Scopes(UserTable(user)).Create(user).Error
	err := DB.Table("users").Create(user).Error
	if err != nil {
		log.Println("insert fail : ", err)
	}
}

// CRUD
func GetById(id int64) User {
	var user User
	err := DB.Where("id=?", id).First(&user).Error
	if err != nil {
		log.Println("get user by id fail : ", err)
	}
	return user
}

func UpdateById(id int64) {
	err := DB.Model(&User{}).Where("id=?", id).
		Update("username", "refaw").Error
	if err != nil {
		log.Println("update users  fail : ", err)
	}
}

func DeleteById(id int64) {
	err := DB.Where("id=?", id).Delete(&User{}).Error
	if err != nil {
		log.Println("delete users  fail : ", err)
	}
}
