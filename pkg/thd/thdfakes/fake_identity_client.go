// Code generated by counterfeiter. DO NOT EDIT.
package thdfakes

import (
	"sync"

	"github.homedepot.com/cd/cloud-runner/pkg/thd"
)

type FakeIdentityClient struct {
	TokenStub        func() (string, error)
	tokenMutex       sync.RWMutex
	tokenArgsForCall []struct {
	}
	tokenReturns struct {
		result1 string
		result2 error
	}
	tokenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	WithClientIDStub        func(string)
	withClientIDMutex       sync.RWMutex
	withClientIDArgsForCall []struct {
		arg1 string
	}
	WithClientSecretStub        func(string)
	withClientSecretMutex       sync.RWMutex
	withClientSecretArgsForCall []struct {
		arg1 string
	}
	WithResourceStub        func(string)
	withResourceMutex       sync.RWMutex
	withResourceArgsForCall []struct {
		arg1 string
	}
	WithTokenEndpointStub        func(string)
	withTokenEndpointMutex       sync.RWMutex
	withTokenEndpointArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIdentityClient) Token() (string, error) {
	fake.tokenMutex.Lock()
	ret, specificReturn := fake.tokenReturnsOnCall[len(fake.tokenArgsForCall)]
	fake.tokenArgsForCall = append(fake.tokenArgsForCall, struct {
	}{})
	stub := fake.TokenStub
	fakeReturns := fake.tokenReturns
	fake.recordInvocation("Token", []interface{}{})
	fake.tokenMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIdentityClient) TokenCallCount() int {
	fake.tokenMutex.RLock()
	defer fake.tokenMutex.RUnlock()
	return len(fake.tokenArgsForCall)
}

func (fake *FakeIdentityClient) TokenCalls(stub func() (string, error)) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = stub
}

func (fake *FakeIdentityClient) TokenReturns(result1 string, result2 error) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = nil
	fake.tokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeIdentityClient) TokenReturnsOnCall(i int, result1 string, result2 error) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = nil
	if fake.tokenReturnsOnCall == nil {
		fake.tokenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.tokenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeIdentityClient) WithClientID(arg1 string) {
	fake.withClientIDMutex.Lock()
	fake.withClientIDArgsForCall = append(fake.withClientIDArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.WithClientIDStub
	fake.recordInvocation("WithClientID", []interface{}{arg1})
	fake.withClientIDMutex.Unlock()
	if stub != nil {
		fake.WithClientIDStub(arg1)
	}
}

func (fake *FakeIdentityClient) WithClientIDCallCount() int {
	fake.withClientIDMutex.RLock()
	defer fake.withClientIDMutex.RUnlock()
	return len(fake.withClientIDArgsForCall)
}

func (fake *FakeIdentityClient) WithClientIDCalls(stub func(string)) {
	fake.withClientIDMutex.Lock()
	defer fake.withClientIDMutex.Unlock()
	fake.WithClientIDStub = stub
}

func (fake *FakeIdentityClient) WithClientIDArgsForCall(i int) string {
	fake.withClientIDMutex.RLock()
	defer fake.withClientIDMutex.RUnlock()
	argsForCall := fake.withClientIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIdentityClient) WithClientSecret(arg1 string) {
	fake.withClientSecretMutex.Lock()
	fake.withClientSecretArgsForCall = append(fake.withClientSecretArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.WithClientSecretStub
	fake.recordInvocation("WithClientSecret", []interface{}{arg1})
	fake.withClientSecretMutex.Unlock()
	if stub != nil {
		fake.WithClientSecretStub(arg1)
	}
}

func (fake *FakeIdentityClient) WithClientSecretCallCount() int {
	fake.withClientSecretMutex.RLock()
	defer fake.withClientSecretMutex.RUnlock()
	return len(fake.withClientSecretArgsForCall)
}

func (fake *FakeIdentityClient) WithClientSecretCalls(stub func(string)) {
	fake.withClientSecretMutex.Lock()
	defer fake.withClientSecretMutex.Unlock()
	fake.WithClientSecretStub = stub
}

func (fake *FakeIdentityClient) WithClientSecretArgsForCall(i int) string {
	fake.withClientSecretMutex.RLock()
	defer fake.withClientSecretMutex.RUnlock()
	argsForCall := fake.withClientSecretArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIdentityClient) WithResource(arg1 string) {
	fake.withResourceMutex.Lock()
	fake.withResourceArgsForCall = append(fake.withResourceArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.WithResourceStub
	fake.recordInvocation("WithResource", []interface{}{arg1})
	fake.withResourceMutex.Unlock()
	if stub != nil {
		fake.WithResourceStub(arg1)
	}
}

func (fake *FakeIdentityClient) WithResourceCallCount() int {
	fake.withResourceMutex.RLock()
	defer fake.withResourceMutex.RUnlock()
	return len(fake.withResourceArgsForCall)
}

func (fake *FakeIdentityClient) WithResourceCalls(stub func(string)) {
	fake.withResourceMutex.Lock()
	defer fake.withResourceMutex.Unlock()
	fake.WithResourceStub = stub
}

func (fake *FakeIdentityClient) WithResourceArgsForCall(i int) string {
	fake.withResourceMutex.RLock()
	defer fake.withResourceMutex.RUnlock()
	argsForCall := fake.withResourceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIdentityClient) WithTokenEndpoint(arg1 string) {
	fake.withTokenEndpointMutex.Lock()
	fake.withTokenEndpointArgsForCall = append(fake.withTokenEndpointArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.WithTokenEndpointStub
	fake.recordInvocation("WithTokenEndpoint", []interface{}{arg1})
	fake.withTokenEndpointMutex.Unlock()
	if stub != nil {
		fake.WithTokenEndpointStub(arg1)
	}
}

func (fake *FakeIdentityClient) WithTokenEndpointCallCount() int {
	fake.withTokenEndpointMutex.RLock()
	defer fake.withTokenEndpointMutex.RUnlock()
	return len(fake.withTokenEndpointArgsForCall)
}

func (fake *FakeIdentityClient) WithTokenEndpointCalls(stub func(string)) {
	fake.withTokenEndpointMutex.Lock()
	defer fake.withTokenEndpointMutex.Unlock()
	fake.WithTokenEndpointStub = stub
}

func (fake *FakeIdentityClient) WithTokenEndpointArgsForCall(i int) string {
	fake.withTokenEndpointMutex.RLock()
	defer fake.withTokenEndpointMutex.RUnlock()
	argsForCall := fake.withTokenEndpointArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIdentityClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tokenMutex.RLock()
	defer fake.tokenMutex.RUnlock()
	fake.withClientIDMutex.RLock()
	defer fake.withClientIDMutex.RUnlock()
	fake.withClientSecretMutex.RLock()
	defer fake.withClientSecretMutex.RUnlock()
	fake.withResourceMutex.RLock()
	defer fake.withResourceMutex.RUnlock()
	fake.withTokenEndpointMutex.RLock()
	defer fake.withTokenEndpointMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIdentityClient) recordInvocation(key string, args []interface{}) {
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

var _ thd.IdentityClient = new(FakeIdentityClient)