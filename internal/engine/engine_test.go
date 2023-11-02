package engine_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/base-org/pessimism/internal/core"
	"github.com/base-org/pessimism/internal/engine"
	"github.com/base-org/pessimism/internal/engine/heuristic"
	"github.com/base-org/pessimism/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testSuite struct {
	ctrl          *gomock.Controller
	re            engine.RiskEngine
	mockHeuristic *mocks.MockHeuristic
}

func createTestSuite(t *testing.T) *testSuite {
	ctrl := gomock.NewController(t)

	return &testSuite{
		ctrl:          ctrl,
		re:            engine.NewHardCodedEngine(make(chan core.Alert)),
		mockHeuristic: mocks.NewMockHeuristic(ctrl),
	}
}

func testErr() error {
	return fmt.Errorf("test error")
}

func Test_HardCodedEngine(t *testing.T) {
	var tests = []struct {
		name string
		test func(t *testing.T, ts *testSuite)
	}{
		{
			name: "Activation Failure From Error",
			test: func(t *testing.T, ts *testSuite) {
				td := core.TransitData{}

				ts.mockHeuristic.EXPECT().Assess(td).
					Return(heuristic.NoActivations(), testErr()).Times(1)

				ts.mockHeuristic.EXPECT().SUUID().
					Return(core.NilSUUID()).Times(2)

				as := ts.re.Execute(context.Background(), td, ts.mockHeuristic)
				assert.False(t, as.Activated())

			}},
		{
			name: "Successful Activation",
			test: func(t *testing.T, ts *testSuite) {
				td := core.TransitData{}

				expectedOut := heuristic.NewActivationSet().Add(
					&heuristic.Activation{
						Message: "20 inch blade on the Impala",
					})

				ts.mockHeuristic.EXPECT().Assess(td).
					Return(expectedOut, nil).Times(1)

				ts.mockHeuristic.EXPECT().SUUID().
					Return(core.NilSUUID()).Times(1)

				as := ts.re.Execute(context.Background(), td, ts.mockHeuristic)
				assert.NotNil(t, as)
				assert.True(t, as.Activated())
				assert.Equal(t, expectedOut, as)
			}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d-%s", i, test.name), func(t *testing.T) {
			ts := createTestSuite(t)
			test.test(t, ts)
		})
	}
}
