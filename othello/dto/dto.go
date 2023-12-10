package othellodto

import "othello/othello/internal/domain"

type Action int8

const (
	ActionUnknown Action = iota
	ActionPutDisk
	ActionGiveUp
	ActionPass
)

func (a Action) Domain() domain.Action {
	return domain.Action(a)
}

type Disk int8

const (
	DiskUnknown Disk = iota
	DiskBlack
	DiskWhite
	DiskNone
)

func DiskFromDomain(d domain.Disk) Disk {
	return Disk(d)
}

const BoardSize = domain.BoardSize

type Board [BoardSize][BoardSize]Disk

type GameState struct {
	CurrentTurn  Disk
	Board        *Board
	GameFinished bool
	Winner       Disk // Drow の場合は DiskNone が返る
}

func BoardFromDomain(bd *domain.Board) *Board {
	b := Board{}
	for i := 0; i < domain.BoardSize; i++ {
		for j := 0; j < domain.BoardSize; j++ {
			b[i][j] = Disk(bd[i][j])
		}
	}
	return &b
}
