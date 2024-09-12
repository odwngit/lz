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
	var id FileIdentification = FileIdentification{Ansi("White"), "    ", "File"}
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
			id = FileIdentification{Ansi("LightYellow"), "[JS]", "Javascript"}
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
		case "bat", "cmd", "sh":
			id = FileIdentification{Ansi("LightGreen"), "[>_]", "Command-Line Script"}
		case "c":
			id = FileIdentification{Ansi("LightBlue"), "(C) ", "C"}
		case "h", "hpp", "hh":
			id = FileIdentification{Ansi("Cyan"), "(h) ", "C Header File"}
		case "cpp", "cc":
			id = FileIdentification{Ansi("LightBlue"), "*C++", "C++"}
		case "json":
			id = FileIdentification{Ansi("Yellow"), "{;} ", "JSON"}
		case "py", "pyc", "ipynb":
			id = FileIdentification{Ansi("LightBlue"), "&"+Ansi("Cyan")+"p"+Ansi("LightYellow")+"y"+Ansi("Yellow")+">", "Python"}
		case "txt", "text":
			id = FileIdentification{Ansi("White"), "... ", "Raw Text File"}
		case "rb":
			id = FileIdentification{Ansi("Red"), "/rb>", "Ruby"}
		case "log":
			id = FileIdentification{Ansi("White"), "LOG"+Ansi("Red")+"!", "Log File"}
		case "lua":
			id = FileIdentification{Ansi("Blue"), "oO()", "Lua Script"}
		case "nvim":
			id = FileIdentification{Ansi("LightBlue"), "|"+Ansi("LightGreen")+"\\"+Ansi("Green")+"| ", "Neovim Config"}
		case "gz", "zip", "7z", "rar":
			id = FileIdentification{Ansi("Red"), "0"+Ansi("Green")+"0"+Ansi("LightBlue")+"0 ", "Compressed Archive"}
		case "lnk":
			id = FileIdentification{Ansi("LightCyan"), "\\_/>", "Shortcut"}
		case "cfg", "config", "ini":
			id = FileIdentification{Ansi("Gray"), "{%} ", "Configuration File"}
		case "pdf":
			id = FileIdentification{Ansi("LightRed"), "(&) ", "PDF Document"}
	}

	return id
}
