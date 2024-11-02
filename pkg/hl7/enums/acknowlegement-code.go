package enums

type AcknowledgementCode int

const (

	/** Original mode: Application Accept - Enhanced mode: Application acknowledgment: Accept */
	ApplicationAccept = iota

	/** Original mode: Application Error - Enhanced mode: Application acknowledgment: Error */
	ApplicationError

	/** Original mode: Application Reject - Enhanced mode: Application acknowledgment: Reject */
	ApplicationReject

	/** Enhanced mode: Accept acknowledgment: Commit Accept */
	CommitAccept

	/** Enhanced mode: Accept acknowledgment: Commit Error */
	CommitError

	/** Enhanced mode: Accept acknowledgment: Commit Reject */
	CommitReject
)

var acknowledgementCodeStrings = map[AcknowledgementCode]string{
	ApplicationAccept: "AA",
	ApplicationError:  "AE",
	ApplicationReject: "AR",
	CommitAccept:      "CA",
	CommitError:       "CE",
	CommitReject:      "CR",
}

func (a AcknowledgementCode) String() string {
	return acknowledgementCodeStrings[a]
}
