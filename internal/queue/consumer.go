// Copyright (c) 2020. Agylia Ltd. All rights reserved

// Package queue contains the types for writing and listening on a queue.
package queue

// Consumer abstracts the queue consumer.
type Consumer interface {
	// Open create a connection to the queue and ensures the queue exists. If the
	// connection closes the supplied function is called.
	Open(onClosed func(err error)) error

	// Listen starts consuming the queue and when a message is received the
	// supplied onReceived function is called.
	Listen(onReceived func(buf []byte), onClosed func(err error)) error
}
