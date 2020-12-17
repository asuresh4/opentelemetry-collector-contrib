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

func TestBasicStringFilter(t *testing.T) {
	for _, tc := range []struct {
		filter      []string
		input       string
		shouldMatch bool
		shouldError bool
	}{
		{
			filter:      []string{},
			input:       "process_",
			shouldMatch: false,
		},
		{
			filter: []string{
				"!app",
			},
			input:       "app",
			shouldMatch: false,
		},
		{
			filter: []string{
				"!app",
			},
			input:       "something",
			shouldMatch: true,
		},
		{
			filter: []string{
				"other",
				"!app",
			},
			input:       "something",
			shouldMatch: true,
		},
		{
			filter: []string{
				"other",
				"!app",
			},
			input:       "app",
			shouldMatch: false,
		},
		{
			filter: []string{
				"/^process_/",
				"/^node_/",
			},
			input:       "process_",
			shouldMatch: true,
		},
		{
			filter: []string{
				"!/^process_/",
			},
			input:       "process_",
			shouldMatch: false,
		},
		{
			filter: []string{
				"!app",
				"!/^process_/",
			},
			input:       "other",
			shouldMatch: true,
		},
		{
			filter: []string{
				"!other",
				"!/^process_/",
			},
			input: "other",
			// Since "other" is explicitly excluded, it should not ever match.
			shouldMatch: false,
		},
		{
			filter: []string{
				"app",
				"!/^process_/",
			},
			input:       "other",
			shouldMatch: true,
		},
		{
			filter: []string{
				"asdfdfasdf",
				"!/^node_/",
			},
			input:       "process_",
			shouldMatch: true,
		},
		{
			filter: []string{
				"asdfdfasdf",
				"/^node_/",
			},
			input:       "process_",
			shouldMatch: false,
		},
	} {
		f, err := newBasicStringFilter(tc.filter)
		if tc.shouldError {
			assert.NotNil(t, err, spew.Sdump(tc))
		} else {
			assert.Nil(t, err, spew.Sdump(tc))
		}

		assert.Equal(t, tc.shouldMatch, f.Matches(tc.input), "%s\n%s", spew.Sdump(tc), spew.Sdump(f))
	}
}

func TestOverridableStringFilter(t *testing.T) {
	for _, tc := range []struct {
		filter      []string
		inputs      []string
		shouldMatch []bool
		shouldError bool
	}{
		{
			filter:      []string{},
			inputs:      []string{"process_", "", "asdf"},
			shouldMatch: []bool{false, false, false},
		},
		{
			filter: []string{
				"*",
			},
			inputs:      []string{"app", "asdf", "", "*"},
			shouldMatch: []bool{true, true, true, true},
		},
		{
			filter: []string{
				"!app",
			},
			inputs:      []string{"app", "other"},
			shouldMatch: []bool{false, false},
		},
		{
			filter: []string{
				// A positive and negative literal match cancel each other out
				// and don't match.
				"app",
				"!app",
			},
			inputs:      []string{"app", "other"},
			shouldMatch: []bool{false, false},
		},
		{
			filter: []string{
				"other",
				"!app",
			},
			inputs:      []string{"other", "something", "app"},
			shouldMatch: []bool{true, false, false},
		},
		{
			filter: []string{
				"/^process_/",
				"/^node_/",
			},
			inputs:      []string{"process_", "node_", "process_asdf", "other"},
			shouldMatch: []bool{true, true, true, false},
		},
		{
			filter: []string{
				"!/^process_/",
			},
			inputs:      []string{"process_", "other"},
			shouldMatch: []bool{false, false},
		},
		{
			filter: []string{
				"app",
				"!/^process_/",
				"process_",
			},
			inputs:      []string{"other", "app", "process_cpu", "process_"},
			shouldMatch: []bool{false, true, false, false},
		},
		{
			filter: []string{
				"asdfdfasdf",
				"/^node_/",
			},
			inputs:      []string{"node_test"},
			shouldMatch: []bool{true},
		},
		{
			filter: []string{
				"process_*",
				"!process_cpu",
			},
			inputs:      []string{"process_mem", "process_cpu", "asdf"},
			shouldMatch: []bool{true, false, false},
		},
		{
			filter: []string{
				"*",
				"!process_cpu",
			},
			inputs:      []string{"process_mem", "process_cpu", "asdf"},
			shouldMatch: []bool{true, false, true},
		},
		{
			filter: []string{
				"metric_?",
				"!metric_a",
				"!metric_b",
				"random",
			},
			inputs:      []string{"metric_a", "metric_b", "metric_c", "asdf", "random"},
			shouldMatch: []bool{false, false, true, false, true},
		},
		{
			filter: []string{
				"!process_cpu",
				// Order doesn't matter
				"*",
			},
			inputs:      []string{"process_mem", "process_cpu", "asdf"},
			shouldMatch: []bool{true, false, true},
		},
		{
			filter: []string{
				"/a.*/",
				"!/.*z/",
				"b",
				// Static match should not override the negated regex above
				"alz",
			},
			inputs:      []string{"", "asdf", "asdz", "b", "wrong", "alz"},
			shouldMatch: []bool{false, true, false, true, false, false},
		},
	} {
		f, err := newOverridableStringFilter(tc.filter)
		if tc.shouldError {
			assert.NotNil(t, err, spew.Sdump(tc))
		} else {
			assert.Nil(t, err, spew.Sdump(tc))
		}
		for i := range tc.inputs {
			assert.Equal(t, tc.shouldMatch[i], f.Matches(tc.inputs[i]), "input[%d] of %s\n%s", i, spew.Sdump(tc), spew.Sdump(f))
		}
	}
}
