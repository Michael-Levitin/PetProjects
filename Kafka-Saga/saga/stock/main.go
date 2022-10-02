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

func NewStock(ctx context.Context) (*Store, error) {
	Data := dto.NewMap()

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	syncProducer, err := sarama.NewSyncProducer(consts.Brokers, cfg)
	if err != nil {
		return nil, err
	}

	reserve, err := sarama.NewConsumerGroup(consts.Brokers, "store", cfg)
	if err != nil {
		return nil, err
	}
	iHandler := &handlers.IncomeHandler{
		P:    syncProducer,
		Data: Data,
	}
	go func() {
		for {
			err := reserve.Consume(ctx, []string{"order_send"}, iHandler)
			if err != nil {
				log.Printf("reserve consumer error: %v", err)
				time.Sleep(time.Second * 5)
			}
			log.Printf("reserve consumer done")

		}
	}()

	reset, err := sarama.NewConsumerGroup(consts.Brokers, "stockReset", cfg)
	if err != nil {
		return nil, err
	}
	rHandler := &handlers.ResetHandler{
		P:    syncProducer,
		Data: Data,
	}
	go func() {
		for {
			err := reset.Consume(ctx, []string{"shop_order_reset"}, rHandler)
			log.Printf("order reset")
			if err != nil {
				log.Printf("reset order error: %v", err)
				time.Sleep(time.Second * 5)
			}
		}
	}()
	return &Store{
		data:           dto.NewMap(),
		producer:       syncProducer,
		incomeConsumer: iHandler,
		resetConsumer:  rHandler,
	}, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	_, err := NewStock(ctx)
	if err != nil {
		log.Fatalf("NewStock: %v", err)
	}
	<-ctx.Done()
}
