package main

import (
	"kafka-saga/saga/stock/dto"
	"kafka-saga/saga/stock/handlers"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"golang.org/x/net/context"
	"kafka-saga/saga/consts"
)

type Store struct {
	data           *dto.Map
	producer       sarama.SyncProducer
	incomeConsumer *handlers.IncomeHandler
	resetConsumer  *handlers.ResetHandler
}

func NewStore(ctx context.Context) (*Store, error) {
	Data := dto.NewMap()

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(consts.Brokers, cfg)
	if err != nil {
		return nil, err
	}

	income, err := sarama.NewConsumerGroup(consts.Brokers, "store", cfg)
	if err != nil {
		return nil, err
	}
	iHandler := &handlers.IncomeHandler{
		P:    producer,
		Data: Data,
	}
	go func() {
		for {
			err := income.Consume(ctx, []string{"income_orders"}, iHandler)
			if err != nil {
				log.Printf("income consumer error: %v", err)
				time.Sleep(time.Second * 5)
			}
			log.Printf("income consumer done")
		}
	}()

	reset, err := sarama.NewConsumerGroup(consts.Brokers, "storeReset", cfg)
	if err != nil {
		return nil, err
	}
	rHandler := &handlers.ResetHandler{
		P:    producer,
		Data: Data,
	}
	go func() {
		for {
			err := reset.Consume(ctx, []string{"reset_orders"}, rHandler)
			log.Printf("order reset")
			if err != nil {
				log.Printf("reset consumer error: %v", err)
				time.Sleep(time.Second * 5)
			}
		}
	}()
	return &Store{
		data:           dto.NewMap(),
		producer:       producer,
		incomeConsumer: iHandler,
		resetConsumer:  rHandler,
	}, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	_, err := NewStore(ctx)
	if err != nil {
		log.Fatalf("NewStore: %v", err)
	}
	<-ctx.Done()
}
