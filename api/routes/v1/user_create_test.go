package v1_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	. "github.com/replicatedcom/gin-example/api/routes/v1"
	"github.com/replicatedcom/gin-example/api/routes/v1/params"
	"github.com/replicatedcom/gin-example/api/routes/v1/resource"
	"github.com/replicatedcom/gin-example/inject"
	"github.com/replicatedcom/gin-example/models"
	servicesmock "github.com/replicatedcom/gin-example/services/mock"

	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
	"gopkg.in/gin-gonic/gin.v1"
)

func TestUserCreate(t *testing.T) {
	mockUserService := new(servicesmock.MockUserService)
	env := &inject.Env{
		UserService: mockUserService,
	}

	router := gin.New()
	router.POST("/user", UserCreate(env))

	userEmail := "ethan@replicated.com"
	userPassword := "testing123"

	mockUserService.
		On("Create", &models.User{
			Email: userEmail,
		}, userPassword).
		Return(nil)

	w := httptest.NewRecorder()
	req, err := gorequest.New().Post("http://example.com/user").
		Send(params.UserCreate{
			Email:    userEmail,
			Password: userPassword,
		}).
		MakeRequest()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Regexp(t, regexp.MustCompile(`^Bearer (.+)`), w.Header().Get("Authorization")) // TODO

	mockUserService.AssertExpectations(t)

	var user resource.User
	err = json.Unmarshal(w.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, userEmail, user.Email)
}
