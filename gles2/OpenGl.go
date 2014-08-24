/*
 * The gles2 package contains the constants and interface compatible with
 * OpenGL ES2.
 *
 * All values are based on the public header from the Khronos group:
 * http://www.khronos.org/registry/gles/
 */
package gles2

type OpenGl interface {
	AttachShader(program uint, shader uint)

	BindAttribLocation(program uint, index int, name string)
	BindBuffer(target uint, buffer uint)
	BufferData(target uint, size int, data interface{}, usage uint)

	Clear(mask uint)
	ClearColor(red float32, green float32, blue float32, alpha float32)

	CompileShader(shader uint)

	CreateProgram() uint
	CreateShader(shaderType uint) uint

	DeleteBuffers(buffers []uint)
	DeleteProgram(program uint)
	DeleteShader(shader uint)

	DrawArrays(mode uint, first int, count int)

	Enable(cap uint)
	EnableVertexAttribArray(index int)

	GenBuffers(n int) []uint

	GetAttribLocation(program uint, name string) int
	GetError() uint
	GetShaderInfoLog(shader uint) string
	GetShaderParameter(shader uint, param uint) int
	GetProgramParameter(program uint, param uint) int
	GetUniformLocation(program uint, name string) int

	LinkProgram(program uint)

	ShaderSource(shader uint, source string)

	UniformMatrix4fv(location int, transpose bool, value ...[]float32)
	UseProgram(program uint)

	VertexAttribOffset(index int, size int, attribType uint, normalized bool, stride int, offset int)
	Viewport(x int, y int, width int, height int)
}
