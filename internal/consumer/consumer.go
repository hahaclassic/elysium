package consumer

import (
	"log"
	"sync"
	"time"

	"github.com/hahaclassic/elysium/config"
	"github.com/hahaclassic/elysium/internal/model"
)

type Fetcher interface {
	Fetch(limit int) ([]*model.Event, error)
}

type Processor interface {
	Process(e *model.Event, errors chan error)
}

type Consumer struct {
	fetcher   Fetcher
	processor Processor
	batchSize int
}

func New(config *config.ConsumerConfig, fetcher Fetcher, processor Processor) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: config.BatchSize,
	}
}

func (c *Consumer) Start() error {
	for {
		events, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(events) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		go c.handleEvents(events)
	}
}

func (c *Consumer) handleEvents(events []*model.Event) {
	errors := make(chan error, c.batchSize)

	wg := sync.WaitGroup{}
	wg.Add(len(events))

	for _, event := range events {
		go func() {
			defer wg.Done()

			c.processor.Process(event, errors)

			if err := <-errors; err != nil {
				log.Print()
			}
		}()
	}

	wg.Wait()
}
