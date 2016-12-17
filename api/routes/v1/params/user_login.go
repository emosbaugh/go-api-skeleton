package params

// UserLogin contains parameters to the user login route.
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
