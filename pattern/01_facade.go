package main

import "fmt"

// SubsystemA - Подсистема A
type SubsystemA struct{}

func (a *SubsystemA) OperationA() {
    fmt.Println("Subsystem A: выполнение операции A")
}

// SubsystemB - Подсистема B
type SubsystemB struct{}

func (b *SubsystemB) OperationB() {
    fmt.Println("Subsystem B: выполнение операции B")
}

// SubsystemC - Подсистема C
type SubsystemC struct{}

func (c *SubsystemC) OperationC() {
    fmt.Println("Subsystem C: выполнение операции C")
}

// Facade - Фасад
type Facade struct {
    subsystemA *SubsystemA
    subsystemB *SubsystemB
    subsystemC *SubsystemC
}

func NewFacade() *Facade {
    return &Facade{
        subsystemA: &SubsystemA{},
        subsystemB: &SubsystemB{},
        subsystemC: &SubsystemC{},
    }
}

func (f *Facade) Operation1() {
    fmt.Println("Facade: выполнение операции 1")
    f.subsystemA.OperationA()
    f.subsystemB.OperationB()
}

func (f *Facade) Operation2() {
    fmt.Println("Facade: выполнение операции 2")
    f.subsystemB.OperationB()
    f.subsystemC.OperationC()
}

func main() {
    facade := NewFacade()
    facade.Operation1()
    facade.Operation2()
}

// Паттерн Фасад используется для предоставления унифицированного интерфейса
// для набора взаимосвязанных подсистем, упрощая тем самым их использование.
// Фасад скрывает сложность взаимодействия между подсистемами и предоставляет
// более простой интерфейс для клиента, что уменьшает зависимости между
// клиентским кодом и подсистемами, а также снижает сложность их взаимодействия.
// В этом примере Facade предоставляет два метода Operation1 и Operation2,
// которые представляют собой более простые операции, скрывая детали
// внутренней реализации подсистем A, B и C.
