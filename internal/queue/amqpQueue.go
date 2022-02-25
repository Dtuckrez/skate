// Copyright (c) 2020. Agylia Ltd. All rights reserved

// Package queue contains the types for writing and listening on a queue.
package queue

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

// AMQPQueue represents a AMQP queue client.
type AMQPQueue struct {
	url       string
	queueName string

	// AMQP instance values
	conn  *amqp.Connection
	chnl  *amqp.Channel
	queue *amqp.Queue
}

// NewAMQPQueue creates a new instance of the AMQPQueue type with the
// supplied logger, url and queue name.
func NewAMQPQueue(url, queueName string) *AMQPQueue {
	return &AMQPQueue{
		url:       url,
		queueName: queueName,
	}
}

// Open create a connection to the queue and ensures the queue exists. If the
// connection closes the supplied function is called.
func (i *AMQPQueue) Open(onClosed func(err error)) error {
	var err error
	// Connect
	i.conn, err = amqp.Dial(i.url)
	if err != nil {
		return fmt.Errorf("[Open][failed to dial %s queue: %v]", i.queueName, err)
	}

	// create channel
	i.chnl, err = i.conn.Channel()
	if err != nil {
		return fmt.Errorf("[Open][failed to create channel for %s queue: %v]", i.queueName, err)
	}

	// create queue
	i.queue, err = declare(i.chnl, i.queueName)
	if err != nil {
		return fmt.Errorf("[Open]%v", err)
	}

	go func() {
		err := <-i.conn.NotifyClose(make(chan *amqp.Error))
		onClosed(fmt.Errorf("[connection closed for %s queue: %v]", i.queueName, err))
	}()

	return nil
}

// Enqueue appends a single EnqueueMessage object to the queue.
func (i *AMQPQueue) Enqueue(buf []byte) error {
	if err := i.chnl.Publish(
		"",           // exchange (default)
		i.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         buf,
			AppId:        "skate",
			Timestamp:    time.Now(),
		},
	); err != nil {
		return fmt.Errorf("[Enqueue][failed to publish to %s queue %v]", i.queueName, err)
	}

	return nil
}

// declare delares a new AMQP queue with the initialisation parameters required
// for the serivce.
func declare(ch *amqp.Channel, queueName string) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return nil, fmt.Errorf("[declare][failed to declare %s queue: %v]", queueName, err)
	}

	return &q, nil
}
