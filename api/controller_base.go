package api

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type controllerBase struct {
	Context *gin.Context
}

func (c *controllerBase) get() {
	c.Context.JSON(http.StatusOK, "unhandled")
}

func controller(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		p := &pingController{}
		p.controllerBase.Context = c
		st := reflect.TypeOf(p)
		_, ok := st.MethodByName("Get")
		if !ok {
			p.controllerBase.get()
		} else {
			reflect.ValueOf(p).MethodByName("Get").Call([]reflect.Value{})
		}
	}

}
