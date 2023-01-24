package mendix

import "fmt"

const (
	NAMESPACES       = "clusters/%s/namespaces"
	NAMESPACES_BY_ID = "clusters/%s/namespaces/%s"
)

type NameSpaceService struct {
	client *Client
}

type ListOfNamespace struct {
	Namespaces []struct {
		ManifestVersion string `json:"manifestVersion"`
		Namespace       struct {
			NamespaceID      string `json:"namespaceId"`
			Name             string `json:"name"`
			InstallationType string `json:"installationType"`
			Secret           string `json:"secret"`
		} `json:"namespace"`
	} `json:"namespaces"`
}

// Retrieves the manifests of all namespaces of a cluster.
func (s *NameSpaceService) GetListOfNameSpaces(clusterId string) (*ListOfNamespace, error) {
	var out ListOfNamespace
	u := fmt.Sprintf(NAMESPACES, clusterId)
	req, _ := s.client.NewRequest("GET", u, nil)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// Creates a new namespace based on a manifest.
func (s *NameSpaceService) CreateNameSpace(clusterId string, namespace *Namespace) (*Namespace, error) {
	var out Namespace
	u := fmt.Sprintf(NAMESPACES, clusterId)
	req, _ := s.client.NewRequest("POST", u, namespace)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieves the manifest of a namespace.
func (s *NameSpaceService) GetNameSpaceById(clusterId string, namespaceId string) (*Namespace, error) {
	var out Namespace
	u := fmt.Sprintf(NAMESPACES_BY_ID, clusterId, namespaceId)
	req, _ := s.client.NewRequest("GET", u, nil)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// Deletes a specific namespace.
func (s *NameSpaceService) DeleteNameSpace(clusterId string, namespaceId string) (*JobResponse, error) {
	var out JobResponse
	u := fmt.Sprintf(NAMESPACES_BY_ID, clusterId, namespaceId)
	req, _ := s.client.NewRequest("DELETE", u, nil)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
