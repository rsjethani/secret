package secret_test

import (
	"encoding/json"
	"fmt"

	"github.com/rsjethani/secret/v3"
)

func ExampleNew() {
	s := secret.New("$ecre!")
	fmt.Println(s, s.Secret())

	// Output: ***** $ecre!
}

func ExampleRedactAs() {
	s := secret.New("$ecre!", secret.RedactAs(secret.FiveX))
	fmt.Println(s, s.Secret())

	s = secret.New("$ecre!", secret.RedactAs(secret.Redacted))
	fmt.Println(s, s.Secret())

	s = secret.New("$ecre!", secret.RedactAs("my redact hint"))
	fmt.Println(s, s.Secret())

	// Output:
	// XXXXX $ecre!
	// [REDACTED] $ecre!
	// my redact hint $ecre!
}

func ExampleText_MarshalText() {
	login := struct {
		User     string
		Password secret.Text
	}{
		User:     "John",
		Password: secret.New("shh!"),
	}

	bytes, err := json.Marshal(&login)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	// Output: {"User":"John","Password":"*****"}
}

func ExampleText_UnmarshalText() {
	login := struct {
		User     string
		Password secret.Text
	}{}

	err := json.Unmarshal([]byte(`{"User":"John","Password":"$ecre!"}`), &login)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", login)
	fmt.Println(login.Password.Secret())

	// Output:
	// {User:John Password:*****}
	// $ecre!
}

func ExampleEqual() {
	// Empty Texts are equal.
	fmt.Println(secret.Equal(secret.Text{}, secret.Text{}))

	// Initialsed Text is not equal to an empty one.
	fmt.Println(secret.Equal(secret.New("hello"), secret.Text{}))

	// Texts with different secret strings are not equal.
	fmt.Println(secret.Equal(secret.New("hello"), secret.New("world")))

	// Texts with different redact strings but same secret string are equal.
	fmt.Println(secret.Equal(secret.New("hello"), secret.New("hello", secret.RedactAs(secret.FiveX))))

	// Output:
	// true
	// false
	// false
	// true
}
