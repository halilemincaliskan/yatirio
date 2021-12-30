package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/halilemincaliskan/yatirio/helpers"
	"github.com/halilemincaliskan/yatirio/model"
	"log"
	"strconv"
)

type Mainpage struct {}

func (mainpage Mainpage) Index(c *fiber.Ctx) error {
	if !helpers.CheckUser(c){
		return nil
	}
	user,err := model.User{}.GetUser("user_name = ?",helpers.GetUser(c))
	wallet,_ := model.Wallet{}.GetWallet(user.ID)
	rate,_ := model.Rate{}.GetRate(1)
	if err != nil{
		log.Fatal(err)
	}
	userBalance := fmt.Sprintf("Tl : %f Dolar : %f Bitcoin : %f Ethereum : %f Doge : %f", wallet.TlBalance,wallet.DollarBalance,wallet.BitcoinBalance,wallet.EthereumBalance,wallet.DogeBalance)
	return c.Render("mainpage/index", fiber.Map{
		"Alert": helpers.GetAlert(c),
		"User": user.UserName,
		"Balance": userBalance,
		"TlBalance": wallet.TlBalance,
		"DollarBalance": wallet.DollarBalance,
		"BitcoinBalance": wallet.BitcoinBalance,
		"EthereumBalance": wallet.EthereumBalance,
		"DogeBalance": wallet.DogeBalance,
		"DollarRate": rate.DollarRate,
		"BitcoinRate": rate.BitcoinRate,
		"EthereumRate": rate.EthereumRate,
		"DogeRate": rate.DogeRate,
		"DollarValue": wallet.DollarBalance * rate.DollarRate,
		"BitcoinValue": wallet.BitcoinBalance * rate.BitcoinRate,
		"EthereumValue": wallet.EthereumBalance * rate.EthereumRate,
		"DogeValue": wallet.DogeBalance * rate.DogeRate,
	})
}

func (mainpage Mainpage) Islem(c *fiber.Ctx) error {
	if !helpers.CheckUser(c){
		return nil
	}
	currency := c.Params("currency")
	currencyBalance := currency + "_balance"
	process := c.Params("process")
	getValue := currency + "-" + process + "-" + "value"
	value := c.FormValue(getValue)
	getRate,_ := model.Rate{}.GetRate(1)
	var rate float32
	var newBalance float32
	var newTlBalance float32
	userId, err := model.User{}.GetUser("user_name", helpers.GetUser(c))
	user,err := model.Wallet{}.GetWallet(userId.ID)
	switch currency {
	case "dollar":
		rate = getRate.DollarRate
		if process == "buy" {
			floatValue,_ := strconv.ParseFloat(value, 32)
			newBalance = user.DollarBalance + float32(floatValue)
			tlFloatValue,_ := strconv.ParseFloat(value, 32)
			newTlBalance = user.TlBalance - (float32(tlFloatValue)*float32(rate))
		}else if process == "sell"{
			floatValue,_ := strconv.ParseFloat(value, 32)
			newBalance = user.DollarBalance - float32(floatValue)
			tlFloatValue,_ := strconv.ParseFloat(value, 32)
			newTlBalance = user.TlBalance + (float32(tlFloatValue)*float32(rate))
		}else {
			fmt.Println("Something Gone Wrong")
		}
	case "bitcoin":
		rate = getRate.BitcoinRate
		if process == "buy" {
			floatValue,_ := strconv.ParseFloat(value, 32)
			newBalance = user.BitcoinBalance + float32(floatValue)
			tlFloatValue,_ := strconv.ParseFloat(value, 32)
			newTlBalance = user.TlBalance - (float32(tlFloatValue)*rate)
		}else if process == "sell"{
			floatValue,_ := strconv.ParseFloat(value, 32)
			newBalance = user.BitcoinBalance - float32(floatValue)
			tlFloatValue,_ := strconv.ParseFloat(value, 32)
			newTlBalance = user.TlBalance + (float32(tlFloatValue)*rate)
		}else {
			fmt.Println("Something Gone Wrong")
		}
	case "ethereum":
		rate = getRate.EthereumRate
		if process == "buy" {
			floatValue,_ := strconv.ParseFloat(value, 32)
			newBalance = user.EthereumBalance + float32(floatValue)
			tlFloatValue,_ := strconv.ParseFloat(value, 32)
			newTlBalance = user.TlBalance - (float32(tlFloatValue)*rate)
		}else if process == "sell"{
			floatValue,_ := strconv.ParseFloat(value, 32)
			newBalance = user.EthereumBalance - float32(floatValue)
			tlFloatValue,_ := strconv.ParseFloat(value, 32)
			newTlBalance = user.TlBalance + (float32(tlFloatValue)*rate)
		}else {
			fmt.Println("Something Gone Wrong")
		}
	case "doge":
		rate = getRate.DogeRate
		if process == "buy" {
			floatValue,_ := strconv.ParseFloat(value, 32)
			newBalance = user.DogeBalance + float32(floatValue)
			tlFloatValue,_ := strconv.ParseFloat(value, 32)
			newTlBalance = user.TlBalance - (float32(tlFloatValue)*rate)
		}else if process == "sell"{
			floatValue,_ := strconv.ParseFloat(value, 32)
			newBalance = user.DogeBalance - float32(floatValue)
			tlFloatValue,_ := strconv.ParseFloat(value, 32)
			newTlBalance = user.TlBalance + (float32(tlFloatValue)*rate)
		}else {
			fmt.Println("Something Gone Wrong")
		}
	default:
		fmt.Println("INVALID CURRENCY")
	}
	user.Update("tl_balance", newTlBalance)
	user.Update(currencyBalance,newBalance)
	if err != nil {
		log.Fatal(err)
		return err
	}
	c.Redirect("/mainpage")
	return nil
	//c.Params("currency"), c.Params("process"))
}