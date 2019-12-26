package api

import (
	"net/http"
)

type pingController struct {
	controllerBase
}

func (c *pingController) Get() {
	c.Context.JSON(http.StatusOK, "pong from ping controller")
}
