// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by "stringer -type StepType"; DO NOT EDIT.

package testutils

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StepTypeOutgoingPlcMessage-1]
	_ = x[StepTypeOutgoingPlcBytes-2]
	_ = x[StepTypeIncomingPlcMessage-3]
	_ = x[StepTypeIncomingPlcBytes-4]
	_ = x[StepTypeApiRequest-5]
	_ = x[StepTypeApiResponse-6]
	_ = x[StepTypeDelay-7]
	_ = x[StepTypeTerminate-8]
}

const _StepType_name = "StepTypeOutgoingPlcMessageStepTypeOutgoingPlcBytesStepTypeIncomingPlcMessageStepTypeIncomingPlcBytesStepTypeApiRequestStepTypeApiResponseStepTypeDelayStepTypeTerminate"

var _StepType_index = [...]uint8{0, 26, 50, 76, 100, 118, 137, 150, 167}

func (i StepType) String() string {
	i -= 1
	if i >= StepType(len(_StepType_index)-1) {
		return "StepType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _StepType_name[_StepType_index[i]:_StepType_index[i+1]]
}
