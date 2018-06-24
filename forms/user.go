package forms

// UserSignupForm - Sign Up form
type UserSignupForm struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
