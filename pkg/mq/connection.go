package mq

import (
	"log"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
)


type Connection struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewConnection(url string) (*Connection, error) {
	var conn *amqp.Connection
	var ch *amqp.Channel
	var err error

	// Retry logic
for i := 0; i < 5; i++ {
		conn, err = amqp.Dial(url)
		if err != nil {
			log.Printf("RabbitMQ connection failed: %v. Retrying...", err)
			time.Sleep(2 * time.Second)
			continue
		}

		ch, err = conn.Channel()
		if err != nil {
			log.Printf("Failed to open channel: %v. Retrying...", err)
			conn.Close() // be sure to close bad connection
			time.Sleep(2 * time.Second)
			continue
		}

		// Success
		return &Connection{Conn: conn, Channel: ch}, nil
	}

	// If we reach here, all retries failed
	return nil, err
}

func (c *Connection) Close() {
	c.Channel.Close()
	c.Conn.Close()
}