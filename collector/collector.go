package collector

type Collector interface {
	Collect() (map[string]interface{}, error)
}

type NullCollector struct {
}

func (c *NullCollector) Collect() (map[string]interface{}, error) {
	//do-nothing
	return nil, nil
}

type CPUCollector struct {
}

func (c *CPUCollector) Collect() (map[string]interface{}, error) {
	// get system cpu information

	return nil, nil
}

type MemoryCollector struct {
}

func (m *MemoryCollector) Collect() (map[string]interface{}, error) {
	// get system memory information

	return nil, nil
}
