// Code generated by MockGen. DO NOT EDIT.
// Source: expected_keepers.go

// Package types is a generated GoMock package.
package types

import (
	context "context"
	reflect "reflect"

	math "cosmossdk.io/math"
	types "github.com/bnb-chain/greenfield/x/sp/types"
	types0 "github.com/bnb-chain/greenfield/x/storage/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/cosmos/cosmos-sdk/x/staking/types"
	gomock "github.com/golang/mock/gomock"
)

// MockSpKeeper is a mock of SpKeeper interface.
type MockSpKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockSpKeeperMockRecorder
}

// MockSpKeeperMockRecorder is the mock recorder for MockSpKeeper.
type MockSpKeeperMockRecorder struct {
	mock *MockSpKeeper
}

// NewMockSpKeeper creates a new mock instance.
func NewMockSpKeeper(ctrl *gomock.Controller) *MockSpKeeper {
	mock := &MockSpKeeper{ctrl: ctrl}
	mock.recorder = &MockSpKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpKeeper) EXPECT() *MockSpKeeperMockRecorder {
	return m.recorder
}

// DepositDenomForSP mocks base method.
func (m *MockSpKeeper) DepositDenomForSP(ctx types1.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DepositDenomForSP", ctx)
	ret0, _ := ret[0].(string)
	return ret0
}

// DepositDenomForSP indicates an expected call of DepositDenomForSP.
func (mr *MockSpKeeperMockRecorder) DepositDenomForSP(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepositDenomForSP", reflect.TypeOf((*MockSpKeeper)(nil).DepositDenomForSP), ctx)
}

// GetStorageProvider mocks base method.
func (m *MockSpKeeper) GetStorageProvider(ctx types1.Context, addr types1.AccAddress) (*types.StorageProvider, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageProvider", ctx, addr)
	ret0, _ := ret[0].(*types.StorageProvider)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStorageProvider indicates an expected call of GetStorageProvider.
func (mr *MockSpKeeperMockRecorder) GetStorageProvider(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageProvider", reflect.TypeOf((*MockSpKeeper)(nil).GetStorageProvider), ctx, addr)
}

// Slash mocks base method.
func (m *MockSpKeeper) Slash(ctx types1.Context, spAcc types1.AccAddress, rewardInfos []types.RewardInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Slash", ctx, spAcc, rewardInfos)
	ret0, _ := ret[0].(error)
	return ret0
}

// Slash indicates an expected call of Slash.
func (mr *MockSpKeeperMockRecorder) Slash(ctx, spAcc, rewardInfos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Slash", reflect.TypeOf((*MockSpKeeper)(nil).Slash), ctx, spAcc, rewardInfos)
}

// MockStakingKeeper is a mock of StakingKeeper interface.
type MockStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStakingKeeperMockRecorder
}

// MockStakingKeeperMockRecorder is the mock recorder for MockStakingKeeper.
type MockStakingKeeperMockRecorder struct {
	mock *MockStakingKeeper
}

// NewMockStakingKeeper creates a new mock instance.
func NewMockStakingKeeper(ctrl *gomock.Controller) *MockStakingKeeper {
	mock := &MockStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingKeeper) EXPECT() *MockStakingKeeperMockRecorder {
	return m.recorder
}

// GetHistoricalInfo mocks base method.
func (m *MockStakingKeeper) GetHistoricalInfo(ctx types1.Context, height int64) (types2.HistoricalInfo, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoricalInfo", ctx, height)
	ret0, _ := ret[0].(types2.HistoricalInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetHistoricalInfo indicates an expected call of GetHistoricalInfo.
func (mr *MockStakingKeeperMockRecorder) GetHistoricalInfo(ctx, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoricalInfo", reflect.TypeOf((*MockStakingKeeper)(nil).GetHistoricalInfo), ctx, height)
}

// GetLastValidators mocks base method.
func (m *MockStakingKeeper) GetLastValidators(ctx types1.Context) []types2.Validator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastValidators", ctx)
	ret0, _ := ret[0].([]types2.Validator)
	return ret0
}

// GetLastValidators indicates an expected call of GetLastValidators.
func (mr *MockStakingKeeperMockRecorder) GetLastValidators(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastValidators", reflect.TypeOf((*MockStakingKeeper)(nil).GetLastValidators), ctx)
}

// MockStorageKeeper is a mock of StorageKeeper interface.
type MockStorageKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStorageKeeperMockRecorder
}

// MockStorageKeeperMockRecorder is the mock recorder for MockStorageKeeper.
type MockStorageKeeperMockRecorder struct {
	mock *MockStorageKeeper
}

// NewMockStorageKeeper creates a new mock instance.
func NewMockStorageKeeper(ctrl *gomock.Controller) *MockStorageKeeper {
	mock := &MockStorageKeeper{ctrl: ctrl}
	mock.recorder = &MockStorageKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageKeeper) EXPECT() *MockStorageKeeperMockRecorder {
	return m.recorder
}

// GetBucketInfo mocks base method.
func (m *MockStorageKeeper) GetBucketInfo(ctx types1.Context, bucketName string) (*types0.BucketInfo, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBucketInfo", ctx, bucketName)
	ret0, _ := ret[0].(*types0.BucketInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetBucketInfo indicates an expected call of GetBucketInfo.
func (mr *MockStorageKeeperMockRecorder) GetBucketInfo(ctx, bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketInfo", reflect.TypeOf((*MockStorageKeeper)(nil).GetBucketInfo), ctx, bucketName)
}

// GetObjectInfo mocks base method.
func (m *MockStorageKeeper) GetObjectInfo(ctx types1.Context, bucketName, objectName string) (*types0.ObjectInfo, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectInfo", ctx, bucketName, objectName)
	ret0, _ := ret[0].(*types0.ObjectInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetObjectInfo indicates an expected call of GetObjectInfo.
func (mr *MockStorageKeeperMockRecorder) GetObjectInfo(ctx, bucketName, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectInfo", reflect.TypeOf((*MockStorageKeeper)(nil).GetObjectInfo), ctx, bucketName, objectName)
}

// GetObjectInfoById mocks base method.
func (m *MockStorageKeeper) GetObjectInfoById(ctx types1.Context, objectId math.Uint) (*types0.ObjectInfo, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectInfoById", ctx, objectId)
	ret0, _ := ret[0].(*types0.ObjectInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetObjectInfoById indicates an expected call of GetObjectInfoById.
func (mr *MockStorageKeeperMockRecorder) GetObjectInfoById(ctx, objectId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectInfoById", reflect.TypeOf((*MockStorageKeeper)(nil).GetObjectInfoById), ctx, objectId)
}

// GetObjectInfoCount mocks base method.
func (m *MockStorageKeeper) GetObjectInfoCount(ctx types1.Context) math.Uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectInfoCount", ctx)
	ret0, _ := ret[0].(math.Uint)
	return ret0
}

// GetObjectInfoCount indicates an expected call of GetObjectInfoCount.
func (mr *MockStorageKeeperMockRecorder) GetObjectInfoCount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectInfoCount", reflect.TypeOf((*MockStorageKeeper)(nil).GetObjectInfoCount), ctx)
}

// MaxSegmentSize mocks base method.
func (m *MockStorageKeeper) MaxSegmentSize(ctx types1.Context) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxSegmentSize", ctx)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// MaxSegmentSize indicates an expected call of MaxSegmentSize.
func (mr *MockStorageKeeperMockRecorder) MaxSegmentSize(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxSegmentSize", reflect.TypeOf((*MockStorageKeeper)(nil).MaxSegmentSize), ctx)
}

// MockPaymentKeeper is a mock of PaymentKeeper interface.
type MockPaymentKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentKeeperMockRecorder
}

// MockPaymentKeeperMockRecorder is the mock recorder for MockPaymentKeeper.
type MockPaymentKeeperMockRecorder struct {
	mock *MockPaymentKeeper
}

// NewMockPaymentKeeper creates a new mock instance.
func NewMockPaymentKeeper(ctrl *gomock.Controller) *MockPaymentKeeper {
	mock := &MockPaymentKeeper{ctrl: ctrl}
	mock.recorder = &MockPaymentKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentKeeper) EXPECT() *MockPaymentKeeperMockRecorder {
	return m.recorder
}

// QueryDynamicBalance mocks base method.
func (m *MockPaymentKeeper) QueryDynamicBalance(ctx types1.Context, addr types1.AccAddress) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryDynamicBalance", ctx, addr)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryDynamicBalance indicates an expected call of QueryDynamicBalance.
func (mr *MockPaymentKeeperMockRecorder) QueryDynamicBalance(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryDynamicBalance", reflect.TypeOf((*MockPaymentKeeper)(nil).QueryDynamicBalance), ctx, addr)
}

// Withdraw mocks base method.
func (m *MockPaymentKeeper) Withdraw(ctx types1.Context, fromAddr, toAddr types1.AccAddress, amount math.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", ctx, fromAddr, toAddr, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockPaymentKeeperMockRecorder) Withdraw(ctx, fromAddr, toAddr, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockPaymentKeeper)(nil).Withdraw), ctx, fromAddr, toAddr, amount)
}

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx context.Context, addr types1.AccAddress) types1.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types1.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx types1.Context, addr types1.AccAddress) types1.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types1.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}
