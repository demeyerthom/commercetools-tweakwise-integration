package receiver

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
)

type GcpReceiver struct {
	productProjectionSubscription pubsub.Subscription
}

func (g *GcpReceiver) ReceiveProductChanges(ctx context.Context, payloads chan<- *Payload) error {
	return g.productProjectionSubscription.Receive(ctx, func(ctx context.Context, message *pubsub.Message) {
		var change Change
		//TODO: what will we do with errors?
		_ = json.Unmarshal(message.Data, &change)

		ackFunc := func() {
			message.Ack()
		}

		nackFunc := func() {
			message.Nack()
		}

		payloads <- &Payload{
			Change: change,
			Ack:    ackFunc,
			Nack:   nackFunc,
		}
	})
}
