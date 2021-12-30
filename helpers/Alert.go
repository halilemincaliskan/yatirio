package helpers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store_alert = session.New(session.Config{
	KeyLookup:   "cookie:alert",
})

func SetAlert(c *fiber.Ctx,message string) error {
	sess,err := store_alert.Get(c)
	if err != nil{
		fmt.Println(err)
		return err
	}
	sess.Set("message",message)

	return sess.Save()
}

func GetAlert(c *fiber.Ctx) map[string]interface{}{
	sess,err := store_alert.Get(c)
	if err != nil{
		fmt.Println(err)
		return nil
	}

	data := make(map[string]interface{})
	message := sess.Get("message")
	if message != nil{
		data["isAlert"] = true
		data["message"] = message
	} else {
		data["isAlert"] = false
		data["message"] = nil
	}
	sess.Save()
	return data
}
