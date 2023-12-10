package othelloerror

import "errors"

var (
	ErrInvalidPlaceForDisk = errors.New("invalid place to put the disk")
)
