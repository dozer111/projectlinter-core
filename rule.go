package projectlinter

// Rule - the main thing in all the project. It`s the one that actually checks for something.
//
// Examples:
//
//	file <fileName> exists
//	composer has <dependencyName> dependency
//	composer dependencies are correct
//	file .editorconfig is correct
//	...
type Rule interface {
	// ID to run/ignore via configs
	ID() string
	// Title to show in output
	Title() string
	// Validate here we write validation logic
	Validate()
	IsPassed() bool
	// FailedMessage - if IsPassed -false - use this method to show failed message to user
	// The signature is so, to make the output maximum flexible
	// You can write your own custom output or use one of predefined printers
	FailedMessage() []string
}
