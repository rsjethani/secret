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

// New returns [Text] for the secret with [FiveStar] as the default redact hint.
func New(secret string) Text {
	sec := Text{
		v: new(string),
		r: new(string),
	}

	*sec.v = secret
	*sec.r = FiveStar

	return sec
}

// Some common redact hints.
const (
	FiveX    string = "XXXXX"
	FiveStar string = "*****"
	Redacted string = "[REDACTED]"
)

// WithRedact returns a copy of [Text] but with r as the redact string.
// Some common redact hints like [FiveX] etc. are provided for convenience.
func (tx Text) WithRedact(r string) Text {
	tx2 := tx
	*tx2.r = r
	return tx2
}

// String implements the [fmt.Stringer] interface and returns only the redact hint. This prevents the
// secret value from being printed to std*, logs etc.
func (s Text) String() string {
	if s.r == nil {
		return FiveStar
	}
	return *s.r
}

// Value gives you access to the actual secret value stored inside Text.
func (s Text) Value() string {
	if s.v == nil {
		return ""
	}
	return *s.v
}

// MarshalText implements [encoding.TextMarshaler]. It marshals redact string into bytes rather than
// the actual secret value.
func (s Text) MarshalText() ([]byte, error) {
	return []byte(*s.r), nil
}

// UnmarshalText implements [encoding.TextUnmarshaler]. It unmarshals b into receiver's new secret
// value. If redact string is present then it is reused.
func (tx *Text) UnmarshalText(b []byte) error {
	s := string(b)

	// If the original redact is not nil then use it otherwise fallback to default.
	if tx.r != nil {
		*tx = New(s).WithRedact(*tx.r)
	} else {
		*tx = New(s)
	}
	return nil
}

// Equals checks whether s2 has same secret string or not.
func (s *Text) Equals(s2 Text) bool {
	return *s.v == *s2.v
}
