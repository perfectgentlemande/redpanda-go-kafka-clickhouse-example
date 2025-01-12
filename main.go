package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
)

const topic = "hotels_snapshot"

var brokers = []string{"localhost:9093"}

type Content struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Hotel struct {
	ID        uuid.UUID `json:"id"`
	GeoID     uint64    `json:"geo_id"`
	Emails    []string  `json:"emails"`
	Type      uint8     `json:"type"`
	ContentRu *Content  `json:"content_ru"`
	CreatedAt time.Time `json:"created_at"`
}

const (
	hotelTypeHotel     = 1
	hotelTypeApartment = 2
	hotelTypeResort    = 3
)

func main() {
	producer, err := newProducer()
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing producer: %v", err)
		}
	}()

	for _, hotel := range getSampledHotels() {
		data, err := json.Marshal(hotel)
		if err != nil {
			log.Fatalf("error marshalling hotel: %v", err)
		}

		_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic:     topic,
			Partition: -1,
			Value:     sarama.ByteEncoder(data),
		})
		if err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
	}

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.ByteEncoder("some bad document"),
	})
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	log.Println("Messages sent successfully")
}

func newProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}
