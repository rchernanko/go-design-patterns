package singleton

/*

Singleton design pattern - having a unique instance of a type in the entire program

Description:
- This pattern will provide you with a single instance of an object, and guarantee that there are no duplicates
- At the first call to use that instance, it is created and then reused between all the parts in the application that need
to use that particular behaviour
- This creation hides the complexity of creating the object from the caller
- In this example, we are learning the classic singleton implementation for single threaded context. We will see a concurrent
singleton implementation when we reach the chapters about concurrency, because the below implementation is not thread safe!

Usage examples:
- When you want to use the same connection to a database to make every query
- When you open a secure shell (SSH) connection to a server to do a few tasks, and don't want to reopen the connection for each task

Objectives:
- As a general guide, consider using the Singleton pattern when the following rule applies:
	- We need a single, shared value, of some particular type
	- We need to restrict object creation of some type to a single unit along the entire program

Example:
- A unique counter

 */

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton //must be a pointer to a struct as otherwise we cannot do the nil check below in GetInstance()

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton) //using the 'new' keyword, we are creating a pointer to an instance of the singleton type
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}