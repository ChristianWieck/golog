package log

import (
	"sync"
)

var handlerLock = sync.Mutex{}

type handlerEntry struct {
	LogHandler LogHandler
	LogLevel   LogLevel
}

var logHandlers = []handlerEntry{}

type LogHandler interface {
	Handle(LogLevel, string, LogDetails) error
}

func RegisterHandler(handler LogHandler) {
	RegisterHandlerVerbosity(handler, LogLevelDebug)
}

func RegisterHandlerVerbosity(handler LogHandler, verbosity LogLevel) {
	handlerLock.Lock()
	defer handlerLock.Unlock()
	logHandlers = append(logHandlers, handlerEntry{handler, verbosity})
}

func RemoveHandler(handler LogHandler) {
	handlerLock.Lock()
	defer handlerLock.Unlock()
	newHandlerList := make([]handlerEntry, 0)
	for _, h := range logHandlers {
		if h.LogHandler != handler {
			newHandlerList = append(newHandlerList, h)
		}
	}
	logHandlers = newHandlerList
}
