package mendix

import "fmt"

const (
	ENVIRONMENT       = "apps/%s/environments"
	ENVIRONMENT_BY_ID = "apps/%s/environments/%s"
)

type EnvironmentService struct {
	client *Client
}

type ListOfEnvironment struct {
	Environments []struct {
		ManifestVersion string `json:"manifestVersion"`
		Environment     struct {
			ID         string `json:"id"`
			Properties struct {
				Name                string `json:"name"`
				Production          bool   `json:"production"`
				DefaultStudioTarget bool   `json:"defaultStudioTarget"`
			} `json:"properties"`
			Deployment struct {
				AppURL    string `json:"appUrl"`
				PackageID string `json:"packageId"`
				Constants []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"constants"`
				ScheduledEvents []struct {
					Name    string `json:"name"`
					Enabled bool   `json:"enabled"`
				} `json:"scheduledEvents"`
				AppLogLevels []struct {
					Name  string `json:"name"`
					Level string `json:"level"`
				} `json:"appLogLevels"`
			} `json:"deployment"`
			Container struct {
				State     string `json:"state"`
				Instances int    `json:"instances"`
				Resources struct {
					Limits struct {
						CPU    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"limits"`
					Requests struct {
						CPU    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"requests"`
				} `json:"resources"`
			} `json:"container"`
			Network struct {
				OutgoingConnectionCertificate []struct {
					Enabled         bool   `json:"enabled"`
					Password        string `json:"password"`
					CertificateType string `json:"certificateType"`
					Key             string `json:"key"`
				} `json:"outgoingConnectionCertificate"`
			} `json:"network"`
			Runtime struct {
				SubscriptionSecret   string `json:"subscriptionSecret"`
				DebuggerPassword     string `json:"debuggerPassword"`
				MxAdminPassword      string `json:"mxAdminPassword"`
				EnvironmentVariables []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"environmentVariables"`
				CustomSettings []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"customSettings"`
			} `json:"runtime"`
			Provider struct {
				ClusterID string `json:"clusterId"`
				Namespace string `json:"namespace"`
			} `json:"provider"`
			Annotations []struct {
				Key            string `json:"key"`
				Value          string `json:"value"`
				AnnotationType string `json:"annotationType"`
			} `json:"annotations"`
			Services []struct {
				Name   string `json:"name"`
				Plan   string `json:"plan"`
				Config string `json:"config"`
			} `json:"services"`
			RuntimeMetricsConfiguration struct {
				Mode                         string `json:"mode"`
				Interval                     string `json:"interval"`
				MxAgentConfig                string `json:"mxAgentConfig"`
				MxAgentInstrumentationConfig string `json:"mxAgentInstrumentationConfig"`
			} `json:"runtimeMetricsConfiguration"`
		} `json:"environment"`
	} `json:"environments"`
}

// Retrieves the manifests of all environments of an application.
func (s *EnvironmentService) GetListOfEnvironments(appId string) (*ListOfEnvironment, error) {
	var out ListOfEnvironment
	u := fmt.Sprintf(ENVIRONMENT, appId)
	req, _ := s.client.NewRequest("GET", u, nil)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// Creates a new environment based on a manifest.
func (s *EnvironmentService) CreateEnvironment(appId string, environment *Environment) (*Environment, error) {
	var out Environment
	u := fmt.Sprintf(ENVIRONMENT, appId)
	req, _ := s.client.NewRequest("POST", u, environment)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// Retrieves the manifest of an environment.
func (s *EnvironmentService) GetEnvironmentById(appId string, environmentId string) (*Environment, error) {
	var out Environment
	u := fmt.Sprintf(ENVIRONMENT_BY_ID, appId, environmentId)
	req, _ := s.client.NewRequest("GET", u, nil)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// Updates a specific environment based on a manifest.
func (s *EnvironmentService) UpdateEnvironment(appId string, environmentId string, environment *Environment) (*Environment, error) {
	var out Environment
	u := fmt.Sprintf(ENVIRONMENT_BY_ID, appId, environmentId)
	req, _ := s.client.NewRequest("PUT", u, environment)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// Deletes a specific environment.
func (s *EnvironmentService) DeleteEnvironment(appId string, environmentId string) (*JobResponse, error) {
	var out JobResponse
	u := fmt.Sprintf(ENVIRONMENT_BY_ID, appId, environmentId)
	req, _ := s.client.NewRequest("DELETE", u, nil)
	_, err := s.client.Do(req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
