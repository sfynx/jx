// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/jenkins-x/jx/pkg/auth (interfaces: ConfigSaver)

package auth_test

import (
	"reflect"
	"time"

	auth "github.com/jenkins-x/jx/pkg/auth"
	pegomock "github.com/petergtz/pegomock"
)

type MockConfigSaver struct {
	fail func(message string, callerSkip ...int)
}

func NewMockConfigSaver(options ...pegomock.Option) *MockConfigSaver {
	mock := &MockConfigSaver{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockConfigSaver) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockConfigSaver) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockConfigSaver) LoadConfig() (*auth.AuthConfig, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConfigSaver().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("LoadConfig", params, []reflect.Type{reflect.TypeOf((**auth.AuthConfig)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *auth.AuthConfig
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*auth.AuthConfig)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockConfigSaver) SaveConfig(_param0 *auth.AuthConfig) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConfigSaver().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("SaveConfig", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockConfigSaver) VerifyWasCalledOnce() *VerifierMockConfigSaver {
	return &VerifierMockConfigSaver{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockConfigSaver) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockConfigSaver {
	return &VerifierMockConfigSaver{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockConfigSaver) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockConfigSaver {
	return &VerifierMockConfigSaver{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockConfigSaver) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockConfigSaver {
	return &VerifierMockConfigSaver{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockConfigSaver struct {
	mock                   *MockConfigSaver
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockConfigSaver) LoadConfig() *MockConfigSaver_LoadConfig_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "LoadConfig", params, verifier.timeout)
	return &MockConfigSaver_LoadConfig_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConfigSaver_LoadConfig_OngoingVerification struct {
	mock              *MockConfigSaver
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConfigSaver_LoadConfig_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConfigSaver_LoadConfig_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockConfigSaver) SaveConfig(_param0 *auth.AuthConfig) *MockConfigSaver_SaveConfig_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SaveConfig", params, verifier.timeout)
	return &MockConfigSaver_SaveConfig_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConfigSaver_SaveConfig_OngoingVerification struct {
	mock              *MockConfigSaver
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConfigSaver_SaveConfig_OngoingVerification) GetCapturedArguments() *auth.AuthConfig {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockConfigSaver_SaveConfig_OngoingVerification) GetAllCapturedArguments() (_param0 []*auth.AuthConfig) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*auth.AuthConfig, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*auth.AuthConfig)
		}
	}
	return
}
