package main

import (
	"context"
	"fmt"
	"github.com/demeyerthom/go-tweakwise-sdk/backend"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"log"
	"os"
)

var (
	commercetoolsProjectKey   = os.Getenv("COMMERCETOOLS_PROJECT_KEY")
	commercetoolsClientId     = os.Getenv("COMMERCETOOLS_CLIENT_ID")
	commercetoolsClientSecret = os.Getenv("COMMERCETOOLS_CLIENT_SECRET")
	commercetoolsScopes       = []string{os.Getenv("COMMERCETOOLS_SCOPES")}

	tweakwiseInstanceKey = os.Getenv("TWEAKWISE_INSTANCE_KEY")
	tweakwiseApiKey      = os.Getenv("TWEAKWISE_API_KEY")

	commercetoolsClient *commercetools.Client

	tweakwiseClient *backend.APIClient
)

func init() {
	var err error

	commercetoolsClient, err = commercetools.NewClient(&commercetools.ClientConfig{
		ProjectKey: commercetoolsProjectKey,
		Endpoints:  commercetools.NewClientEndpoints("europe-west1", "gcp"),
		Credentials: &commercetools.ClientCredentials{
			ClientID:     commercetoolsClientId,
			ClientSecret: commercetoolsClientSecret,
			Scopes:       commercetoolsScopes,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	tweakwiseClient = backend.NewAPIClient(backend.NewConfiguration())
}

func main() {
	fmt.Println("Hello world!")

	product, err := commercetoolsClient.ProductGetWithID(context.Background(), "01b65ae3-d9c2-4a21-a2bb-d5418b4ae6a9")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(product)
}
