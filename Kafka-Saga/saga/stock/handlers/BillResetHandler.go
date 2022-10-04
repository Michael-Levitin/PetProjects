package handlers

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"kafka-saga/saga/consts"
	"kafka-saga/saga/stock/dto"
	"log"
)

type BillResetHandler struct {
	P    sarama.SyncProducer
	Data *dto.Map
}

func (br *BillResetHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (br *BillResetHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (br *BillResetHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var d consts.Order
		err := json.Unmarshal(msg.Value, &d)
		if err != nil {
			log.Print("reset data %v: %v", string(msg.Value), err)
			continue
		}
		if err := br.Data.Delete(d.Id); err != nil {
			log.Printf("bad order: %v, %s", d.Id, err)
			continue
		}
		log.Printf("billing resets order in stock %v", d.Id)
	}
	return nil
}
