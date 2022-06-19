package processor

import (
	"context"
	"github.com/demeyerthom/commercetools-tweakwise-integration/internal/receiver"
	"github.com/demeyerthom/commercetools-tweakwise-integration/pkg"
	tweakwise "github.com/demeyerthom/go-tweakwise-sdk/backend"
	commercetools "github.com/labd/commercetools-go-sdk/platform"
)

func ProcessProductChanges(ctx context.Context, commercetoolsClient *commercetools.ByProjectKeyRequestBuilder, tweakwiseClient *tweakwise.APIClient, mapper pkg.ProductProjectionMapper, changes <-chan *receiver.Payload) error {
	return nil
}
