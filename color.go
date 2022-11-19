package todo

import "fmt"

const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
)

func red(text string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, text, ColorDefault)
}

func blue(text string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, text, ColorDefault)
}

func green(text string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, text, ColorDefault)
}

func gray(text string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, text, ColorDefault)
}
