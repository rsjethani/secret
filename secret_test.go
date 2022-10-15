package secret

import (
	"encoding/json"
	"testing"
)

func TestTextUnmarshalJSON_allocates_new_data_rather_than_overwriting_existing(t *testing.T) {
	s1 := NewText("hello")

	oldRedact := s1.r
	oldValue := s1.v

	err := json.Unmarshal([]byte(`"hello"`), &s1)
	if err != nil {
		t.Fatal(err)
	}

	if s1.r == oldRedact {
		t.Fatal("UnmarshalJSON did not allocate new redact string instead it overwrote exitsing")
	}

	if s1.v == oldValue {
		t.Fatal("UnmarshalJSON did not allocate new value string instead it overwrote exitsing")
	}
}
