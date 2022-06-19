/*
 * Commercetools Tweakwise Integration
 *
 * A service built for managing the Commercetools to Tweakwise integration
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"context"
	"net/http"
)

// ConfigurationApiRouter defines the required methods for binding the api requests to a responses for the ConfigurationApi
// The ConfigurationApiRouter implementation should parse necessary information from the http request,
// pass the data to a ConfigurationApiServicer to perform the required actions, then write the service results to the http response.
type ConfigurationApiRouter interface {
	ConfigurationGet(http.ResponseWriter, *http.Request)
	ConfigurationPost(http.ResponseWriter, *http.Request)
}

// ManageApiRouter defines the required methods for binding the api requests to a responses for the ManageApi
// The ManageApiRouter implementation should parse necessary information from the http request,
// pass the data to a ManageApiServicer to perform the required actions, then write the service results to the http response.
type ManageApiRouter interface {
	ConnectPost(http.ResponseWriter, *http.Request)
	StartPost(http.ResponseWriter, *http.Request)
	StopPost(http.ResponseWriter, *http.Request)
}

// ConfigurationApiServicer defines the api actions for the ConfigurationApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ConfigurationApiServicer interface {
	ConfigurationGet(context.Context) (ImplResponse, error)
	ConfigurationPost(context.Context, Configuration) (ImplResponse, error)
}

// ManageApiServicer defines the api actions for the ManageApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ManageApiServicer interface {
	ConnectPost(context.Context) (ImplResponse, error)
	StartPost(context.Context) (ImplResponse, error)
	StopPost(context.Context) (ImplResponse, error)
}