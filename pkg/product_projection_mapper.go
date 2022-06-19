package pkg

import (
	tweakwise "github.com/demeyerthom/go-tweakwise-sdk/backend"
	commercetools "github.com/labd/commercetools-go-sdk/platform"
)

// ProductProjectionMapper is the interface that we're exposing as a plugin.
type ProductProjectionMapper interface {
	Map(languages []string, product *commercetools.ProductProjection) (map[string]*tweakwise.ProductApiModel, error)
}
