// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	candiutils "github.com/golangid/candi/candiutils"
	mock "github.com/stretchr/testify/mock"
)

// LockerOption is an autogenerated mock type for the LockerOption type
type LockerOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *LockerOption) Execute(_a0 *candiutils.LockerOptions) {
	_m.Called(_a0)
}

// NewLockerOption creates a new instance of LockerOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLockerOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *LockerOption {
	mock := &LockerOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
