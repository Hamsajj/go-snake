package main

func createTriangle(v1, v2, v3 Vertex2D) []float32 {
	return []float32{
		v1[0], v1[1], 0,
		v2[0], v2[1], 0,
		v3[0], v3[1], 0,
	}
}

func createSquare(v1, v2, v3, v4 Vertex2D) []float32 {
	return []float32{
		v1[0], v1[1], 0,
		v2[0], v2[1], 0,
		v3[0], v3[1], 0,
		v3[0], v3[1], 0,
		v4[0], v4[1], 0,
		v1[0], v1[1], 0,
	}
}
