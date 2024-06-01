package secret

// Text provides a way to safely store your secret string along with a proxy/redact string. This
// redact string is what is used in operations like printing and serializing there by avoiding
// leaking the secret. Once created, the instance is readonly except for the [Text.UnmarshalText]
// operation, but that too only modifies the local copy. Hence the type is concurrent safe.
type Text struct {
	secret *string
	redact *string
}

// RedactAs is a functional option to set r as the redact string for [Text]. You can use one of
// the common redact strings provided with this package like [FiveX] or provide your own.
func RedactAs(r string) func(*Text) {
	return func(t *Text) {
		*t.redact = r
	}
}

// New returns [Text] for the secret with [FiveStar] as the default redact string. Provide options
// like [RedactAs] to modify default behavior.
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

// String implements the [fmt.Stringer] interface and returns only the redact string. This prevents
// the actual secret string from being sent to std*, logs etc.
func (tx Text) String() string {
	if tx.redact != nil {
		return *tx.redact
	}
	return FiveStar
}

// Value gives you access to the secret string stored inside [Text].
func (tx Text) Value() string {
	if tx.secret != nil {
		return *tx.secret
	}
	return ""
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
		*tx = New(s, RedactAs(*tx.redact))
	} else {
		*tx = New(s)
	}
	return nil
}

// Equal returns true if both arguments have the same secret. The redact strings are not considered.
func Equal(tx1, tx2 Text) bool {
	return *tx1.secret == *tx2.secret
}
