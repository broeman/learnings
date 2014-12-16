// Using GoGL and GLFW for rendering a window
// Unfortunately Ubuntu doesn't have GLFW3 in repos, so I'll revisit this later
// Author: Jesper Brodersen 2014

package main

import (
	"fmt"
	gl "github.com/chsc/gogl/gl43"
	"github.com/go-gl/glfw"
	"os"
	"runtime"
)

func main() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.Terminate()

	glfw.OpenWindowHint(glfw.WindowNoResize, 1)

	if err := glfw.OpenWindow(640, 480, 0, 0, 0, 0, 16, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle("GoGL")

	if err := gl.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "GoGL: %s\n", err)
		return
	}
}
