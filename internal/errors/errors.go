package errors

import "github.com/TylerHaigh/go-simple-hl7/pkg/hl7/enums"

func MapErrorCode(code enums.MessageErrorConditionCodes) enums.AcknowledgementCode {
	switch code {
	case enums.MessageAcceptedSuccess:
		return enums.ApplicationAccept

	case enums.SegmentSequenceError:
		fallthrough
	case enums.RequiredFieldMissing:
		fallthrough
	case enums.DataTypeError:
		fallthrough
	case enums.TableValueNotFound:
		return enums.ApplicationError

	case enums.UnsupportedMessageType:
		fallthrough
	case enums.UnsupportedEventCode:
		fallthrough
	case enums.UnsupportedProcessingId:
		fallthrough
	case enums.UnsupportedVersionId:
		fallthrough
	case enums.UnknownKeyIdentifier:
		fallthrough
	case enums.DuplicateKeyIdentifier:
		fallthrough
	case enums.ApplicationRecordLocked:
		fallthrough
	case enums.ApplicationInternalError:
		fallthrough

	default:
		return enums.ApplicationReject
	}
}

func MapErrorCodeDescription(code enums.MessageErrorConditionCodes) string {
	switch code {

	case enums.MessageAcceptedSuccess:
		return "Message accepted Success"

	case enums.SegmentSequenceError:
		return "Segment sequence error"
	case enums.RequiredFieldMissing:
		return "Required field missing"
	case enums.DataTypeError:
		return "Data type error"
	case enums.TableValueNotFound:
		return "Table value not found"

	case enums.UnsupportedMessageType:
		return "Unsupported message type"
	case enums.UnsupportedEventCode:
		return "Unsupported event code"
	case enums.UnsupportedProcessingId:
		return "Unsupported processing id"
	case enums.UnsupportedVersionId:
		return "Unsupported version id"
	case enums.UnknownKeyIdentifier:
		return "Unknown key identifier"
	case enums.DuplicateKeyIdentifier:
		return "Duplicate key identifier"
	case enums.ApplicationRecordLocked:
		return "Application record locked"
	case enums.ApplicationInternalError:
		return "Application internal error"

	default:
		return "Application internal error"

	}
}
