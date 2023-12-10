package othelloerror

import (
	"fmt"
)

type ErrorInvalidPlaceForDisk string

func (e ErrorInvalidPlaceForDisk) Error() string {
	return fmt.Sprintf("invalid place to put the disk: %s", string(e))
}

var (
	ErrInvalidPlaceForDisk = ErrorInvalidPlaceForDisk("")
)
