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

package translation

import (
	"testing"

	sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/translation/dpfilters"
)

func TestGetExcludeMetricsRule(t *testing.T) {
	rule := GetExcludeMetricsRule([]dpfilters.MetricFilter{{MetricNames: []string{"m1", "m2"}}})
	require.Equal(t, 2, len(rule.MetricFilters[0].MetricNames))
	fs, err := dpfilters.NewFilterSet(rule.MetricFilters)
	require.NoError(t, err)
	assert.Equal(t, rule.Action, ActionDropMetrics)
	assert.False(t, fs.Matches(&sfxpb.DataPoint{Metric: "m0"}))
	assert.True(t, fs.Matches(&sfxpb.DataPoint{Metric: "m1"}))
}
