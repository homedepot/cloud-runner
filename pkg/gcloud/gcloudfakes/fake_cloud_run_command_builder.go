// Code generated by counterfeiter. DO NOT EDIT.
package gcloudfakes

import (
	"sync"

	"github.homedepot.com/cd/cloud-runner/pkg/gcloud"
)

type FakeCloudRunCommandBuilder struct {
	AllowUnauthenticatedStub        func(bool) gcloud.CloudRunCommandBuilder
	allowUnauthenticatedMutex       sync.RWMutex
	allowUnauthenticatedArgsForCall []struct {
		arg1 bool
	}
	allowUnauthenticatedReturns struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	allowUnauthenticatedReturnsOnCall map[int]struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	BuildStub        func() (gcloud.CloudRunCommand, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
	}
	buildReturns struct {
		result1 gcloud.CloudRunCommand
		result2 error
	}
	buildReturnsOnCall map[int]struct {
		result1 gcloud.CloudRunCommand
		result2 error
	}
	ImageStub        func(string) gcloud.CloudRunCommandBuilder
	imageMutex       sync.RWMutex
	imageArgsForCall []struct {
		arg1 string
	}
	imageReturns struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	imageReturnsOnCall map[int]struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	MaxInstancesStub        func(int) gcloud.CloudRunCommandBuilder
	maxInstancesMutex       sync.RWMutex
	maxInstancesArgsForCall []struct {
		arg1 int
	}
	maxInstancesReturns struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	maxInstancesReturnsOnCall map[int]struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	MemoryStub        func(string) gcloud.CloudRunCommandBuilder
	memoryMutex       sync.RWMutex
	memoryArgsForCall []struct {
		arg1 string
	}
	memoryReturns struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	memoryReturnsOnCall map[int]struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	ProjectIDStub        func(string) gcloud.CloudRunCommandBuilder
	projectIDMutex       sync.RWMutex
	projectIDArgsForCall []struct {
		arg1 string
	}
	projectIDReturns struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	projectIDReturnsOnCall map[int]struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	RegionStub        func(string) gcloud.CloudRunCommandBuilder
	regionMutex       sync.RWMutex
	regionArgsForCall []struct {
		arg1 string
	}
	regionReturns struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	regionReturnsOnCall map[int]struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	ServiceNameStub        func(string) gcloud.CloudRunCommandBuilder
	serviceNameMutex       sync.RWMutex
	serviceNameArgsForCall []struct {
		arg1 string
	}
	serviceNameReturns struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	serviceNameReturnsOnCall map[int]struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	VPCConnectorStub        func(string) gcloud.CloudRunCommandBuilder
	vPCConnectorMutex       sync.RWMutex
	vPCConnectorArgsForCall []struct {
		arg1 string
	}
	vPCConnectorReturns struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	vPCConnectorReturnsOnCall map[int]struct {
		result1 gcloud.CloudRunCommandBuilder
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCloudRunCommandBuilder) AllowUnauthenticated(arg1 bool) gcloud.CloudRunCommandBuilder {
	fake.allowUnauthenticatedMutex.Lock()
	ret, specificReturn := fake.allowUnauthenticatedReturnsOnCall[len(fake.allowUnauthenticatedArgsForCall)]
	fake.allowUnauthenticatedArgsForCall = append(fake.allowUnauthenticatedArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.AllowUnauthenticatedStub
	fakeReturns := fake.allowUnauthenticatedReturns
	fake.recordInvocation("AllowUnauthenticated", []interface{}{arg1})
	fake.allowUnauthenticatedMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCloudRunCommandBuilder) AllowUnauthenticatedCallCount() int {
	fake.allowUnauthenticatedMutex.RLock()
	defer fake.allowUnauthenticatedMutex.RUnlock()
	return len(fake.allowUnauthenticatedArgsForCall)
}

func (fake *FakeCloudRunCommandBuilder) AllowUnauthenticatedCalls(stub func(bool) gcloud.CloudRunCommandBuilder) {
	fake.allowUnauthenticatedMutex.Lock()
	defer fake.allowUnauthenticatedMutex.Unlock()
	fake.AllowUnauthenticatedStub = stub
}

func (fake *FakeCloudRunCommandBuilder) AllowUnauthenticatedArgsForCall(i int) bool {
	fake.allowUnauthenticatedMutex.RLock()
	defer fake.allowUnauthenticatedMutex.RUnlock()
	argsForCall := fake.allowUnauthenticatedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCloudRunCommandBuilder) AllowUnauthenticatedReturns(result1 gcloud.CloudRunCommandBuilder) {
	fake.allowUnauthenticatedMutex.Lock()
	defer fake.allowUnauthenticatedMutex.Unlock()
	fake.AllowUnauthenticatedStub = nil
	fake.allowUnauthenticatedReturns = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) AllowUnauthenticatedReturnsOnCall(i int, result1 gcloud.CloudRunCommandBuilder) {
	fake.allowUnauthenticatedMutex.Lock()
	defer fake.allowUnauthenticatedMutex.Unlock()
	fake.AllowUnauthenticatedStub = nil
	if fake.allowUnauthenticatedReturnsOnCall == nil {
		fake.allowUnauthenticatedReturnsOnCall = make(map[int]struct {
			result1 gcloud.CloudRunCommandBuilder
		})
	}
	fake.allowUnauthenticatedReturnsOnCall[i] = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) Build() (gcloud.CloudRunCommand, error) {
	fake.buildMutex.Lock()
	ret, specificReturn := fake.buildReturnsOnCall[len(fake.buildArgsForCall)]
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
	}{})
	stub := fake.BuildStub
	fakeReturns := fake.buildReturns
	fake.recordInvocation("Build", []interface{}{})
	fake.buildMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCloudRunCommandBuilder) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakeCloudRunCommandBuilder) BuildCalls(stub func() (gcloud.CloudRunCommand, error)) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = stub
}

func (fake *FakeCloudRunCommandBuilder) BuildReturns(result1 gcloud.CloudRunCommand, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 gcloud.CloudRunCommand
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudRunCommandBuilder) BuildReturnsOnCall(i int, result1 gcloud.CloudRunCommand, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 gcloud.CloudRunCommand
			result2 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 gcloud.CloudRunCommand
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudRunCommandBuilder) Image(arg1 string) gcloud.CloudRunCommandBuilder {
	fake.imageMutex.Lock()
	ret, specificReturn := fake.imageReturnsOnCall[len(fake.imageArgsForCall)]
	fake.imageArgsForCall = append(fake.imageArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ImageStub
	fakeReturns := fake.imageReturns
	fake.recordInvocation("Image", []interface{}{arg1})
	fake.imageMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCloudRunCommandBuilder) ImageCallCount() int {
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	return len(fake.imageArgsForCall)
}

func (fake *FakeCloudRunCommandBuilder) ImageCalls(stub func(string) gcloud.CloudRunCommandBuilder) {
	fake.imageMutex.Lock()
	defer fake.imageMutex.Unlock()
	fake.ImageStub = stub
}

func (fake *FakeCloudRunCommandBuilder) ImageArgsForCall(i int) string {
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	argsForCall := fake.imageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCloudRunCommandBuilder) ImageReturns(result1 gcloud.CloudRunCommandBuilder) {
	fake.imageMutex.Lock()
	defer fake.imageMutex.Unlock()
	fake.ImageStub = nil
	fake.imageReturns = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) ImageReturnsOnCall(i int, result1 gcloud.CloudRunCommandBuilder) {
	fake.imageMutex.Lock()
	defer fake.imageMutex.Unlock()
	fake.ImageStub = nil
	if fake.imageReturnsOnCall == nil {
		fake.imageReturnsOnCall = make(map[int]struct {
			result1 gcloud.CloudRunCommandBuilder
		})
	}
	fake.imageReturnsOnCall[i] = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) MaxInstances(arg1 int) gcloud.CloudRunCommandBuilder {
	fake.maxInstancesMutex.Lock()
	ret, specificReturn := fake.maxInstancesReturnsOnCall[len(fake.maxInstancesArgsForCall)]
	fake.maxInstancesArgsForCall = append(fake.maxInstancesArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.MaxInstancesStub
	fakeReturns := fake.maxInstancesReturns
	fake.recordInvocation("MaxInstances", []interface{}{arg1})
	fake.maxInstancesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCloudRunCommandBuilder) MaxInstancesCallCount() int {
	fake.maxInstancesMutex.RLock()
	defer fake.maxInstancesMutex.RUnlock()
	return len(fake.maxInstancesArgsForCall)
}

func (fake *FakeCloudRunCommandBuilder) MaxInstancesCalls(stub func(int) gcloud.CloudRunCommandBuilder) {
	fake.maxInstancesMutex.Lock()
	defer fake.maxInstancesMutex.Unlock()
	fake.MaxInstancesStub = stub
}

func (fake *FakeCloudRunCommandBuilder) MaxInstancesArgsForCall(i int) int {
	fake.maxInstancesMutex.RLock()
	defer fake.maxInstancesMutex.RUnlock()
	argsForCall := fake.maxInstancesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCloudRunCommandBuilder) MaxInstancesReturns(result1 gcloud.CloudRunCommandBuilder) {
	fake.maxInstancesMutex.Lock()
	defer fake.maxInstancesMutex.Unlock()
	fake.MaxInstancesStub = nil
	fake.maxInstancesReturns = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) MaxInstancesReturnsOnCall(i int, result1 gcloud.CloudRunCommandBuilder) {
	fake.maxInstancesMutex.Lock()
	defer fake.maxInstancesMutex.Unlock()
	fake.MaxInstancesStub = nil
	if fake.maxInstancesReturnsOnCall == nil {
		fake.maxInstancesReturnsOnCall = make(map[int]struct {
			result1 gcloud.CloudRunCommandBuilder
		})
	}
	fake.maxInstancesReturnsOnCall[i] = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) Memory(arg1 string) gcloud.CloudRunCommandBuilder {
	fake.memoryMutex.Lock()
	ret, specificReturn := fake.memoryReturnsOnCall[len(fake.memoryArgsForCall)]
	fake.memoryArgsForCall = append(fake.memoryArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.MemoryStub
	fakeReturns := fake.memoryReturns
	fake.recordInvocation("Memory", []interface{}{arg1})
	fake.memoryMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCloudRunCommandBuilder) MemoryCallCount() int {
	fake.memoryMutex.RLock()
	defer fake.memoryMutex.RUnlock()
	return len(fake.memoryArgsForCall)
}

func (fake *FakeCloudRunCommandBuilder) MemoryCalls(stub func(string) gcloud.CloudRunCommandBuilder) {
	fake.memoryMutex.Lock()
	defer fake.memoryMutex.Unlock()
	fake.MemoryStub = stub
}

func (fake *FakeCloudRunCommandBuilder) MemoryArgsForCall(i int) string {
	fake.memoryMutex.RLock()
	defer fake.memoryMutex.RUnlock()
	argsForCall := fake.memoryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCloudRunCommandBuilder) MemoryReturns(result1 gcloud.CloudRunCommandBuilder) {
	fake.memoryMutex.Lock()
	defer fake.memoryMutex.Unlock()
	fake.MemoryStub = nil
	fake.memoryReturns = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) MemoryReturnsOnCall(i int, result1 gcloud.CloudRunCommandBuilder) {
	fake.memoryMutex.Lock()
	defer fake.memoryMutex.Unlock()
	fake.MemoryStub = nil
	if fake.memoryReturnsOnCall == nil {
		fake.memoryReturnsOnCall = make(map[int]struct {
			result1 gcloud.CloudRunCommandBuilder
		})
	}
	fake.memoryReturnsOnCall[i] = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) ProjectID(arg1 string) gcloud.CloudRunCommandBuilder {
	fake.projectIDMutex.Lock()
	ret, specificReturn := fake.projectIDReturnsOnCall[len(fake.projectIDArgsForCall)]
	fake.projectIDArgsForCall = append(fake.projectIDArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ProjectIDStub
	fakeReturns := fake.projectIDReturns
	fake.recordInvocation("ProjectID", []interface{}{arg1})
	fake.projectIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCloudRunCommandBuilder) ProjectIDCallCount() int {
	fake.projectIDMutex.RLock()
	defer fake.projectIDMutex.RUnlock()
	return len(fake.projectIDArgsForCall)
}

func (fake *FakeCloudRunCommandBuilder) ProjectIDCalls(stub func(string) gcloud.CloudRunCommandBuilder) {
	fake.projectIDMutex.Lock()
	defer fake.projectIDMutex.Unlock()
	fake.ProjectIDStub = stub
}

func (fake *FakeCloudRunCommandBuilder) ProjectIDArgsForCall(i int) string {
	fake.projectIDMutex.RLock()
	defer fake.projectIDMutex.RUnlock()
	argsForCall := fake.projectIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCloudRunCommandBuilder) ProjectIDReturns(result1 gcloud.CloudRunCommandBuilder) {
	fake.projectIDMutex.Lock()
	defer fake.projectIDMutex.Unlock()
	fake.ProjectIDStub = nil
	fake.projectIDReturns = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) ProjectIDReturnsOnCall(i int, result1 gcloud.CloudRunCommandBuilder) {
	fake.projectIDMutex.Lock()
	defer fake.projectIDMutex.Unlock()
	fake.ProjectIDStub = nil
	if fake.projectIDReturnsOnCall == nil {
		fake.projectIDReturnsOnCall = make(map[int]struct {
			result1 gcloud.CloudRunCommandBuilder
		})
	}
	fake.projectIDReturnsOnCall[i] = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) Region(arg1 string) gcloud.CloudRunCommandBuilder {
	fake.regionMutex.Lock()
	ret, specificReturn := fake.regionReturnsOnCall[len(fake.regionArgsForCall)]
	fake.regionArgsForCall = append(fake.regionArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RegionStub
	fakeReturns := fake.regionReturns
	fake.recordInvocation("Region", []interface{}{arg1})
	fake.regionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCloudRunCommandBuilder) RegionCallCount() int {
	fake.regionMutex.RLock()
	defer fake.regionMutex.RUnlock()
	return len(fake.regionArgsForCall)
}

func (fake *FakeCloudRunCommandBuilder) RegionCalls(stub func(string) gcloud.CloudRunCommandBuilder) {
	fake.regionMutex.Lock()
	defer fake.regionMutex.Unlock()
	fake.RegionStub = stub
}

func (fake *FakeCloudRunCommandBuilder) RegionArgsForCall(i int) string {
	fake.regionMutex.RLock()
	defer fake.regionMutex.RUnlock()
	argsForCall := fake.regionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCloudRunCommandBuilder) RegionReturns(result1 gcloud.CloudRunCommandBuilder) {
	fake.regionMutex.Lock()
	defer fake.regionMutex.Unlock()
	fake.RegionStub = nil
	fake.regionReturns = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) RegionReturnsOnCall(i int, result1 gcloud.CloudRunCommandBuilder) {
	fake.regionMutex.Lock()
	defer fake.regionMutex.Unlock()
	fake.RegionStub = nil
	if fake.regionReturnsOnCall == nil {
		fake.regionReturnsOnCall = make(map[int]struct {
			result1 gcloud.CloudRunCommandBuilder
		})
	}
	fake.regionReturnsOnCall[i] = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) ServiceName(arg1 string) gcloud.CloudRunCommandBuilder {
	fake.serviceNameMutex.Lock()
	ret, specificReturn := fake.serviceNameReturnsOnCall[len(fake.serviceNameArgsForCall)]
	fake.serviceNameArgsForCall = append(fake.serviceNameArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ServiceNameStub
	fakeReturns := fake.serviceNameReturns
	fake.recordInvocation("ServiceName", []interface{}{arg1})
	fake.serviceNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCloudRunCommandBuilder) ServiceNameCallCount() int {
	fake.serviceNameMutex.RLock()
	defer fake.serviceNameMutex.RUnlock()
	return len(fake.serviceNameArgsForCall)
}

func (fake *FakeCloudRunCommandBuilder) ServiceNameCalls(stub func(string) gcloud.CloudRunCommandBuilder) {
	fake.serviceNameMutex.Lock()
	defer fake.serviceNameMutex.Unlock()
	fake.ServiceNameStub = stub
}

func (fake *FakeCloudRunCommandBuilder) ServiceNameArgsForCall(i int) string {
	fake.serviceNameMutex.RLock()
	defer fake.serviceNameMutex.RUnlock()
	argsForCall := fake.serviceNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCloudRunCommandBuilder) ServiceNameReturns(result1 gcloud.CloudRunCommandBuilder) {
	fake.serviceNameMutex.Lock()
	defer fake.serviceNameMutex.Unlock()
	fake.ServiceNameStub = nil
	fake.serviceNameReturns = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) ServiceNameReturnsOnCall(i int, result1 gcloud.CloudRunCommandBuilder) {
	fake.serviceNameMutex.Lock()
	defer fake.serviceNameMutex.Unlock()
	fake.ServiceNameStub = nil
	if fake.serviceNameReturnsOnCall == nil {
		fake.serviceNameReturnsOnCall = make(map[int]struct {
			result1 gcloud.CloudRunCommandBuilder
		})
	}
	fake.serviceNameReturnsOnCall[i] = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) VPCConnector(arg1 string) gcloud.CloudRunCommandBuilder {
	fake.vPCConnectorMutex.Lock()
	ret, specificReturn := fake.vPCConnectorReturnsOnCall[len(fake.vPCConnectorArgsForCall)]
	fake.vPCConnectorArgsForCall = append(fake.vPCConnectorArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.VPCConnectorStub
	fakeReturns := fake.vPCConnectorReturns
	fake.recordInvocation("VPCConnector", []interface{}{arg1})
	fake.vPCConnectorMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCloudRunCommandBuilder) VPCConnectorCallCount() int {
	fake.vPCConnectorMutex.RLock()
	defer fake.vPCConnectorMutex.RUnlock()
	return len(fake.vPCConnectorArgsForCall)
}

func (fake *FakeCloudRunCommandBuilder) VPCConnectorCalls(stub func(string) gcloud.CloudRunCommandBuilder) {
	fake.vPCConnectorMutex.Lock()
	defer fake.vPCConnectorMutex.Unlock()
	fake.VPCConnectorStub = stub
}

func (fake *FakeCloudRunCommandBuilder) VPCConnectorArgsForCall(i int) string {
	fake.vPCConnectorMutex.RLock()
	defer fake.vPCConnectorMutex.RUnlock()
	argsForCall := fake.vPCConnectorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCloudRunCommandBuilder) VPCConnectorReturns(result1 gcloud.CloudRunCommandBuilder) {
	fake.vPCConnectorMutex.Lock()
	defer fake.vPCConnectorMutex.Unlock()
	fake.VPCConnectorStub = nil
	fake.vPCConnectorReturns = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) VPCConnectorReturnsOnCall(i int, result1 gcloud.CloudRunCommandBuilder) {
	fake.vPCConnectorMutex.Lock()
	defer fake.vPCConnectorMutex.Unlock()
	fake.VPCConnectorStub = nil
	if fake.vPCConnectorReturnsOnCall == nil {
		fake.vPCConnectorReturnsOnCall = make(map[int]struct {
			result1 gcloud.CloudRunCommandBuilder
		})
	}
	fake.vPCConnectorReturnsOnCall[i] = struct {
		result1 gcloud.CloudRunCommandBuilder
	}{result1}
}

func (fake *FakeCloudRunCommandBuilder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allowUnauthenticatedMutex.RLock()
	defer fake.allowUnauthenticatedMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	fake.maxInstancesMutex.RLock()
	defer fake.maxInstancesMutex.RUnlock()
	fake.memoryMutex.RLock()
	defer fake.memoryMutex.RUnlock()
	fake.projectIDMutex.RLock()
	defer fake.projectIDMutex.RUnlock()
	fake.regionMutex.RLock()
	defer fake.regionMutex.RUnlock()
	fake.serviceNameMutex.RLock()
	defer fake.serviceNameMutex.RUnlock()
	fake.vPCConnectorMutex.RLock()
	defer fake.vPCConnectorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCloudRunCommandBuilder) recordInvocation(key string, args []interface{}) {
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

var _ gcloud.CloudRunCommandBuilder = new(FakeCloudRunCommandBuilder)
