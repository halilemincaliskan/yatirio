package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	UserId int
	TlBalance,DollarBalance,BitcoinBalance,EthereumBalance,DogeBalance float32
}

func (wallet Wallet) Migrate() error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.AutoMigrate(&wallet)
	return nil
}

func (wallet Wallet) Add(userId int, tlbalance float32, dollarbalance float32, bitcoinbalance float32, ethereumbalance float32, dogebalance float32) error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Create(&Wallet{UserId: userId, TlBalance: tlbalance, DollarBalance: dollarbalance, BitcoinBalance: bitcoinbalance, EthereumBalance: ethereumbalance, DogeBalance: dogebalance})
	return nil
}

func (wallet Wallet) GetWallet(where ...interface{}) (Wallet,error){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return wallet,err
	}
	db.First(&wallet,where...)
	return wallet,err
}

func (wallet Wallet) GetAllWallets(where ...interface{}) ([]Wallet,error){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	var wallets []Wallet
	db.Find(&wallets,where...)
	return wallets,err
}

func (wallet Wallet) Update(column string, value interface{}) error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Model(&wallet).Update(column,value)
	return nil
}

func (wallet Wallet) UpdateMultiple(data Wallet) error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Model(&wallet).Updates(data)
	return nil
}

func (wallet Wallet) Delete() error{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.Delete(&wallet,wallet.ID)
	return nil
}