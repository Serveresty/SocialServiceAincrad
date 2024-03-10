package cerr

import "errors"

var ErrAlreadyAuthorized = errors.New("already authorized")
var ErrHashError = errors.New("password is not equal hash")
var ErrUnauthorized = errors.New("unauthorized")
var ErrorClaims = errors.New("eError while getting claims from token")
var ErrNoAccessByPrivacy = errors.New("no access by privacy")
var ErrBadRequest = errors.New("wrong request")
