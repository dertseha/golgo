package native

import (
	"fmt"

	"github.com/go-gl/gl"
)

type NativeOpenGl struct {
	buffers            map[uint]gl.Buffer
	locations          map[int]gl.AttribLocation
	locationsByProgram map[uint][]int
	programs           map[uint]gl.Program
	shaders            map[uint]gl.Shader
	uniforms           map[int]gl.UniformLocation
	uniformsByProgram  map[uint][]int
}

func CreateGles2Wrapper() *NativeOpenGl {
	result := &NativeOpenGl{
		buffers:            make(map[uint]gl.Buffer),
		locations:          make(map[int]gl.AttribLocation),
		locationsByProgram: make(map[uint][]int),
		programs:           make(map[uint]gl.Program),
		shaders:            make(map[uint]gl.Shader),
		uniforms:           make(map[int]gl.UniformLocation),
		uniformsByProgram:  make(map[uint][]int)}

	return result
}

func (native *NativeOpenGl) fetchAttribLocation(key int) gl.AttribLocation {
	obj, ok := native.locations[key]

	if !ok {
		panic("AttribLocation not found with key " + fmt.Sprintf("%d", key))
	}

	return obj
}

func (native *NativeOpenGl) fetchBuffer(key uint) gl.Buffer {
	obj, ok := native.buffers[key]

	if !ok {
		panic("Buffer not found with key " + fmt.Sprintf("%d", key))
	}

	return obj
}

func (native *NativeOpenGl) fetchProgram(key uint) gl.Program {
	obj, ok := native.programs[key]

	if !ok {
		panic("Program not found with key " + fmt.Sprintf("%d", key))
	}

	return obj
}

func (native *NativeOpenGl) fetchShader(key uint) gl.Shader {
	obj, ok := native.shaders[key]

	if !ok {
		panic("Shader not found with key " + fmt.Sprintf("%d", key))
	}

	return obj
}

func (native *NativeOpenGl) fetchUniform(key int) gl.UniformLocation {
	obj, ok := native.uniforms[key]

	if !ok {
		panic("Uniform not found with key " + fmt.Sprintf("%d", key))
	}

	return obj
}

func (native *NativeOpenGl) AttachShader(program uint, shader uint) {
	objProgram := native.fetchProgram(program)
	objShader := native.fetchShader(shader)

	objProgram.AttachShader(objShader)
}

func (native *NativeOpenGl) BindAttribLocation(program uint, index int, name string) {
	objProgram := native.fetchProgram(program)
	objAttribLocation := native.fetchAttribLocation(index)

	objProgram.BindAttribLocation(objAttribLocation, name)
}

func (native *NativeOpenGl) BindBuffer(target uint, buffer uint) {
	native.fetchBuffer(buffer).Bind(gl.GLenum(target))
}

func (native *NativeOpenGl) BufferData(target uint, size int, data interface{}, usage uint) {
	gl.BufferData(gl.GLenum(target), size, data, gl.GLenum(usage))
}

func (*NativeOpenGl) Clear(mask uint) {
	gl.Clear(gl.GLbitfield(mask))
}

func (*NativeOpenGl) ClearColor(red float32, green float32, blue float32, alpha float32) {
	gl.ClearColor(gl.GLclampf(red), gl.GLclampf(green), gl.GLclampf(blue), gl.GLclampf(alpha))
}

func (native *NativeOpenGl) CompileShader(shader uint) {
	native.fetchShader(shader).Compile()
}

func (native *NativeOpenGl) CreateProgram() uint {
	program := gl.CreateProgram()
	key := uint(program)

	native.programs[key] = program
	native.locationsByProgram[key] = make([]int, 0)

	return key
}

func (native *NativeOpenGl) CreateShader(shaderType uint) uint {
	shader := gl.CreateShader(gl.GLenum(shaderType))
	key := uint(shader)

	native.shaders[key] = shader

	return key
}

func (native *NativeOpenGl) DeleteBuffers(buffers []uint) {
	for _, key := range buffers {
		native.fetchBuffer(key).Delete()
		delete(native.buffers, key)
	}
}

func (native *NativeOpenGl) DeleteProgram(program uint) {
	native.fetchProgram(program).Delete()
	delete(native.programs, program)

	for _, key := range native.locationsByProgram[program] {
		delete(native.locations, key)
	}
	delete(native.locationsByProgram, program)
	for _, key := range native.uniformsByProgram[program] {
		delete(native.uniforms, key)
	}
	delete(native.uniformsByProgram, program)
}

func (native *NativeOpenGl) DeleteShader(shader uint) {
	native.fetchShader(shader).Delete()
	delete(native.shaders, shader)
}

func (native *NativeOpenGl) DrawArrays(mode uint, first int, count int) {
	gl.DrawArrays(gl.GLenum(mode), first, count)
}

func (native *NativeOpenGl) Enable(cap uint) {
	gl.Enable(gl.GLenum(cap))
}

func (native *NativeOpenGl) EnableVertexAttribArray(index int) {
	native.fetchAttribLocation(index).EnableArray()
}

func (native *NativeOpenGl) GenBuffers(n int) []uint {
	buffers := make([]gl.Buffer, n)
	gl.GenBuffers(buffers)
	result := make([]uint, 0)

	for _, obj := range buffers {
		key := uint(obj)
		native.buffers[key] = obj
		result = append(result, key)
	}

	return result
}

func (native *NativeOpenGl) GetAttribLocation(program uint, name string) int {
	obj := native.fetchProgram(program).GetAttribLocation(name)
	key := int(obj)

	native.locationsByProgram[program] = append(native.locationsByProgram[program], key)
	native.locations[key] = obj

	return key
}

func (native *NativeOpenGl) GetError() uint {
	return uint(gl.GetError())
}

func (native *NativeOpenGl) GetProgramParameter(program uint, param uint) int {
	return native.fetchProgram(program).Get(gl.GLenum(param))
}

func (native *NativeOpenGl) GetShaderInfoLog(shader uint) string {
	return native.fetchShader(shader).GetInfoLog()
}

func (native *NativeOpenGl) GetShaderParameter(shader uint, param uint) int {
	return native.fetchShader(shader).Get(gl.GLenum(param))
}

func (native *NativeOpenGl) GetUniformLocation(program uint, name string) int {
	obj := native.fetchProgram(program).GetUniformLocation(name)
	key := int(obj)

	native.uniforms[key] = obj
	native.uniformsByProgram[program] = append(native.uniformsByProgram[program], key)

	return key
}

func (native *NativeOpenGl) LinkProgram(program uint) {
	native.fetchProgram(program).Link()
}

func (native *NativeOpenGl) ReadPixels(x int, y int, width int, height int, format uint, pixelType uint, pixels interface{}) {
	gl.ReadPixels(x, y, width, height, gl.GLenum(format), gl.GLenum(pixelType), pixels)
}

func (native *NativeOpenGl) ShaderSource(shader uint, source string) {
	native.fetchShader(shader).Source(source)
}

func (native *NativeOpenGl) UniformMatrix4fv(location int, transpose bool, value *[16]float32) {
	native.fetchUniform(location).UniformMatrix4f(transpose, value)
}

func (native *NativeOpenGl) UseProgram(program uint) {
	native.fetchProgram(program).Use()
}

func (native *NativeOpenGl) VertexAttribOffset(index int, size int, attribType uint, normalized bool, stride int, offset int) {
	native.fetchAttribLocation(index).AttribPointer(uint(size), gl.GLenum(attribType), normalized, stride, uintptr(offset))
}

func (native *NativeOpenGl) Viewport(x int, y int, width int, height int) {
	gl.Viewport(x, y, width, height)
}
