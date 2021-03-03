// Code generated by counterfeiter. DO NOT EDIT.
package gcloudfakes

import (
	"sync"

	"github.homedepot.com/cd/cloud-runner/pkg/gcloud"
)

type FakeCloudRunCommand struct {
	CombinedOutputStub        func() ([]byte, error)
	combinedOutputMutex       sync.RWMutex
	combinedOutputArgsForCall []struct {
	}
	combinedOutputReturns struct {
		result1 []byte
		result2 error
	}
	combinedOutputReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	StringStub        func() string
	stringMutex       sync.RWMutex
	stringArgsForCall []struct {
	}
	stringReturns struct {
		result1 string
	}
	stringReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCloudRunCommand) CombinedOutput() ([]byte, error) {
	fake.combinedOutputMutex.Lock()
	ret, specificReturn := fake.combinedOutputReturnsOnCall[len(fake.combinedOutputArgsForCall)]
	fake.combinedOutputArgsForCall = append(fake.combinedOutputArgsForCall, struct {
	}{})
	stub := fake.CombinedOutputStub
	fakeReturns := fake.combinedOutputReturns
	fake.recordInvocation("CombinedOutput", []interface{}{})
	fake.combinedOutputMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCloudRunCommand) CombinedOutputCallCount() int {
	fake.combinedOutputMutex.RLock()
	defer fake.combinedOutputMutex.RUnlock()
	return len(fake.combinedOutputArgsForCall)
}

func (fake *FakeCloudRunCommand) CombinedOutputCalls(stub func() ([]byte, error)) {
	fake.combinedOutputMutex.Lock()
	defer fake.combinedOutputMutex.Unlock()
	fake.CombinedOutputStub = stub
}

func (fake *FakeCloudRunCommand) CombinedOutputReturns(result1 []byte, result2 error) {
	fake.combinedOutputMutex.Lock()
	defer fake.combinedOutputMutex.Unlock()
	fake.CombinedOutputStub = nil
	fake.combinedOutputReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudRunCommand) CombinedOutputReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.combinedOutputMutex.Lock()
	defer fake.combinedOutputMutex.Unlock()
	fake.CombinedOutputStub = nil
	if fake.combinedOutputReturnsOnCall == nil {
		fake.combinedOutputReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.combinedOutputReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudRunCommand) String() string {
	fake.stringMutex.Lock()
	ret, specificReturn := fake.stringReturnsOnCall[len(fake.stringArgsForCall)]
	fake.stringArgsForCall = append(fake.stringArgsForCall, struct {
	}{})
	stub := fake.StringStub
	fakeReturns := fake.stringReturns
	fake.recordInvocation("String", []interface{}{})
	fake.stringMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCloudRunCommand) StringCallCount() int {
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return len(fake.stringArgsForCall)
}

func (fake *FakeCloudRunCommand) StringCalls(stub func() string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = stub
}

func (fake *FakeCloudRunCommand) StringReturns(result1 string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = nil
	fake.stringReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCloudRunCommand) StringReturnsOnCall(i int, result1 string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = nil
	if fake.stringReturnsOnCall == nil {
		fake.stringReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.stringReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCloudRunCommand) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.combinedOutputMutex.RLock()
	defer fake.combinedOutputMutex.RUnlock()
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCloudRunCommand) recordInvocation(key string, args []interface{}) {
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

var _ gcloud.CloudRunCommand = new(FakeCloudRunCommand)
