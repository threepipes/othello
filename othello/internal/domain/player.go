package domain

type Player int8

const (
	PlayerUnknown Player = iota
	PlayerBlack
	PlayerWhite
)

func (p Player) Disk() Disk {
	switch p {
	case PlayerBlack:
		return DiskBlack
	case PlayerWhite:
		return DiskWhite
	default:
		return DiskUnknown
	}
}
