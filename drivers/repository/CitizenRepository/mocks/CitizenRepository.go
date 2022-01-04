// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	records "vaccine-app-be/drivers/records"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// CitizenRepository is an autogenerated mock type for the CitizenRepository type
type CitizenRepository struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *CitizenRepository) FindByEmail(ctx context.Context, email string) (records.Citizen, error) {
	ret := _m.Called(ctx, email)

	var r0 records.Citizen
	if rf, ok := ret.Get(0).(func(context.Context, string) records.Citizen); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(records.Citizen)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, citizens
func (_m *CitizenRepository) Register(ctx context.Context, citizens records.Citizen) (records.Citizen, error) {
	ret := _m.Called(ctx, citizens)

	var r0 records.Citizen
	if rf, ok := ret.Get(0).(func(context.Context, records.Citizen) records.Citizen); ok {
		r0 = rf(ctx, citizens)
	} else {
		r0 = ret.Get(0).(records.Citizen)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, records.Citizen) error); ok {
		r1 = rf(ctx, citizens)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, userId, birthDay, address
func (_m *CitizenRepository) Update(ctx context.Context, userId int, birthDay time.Time, address string) (records.Citizen, error) {
	ret := _m.Called(ctx, userId, birthDay, address)

	var r0 records.Citizen
	if rf, ok := ret.Get(0).(func(context.Context, int, time.Time, string) records.Citizen); ok {
		r0 = rf(ctx, userId, birthDay, address)
	} else {
		r0 = ret.Get(0).(records.Citizen)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, time.Time, string) error); ok {
		r1 = rf(ctx, userId, birthDay, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
