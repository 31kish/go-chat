package controllers

import "github.com/revel/revel"

// App -
type App struct {
	*revel.Controller
}

// Index -
func (c App) Index() revel.Result {
	return c.Render()
}

// Signup -
func (c App) Signup() revel.Result {
	return c.Render()
}

// Create -
func (c App) Create() revel.Result {
	return c.Render()
}
