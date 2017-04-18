// This file was generated by counterfeiter
package originfakes

import (
	"sync"

	"github.com/dpb587/metalink/origin"
)

type FakeOriginFactory struct {
	CreateStub        func(string) (origin.Origin, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
	}
	createReturns struct {
		result1 origin.Origin
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOriginFactory) Create(arg1 string) (origin.Origin, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *FakeOriginFactory) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeOriginFactory) CreateArgsForCall(i int) string {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *FakeOriginFactory) CreateReturns(result1 origin.Origin, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 origin.Origin
		result2 error
	}{result1, result2}
}

func (fake *FakeOriginFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeOriginFactory) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ origin.OriginFactory = new(FakeOriginFactory)
