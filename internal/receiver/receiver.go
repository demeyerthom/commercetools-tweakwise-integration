package receiver

import (
	"context"
	"time"
)

const (
	Product  ReferenceType = "product"
	Category               = "category"
	Stock                  = "stock"
)

type ReferenceType string

type Change struct {
	NotificationType string `json:"notificationType"`
	ProjectKey       string `json:"projectKey"`
	Resource         struct {
		ID     string        `json:"id"`
		TypeID ReferenceType `json:"typeId"`
	} `json:"resource"`
	ResourceUserProvidedIdentifiers struct {
		Key            *string `json:"key,omitempty"`
		ExternalId     *string `json:"externalId,omitempty"`
		OrderNumber    *string `json:"orderNumber,omitempty"`
		CustomerNumber *string `json:"customerNumber,omitempty"`
		Sku            *string `json:"sku,omitempty"`
	} `json:"resourceUserProvidedIdentifiers,omitempty"`
	Version    int       `json:"version"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

type Ack func()

type Nack func()

type Payload struct {
	Change Change
	Ack    Ack
	Nack   Nack
}

type ProductChangeReceiver interface {
	ReceiveProductChanges(ctx context.Context, payload chan<- *Payload) error
}
