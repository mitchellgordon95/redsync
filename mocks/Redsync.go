package mocks

import "github.com/mitchellgordon95/redsync"
import "github.com/stretchr/testify/mock"

type Redsync struct {
	mock.Mock
}

// NewMutex provides a mock function with given fields: name, options
func (_m *Redsync) NewMutex(name string, options ...redsync.Option) redsync.Mutex {
	ret := _m.Called(name, options)

	var r0 redsync.Mutex
	if rf, ok := ret.Get(0).(func(string, ...redsync.Option) redsync.Mutex); ok {
		r0 = rf(name, options...)
	} else {
		r0 = ret.Get(0).(redsync.Mutex)
	}

	return r0
}
