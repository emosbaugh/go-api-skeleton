package inject

import (
	"github.com/replicatedcom/gin-example/services"
)

type Env struct {
	UserService services.IUserService
}
