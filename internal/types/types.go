package types

type FileType string

const (
	EMPTY_FILE = "EMPTY_FILE"
)

type COLOR string

const (
	ColorReset   COLOR = "\033[0m"
	ColorRed     COLOR = "\033[31m"
	ColorGreen   COLOR = "\033[32m"
	ColorYellow  COLOR = "\033[33m"
	ColorBlue    COLOR = "\033[34m"
	ColorMagneta COLOR = "\033[35m"
	ColorCyan    COLOR = "\033[36m"
	ColorWhite   COLOR = "\033[37m"
	BrightBlue   COLOR = "\033[94m"
	BoldBlue     COLOR = "\033[1;34m"
	BlueOnYellow COLOR = "\033[34;43m"
)
