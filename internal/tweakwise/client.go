package tweakwise

import tweakwise "github.com/demeyerthom/go-tweakwise-sdk/backend"

const (
	AuthenticationTokenHeader = "TWN-Authentication"
	InstanceKeyHeader         = "TWN-InstanceKey"
)

func Create(instance string, token string) *tweakwise.APIClient {
	var config = tweakwise.NewConfiguration()
	config.AddDefaultHeader(InstanceKeyHeader, instance)
	config.AddDefaultHeader(AuthenticationTokenHeader, token)

	return tweakwise.NewAPIClient(config)
}
