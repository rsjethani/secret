package secret

import (
	"encoding/json"
	"fmt"
)

const DefaultRedact string = "*****"

type Secret struct {
	v *string
	r *string
}

func New(s string, options ...func(*Secret)) Secret {
	sec := Secret{}

	sec.init()
	*sec.v = s

	for _, opt := range options {
		opt(&sec)
	}
	return sec
}

func (s *Secret) init() {
	s.v = new(string)
	s.r = new(string)
	*s.r = DefaultRedact
}

func CustomRedact(r string) func(*Secret) {
	return func(s *Secret) {
		*s.r = r
	}
}

func (s Secret) String() string {
	return *s.r
}

func (s Secret) Value() string {
	return *s.v
}

func (s Secret) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, *s.r)), nil
}

func (s *Secret) UnmarshalJSON(b []byte) error {
	s.init()
	return json.Unmarshal(b, s.v)
}

func Redacted(s *Secret) {
	*s.r = "[REDACTED]"
}

func FiveXs(s *Secret) {
	*s.r = "XXXXX"
}

func (s *Secret) Copy() Secret {
	return New(*s.v, CustomRedact(*s.r))
}
