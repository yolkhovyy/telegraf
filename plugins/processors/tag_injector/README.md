# Tag Injector Processor Plugin

Use the `tag_injector` processor to einject extar tags based on the existsing tags.

This can be useful when in IoT scenarios:
- Device replacement
- Metric enrichment with metadata independent of physical device properties (such as mac addresses)
- Metric aggregation based on the injected tags, e.g. inject a location tag which might group multiple metrics

### Configuration

```toml
[[processors.tag_injector]]
  order = 2
  driver_name = "mysql"
  data_source_name = "user:password@/tag_injector?charset=utf8mb4&collation=utf8mb4_unicode_ci"
```

### Example


### Database

#### Volume create

```
docker volume create --name maria-db -o type=none -o device=/home/yo/tmp/tag_injector/db -o o=bind
```
