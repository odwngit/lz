package main

func Ansi(color string) string {
	switch color {
		case "Reset":
			return "\033[0m"
		case "Black":
			return "\033[0;30m"
		case "Red":
			return "\033[0;31m"
		case "Green":
			return "\033[0;32m"
		case "Yellow":
			return "\033[0;33m"
		case "Blue":
			return "\033[0;34m"
		case "Purple":
			return "\033[0;35m"
		case "Cyan":
			return "\033[0;36m"
		case "White":
			return "\033[0;37m"
		case "Gray":
			return "\033[0;90m"
		case "LightRed":
			return "\033[0;91m"
		case "LightGreen":
			return "\033[0;92m"
		case "LightYellow":
			return "\033[0;93m"
		case "LightBlue":
			return "\033[0;94m"
		case "Pink":
			return "\033[0;95m"
		case "LightCyan":
			return "\033[0;96m"
		default:
			return "\033[0m"
	}
}
