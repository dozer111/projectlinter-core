package painter

// Painter - abstraction
// Used in projectlinter.Rule(FailedMessage) to make the output more convenient and visually pretty
type Painter interface {
	Red(text string) string
	Green(text string) string
	Blue(text string) string
	White(text string) string
	Yellow(text string) string

	Warning(ruleName string) string
}
