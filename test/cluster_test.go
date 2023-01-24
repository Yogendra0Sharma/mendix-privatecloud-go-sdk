package mendix_test

import "testing"

func TestListClusters(t *testing.T) {
	client := setup(t)
	out, err := client.Cluster.GetListOfClusters()
	prettyPrint(out)
	if err != nil {
		t.Errorf("ERROR List Clusters : %v", err)
	}
}

func TestGetClusterById(t *testing.T) {
	client := setup(t)
	clusterId := "a6d13bb5-7afc-4c9a-87e8-dba9a679a8ca"
	out, err := client.Cluster.GetClusterById(clusterId)
	prettyPrint(out)
	if err != nil {
		t.Errorf("ERROR List Clusters : %v", err)
	}
}
