package secret_test

import (
	"encoding/json"
	"fmt"

	"github.com/rsjethani/secret/v2"
)

func ExampleNew() {
	s := secret.New("$ecre!")
	fmt.Println(s, s.Value())

	// Output: ***** $ecre!
}

func ExampleRedactAs() {
	s := secret.New("$ecre!", secret.RedactAs(secret.FiveX))
	fmt.Println(s, s.Value())

	s = secret.New("$ecre!", secret.RedactAs(secret.Redacted))
	fmt.Println(s, s.Value())

	s = secret.New("$ecre!", secret.RedactAs("my redact hint"))
	fmt.Println(s, s.Value())

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
	fmt.Println(login.Password.Value())

	// Output:
	// {User:John Password:*****}
	// $ecre!
}

func ExampleEqual() {
	tx1 := secret.New("hello")
	tx2 := secret.New("hello", secret.RedactAs(secret.Redacted))
	tx3 := secret.New("world")
	fmt.Println(secret.Equal(tx1, tx2))
	fmt.Println(secret.Equal(tx1, tx3))

	// Output:
	// true
	// false
}
