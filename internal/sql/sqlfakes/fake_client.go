// Code generated by counterfeiter. DO NOT EDIT.
package sqlfakes

import (
	"sync"

	"github.com/homedepot/cloud-runner/internal/sql"
	cloudrunner "github.com/homedepot/cloud-runner/pkg"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type FakeClient struct {
	ConnectStub        func(string, interface{}) error
	connectMutex       sync.RWMutex
	connectArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	connectReturns struct {
		result1 error
	}
	connectReturnsOnCall map[int]struct {
		result1 error
	}
	ConnectionStub        func() (string, string)
	connectionMutex       sync.RWMutex
	connectionArgsForCall []struct{}
	connectionReturns     struct {
		result1 string
		result2 string
	}
	connectionReturnsOnCall map[int]struct {
		result1 string
		result2 string
	}
	CreateCredentialsStub        func(cloudrunner.Credentials) error
	createCredentialsMutex       sync.RWMutex
	createCredentialsArgsForCall []struct {
		arg1 cloudrunner.Credentials
	}
	createCredentialsReturns struct {
		result1 error
	}
	createCredentialsReturnsOnCall map[int]struct {
		result1 error
	}
	CreateDeploymentStub        func(cloudrunner.Deployment) error
	createDeploymentMutex       sync.RWMutex
	createDeploymentArgsForCall []struct {
		arg1 cloudrunner.Deployment
	}
	createDeploymentReturns struct {
		result1 error
	}
	createDeploymentReturnsOnCall map[int]struct {
		result1 error
	}
	CreateReadPermissionStub        func(cloudrunner.CredentialsReadPermission) error
	createReadPermissionMutex       sync.RWMutex
	createReadPermissionArgsForCall []struct {
		arg1 cloudrunner.CredentialsReadPermission
	}
	createReadPermissionReturns struct {
		result1 error
	}
	createReadPermissionReturnsOnCall map[int]struct {
		result1 error
	}
	CreateWritePermissionStub        func(cloudrunner.CredentialsWritePermission) error
	createWritePermissionMutex       sync.RWMutex
	createWritePermissionArgsForCall []struct {
		arg1 cloudrunner.CredentialsWritePermission
	}
	createWritePermissionReturns struct {
		result1 error
	}
	createWritePermissionReturnsOnCall map[int]struct {
		result1 error
	}
	DBStub        func() *gorm.DB
	dBMutex       sync.RWMutex
	dBArgsForCall []struct{}
	dBReturns     struct {
		result1 *gorm.DB
	}
	dBReturnsOnCall map[int]struct {
		result1 *gorm.DB
	}
	DeleteCredentialsStub        func(string) error
	deleteCredentialsMutex       sync.RWMutex
	deleteCredentialsArgsForCall []struct {
		arg1 string
	}
	deleteCredentialsReturns struct {
		result1 error
	}
	deleteCredentialsReturnsOnCall map[int]struct {
		result1 error
	}
	GetCredentialsStub        func(string) (cloudrunner.Credentials, error)
	getCredentialsMutex       sync.RWMutex
	getCredentialsArgsForCall []struct {
		arg1 string
	}
	getCredentialsReturns struct {
		result1 cloudrunner.Credentials
		result2 error
	}
	getCredentialsReturnsOnCall map[int]struct {
		result1 cloudrunner.Credentials
		result2 error
	}
	GetDeploymentStub        func(string) (cloudrunner.Deployment, error)
	getDeploymentMutex       sync.RWMutex
	getDeploymentArgsForCall []struct {
		arg1 string
	}
	getDeploymentReturns struct {
		result1 cloudrunner.Deployment
		result2 error
	}
	getDeploymentReturnsOnCall map[int]struct {
		result1 cloudrunner.Deployment
		result2 error
	}
	ListCredentialsStub        func() ([]cloudrunner.Credentials, error)
	listCredentialsMutex       sync.RWMutex
	listCredentialsArgsForCall []struct{}
	listCredentialsReturns     struct {
		result1 []cloudrunner.Credentials
		result2 error
	}
	listCredentialsReturnsOnCall map[int]struct {
		result1 []cloudrunner.Credentials
		result2 error
	}
	UpdateDeploymentStub        func(cloudrunner.Deployment) error
	updateDeploymentMutex       sync.RWMutex
	updateDeploymentArgsForCall []struct {
		arg1 cloudrunner.Deployment
	}
	updateDeploymentReturns struct {
		result1 error
	}
	updateDeploymentReturnsOnCall map[int]struct {
		result1 error
	}
	WithHostStub        func(string)
	withHostMutex       sync.RWMutex
	withHostArgsForCall []struct {
		arg1 string
	}
	WithNameStub        func(string)
	withNameMutex       sync.RWMutex
	withNameArgsForCall []struct {
		arg1 string
	}
	WithPassStub        func(string)
	withPassMutex       sync.RWMutex
	withPassArgsForCall []struct {
		arg1 string
	}
	WithUserStub        func(string)
	withUserMutex       sync.RWMutex
	withUserArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Connect(arg1 string, arg2 interface{}) error {
	fake.connectMutex.Lock()
	ret, specificReturn := fake.connectReturnsOnCall[len(fake.connectArgsForCall)]
	fake.connectArgsForCall = append(fake.connectArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("Connect", []interface{}{arg1, arg2})
	fake.connectMutex.Unlock()
	if fake.ConnectStub != nil {
		return fake.ConnectStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.connectReturns.result1
}

func (fake *FakeClient) ConnectCallCount() int {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	return len(fake.connectArgsForCall)
}

func (fake *FakeClient) ConnectArgsForCall(i int) (string, interface{}) {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	return fake.connectArgsForCall[i].arg1, fake.connectArgsForCall[i].arg2
}

func (fake *FakeClient) ConnectReturns(result1 error) {
	fake.ConnectStub = nil
	fake.connectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) ConnectReturnsOnCall(i int, result1 error) {
	fake.ConnectStub = nil
	if fake.connectReturnsOnCall == nil {
		fake.connectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.connectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Connection() (string, string) {
	fake.connectionMutex.Lock()
	ret, specificReturn := fake.connectionReturnsOnCall[len(fake.connectionArgsForCall)]
	fake.connectionArgsForCall = append(fake.connectionArgsForCall, struct{}{})
	fake.recordInvocation("Connection", []interface{}{})
	fake.connectionMutex.Unlock()
	if fake.ConnectionStub != nil {
		return fake.ConnectionStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.connectionReturns.result1, fake.connectionReturns.result2
}

func (fake *FakeClient) ConnectionCallCount() int {
	fake.connectionMutex.RLock()
	defer fake.connectionMutex.RUnlock()
	return len(fake.connectionArgsForCall)
}

func (fake *FakeClient) ConnectionReturns(result1 string, result2 string) {
	fake.ConnectionStub = nil
	fake.connectionReturns = struct {
		result1 string
		result2 string
	}{result1, result2}
}

func (fake *FakeClient) ConnectionReturnsOnCall(i int, result1 string, result2 string) {
	fake.ConnectionStub = nil
	if fake.connectionReturnsOnCall == nil {
		fake.connectionReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
		})
	}
	fake.connectionReturnsOnCall[i] = struct {
		result1 string
		result2 string
	}{result1, result2}
}

func (fake *FakeClient) CreateCredentials(arg1 cloudrunner.Credentials) error {
	fake.createCredentialsMutex.Lock()
	ret, specificReturn := fake.createCredentialsReturnsOnCall[len(fake.createCredentialsArgsForCall)]
	fake.createCredentialsArgsForCall = append(fake.createCredentialsArgsForCall, struct {
		arg1 cloudrunner.Credentials
	}{arg1})
	fake.recordInvocation("CreateCredentials", []interface{}{arg1})
	fake.createCredentialsMutex.Unlock()
	if fake.CreateCredentialsStub != nil {
		return fake.CreateCredentialsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createCredentialsReturns.result1
}

func (fake *FakeClient) CreateCredentialsCallCount() int {
	fake.createCredentialsMutex.RLock()
	defer fake.createCredentialsMutex.RUnlock()
	return len(fake.createCredentialsArgsForCall)
}

func (fake *FakeClient) CreateCredentialsArgsForCall(i int) cloudrunner.Credentials {
	fake.createCredentialsMutex.RLock()
	defer fake.createCredentialsMutex.RUnlock()
	return fake.createCredentialsArgsForCall[i].arg1
}

func (fake *FakeClient) CreateCredentialsReturns(result1 error) {
	fake.CreateCredentialsStub = nil
	fake.createCredentialsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateCredentialsReturnsOnCall(i int, result1 error) {
	fake.CreateCredentialsStub = nil
	if fake.createCredentialsReturnsOnCall == nil {
		fake.createCredentialsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createCredentialsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateDeployment(arg1 cloudrunner.Deployment) error {
	fake.createDeploymentMutex.Lock()
	ret, specificReturn := fake.createDeploymentReturnsOnCall[len(fake.createDeploymentArgsForCall)]
	fake.createDeploymentArgsForCall = append(fake.createDeploymentArgsForCall, struct {
		arg1 cloudrunner.Deployment
	}{arg1})
	fake.recordInvocation("CreateDeployment", []interface{}{arg1})
	fake.createDeploymentMutex.Unlock()
	if fake.CreateDeploymentStub != nil {
		return fake.CreateDeploymentStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createDeploymentReturns.result1
}

func (fake *FakeClient) CreateDeploymentCallCount() int {
	fake.createDeploymentMutex.RLock()
	defer fake.createDeploymentMutex.RUnlock()
	return len(fake.createDeploymentArgsForCall)
}

func (fake *FakeClient) CreateDeploymentArgsForCall(i int) cloudrunner.Deployment {
	fake.createDeploymentMutex.RLock()
	defer fake.createDeploymentMutex.RUnlock()
	return fake.createDeploymentArgsForCall[i].arg1
}

func (fake *FakeClient) CreateDeploymentReturns(result1 error) {
	fake.CreateDeploymentStub = nil
	fake.createDeploymentReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateDeploymentReturnsOnCall(i int, result1 error) {
	fake.CreateDeploymentStub = nil
	if fake.createDeploymentReturnsOnCall == nil {
		fake.createDeploymentReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createDeploymentReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateReadPermission(arg1 cloudrunner.CredentialsReadPermission) error {
	fake.createReadPermissionMutex.Lock()
	ret, specificReturn := fake.createReadPermissionReturnsOnCall[len(fake.createReadPermissionArgsForCall)]
	fake.createReadPermissionArgsForCall = append(fake.createReadPermissionArgsForCall, struct {
		arg1 cloudrunner.CredentialsReadPermission
	}{arg1})
	fake.recordInvocation("CreateReadPermission", []interface{}{arg1})
	fake.createReadPermissionMutex.Unlock()
	if fake.CreateReadPermissionStub != nil {
		return fake.CreateReadPermissionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReadPermissionReturns.result1
}

func (fake *FakeClient) CreateReadPermissionCallCount() int {
	fake.createReadPermissionMutex.RLock()
	defer fake.createReadPermissionMutex.RUnlock()
	return len(fake.createReadPermissionArgsForCall)
}

func (fake *FakeClient) CreateReadPermissionArgsForCall(i int) cloudrunner.CredentialsReadPermission {
	fake.createReadPermissionMutex.RLock()
	defer fake.createReadPermissionMutex.RUnlock()
	return fake.createReadPermissionArgsForCall[i].arg1
}

func (fake *FakeClient) CreateReadPermissionReturns(result1 error) {
	fake.CreateReadPermissionStub = nil
	fake.createReadPermissionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateReadPermissionReturnsOnCall(i int, result1 error) {
	fake.CreateReadPermissionStub = nil
	if fake.createReadPermissionReturnsOnCall == nil {
		fake.createReadPermissionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReadPermissionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateWritePermission(arg1 cloudrunner.CredentialsWritePermission) error {
	fake.createWritePermissionMutex.Lock()
	ret, specificReturn := fake.createWritePermissionReturnsOnCall[len(fake.createWritePermissionArgsForCall)]
	fake.createWritePermissionArgsForCall = append(fake.createWritePermissionArgsForCall, struct {
		arg1 cloudrunner.CredentialsWritePermission
	}{arg1})
	fake.recordInvocation("CreateWritePermission", []interface{}{arg1})
	fake.createWritePermissionMutex.Unlock()
	if fake.CreateWritePermissionStub != nil {
		return fake.CreateWritePermissionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createWritePermissionReturns.result1
}

func (fake *FakeClient) CreateWritePermissionCallCount() int {
	fake.createWritePermissionMutex.RLock()
	defer fake.createWritePermissionMutex.RUnlock()
	return len(fake.createWritePermissionArgsForCall)
}

func (fake *FakeClient) CreateWritePermissionArgsForCall(i int) cloudrunner.CredentialsWritePermission {
	fake.createWritePermissionMutex.RLock()
	defer fake.createWritePermissionMutex.RUnlock()
	return fake.createWritePermissionArgsForCall[i].arg1
}

func (fake *FakeClient) CreateWritePermissionReturns(result1 error) {
	fake.CreateWritePermissionStub = nil
	fake.createWritePermissionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateWritePermissionReturnsOnCall(i int, result1 error) {
	fake.CreateWritePermissionStub = nil
	if fake.createWritePermissionReturnsOnCall == nil {
		fake.createWritePermissionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createWritePermissionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DB() *gorm.DB {
	fake.dBMutex.Lock()
	ret, specificReturn := fake.dBReturnsOnCall[len(fake.dBArgsForCall)]
	fake.dBArgsForCall = append(fake.dBArgsForCall, struct{}{})
	fake.recordInvocation("DB", []interface{}{})
	fake.dBMutex.Unlock()
	if fake.DBStub != nil {
		return fake.DBStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.dBReturns.result1
}

func (fake *FakeClient) DBCallCount() int {
	fake.dBMutex.RLock()
	defer fake.dBMutex.RUnlock()
	return len(fake.dBArgsForCall)
}

func (fake *FakeClient) DBReturns(result1 *gorm.DB) {
	fake.DBStub = nil
	fake.dBReturns = struct {
		result1 *gorm.DB
	}{result1}
}

func (fake *FakeClient) DBReturnsOnCall(i int, result1 *gorm.DB) {
	fake.DBStub = nil
	if fake.dBReturnsOnCall == nil {
		fake.dBReturnsOnCall = make(map[int]struct {
			result1 *gorm.DB
		})
	}
	fake.dBReturnsOnCall[i] = struct {
		result1 *gorm.DB
	}{result1}
}

func (fake *FakeClient) DeleteCredentials(arg1 string) error {
	fake.deleteCredentialsMutex.Lock()
	ret, specificReturn := fake.deleteCredentialsReturnsOnCall[len(fake.deleteCredentialsArgsForCall)]
	fake.deleteCredentialsArgsForCall = append(fake.deleteCredentialsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteCredentials", []interface{}{arg1})
	fake.deleteCredentialsMutex.Unlock()
	if fake.DeleteCredentialsStub != nil {
		return fake.DeleteCredentialsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteCredentialsReturns.result1
}

func (fake *FakeClient) DeleteCredentialsCallCount() int {
	fake.deleteCredentialsMutex.RLock()
	defer fake.deleteCredentialsMutex.RUnlock()
	return len(fake.deleteCredentialsArgsForCall)
}

func (fake *FakeClient) DeleteCredentialsArgsForCall(i int) string {
	fake.deleteCredentialsMutex.RLock()
	defer fake.deleteCredentialsMutex.RUnlock()
	return fake.deleteCredentialsArgsForCall[i].arg1
}

func (fake *FakeClient) DeleteCredentialsReturns(result1 error) {
	fake.DeleteCredentialsStub = nil
	fake.deleteCredentialsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DeleteCredentialsReturnsOnCall(i int, result1 error) {
	fake.DeleteCredentialsStub = nil
	if fake.deleteCredentialsReturnsOnCall == nil {
		fake.deleteCredentialsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteCredentialsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) GetCredentials(arg1 string) (cloudrunner.Credentials, error) {
	fake.getCredentialsMutex.Lock()
	ret, specificReturn := fake.getCredentialsReturnsOnCall[len(fake.getCredentialsArgsForCall)]
	fake.getCredentialsArgsForCall = append(fake.getCredentialsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetCredentials", []interface{}{arg1})
	fake.getCredentialsMutex.Unlock()
	if fake.GetCredentialsStub != nil {
		return fake.GetCredentialsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getCredentialsReturns.result1, fake.getCredentialsReturns.result2
}

func (fake *FakeClient) GetCredentialsCallCount() int {
	fake.getCredentialsMutex.RLock()
	defer fake.getCredentialsMutex.RUnlock()
	return len(fake.getCredentialsArgsForCall)
}

func (fake *FakeClient) GetCredentialsArgsForCall(i int) string {
	fake.getCredentialsMutex.RLock()
	defer fake.getCredentialsMutex.RUnlock()
	return fake.getCredentialsArgsForCall[i].arg1
}

func (fake *FakeClient) GetCredentialsReturns(result1 cloudrunner.Credentials, result2 error) {
	fake.GetCredentialsStub = nil
	fake.getCredentialsReturns = struct {
		result1 cloudrunner.Credentials
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetCredentialsReturnsOnCall(i int, result1 cloudrunner.Credentials, result2 error) {
	fake.GetCredentialsStub = nil
	if fake.getCredentialsReturnsOnCall == nil {
		fake.getCredentialsReturnsOnCall = make(map[int]struct {
			result1 cloudrunner.Credentials
			result2 error
		})
	}
	fake.getCredentialsReturnsOnCall[i] = struct {
		result1 cloudrunner.Credentials
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetDeployment(arg1 string) (cloudrunner.Deployment, error) {
	fake.getDeploymentMutex.Lock()
	ret, specificReturn := fake.getDeploymentReturnsOnCall[len(fake.getDeploymentArgsForCall)]
	fake.getDeploymentArgsForCall = append(fake.getDeploymentArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetDeployment", []interface{}{arg1})
	fake.getDeploymentMutex.Unlock()
	if fake.GetDeploymentStub != nil {
		return fake.GetDeploymentStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDeploymentReturns.result1, fake.getDeploymentReturns.result2
}

func (fake *FakeClient) GetDeploymentCallCount() int {
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	return len(fake.getDeploymentArgsForCall)
}

func (fake *FakeClient) GetDeploymentArgsForCall(i int) string {
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	return fake.getDeploymentArgsForCall[i].arg1
}

func (fake *FakeClient) GetDeploymentReturns(result1 cloudrunner.Deployment, result2 error) {
	fake.GetDeploymentStub = nil
	fake.getDeploymentReturns = struct {
		result1 cloudrunner.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetDeploymentReturnsOnCall(i int, result1 cloudrunner.Deployment, result2 error) {
	fake.GetDeploymentStub = nil
	if fake.getDeploymentReturnsOnCall == nil {
		fake.getDeploymentReturnsOnCall = make(map[int]struct {
			result1 cloudrunner.Deployment
			result2 error
		})
	}
	fake.getDeploymentReturnsOnCall[i] = struct {
		result1 cloudrunner.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListCredentials() ([]cloudrunner.Credentials, error) {
	fake.listCredentialsMutex.Lock()
	ret, specificReturn := fake.listCredentialsReturnsOnCall[len(fake.listCredentialsArgsForCall)]
	fake.listCredentialsArgsForCall = append(fake.listCredentialsArgsForCall, struct{}{})
	fake.recordInvocation("ListCredentials", []interface{}{})
	fake.listCredentialsMutex.Unlock()
	if fake.ListCredentialsStub != nil {
		return fake.ListCredentialsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listCredentialsReturns.result1, fake.listCredentialsReturns.result2
}

func (fake *FakeClient) ListCredentialsCallCount() int {
	fake.listCredentialsMutex.RLock()
	defer fake.listCredentialsMutex.RUnlock()
	return len(fake.listCredentialsArgsForCall)
}

func (fake *FakeClient) ListCredentialsReturns(result1 []cloudrunner.Credentials, result2 error) {
	fake.ListCredentialsStub = nil
	fake.listCredentialsReturns = struct {
		result1 []cloudrunner.Credentials
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListCredentialsReturnsOnCall(i int, result1 []cloudrunner.Credentials, result2 error) {
	fake.ListCredentialsStub = nil
	if fake.listCredentialsReturnsOnCall == nil {
		fake.listCredentialsReturnsOnCall = make(map[int]struct {
			result1 []cloudrunner.Credentials
			result2 error
		})
	}
	fake.listCredentialsReturnsOnCall[i] = struct {
		result1 []cloudrunner.Credentials
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) UpdateDeployment(arg1 cloudrunner.Deployment) error {
	fake.updateDeploymentMutex.Lock()
	ret, specificReturn := fake.updateDeploymentReturnsOnCall[len(fake.updateDeploymentArgsForCall)]
	fake.updateDeploymentArgsForCall = append(fake.updateDeploymentArgsForCall, struct {
		arg1 cloudrunner.Deployment
	}{arg1})
	fake.recordInvocation("UpdateDeployment", []interface{}{arg1})
	fake.updateDeploymentMutex.Unlock()
	if fake.UpdateDeploymentStub != nil {
		return fake.UpdateDeploymentStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateDeploymentReturns.result1
}

func (fake *FakeClient) UpdateDeploymentCallCount() int {
	fake.updateDeploymentMutex.RLock()
	defer fake.updateDeploymentMutex.RUnlock()
	return len(fake.updateDeploymentArgsForCall)
}

func (fake *FakeClient) UpdateDeploymentArgsForCall(i int) cloudrunner.Deployment {
	fake.updateDeploymentMutex.RLock()
	defer fake.updateDeploymentMutex.RUnlock()
	return fake.updateDeploymentArgsForCall[i].arg1
}

func (fake *FakeClient) UpdateDeploymentReturns(result1 error) {
	fake.UpdateDeploymentStub = nil
	fake.updateDeploymentReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) UpdateDeploymentReturnsOnCall(i int, result1 error) {
	fake.UpdateDeploymentStub = nil
	if fake.updateDeploymentReturnsOnCall == nil {
		fake.updateDeploymentReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateDeploymentReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) WithHost(arg1 string) {
	fake.withHostMutex.Lock()
	fake.withHostArgsForCall = append(fake.withHostArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("WithHost", []interface{}{arg1})
	fake.withHostMutex.Unlock()
	if fake.WithHostStub != nil {
		fake.WithHostStub(arg1)
	}
}

func (fake *FakeClient) WithHostCallCount() int {
	fake.withHostMutex.RLock()
	defer fake.withHostMutex.RUnlock()
	return len(fake.withHostArgsForCall)
}

func (fake *FakeClient) WithHostArgsForCall(i int) string {
	fake.withHostMutex.RLock()
	defer fake.withHostMutex.RUnlock()
	return fake.withHostArgsForCall[i].arg1
}

func (fake *FakeClient) WithName(arg1 string) {
	fake.withNameMutex.Lock()
	fake.withNameArgsForCall = append(fake.withNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("WithName", []interface{}{arg1})
	fake.withNameMutex.Unlock()
	if fake.WithNameStub != nil {
		fake.WithNameStub(arg1)
	}
}

func (fake *FakeClient) WithNameCallCount() int {
	fake.withNameMutex.RLock()
	defer fake.withNameMutex.RUnlock()
	return len(fake.withNameArgsForCall)
}

func (fake *FakeClient) WithNameArgsForCall(i int) string {
	fake.withNameMutex.RLock()
	defer fake.withNameMutex.RUnlock()
	return fake.withNameArgsForCall[i].arg1
}

func (fake *FakeClient) WithPass(arg1 string) {
	fake.withPassMutex.Lock()
	fake.withPassArgsForCall = append(fake.withPassArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("WithPass", []interface{}{arg1})
	fake.withPassMutex.Unlock()
	if fake.WithPassStub != nil {
		fake.WithPassStub(arg1)
	}
}

func (fake *FakeClient) WithPassCallCount() int {
	fake.withPassMutex.RLock()
	defer fake.withPassMutex.RUnlock()
	return len(fake.withPassArgsForCall)
}

func (fake *FakeClient) WithPassArgsForCall(i int) string {
	fake.withPassMutex.RLock()
	defer fake.withPassMutex.RUnlock()
	return fake.withPassArgsForCall[i].arg1
}

func (fake *FakeClient) WithUser(arg1 string) {
	fake.withUserMutex.Lock()
	fake.withUserArgsForCall = append(fake.withUserArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("WithUser", []interface{}{arg1})
	fake.withUserMutex.Unlock()
	if fake.WithUserStub != nil {
		fake.WithUserStub(arg1)
	}
}

func (fake *FakeClient) WithUserCallCount() int {
	fake.withUserMutex.RLock()
	defer fake.withUserMutex.RUnlock()
	return len(fake.withUserArgsForCall)
}

func (fake *FakeClient) WithUserArgsForCall(i int) string {
	fake.withUserMutex.RLock()
	defer fake.withUserMutex.RUnlock()
	return fake.withUserArgsForCall[i].arg1
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	fake.connectionMutex.RLock()
	defer fake.connectionMutex.RUnlock()
	fake.createCredentialsMutex.RLock()
	defer fake.createCredentialsMutex.RUnlock()
	fake.createDeploymentMutex.RLock()
	defer fake.createDeploymentMutex.RUnlock()
	fake.createReadPermissionMutex.RLock()
	defer fake.createReadPermissionMutex.RUnlock()
	fake.createWritePermissionMutex.RLock()
	defer fake.createWritePermissionMutex.RUnlock()
	fake.dBMutex.RLock()
	defer fake.dBMutex.RUnlock()
	fake.deleteCredentialsMutex.RLock()
	defer fake.deleteCredentialsMutex.RUnlock()
	fake.getCredentialsMutex.RLock()
	defer fake.getCredentialsMutex.RUnlock()
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	fake.listCredentialsMutex.RLock()
	defer fake.listCredentialsMutex.RUnlock()
	fake.updateDeploymentMutex.RLock()
	defer fake.updateDeploymentMutex.RUnlock()
	fake.withHostMutex.RLock()
	defer fake.withHostMutex.RUnlock()
	fake.withNameMutex.RLock()
	defer fake.withNameMutex.RUnlock()
	fake.withPassMutex.RLock()
	defer fake.withPassMutex.RUnlock()
	fake.withUserMutex.RLock()
	defer fake.withUserMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ sql.Client = new(FakeClient)
