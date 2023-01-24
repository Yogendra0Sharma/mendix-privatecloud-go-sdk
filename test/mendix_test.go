package mendix_test

import (
	"encoding/json"
	"fmt"
	"testing"

	mendix "github.com/yogendra0sharma/mendix-privatecloud-go-sdk"
)

func prettyPrint(data interface{}) {
	dbg, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(dbg))
}

func setup(t *testing.T) *mendix.Client {
	// client is the Gitlab client being tested.
	client := mendix.NewClient("")
	return client
}
