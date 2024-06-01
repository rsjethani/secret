// Package secret provides types to guard your secret values from leaking into logs, std* etc.
//
// The objective is to disallow writing/serializing of secret values to std*, logs, JSON string
// etc. but provide access to the secret when requested explicitly.
package secret

// Text provides a way to safely store your secret value and a corresponding redact hint. This
// redact hint is what is used in operations like printing and serializing.
type Text struct {
	secret *string
	redact *string
}

// New returns [Text] for the secret with [FiveStar] as the default redact hint. Provide options
// like [RedactHint] to modify default behavior.
func New(secret string, options ...func(*Text)) Text {
	tx := Text{
		secret: new(string),
		redact: new(string),
	}

	*tx.secret = secret
	*tx.redact = FiveStar

	for _, o := range options {
		o(&tx)
	}

	return tx
}

// Some common redact hints.
const (
	FiveX    string = "XXXXX"
	FiveStar string = "*****"
	Redacted string = "[REDACTED]"
)

// RedactHint is a functional option to set r as the redact hint for the [Text]. You can use one of
// the common redact hints provided with this package like [FiveX] or provide your own string.
func RedactHint(r string) func(*Text) {
	return func(t *Text) {
		*t.redact = r
	}
}

// String implements the [fmt.Stringer] interface and returns only the redact hint. This prevents the
// secret value from being printed to std*, logs etc.
func (tx Text) String() string {
	if tx.redact == nil {
		return FiveStar
	}
	return *tx.redact
}

// Value gives you access to the actual secret value stored inside Text.
func (tx Text) Value() string {
	if tx.secret == nil {
		return ""
	}
	return *tx.secret
}

// MarshalText implements [encoding.TextMarshaler]. It marshals redact string into bytes rather than
// the actual secret value.
func (tx Text) MarshalText() ([]byte, error) {
	return []byte(tx.String()), nil
}

// UnmarshalText implements [encoding.TextUnmarshaler]. It unmarshals b into receiver's new secret
// value. If redact string is present then it is reused.
func (tx *Text) UnmarshalText(b []byte) error {
	s := string(b)

	// If the original redact is not nil then use it otherwise fallback to default.
	if tx.redact != nil {
		*tx = New(s, RedactHint(*tx.redact))
	} else {
		*tx = New(s)
	}
	return nil
}

// Equal returns true if both arguments have the same secret. The redact strings are not considered.
func Equal(tx1, tx2 Text) bool {
	return *tx1.secret == *tx2.secret
}
