package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"math/rand"
	"time"

	"github.com/Shopify/sarama"
	"kafka-saga/saga/consts"
)

func NewShop(ctx context.Context) error {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	syncProducer, err := sarama.NewSyncProducer(consts.Brokers, cfg)
	if err != nil {
		log.Fatalf("sync kafka: %v", err)
	}

	go func() {
		for {
			d := consts.Order{
				Id:   int(time.Now().UnixNano()),
				Data: time.Now().Format("order 150405.000"),
			}
			b, err := json.Marshal(d)
			if err != nil {
				log.Printf("marshal unsuccessful %v", err)
				continue
			}
			par, off, err := syncProducer.SendMessage(&sarama.ProducerMessage{
				Topic: "reserve_orders",
				Key:   sarama.StringEncoder(fmt.Sprintf("%v", d.Id)),
				Value: sarama.ByteEncoder(b),
			})
			log.Printf("order %v -> %v; %v", par, off, err)

			time.Sleep(time.Millisecond * 500)

			if rand.Intn(10) == 9 {
				par, off, err = syncProducer.SendMessage(&sarama.ProducerMessage{
					Topic: "order_reset",
					Key:   sarama.StringEncoder(fmt.Sprintf("%v", d.Id)),
					Value: sarama.ByteEncoder(b),
				})
				log.Printf("reset %v -> %v; %v", par, off, err)
			}
		}
	}()

	return nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	err := NewShop(ctx)
	if err != nil {
		log.Fatalf("NewStock: %v", err)
	}
	<-ctx.Done()
}
