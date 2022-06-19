package main

import (
	"context"
	"github.com/demeyerthom/commercetools-tweakwise-integration/internal/configuration"
	"github.com/demeyerthom/commercetools-tweakwise-integration/internal/processor"
	"github.com/demeyerthom/commercetools-tweakwise-integration/internal/receiver"
	tweakwisehelper "github.com/demeyerthom/commercetools-tweakwise-integration/internal/tweakwise"
	"github.com/demeyerthom/commercetools-tweakwise-integration/pkg"
	tweakwise "github.com/demeyerthom/go-tweakwise-sdk/backend"
	commercetools "github.com/labd/commercetools-go-sdk/platform"
	"github.com/spf13/viper"
	"golang.org/x/oauth2/clientcredentials"
	"log"
)

var (
	ctx = context.Background()

	projectClient   *commercetools.ByProjectKeyRequestBuilder
	tweakwiseClient *tweakwise.APIClient

	config *configuration.Configuration

	receiverImpl receiver.ProductChangeReceiver
	mapper       pkg.ProductProjectionMapper
)

func init() {
	viper.AutomaticEnv()
}

func startup() {
	commercetoolsClient, err := commercetools.NewClient(&commercetools.ClientConfig{
		Credentials: &clientcredentials.Config{
			ClientID:     viper.GetString("CTP_CLIENT_ID"),
			ClientSecret: viper.GetString("CTP_CLIENT_SECRET"),
			TokenURL:     viper.GetString("CTP_AUTH_URL") + "/oauth/token",
			Scopes:       viper.GetStringSlice("CTP_SCOPES"),
		},
		URL: viper.GetString("CTP_API_URL"),
	})

	if err != nil {
		log.Fatal(err)
	}

	projectClient = commercetoolsClient.WithProjectKey(viper.GetString("CTP_PROJECT_KEY"))

	config, err = configuration.Get(ctx, projectClient)
	if err != nil {
		log.Fatal(err)
	}

	tweakwiseClient = tweakwisehelper.Create(config.Tweakwise.Instance, config.Tweakwise.Token)
}

func main() {
	startup()
	var err error
	log.Println("Starting synchronization")

	//TODO: determine if we want bounded or unbounded channel
	var productChangesChan = make(chan *receiver.Payload, 1)

	err = receiverImpl.ReceiveProductChanges(ctx, productChangesChan)
	if err != nil {
		log.Fatal(err)
	}

	err = processor.ProcessProductChanges(ctx, projectClient, tweakwiseClient, mapper, productChangesChan)

	log.Println("Finished synchronization")
}
