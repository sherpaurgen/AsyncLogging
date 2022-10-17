package AsyncLogging

import (
	"io"
	"sync"
)

// Basic logger struct. Use message channel for async logging and writer for synchronous logging
type TLog struct {
	dest             io.Writer
	m                *sync.Mutex
	msgCh            chan string
	errCh            chan error
	shutdownCh       chan struct{}
	shutdownComplete chan struct{}
}
