package main

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const Height = 600
const Width = 800

type HelloTriangleApplication struct {
	window *glfw.Window
}

func (h *HelloTriangleApplication) cleanup() error {
	return nil
}

func (h *HelloTriangleApplication) initVulkan() error {
	return nil
}

func (h *HelloTriangleApplication) initWindow() error {
	err := glfw.Init()

	if err != nil {
		return err
	}

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	h.window, err = glfw.CreateWindow(Width, Height, "Vulkan", nil, nil)

	if err != nil {
		return err
	}

	return nil
}

func (h *HelloTriangleApplication) mainLoop() {
	// h.window.MakeContextCurrent()

	for !h.window.ShouldClose() {
		// h.window.SwapBuffers()
		glfw.PollEvents()
	}
}

func (h *HelloTriangleApplication) Run() error {
	defer h.cleanup()

	err := h.initWindow()

	if err != nil {
		return err
	}

	err = h.initVulkan()

	if err != nil {
		return err
	}

	h.mainLoop()

	return nil
}

// func init() {
// 	// This is needed to arrange that main() runs on main thread.
// 	// See documentation for functions that are only allowed to be called from the main thread.
// 	runtime.LockOSThread()
// }

func main() {
	app := HelloTriangleApplication{}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
