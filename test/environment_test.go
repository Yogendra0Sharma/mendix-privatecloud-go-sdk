package mendix_test

import "testing"

func TestGetListOfEnvironments(t *testing.T) {
	client := setup(t)
	appid := ""
	out, err := client.Environment.GetListOfEnvironments(appid)
	prettyPrint(out)
	if err != nil {
		t.Errorf("ERROR List Clusters : %v", err)
	}
}

func TestGetEnvironmentById(t *testing.T) {
	client := setup(t)
	appId := ""
	environmentId := ""
	out, err := client.Environment.GetEnvironmentById(appId, environmentId)
	prettyPrint(out)
	if err != nil {
		t.Errorf("ERROR List Clusters : %v", err)
	}
}
