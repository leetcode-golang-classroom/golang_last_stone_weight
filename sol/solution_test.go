package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	stones := []int{2, 7, 4, 1, 8, 1}
	for idx := 0; idx < b.N; idx++ {
		lastStoneWeight(stones)
	}
}
func Test_lastStoneWeight(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "stones = [2,7,4,1,8,1]",
			args: args{stones: []int{2, 7, 4, 1, 8, 1}},
			want: 1,
		},
		{
			name: "stones = [1]",
			args: args{stones: []int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeight(tt.args.stones); got != tt.want {
				t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
