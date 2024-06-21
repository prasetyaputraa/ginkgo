package context

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	params map[string]interface{}
	gin.Context
}
