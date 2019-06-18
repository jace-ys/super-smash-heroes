package errors

import (
	"fmt"
)

var (
	InternalServerError = fmt.Errorf("Internal server error")
	MissingAccessToken  = fmt.Errorf("No access token found")
	InvalidRequest      = fmt.Errorf("Invalid request format")

	MissingID             = fmt.Errorf("Missing superhero ID")
	InvalidIDFormat       = fmt.Errorf("Provided ID is of invalid format")
	SuperheroesNotFound   = fmt.Errorf("No superheroes could be found")
	SuperheroNotFound     = fmt.Errorf("Requested superhero could not be found")
	SuperheroDoesNotExist = fmt.Errorf("Requested superhero does not exist")
	SuperheroExists       = fmt.Errorf("Requested superhero already added")

	UndeterminedWinner = fmt.Errorf("Battle winner could not be determined")
)
