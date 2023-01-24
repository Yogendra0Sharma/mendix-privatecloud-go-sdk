package mendix

import "fmt"

const (
	CLUSTERS      = "clusters"
	CLUSTER_BY_ID = "clusters/%s"
)

type ClusterService struct {
	client *Client
}

type ListOfClusterResponse struct {
	Clusters []struct {
		ManifestVersion string `json:"manifestVersion"`
		Cluster         struct {
			ClusterID   string `json:"clusterId"`
			Name        string `json:"name"`
			ClusterType string `json:"clusterType"`
			Description string `json:"description"`
		} `json:"cluster"`
	} `json:"clusters"`
}

// Retrieves all the clusters for which the user is a Cluster Manager.
func (s *ClusterService) GetListOfClusters() (*ListOfClusterResponse, error) {
	var out ListOfClusterResponse
	req, _ := s.client.NewRequest("GET", CLUSTERS, nil)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// Creates a new cluster based on a manifest.
func (s *ClusterService) CreateCluster(cluster *ClusterObj) (*ClusterObj, error) {
	var out ClusterObj

	req, _ := s.client.NewRequest("POST", CLUSTERS, &out)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return &out, err
	}

	return &out, nil
}

// Retrieves the manifest of a cluster
func (s *ClusterService) GetClusterById(clusterId string) (*ClusterObj, error) {
	var out ClusterObj
	u := fmt.Sprintf(CLUSTER_BY_ID, clusterId)
	req, _ := s.client.NewRequest("GET", u, &out)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return &out, err
	}

	return &out, nil
}

// Updates a specific cluster
func (s *ClusterService) UpdateCluster(clusterId string, cluster *ClusterObj) (*ClusterObj, error) {
	var out ClusterObj
	u := fmt.Sprintf(CLUSTER_BY_ID, clusterId)
	req, _ := s.client.NewRequest("PUT", u, &out)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// Deletes a specific cluster.
func (s *ClusterService) DeleteCluster(clusterId string) (*JobResponse, error) {
	var out JobResponse
	u := fmt.Sprintf(CLUSTER_BY_ID, clusterId)
	req, _ := s.client.NewRequest("DELETE", u, &out)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
