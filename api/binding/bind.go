package binding

import (
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/gin-gonic/gin.v1"
)

// Bind offers binding only for JSON. This plus JWT implementation offers CSRF
// protection.
// TODO: provide link to docs on this
func Bind(c *gin.Context, obj interface{}) error {
	return c.BindWith(obj, binding.JSON)
}
