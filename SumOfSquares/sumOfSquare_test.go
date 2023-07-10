package sumofsquares

import "testing"

func Test_sumOfSquares(t *testing.T) {
	type args struct {
		lengthOfSide []string
		iterate      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfSquares(tt.args.lengthOfSide, tt.args.iterate); got != tt.want {
				t.Errorf("sumOfSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
