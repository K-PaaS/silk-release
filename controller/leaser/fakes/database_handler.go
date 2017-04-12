// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/silk/controller"
)

type DatabaseHandler struct {
	MigrateStub        func() (int, error)
	migrateMutex       sync.RWMutex
	migrateArgsForCall []struct{}
	migrateReturns     struct {
		result1 int
		result2 error
	}
	migrateReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	AddEntryStub        func(*controller.Lease) error
	addEntryMutex       sync.RWMutex
	addEntryArgsForCall []struct {
		arg1 *controller.Lease
	}
	addEntryReturns struct {
		result1 error
	}
	addEntryReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteEntryStub        func(string) error
	deleteEntryMutex       sync.RWMutex
	deleteEntryArgsForCall []struct {
		arg1 string
	}
	deleteEntryReturns struct {
		result1 error
	}
	deleteEntryReturnsOnCall map[int]struct {
		result1 error
	}
	LeaseForUnderlayIPStub        func(string) (*controller.Lease, error)
	leaseForUnderlayIPMutex       sync.RWMutex
	leaseForUnderlayIPArgsForCall []struct {
		arg1 string
	}
	leaseForUnderlayIPReturns struct {
		result1 *controller.Lease
		result2 error
	}
	leaseForUnderlayIPReturnsOnCall map[int]struct {
		result1 *controller.Lease
		result2 error
	}
	AllStub        func() ([]controller.Lease, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct{}
	allReturns     struct {
		result1 []controller.Lease
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []controller.Lease
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DatabaseHandler) Migrate() (int, error) {
	fake.migrateMutex.Lock()
	ret, specificReturn := fake.migrateReturnsOnCall[len(fake.migrateArgsForCall)]
	fake.migrateArgsForCall = append(fake.migrateArgsForCall, struct{}{})
	fake.recordInvocation("Migrate", []interface{}{})
	fake.migrateMutex.Unlock()
	if fake.MigrateStub != nil {
		return fake.MigrateStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.migrateReturns.result1, fake.migrateReturns.result2
}

func (fake *DatabaseHandler) MigrateCallCount() int {
	fake.migrateMutex.RLock()
	defer fake.migrateMutex.RUnlock()
	return len(fake.migrateArgsForCall)
}

func (fake *DatabaseHandler) MigrateReturns(result1 int, result2 error) {
	fake.MigrateStub = nil
	fake.migrateReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) MigrateReturnsOnCall(i int, result1 int, result2 error) {
	fake.MigrateStub = nil
	if fake.migrateReturnsOnCall == nil {
		fake.migrateReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.migrateReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) AddEntry(arg1 *controller.Lease) error {
	fake.addEntryMutex.Lock()
	ret, specificReturn := fake.addEntryReturnsOnCall[len(fake.addEntryArgsForCall)]
	fake.addEntryArgsForCall = append(fake.addEntryArgsForCall, struct {
		arg1 *controller.Lease
	}{arg1})
	fake.recordInvocation("AddEntry", []interface{}{arg1})
	fake.addEntryMutex.Unlock()
	if fake.AddEntryStub != nil {
		return fake.AddEntryStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addEntryReturns.result1
}

func (fake *DatabaseHandler) AddEntryCallCount() int {
	fake.addEntryMutex.RLock()
	defer fake.addEntryMutex.RUnlock()
	return len(fake.addEntryArgsForCall)
}

func (fake *DatabaseHandler) AddEntryArgsForCall(i int) *controller.Lease {
	fake.addEntryMutex.RLock()
	defer fake.addEntryMutex.RUnlock()
	return fake.addEntryArgsForCall[i].arg1
}

func (fake *DatabaseHandler) AddEntryReturns(result1 error) {
	fake.AddEntryStub = nil
	fake.addEntryReturns = struct {
		result1 error
	}{result1}
}

func (fake *DatabaseHandler) AddEntryReturnsOnCall(i int, result1 error) {
	fake.AddEntryStub = nil
	if fake.addEntryReturnsOnCall == nil {
		fake.addEntryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addEntryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DatabaseHandler) DeleteEntry(arg1 string) error {
	fake.deleteEntryMutex.Lock()
	ret, specificReturn := fake.deleteEntryReturnsOnCall[len(fake.deleteEntryArgsForCall)]
	fake.deleteEntryArgsForCall = append(fake.deleteEntryArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteEntry", []interface{}{arg1})
	fake.deleteEntryMutex.Unlock()
	if fake.DeleteEntryStub != nil {
		return fake.DeleteEntryStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteEntryReturns.result1
}

func (fake *DatabaseHandler) DeleteEntryCallCount() int {
	fake.deleteEntryMutex.RLock()
	defer fake.deleteEntryMutex.RUnlock()
	return len(fake.deleteEntryArgsForCall)
}

func (fake *DatabaseHandler) DeleteEntryArgsForCall(i int) string {
	fake.deleteEntryMutex.RLock()
	defer fake.deleteEntryMutex.RUnlock()
	return fake.deleteEntryArgsForCall[i].arg1
}

func (fake *DatabaseHandler) DeleteEntryReturns(result1 error) {
	fake.DeleteEntryStub = nil
	fake.deleteEntryReturns = struct {
		result1 error
	}{result1}
}

func (fake *DatabaseHandler) DeleteEntryReturnsOnCall(i int, result1 error) {
	fake.DeleteEntryStub = nil
	if fake.deleteEntryReturnsOnCall == nil {
		fake.deleteEntryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteEntryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DatabaseHandler) LeaseForUnderlayIP(arg1 string) (*controller.Lease, error) {
	fake.leaseForUnderlayIPMutex.Lock()
	ret, specificReturn := fake.leaseForUnderlayIPReturnsOnCall[len(fake.leaseForUnderlayIPArgsForCall)]
	fake.leaseForUnderlayIPArgsForCall = append(fake.leaseForUnderlayIPArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("LeaseForUnderlayIP", []interface{}{arg1})
	fake.leaseForUnderlayIPMutex.Unlock()
	if fake.LeaseForUnderlayIPStub != nil {
		return fake.LeaseForUnderlayIPStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.leaseForUnderlayIPReturns.result1, fake.leaseForUnderlayIPReturns.result2
}

func (fake *DatabaseHandler) LeaseForUnderlayIPCallCount() int {
	fake.leaseForUnderlayIPMutex.RLock()
	defer fake.leaseForUnderlayIPMutex.RUnlock()
	return len(fake.leaseForUnderlayIPArgsForCall)
}

func (fake *DatabaseHandler) LeaseForUnderlayIPArgsForCall(i int) string {
	fake.leaseForUnderlayIPMutex.RLock()
	defer fake.leaseForUnderlayIPMutex.RUnlock()
	return fake.leaseForUnderlayIPArgsForCall[i].arg1
}

func (fake *DatabaseHandler) LeaseForUnderlayIPReturns(result1 *controller.Lease, result2 error) {
	fake.LeaseForUnderlayIPStub = nil
	fake.leaseForUnderlayIPReturns = struct {
		result1 *controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) LeaseForUnderlayIPReturnsOnCall(i int, result1 *controller.Lease, result2 error) {
	fake.LeaseForUnderlayIPStub = nil
	if fake.leaseForUnderlayIPReturnsOnCall == nil {
		fake.leaseForUnderlayIPReturnsOnCall = make(map[int]struct {
			result1 *controller.Lease
			result2 error
		})
	}
	fake.leaseForUnderlayIPReturnsOnCall[i] = struct {
		result1 *controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) All() ([]controller.Lease, error) {
	fake.allMutex.Lock()
	ret, specificReturn := fake.allReturnsOnCall[len(fake.allArgsForCall)]
	fake.allArgsForCall = append(fake.allArgsForCall, struct{}{})
	fake.recordInvocation("All", []interface{}{})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allReturns.result1, fake.allReturns.result2
}

func (fake *DatabaseHandler) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *DatabaseHandler) AllReturns(result1 []controller.Lease, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) AllReturnsOnCall(i int, result1 []controller.Lease, result2 error) {
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []controller.Lease
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.migrateMutex.RLock()
	defer fake.migrateMutex.RUnlock()
	fake.addEntryMutex.RLock()
	defer fake.addEntryMutex.RUnlock()
	fake.deleteEntryMutex.RLock()
	defer fake.deleteEntryMutex.RUnlock()
	fake.leaseForUnderlayIPMutex.RLock()
	defer fake.leaseForUnderlayIPMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return fake.invocations
}

func (fake *DatabaseHandler) recordInvocation(key string, args []interface{}) {
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