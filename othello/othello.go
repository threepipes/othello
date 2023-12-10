package othello

import (
	"fmt"
	othellodto "othello/othello/dto"
	"othello/othello/internal/usecase"
)

type GameController struct {
	game *usecase.Game
}

func NewGame() (*GameController, *othellodto.GameState) {
	ctl := &GameController{usecase.NewGame()}
	bd := ctl.game.Board()
	return ctl, &othellodto.GameState{
		CurrentTurn: othellodto.DiskFromDomain(ctl.game.CurrentTurn()),
		Board:       othellodto.BoardFromDomain(&bd),
	}
}

func (c *GameController) UpdateGame(a othellodto.Action, x int, y int) (*othellodto.GameState, error) {
	finished, winner, err := c.game.Update(a.Domain(), x, y)
	if err != nil {
		return nil, fmt.Errorf("failed to update game: %w", err)
	}

	bd := c.game.Board()
	return &othellodto.GameState{
		CurrentTurn:  othellodto.DiskFromDomain(c.game.CurrentTurn()),
		Board:        othellodto.BoardFromDomain(&bd),
		GameFinished: finished,
		Winner:       othellodto.DiskFromDomain(winner),
	}, nil
}

// 現在のプレイヤーが石を置ける状態か
func (c *GameController) Playable() bool {
	bd := c.game.Board()
	return len(usecase.AvailableSpaces(&bd, c.game.CurrentTurn())) > 0
}
