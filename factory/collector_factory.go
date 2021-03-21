package factory

import (
	"reflect"

	"gitlab.com/sdk-go/collector"
)

type CollectorFactory struct {
	CollectorType string
}

func NewCollectorFactory() *CollectorFactory {
	return &CollectorFactory{}
}

func (f *CollectorFactory) SetCollectorType(collectorType string) {
	f.CollectorType = collectorType
}

var collectors = make(map[string]reflect.Type)

func init() {
	collectors["cpu"] = reflect.TypeOf(collector.CPUCollector{})
	collectors["memory"] = reflect.TypeOf(collector.MemoryCollector{})
}

func (f *CollectorFactory) RegisterCollector(key string, c collector.Collector) {
	collectors[key] = reflect.TypeOf(c)
}

func (f *CollectorFactory) CreateCollector() collector.Collector {
	c, ok := collectors[f.CollectorType]
	if !ok {
		return &collector.NullCollector{}
	}
	return reflect.New(c).Interface().(collector.Collector)
}
