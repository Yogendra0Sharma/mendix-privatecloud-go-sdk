package mendix_test

import (
	"testing"
)

func TestListCLI(t *testing.T) {
	client := setup(t)
	out, err := client.Cli.List()
	prettyPrint(out)
	if err != nil {
		t.Errorf("ERROR List CLI : %v", err)
	}
}
