package controllers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/mrkaspa/socialdonor/app/models"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

// type: revel controller with `*gorm.DB`
// c.Txn will keep `Gdb *gorm.DB`
type GormController struct {
	*revel.Controller
	Txn *gorm.DB
}

var BloodTypes = []string{"O+", "O-", "A+", "A-", "B+", "B-", "AB+", "AB-"}

// SetUser loads the session user
func (c *GormController) SetUser() revel.Result {
	revel.INFO.Println("Entro a setUser!!!")
	user := models.User{}
	if _, ok := c.Session["uid"]; ok {
		uid, _ := strconv.ParseInt(c.Session["uid"], 10, 0)
		c.Txn.First(&user, uid)
	}
	c.RenderArgs["user"] = &user
	return nil
}

func (c *GormController) Connected() *models.User {
	return c.RenderArgs["user"].(*models.User)
}

func (c GormController) LoadBloodTypes() {
	c.RenderArgs["bloodTypes"] = BloodTypes
}

// transactions

// This method fills the c.Txn before each transaction
func (c *GormController) Begin() revel.Result {
	fmt.Println("Entro a BEGIN!!!")
	txn := models.Gdb.Begin()
	if txn.Error != nil {
		fmt.Println("Unable to open a connection")
		fmt.Println(txn.Error)
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
