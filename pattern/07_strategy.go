package main

import "fmt"

// Strategy - интерфейс стратегии
type Strategy interface {
	ExecuteStrategy(int, int) int
}

// ConcreteStrategyAdd - конкретная стратегия сложения
type ConcreteStrategyAdd struct{}

func (s *ConcreteStrategyAdd) ExecuteStrategy(a, b int) int {
	return a + b
}

// ConcreteStrategySubtract - конкретная стратегия вычитания
type ConcreteStrategySubtract struct{}

func (s *ConcreteStrategySubtract) ExecuteStrategy(a, b int) int {
	return a - b
}

// Context - контекст, использующий стратегию
type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.ExecuteStrategy(a, b)
}

func main() {
	context := NewContext(&ConcreteStrategyAdd{})
	result := context.ExecuteStrategy(10, 5)
	fmt.Println("Результат сложения:", result)

	context.SetStrategy(&ConcreteStrategySubtract{})
	result = context.ExecuteStrategy(10, 5)
	fmt.Println("Результат вычитания:", result)
}

// Паттерн "Стратегия" относится к поведенческим паттернам проектирования и позволяет определить семейство алгоритмов,
// инкапсулировать каждый из них и сделать их взаимозаменяемыми. Он предоставляет возможность выбирать алгоритм
// во время выполнения программы. В этом примере интерфейс Strategy определяет общий метод для всех стратегий,
// а ConcreteStrategyAdd и ConcreteStrategySubtract предоставляют конкретные реализации алгоритмов. Класс Context использует
// стратегию для выполнения определенной операции. Паттерн "Стратегия" позволяет клиентам выбирать нужный алгоритм
// динамически, а также обеспечивает гибкость и расширяемость программы за счет возможности добавления новых
// стратегий без изменения кода контекста.
