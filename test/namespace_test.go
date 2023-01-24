package mendix_test

import "testing"

func TestGetListOfNameSpaces(t *testing.T) {
	client := setup(t)
	clusterId := ""
	out, err := client.Namespace.GetListOfNameSpaces(clusterId)
	prettyPrint(out)
	if err != nil {
		t.Errorf("ERROR List Clusters : %v", err)
	}
}

func TestGetNameSpaceById(t *testing.T) {
	client := setup(t)
	clusterId := ""
	namespaceId := ""
	out, err := client.Namespace.GetNameSpaceById(clusterId, namespaceId)
	prettyPrint(out)
	if err != nil {
		t.Errorf("ERROR List Clusters : %v", err)
	}
}
