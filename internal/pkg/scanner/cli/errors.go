package cli

import "errors"

var ErrCtxTimedOut = errors.New("context deadline exceeded")
var ErrSignalInterrupt = errors.New("CTRL+C pressed: Exiting with partial output")
