package services

import (
	"reflect"
)

var container = newContainer()

func GetContainer() *Container {
	return container
}

func CreateContainer() *Container {
	return newContainer()
}

func AddService(descriptor ServiceDescriptor) *Container {
	container.Add(descriptor)
	return container
}

func TryAddService(descriptor ServiceDescriptor) *Container {
	container.TryAdd(descriptor)
	return container
}

func Add[T interface{}](scope ServiceScope, creator func(*Container) *T) {
	AddService(ServiceDescriptor{
		ServiceType: getServiceType[T](),
		Creator:     func(c *Container) interface{} { return creator(c) },
		Scope:       scope,
	})
}

func TryAdd[T interface{}](scope ServiceScope, creator func(*Container) *T) {
	TryAddService(ServiceDescriptor{
		ServiceType: getServiceType[T](),
		Creator:     func(c *Container) interface{} { return creator(c) },
		Scope:       scope,
	})
}

func AddTransient[T interface{}](creator func(*Container) *T) {
	Add(Transient, creator)
}
func TryAddTransient[T interface{}](creator func(*Container) *T) {
	TryAdd(Transient, creator)
}

func AddSingleton[T interface{}](creator func(*Container) *T) {
	Add(Singleton, creator)
}
func TryAddSingleton[T interface{}](creator func(*Container) *T) {
	TryAdd(Singleton, creator)
}

func AddValue(value interface{}) {
	AddService(ServiceDescriptor{
		ServiceType: getInterfaceType(value),
		Value:       value,
		Scope:       Singleton,
	})
}

func TryAddValue(value interface{}) {
	TryAddService(ServiceDescriptor{
		ServiceType: getInterfaceType(value),
		Value:       value,
		Scope:       Singleton,
	})
}

func AddByType[T interface{}](serviceType reflect.Type, scope ServiceScope, creator func(*Container) *T) {
	AddService(ServiceDescriptor{
		ServiceType: serviceType,
		Creator:     func(c *Container) interface{} { return creator(c) },
		Scope:       scope,
	})
}

func TryAddByType[T interface{}](serviceType reflect.Type, scope ServiceScope, creator func(*Container) *T) {
	TryAddService(ServiceDescriptor{
		ServiceType: serviceType,
		Creator:     func(c *Container) interface{} { return creator(c) },
		Scope:       scope,
	})
}

func Get[T interface{}]() *T {
	return container.Get(getServiceType[T]()).(*T)
}

func GetByType[T interface{}](serviceType reflect.Type) *T {
	return container.Get(serviceType).(*T)
}

func getServiceType[T interface{}]() reflect.Type {
	return reflect.TypeOf((*T)(nil))
}

func getInterfaceType(value interface{}) reflect.Type {
	return reflect.TypeOf(value)
}
