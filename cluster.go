package mendix

const (
	CLUSTERS      = "clusters"
	CLUSTER_BY_ID = "clusters/%s"
)

type ClusterService struct {
	client *Client
}
