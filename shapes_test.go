package main

import "testing"

func TestCreateTriangle(t *testing.T) {
	type args struct {
		v1 Vertex2D
		v2 Vertex2D
		v3 Vertex2D
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "simple triangle",
			args: args{
				v1: Vertex2D{0, 0},
				v2: Vertex2D{1, 0},
				v3: Vertex2D{0, 1},
			},
			want: []float32{
				0, 0, 0,
				1, 0, 0,
				0, 1, 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createTriangle(tt.args.v1, tt.args.v2, tt.args.v3); !AreArraysEqual(got, tt.want) {
				t.Errorf("createTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSquare(t *testing.T) {
	type args struct {
		v1 Vertex2D
		v2 Vertex2D
		v3 Vertex2D
		v4 Vertex2D
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "simple square",
			args: args{
				v1: Vertex2D{0, 0},
				v2: Vertex2D{1, 0},
				v3: Vertex2D{1, 1},
				v4: Vertex2D{0, 1},
			},
			want: []float32{
				0, 0, 0,
				1, 0, 0,
				1, 1, 0,
				1, 1, 0,
				0, 1, 0,
				0, 0, 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createSquare(tt.args.v1, tt.args.v2, tt.args.v3, tt.args.v4); !AreArraysEqual(got, tt.want) {
				t.Errorf("createSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
