package event

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
)

type Event interface {
	ID() uint
	TriggerAt() enum.TriggerType
	Triggered(context.CallbackContext) bool
	Callback() func(*context.CallbackContext)
}
