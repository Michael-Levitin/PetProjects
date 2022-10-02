package handlers

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"kafka-saga/saga/consts"
	"log"
)

type BillingHandler struct {
	P sarama.SyncProducer
}

func (b *BillingHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (b *BillingHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (b *BillingHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var d consts.Order
		err := json.Unmarshal(msg.Value, &d)
		if err != nil {
			log.Print("reserve data %v: %v", string(msg.Value), err)
			continue
		}
		log.Printf("billing reports - recieved payment for order %v: %v", d.Id, err)
	}
	return nil
}
