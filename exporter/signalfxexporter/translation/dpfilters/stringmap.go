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
	"errors"
	"strings"
)

// stringMapFilter matches against the values of a map[string]string.
type stringMapFilter interface {
	Matches(map[string]string) bool
}

// newStringMapFilter returns a filter that matches against the provided map.
// All key/value pairs must match the spec given in m for a map to be
// considered a match.
func newStringMapFilter(m map[string][]string) (stringMapFilter, error) {
	filterMap := map[string]*overridableStringFilter{}
	okMissing := map[string]bool{}
	for k := range m {
		if len(m[k]) == 0 {
			return nil, errors.New("string map value in filter cannot be empty")
		}

		realKey := strings.TrimSuffix(k, "?")

		var err error
		filterMap[realKey], err = newOverridableStringFilter(m[k])
		if err != nil {
			return nil, err
		}

		if len(realKey) != len(k) {
			okMissing[realKey] = true
		}
	}

	return &fullStringMapFilter{
		filterMap: filterMap,
		okMissing: okMissing,
	}, nil
}

// Each key/value pair must match the filter for the whole map to match.
type fullStringMapFilter struct {
	filterMap map[string]*overridableStringFilter
	okMissing map[string]bool
}

func (f *fullStringMapFilter) Matches(m map[string]string) bool {
	// Empty map input never matches
	if len(m) == 0 && len(f.okMissing) == 0 {
		return false
	}

	for k, filter := range f.filterMap {
		if v, ok := m[k]; ok {
			if !filter.Matches(v) {
				return false
			}
		} else {
			return f.okMissing[k]
		}
	}
	return true
}
