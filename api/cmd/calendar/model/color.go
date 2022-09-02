package model

type Color string

const (
	RED    Color = "red"
	BLUE   Color = "blue"
	YELLOW Color = "yellow"
	GREEN  Color = "green"
)

func ConvertToColor(c string) Color {
	switch Color(c) {
	case RED:
		return RED
	case BLUE:
		return BLUE
	case YELLOW:
		return YELLOW
	case GREEN:
		return GREEN
	}
	return Color("")
}