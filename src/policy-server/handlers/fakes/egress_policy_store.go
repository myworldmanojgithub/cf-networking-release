// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/store"
	"sync"
)

type EgressPolicyStore struct {
	AllStub        func() ([]store.EgressPolicy, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct{}
	allReturns     struct {
		result1 []store.EgressPolicy
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	ByGuidsStub        func(ids []string) ([]store.EgressPolicy, error)
	byGuidsMutex       sync.RWMutex
	byGuidsArgsForCall []struct {
		ids []string
	}
	byGuidsReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	byGuidsReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	CreateStub        func(egressPolicies []store.EgressPolicy) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		egressPolicies []store.EgressPolicy
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressPolicyStore) All() ([]store.EgressPolicy, error) {
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

func (fake *EgressPolicyStore) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *EgressPolicyStore) AllReturns(result1 []store.EgressPolicy, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) AllReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) ByGuids(ids []string) ([]store.EgressPolicy, error) {
	var idsCopy []string
	if ids != nil {
		idsCopy = make([]string, len(ids))
		copy(idsCopy, ids)
	}
	fake.byGuidsMutex.Lock()
	ret, specificReturn := fake.byGuidsReturnsOnCall[len(fake.byGuidsArgsForCall)]
	fake.byGuidsArgsForCall = append(fake.byGuidsArgsForCall, struct {
		ids []string
	}{idsCopy})
	fake.recordInvocation("ByGuids", []interface{}{idsCopy})
	fake.byGuidsMutex.Unlock()
	if fake.ByGuidsStub != nil {
		return fake.ByGuidsStub(ids)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.byGuidsReturns.result1, fake.byGuidsReturns.result2
}

func (fake *EgressPolicyStore) ByGuidsCallCount() int {
	fake.byGuidsMutex.RLock()
	defer fake.byGuidsMutex.RUnlock()
	return len(fake.byGuidsArgsForCall)
}

func (fake *EgressPolicyStore) ByGuidsArgsForCall(i int) []string {
	fake.byGuidsMutex.RLock()
	defer fake.byGuidsMutex.RUnlock()
	return fake.byGuidsArgsForCall[i].ids
}

func (fake *EgressPolicyStore) ByGuidsReturns(result1 []store.EgressPolicy, result2 error) {
	fake.ByGuidsStub = nil
	fake.byGuidsReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) ByGuidsReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.ByGuidsStub = nil
	if fake.byGuidsReturnsOnCall == nil {
		fake.byGuidsReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.byGuidsReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) Create(egressPolicies []store.EgressPolicy) error {
	var egressPoliciesCopy []store.EgressPolicy
	if egressPolicies != nil {
		egressPoliciesCopy = make([]store.EgressPolicy, len(egressPolicies))
		copy(egressPoliciesCopy, egressPolicies)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		egressPolicies []store.EgressPolicy
	}{egressPoliciesCopy})
	fake.recordInvocation("Create", []interface{}{egressPoliciesCopy})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(egressPolicies)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *EgressPolicyStore) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *EgressPolicyStore) CreateArgsForCall(i int) []store.EgressPolicy {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].egressPolicies
}

func (fake *EgressPolicyStore) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyStore) CreateReturnsOnCall(i int, result1 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.byGuidsMutex.RLock()
	defer fake.byGuidsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EgressPolicyStore) recordInvocation(key string, args []interface{}) {
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
