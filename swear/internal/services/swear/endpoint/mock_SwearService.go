// Code generated by mockery v2.46.2. DO NOT EDIT.

package endpoint

import (
	context "context"
	model "swearBot/internal/services/swear/model"

	mock "github.com/stretchr/testify/mock"
)

// MockSwearService is an autogenerated mock type for the SwearService type
type MockSwearService struct {
	mock.Mock
}

// AddSwears provides a mock function with given fields: _a0, _a1
func (_m *MockSwearService) AddSwears(_a0 context.Context, _a1 model.AddSwearsReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddSwears")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.AddSwearsReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockSwearService creates a new instance of MockSwearService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSwearService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSwearService {
	mock := &MockSwearService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
