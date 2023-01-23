package mendix

const (
	ENVIRONMENT       = "apps/%s/environments"
	ENVIRONMENT_BY_ID = "apps/%s/environments/%s"
)

type EnvironmentService struct {
	client *Client
}
