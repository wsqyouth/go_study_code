package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int64
		b    int64
		want int64
	}{
		{
			name: "test",
			a:    1,
			b:    2,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("TestAdd() got = %v, want %v", got, tt.want)
			}

		})
	}
}
