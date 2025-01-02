package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type Vertex2D [2]float32

const (
	width  = 500
	height = 500
)

func main() {
	runtime.LockOSThread()

	window, terminate := initGlfw()
	defer terminate()

	program, err := initOpenGL()
	if err != nil {
		log.Fatalln("failed to initialize OpenGL:", err)
	}

	framesNumber := 100
	squares := createSetOfSquaresVAO(framesNumber)

	drawn := 0
	objectSize := 6 // 2 triangles (len of square / 3)
	frameCounter := 0
	for !window.ShouldClose() {
		n := frameCounter / (framesNumber - 1)
		if n%2 == 0 {
			drawn++
		} else {
			drawn--
		}
		vaoToDraw := squares[drawn]
		draw(vaoToDraw, objectSize, window, program)

		frameCounter++
	}
}

func createSetOfSquaresVAO(number int) []uint32 {
	squares := make([]uint32, 0, number)
	maxRadius := 0.5
	minRadius := 0.0
	stepSize := float32(maxRadius-minRadius) / float32(number)
	for i := 0; i < number; i++ {
		point := 0.0 + stepSize*float32(i)
		square := createSquare(
			Vertex2D{-0.5, 0.5},
			Vertex2D{-0.5, -0.5},
			Vertex2D{0.5, -0.5},
			Vertex2D{point, point},
		)
		vao := makeVao(square)
		squares = append(squares, vao)
	}

	return squares
}

func draw(vao uint32, arraySize int, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(arraySize))

	glfw.PollEvents()
	window.SwapBuffers()
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() (*glfw.Window, func()) {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Snake", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window, func() {
		glfw.Terminate()
	}
}

// initOpenGL initializes OpenGL and returns a program.
func initOpenGL() (uint32, error) {
	if err := gl.Init(); err != nil {
		return 0, err
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShaderSource, err := loadFile("./assets/shaders/vertex.glsl")
	if err != nil {
		return 0, err
	}
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}
	fragmentShaderSource, err := loadFile("./assets/shaders/fragment.glsl")
	if err != nil {
		return 0, err
	}
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog, nil
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

// loadShaderFile reads a file and returns the content as a string.
func loadFile(filename string) (string, error) {
	// Read the file
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
