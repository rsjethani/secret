// Package secret provides types to guard your secret values from leaking into logs, std* etc.
//
// The objective is to disallow writing/serializing of secret values to std*, logs, JSON string
// etc. but provide access to the secret when requested explicitly.
package secret

// Text provides a way to safely store your secret value and a corresponding redact hint. This
// redact hint is what is used in operations like printing and serializing.
type Text struct {
	// v is the actual secret values.
	v *string
	// r is the redact hint to be used in place of secret.
	r *string
}

// New returns [Text] for the secret with [FiveStar] as the default redact hint. Provide options
// like [RedactHint] to modify default behavior.
func New(secret string, options ...func(*Text)) Text {
	tx := Text{
		v: new(string),
		r: new(string),
	}

	*tx.v = secret
	*tx.r = FiveStar

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
		*t.r = r
	}
}

// String implements the [fmt.Stringer] interface and returns only the redact hint. This prevents the
// secret value from being printed to std*, logs etc.
func (tx Text) String() string {
	if tx.r == nil {
		return FiveStar
	}
	return *tx.r
}

// Value gives you access to the actual secret value stored inside Text.
func (tx Text) Value() string {
	if tx.v == nil {
		return ""
	}
	return *tx.v
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
	if tx.r != nil {
		*tx = New(s, RedactHint(*tx.r))
	} else {
		*tx = New(s)
	}
	return nil
}

// Equals checks whether s2 has same secret string or not.
func (tx *Text) Equals(s2 Text) bool {
	return *tx.v == *s2.v
}
