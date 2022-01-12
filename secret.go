package secret

import (
	"encoding/json"
	"fmt"
)

const DefaultRedact string = "*****"

type Text struct {
	v *string
	r *string
}

func NewText(s string, options ...func(*Text)) Text {
	sec := Text{}

	sec.init()
	*sec.v = s

	for _, opt := range options {
		opt(&sec)
	}
	return sec
}

func (s *Text) init() {
	s.v = new(string)
	s.r = new(string)
	*s.r = DefaultRedact
}

func CustomRedact(r string) func(*Text) {
	return func(s *Text) {
		*s.r = r
	}
}

func (s Text) String() string {
	return *s.r
}

func (s Text) Value() string {
	return *s.v
}

func (s Text) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, *s.r)), nil
}

func (s *Text) UnmarshalJSON(b []byte) error {
	s.init()
	return json.Unmarshal(b, s.v)
}

func Redacted(s *Text) {
	*s.r = "[REDACTED]"
}

func FiveXs(s *Text) {
	*s.r = "XXXXX"
}
