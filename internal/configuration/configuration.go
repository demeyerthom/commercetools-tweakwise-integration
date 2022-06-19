package configuration

import (
	"context"
	"encoding/json"
	"fmt"
	commercetools "github.com/labd/commercetools-go-sdk/platform"
	"log"
)

var (
	//TODO: do we want these configurable?
	container = "tweakwise"
	key       = "configuration"
)

//TODO: do we want to split the config elements to separate keys within object container?
func Create(ctx context.Context, client *commercetools.ByProjectKeyRequestBuilder, configuration *Configuration) error {
	var body = commercetools.CustomObjectDraft{
		Container: container,
		Key:       key,
		Value:     configuration,
	}

	res, err := client.CustomObjects().Post(body).Execute(ctx)
	if err != nil {
		return err
	}

	log.Println(fmt.Printf("Created new object with id %s", res.ID))

	return nil
}

func Get(ctx context.Context, client *commercetools.ByProjectKeyRequestBuilder) (*Configuration, error) {
	res, err := client.CustomObjects().WithContainerAndKey(container, key).Get().Execute(ctx)
	if err != nil {
		return nil, err
	}

	var config Configuration

	err = json.Unmarshal([]byte(res.Value.(string)), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

type SourceType string

const (
	GCP  SourceType = "gcp"
	HTTP            = "http"
)

type Configuration struct {
	Tweakwise struct {
		Instance string `json:"instance"`
		Token    string `json:"token"`
	}

	Source struct {
		Type SourceType `json:"type"`
		//TODO: fill in additional source information (what is the nicest way to do this?)
	}
}
