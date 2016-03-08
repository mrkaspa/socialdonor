package controllers

import "github.com/revel/revel"

type SecuredController struct {
	GormController
}

func (c *SecuredController) Check() revel.Result {
	if user := c.Connected(); user.ID == 0 || c.Connected() == nil {
    c.Flash.Error("Ingresa con tu facebook")
		return c.Redirect(App.Index)
	}
	return nil
}
