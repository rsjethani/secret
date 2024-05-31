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

func ExampleText_WithRedact() {
	s := secret.New("$ecre!").WithRedact(secret.FiveX)
	fmt.Println(s, s.Value())

	s = secret.New("$ecre!").WithRedact(secret.Redacted)
	fmt.Println(s, s.Value())

	s = secret.New("$ecre!").WithRedact("my redact hint")
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

func ExampleText_Equals() {
	s1 := secret.New("hello")
	s2 := secret.New("hello")
	s3 := secret.New("hi")
	fmt.Println(s1.Equals(s2))
	fmt.Println(s1.Equals(s3))

	// Output:
	// true
	// false
}
