package main

import "fmt"

// Handler - общий интерфейс для всех обработчиков
type Handler interface {
    SetNext(handler Handler)
    Handle(request int)
}

// ConcreteHandler - базовый тип для всех обработчиков
type ConcreteHandler struct {
    next Handler
}

func (h *ConcreteHandler) SetNext(handler Handler) {
    h.next = handler
}

func (h *ConcreteHandler) Handle(request int) {
    if h.canHandleRequest(request) {
        fmt.Printf("Обработчик %T обрабатывает запрос %d\n", h, request)
    } else if h.next != nil {
        fmt.Printf("Обработчик %T передает запрос %d дальше\n", h, request)
        h.next.Handle(request)
    } else {
        fmt.Printf("Ни один из обработчиков не может обработать запрос %d\n", request)
    }
}

// Этот метод должен быть переопределен в конкретных обработчиках
func (h *ConcreteHandler) canHandleRequest(request int) bool {
    return false
}

// ConcreteHandlerA - конкретный обработчик A
type ConcreteHandlerA struct {
    ConcreteHandler
}

func (h *ConcreteHandlerA) canHandleRequest(request int) bool {
    return request >= 0 && request < 10
}

// ConcreteHandlerB - конкретный обработчик B
type ConcreteHandlerB struct {
    ConcreteHandler
}

func (h *ConcreteHandlerB) canHandleRequest(request int) bool {
    return request >= 10 && request < 20
}

// ConcreteHandlerC - конкретный обработчик C
type ConcreteHandlerC struct {
    ConcreteHandler
}

func (h *ConcreteHandlerC) canHandleRequest(request int) bool {
    return request >= 20
}

func main() {
    handlerA := &ConcreteHandlerA{}
    handlerB := &ConcreteHandlerB{}
    handlerC := &ConcreteHandlerC{}

    handlerA.SetNext(handlerB)
    handlerB.SetNext(handlerC)

    // Последовательно передаем запросы обработчикам
    requests := []int{5, 12, 25}
    for _, request := range requests {
        handlerA.Handle(request)
    }
}
