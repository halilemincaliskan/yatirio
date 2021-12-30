package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/halilemincaliskan/yatirio/model"
	"time"
)

var store_user = session.New(session.Config{
	KeyLookup:   "cookie:user",
	Expiration:   3 * time.Hour,
})

func SetUser(c *fiber.Ctx,username string,password string) error {
	sess,err := store_user.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Set("username", username)
	sess.Set("password", password)

	return sess.Save()
}

func GetUser(c *fiber.Ctx) interface{} {
	sess,err := store_user.Get(c)
	if err != nil {
		panic(err)
	}
	username := sess.Get("username")
	return username
}

func CheckUser(c *fiber.Ctx) bool {
	sess,err := store_user.Get(c)
	if err != nil {
		return false
	}
	username := sess.Get("username")
	password := sess.Get("password")

	user, err := model.User{}.GetUser("user_name = ? AND pass = ?",username,password)
	if user.UserName == username && user.Pass == password{
		return true
	}
	SetAlert(c,"Lütfen Giriş Yapın")
	c.Redirect("/login")
	return false
}