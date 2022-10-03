package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"kafka-saga/saga/consts"
	"log"
	"math/rand"
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

		if rand.Intn(10) == 1 {
			_, _, err := b.P.SendMessage(&sarama.ProducerMessage{
				Topic: "billing_order_reset",
				Key:   sarama.StringEncoder(fmt.Sprintf("%v", d.Id)),
				Value: sarama.ByteEncoder(msg.Value),
			})
			if err != nil {
				log.Printf("Cant send reset: %v", err)
			}
			log.Printf("billing resets order %v", d.Id)
			continue
		}

		_, _, err = b.P.SendMessage(&sarama.ProducerMessage{
			Topic: "bill_confirmed",
			Key:   sarama.StringEncoder(fmt.Sprintf("%v", d.Id)),
			Value: sarama.ByteEncoder(msg.Value),
		})
		if err != nil {
			log.Printf("cannot send bill confiramtion %v", err)
		}
	}
	return nil
}
