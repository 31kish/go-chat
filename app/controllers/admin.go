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

// Admin -
type Admin struct {
	*revel.Controller
}

func (c Admin) before() revel.Result {
	c.ViewArgs["title"] = "Go Chat"
	return c.Result
}

func (c Admin) loggedIn() bool {
	_, contains := c.Session["user_admin_id"]
	return contains
}

func init() {
	revel.InterceptMethod(Admin.before, revel.BEFORE)
}

// Show - Top page
func (c Admin) Show() revel.Result {
	if !c.loggedIn() {
		return c.Redirect(routes.Admin.Index())
	}

	s := services.UserAdmin{}
	userAdmins, err := s.GetAll()

	if err != nil {
		c.Flash.Error("%s", err)
		return c.Redirect(routes.Admin.Show())
	}

	c.ViewArgs["name"] = c.Session["user_admin_name"]
	c.ViewArgs["id"] = c.Session["user_admin_id"]
	return c.Render(userAdmins)
}

// Index - signin page
func (c Admin) Index() revel.Result {
	if c.loggedIn() {
		return c.Redirect(routes.Admin.Show())
	}

	return c.Render()
}

// Signin - signin action
func (c Admin) Signin(mailAdress string, password string) revel.Result {
	s := services.UserAdmin{}

	user, err := s.Get(mailAdress, password)
	if err != nil {
		c.Flash.Error("%s", err)
	}

	c.Session["user_admin_id"] = fmt.Sprint(user.ID)
	c.Session["user_admin_name"] = user.Name
	return c.Redirect(routes.Admin.Show())
}

// Signup - signup page
func (c Admin) Signup() revel.Result {
	return c.Render()
}

// Create - signup action
func (c Admin) Create(userAdmin models.UserAdmin, verifyPassword string) revel.Result {
	userAdmin.HashedPassword, _ = utils.EncryptPassword(userAdmin.Password)

	userAdmin.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Admin.Signup())
	}

	// insert
	// ex. service.UserAdmin{}.Save(interface)
	s := services.UserAdmin{}
	user, err := s.Create(userAdmin)
	if err != nil {
		c.Flash.Error("%s", err)
		return c.Redirect(routes.Admin.Signup())
	}

	// session
	c.Session["user_admin_id"] = fmt.Sprint(user.ID)
	c.Session["user_admin_name"] = user.Name
	return c.Redirect(routes.Admin.Index())
}

// Delete - user delete action
func (c Admin) Delete(id int) revel.Result {
	s := services.UserAdmin{}
	err := s.Delete(id)

	if err != nil {
		c.Response.Status = http.StatusInternalServerError
		text := err.Error()
		return c.RenderText(text)
	}

	return c.RenderText("success")
}

// Update - user update action
func (c Admin) Update(id int, name string, mailAdress string) revel.Result {
	s := services.UserAdmin{}

	err := s.Update(id, name, mailAdress)

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		text := err.Error()
		return c.RenderText(text)
	}

	return c.Render()
}

// Signout - sign out
func (c Admin) Signout(id string) revel.Result {
	delete(c.Session, "user_admin_id")
	delete(c.Session, "user_admin_name")
	return c.Redirect(routes.Admin.Index())
}
