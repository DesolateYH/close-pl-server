package processor

import (
	"close-pl-server/internal/connection"
)

type Processor interface {
	Run()
}

type processor struct {
	connection connection.Connection
}

func NewProcessor(connection connection.Connection) Processor {
	return &processor{connection: connection}
}
