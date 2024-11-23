package enums

type MessageErrorConditionCodes int

const (

	// Message accepted Success. Optional, as the AA conveys success. Used for
	// systems that must always return a status code.
	MessageAcceptedSuccess = iota

	// Error Codes

	// Segment sequence error.
	// Error: The message segments were not in the proper order, or required
	// segments are missing.
	SegmentSequenceError

	// Required field missing.
	// Error: A required field is missing from a segment
	RequiredFieldMissing

	// Data type error.
	// Error: The field contained data of the wrong data type, e.g., an NM
	// field contained "FOO".
	DataTypeError

	// Table value not found.
	// Error: A field of data type ID or IS was compared against the corresponding
	// table, and no match was found.
	TableValueNotFound

	// Rejection Codes

	// Unsupported message type.
	// Rejection: The Message Type is not supported.
	UnsupportedMessageType

	// Unsupported event code.
	// Rejection: The Event Code is not supported.
	UnsupportedEventCode

	// Unsupported processing id.
	// Rejection: The Processing ID is not supported.
	UnsupportedProcessingId

	// Unsupported version id.
	// Rejection:  The Version ID is not supported.
	UnsupportedVersionId

	// Unknown key identifier.
	// Rejection: The ID of the patient, order, etc., was not found.
	// Used for transactions other than additions, e.g., transfer of a
	// non-existent patient.
	UnknownKeyIdentifier

	// Duplicate key identifier.
	// Rejection: The ID of the patient, order, etc., already exists.
	// Used in response to addition transactions (Admit, New Order, etc.).
	DuplicateKeyIdentifier

	// Application record locked.
	// Rejection: The transaction could not be performed at the application
	// storage level, e.g., database locked.
	ApplicationRecordLocked

	// Application internal error.
	// Rejection: A catchall for internal errors not explicitly covered by
	// other codes.
	ApplicationInternalError
)

var messageErrorConditionCodes = map[MessageErrorConditionCodes]string{

	MessageAcceptedSuccess:   "0",
	SegmentSequenceError:     "100",
	RequiredFieldMissing:     "101",
	DataTypeError:            "102",
	TableValueNotFound:       "103",
	UnsupportedMessageType:   "200",
	UnsupportedEventCode:     "201",
	UnsupportedProcessingId:  "202",
	UnsupportedVersionId:     "203",
	UnknownKeyIdentifier:     "204",
	DuplicateKeyIdentifier:   "205",
	ApplicationRecordLocked:  "206",
	ApplicationInternalError: "207",
}

func (e MessageErrorConditionCodes) String() string {
	return messageErrorConditionCodes[e]
}
