package handlers

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"kafka-saga/saga/consts"
	"log"
)

type ResetHandler struct {
	P sarama.SyncProducer
}

func (r *ResetHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (r *ResetHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (r *ResetHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var d consts.Order
		err := json.Unmarshal(msg.Value, &d)
		if err != nil {
			log.Print("reset data %v: %v", string(msg.Value), err)
			continue
		}
		log.Printf("Stock repors - order %v deleted", d.Id)
	}
	return nil
}
