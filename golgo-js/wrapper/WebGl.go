package wrapper

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
)

type WebGl struct {
	gl *webgl.Context

	buffers           ObjectMapper
	programs          ObjectMapper
	shaders           ObjectMapper
	uniforms          ObjectMapper
	uniformsByProgram map[uint][]uint
}

func CreateGles2Wrapper(gl *webgl.Context) *WebGl {
	result := &WebGl{
		gl:                gl,
		buffers:           NewObjectMapper(),
		programs:          NewObjectMapper(),
		shaders:           NewObjectMapper(),
		uniforms:          NewObjectMapper(),
		uniformsByProgram: make(map[uint][]uint)}

	return result
}

func (gl *WebGl) AttachShader(program uint, shader uint) {
	objShader := gl.shaders.get(shader)
	objProgram := gl.programs.get(program)

	gl.gl.AttachShader(objProgram, objShader)
}

func (gl *WebGl) BindAttribLocation(program uint, index int, name string) {
	gl.gl.BindAttribLocation(gl.programs.get(program), index, name)
}

func (gl *WebGl) BindBuffer(target uint, buffer uint) {
	gl.gl.BindBuffer(int(target), gl.buffers.get(buffer))
}

func (gl *WebGl) BufferData(target uint, size int, data interface{}, usage uint) {
	gl.gl.BufferData(int(target), data, int(usage))
}

func (gl *WebGl) Clear(mask uint) {
	gl.gl.Clear(int(mask))
}

func (gl *WebGl) ClearColor(red float32, green float32, blue float32, alpha float32) {
	gl.gl.ClearColor(red, green, blue, alpha)
}

func (gl *WebGl) CompileShader(shader uint) {
	gl.gl.CompileShader(gl.shaders.get(shader))
}

func (gl *WebGl) CreateProgram() uint {
	key := gl.programs.put(gl.gl.CreateProgram())
	gl.uniformsByProgram[key] = make([]uint, 0)

	return key
}

func (gl *WebGl) CreateShader(shaderType uint) uint {
	return gl.shaders.put(gl.gl.CreateShader(int(shaderType)))
}

func (gl *WebGl) DeleteBuffers(buffers []uint) {
	for _, buffer := range buffers {
		gl.gl.DeleteBuffer(gl.buffers.del(buffer))
	}
}

func (gl *WebGl) DeleteProgram(program uint) {
	gl.gl.DeleteProgram(gl.programs.del(program))
	for _, value := range gl.uniformsByProgram[program] {
		gl.uniforms.del(value)
	}
	delete(gl.uniformsByProgram, program)
}

func (gl *WebGl) DeleteShader(shader uint) {
	gl.gl.DeleteShader(gl.shaders.del(shader))
}

func (gl *WebGl) DrawArrays(mode uint, first int, count int) {
	gl.gl.DrawArrays(int(mode), first, count)
}

func (gl *WebGl) Enable(cap uint) {
	gl.gl.Enable(int(cap))
}

func (gl *WebGl) EnableVertexAttribArray(index int) {
	gl.gl.EnableVertexAttribArray(index)
}

func (gl *WebGl) GenBuffers(n int) []uint {
	ids := make([]uint, n)

	for i := 0; i < n; i++ {
		ids[i] = gl.buffers.put(gl.gl.CreateBuffer())
	}

	return ids
}

func (gl *WebGl) GetAttribLocation(program uint, name string) int {
	return gl.gl.GetAttribLocation(gl.programs.get(program), name)
}

func (gl *WebGl) GetError() uint {
	return uint(gl.gl.GetError())
}

func (gl *WebGl) GetShaderInfoLog(shader uint) string {
	return gl.gl.GetShaderInfoLog(gl.shaders.get(shader))
}

func paramToInt(value js.Object) int {
	result := value.Int()

	if value.Str() == "true" {
		result = 1
	}

	return result
}

func (gl *WebGl) GetShaderParameter(shader uint, param uint) int {
	value := gl.gl.GetShaderParameter(gl.shaders.get(shader), int(param))

	return paramToInt(value)
}

func (gl *WebGl) GetProgramParameter(program uint, param uint) int {
	value := gl.gl.Call("getProgramParameter", gl.programs.get(program), int(param))

	return paramToInt(value)
}

func (gl *WebGl) GetUniformLocation(program uint, name string) int {
	uniform := gl.gl.GetUniformLocation(gl.programs.get(program), name)
	key := gl.uniforms.put(uniform)

	gl.uniformsByProgram[program] = append(gl.uniformsByProgram[program], key)

	return int(key)
}

func (gl *WebGl) LinkProgram(program uint) {
	gl.gl.LinkProgram(gl.programs.get(program))
}

func (gl *WebGl) ReadPixels(x int, y int, width int, height int, format uint, pixelType uint, pixels interface{}) {
	gl.gl.ReadPixels(x, y, width, height, format, pixelType, pixels)
}

func (gl *WebGl) ShaderSource(shader uint, source string) {
	gl.gl.ShaderSource(gl.shaders.get(shader), source)
}

func (gl *WebGl) UniformMatrix4fv(location int, transpose bool, value ...[]float32) {
	var args []interface{} = make([]interface{}, 0, 2+len(value))

	args = append(args, gl.uniforms.get(uint(location)), transpose)
	for _, entry := range value {
		args = append(args, entry)
	}
	gl.gl.Call("uniformMatrix4fv", args...)
}
func (gl *WebGl) UseProgram(program uint) {
	gl.gl.UseProgram(gl.programs.get(program))
}

func (gl *WebGl) VertexAttribOffset(index int, size int, attribType uint, normalized bool, stride int, offset int) {
	gl.gl.VertexAttribPointer(index, size, int(attribType), normalized, stride, offset)
}

func (gl *WebGl) Viewport(x int, y int, width int, height int) {
	gl.gl.Viewport(x, y, width, height)
}
