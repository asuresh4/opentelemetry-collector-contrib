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

	sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"
)

// datapointFilter can be used to filter out datapoints
type datapointFilter interface {
	// Matches takes a datapoint and returns whether it is matched by the
	// filter
	Matches(*sfxpb.DataPoint) bool
}

type overridableDatapointFilter struct {
	dimFilter    stringMapFilter
	metricFilter stringFilter
}

// newOverridable returns a new overridable filter with the given configuration
func newOverridable(metricNames []string, dimensions map[string][]string) (datapointFilter, error) {
	var dimFilter stringMapFilter
	if len(dimensions) > 0 {
		var err error
		dimFilter, err = newStringMapFilter(dimensions)
		if err != nil {
			return nil, err
		}
	}

	var metricFilter stringFilter
	if len(metricNames) > 0 {
		var err error
		metricFilter, err = newOverridableStringFilter(metricNames)
		if err != nil {
			return nil, err
		}
	}

	if metricFilter == nil && dimFilter == nil {
		return nil, errors.New("metric filter must have at least one metric or dimension defined on it")
	}

	return &overridableDatapointFilter{
		metricFilter: metricFilter,
		dimFilter:    dimFilter,
	}, nil
}

// Matches tests a datapoint to see whether it is excluded by this 
func (f *overridableDatapointFilter) Matches(dp *sfxpb.DataPoint) bool {
	return (f.metricFilter == nil || f.metricFilter.Matches(dp.Metric)) &&
		(f.dimFilter == nil || f.dimFilter.Matches(getDimensionsMap(dp.Dimensions)))
}

func getDimensionsMap(dimensions []*sfxpb.Dimension) map[string]string {
	out := make(map[string]string, len(dimensions))
	for _, dim := range dimensions {
		out[dim.Key] = dim.Value
	}
	return out
}
