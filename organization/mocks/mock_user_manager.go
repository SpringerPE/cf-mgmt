// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/pivotalservices/cf-mgmt/organization (interfaces: UserMgr)

package mock_organization

import (
	gomock "github.com/golang/mock/gomock"
	ldap "github.com/pivotalservices/cf-mgmt/ldap"
	organization "github.com/pivotalservices/cf-mgmt/organization"
)

// Mock of UserMgr interface
type MockUserMgr struct {
	ctrl     *gomock.Controller
	recorder *_MockUserMgrRecorder
}

// Recorder for MockUserMgr (not exported)
type _MockUserMgrRecorder struct {
	mock *MockUserMgr
}

func NewMockUserMgr(ctrl *gomock.Controller) *MockUserMgr {
	mock := &MockUserMgr{ctrl: ctrl}
	mock.recorder = &_MockUserMgrRecorder{mock}
	return mock
}

func (_m *MockUserMgr) EXPECT() *_MockUserMgrRecorder {
	return _m.recorder
}

func (_m *MockUserMgr) UpdateOrgUsers(_param0 *ldap.Config, _param1 map[string]string, _param2 organization.UpdateUsersInput) error {
	ret := _m.ctrl.Call(_m, "UpdateOrgUsers", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockUserMgrRecorder) UpdateOrgUsers(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateOrgUsers", arg0, arg1, arg2)
}
