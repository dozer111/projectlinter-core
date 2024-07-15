package painter

// fakePainter - for test purposes
//
// Tests of printers is much more easy and understandable when you compare text with more verbose color codes
type fakePainter struct {
}

var _ Painter = (*fakePainter)(nil)

func (p fakePainter) Red(text string) string {
	return "(red)" + text
}

func (p fakePainter) Green(text string) string {
	return "(green)" + text
}

func (p fakePainter) Blue(text string) string {
	return "(blue)" + text
}

func (p fakePainter) White(text string) string {
	return "(white)" + text
}

func (p fakePainter) Yellow(text string) string {
	return "(yellow)" + text
}

func (p fakePainter) Warning(text string) string {
	return "(warn)" + text
}
