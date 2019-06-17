package srvc

import (
	"fmt"
)

var (
	errInternalServerError   = fmt.Errorf("Internal server error")
	errSuperheroDoesNotExist = fmt.Errorf("Requested superhero does not exist")
)
