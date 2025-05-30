package chain

import "fmt"

// Request 请求结构
type Request struct {
	Level   int
	Message string
}

// Handler 处理器接口
type Handler interface {
	SetNext(handler Handler)
	Handle(request *Request)
}

// BaseHandler 基础处理器
type BaseHandler struct {
	next Handler
}

func (bh *BaseHandler) SetNext(handler Handler) {
	bh.next = handler
}

func (bh *BaseHandler) Handle(request *Request) {
	if bh.next != nil {
		bh.next.Handle(request)
	}
}

// InfoHandler 信息处理器
type InfoHandler struct {
	BaseHandler
}

func (ih *InfoHandler) Handle(request *Request) {
	if request.Level <= 1 {
		fmt.Printf("INFO: %s\n", request.Message)
	} else {
		ih.BaseHandler.Handle(request)
	}
}

// WarningHandler 警告处理器
type WarningHandler struct {
	BaseHandler
}

func (wh *WarningHandler) Handle(request *Request) {
	if request.Level == 2 {
		fmt.Printf("WARNING: %s\n", request.Message)
	} else {
		wh.BaseHandler.Handle(request)
	}
}

// ErrorHandler 错误处理器
type ErrorHandler struct {
	BaseHandler
}

func (eh *ErrorHandler) Handle(request *Request) {
	if request.Level >= 3 {
		fmt.Printf("ERROR: %s\n", request.Message)
	} else {
		eh.BaseHandler.Handle(request)
	}
}

// Logger 日志记录器
type Logger struct {
	chain Handler
}

func NewLogger() *Logger {
	infoHandler := &InfoHandler{}
	warningHandler := &WarningHandler{}
	errorHandler := &ErrorHandler{}

	// 设置责任链
	infoHandler.SetNext(warningHandler)
	warningHandler.SetNext(errorHandler)

	return &Logger{chain: infoHandler}
}

func (l *Logger) Log(level int, message string) {
	request := &Request{
		Level:   level,
		Message: message,
	}
	l.chain.Handle(request)
}