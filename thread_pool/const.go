package main

import "errors"

var ErrInvalidPoolSize = errors.New("invalid pool size")
var ErrInvalidQueueSize = errors.New("invalid queue size")
var ErrInvalidJobType = errors.New("invalid job type")
