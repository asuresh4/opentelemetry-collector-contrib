// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dpfilters

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestStringMapFilter(t *testing.T) {
	for _, tc := range []struct {
		filter      map[string][]string
		input       map[string]string
		shouldMatch bool
		shouldError bool
	}{
		{
			filter: map[string][]string{},
			input:  map[string]string{},
			// Empty map never matches anything, even blank filter
			shouldMatch: false,
		},
		{
			filter: map[string][]string{
				"app": {"test"},
			},
			input:       map[string]string{},
			shouldMatch: false,
		},
		{
			filter: map[string][]string{
				"app?": {"test"},
			},
			input:       map[string]string{},
			shouldMatch: true,
		},
		{
			filter: map[string][]string{
				"app?": {"test"},
			},
			input: map[string]string{
				"version": "latest",
			},
			shouldMatch: true,
		},
		{
			filter: map[string][]string{
				"app":     {"test"},
				"version": {"*"},
			},
			input: map[string]string{
				"app": "test",
			},
			shouldMatch: false,
		},
		{
			filter: map[string][]string{
				"app": {"test"},
			},
			input: map[string]string{
				"app":     "test",
				"version": "2.0",
			},
			shouldMatch: true,
		},
		{
			filter: map[string][]string{
				"version": {`/\d+\.\d+/`},
			},
			input: map[string]string{
				"app":     "test",
				"version": "2.0",
			},
			shouldMatch: true,
		},
		{
			filter: map[string][]string{
				"version": {`/\d+\.\d+/`},
			},
			input: map[string]string{
				"app":     "test",
				"version": "bad",
			},
			shouldMatch: false,
		},
	} {
		f, err := newStringMapFilter(tc.filter)
		if tc.shouldError {
			assert.NotNil(t, err, spew.Sdump(tc))
		} else {
			assert.Nil(t, err, spew.Sdump(tc))
		}

		assert.Equal(t, tc.shouldMatch, f.Matches(tc.input), "%s\n%s", spew.Sdump(tc), spew.Sdump(f))
	}
}
