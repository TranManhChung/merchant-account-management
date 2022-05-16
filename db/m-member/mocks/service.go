// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "main.go/model"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, account
func (_m *Service) Add(ctx context.Context, account model.MerchantMember) error {
	ret := _m.Called(ctx, account)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.MerchantMember) error); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, email
func (_m *Service) Delete(ctx context.Context, email string) error {
	ret := _m.Called(ctx, email)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, email
func (_m *Service) Get(ctx context.Context, email string) (model.MerchantMemberEntity, error) {
	ret := _m.Called(ctx, email)

	var r0 model.MerchantMemberEntity
	if rf, ok := ret.Get(0).(func(context.Context, string) model.MerchantMemberEntity); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(model.MerchantMemberEntity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsExisted provides a mock function with given fields: ctx, email
func (_m *Service) IsExisted(ctx context.Context, email string) (bool, error) {
	ret := _m.Called(ctx, email)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, account
func (_m *Service) Update(ctx context.Context, account model.MerchantMember) error {
	ret := _m.Called(ctx, account)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.MerchantMember) error); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}