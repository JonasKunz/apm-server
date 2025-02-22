// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
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

package modelprocessor_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/apm-server/internal/model"
	"github.com/elastic/apm-server/internal/model/modelprocessor"
)

func TestSetErrorMessage(t *testing.T) {
	tests := []struct {
		input   model.Error
		message string
	}{{
		input:   model.Error{},
		message: "",
	}, {
		input:   model.Error{Log: &model.ErrorLog{Message: "log_message"}},
		message: "log_message",
	}, {
		input:   model.Error{Exception: &model.Exception{Message: "exception_message"}},
		message: "exception_message",
	}, {
		input: model.Error{
			Log:       &model.ErrorLog{},
			Exception: &model.Exception{Message: "exception_message"},
		},
		message: "exception_message",
	}, {
		input: model.Error{
			Log:       &model.ErrorLog{Message: "log_message"},
			Exception: &model.Exception{Message: "exception_message"},
		},
		message: "log_message",
	}}

	for _, test := range tests {
		batch := model.Batch{{Error: &test.input}}
		processor := modelprocessor.SetErrorMessage{}
		err := processor.ProcessBatch(context.Background(), &batch)
		assert.NoError(t, err)
		assert.Equal(t, test.message, batch[0].Message)
	}

}
