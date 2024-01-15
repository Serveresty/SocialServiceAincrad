package cerr

import "errors"

var AlreadyAuthorized = errors.New("Already authorized")
var HashError = errors.New("Password is not equal hash")
