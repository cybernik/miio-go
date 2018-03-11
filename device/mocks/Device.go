// Code generated by mockery v1.0.0
package mocks

import common "github.com/nickw444/miio-go/common"

import mock "github.com/stretchr/testify/mock"
import packet "github.com/nickw444/miio-go/protocol/packet"
import product "github.com/nickw444/miio-go/device/product"
import subscriptioncommon "github.com/nickw444/miio-go/subscription/common"
import time "time"
import transport "github.com/nickw444/miio-go/protocol/transport"

// Device is an autogenerated mock type for the Device type
type Device struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Device) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CloseAllSubscriptions provides a mock function with given fields:
func (_m *Device) CloseAllSubscriptions() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Discover provides a mock function with given fields:
func (_m *Device) Discover() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetInfo provides a mock function with given fields:
func (_m *Device) GetInfo() (common.DeviceInfo, error) {
	ret := _m.Called()

	var r0 common.DeviceInfo
	if rf, ok := ret.Get(0).(func() common.DeviceInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(common.DeviceInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLabel provides a mock function with given fields:
func (_m *Device) GetLabel() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProduct provides a mock function with given fields:
func (_m *Device) GetProduct() (product.Product, error) {
	ret := _m.Called()

	var r0 product.Product
	if rf, ok := ret.Get(0).(func() product.Product); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(product.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetToken provides a mock function with given fields:
func (_m *Device) GetToken() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// Handle provides a mock function with given fields: _a0
func (_m *Device) Handle(_a0 *packet.Packet) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*packet.Packet) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HasSubscribers provides a mock function with given fields:
func (_m *Device) HasSubscribers() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Device) ID() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// NewSubscription provides a mock function with given fields:
func (_m *Device) NewSubscription() (subscriptioncommon.Subscription, error) {
	ret := _m.Called()

	var r0 subscriptioncommon.Subscription
	if rf, ok := ret.Get(0).(func() subscriptioncommon.Subscription); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(subscriptioncommon.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Outbound provides a mock function with given fields:
func (_m *Device) Outbound() transport.Outbound {
	ret := _m.Called()

	var r0 transport.Outbound
	if rf, ok := ret.Get(0).(func() transport.Outbound); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(transport.Outbound)
		}
	}

	return r0
}

// Provisional provides a mock function with given fields:
func (_m *Device) Provisional() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Publish provides a mock function with given fields: event
func (_m *Device) Publish(event interface{}) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RefreshThrottle provides a mock function with given fields:
func (_m *Device) RefreshThrottle() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// RemoveSubscription provides a mock function with given fields: s
func (_m *Device) RemoveSubscription(s subscriptioncommon.Subscription) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(subscriptioncommon.Subscription) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Seen provides a mock function with given fields:
func (_m *Device) Seen() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// SetProvisional provides a mock function with given fields: _a0
func (_m *Device) SetProvisional(_a0 bool) {
	_m.Called(_a0)
}
