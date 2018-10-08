// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v6"
)

type FakeV3CancelZdtPushActor struct {
	CancelDeploymentByAppNameAndSpaceStub        func(appName string, spaceGUID string) (v3action.Warnings, error)
	cancelDeploymentByAppNameAndSpaceMutex       sync.RWMutex
	cancelDeploymentByAppNameAndSpaceArgsForCall []struct {
		appName   string
		spaceGUID string
	}
	cancelDeploymentByAppNameAndSpaceReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	cancelDeploymentByAppNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3CancelZdtPushActor) CancelDeploymentByAppNameAndSpace(appName string, spaceGUID string) (v3action.Warnings, error) {
	fake.cancelDeploymentByAppNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.cancelDeploymentByAppNameAndSpaceReturnsOnCall[len(fake.cancelDeploymentByAppNameAndSpaceArgsForCall)]
	fake.cancelDeploymentByAppNameAndSpaceArgsForCall = append(fake.cancelDeploymentByAppNameAndSpaceArgsForCall, struct {
		appName   string
		spaceGUID string
	}{appName, spaceGUID})
	fake.recordInvocation("CancelDeploymentByAppNameAndSpace", []interface{}{appName, spaceGUID})
	fake.cancelDeploymentByAppNameAndSpaceMutex.Unlock()
	if fake.CancelDeploymentByAppNameAndSpaceStub != nil {
		return fake.CancelDeploymentByAppNameAndSpaceStub(appName, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.cancelDeploymentByAppNameAndSpaceReturns.result1, fake.cancelDeploymentByAppNameAndSpaceReturns.result2
}

func (fake *FakeV3CancelZdtPushActor) CancelDeploymentByAppNameAndSpaceCallCount() int {
	fake.cancelDeploymentByAppNameAndSpaceMutex.RLock()
	defer fake.cancelDeploymentByAppNameAndSpaceMutex.RUnlock()
	return len(fake.cancelDeploymentByAppNameAndSpaceArgsForCall)
}

func (fake *FakeV3CancelZdtPushActor) CancelDeploymentByAppNameAndSpaceArgsForCall(i int) (string, string) {
	fake.cancelDeploymentByAppNameAndSpaceMutex.RLock()
	defer fake.cancelDeploymentByAppNameAndSpaceMutex.RUnlock()
	return fake.cancelDeploymentByAppNameAndSpaceArgsForCall[i].appName, fake.cancelDeploymentByAppNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV3CancelZdtPushActor) CancelDeploymentByAppNameAndSpaceReturns(result1 v3action.Warnings, result2 error) {
	fake.CancelDeploymentByAppNameAndSpaceStub = nil
	fake.cancelDeploymentByAppNameAndSpaceReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3CancelZdtPushActor) CancelDeploymentByAppNameAndSpaceReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.CancelDeploymentByAppNameAndSpaceStub = nil
	if fake.cancelDeploymentByAppNameAndSpaceReturnsOnCall == nil {
		fake.cancelDeploymentByAppNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.cancelDeploymentByAppNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3CancelZdtPushActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeV3CancelZdtPushActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3CancelZdtPushActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3CancelZdtPushActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3CancelZdtPushActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cancelDeploymentByAppNameAndSpaceMutex.RLock()
	defer fake.cancelDeploymentByAppNameAndSpaceMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3CancelZdtPushActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.V3CancelZdtPushActor = new(FakeV3CancelZdtPushActor)