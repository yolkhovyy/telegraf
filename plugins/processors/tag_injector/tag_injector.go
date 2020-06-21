package TagInjector

import (
	"log"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/processors"
)

const sampleConfig = `
  ## List of tags to inject
  inject = ["foo", "bar", "baz"]
`

type TagInjector struct {
	Inject     []string `toml:"inject"`
	init       bool
	injectTags map[string]string
}

func (d *TagInjector) SampleConfig() string {
	return sampleConfig
}

func (d *TagInjector) Description() string {
	return "Injects tags."
}

func (d *TagInjector) initOnce() error {
	if d.init {
		return nil
	}
	d.injectTags = make(map[string]string)
	// convert list of tags-to-inject to a map so we can do constant-time lookups
	for _, tag_key := range d.Inject {
		d.injectTags[tag_key] = ""
	}
	d.init = true
	return nil
}

func (d *TagInjector) Apply(in ...telegraf.Metric) []telegraf.Metric {
	err := d.initOnce()
	if err != nil {
		log.Printf("E! [processors.tag_injector] could not create tag_injector processor: %v", err)
		return in
	}
	for _, point := range in {
		tagValue := "dummy"
		for tagKey := range d.injectTags {
			point.AddTag(tagKey, tagValue)
		}
	}

	return in
}

func init() {
	processors.Add("tag_injector", func() telegraf.Processor {
		return &TagInjector{}
	})
}
