package controllers

import (
	"go-chat/app/models"
	"go-chat/app/routes"
	"go-chat/app/services"
	"go-chat/app/utils"

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

	userAdmin, err := s.GetUserAdmin(mailAdress, password)
	if err != nil {
		c.Flash.Error("%s", err)
	}

	c.Session["user_admin_id"] = string(userAdmin.(models.UserAdmin).ID)
	c.Session["user_admin_name"] = userAdmin.(models.UserAdmin).Name
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
	r, err := s.Save(userAdmin)
	if err != nil {
		c.Flash.Error("%s", err)
		return c.Redirect(routes.Admin.Signup())
	}

	// session
	c.Session["user_admin_id"] = string(r.(*models.UserAdmin).ID)
	c.Session["user_admin_name"] = r.(*models.UserAdmin).Name

	return c.Redirect(routes.Admin.Index())
}

// Show - admin top page
func (c Admin) Show() revel.Result {
	if !c.loggedIn() {
		return c.Redirect(routes.Admin.Index())
	}

	return c.Render()
}

// Delete - user delete action
func (c Admin) Delete(id int) revel.Result {
	return c.Render()
}
