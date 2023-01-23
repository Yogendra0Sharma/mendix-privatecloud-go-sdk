package mendix_test

import (
	"testing"

	mendix "github.com/yogendra0sharma/mendix-privatecloud-go-sdk"
)

func setup(t *testing.T) *mendix.Client {
	// client is the Gitlab client being tested.
	client := mendix.NewClient("")
	return client
}
