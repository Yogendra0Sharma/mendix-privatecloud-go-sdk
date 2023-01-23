package mendix

const (
	NAMESPACES       = "clusters/%s/namespaces"
	NAMESPACES_BY_ID = "clusters/%s/namespaces/%s"
)

type NameSpaceService struct {
	client *Client
}
