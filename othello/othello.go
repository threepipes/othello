package othello

import (
	"fmt"
	othellodto "othello/othello/dto"
	"othello/othello/internal/usecase"
)

func NewGame() *usecase.Game {
	return usecase.NewGame()
}

func UpdateGame(g *usecase.Game, a othellodto.Action, x int, y int) (*othellodto.GameState, error) {
	finished, winner, err := g.Update(a, x, y)
	if err != nil {
		return nil, fmt.Errorf("failed to update game: %w", err)
	}

	bd := g.Board()
	return &othellodto.GameState{
		CurrentTurn:  othellodto.DiskFromDomain(g.CurrentTurn()),
		Board:        othellodto.BoardFromDomain(&bd),
		GameFinished: finished,
		Winner:       othellodto.DiskFromDomain(winner),
	}, nil
}
