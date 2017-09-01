package controllers

import (
	"fmt"
	"go-chat/app/models"
	"go-chat/app/routes"
	"go-chat/app/services"
	"go-chat/app/utils"
	"net/http"
	"strings"

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

	admin := services.UserAdmin{}
	userAdmins, err := admin.GetAll()

	if err != nil {
		c.Flash.Data["error_user_admins"] = err.Error()
	}

	user := services.User{}
	users, err := user.GetAll()

	if err != nil {
		c.Flash.Data["error_users"] = err.Error()
	}

	c.ViewArgs["name"] = c.Session["user_admin_name"]
	c.ViewArgs["id"] = c.Session["user_admin_id"]
	return c.Render(userAdmins, users)
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
		return c.Redirect(routes.Admin.Index())
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
	m := models.UserAdmin{Name: name, MailAdress: mailAdress}
	m.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Response.Status = http.StatusUnprocessableEntity
		var msg string
		for _, err := range c.Validation.Errors {
			k := strings.Split(err.Key, ".")[1]
			msg = fmt.Sprintf("%sは%s", k, err.Message)
		}
		return c.RenderText(msg)
	}

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
