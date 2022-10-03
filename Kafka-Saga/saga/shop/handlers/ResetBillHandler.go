package handlers

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"kafka-saga/saga/consts"
	"log"
)

type ResetBillHandler struct {
	P sarama.SyncProducer
}

func (rb *ResetBillHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (rb *ResetBillHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (rb *ResetBillHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var d consts.Order
		err := json.Unmarshal(msg.Value, &d)
		if err != nil {
			log.Print("reset data %v: %v", string(msg.Value), err)
			continue
		}
		log.Printf("billing repors - order %v deleted", d.Id)
	}
	return nil
}
