package params

// UserCreate contains parameters to the user create route.
type UserCreate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
