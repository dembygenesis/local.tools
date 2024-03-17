// Code generated by counterfeiter. DO NOT EDIT.
package categorylogicfakes

import (
	"context"
	"sync"

	"github.com/dembygenesis/local.tools/internal/model"
	"github.com/dembygenesis/local.tools/internal/persistence"
)

type FakePersistor struct {
	CreateCategoryStub        func(context.Context, persistence.TransactionHandler, *model.Category) (*model.Category, error)
	createCategoryMutex       sync.RWMutex
	createCategoryArgsForCall []struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 *model.Category
	}
	createCategoryReturns struct {
		result1 *model.Category
		result2 error
	}
	createCategoryReturnsOnCall map[int]struct {
		result1 *model.Category
		result2 error
	}
	GetCategoriesStub        func(context.Context, persistence.TransactionHandler, *model.CategoryFilters) (*model.PaginatedCategories, error)
	getCategoriesMutex       sync.RWMutex
	getCategoriesArgsForCall []struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 *model.CategoryFilters
	}
	getCategoriesReturns struct {
		result1 *model.PaginatedCategories
		result2 error
	}
	getCategoriesReturnsOnCall map[int]struct {
		result1 *model.PaginatedCategories
		result2 error
	}
	GetCategoryByNameStub        func(context.Context, persistence.TransactionHandler, string) (*model.Category, error)
	getCategoryByNameMutex       sync.RWMutex
	getCategoryByNameArgsForCall []struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 string
	}
	getCategoryByNameReturns struct {
		result1 *model.Category
		result2 error
	}
	getCategoryByNameReturnsOnCall map[int]struct {
		result1 *model.Category
		result2 error
	}
	GetCategoryTypeByIdStub        func(context.Context, persistence.TransactionHandler, int) (*model.CategoryType, error)
	getCategoryTypeByIdMutex       sync.RWMutex
	getCategoryTypeByIdArgsForCall []struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 int
	}
	getCategoryTypeByIdReturns struct {
		result1 *model.CategoryType
		result2 error
	}
	getCategoryTypeByIdReturnsOnCall map[int]struct {
		result1 *model.CategoryType
		result2 error
	}
	UpdateCategoryStub        func(context.Context, persistence.TransactionHandler, *model.UpdateCategory) (*model.Category, error)
	updateCategoryMutex       sync.RWMutex
	updateCategoryArgsForCall []struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 *model.UpdateCategory
	}
	updateCategoryReturns struct {
		result1 *model.Category
		result2 error
	}
	updateCategoryReturnsOnCall map[int]struct {
		result1 *model.Category
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePersistor) CreateCategory(arg1 context.Context, arg2 persistence.TransactionHandler, arg3 *model.Category) (*model.Category, error) {
	fake.createCategoryMutex.Lock()
	ret, specificReturn := fake.createCategoryReturnsOnCall[len(fake.createCategoryArgsForCall)]
	fake.createCategoryArgsForCall = append(fake.createCategoryArgsForCall, struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 *model.Category
	}{arg1, arg2, arg3})
	stub := fake.CreateCategoryStub
	fakeReturns := fake.createCategoryReturns
	fake.recordInvocation("CreateCategory", []interface{}{arg1, arg2, arg3})
	fake.createCategoryMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistor) CreateCategoryCallCount() int {
	fake.createCategoryMutex.RLock()
	defer fake.createCategoryMutex.RUnlock()
	return len(fake.createCategoryArgsForCall)
}

func (fake *FakePersistor) CreateCategoryCalls(stub func(context.Context, persistence.TransactionHandler, *model.Category) (*model.Category, error)) {
	fake.createCategoryMutex.Lock()
	defer fake.createCategoryMutex.Unlock()
	fake.CreateCategoryStub = stub
}

func (fake *FakePersistor) CreateCategoryArgsForCall(i int) (context.Context, persistence.TransactionHandler, *model.Category) {
	fake.createCategoryMutex.RLock()
	defer fake.createCategoryMutex.RUnlock()
	argsForCall := fake.createCategoryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistor) CreateCategoryReturns(result1 *model.Category, result2 error) {
	fake.createCategoryMutex.Lock()
	defer fake.createCategoryMutex.Unlock()
	fake.CreateCategoryStub = nil
	fake.createCategoryReturns = struct {
		result1 *model.Category
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) CreateCategoryReturnsOnCall(i int, result1 *model.Category, result2 error) {
	fake.createCategoryMutex.Lock()
	defer fake.createCategoryMutex.Unlock()
	fake.CreateCategoryStub = nil
	if fake.createCategoryReturnsOnCall == nil {
		fake.createCategoryReturnsOnCall = make(map[int]struct {
			result1 *model.Category
			result2 error
		})
	}
	fake.createCategoryReturnsOnCall[i] = struct {
		result1 *model.Category
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) GetCategories(arg1 context.Context, arg2 persistence.TransactionHandler, arg3 *model.CategoryFilters) (*model.PaginatedCategories, error) {
	fake.getCategoriesMutex.Lock()
	ret, specificReturn := fake.getCategoriesReturnsOnCall[len(fake.getCategoriesArgsForCall)]
	fake.getCategoriesArgsForCall = append(fake.getCategoriesArgsForCall, struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 *model.CategoryFilters
	}{arg1, arg2, arg3})
	stub := fake.GetCategoriesStub
	fakeReturns := fake.getCategoriesReturns
	fake.recordInvocation("GetCategories", []interface{}{arg1, arg2, arg3})
	fake.getCategoriesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistor) GetCategoriesCallCount() int {
	fake.getCategoriesMutex.RLock()
	defer fake.getCategoriesMutex.RUnlock()
	return len(fake.getCategoriesArgsForCall)
}

func (fake *FakePersistor) GetCategoriesCalls(stub func(context.Context, persistence.TransactionHandler, *model.CategoryFilters) (*model.PaginatedCategories, error)) {
	fake.getCategoriesMutex.Lock()
	defer fake.getCategoriesMutex.Unlock()
	fake.GetCategoriesStub = stub
}

func (fake *FakePersistor) GetCategoriesArgsForCall(i int) (context.Context, persistence.TransactionHandler, *model.CategoryFilters) {
	fake.getCategoriesMutex.RLock()
	defer fake.getCategoriesMutex.RUnlock()
	argsForCall := fake.getCategoriesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistor) GetCategoriesReturns(result1 *model.PaginatedCategories, result2 error) {
	fake.getCategoriesMutex.Lock()
	defer fake.getCategoriesMutex.Unlock()
	fake.GetCategoriesStub = nil
	fake.getCategoriesReturns = struct {
		result1 *model.PaginatedCategories
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) GetCategoriesReturnsOnCall(i int, result1 *model.PaginatedCategories, result2 error) {
	fake.getCategoriesMutex.Lock()
	defer fake.getCategoriesMutex.Unlock()
	fake.GetCategoriesStub = nil
	if fake.getCategoriesReturnsOnCall == nil {
		fake.getCategoriesReturnsOnCall = make(map[int]struct {
			result1 *model.PaginatedCategories
			result2 error
		})
	}
	fake.getCategoriesReturnsOnCall[i] = struct {
		result1 *model.PaginatedCategories
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) GetCategoryByName(arg1 context.Context, arg2 persistence.TransactionHandler, arg3 string) (*model.Category, error) {
	fake.getCategoryByNameMutex.Lock()
	ret, specificReturn := fake.getCategoryByNameReturnsOnCall[len(fake.getCategoryByNameArgsForCall)]
	fake.getCategoryByNameArgsForCall = append(fake.getCategoryByNameArgsForCall, struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GetCategoryByNameStub
	fakeReturns := fake.getCategoryByNameReturns
	fake.recordInvocation("GetCategoryByName", []interface{}{arg1, arg2, arg3})
	fake.getCategoryByNameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistor) GetCategoryByNameCallCount() int {
	fake.getCategoryByNameMutex.RLock()
	defer fake.getCategoryByNameMutex.RUnlock()
	return len(fake.getCategoryByNameArgsForCall)
}

func (fake *FakePersistor) GetCategoryByNameCalls(stub func(context.Context, persistence.TransactionHandler, string) (*model.Category, error)) {
	fake.getCategoryByNameMutex.Lock()
	defer fake.getCategoryByNameMutex.Unlock()
	fake.GetCategoryByNameStub = stub
}

func (fake *FakePersistor) GetCategoryByNameArgsForCall(i int) (context.Context, persistence.TransactionHandler, string) {
	fake.getCategoryByNameMutex.RLock()
	defer fake.getCategoryByNameMutex.RUnlock()
	argsForCall := fake.getCategoryByNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistor) GetCategoryByNameReturns(result1 *model.Category, result2 error) {
	fake.getCategoryByNameMutex.Lock()
	defer fake.getCategoryByNameMutex.Unlock()
	fake.GetCategoryByNameStub = nil
	fake.getCategoryByNameReturns = struct {
		result1 *model.Category
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) GetCategoryByNameReturnsOnCall(i int, result1 *model.Category, result2 error) {
	fake.getCategoryByNameMutex.Lock()
	defer fake.getCategoryByNameMutex.Unlock()
	fake.GetCategoryByNameStub = nil
	if fake.getCategoryByNameReturnsOnCall == nil {
		fake.getCategoryByNameReturnsOnCall = make(map[int]struct {
			result1 *model.Category
			result2 error
		})
	}
	fake.getCategoryByNameReturnsOnCall[i] = struct {
		result1 *model.Category
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) GetCategoryTypeById(arg1 context.Context, arg2 persistence.TransactionHandler, arg3 int) (*model.CategoryType, error) {
	fake.getCategoryTypeByIdMutex.Lock()
	ret, specificReturn := fake.getCategoryTypeByIdReturnsOnCall[len(fake.getCategoryTypeByIdArgsForCall)]
	fake.getCategoryTypeByIdArgsForCall = append(fake.getCategoryTypeByIdArgsForCall, struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 int
	}{arg1, arg2, arg3})
	stub := fake.GetCategoryTypeByIdStub
	fakeReturns := fake.getCategoryTypeByIdReturns
	fake.recordInvocation("GetCategoryTypeById", []interface{}{arg1, arg2, arg3})
	fake.getCategoryTypeByIdMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistor) GetCategoryTypeByIdCallCount() int {
	fake.getCategoryTypeByIdMutex.RLock()
	defer fake.getCategoryTypeByIdMutex.RUnlock()
	return len(fake.getCategoryTypeByIdArgsForCall)
}

func (fake *FakePersistor) GetCategoryTypeByIdCalls(stub func(context.Context, persistence.TransactionHandler, int) (*model.CategoryType, error)) {
	fake.getCategoryTypeByIdMutex.Lock()
	defer fake.getCategoryTypeByIdMutex.Unlock()
	fake.GetCategoryTypeByIdStub = stub
}

func (fake *FakePersistor) GetCategoryTypeByIdArgsForCall(i int) (context.Context, persistence.TransactionHandler, int) {
	fake.getCategoryTypeByIdMutex.RLock()
	defer fake.getCategoryTypeByIdMutex.RUnlock()
	argsForCall := fake.getCategoryTypeByIdArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistor) GetCategoryTypeByIdReturns(result1 *model.CategoryType, result2 error) {
	fake.getCategoryTypeByIdMutex.Lock()
	defer fake.getCategoryTypeByIdMutex.Unlock()
	fake.GetCategoryTypeByIdStub = nil
	fake.getCategoryTypeByIdReturns = struct {
		result1 *model.CategoryType
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) GetCategoryTypeByIdReturnsOnCall(i int, result1 *model.CategoryType, result2 error) {
	fake.getCategoryTypeByIdMutex.Lock()
	defer fake.getCategoryTypeByIdMutex.Unlock()
	fake.GetCategoryTypeByIdStub = nil
	if fake.getCategoryTypeByIdReturnsOnCall == nil {
		fake.getCategoryTypeByIdReturnsOnCall = make(map[int]struct {
			result1 *model.CategoryType
			result2 error
		})
	}
	fake.getCategoryTypeByIdReturnsOnCall[i] = struct {
		result1 *model.CategoryType
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) UpdateCategory(arg1 context.Context, arg2 persistence.TransactionHandler, arg3 *model.UpdateCategory) (*model.Category, error) {
	fake.updateCategoryMutex.Lock()
	ret, specificReturn := fake.updateCategoryReturnsOnCall[len(fake.updateCategoryArgsForCall)]
	fake.updateCategoryArgsForCall = append(fake.updateCategoryArgsForCall, struct {
		arg1 context.Context
		arg2 persistence.TransactionHandler
		arg3 *model.UpdateCategory
	}{arg1, arg2, arg3})
	stub := fake.UpdateCategoryStub
	fakeReturns := fake.updateCategoryReturns
	fake.recordInvocation("UpdateCategory", []interface{}{arg1, arg2, arg3})
	fake.updateCategoryMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistor) UpdateCategoryCallCount() int {
	fake.updateCategoryMutex.RLock()
	defer fake.updateCategoryMutex.RUnlock()
	return len(fake.updateCategoryArgsForCall)
}

func (fake *FakePersistor) UpdateCategoryCalls(stub func(context.Context, persistence.TransactionHandler, *model.UpdateCategory) (*model.Category, error)) {
	fake.updateCategoryMutex.Lock()
	defer fake.updateCategoryMutex.Unlock()
	fake.UpdateCategoryStub = stub
}

func (fake *FakePersistor) UpdateCategoryArgsForCall(i int) (context.Context, persistence.TransactionHandler, *model.UpdateCategory) {
	fake.updateCategoryMutex.RLock()
	defer fake.updateCategoryMutex.RUnlock()
	argsForCall := fake.updateCategoryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistor) UpdateCategoryReturns(result1 *model.Category, result2 error) {
	fake.updateCategoryMutex.Lock()
	defer fake.updateCategoryMutex.Unlock()
	fake.UpdateCategoryStub = nil
	fake.updateCategoryReturns = struct {
		result1 *model.Category
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) UpdateCategoryReturnsOnCall(i int, result1 *model.Category, result2 error) {
	fake.updateCategoryMutex.Lock()
	defer fake.updateCategoryMutex.Unlock()
	fake.UpdateCategoryStub = nil
	if fake.updateCategoryReturnsOnCall == nil {
		fake.updateCategoryReturnsOnCall = make(map[int]struct {
			result1 *model.Category
			result2 error
		})
	}
	fake.updateCategoryReturnsOnCall[i] = struct {
		result1 *model.Category
		result2 error
	}{result1, result2}
}

func (fake *FakePersistor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCategoryMutex.RLock()
	defer fake.createCategoryMutex.RUnlock()
	fake.getCategoriesMutex.RLock()
	defer fake.getCategoriesMutex.RUnlock()
	fake.getCategoryByNameMutex.RLock()
	defer fake.getCategoryByNameMutex.RUnlock()
	fake.getCategoryTypeByIdMutex.RLock()
	defer fake.getCategoryTypeByIdMutex.RUnlock()
	fake.updateCategoryMutex.RLock()
	defer fake.updateCategoryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePersistor) recordInvocation(key string, args []interface{}) {
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
