package cerr

import "errors"

var AlreadyAuthorized = errors.New("Already authorized")
var HashError = errors.New("Password is not equal hash")
var Unauthorized = errors.New("Unauthorized")
var ErrorClaims = errors.New("Error while getting claims from token")
