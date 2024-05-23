package main

import "fmt"

// Element - интерфейс элемента, который может быть посещен
type Element interface {
    Accept(visitor Visitor)
}

// ConcreteElementA - конкретный элемент A
type ConcreteElementA struct{}

func (a *ConcreteElementA) Accept(visitor Visitor) {
    visitor.VisitConcreteElementA(a)
}

// ConcreteElementB - конкретный элемент B
type ConcreteElementB struct{}

func (b *ConcreteElementB) Accept(visitor Visitor) {
    visitor.VisitConcreteElementB(b)
}

// Visitor - интерфейс посетителя
type Visitor interface {
    VisitConcreteElementA(element *ConcreteElementA)
    VisitConcreteElementB(element *ConcreteElementB)
}

// ConcreteVisitor - конкретный посетитель
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
    fmt.Println("Посетитель посещает ConcreteElementA")
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
    fmt.Println("Посетитель посещает ConcreteElementB")
}

// ObjectStructure - структура объектов, которые могут быть посещены
type ObjectStructure struct {
    elements []Element
}

func (os *ObjectStructure) Attach(element Element) {
    os.elements = append(os.elements, element)
}

func (os *ObjectStructure) Detach(element Element) {
    for i, e := range os.elements {
        if e == element {
            os.elements = append(os.elements[:i], os.elements[i+1:]...)
            break
        }
    }
}

func (os *ObjectStructure) Accept(visitor Visitor) {
    for _, element := range os.elements {
        element.Accept(visitor)
    }
}

func main() {
    objectStructure := &ObjectStructure{}

    objectStructure.Attach(&ConcreteElementA{})
    objectStructure.Attach(&ConcreteElementB{})

    visitor := &ConcreteVisitor{}

    objectStructure.Accept(visitor)
}

// Посетитель (Visitor) - это поведенческий паттерн проектирования, который позволяет добавлять новые операции
// над набором объектов, не изменяя их классы. Посетитель позволяет определить новую операцию, не изменяя классы
// объектов, над которыми эта операция выполняется. Он достигается путем выделения этой операции в отдельный класс
// (Посетитель), который реализует методы для обработки каждого конкретного типа объекта. Клиенты могут передавать
// объекты этого класса посетителям и вызывать соответствующие операции, не заботясь о типе объекта.