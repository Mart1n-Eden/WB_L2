package main

import "fmt"

// State - интерфейс состояния
type State interface {
    Handle() string
}

// ConcreteStateA - конкретное состояние A
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() string {
    return "Обработка состояния A"
}

// ConcreteStateB - конкретное состояние B
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() string {
    return "Обработка состояния B"
}

// Context - контекст, который содержит состояние
type Context struct {
    state State
}

func NewContext(state State) *Context {
    return &Context{state: state}
}

func (c *Context) Request() string {
    return c.state.Handle()
}

func (c *Context) SetState(state State) {
    c.state = state
}

func main() {
    context := NewContext(&ConcreteStateA{})
    fmt.Println(context.Request())

    context.SetState(&ConcreteStateB{})
    fmt.Println(context.Request())
}

// Паттерн "Состояние" относится к поведенческим паттернам проектирования и позволяет объекту изменять свое поведение
// в зависимости от своего внутреннего состояния. Он представляет собой альтернативу условным операторам, связанным с
// изменяющимися состояниями объекта. В этом примере интерфейс State определяет общий метод для всех состояний,
// а ConcreteStateA и ConcreteStateB предоставляют конкретные реализации состояний. Класс Context содержит ссылку на текущее
// состояние и вызывает его методы в зависимости от внутреннего состояния. Паттерн "Состояние" позволяет объекту
// переключаться между различными состояниями без изменения его интерфейса, что делает его код более гибким и
// легким для поддержки и изменений.
