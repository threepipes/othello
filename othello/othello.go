package othello

import (
	"fmt"
	othellodto "othello/othello/dto"
	"othello/othello/internal/usecase"
)

type GameController struct {
	game *usecase.Game
}

func NewGame() *GameController {
	return &GameController{usecase.NewGame()}
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
