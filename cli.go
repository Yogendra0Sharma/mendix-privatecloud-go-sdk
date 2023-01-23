package mendix

const (
	// Get List of Available Versions of mxpc-cli
	GET_CLI_LIST = "cli"

	// Get/Download mxpc-cli by cli version
	DOWNLOAD_CLI = "cli/%s"
)

type CLIService struct {
	client *Client
}

type CliVersionListResponse struct {
	CliVersionList []struct {
		OperatorVersion           string `json:"operatorVersion"`
		CliVersion                string `json:"cliVersion"`
		AvailableOperatingSystems string `json:"availableOperatingSystems"`
		AvailableArchitectures    string `json:"availableArchitectures"`
	} `json:"cliVersionList"`
}

// Retrieves all available versions of mxpc-cli.
func (s *CLIService) List() (*CliVersionListResponse, error) {
	var out CliVersionListResponse
	req, err := s.client.NewRequest("GET", GET_CLI_LIST, nil)
	resp, err := s.client.Do(req, &out)
	prettyPrint(resp)
	if err != nil {
		return &out, err
	}
	return &out, nil
}
