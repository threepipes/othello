package cli

import (
	"errors"
	"fmt"
	"os"
	"othello/cli/prompt"
	"othello/othello"
	othellodto "othello/othello/dto"
	othelloerror "othello/othello/error"
	"strings"
)

func playerToStr(d othellodto.Disk) string {
	switch d {
	case othellodto.DiskBlack:
		return "x"
	case othellodto.DiskWhite:
		return "o"
	default:
		return "Unknown" // TODO: error
	}
}

func boardToStr(bd *othellodto.Board) string {
	builder := strings.Builder{}
	builder.WriteString("   1 2 3 4 5 6 7 8\n")
	for i := 0; i < len(bd); i++ {
		builder.WriteRune(rune('A' + i))
		builder.WriteString(": ")
		for j := 0; j < len(bd[i]); j++ {
			switch bd[i][j] {
			case othellodto.DiskBlack:
				builder.WriteRune('x')
			case othellodto.DiskWhite:
				builder.WriteRune('o')
			case othellodto.DiskNone:
				builder.WriteRune('.')
			default:
				builder.WriteRune('?')
			}
			builder.WriteRune(' ')
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func readAction(ppt *prompt.Prompt, playable bool) (act othellodto.Action, x int, y int, err error) {
	if playable {
		// 石を置ける場合
		pick, err := ppt.Choose([]string{"Put a stone", "Give up"}, "")
		if err != nil {
			return 0, 0, 0, fmt.Errorf("failed to read input: %w", err)
		}
		if pick == 1 {
			act = othellodto.ActionGiveUp
		} else {
			match, err := ppt.InputStringRegexMatch("([A-H])([1-8])", "Input the point to put your disk (ex. D6): ")
			if err != nil || len(match) != 3 {
				return 0, 0, 0, fmt.Errorf("failed to read input")
			}
			row := match[1][0] - 'A'
			col := match[2][0] - '1'
			act = othellodto.ActionPutDisk
			y = int(row)
			x = int(col)
		}
	} else {
		// 石を置けない場合
		pick, err := ppt.Choose([]string{"Pass", "Give up"}, "You cannot put a stone.")
		if err != nil {
			return 0, 0, 0, fmt.Errorf("failed to read input: %w", err)
		}
		if pick == 0 {
			act = othellodto.ActionPass
		} else {
			act = othellodto.ActionGiveUp
		}
	}
	return
}

// singleGamePlay は Othello の 1 ゲームを管理する
func singleGamePlay(ppt *prompt.Prompt) error {
	fmt.Println("Game Start!")
	ctl, curState := othello.NewGame()
	for true {
		// 画面描画
		fmt.Printf("\nCurrent Player: '%s'\n", playerToStr(curState.CurrentTurn))
		fmt.Print(boardToStr(curState.Board))

		// 入力受付
		act, x, y, err := readAction(ppt, ctl.Playable())
		if err != nil {
			return fmt.Errorf("failed to read an action: %w", err)
		}

		// 更新
		state, err := ctl.UpdateGame(act, x, y)
		if errors.As(err, &othelloerror.ErrInvalidPlaceForDisk) {
			fmt.Println("You cannot put a stone there. Please try again.")
			continue
		}
		if err != nil {
			return fmt.Errorf("failed to update the game: %w", err)
		}
		curState = state

		// ゲーム終了判定
		if curState.GameFinished {
			fmt.Printf("Game Over! \nWinner Player: '%s'\n", playerToStr(curState.Winner))
			break
		}
	}
	return nil
}

// GamePlay は Othello の CLI 立ち上げから終了までを管理する
func GamePlay() error {
	ppt := prompt.NewPrompt(os.Stdin, os.Stdout)
	// タイトル表示
	fmt.Println("=== O T H E L L O ===")

	for true {
		// ゲーム開始
		if err := singleGamePlay(ppt); err != nil {
			return err
		}

		// 続けてもう一度ゲームするか確認
		pick, err := ppt.Choose([]string{"Yes", "No"}, "Do you play again?")
		if err != nil {
			return fmt.Errorf("failed to choose option: %w", err)
		}
		if pick == 0 {
			fmt.Println("Replay!")
		} else {
			fmt.Println("Bye!")
			break
		}
	}
	fmt.Println("Finish Othello.")
	return nil
}
