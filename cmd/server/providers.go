package main

import (
	"github.com/google/wire"
)

var providers = wire.NewSet(httpHandler,
	processorProvider,
	writerProvider)
