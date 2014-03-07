package shaders

import (
	"log"

	gl "github.com/remogatto/opengles2"
)

type VertexShader string
type FragmentShader string

func checkShaderCompileStatus(shader uint32) {
	var stat int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &stat)
	if stat == 0 {
		var length int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &length)
		infoLog := gl.GetShaderInfoLog(shader, gl.Sizei(length), nil)
		if len(infoLog) > 0 {
			log.Fatalf("Compile error in shader %d: \"%s\"\n", shader, infoLog[:len(infoLog)])
		}
	}
}

func checkProgramLinkStatus(pid uint32) {
	var stat int32
	gl.GetProgramiv(pid, gl.LINK_STATUS, &stat)
	if stat == 0 {
		var length int32
		gl.GetProgramiv(pid, gl.INFO_LOG_LENGTH, &length)
		infoLog := gl.GetProgramInfoLog(pid, gl.Sizei(length), nil)
		if len(infoLog) > 0 {
			log.Fatalf("Link error in program %d: \"%s\"\n", pid, infoLog[:len(infoLog)])
		}
	}
}

func compileShader(typeOfShader gl.Enum, source string) uint32 {
	if shader := gl.CreateShader(typeOfShader); shader != 0 {
		gl.ShaderSource(shader, 1, &source, nil)
		gl.CompileShader(shader)
		checkShaderCompileStatus(shader)
		return shader
	}
	return 0
}

func (s VertexShader) Compile() uint32 {
	shaderId := compileShader(gl.VERTEX_SHADER, (string)(s))
	return shaderId
}

func (s FragmentShader) Compile() uint32 {
	shaderId := compileShader(gl.FRAGMENT_SHADER, (string)(s))
	return shaderId
}

type Program uint32

func NewProgram(fsh FragmentShader, vsh VertexShader) Program {
	pid := gl.CreateProgram()
	gl.AttachShader(pid, fsh.Compile())
	gl.AttachShader(pid, vsh.Compile())
	gl.LinkProgram(pid)
	checkProgramLinkStatus(pid)
	return Program(pid)
}

func (p Program) Use() {
	gl.UseProgram(uint32(p))
}

func (p Program) GetAttribute(name string) uint32 {
	return gl.GetAttribLocation(uint32(p), name)
}

func (p Program) GetUniform(name string) uint32 {
	return gl.GetUniformLocation(uint32(p), name)
}
