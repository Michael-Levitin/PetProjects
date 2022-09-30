package handlers

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"kafka-saga/saga/consts"
	"kafka-saga/saga/stock/dto"
	"log"
)

type ResetHandler struct {
	P    sarama.SyncProducer
	Data *dto.Map
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
		if err := r.Data.Delete(d.Id); err != nil {
			log.Printf("bad order: %v", d.Id)
			continue
		}
		log.Printf("Order %v deleted", d.Id)
	}
	return nil
}
