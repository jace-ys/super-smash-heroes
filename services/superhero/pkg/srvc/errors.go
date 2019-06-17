package srvc

import (
	"fmt"
)

var (
	errInternalServerError   = fmt.Errorf("Internal server error")
	errNoSuperheroes         = fmt.Errorf("No superheroes could be found")
	errSuperheroDoesNotExist = fmt.Errorf("Requested superhero does not exist")
)
