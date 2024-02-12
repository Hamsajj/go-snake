package main

import "testing"

func TestAreArraysEqual(t *testing.T) {
	type args struct {
		a []float32
		b []float32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal arrays",
			args: args{
				a: []float32{1, 2, 3},
				b: []float32{1, 2, 3},
			},
			want: true,
		},
		{
			name: "different arrays",
			args: args{
				a: []float32{1, 2, 3},
				b: []float32{1, 2, 4},
			},
			want: false,
		},
		{
			name: "different length",
			args: args{
				a: []float32{1, 2, 3},
				b: []float32{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "different length",
			args: args{
				a: []float32{},
				b: []float32{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreArraysEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AreArraysEqual(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
