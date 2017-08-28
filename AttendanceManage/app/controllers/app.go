package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Input() revel.Result {
	return c.Render()
}

func (c App) TimeCard() revel.Result {
	return c.Render()
}
