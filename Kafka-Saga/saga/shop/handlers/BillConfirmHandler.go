package handlers

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"kafka-saga/saga/consts"
	"log"
)

type BillConfirmHandler struct {
	P sarama.SyncProducer
}

func (bc *BillConfirmHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (bc *BillConfirmHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (bc *BillConfirmHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var d consts.Order
		err := json.Unmarshal(msg.Value, &d)
		if err != nil {
			log.Print("reserve data %v: %v", string(msg.Value), err)
			continue
		}
		log.Printf("billing reports to shop - order %v payed: %v", d.Id, err)
	}
	return nil
}
