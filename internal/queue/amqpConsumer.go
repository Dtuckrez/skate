// Copyright (c) 2020. Agylia Ltd. All rights reserved

// Package queue contains the types for writing and listening on a queue.
package queue

import (
	"fmt"
)

// AMQPConsumer is a type to model a queue processor node for a queue.
type AMQPConsumer struct {
	*AMQPQueue
}

// NewAMQPConsumer creates a new instance of the AMQPConsumer type with the
// supplied logger, url and queue name.
func NewAMQPConsumer(url, queueName string) *AMQPConsumer {
	cons := &AMQPConsumer{}
	cons.AMQPQueue = NewAMQPQueue(url, queueName)

	return cons
}

// Open create a connection to the queue and ensures the queue exists. If the
// connection closes the supplied function is called.
func (i *AMQPConsumer) Open(onClosed func(err error)) error {
	// open the queue
	if err := i.AMQPQueue.Open(onClosed); err != nil {
		return fmt.Errorf("[Open][%v]", err)
	}

	return nil
}

// Listen starts consuming the queue and when a message is received the
// supplied onReceived function is called.
func (i *AMQPConsumer) Listen(onReceived func(buf []byte)) error {
	// set QoS
	if err := i.chnl.Qos(
		4,    // prefetch count
		0,    // prefetch size
		true, // global
	); err != nil {
		return fmt.Errorf("[Listen][failed to set QoS for %s queue: %v]", i.queueName, err)
	}

	// consume channel
	msgs, err := i.chnl.Consume(
		i.queue.Name, // queue
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return fmt.Errorf("[Listen][failed to consume %s queue: %v]", i.queueName, err)
	}

	// listen for messages
	go func() {
		for msg := range msgs {
			onReceived(msg.Body)
			func() { _ = msg.Ack(false) }() // nolint: gosec
		}
	}()

	return nil
}
