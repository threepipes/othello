package usecase

import (
	"fmt"
	othellodto "othello/othello/dto"
	othelloerror "othello/othello/error"
	"othello/othello/internal/domain"
)

type Game struct {
	current domain.Disk
	board   *domain.Board
}

type point struct {
	x int
	y int
}

func (g *Game) CurrentTurn() domain.Disk {
	return g.current
}

func (g *Game) Board() domain.Board {
	return *g.board
}

func outOfBoard(x, y int) bool {
	return x < 0 || x >= domain.BoardSize || y < 0 || y >= domain.BoardSize
}

func NewPoint(x, y int) (point, error) {
	if outOfBoard(x, y) {
		return point{}, fmt.Errorf("invalid range: (%d, %d)", x, y)
	}
	return point{x, y}, nil
}

var (
	dx = []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}
	dy = []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}
)

// putDisk は disk を座標 p に石を置き、盤面を更新する
func putDisk(bd *domain.Board, d domain.Disk, p point) error {
	op, err := d.OppositeColer()
	if err != nil {
		return fmt.Errorf("invalid disk: %v", op)
	}
	if bd[p.y][p.x] != domain.DiskNone {
		return othelloerror.ErrInvalidPlaceForDisk
	}
	reversedCount := 0
	// 八方を確認し、ひっくり返す対象の石を記録する
	targets := make(map[point]struct{})
	targets[p] = struct{}{}
	for dir := 0; dir < len(dx); dir++ {
		ny := p.y + dy[dir]
		nx := p.x + dx[dir]
		ops := 0
		for !outOfBoard(nx, ny) && bd[ny][nx] == op {
			ops++
			ny += dy[dir]
			nx += dx[dir]
		}
		if ops > 0 && !outOfBoard(nx, ny) && bd[ny][nx] == d {
			reversedCount += ops
			for i := 0; i < ops; i++ {
				ny -= dy[dir]
				nx -= dx[dir]
				targets[point{x: nx, y: ny}] = struct{}{}
			}
		}
	}

	// ひっくり返した石がなければエラーを返す
	if reversedCount == 0 {
		return fmt.Errorf("no reversed disks found")
	}

	// ひっくり返す場所と石を置く場所を自分の石にする
	for t := range targets {
		bd[t.y][t.x] = d
	}

	return nil
}

// 石を置くことができる場所の一覧 (計算量: BoardSize^3)
func availableSpaces(bd *domain.Board, d domain.Disk) map[point]struct{} {
	res := make(map[point]struct{})
	op, _ := d.OppositeColer() // TODO: error handling
	for y := 0; y < domain.BoardSize; y++ {
		for x := 0; x < domain.BoardSize; x++ {
			if bd[y][x] != domain.DiskNone {
				continue
			}
			// 置ける箇所から八方に進めて、他色をまたいで自色があれば置けると判定
			for dir := 0; dir < len(dx); dir++ {
				ny := y + dy[dir]
				nx := x + dx[dir]
				ops := 0
				for !outOfBoard(nx, ny) && bd[ny][nx] == op {
					ops++
					ny += dy[dir]
					nx += dx[dir]
				}
				if ops > 0 && !outOfBoard(nx, ny) && bd[ny][nx] == d {
					res[point{x, y}] = struct{}{}
				}
			}
		}
	}
	return res
}

// ゲーム終了であれば true を返す
func (g *Game) checkFinished() bool {
	avl1 := availableSpaces(g.board, domain.DiskBlack)
	avl2 := availableSpaces(g.board, domain.DiskWhite)
	return len(avl1) == 0 && len(avl2) == 0
}

func countDisks(b *domain.Board) (blackCount, whiteCount int) {
	for i := 0; i < domain.BoardSize; i++ {
		for j := 0; j < domain.BoardSize; j++ {
			switch b[i][j] {
			case domain.DiskBlack:
				blackCount++
			case domain.DiskWhite:
				whiteCount++
			}
		}
	}
	return
}

func NewGame() *Game {
	return &Game{
		current: domain.DiskBlack, // 最初は黒から始める
		board:   domain.NewBoard(),
	}
}

type GameController struct {
	g *Game
}

// Update は入力をもとに現在の盤面を update する。
func (g *Game) Update(a othellodto.Action, x int, y int) (finished bool, winner domain.Disk, err error) {
	cur := g.current
	op, err := cur.OppositeColer()
	if err != nil {
		return false, 0, fmt.Errorf("invalid current state: %w", err)
	}
	switch a {
	case othellodto.ActionGiveUp:
		return true, op, nil
	case othellodto.ActionPass:
		g.current = op
		if len(availableSpaces(g.board, cur)) != 0 {
			return false, 0, fmt.Errorf("pass is not allowed")
		}
		return false, 0, nil
	case othellodto.ActionPutDisk:
		g.current = op
		break
	default:
		return false, 0, fmt.Errorf("unknow action is passed")
	}
	p, err := NewPoint(x, y)
	if err != nil {
		return false, 0, err
	}
	// 石をおいた場合にひっくり返すことができる場所かチェックする
	aval := availableSpaces(g.board, cur)
	if _, ok := aval[p]; !ok {
		return false, 0, othelloerror.ErrInvalidPlaceForDisk
	}

	if err := putDisk(g.board, cur, p); err != nil {
		return false, 0, err
	}
	g.current = op
	// 勝利判定
	if g.checkFinished() {
		blacks, whites := countDisks(g.board)
		var winner domain.Disk
		if blacks == whites {
			winner = domain.DiskNone
		} else if blacks > whites && cur == domain.DiskBlack || whites > blacks && cur == domain.DiskWhite {
			winner = cur
		} else {
			winner = op
		}
		return true, winner, nil
	}
	return false, 0, nil
}
