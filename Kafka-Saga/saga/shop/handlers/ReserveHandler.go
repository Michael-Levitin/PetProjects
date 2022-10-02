package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"kafka-saga/saga/consts"
	"log"
)

type ReserveHandler struct {
	P sarama.SyncProducer
}

func (r *ReserveHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (r *ReserveHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (r *ReserveHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var d consts.Order
		err := json.Unmarshal(msg.Value, &d)
		if err != nil {
			log.Print("reserve data %v: %v", string(msg.Value), err)
			continue
		}
		log.Printf("stock reports - order %v reserved: %v", d.Id, err)

		_, _, err = r.P.SendMessage(&sarama.ProducerMessage{
			Topic: "bill_send",
			Key:   sarama.StringEncoder(fmt.Sprintf("%v", d.Id)),
			Value: sarama.ByteEncoder(msg.Value),
		})
		if err != nil {
			log.Printf("cannot send bill %v", err)
		}
	}
	return nil
}
