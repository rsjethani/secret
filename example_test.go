package secret_test

import (
	"encoding/json"
	"fmt"

	"github.com/rsjethani/secret/v2"
)

func ExampleNewText() {
	s := secret.NewText("$ecre!")
	fmt.Println(s, s.Value())
	//Output: ***** $ecre!
}

func ExampleFiveXs() {
	s := secret.NewText("$ecre!", secret.FiveXs)
	fmt.Println(s, s.Value())
	// Output: XXXXX $ecre!
}

func ExampleRedacted() {
	s := secret.NewText("$ecre!", secret.Redacted)
	fmt.Println(s, s.Value())
	// Output: [REDACTED] $ecre!
}

func ExampleCustomRedact() {
	s := secret.NewText("$ecre!", secret.CustomRedact("HIDDEN"))
	fmt.Println(s, s.Value())
	// Output: HIDDEN $ecre!
}

func ExampleText_MarshalJSON() {
	login := struct {
		User     string
		Password secret.Text
	}{
		User:     "John",
		Password: secret.NewText("shh!"),
	}

	bytes, err := json.Marshal(&login)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
	// Output: {"User":"John","Password":"*****"}
}

func ExampleText_UnmarshalJSON() {
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
