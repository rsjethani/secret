package secret_test

import (
	"encoding/json"
	"fmt"

	"github.com/rsjethani/secret/v2"
)

func Example() {
	type login struct {
		User      string
		Password1 secret.Text
		Password2 secret.Text
		Password3 secret.Text
		Password4 secret.Text
	}

	x := login{
		User:      "John",
		Password1: secret.NewText("pass1"),
		Password2: secret.NewText("pass2", secret.Redacted),
		Password3: secret.NewText("pass3", secret.FiveXs),
		Password4: secret.NewText("pass4", secret.CustomRedact("^^^^^")),
	}

	bytes, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", x)
	fmt.Printf("%v, %v, %v, %v\n", x.Password1, x.Password2, x.Password3, x.Password4)
	fmt.Printf("%v, %v, %v, %v\n", x.Password1.Value(), x.Password2.Value(), x.Password3.Value(), x.Password4.Value())
	fmt.Printf("%v\n", string(bytes))

	// Unmarshaling a plain string into a Secret also works.
	y := struct {
		User       string
		Credential secret.Text
	}{}
	err = json.Unmarshal([]byte(`{"User": "Doe", "Credential": "secret"}`), &y)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", y.Credential.Value())

	// Output:
	// {User:John Password1:***** Password2:[REDACTED] Password3:XXXXX Password4:^^^^^}
	// *****, [REDACTED], XXXXX, ^^^^^
	// pass1, pass2, pass3, pass4
	// {"User":"John","Password1":"*****","Password2":"[REDACTED]","Password3":"XXXXX","Password4":"^^^^^"}
	// secret
}
