package banana

import "errors"

var (
	UserConflict = errors.New("User already exist")
	SignUpFail = errors.New("Sign up fail")
)
