package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "aloha"
	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	// 空を許さない
	c.Validation.Required(myName).Message("Your name is required")
	// 最小の長さ
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	// errrorの時の分岐
	if c.Validation.HasErrors() {
		// これはいったい
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}
