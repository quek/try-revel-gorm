package controllers

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"github.com/myaccount/my-app/app"
)

// type: revel controller with `*gorm.DB`
// c.Txn will keep `Gdb *gorm.DB`
type GormController struct {
	*revel.Controller
	Txn *gorm.DB
}

// transactions

// This method fills the c.Txn before each transaction
func (c *GormController) Begin() revel.Result {
	txn := app.DB.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	c.Txn = txn
	return nil
}

// This method clears the c.Txn after each transaction
func (c *GormController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Commit()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

// This method clears the c.Txn after each transaction, too
func (c *GormController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Rollback()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
