package controllers

import (
	"go-chat/app/models"
	"go-chat/app/routes"
	"log"

	"github.com/revel/revel"
)

// Admin --
type Admin struct {
	*revel.Controller
}

// Index -- signin page
func (c Admin) Index() revel.Result {
	c.ViewArgs["title"] = "Go Chat"
	return c.Render()
}

// Signin -- signin
func (c Admin) Signin() revel.Result {
	return c.Redirect(c.Show())
}

// Signup -- signup
func (c Admin) Signup() revel.Result {
	return c.Render()
}

// Create -- create user admin
func (c Admin) Create(userAdmin models.UserAdmin, confirmPassword string) revel.Result {
	log.Printf("%s", userAdmin.Name)
	log.Printf("%v", c.Params.Form)
	c.FlashParams()

	return c.Redirect(routes.Admin.Signup())
}

// Show --
func (c Admin) Show() revel.Result {
	return c.Render()
}

// Delete --
func (c Admin) Delete(id int) revel.Result {
	return c.Render()
}
