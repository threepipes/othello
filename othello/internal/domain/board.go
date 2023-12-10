package domain

import "fmt"

const BoardSize = 8

type Board [BoardSize][BoardSize]Disk

func NewBoard() *Board {
	b := Board{}
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			b[i][j] = DiskNone
		}
	}
	return &b
}

type Disk int8

const (
	DiskUnknown Disk = iota
	DiskBlack
	DiskWhite
	DiskNone
)

func (d Disk) OppositeColer() (Disk, error) {
	switch d {
	case DiskBlack:
		return DiskWhite, nil
	case DiskWhite:
		return DiskBlack, nil
	}
	return DiskUnknown, fmt.Errorf("cannot define the opposite disk: %v", d)
}
