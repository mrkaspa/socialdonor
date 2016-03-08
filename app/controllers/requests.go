package controllers

import (
	"os"

	"github.com/mrkaspa/socialdonor/app/models"
	"github.com/mrkaspa/socialdonor/app/routes"
	"github.com/haisum/recaptcha"
	"github.com/revel/revel"
)

// Controller is the main controller
type Requests struct {
	SecuredController
}

// Index for request
func (c Requests) Index() revel.Result {
	c.LoadBloodTypes()
	return c.Render()
}

// Save the request for Blood
func (c Requests) Save(request models.Request) revel.Result {
	if result := c.Recaptcha(); result != nil {
		return result
	}

	userSession := (&c).Connected()
	request.Validate(c.Validation)

	// Handle errors
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		c.Flash.Error("Corrige los errores por favor")
		return c.Redirect(Requests.Index)
	}

	request.UserID = userSession.ID

	if c.Txn.Save(&request).Error != nil {
		c.FlashParams()
		c.Flash.Error("No se pudo guardar la solicitud")
		return c.Redirect(Profile.Index)
	}

	c.Flash.Success("Ya se estan buscando los donantes mas cercanos")

	return c.Redirect(routes.App.Show(request.UUID))

}

// Reacptcha validates the human user
func (c Requests) Recaptcha() revel.Result {
	re := recaptcha.R{Secret: os.Getenv("RECAPTCHA_KEY")}
	isValid := re.Verify(*c.Request.Request)
	if !isValid {
		c.FlashParams()
		c.Flash.Error("Eres un bot")
		return c.Redirect(Requests.Index)
	}
	return nil
}
