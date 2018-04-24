package cargo

import (
	"testing"
	"reflect"
)

var (
	testCaseRateLimit       = 777
	testCaseWorkerBatchSize = 666
)

var optTestCases = []struct {
	message  string
	opt      Option
	expected *poolCfg
}{
	{
		"RateLimit failed: expected: `%+v`, given: `%+v`",
		RateLimit(testCaseRateLimit),
		&poolCfg{rateLimit: &testCaseRateLimit},
	},
	{
		"WorkerBatchSize: expected: `%+v`, given: `%+v`",
		WorkerBatchSize(testCaseWorkerBatchSize),
		&poolCfg{batchSize: &testCaseWorkerBatchSize},
	},
	{
		"WithStats: expected: : `%+v`, given: `%+v`",
		WithStats(),
		&poolCfg{collectStats: true},
	},
}

func TestPoolCfg_Set(t *testing.T) {
	for _, tc := range optTestCases {
		cfg := &poolCfg{}
		tc.opt(cfg)

		if !reflect.DeepEqual(cfg, tc.expected) {
			t.Fatalf(tc.message, cfg, tc.expected)
		}
	}
}
