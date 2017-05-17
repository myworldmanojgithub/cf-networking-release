// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"lib/rules"
	"sync"

	"code.cloudfoundry.org/garden"
)

type NetOutRuleConverter struct {
	ConvertStub        func(rule garden.NetOutRule, logChainName string, logging bool) []rules.IPTablesRule
	convertMutex       sync.RWMutex
	convertArgsForCall []struct {
		rule         garden.NetOutRule
		logChainName string
		logging      bool
	}
	convertReturns struct {
		result1 []rules.IPTablesRule
	}
	convertReturnsOnCall map[int]struct {
		result1 []rules.IPTablesRule
	}
	BulkConvertStub        func(rules []garden.NetOutRule, logChainName string, logging bool) []rules.IPTablesRule
	bulkConvertMutex       sync.RWMutex
	bulkConvertArgsForCall []struct {
		rules        []garden.NetOutRule
		logChainName string
		logging      bool
	}
	bulkConvertReturns struct {
		result1 []rules.IPTablesRule
	}
	bulkConvertReturnsOnCall map[int]struct {
		result1 []rules.IPTablesRule
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NetOutRuleConverter) Convert(rule garden.NetOutRule, logChainName string, logging bool) []rules.IPTablesRule {
	fake.convertMutex.Lock()
	ret, specificReturn := fake.convertReturnsOnCall[len(fake.convertArgsForCall)]
	fake.convertArgsForCall = append(fake.convertArgsForCall, struct {
		rule         garden.NetOutRule
		logChainName string
		logging      bool
	}{rule, logChainName, logging})
	fake.recordInvocation("Convert", []interface{}{rule, logChainName, logging})
	fake.convertMutex.Unlock()
	if fake.ConvertStub != nil {
		return fake.ConvertStub(rule, logChainName, logging)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.convertReturns.result1
}

func (fake *NetOutRuleConverter) ConvertCallCount() int {
	fake.convertMutex.RLock()
	defer fake.convertMutex.RUnlock()
	return len(fake.convertArgsForCall)
}

func (fake *NetOutRuleConverter) ConvertArgsForCall(i int) (garden.NetOutRule, string, bool) {
	fake.convertMutex.RLock()
	defer fake.convertMutex.RUnlock()
	return fake.convertArgsForCall[i].rule, fake.convertArgsForCall[i].logChainName, fake.convertArgsForCall[i].logging
}

func (fake *NetOutRuleConverter) ConvertReturns(result1 []rules.IPTablesRule) {
	fake.ConvertStub = nil
	fake.convertReturns = struct {
		result1 []rules.IPTablesRule
	}{result1}
}

func (fake *NetOutRuleConverter) ConvertReturnsOnCall(i int, result1 []rules.IPTablesRule) {
	fake.ConvertStub = nil
	if fake.convertReturnsOnCall == nil {
		fake.convertReturnsOnCall = make(map[int]struct {
			result1 []rules.IPTablesRule
		})
	}
	fake.convertReturnsOnCall[i] = struct {
		result1 []rules.IPTablesRule
	}{result1}
}

func (fake *NetOutRuleConverter) BulkConvert(rules []garden.NetOutRule, logChainName string, logging bool) []rules.IPTablesRule {
	var rulesCopy []garden.NetOutRule
	if rules != nil {
		rulesCopy = make([]garden.NetOutRule, len(rules))
		copy(rulesCopy, rules)
	}
	fake.bulkConvertMutex.Lock()
	ret, specificReturn := fake.bulkConvertReturnsOnCall[len(fake.bulkConvertArgsForCall)]
	fake.bulkConvertArgsForCall = append(fake.bulkConvertArgsForCall, struct {
		rules        []garden.NetOutRule
		logChainName string
		logging      bool
	}{rulesCopy, logChainName, logging})
	fake.recordInvocation("BulkConvert", []interface{}{rulesCopy, logChainName, logging})
	fake.bulkConvertMutex.Unlock()
	if fake.BulkConvertStub != nil {
		return fake.BulkConvertStub(rules, logChainName, logging)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.bulkConvertReturns.result1
}

func (fake *NetOutRuleConverter) BulkConvertCallCount() int {
	fake.bulkConvertMutex.RLock()
	defer fake.bulkConvertMutex.RUnlock()
	return len(fake.bulkConvertArgsForCall)
}

func (fake *NetOutRuleConverter) BulkConvertArgsForCall(i int) ([]garden.NetOutRule, string, bool) {
	fake.bulkConvertMutex.RLock()
	defer fake.bulkConvertMutex.RUnlock()
	return fake.bulkConvertArgsForCall[i].rules, fake.bulkConvertArgsForCall[i].logChainName, fake.bulkConvertArgsForCall[i].logging
}

func (fake *NetOutRuleConverter) BulkConvertReturns(result1 []rules.IPTablesRule) {
	fake.BulkConvertStub = nil
	fake.bulkConvertReturns = struct {
		result1 []rules.IPTablesRule
	}{result1}
}

func (fake *NetOutRuleConverter) BulkConvertReturnsOnCall(i int, result1 []rules.IPTablesRule) {
	fake.BulkConvertStub = nil
	if fake.bulkConvertReturnsOnCall == nil {
		fake.bulkConvertReturnsOnCall = make(map[int]struct {
			result1 []rules.IPTablesRule
		})
	}
	fake.bulkConvertReturnsOnCall[i] = struct {
		result1 []rules.IPTablesRule
	}{result1}
}

func (fake *NetOutRuleConverter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.convertMutex.RLock()
	defer fake.convertMutex.RUnlock()
	fake.bulkConvertMutex.RLock()
	defer fake.bulkConvertMutex.RUnlock()
	return fake.invocations
}

func (fake *NetOutRuleConverter) recordInvocation(key string, args []interface{}) {
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
