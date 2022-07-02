package di

import "reflect"

type ServiceScope int

const (
	Transient ServiceScope = 0
	Singleton ServiceScope = 1
	Scoped    ServiceScope = 2
)

type ServiceDescriptor struct {
	ServiceType reflect.Type
	Creator     func(c *Container) interface{}
	Scope       ServiceScope
	Value       interface{}
}
