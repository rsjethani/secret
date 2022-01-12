package secret

import (
	"encoding/json"
	"testing"
)

func TestText_secret_and_its_copy_share_the_same_data(t *testing.T) {
	s1 := NewText("hello")
	s2 := s1

	if s2.r != s1.r {
		t.Fatal("redact field of the secret and its copy do not point to same string")
	}

	if s2.v != s1.v {
		t.Fatal("value field of the secret and its copy do not point to same string")
	}
}

func TestText_UnmarshalJSON_allocates_new_data_rather_than_overwiting_existing(t *testing.T) {
	s1 := NewText("hello")

	oldRedact := s1.r
	oldValue := s1.v

	err := json.Unmarshal([]byte(`"hell"`), &s1)
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
