package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Rate struct {
	gorm.Model
	DollarRate,BitcoinRate,EthereumRate,DogeRate float32
}

func (rate Rate) Migrate() error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.AutoMigrate(&rate)
	return nil
}

func (rate Rate) Add(dollarrate float32, bitcoinrate float32, ethereumrate float32, dogerate float32) error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Create(&Rate{DollarRate: dollarrate, BitcoinRate: bitcoinrate, EthereumRate: ethereumrate, DogeRate: dogerate})
	return nil
}

func (rate Rate) GetRate(where ...interface{}) (Rate,error){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return rate,err
	}
	db.First(&rate,where...)
	return rate,err
}

func (rate Rate) GetAllRates(where ...interface{}) ([]Rate,error){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	var rates []Rate
	db.Find(&rates,where...)
	return rates,err
}

func (rate Rate) Update(column string, value interface{}) error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Model(&rate).Update(column,value)
	return nil
}

func (rate Rate) UpdateMultiple(data Wallet) error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Model(&rate).Updates(data)
	return nil
}

func (rate Rate) Delete() error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Delete(&rate,rate.ID)
	return nil
}