package secret_test

import (
	"encoding/json"
	"fmt"

	"github.com/rsjethani/secret/v2"
)

func ExampleText() {
	s := secret.Text{}
	fmt.Println(s, s.Value())
	// Output: *****
}

func ExampleNew() {
	s := secret.New("$ecre!")
	fmt.Println(s, s.Value())
	// Output: ***** $ecre!
}

func ExampleFiveXs() {
	s := secret.New("$ecre!", secret.FiveXs)
	fmt.Println(s, s.Value())
	// Output: XXXXX $ecre!
}

func ExampleRedacted() {
	s := secret.New("$ecre!", secret.Redacted)
	fmt.Println(s, s.Value())
	// Output: [REDACTED] $ecre!
}

func ExampleCustomRedact() {
	s := secret.New("$ecre!", secret.CustomRedact("HIDDEN"))
	fmt.Println(s, s.Value())
	// Output: HIDDEN $ecre!
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
