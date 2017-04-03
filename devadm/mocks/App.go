// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package mocks

import context "context"
import devadm "github.com/mendersoftware/deviceadm/devadm"
import mock "github.com/stretchr/testify/mock"
import model "github.com/mendersoftware/deviceadm/model"

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// AcceptDeviceAuth provides a mock function with given fields: id
func (_m *App) AcceptDeviceAuth(id model.AuthID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.AuthID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDeviceAuth provides a mock function with given fields: id
func (_m *App) DeleteDeviceAuth(id model.AuthID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.AuthID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDeviceData provides a mock function with given fields: id
func (_m *App) DeleteDeviceData(id model.DeviceID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.DeviceID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDeviceAuth provides a mock function with given fields: id
func (_m *App) GetDeviceAuth(id model.AuthID) (*model.DeviceAuth, error) {
	ret := _m.Called(id)

	var r0 *model.DeviceAuth
	if rf, ok := ret.Get(0).(func(model.AuthID) *model.DeviceAuth); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeviceAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.AuthID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDeviceAuths provides a mock function with given fields: skip, limit, status
func (_m *App) ListDeviceAuths(skip int, limit int, status string) ([]model.DeviceAuth, error) {
	ret := _m.Called(skip, limit, status)

	var r0 []model.DeviceAuth
	if rf, ok := ret.Get(0).(func(int, int, string) []model.DeviceAuth); ok {
		r0 = rf(skip, limit, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DeviceAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(skip, limit, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RejectDeviceAuth provides a mock function with given fields: id
func (_m *App) RejectDeviceAuth(id model.AuthID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.AuthID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubmitDeviceAuth provides a mock function with given fields: d
func (_m *App) SubmitDeviceAuth(d model.DeviceAuth) error {
	ret := _m.Called(d)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.DeviceAuth) error); ok {
		r0 = rf(d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithContext provides a mock function with given fields: c
func (_m *App) WithContext(c context.Context) devadm.App {
	ret := _m.Called(c)

	var r0 devadm.App
	if rf, ok := ret.Get(0).(func(context.Context) devadm.App); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(devadm.App)
		}
	}

	return r0
}

var _ devadm.App = (*App)(nil)
