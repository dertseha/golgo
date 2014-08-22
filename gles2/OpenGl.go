/*
 * The gles2 package contains the constants and interface compatible with
 * OpenGL ES2.
 *
 * All values are based on the public header from the Khronos group:
 * http://www.khronos.org/registry/gles/
 */
package gles2

type OpenGl interface {
	Clear(mask uint)
	ClearColor(red float32, green float32, blue float32, alpha float32)
}
