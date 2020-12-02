// Copyright 2019 whizkid77, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	. "github.com/whizkid77/check"
	"go.uber.org/zap"
)

var _ = Suite(&testLogSuite{})

type testLogSuite struct{}

func (t *testLogSuite) TestExport(c *C) {
	conf := &Config{Level: "debug", File: FileLogConfig{}, DisableTimestamp: true}
	lg := newZapTestLogger(conf, c)
	ReplaceGlobals(lg.Logger, nil)
	Info("Testing")
	Debug("Testing")
	Warn("Testing")
	Error("Testing")
	lg.AssertContains("log_test.go:")

	lg = newZapTestLogger(conf, c)
	ReplaceGlobals(lg.Logger, nil)
	logger := With(zap.String("name", "tester"), zap.Int64("age", 42))
	logger.Info("hello")
	logger.Debug("world")
	lg.AssertContains(`name=tester`)
	lg.AssertContains(`age=42`)
}
