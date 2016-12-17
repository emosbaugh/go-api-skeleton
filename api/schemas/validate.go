package schemas

import (
	"net/http"

	"github.com/replicatedcom/gin-example/api/errors"

	"github.com/xeipuuv/gojsonschema"
)

// Validate will validate the document for the provided schema. It will return
// a slice of api.Err for all failed schema validations.
func Validate(schema *gojsonschema.Schema, document gojsonschema.JSONLoader) (*errors.Errs, error) {
	result, err := schema.Validate(document)
	if err != nil {
		return nil, err
	}
	if result.Valid() {
		return nil, nil
	}
	var errs errors.Errs
	for _, resultErr := range result.Errors() {
		errs = append(errs, resultErrorToError(resultErr))
	}
	return &errs, nil
}

func resultErrorToError(resultErr gojsonschema.ResultError) *errors.Err {
	return errors.New(
		http.StatusBadRequest,
		getResultErrorCode(resultErr.Type()),
		resultErr.Description(),
		getResultErrorArgs(resultErr.Details()),
	)
}

func getResultErrorCode(errorType string) string {
	if errorType == "" {
		return ""
	}
	return "schema." + errorType
}

func getResultErrorArgs(details gojsonschema.ErrorDetails) map[string]interface{} {
	args := map[string]interface{}{}
	for key, value := range details {
		args[key] = value
	}
	return args
}
