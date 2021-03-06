// Copyright 2017 Northern.tech AS
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
import mock "github.com/stretchr/testify/mock"
import model "github.com/mendersoftware/deviceadm/model"
import store "github.com/mendersoftware/deviceadm/store"

// DataStore is an autogenerated mock type for the DataStore type
type DataStore struct {
	mock.Mock
}

// DeleteDeviceAuth provides a mock function with given fields: ctx, id
func (_m *DataStore) DeleteDeviceAuth(ctx context.Context, id model.AuthID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.AuthID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDeviceAuthByDevice provides a mock function with given fields: ctx, id
func (_m *DataStore) DeleteDeviceAuthByDevice(ctx context.Context, id model.DeviceID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DeviceID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDeviceAuth provides a mock function with given fields: ctx, id
func (_m *DataStore) GetDeviceAuth(ctx context.Context, id model.AuthID) (*model.DeviceAuth, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.DeviceAuth
	if rf, ok := ret.Get(0).(func(context.Context, model.AuthID) *model.DeviceAuth); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeviceAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.AuthID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceAuths provides a mock function with given fields: ctx, skip, limit, filter
func (_m *DataStore) GetDeviceAuths(ctx context.Context, skip int, limit int, filter store.Filter) ([]model.DeviceAuth, error) {
	ret := _m.Called(ctx, skip, limit, filter)

	var r0 []model.DeviceAuth
	if rf, ok := ret.Get(0).(func(context.Context, int, int, store.Filter) []model.DeviceAuth); ok {
		r0 = rf(ctx, skip, limit, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DeviceAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int, store.Filter) error); ok {
		r1 = rf(ctx, skip, limit, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceAuthsByIdentityData provides a mock function with given fields: ctx, idata
func (_m *DataStore) GetDeviceAuthsByIdentityData(ctx context.Context, idata string) ([]model.DeviceAuth, error) {
	ret := _m.Called(ctx, idata)

	var r0 []model.DeviceAuth
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.DeviceAuth); ok {
		r0 = rf(ctx, idata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DeviceAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, idata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertDeviceAuth provides a mock function with given fields: ctx, dev
func (_m *DataStore) InsertDeviceAuth(ctx context.Context, dev *model.DeviceAuth) error {
	ret := _m.Called(ctx, dev)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.DeviceAuth) error); ok {
		r0 = rf(ctx, dev)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MigrateTenant provides a mock function with given fields: ctx, version, tenant
func (_m *DataStore) MigrateTenant(ctx context.Context, version string, tenant string) error {
	ret := _m.Called(ctx, version, tenant)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, version, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutDeviceAuth provides a mock function with given fields: ctx, dev
func (_m *DataStore) PutDeviceAuth(ctx context.Context, dev *model.DeviceAuth) error {
	ret := _m.Called(ctx, dev)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.DeviceAuth) error); ok {
		r0 = rf(ctx, dev)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeviceAuth provides a mock function with given fields: ctx, dev
func (_m *DataStore) UpdateDeviceAuth(ctx context.Context, dev *model.DeviceAuth) error {
	ret := _m.Called(ctx, dev)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.DeviceAuth) error); ok {
		r0 = rf(ctx, dev)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithAutomigrate provides a mock function with given fields:
func (_m *DataStore) WithAutomigrate() store.DataStore {
	ret := _m.Called()

	var r0 store.DataStore
	if rf, ok := ret.Get(0).(func() store.DataStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.DataStore)
		}
	}

	return r0
}
