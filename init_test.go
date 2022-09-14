package cardGame

import "testing"

func TestTeenPati_InitGame(t *testing.T) {
	type args struct {
		playerName      string
		numberOfPlayers int
	}
	tests := []struct {
		name string
		c    *TeenPati
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.InitGame(tt.args.playerName, tt.args.numberOfPlayers)
		})
	}
}

func TestTeenPati_StartGame(t *testing.T) {
	tests := []struct {
		name string
		c    *TeenPati
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.StartGame()
		})
	}
}
