package services

import (
	"fmt"
	"reflect"
)

func newContainer() *Container {
	return &Container{
		Services: make([]ServiceDescriptor, 0),
		values:   make(map[reflect.Type]interface{}),
	}
}

type Container struct {
	Services []ServiceDescriptor
	values   map[reflect.Type]interface{}
}

func (c *Container) Get(serviceType reflect.Type) interface{} {
	descriptor := c.firstOrDefault(serviceType)
	if descriptor == nil {
		panic(fmt.Sprintf("service %s not found", serviceType))
	}
	var value interface{}
	if descriptor.Scope == Singleton {
		v, ok := c.values[serviceType]
		if ok {
			value = v
		} else {
			value = c.create(descriptor)
			c.values[serviceType] = value
		}
	} else if descriptor.Scope == Transient {
		value = c.create(descriptor)
	}
	return value
}

func (c *Container) Add(descriptor ServiceDescriptor) {
	index := c.find(&descriptor.ServiceType)
	if index >= 0 {
		c.Services[index] = descriptor
	} else {
		c.Services = append(c.Services, descriptor)
	}
}

func (c *Container) TryAdd(descriptor ServiceDescriptor) {
	index := c.find(&descriptor.ServiceType)
	if index < 0 {
		c.Services = append(c.Services, descriptor)
	}
}

func (c *Container) firstOrDefault(serviceType reflect.Type) *ServiceDescriptor {
	for _, v := range c.Services {
		if v.ServiceType == serviceType {
			return &v
		}
	}
	return nil
}

func (c *Container) find(serviceType *reflect.Type) int {
	for i, v := range c.Services {
		if v.ServiceType == *serviceType {
			return i
		}
	}
	return -1
}

func (c *Container) create(descriptor *ServiceDescriptor) interface{} {
	if descriptor.Value != nil {
		return descriptor.Value
	}
	return descriptor.Creator(c)
}
