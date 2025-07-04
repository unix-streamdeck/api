// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	dbus "github.com/godbus/dbus/v5"
	mock "github.com/stretchr/testify/mock"
)

// BusObject is an autogenerated mock type for the BusObject type
type BusObject struct {
	mock.Mock
}

// AddMatchSignal provides a mock function with given fields: iface, member, options
func (_m *BusObject) AddMatchSignal(iface string, member string, options ...dbus.MatchOption) *dbus.Call {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, iface, member)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *dbus.Call
	if rf, ok := ret.Get(0).(func(string, string, ...dbus.MatchOption) *dbus.Call); ok {
		r0 = rf(iface, member, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dbus.Call)
		}
	}

	return r0
}

// Call provides a mock function with given fields: method, flags, args
func (_m *BusObject) Call(method string, flags dbus.Flags, args ...interface{}) *dbus.Call {
	var _ca []interface{}
	_ca = append(_ca, method, flags)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *dbus.Call
	if rf, ok := ret.Get(0).(func(string, dbus.Flags, ...interface{}) *dbus.Call); ok {
		r0 = rf(method, flags, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dbus.Call)
		}
	}

	return r0
}

// CallWithContext provides a mock function with given fields: ctx, method, flags, args
func (_m *BusObject) CallWithContext(ctx context.Context, method string, flags dbus.Flags, args ...interface{}) *dbus.Call {
	var _ca []interface{}
	_ca = append(_ca, ctx, method, flags)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *dbus.Call
	if rf, ok := ret.Get(0).(func(context.Context, string, dbus.Flags, ...interface{}) *dbus.Call); ok {
		r0 = rf(ctx, method, flags, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dbus.Call)
		}
	}

	return r0
}

// Destination provides a mock function with given fields:
func (_m *BusObject) Destination() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetProperty provides a mock function with given fields: p
func (_m *BusObject) GetProperty(p string) (dbus.Variant, error) {
	ret := _m.Called(p)

	var r0 dbus.Variant
	if rf, ok := ret.Get(0).(func(string) dbus.Variant); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(dbus.Variant)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Go provides a mock function with given fields: method, flags, ch, args
func (_m *BusObject) Go(method string, flags dbus.Flags, ch chan *dbus.Call, args ...interface{}) *dbus.Call {
	var _ca []interface{}
	_ca = append(_ca, method, flags, ch)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *dbus.Call
	if rf, ok := ret.Get(0).(func(string, dbus.Flags, chan *dbus.Call, ...interface{}) *dbus.Call); ok {
		r0 = rf(method, flags, ch, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dbus.Call)
		}
	}

	return r0
}

// GoWithContext provides a mock function with given fields: ctx, method, flags, ch, args
func (_m *BusObject) GoWithContext(ctx context.Context, method string, flags dbus.Flags, ch chan *dbus.Call, args ...interface{}) *dbus.Call {
	var _ca []interface{}
	_ca = append(_ca, ctx, method, flags, ch)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *dbus.Call
	if rf, ok := ret.Get(0).(func(context.Context, string, dbus.Flags, chan *dbus.Call, ...interface{}) *dbus.Call); ok {
		r0 = rf(ctx, method, flags, ch, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dbus.Call)
		}
	}

	return r0
}

// Path provides a mock function with given fields:
func (_m *BusObject) Path() dbus.ObjectPath {
	ret := _m.Called()

	var r0 dbus.ObjectPath
	if rf, ok := ret.Get(0).(func() dbus.ObjectPath); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(dbus.ObjectPath)
	}

	return r0
}

// RemoveMatchSignal provides a mock function with given fields: iface, member, options
func (_m *BusObject) RemoveMatchSignal(iface string, member string, options ...dbus.MatchOption) *dbus.Call {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, iface, member)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *dbus.Call
	if rf, ok := ret.Get(0).(func(string, string, ...dbus.MatchOption) *dbus.Call); ok {
		r0 = rf(iface, member, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dbus.Call)
		}
	}

	return r0
}

// SetProperty provides a mock function with given fields: p, v
func (_m *BusObject) SetProperty(p string, v interface{}) error {
	ret := _m.Called(p, v)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(p, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreProperty provides a mock function with given fields: p, value
func (_m *BusObject) StoreProperty(p string, value interface{}) error {
	ret := _m.Called(p, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(p, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBusObject interface {
	mock.TestingT
	Cleanup(func())
}

// NewBusObject creates a new instance of BusObject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBusObject(t mockConstructorTestingTNewBusObject) *BusObject {
	mock := &BusObject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
