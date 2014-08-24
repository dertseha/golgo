package view

import (
	"fmt"
	"math"

	"github.com/ungerik/go3d/mat4"
	"github.com/ungerik/go3d/vec3"

	gles "github.com/dertseha/golgo/gles2"
)

var vertexShaderSource = `
  attribute vec3 aVertexPosition;
  attribute vec4 aVertexColor;

  uniform mat4 uMVMatrix;
  uniform mat4 uPMatrix;

  varying vec4 vColor;

  void main(void) {
    gl_Position = uPMatrix * uMVMatrix * vec4(aVertexPosition, 1.0);
    vColor = aVertexColor;
  }
`

var fragmentShaderSource = `
  #ifdef GL_ES
    precision mediump float;
  #endif

  varying vec4 vColor;

  void main(void) {
    gl_FragColor = vColor;
  }
`

type NeheExample02 struct {
	gl gles.OpenGl

	width  int
	height int

	vertexPosition               int
	triangleVertexPositionBuffer uint
	vertexColor                  int
	triangleVertexColorBuffer    uint

	pMatrix         mat4.T
	pMatrixUniform  int
	mvMatrix        mat4.T
	mvMatrixUniform int
}

func NewNeheExample02(gl gles.OpenGl, width int, height int) *NeheExample02 {
	result := &NeheExample02{
		gl:     gl,
		width:  width,
		height: height}

	return result
}

func checkError(gl gles.OpenGl, stage string) {
	result := gl.GetError()

	if result != gles.NO_ERROR {
		fmt.Println("!!!!! ERROR " + fmt.Sprintf("0x%04X", result) + " at " + stage)
	}
}

func (example *NeheExample02) prepareShader(shaderType uint, source string) uint {
	gl := example.gl
	shader := gl.CreateShader(shaderType)

	gl.ShaderSource(shader, source)
	gl.CompileShader(shader)

	compileStatus := gl.GetShaderParameter(shader, gles.COMPILE_STATUS)
	if compileStatus == 0 {
		fmt.Println("Error: compile of " + fmt.Sprintf("0x%04X", shaderType) + " failed: " +
			fmt.Sprintf("%d", compileStatus) + "  - " + gl.GetShaderInfoLog(shader))
	}

	return shader
}

func (example *NeheExample02) initShaders() {
	gl := example.gl
	fragmentShader := example.prepareShader(gles.FRAGMENT_SHADER, fragmentShaderSource)
	vertexShader := example.prepareShader(gles.VERTEX_SHADER, vertexShaderSource)
	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	if gl.GetProgramParameter(program, gles.LINK_STATUS) == 0 {
		fmt.Println("Error: link failed")
	}

	gl.UseProgram(program)

	checkError(gl, "using program")

	example.vertexPosition = gl.GetAttribLocation(program, "aVertexPosition")
	checkError(gl, "get attrib loc 1")
	gl.EnableVertexAttribArray(example.vertexPosition)
	checkError(gl, "enable attrib loc 1")
	example.vertexColor = gl.GetAttribLocation(program, "aVertexColor")
	gl.EnableVertexAttribArray(example.vertexColor)

	example.pMatrixUniform = gl.GetUniformLocation(program, "uPMatrix")
	example.mvMatrixUniform = gl.GetUniformLocation(program, "uMVMatrix")
	checkError(gl, "uniforms")
}

func (example *NeheExample02) initBuffers() {
	gl := example.gl
	example.triangleVertexPositionBuffer = gl.GenBuffers(1)[0]

	gl.BindBuffer(gles.ARRAY_BUFFER, example.triangleVertexPositionBuffer)
	var vertices = []float32{
		0.0, 1.0, 0.0,
		-1.0, -1.0, 0.0,
		1.0, -1.0, 0.0}
	gl.BufferData(gles.ARRAY_BUFFER, len(vertices)*4, vertices, gles.STATIC_DRAW)
	checkError(gl, "buffered data 1")

	example.triangleVertexColorBuffer = gl.GenBuffers(1)[0]
	gl.BindBuffer(gles.ARRAY_BUFFER, example.triangleVertexColorBuffer)
	var colors = []float32{
		1.0, 0.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
		0.0, 0.0, 1.0, 1.0}
	gl.BufferData(gles.ARRAY_BUFFER, len(colors)*4, colors, gles.STATIC_DRAW)
	checkError(gl, "buffered data 2")
}

func (example *NeheExample02) setMatrixUniforms() {
	gl := example.gl
	gl.UniformMatrix4fv(example.pMatrixUniform, false, example.pMatrix.Slice())
	checkError(gl, "set uniforms 1")
	gl.UniformMatrix4fv(example.mvMatrixUniform, false, example.mvMatrix.Slice())
	checkError(gl, "set uniforms 2")
}

func (example *NeheExample02) perspective(fovy float32, aspect float32, near float32, far float32) {
	top := near * float32(math.Tan(float64(fovy*math.Pi/360.0)))
	right := top * aspect

	example.pMatrix.AssignPerspectiveProjection(-right, right, -top, top, 0.1, 100.0)
}

func (example *NeheExample02) Init() {
	gl := example.gl
	example.initShaders()
	example.initBuffers()

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Enable(gles.DEPTH_TEST)
}

func (example *NeheExample02) DrawScene() {
	gl := example.gl
	gl.Viewport(0, 0, example.width, example.height)
	checkError(gl, "viewport")
	gl.Clear(gles.COLOR_BUFFER_BIT | gles.DEPTH_BUFFER_BIT)
	checkError(gl, "clear")

	example.perspective(45.0, float32(example.width)/float32(example.height), 0.1, 100.0)
	example.mvMatrix = mat4.Ident
	example.mvMatrix.Translate(&vec3.T{-1.0, 0.0, -7.0})
	example.setMatrixUniforms()

	gl.BindBuffer(gles.ARRAY_BUFFER, example.triangleVertexPositionBuffer)
	checkError(gl, "draw bind 1")
	gl.VertexAttribOffset(example.vertexPosition, 3, gles.FLOAT, false, 0, 0)
	checkError(gl, "draw offset 1")

	gl.BindBuffer(gles.ARRAY_BUFFER, example.triangleVertexColorBuffer)
	checkError(gl, "draw bind 2")
	gl.VertexAttribOffset(example.vertexColor, 4, gles.FLOAT, false, 0, 0)
	checkError(gl, "draw offset 2")

	gl.DrawArrays(gles.TRIANGLES, 0, 3)
	checkError(gl, "draw arrays")
}
