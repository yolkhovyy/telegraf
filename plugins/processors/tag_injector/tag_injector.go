package taginjector

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/processors"
)

const sampleConfig = `
  ## List of tags to inject
  inject = ["foo", "bar", "baz"]
`

// TagInjector is ...
type TagInjector struct {
	DriverName     string `toml:"driver_name"`
	DataSourceName string `toml:"data_source_name"`
	init           bool
}

// SampleConfig is ...
func (d *TagInjector) SampleConfig() string {
	return sampleConfig
}

// Description is ...
func (d *TagInjector) Description() string {
	return "Injects tags."
}

func (d *TagInjector) initOnce() error {
	if d.init {
		return nil
	}

	d.init = true
	return nil
}

// Apply is ...
func (d *TagInjector) Apply(in ...telegraf.Metric) []telegraf.Metric {

	err := d.initOnce()
	if err != nil {
		log.Printf("E! [processors.tag_injector] could not create tag_injector processor: %v", err)
		return in
	}

	db, err := sql.Open(d.DriverName, d.DataSourceName)
	if err != nil {
		log.Printf("E! [processors.tag_injector] could not open database: %v", err)
		return in
	}
	defer db.Close()

	for _, point := range in {
		for inTagName, inTagValue := range point.Tags() {
			rows, err := db.Query("select out_tags.tag_name, out_tags.tag_value from in_tags left join out_tags on in_tags.id = out_tags.in_tag_id where in_tags.tag_name=? and in_tags.tag_value=?", inTagName, inTagValue)
			if err != nil {
				log.Printf("E! [processors.tag_injector] database query failed: %v", err)
			} else {
				var (
					outTagName  string
					outTagValue string
				)
				for rows.Next() {
					rows.Scan(&outTagName, &outTagValue)
					point.AddTag(outTagName, outTagValue)
				}
			}
		}
	}

	return in
}

func init() {
	processors.Add("tag_injector", func() telegraf.Processor {
		return &TagInjector{}
	})
}
