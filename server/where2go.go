package server

import (
	"go.uber.org/zap"
)

var Application Where2Go

type Where2Go struct {
	Logger *zap.Logger
}
