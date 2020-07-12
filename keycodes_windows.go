package promptui

// source: https://msdn.microsoft.com/en-us/library/aa243025(v=vs.60).aspx

var (
	// KeyEnter is the default key for submission/selection inside a command line prompt.
	KeyEnter rune = 13

	// KeyBackspace is the default key for deleting input text inside a command line prompt.
	KeyBackspace rune = 8

	// FIXME: keys below are not triggered by readline, not working on Windows

	// KeyPrev is the default key to go up during selection inside a command line prompt.
	KeyPrev        rune = 87
	KeyPrevDisplay      = "w"

	// KeyNext is the default key to go down during selection inside a command line prompt.
	KeyNext        rune = 83
	KeyNextDisplay      = "s"

	// KeyBackward is the default key to page up during selection inside a command line prompt.
	KeyBackward        rune = 65
	KeyBackwardDisplay      = "a"

	// KeyForward is the default key to page down during selection inside a command line prompt.
	KeyForward        rune = 68
	KeyForwardDisplay      = "d"
)
