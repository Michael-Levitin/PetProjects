package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"kafka-saga/saga/consts"
	"kafka-saga/saga/stock/dto"
	"log"
	"math/rand"
)

type IncomeHandler struct {
	P    sarama.SyncProducer
	Data *dto.Map
}

func (i *IncomeHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (i *IncomeHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (i *IncomeHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var d consts.Order
		err := json.Unmarshal(msg.Value, &d)
		if err != nil {
			log.Print("income data %v: %v", string(msg.Value), err)
			continue
		}
		if rand.Intn(5) == 4 {
			_, _, err := i.P.SendMessage(&sarama.ProducerMessage{
				Topic: "stock_order_reset",
				Key:   sarama.StringEncoder(fmt.Sprintf("%v", d.Id)),
				Value: sarama.ByteEncoder(msg.Value),
			})
			if err != nil {
				log.Printf("Cant send reset: %v", err)
			}
			log.Printf("stock resets order %v", d.Id)
			continue
		}
		i.Data.Set(d.Id, d.Data)
		log.Printf("оrder %v reserved", d.Id)
	}
	return nil
}
