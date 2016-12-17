package schemas

import (
	"net/http"

	"github.com/replicatedcom/gin-example/api/binding"

	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/gin-gonic/gin.v1"
)

// Bind parses the request's body as JSON and validates it with the provided schema.
func Bind(c *gin.Context, obj interface{}, schema *gojsonschema.Schema) error {
	err := binding.Bind(c, obj)
	if err != nil {
		return err
	}
	result, err := Validate(schema, gojsonschema.NewGoLoader(obj))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return err
	}
	if result != nil {
		result.Response(c)
		return result
	}
	return nil
}
