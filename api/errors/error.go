package errors

import (
	multierror "github.com/hashicorp/go-multierror"
	"gopkg.in/gin-gonic/gin.v1"
)

type Err struct {
	Status  int                    `json:"-"`
	Code    string                 `json:"code"`
	Message string                 `json:"error,omitempty"`
	Args    map[string]interface{} `json:"args,omitempty"`
}

func New(status int, code string, msg string, args map[string]interface{}) *Err {
	return &Err{status, code, msg, args}
}

func (e *Err) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}

func (e *Err) Response(c *gin.Context) {
	if e == nil {
		return
	}
	c.Error(&gin.Error{
		Err:  e,
		Type: gin.ErrorTypePublic,
		Meta: *e,
	})
	c.Status(e.Status)
}

type Errs []*Err

func (e Errs) Error() string {
	if len(e) == 0 {
		return ""
	}
	return e.ToMultiError().Error()
}

func (e Errs) Response(c *gin.Context) {
	for _, ee := range e {
		ee.Response(c)
	}
}

func (e Errs) ToMultiError() *multierror.Error {
	if len(e) == 0 {
		return nil
	}
	var errs []error
	for _, err := range e {
		errs = append(errs, err)
	}
	return multierror.Append(errs[0], errs[1:]...)
}
