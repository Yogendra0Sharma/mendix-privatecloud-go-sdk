package mendix_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func prettyPrint(data interface{}) {
	dbg, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(dbg))
}
func TestListCLI(t *testing.T) {
	client := setup(t)
	out, err := client.Cli.List()
	prettyPrint(out)
	if err != nil {
		t.Errorf("ERROR List CLI : %v", err)
	}
}
