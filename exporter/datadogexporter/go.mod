module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter

go 1.17

require (
	github.com/DataDog/agent-payload/v5 v5.0.22
	github.com/DataDog/datadog-agent/pkg/quantile v0.36.0-rc.4.0.20220426134022-b55bccd5073f
	github.com/DataDog/datadog-agent/pkg/trace/exportable v0.0.0-20201016145401-4646cf596b02
	github.com/DataDog/sketches-go v1.4.2
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v0.32.1
	github.com/aws/aws-sdk-go v1.44.327
	github.com/cenkalti/backoff/v4 v4.1.3
	github.com/gogo/protobuf v1.3.2
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.53.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders v0.53.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry v0.53.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/collector v0.53.0
	go.opentelemetry.io/collector/pdata v0.53.0
	go.opentelemetry.io/collector/semconv v0.53.0
	go.uber.org/multierr v1.8.0
	go.uber.org/zap v1.21.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.58.1
	gopkg.in/zorkian/go-datadog-api.v2 v2.30.0
)

require (
	cloud.google.com/go/compute v1.23.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/DataDog/datadog-agent/pkg/util/log v0.0.0-20201009092105-58e18918b2db // indirect
	github.com/DataDog/datadog-go v4.8.2+incompatible // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/docker v20.10.17+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/golang/glog v1.1.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf v1.4.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/shirou/gopsutil v2.20.9+incompatible // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/stretchr/objx v0.5.1 // indirect
	github.com/tinylib/msgp v1.1.8 // indirect
	github.com/zorkian/go-datadog-api v2.30.0+incompatible // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/otel v1.20.0 // indirect
	go.opentelemetry.io/otel/metric v1.20.0 // indirect
	go.opentelemetry.io/otel/trace v1.20.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.12.1-0.20230815132531-74c255bcf846 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/grpc v1.57.1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders => ../../internal/metadataproviders

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry => ../../pkg/resourcetotelemetry

// see https://github.com/go-chi/chi/issues/713
// see https://github.com/DataDog/dd-trace-go/issues/1220
replace github.com/go-chi/chi/v4 => github.com/go-chi/chi v4.0.0+incompatible
