//go:build js && wasm

package main

import (
	"strings"
	"syscall/js"
)

func runCommand(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return "no command provided"
	}

	input := strings.TrimSpace(args[0].String())

	switch input {
	case "help":
		return "Available commands: help, about, projects, clear"
	case "about":
		return "Hi, I'm Trevor. I build backend systems, Go projects, and data tools."
	case "projects":
		return "Projects: fightdb, Odoo tools, portfolio site"
	case "clear":
		return "__CLEAR__"
	case "":
		return ""
	default:
		return "command not found: " + input
	}
}

func main() {
	js.Global().Set("runCommand", js.FuncOf(runCommand))

	select {}
}
