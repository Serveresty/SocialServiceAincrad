package cerr

import "errors"

var ErrAlreadyAuthorized = errors.New("Already authorized")
var ErrHashError = errors.New("Password is not equal hash")
var ErrUnauthorized = errors.New("Unauthorized")
var ErrorClaims = errors.New("Error while getting claims from token")
var ErrNoAccessByPrivacy = errors.New("No access by privacy")
