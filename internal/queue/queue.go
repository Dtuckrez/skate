// Copyright (c) 2020. Agylia Ltd. All rights reserved

// Package queue contains the types for writing and listening on a queue.
package queue

// Queue abstracts the queue writer.
type Queue interface {
	// Open create a connection to the queue and ensures the queue exists. If the
	// connection closes the supplied function is called.
	Open(onClosed func(err error)) error

	// Enqueue appends a single EnqueueMessage object to the queue.
	Enqueue(buf []byte) error
}
