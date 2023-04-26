// Package secret provides types to guard your secret values from leaking into logs, std* etc.
//
// The objective is to disallow writing/serializing of secret values to std*, logs, JSON string
// etc. but provide access to the secret when requested explicitly.
package secret

import (
	"encoding/json"
	"fmt"
)

// DefaultRedact is used by default if no other redact hint is given.
const DefaultRedact string = "*****"

// Text provides a way to safely store your secret value and a corresponding redact hint. This
// redact hint what is used in operations like printing and serializing. The default
// value of Text is usable.
type Text struct {
	// v is the actual secret values.
	v *string
	// r is the redact hint to be used in place of secret.
	r *string
}

// NewText creates a new Text instance with s as the secret value. Multiple option functions can
// be passed to alter default behavior.
func NewText(s string, options ...func(*Text)) Text {
	sec := Text{
		v: new(string),
		r: new(string),
	}

	*sec.v = s
	*sec.r = DefaultRedact

	// Apply options to override defaults
	for _, opt := range options {
		opt(&sec)
	}

	return sec
}

// String implements the fmt.Stringer interface and returns only the redact hint. This prevents the
// secret value from being printed to std*, logs etc.
func (s Text) String() string {
	if s.r == nil {
		return DefaultRedact
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

func (s Text) MarshalText() ([]byte, error) {
	return []byte(*s.r), nil
}

func (s *Text) UnmarshalText(b []byte) error {
	v := string(b)

	// If the original redact is not nil then use it otherwise fallback to default.
	if s.r != nil {
		*s = NewText(v, CustomRedact(*s.r))
	} else {
		*s = NewText(v)
	}
	return nil
}

// MarshalJSON allows Text to be serialized into a JSON string. Only the redact hint is part of the
// the JSON string.
func (s Text) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, *s.r)), nil
}

// UnmarshalJSON allows a JSON string to be deserialized into a Text value. DefaultRedact is set
// as the redact hint.
func (s *Text) UnmarshalJSON(b []byte) error {
	// Get the new secret value from unmarshalled data.
	var n string
	if err := json.Unmarshal(b, &n); err != nil {
		return err
	}

	// If the original redact is not nil then use it otherwise fallback to default.
	if s.r != nil {
		*s = NewText(n, CustomRedact(*s.r))
	} else {
		*s = NewText(n)
	}

	return nil
}

// Equals checks whether s2 has same secret string or not.
func (s *Text) Equals(s2 Text) bool {
	return *s.v == *s2.v
}

// Redacted option sets "[REDACTED]" as the redact hint.
func Redacted(s *Text) {
	*s.r = "[REDACTED]"
}

// FiveXs option sets "XXXXX" as the redact hint.
func FiveXs(s *Text) {
	*s.r = "XXXXX"
}

// CustomRedact option sets the value of r as the redact hint.
func CustomRedact(r string) func(*Text) {
	return func(s *Text) {
		*s.r = r
	}
}
