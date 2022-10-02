package main

import (
	"golang.org/x/net/context"
	"kafka-saga/saga/billing/handlers"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"kafka-saga/saga/consts"
)

type Pay struct {
	producer     sarama.SyncProducer
	billConsumer *handlers.BillingHandler
}

func NewPay(ctx context.Context) (*Pay, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	syncProducer, err := sarama.NewSyncProducer(consts.Brokers, cfg)
	if err != nil {
		log.Fatalf("sync kafka: %v", err)
		return nil, err
	}

	// receiving payment from shop
	billing, err := sarama.NewConsumerGroup(consts.Brokers, "payBilling", cfg)
	if err != nil {
		return nil, err
	}
	bHandler := &handlers.BillingHandler{
		P: syncProducer,
	}

	go func() {
		for {
			err := billing.Consume(ctx, []string{"bill_send"}, bHandler)
			log.Printf("bill recieved")
			if err != nil {
				log.Printf("bill consumer error: %v", err)
				time.Sleep(time.Second * 5)
			}
		}
	}()

	return &Pay{
		producer:     syncProducer,
		billConsumer: bHandler,
	}, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	_, err := NewPay(ctx)
	if err != nil {
		log.Fatalf("NewPay: %v", err)
	}
	<-ctx.Done()
}
