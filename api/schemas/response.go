package schemas

import (
	"net/http"
	"os"
	"strconv"

	"github.com/replicatedcom/gin-example/api/errors"

	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/gin-gonic/gin.v1"
)

// MaybeValidateResponse conditionally validates the go resource against the
// provided schema if the "VALIDATE_RESPONSE" environment variable is truthy.
func MaybeValidateResponse(c *gin.Context, obj interface{}, schema *gojsonschema.Schema) error {
	ok, _ := strconv.ParseBool(os.Getenv("VALIDATE_RESPONSE"))
	if !ok {
		return nil
	}
	result, err := Validate(schema, gojsonschema.NewGoLoader(obj))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return err
	}
	if result != nil {
		err := &errors.Err{
			Status:  http.StatusInternalServerError,
			Code:    "resource.invalid",
			Message: result.Error(),
			Args: map[string]interface{}{
				"errors": result,
			},
		}
		err.Response(c)
		return err
	}
	return nil
}
