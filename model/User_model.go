package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	FirstName,LastName,EMail,UserName,Pass string
}

func (user User) Migrate() error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.AutoMigrate(&user)
	return nil
}

func (user User) Add(firstName , lastName, eMail, userName, pass string) error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Create(&User{FirstName: firstName, LastName: lastName, EMail: eMail, UserName:userName,Pass: pass})
	return nil
}

func (user User) GetUser(where ...interface{}) (User,error){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),//for close logger because log every no user condition
	})
	if err != nil {
		fmt.Println(err)
		return user,err
	}
	db.First(&user,where...)
	return user,err
}

func (user User) GetAllUsers(where ...interface{}) ([]User,error){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	var users []User
	db.Find(&users,where...)
	return users,err
}

func (user User) Update(column string, value interface{}) error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Model(&user).Update(column,value)
	return nil
}

func (user User) UpdateMultiple(data User) error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Model(&user).Updates(data)
	return nil
}

func (user User) Delete() error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Delete(&user,user.ID)
	return nil
}