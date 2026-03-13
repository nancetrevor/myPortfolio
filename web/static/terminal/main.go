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
	parts := strings.Fields(input)

	switch parts[0] {
	case "help":
		return "Available commands: help, about, projects, clear, goto"
	case "about":
		return "Hi, I'm Trevor. I build Full-Stack applications, Go projects, and data intensive tools."
	case "projects":
		return "Projects: fightdb, Snip, Odn\nTry goto {project name} to get re-directed to the project."
	case "goto":
		if len(parts) < 2 {
			return "goto Commands: github, linkedin, fightdb, snip, odn\nUsage: goto {name}"
		}
		switch parts[1] {
		case "fightdb":
			return "__NAVIGATE__:https://fightdb.org"
		case "github":
			return "__NAVIGATE__:https:https://github.com/nancetrevor"
		case "linkedin":
			return "__NAVIGATE__:https://www.linkedin.com/in/trevor-nance/"
		case "snip":
			return "__NAVIGATE__:https://github.com/nancetrevor/Snip"
		case "odn":
			return "__NAVIGATE__:https://github.com/nancetrevor/cli_journal"
		default:
			return "site not found."
		}
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
