package controllers

import (
	"fmt"
	"go-chat/app/models"
	"go-chat/app/routes"
	"go-chat/app/services"
	"go-chat/app/utils"
	"net/http"

	"github.com/revel/revel"
)

// App -
type App struct {
	*revel.Controller
}

func init() {
	revel.InterceptMethod(App.before, revel.BEFORE)
}

func (c App) before() revel.Result {
	c.ViewArgs["title"] = "Go Chat"
	return c.Result
}

func (c App) loggedIn() bool {
	_, contains := c.Session["user_id"]
	return contains
}

// Index - User top page.
func (c App) Index() revel.Result {
	if !c.loggedIn() {
		return c.Redirect(routes.App.Signin())
	}

	c.ViewArgs["id"] = c.Session["user_id"]
	c.ViewArgs["name"] = c.Session["user_name"]
	return c.Render()
}

// Signin - User Sign in page.
func (c App) Signin() revel.Result {
	if c.loggedIn() {
		return c.Redirect(routes.App.Index())
	}

	return c.Render()
}

// Auth - User Sign in action.
func (c App) Auth(mailAdress string, password string) revel.Result {
	s := services.User{}

	user, err := s.Get(mailAdress, password)

	if err != nil {
		c.Flash.Error("%s", err)
		return c.Redirect(routes.App.Signin())
	}

	c.Session["user_id"] = fmt.Sprint(user.ID)
	c.Session["user_name"] = user.Name
	return c.Redirect(routes.App.Index())
}

// Signup - User Sign up page.
func (c App) Signup() revel.Result {
	return c.Render()
}

// Create - User create action
func (c App) Create(user models.User, verifyPassword string) revel.Result {
	user.HashedPassword, _ = utils.EncryptPassword(user.Password)

	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	s := services.User{}

	u, err := s.Create(user)

	if err != nil {
		c.Flash.Error("%s", err)
		return c.Redirect(routes.App.Signup())
	}

	c.Session["user_id"] = fmt.Sprint(u.ID)
	c.Session["user_name"] = u.Name

	return c.Redirect(routes.App.Index())
}

// Update - User update action
func (c App) Update(id int, name string, mailAdress string) revel.Result {
	s := services.User{}

	err := s.Update(id, name, mailAdress)

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		text := err.Error()
		return c.RenderText(text)
	}
	return c.Render()
}

// Delete - User delete action.
func (c App) Delete(id int) revel.Result {
	s := services.User{}

	err := s.Delete(id)

	if err != nil {
		c.Response.Status = http.StatusInternalServerError
		text := err.Error()
		return c.RenderText(text)
	}

	return c.RenderText("success")
}

// Signout - User Signout
func (c App) Signout() revel.Result {
	delete(c.Session, "user_id")
	delete(c.Session, "user_name")
	return c.Redirect(routes.App.Signin())
}
