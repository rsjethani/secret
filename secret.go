package secret

import "fmt"

const DefaultRedact string = "*****"

type Secret struct {
	v *string
	r string
}

func New(s string, options ...func(*Secret)) Secret {
	sec := Secret{
		v: new(string),
		r: DefaultRedact,
	}

	*sec.v = s

	for _, opt := range options {
		opt(&sec)
	}
	return sec
}

func CustomRedact(r string) func(*Secret) {
	return func(s *Secret) {
		s.r = r
	}
}

func (s Secret) String() string {
	return s.r
}

func (s Secret) Value() string {
	return *s.v
}

func (s Secret) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, s.r)), nil
}

func Redacted(s *Secret) {
	s.r = "[REDACTED]"
}

func FiveXs(s *Secret) {
	s.r = "XXXXX"
}
