package usecase

import (
	"othello/othello/internal/domain"
	"reflect"
	"strings"
	"testing"
)

func strToBoard(s []string) *domain.Board {
	bd := domain.Board{}
	for y, row := range s {
		for x, col := range row {
			switch col {
			case 'o':
				bd[y][x] = domain.DiskWhite
			case 'x':
				bd[y][x] = domain.DiskBlack
			default:
				bd[y][x] = domain.DiskNone
			}
		}
	}
	return &bd
}

func Test_availableSpaces(t *testing.T) {
	type args struct {
		bd *domain.Board
		d  domain.Disk
	}
	tests := []struct {
		name string
		args args
		want map[point]struct{}
	}{
		{
			name: "初期状態",
			args: args{
				bd: strToBoard([]string{
					"........", // 0
					"........", // 1
					"........", // 2
					"...xo...", // 3
					"...ox...", // 4
					"........", // 5
					"........", // 6
					"........", // 7
				}),
				d: domain.DiskBlack,
			},
			want: map[point]struct{}{
				{x: 4, y: 2}: {},
				{x: 5, y: 3}: {},
				{x: 2, y: 4}: {},
				{x: 3, y: 5}: {},
			},
		},
		{
			name: "途中状態",
			args: args{
				bd: strToBoard([]string{
					"........", // 0
					"........", // 1
					"....xo..", // 2
					"...xo...", // 3
					"...oxo..", // 4
					"....x...", // 5
					"........", // 6
					"........", // 7
				}),
				d: domain.DiskBlack,
			},
			want: map[point]struct{}{
				{x: 2, y: 3}: {},
				{x: 2, y: 4}: {},
				{x: 3, y: 5}: {},
				{x: 5, y: 3}: {},
				{x: 6, y: 2}: {},
				{x: 6, y: 3}: {},
				{x: 6, y: 4}: {},
			},
		},
		{
			name: "置けない",
			args: args{
				bd: strToBoard([]string{
					"........", // 0
					"........", // 1
					"....oooo", // 2
					"...oo...", // 3
					"...ooo..", // 4
					"....o...", // 5
					"........", // 6
					"ox......", // 7
				}),
				d: domain.DiskBlack,
			},
			want: map[point]struct{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := availableSpaces(tt.args.bd, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("availableSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func boardToStr(bd *domain.Board) string {
	b := strings.Builder{}
	b.WriteRune('\n')
	for i := 0; i < domain.BoardSize; i++ {
		for j := 0; j < domain.BoardSize; j++ {
			switch bd[i][j] {
			case domain.DiskBlack:
				b.WriteRune('x')
			case domain.DiskWhite:
				b.WriteRune('o')
			case domain.DiskNone:
				b.WriteRune('.')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func Test_putDisk(t *testing.T) {
	type args struct {
		bd *domain.Board
		d  domain.Disk
		p  point
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Board
		wantErr bool
	}{
		{
			name: "初期状態",
			args: args{
				bd: strToBoard([]string{
					"........", // 0
					"........", // 1
					"........", // 2
					"...xo...", // 3
					"...ox...", // 4
					"........", // 5
					"........", // 6
					"........", // 7
				}),
				d: domain.DiskBlack,
				p: point{x: 5, y: 3},
			},
			want: strToBoard([]string{
				"........", // 0
				"........", // 1
				"........", // 2
				"...xxx..", // 3
				"...ox...", // 4
				"........", // 5
				"........", // 6
				"........", // 7
			}),
		},
		{
			name: "途中状態",
			args: args{
				bd: strToBoard([]string{
					"........", // 0
					"........", // 1
					"....xo..", // 2
					"...xo...", // 3
					"...oxo..", // 4
					"....x...", // 5
					"........", // 6
					"........", // 7
				}),
				d: domain.DiskBlack,
				p: point{x: 6, y: 3},
			},
			want: strToBoard([]string{
				"........", // 0
				"........", // 1
				"....xo..", // 2
				"...xo.x.", // 3
				"...oxx..", // 4
				"....x...", // 5
				"........", // 6
				"........", // 7
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := tt.args.bd
			if err := putDisk(board, tt.args.d, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("putDisk() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(board, tt.want) {
				t.Errorf("putDisk() = %v, want: %v", boardToStr(board), boardToStr(tt.want))
			}
		})
	}
}
