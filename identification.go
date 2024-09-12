package main

import "strings"

type FileIdentification struct {
	ansi string
	symbol string
	name string
}

// Returns string array with Ansi code, and Nerd Fonts character.
func Identify(name string) FileIdentification {
	name = strings.ToLower(name)
	var extension string = strings.Split(name, ".")[len(strings.Split(name, "."))-1]
	var id FileIdentification = FileIdentification{Ansi("White"), "\uf15b", "File"}
	switch extension {
		case "go":
			id = FileIdentification{Ansi("Cyan"), "=>Go", "Go"}
		case "exe":
			id = FileIdentification{Ansi("Red"), ".EXE", "Windows Executable"}
		case "mod":
			if name == "go.mod" {
				id = FileIdentification{Ansi("LightCyan"), "=>Go", "Go Module"}
			}
		case "sum":
			if name == "go.sum" {
				id = FileIdentification{Ansi("LightCyan"), "=>Go", "Go Checksum"}
			}
		case "js":
			id = FileIdentification{Ansi("Yellow"), "[JS]", "Javascript"}
		case "html", "htm", "xhtml":
			id = FileIdentification{Ansi("LightRed"), "{5} ", "HTML"}
		case "css", "scss", "sass":
			id = FileIdentification{Ansi("LightBlue"), "{3} ", "CSS"}
		case "md", "markdown":
			if name == "readme.md" {
				id = FileIdentification{Ansi("LightBlue"), "(i) ", "Readme/Info File"}
			} else {
				id = FileIdentification{Ansi("Cyan"), "<M↓>", "Markdown"}
			}
	}

	return id
}
