package cardGame

import "testing"

func TestSortSymbolNumber_Len(t *testing.T) {
	tests := []struct {
		name string
		a    SortSymbolNumber
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Len(); got != tt.want {
				t.Errorf("SortSymbolNumber.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortSymbolNumber_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		a    SortSymbolNumber
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestSortSymbolNumber_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		a    SortSymbolNumber
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("SortSymbolNumber.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortCard(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortCard(tt.args.cards)
		})
	}
}

func TestCard_Info(t *testing.T) {
	tests := []struct {
		name  string
		c     Card
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.c.Info()
			if got != tt.want {
				t.Errorf("Card.Info() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Card.Info() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIsCardHigher(t *testing.T) {
	type args struct {
		c1 Card
		c2 Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCardHigher(tt.args.c1, tt.args.c2); got != tt.want {
				t.Errorf("IsCardHigher() = %v, want %v", got, tt.want)
			}
		})
	}
}
