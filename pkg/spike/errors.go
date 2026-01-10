package spike

import "errors"

var ErrFailedScanner = errors.New("error while running scanner")
var ErrClosingScanner = errors.New("error while closing scanner")
