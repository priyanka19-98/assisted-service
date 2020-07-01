// Code generated by mockery v1.0.0. DO NOT EDIT.

package versions

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockAPI is an autogenerated mock type for the API type
type MockAPI struct {
	mock.Mock
}

// ListComponentVersions provides a mock function with given fields: ctx, params
func (_m *MockAPI) ListComponentVersions(ctx context.Context, params *ListComponentVersionsParams) (*ListComponentVersionsOK, error) {
	ret := _m.Called(ctx, params)

	var r0 *ListComponentVersionsOK
	if rf, ok := ret.Get(0).(func(context.Context, *ListComponentVersionsParams) *ListComponentVersionsOK); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ListComponentVersionsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ListComponentVersionsParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}