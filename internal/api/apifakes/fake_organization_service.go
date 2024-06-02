// Code generated by counterfeiter. DO NOT EDIT.
package apifakes

import (
	"context"
	"sync"

	"github.com/dembygenesis/local.tools/internal/model"
)

type FakeOrganizationService struct {
	CreateOrganizationStub        func(context.Context, *model.CreateOrganization) (*model.Organization, error)
	createOrganizationMutex       sync.RWMutex
	createOrganizationArgsForCall []struct {
		arg1 context.Context
		arg2 *model.CreateOrganization
	}
	createOrganizationReturns struct {
		result1 *model.Organization
		result2 error
	}
	createOrganizationReturnsOnCall map[int]struct {
		result1 *model.Organization
		result2 error
	}
	DeleteOrganizationStub        func(context.Context, *model.DeleteOrganization) error
	deleteOrganizationMutex       sync.RWMutex
	deleteOrganizationArgsForCall []struct {
		arg1 context.Context
		arg2 *model.DeleteOrganization
	}
	deleteOrganizationReturns struct {
		result1 error
	}
	deleteOrganizationReturnsOnCall map[int]struct {
		result1 error
	}
	ListOrganizationsStub        func(context.Context, *model.OrganizationFilters) (*model.PaginatedOrganizations, error)
	listOrganizationsMutex       sync.RWMutex
	listOrganizationsArgsForCall []struct {
		arg1 context.Context
		arg2 *model.OrganizationFilters
	}
	listOrganizationsReturns struct {
		result1 *model.PaginatedOrganizations
		result2 error
	}
	listOrganizationsReturnsOnCall map[int]struct {
		result1 *model.PaginatedOrganizations
		result2 error
	}
	RestoreOrganizationStub        func(context.Context, *model.RestoreOrganization) error
	restoreOrganizationMutex       sync.RWMutex
	restoreOrganizationArgsForCall []struct {
		arg1 context.Context
		arg2 *model.RestoreOrganization
	}
	restoreOrganizationReturns struct {
		result1 error
	}
	restoreOrganizationReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOrganizationService) UpdateOrganization(ctx context.Context, category *model.UpdateOrganization) (*model.Organization, error) {
	//TODO implement me
	panic("implement me")
}

func (fake *FakeOrganizationService) CreateOrganization(arg1 context.Context, arg2 *model.CreateOrganization) (*model.Organization, error) {
	fake.createOrganizationMutex.Lock()
	ret, specificReturn := fake.createOrganizationReturnsOnCall[len(fake.createOrganizationArgsForCall)]
	fake.createOrganizationArgsForCall = append(fake.createOrganizationArgsForCall, struct {
		arg1 context.Context
		arg2 *model.CreateOrganization
	}{arg1, arg2})
	stub := fake.CreateOrganizationStub
	fakeReturns := fake.createOrganizationReturns
	fake.recordInvocation("CreateOrganization", []interface{}{arg1, arg2})
	fake.createOrganizationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOrganizationService) CreateOrganizationCallCount() int {
	fake.createOrganizationMutex.RLock()
	defer fake.createOrganizationMutex.RUnlock()
	return len(fake.createOrganizationArgsForCall)
}

func (fake *FakeOrganizationService) CreateOrganizationCalls(stub func(context.Context, *model.CreateOrganization) (*model.Organization, error)) {
	fake.createOrganizationMutex.Lock()
	defer fake.createOrganizationMutex.Unlock()
	fake.CreateOrganizationStub = stub
}

func (fake *FakeOrganizationService) CreateOrganizationArgsForCall(i int) (context.Context, *model.CreateOrganization) {
	fake.createOrganizationMutex.RLock()
	defer fake.createOrganizationMutex.RUnlock()
	argsForCall := fake.createOrganizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOrganizationService) CreateOrganizationReturns(result1 *model.Organization, result2 error) {
	fake.createOrganizationMutex.Lock()
	defer fake.createOrganizationMutex.Unlock()
	fake.CreateOrganizationStub = nil
	fake.createOrganizationReturns = struct {
		result1 *model.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeOrganizationService) CreateOrganizationReturnsOnCall(i int, result1 *model.Organization, result2 error) {
	fake.createOrganizationMutex.Lock()
	defer fake.createOrganizationMutex.Unlock()
	fake.CreateOrganizationStub = nil
	if fake.createOrganizationReturnsOnCall == nil {
		fake.createOrganizationReturnsOnCall = make(map[int]struct {
			result1 *model.Organization
			result2 error
		})
	}
	fake.createOrganizationReturnsOnCall[i] = struct {
		result1 *model.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeOrganizationService) DeleteOrganization(arg1 context.Context, arg2 *model.DeleteOrganization) error {
	fake.deleteOrganizationMutex.Lock()
	ret, specificReturn := fake.deleteOrganizationReturnsOnCall[len(fake.deleteOrganizationArgsForCall)]
	fake.deleteOrganizationArgsForCall = append(fake.deleteOrganizationArgsForCall, struct {
		arg1 context.Context
		arg2 *model.DeleteOrganization
	}{arg1, arg2})
	stub := fake.DeleteOrganizationStub
	fakeReturns := fake.deleteOrganizationReturns
	fake.recordInvocation("DeleteOrganization", []interface{}{arg1, arg2})
	fake.deleteOrganizationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOrganizationService) DeleteOrganizationCallCount() int {
	fake.deleteOrganizationMutex.RLock()
	defer fake.deleteOrganizationMutex.RUnlock()
	return len(fake.deleteOrganizationArgsForCall)
}

func (fake *FakeOrganizationService) DeleteOrganizationCalls(stub func(context.Context, *model.DeleteOrganization) error) {
	fake.deleteOrganizationMutex.Lock()
	defer fake.deleteOrganizationMutex.Unlock()
	fake.DeleteOrganizationStub = stub
}

func (fake *FakeOrganizationService) DeleteOrganizationArgsForCall(i int) (context.Context, *model.DeleteOrganization) {
	fake.deleteOrganizationMutex.RLock()
	defer fake.deleteOrganizationMutex.RUnlock()
	argsForCall := fake.deleteOrganizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOrganizationService) DeleteOrganizationReturns(result1 error) {
	fake.deleteOrganizationMutex.Lock()
	defer fake.deleteOrganizationMutex.Unlock()
	fake.DeleteOrganizationStub = nil
	fake.deleteOrganizationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationService) DeleteOrganizationReturnsOnCall(i int, result1 error) {
	fake.deleteOrganizationMutex.Lock()
	defer fake.deleteOrganizationMutex.Unlock()
	fake.DeleteOrganizationStub = nil
	if fake.deleteOrganizationReturnsOnCall == nil {
		fake.deleteOrganizationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteOrganizationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationService) ListOrganizations(arg1 context.Context, arg2 *model.OrganizationFilters) (*model.PaginatedOrganizations, error) {
	fake.listOrganizationsMutex.Lock()
	ret, specificReturn := fake.listOrganizationsReturnsOnCall[len(fake.listOrganizationsArgsForCall)]
	fake.listOrganizationsArgsForCall = append(fake.listOrganizationsArgsForCall, struct {
		arg1 context.Context
		arg2 *model.OrganizationFilters
	}{arg1, arg2})
	stub := fake.ListOrganizationsStub
	fakeReturns := fake.listOrganizationsReturns
	fake.recordInvocation("ListOrganizations", []interface{}{arg1, arg2})
	fake.listOrganizationsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOrganizationService) ListOrganizationsCallCount() int {
	fake.listOrganizationsMutex.RLock()
	defer fake.listOrganizationsMutex.RUnlock()
	return len(fake.listOrganizationsArgsForCall)
}

func (fake *FakeOrganizationService) ListOrganizationsCalls(stub func(context.Context, *model.OrganizationFilters) (*model.PaginatedOrganizations, error)) {
	fake.listOrganizationsMutex.Lock()
	defer fake.listOrganizationsMutex.Unlock()
	fake.ListOrganizationsStub = stub
}

func (fake *FakeOrganizationService) ListOrganizationsArgsForCall(i int) (context.Context, *model.OrganizationFilters) {
	fake.listOrganizationsMutex.RLock()
	defer fake.listOrganizationsMutex.RUnlock()
	argsForCall := fake.listOrganizationsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOrganizationService) ListOrganizationsReturns(result1 *model.PaginatedOrganizations, result2 error) {
	fake.listOrganizationsMutex.Lock()
	defer fake.listOrganizationsMutex.Unlock()
	fake.ListOrganizationsStub = nil
	fake.listOrganizationsReturns = struct {
		result1 *model.PaginatedOrganizations
		result2 error
	}{result1, result2}
}

func (fake *FakeOrganizationService) ListOrganizationsReturnsOnCall(i int, result1 *model.PaginatedOrganizations, result2 error) {
	fake.listOrganizationsMutex.Lock()
	defer fake.listOrganizationsMutex.Unlock()
	fake.ListOrganizationsStub = nil
	if fake.listOrganizationsReturnsOnCall == nil {
		fake.listOrganizationsReturnsOnCall = make(map[int]struct {
			result1 *model.PaginatedOrganizations
			result2 error
		})
	}
	fake.listOrganizationsReturnsOnCall[i] = struct {
		result1 *model.PaginatedOrganizations
		result2 error
	}{result1, result2}
}

func (fake *FakeOrganizationService) RestoreOrganization(arg1 context.Context, arg2 *model.RestoreOrganization) error {
	fake.restoreOrganizationMutex.Lock()
	ret, specificReturn := fake.restoreOrganizationReturnsOnCall[len(fake.restoreOrganizationArgsForCall)]
	fake.restoreOrganizationArgsForCall = append(fake.restoreOrganizationArgsForCall, struct {
		arg1 context.Context
		arg2 *model.RestoreOrganization
	}{arg1, arg2})
	stub := fake.RestoreOrganizationStub
	fakeReturns := fake.restoreOrganizationReturns
	fake.recordInvocation("RestoreOrganization", []interface{}{arg1, arg2})
	fake.restoreOrganizationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOrganizationService) RestoreOrganizationCallCount() int {
	fake.restoreOrganizationMutex.RLock()
	defer fake.restoreOrganizationMutex.RUnlock()
	return len(fake.restoreOrganizationArgsForCall)
}

func (fake *FakeOrganizationService) RestoreOrganizationCalls(stub func(context.Context, *model.RestoreOrganization) error) {
	fake.restoreOrganizationMutex.Lock()
	defer fake.restoreOrganizationMutex.Unlock()
	fake.RestoreOrganizationStub = stub
}

func (fake *FakeOrganizationService) RestoreOrganizationArgsForCall(i int) (context.Context, *model.RestoreOrganization) {
	fake.restoreOrganizationMutex.RLock()
	defer fake.restoreOrganizationMutex.RUnlock()
	argsForCall := fake.restoreOrganizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOrganizationService) RestoreOrganizationReturns(result1 error) {
	fake.restoreOrganizationMutex.Lock()
	defer fake.restoreOrganizationMutex.Unlock()
	fake.RestoreOrganizationStub = nil
	fake.restoreOrganizationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationService) RestoreOrganizationReturnsOnCall(i int, result1 error) {
	fake.restoreOrganizationMutex.Lock()
	defer fake.restoreOrganizationMutex.Unlock()
	fake.RestoreOrganizationStub = nil
	if fake.restoreOrganizationReturnsOnCall == nil {
		fake.restoreOrganizationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.restoreOrganizationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createOrganizationMutex.RLock()
	defer fake.createOrganizationMutex.RUnlock()
	fake.deleteOrganizationMutex.RLock()
	defer fake.deleteOrganizationMutex.RUnlock()
	fake.listOrganizationsMutex.RLock()
	defer fake.listOrganizationsMutex.RUnlock()
	fake.restoreOrganizationMutex.RLock()
	defer fake.restoreOrganizationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOrganizationService) recordInvocation(key string, args []interface{}) {
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
