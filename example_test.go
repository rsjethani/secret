package secret_test

import (
	"encoding/json"
	"fmt"

	"github.com/rsjethani/secret"
)

func Example() {
	type login struct {
		User      string
		Password1 secret.Secret
		Password2 secret.Secret
		Password3 secret.Secret
		Password4 secret.Secret
	}

	x := login{
		User:      "John",
		Password1: secret.New("pass1"),
		Password2: secret.New("pass2", secret.Redacted),
		Password3: secret.New("pass3", secret.FiveXs),
		Password4: secret.New("pass4", secret.CustomRedact("^^^^^")),
	}

	bytes, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", x)
	fmt.Printf("%v, %v, %v, %v\n", x.Password1, x.Password2, x.Password3, x.Password4)
	fmt.Printf("%v, %v, %v, %v\n", x.Password1.Value(), x.Password2.Value(), x.Password3.Value(), x.Password4.Value())
	fmt.Printf("%v\n", string(bytes))
	// Output:
	// {User:John Password1:***** Password2:[REDACTED] Password3:XXXXX Password4:^^^^^}
	// *****, [REDACTED], XXXXX, ^^^^^
	// pass1, pass2, pass3, pass4
	// {"User":"John","Password1":"*****","Password2":"[REDACTED]","Password3":"XXXXX","Password4":"^^^^^"}
}
