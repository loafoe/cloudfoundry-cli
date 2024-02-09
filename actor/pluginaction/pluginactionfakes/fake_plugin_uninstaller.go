// Code generated by counterfeiter. DO NOT EDIT.
package pluginactionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/pluginaction"
)

type FakePluginUninstaller struct {
	RunStub        func(string, string) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 string
		arg2 string
	}
	runReturns struct {
		result1 error
	}
	runReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePluginUninstaller) Run(arg1 string, arg2 string) error {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Run", []interface{}{arg1, arg2})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.runReturns
	return fakeReturns.result1
}

func (fake *FakePluginUninstaller) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakePluginUninstaller) RunCalls(stub func(string, string) error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakePluginUninstaller) RunArgsForCall(i int) (string, string) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	argsForCall := fake.runArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePluginUninstaller) RunReturns(result1 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePluginUninstaller) RunReturnsOnCall(i int, result1 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePluginUninstaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePluginUninstaller) recordInvocation(key string, args []interface{}) {
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

var _ pluginaction.PluginUninstaller = new(FakePluginUninstaller)
