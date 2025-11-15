package collatzconjecture

import (
	"reflect"
	"testing"
)

func TestCollatzSequence(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		want    []int
		wantErr bool
	}{
		{
			name:    "n = 1 (base case)",
			n:       1,
			want:    []int{1},
			wantErr: false,
		},
		{
			name:    "n = 3 (classic example)",
			n:       3,
			want:    []int{3, 10, 5, 16, 8, 4, 2, 1},
			wantErr: false,
		},
		{
			name:    "n = 6 (even start)",
			n:       6,
			want:    []int{6, 3, 10, 5, 16, 8, 4, 2, 1},
			wantErr: false,
		},
		{
			name:    "n = 2 (power of 2)",
			n:       2,
			want:    []int{2, 1},
			wantErr: false,
		},
		{
			name:    "n = 16 (larger power of 2)",
			n:       16,
			want:    []int{16, 8, 4, 2, 1},
			wantErr: false,
		},
		{
			name:    "n = 5 (odd start)",
			n:       5,
			want:    []int{5, 16, 8, 4, 2, 1},
			wantErr: false,
		},
		{
			name:    "n = 0 (invalid)",
			n:       0,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "n = -5 (negative)",
			n:       -5,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "n = 27 (long sequence)",
			n:       27,
			want:    []int{27, 82, 41, 124, 62, 31, 94, 47, 142, 71, 214, 107, 322, 161, 484, 242, 121, 364, 182, 91, 274, 137, 412, 206, 103, 310, 155, 466, 233, 700, 350, 175, 526, 263, 790, 395, 1186, 593, 1780, 890, 445, 1336, 668, 334, 167, 502, 251, 754, 377, 1132, 566, 283, 850, 425, 1276, 638, 319, 958, 479, 1438, 719, 2158, 1079, 3238, 1619, 4858, 2429, 7288, 3644, 1822, 911, 2734, 1367, 4102, 2051, 6154, 3077, 9232, 4616, 2308, 1154, 577, 1732, 866, 433, 1300, 650, 325, 976, 488, 244, 122, 61, 184, 92, 46, 23, 70, 35, 106, 53, 160, 80, 40, 20, 10, 5, 16, 8, 4, 2, 1},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CollatzSequence(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("CollatzSequence(%d) error = %v, wantErr %v", tt.n, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CollatzSequence(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestCollatzLength(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		want    int
		wantErr bool
	}{
		{
			name:    "n = 1",
			n:       1,
			want:    1,
			wantErr: false,
		},
		{
			name:    "n = 3",
			n:       3,
			want:    8,
			wantErr: false,
		},
		{
			name:    "n = 6",
			n:       6,
			want:    9,
			wantErr: false,
		},
		{
			name:    "n = 2",
			n:       2,
			want:    2,
			wantErr: false,
		},
		{
			name:    "n = 16",
			n:       16,
			want:    5,
			wantErr: false,
		},
		{
			name:    "n = 27",
			n:       27,
			want:    112,
			wantErr: false,
		},
		{
			name:    "n = 0 (invalid)",
			n:       0,
			want:    0,
			wantErr: true,
		},
		{
			name:    "n = -10 (negative)",
			n:       -10,
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CollatzLength(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("CollatzLength(%d) error = %v, wantErr %v", tt.n, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CollatzLength(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestMaxCollatzInRange(t *testing.T) {
	tests := []struct {
		name       string
		start      int
		end        int
		wantNum    int
		wantLength int
		wantErr    bool
	}{
		{
			name:       "range 1-10",
			start:      1,
			end:        10,
			wantNum:    9,
			wantLength: 20,
			wantErr:    false,
		},
		{
			name:       "single number range",
			start:      1,
			end:        1,
			wantNum:    1,
			wantLength: 1,
			wantErr:    false,
		},
		{
			name:       "range 1-5",
			start:      1,
			end:        5,
			wantNum:    3,
			wantLength: 8,
			wantErr:    false,
		},
		{
			name:       "range 10-15",
			start:      10,
			end:        15,
			wantNum:    15,
			wantLength: 18,
			wantErr:    false,
		},
		{
			name:       "inverted range (start > end)",
			start:      10,
			end:        5,
			wantNum:    0,
			wantLength: 0,
			wantErr:    true,
		},
		{
			name:       "negative start",
			start:      -5,
			end:        10,
			wantNum:    0,
			wantLength: 0,
			wantErr:    true,
		},
		{
			name:       "zero start",
			start:      0,
			end:        10,
			wantNum:    0,
			wantLength: 0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum, gotLength, err := MaxCollatzInRange(tt.start, tt.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxCollatzInRange(%d, %d) error = %v, wantErr %v", tt.start, tt.end, err, tt.wantErr)
				return
			}
			if gotNum != tt.wantNum {
				t.Errorf("MaxCollatzInRange(%d, %d) num = %d, want %d", tt.start, tt.end, gotNum, tt.wantNum)
			}
			if gotLength != tt.wantLength {
				t.Errorf("MaxCollatzInRange(%d, %d) length = %d, want %d", tt.start, tt.end, gotLength, tt.wantLength)
			}
		})
	}
}

// Benchmark tests to measure performance
func BenchmarkCollatzSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CollatzSequence(27)
	}
}

func BenchmarkCollatzLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CollatzLength(27)
	}
}

func BenchmarkMaxCollatzInRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxCollatzInRange(1, 100)
	}
}
