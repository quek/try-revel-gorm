package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/myaccount/my-app/app/models"
)

type App struct {
	GormController
}

func (c App) Index() revel.Result {
	fmt.Println(c.Txn)

	users := []models.User{}
	c.Txn.Find(&users)
	fmt.Println(users)

	return c.Render()
}
