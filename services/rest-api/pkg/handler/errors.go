package handler

import (
	"fmt"
)

var (
	errInvalidRequest = fmt.Errorf("Invalid request format")
	errMissingID      = fmt.Errorf("Missing superhero ID")
)
