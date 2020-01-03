package log4go


import (
	"fmt"
)


type colorClass struct {

}

var Color = colorClass{}

var colorFormat = "\x1b[%dm%s\x1b[0m"

func colorize(colorCode int, s string) string {
	return fmt.Sprintf(colorFormat, colorCode, s)
}

// ColorBlue represents the blue color.
var ColorBlue = 34

// Blue returns the string contents of "s" wrapped by blue foreground color.
func (this *colorClass) Blue(s string) string {
	return colorize(ColorBlue, s)
}

// ColorLightGreen represents a light green color.
var ColorLightGreen = 36

// LightGreen returns the string contents of "s" wrapped by a light green foreground color.
func (this *colorClass) LightGreen(s string) string {
	return colorize(ColorLightGreen, s)
}

// ColorPurple represents the purple color.
var ColorPurple = 35

// Purple returns the string contents of "s" wrapped by purple foreground color.
func (this *colorClass) Purple(s string) string {
	return colorize(ColorPurple, s)
}

// ColorWhite represents the white color.
var ColorWhite = 0

// White returns the string contents of "s" wrapped by white foreground color.
func (this *colorClass) White(s string) string {
	return colorize(ColorWhite, s)
}

// ColorGray represents the gray color.
var ColorGray = 37

// Gray returns the string contents of "s" wrapped by gray foreground color.
func (this *colorClass) Gray(s string) string {
	return colorize(ColorGray, s)
}

// ColorRed represents the red color.
var ColorRed = 31

// Red returns the string contents of "s" wrapped by red foreground color.
func (this *colorClass) Red(s string) string {
	return colorize(ColorRed, s)
}

// ColorRedBackground represents a white foreground color and red background.
var ColorRedBackground = 41

// RedBackground returns the string contents of "s" wrapped by white foreground color and red background.
func (this *colorClass) RedBackground(s string) string {
	return colorize(ColorRedBackground, s)
}

// ColorGreen represents the green color.
var ColorGreen = 32

// Green returns the string contents of "s" wrapped by green foreground color.
func (this *colorClass) Green(s string) string {
	return colorize(ColorGreen, s)
}

// ColorYellow represents the yellow color.
var ColorYellow = 33

// Yellow returns the string contents of "s" wrapped by yellow foreground color.
func (this *colorClass) Yellow(s string) string {
	return colorize(ColorYellow, s)
}


func ColorizedLevelMsg(level Level, msg string) string {
	return [...]string{
		Color.Green(msg), // FNST
		Color.Green(msg), // FINE
		Color.LightGreen(msg), // DEBG
		Color.Purple(msg), // TRAC
		Color.Green(msg), // INFO
		Color.Yellow(msg), // WARN
		Color.Red(msg), // EROR
		Color.Yellow(msg), // CRIT
	}[level]
}
