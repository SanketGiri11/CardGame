package cardGame

import (
	"reflect"
	"testing"
)

func TestPlayers_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    Players
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Players.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_DrawCards(t *testing.T) {
	type args struct {
		deck         *Deck
		numberOfDraw int
	}
	tests := []struct {
		name string
		p    *Player
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.DrawCards(tt.args.deck, tt.args.numberOfDraw)
		})
	}
}

func TestPlayer_ThrowCards(t *testing.T) {
	type args struct {
		cardIndex int
		deck      *Deck
	}
	tests := []struct {
		name string
		p    *Player
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.ThrowCards(tt.args.cardIndex, tt.args.deck)
		})
	}
}

func TestPlayer_ShowHand(t *testing.T) {
	tests := []struct {
		name string
		p    Player
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.ShowHand()
		})
	}
}

func TestRegister(t *testing.T) {
	type args struct {
		pl         []string
		numberOfAI int
	}
	tests := []struct {
		name string
		args args
		want Players
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Register(tt.args.pl, tt.args.numberOfAI); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
