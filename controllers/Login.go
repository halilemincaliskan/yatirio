package controllers

import (
	"crypto/sha256"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/halilemincaliskan/yatirio/helpers"
	"github.com/halilemincaliskan/yatirio/model"
	"log"
	"time"
)

var store = session.New(session.Config{
	Expiration:   1 * time.Minute,
})

type Login struct {}

func (login Login) Index(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}
	is_logged := sess.Get("isLogged")
	if is_logged == 1 {
		return c.Redirect("/mainpage")
	}
	return c.Render("login/index", fiber.Map{
		"Alert": helpers.GetAlert(c),
	})
}

func (login Login) Check(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := fmt.Sprintf("%x",sha256.Sum256([]byte(c.FormValue("password"))))
	user, err := model.User{}.GetUser("user_name = ? AND pass = ?",username,password)
	if err != nil{
		return err
	}
	if user.UserName == username && user.Pass == password {
		err := helpers.SetUser(c,username,password)
		if err != nil {
			log.Fatal(err)
		}
		helpers.SetAlert(c,"Başarıyla Giriş Yaptınız")
		return c.Redirect("/mainpage")
	} else{
		helpers.SetAlert(c,"Kullanıcı Adınız veya Şifreniz Yanlış Lütfen Tekrar Deneyiniz")
		return c.Redirect("/login")
	}
}
func (login Login) Logout(c *fiber.Ctx) error {
	sess,err := store.Get(c)
	if err != nil{
		return err
	}
	sess.Delete("username")
	sess.Delete("password")
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = sess.Save()
	if err != nil {
		fmt.Println(err)
		return err
	}
	helpers.SetAlert(c,"Başarıyla Çıkış Yaptınız")
	return c.Redirect("/login")
}