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
	const center = BoardSize/2 - 1
	b[center][center] = DiskWhite
	b[center+1][center] = DiskBlack
	b[center][center+1] = DiskBlack
	b[center+1][center+1] = DiskWhite
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
