package enums

type ErrorSeverity int

const (
	// Transaction was unsuccessful
	Error = iota

	// Error	Message not processed due to application or network failure condition
	Fatal

	// Transaction was successful but includes information e.g., inform patient
	Information

	// Transaction successful, but there may issues
	Warning
)

var errorSeverityStrings = map[ErrorSeverity]string{
	Error:       "E",
	Fatal:       "F",
	Information: "I",
	Warning:     "W",
}

func (e ErrorSeverity) String() string {
	return errorSeverityStrings[e]
}
