package controllers

import (
	"crypto/sha256"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/halilemincaliskan/yatirio/helpers"
	"github.com/halilemincaliskan/yatirio/model"
)

type Register struct {}

func (register Register) Index(c *fiber.Ctx) error {
	return c.Render("register/index", fiber.Map{})
}

func (register Register) SignUp(c *fiber.Ctx) error{
	firstName := c.FormValue("firstname")
	lastName := c.FormValue("lastname")
	userName := c.FormValue("username")
	email := c.FormValue("email")
	password := fmt.Sprintf("%x",sha256.Sum256([]byte(c.FormValue("password"))))
	repassword := fmt.Sprintf("%x",sha256.Sum256([]byte(c.FormValue("repassword"))))
	user, err := model.User{}.GetUser("user_name = ? OR e_mail = ?",userName,email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if user.UserName != "" {
		helpers.SetAlert(c,"Bu Kullanıcı Adı veya Mail Zaten Kullanılmakta Lütfen Tekrar Deneyiniz")
		return c.Redirect("/register")
	}
	if password == repassword {
		model.User{}.Add(firstName,lastName,email,userName,password)
		user,err := model.User{}.GetUser("user_name = ?",userName)
		if err != nil {
			fmt.Println(err)
			return err
		}
		model.Wallet{}.Add(int(user.ID),1000,0,0,0,0)
		helpers.SetAlert(c,"Başarıyla Kayıt Oldunuz Şimdi Giriş Yapabilirsiniz")
		return c.Redirect("/login")
	}
	fmt.Println("bir hata oldu")
	return c.Redirect("/register")
}