package taginjector

import (
	"testing"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
	"github.com/stretchr/testify/assert"
)

const measuremement = "metrics"

var testTags = []map[string]string{
	{"device_id": "5a2d34fffe393251", "idevice_id": "f4f95c8c-1aa9-4be7-80d4-4761c9d8f486", "location_id": "greenhouse air"},
	{"device_id": "c67c8dfffe65d020", "idevice_id": "6efcce05-ffc5-48be-a4c9-fa25b159436b", "location_id": "greenhouse cucumbers"},
	{"device_id": "c67c8dfffe65cb63", "idevice_id": "a878b91f-bda7-4e10-ba06-d51fb5651bae", "location_id": "greenhouse tomatoes"},
	{"device_id": "5a2d34fffe3682ed", "idevice_id": "5bd540e6-7195-4dcb-aff8-2d23e457c9bd", "location_id": "ground floor"},
	{"device_id": "5a2d34fffe39783f", "idevice_id": "2440e974-cb2a-40c7-8117-05468cfb3eda", "location_id": "first floor"},
}

func buildMetric(name string, tags map[string]string, fields map[string]interface{}, metricTime time.Time) telegraf.Metric {
	if tags == nil {
		tags = map[string]string{}
	}
	if fields == nil {
		fields = map[string]interface{}{}
	}
	m, _ := metric.New(name, tags, fields, metricTime)
	return m
}

func TestJnjector(t *testing.T) {
	currentTime := time.Now()

	tagInjector := TagInjector{
		DriverName:     "mysql",
		DataSourceName: "user:password@/tag_injector?&charset=utf8mb4&collation=utf8mb4_unicode_ci",
	}

	for _, tags := range testTags {
		metric := buildMetric(measuremement, tags, nil, currentTime)
		resultMetric := tagInjector.Apply(metric)
		assert.Equal(t, tags, resultMetric[0].Tags())
	}

}
