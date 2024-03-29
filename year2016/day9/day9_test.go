package day9

import "testing"

func TestPart1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//ADVENT contains no markers and decompresses to itself with no changes, resulting in a decompressed length of 6.
		{"", args{"ADVENT"}, "6"},
		//A(1x5)BC repeats only the B a total of 5 times, becoming ABBBBBC for a decompressed length of 7.
		{"", args{"A(1x5)BC"}, "7"},
		//(3x3)XYZ becomes XYZXYZXYZ for a decompressed length of 9.
		{"", args{"(3x3)XYZ"}, "9"},
		//A(2x2)BCD(2x2)EFG doubles the BC and EF, becoming ABCBCDEFEFG for a decompressed length of 11.
		{"", args{"A(2x2)BCD(2x2)EFG"}, "11"},
		//(6x1)(1x3)A simply becomes (1x3)A - the (1x3) looks like a marker, but because it's within a data section of another marker, it is not treated any differently from the A that comes after it. It has a decompressed length of 6.
		{"", args{"(6x1)(1x3)A"}, "6"},
		//X(8x2)(3x3)ABCY becomes X(3x3)ABC(3x3)ABCY (for a decompressed length of 18), because the decompressed data from the (8x2) marker (the (3x3)ABC) is skipped and not processed further.
		{"", args{"X(8x2)(3x3)ABCY"}, "18"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (Day9{}).Part1(tt.args.input); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//(3x3)XYZ still becomes XYZXYZXYZ, as the decompressed section contains no markers.
		{"", args{"(3x3)XYZ"}, "9"},
		//X(8x2)(3x3)ABCY becomes XABCABCABCABCABCABCY, because the decompressed data from the (8x2) marker is then further decompressed, thus triggering the (3x3) marker twice for a total of six ABC sequences.
		{"", args{"X(8x2)(3x3)ABCY"}, "20"},
		//(27x12)(20x12)(13x14)(7x10)(1x12)A decompresses into a string of A repeated 241920 times.
		{"", args{"(27x12)(20x12)(13x14)(7x10)(1x12)A"}, "241920"},
		//(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN becomes 445 characters long.
		{"", args{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN"}, "445"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (Day9{}).Part2(tt.args.input); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
