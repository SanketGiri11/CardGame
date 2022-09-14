package cardGame

import (
	"reflect"
	"testing"
)

func TestInitDeck(t *testing.T) {
	tests := []struct {
		name string
		want Deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitDeck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_Shuffle(t *testing.T) {
	tests := []struct {
		name string
		deck Deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.deck.Shuffle()
		})
	}
}

func TestDeck_Show(t *testing.T) {
	tests := []struct {
		name string
		deck Deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.deck.Show()
		})
	}
}

func TestDeck_Draw(t *testing.T) {
	type args struct {
		numberOfCards int
	}
	tests := []struct {
		name string
		deck *Deck
		args args
		want Deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.deck.Draw(tt.args.numberOfCards); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deck.Draw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_AddCard(t *testing.T) {
	type args struct {
		card Card
	}
	tests := []struct {
		name string
		deck *Deck
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.deck.AddCard(tt.args.card)
		})
	}
}
