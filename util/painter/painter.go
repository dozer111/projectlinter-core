package painter

// Painter - abstraction
// Used in projectlinter.Rule(FailedMessage) to make the output more convenient and visually pretty
type Painter interface {
	Red(text string) string
	Green(text string) string
	Blue(text string) string
	White(text string) string
	Yellow(text string) string

	Warning(text string) string
}

type PainterState string

const (
	BashCLI PainterState = "bash"
	Fake    PainterState = "fake"
)

var CurrentPainterState PainterState = BashCLI

// NewPainter - factory method
// Create a painter according to your needs
//
// The idea is: you can change the painter realization from outside by changing the CurrentPainterState
func NewPainter() Painter {
	switch CurrentPainterState {
	case BashCLI:
		return &bashPainter{}
	case Fake:
		return &fakePainter{}
	default:
		return &bashPainter{}
	}
}
