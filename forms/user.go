package forms

// UserSignupForm - Sign Up form
type UserSignupForm struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UserLoginForm - Sign Up form
type UserLoginForm struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}
