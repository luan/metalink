// This file was generated by counterfeiter
package storagefakes

import (
	"sync"

	blobreceipt "github.com/dpb587/blob-receipt"
	"github.com/dpb587/blob-receipt/storage"
)

type FakeStorage struct {
	StringStub        func() string
	stringMutex       sync.RWMutex
	stringArgsForCall []struct{}
	stringReturns     struct {
		result1 string
	}
	ExistsStub        func() (bool, error)
	existsMutex       sync.RWMutex
	existsArgsForCall []struct{}
	existsReturns     struct {
		result1 bool
		result2 error
	}
	GetStub        func() (blobreceipt.BlobReceipt, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct{}
	getReturns     struct {
		result1 blobreceipt.BlobReceipt
		result2 error
	}
	PutStub        func(blobreceipt.BlobReceipt) error
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1 blobreceipt.BlobReceipt
	}
	putReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStorage) String() string {
	fake.stringMutex.Lock()
	fake.stringArgsForCall = append(fake.stringArgsForCall, struct{}{})
	fake.recordInvocation("String", []interface{}{})
	fake.stringMutex.Unlock()
	if fake.StringStub != nil {
		return fake.StringStub()
	}
	return fake.stringReturns.result1
}

func (fake *FakeStorage) StringCallCount() int {
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return len(fake.stringArgsForCall)
}

func (fake *FakeStorage) StringReturns(result1 string) {
	fake.StringStub = nil
	fake.stringReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStorage) Exists() (bool, error) {
	fake.existsMutex.Lock()
	fake.existsArgsForCall = append(fake.existsArgsForCall, struct{}{})
	fake.recordInvocation("Exists", []interface{}{})
	fake.existsMutex.Unlock()
	if fake.ExistsStub != nil {
		return fake.ExistsStub()
	}
	return fake.existsReturns.result1, fake.existsReturns.result2
}

func (fake *FakeStorage) ExistsCallCount() int {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return len(fake.existsArgsForCall)
}

func (fake *FakeStorage) ExistsReturns(result1 bool, result2 error) {
	fake.ExistsStub = nil
	fake.existsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) Get() (blobreceipt.BlobReceipt, error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct{}{})
	fake.recordInvocation("Get", []interface{}{})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub()
	}
	return fake.getReturns.result1, fake.getReturns.result2
}

func (fake *FakeStorage) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeStorage) GetReturns(result1 blobreceipt.BlobReceipt, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 blobreceipt.BlobReceipt
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) Put(arg1 blobreceipt.BlobReceipt) error {
	fake.putMutex.Lock()
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1 blobreceipt.BlobReceipt
	}{arg1})
	fake.recordInvocation("Put", []interface{}{arg1})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1)
	}
	return fake.putReturns.result1
}

func (fake *FakeStorage) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeStorage) PutArgsForCall(i int) blobreceipt.BlobReceipt {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].arg1
}

func (fake *FakeStorage) PutReturns(result1 error) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStorage) recordInvocation(key string, args []interface{}) {
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

var _ storage.Storage = new(FakeStorage)
