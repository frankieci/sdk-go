------------
### ***sdk-go Project***

#### Description
***sdk-go can help collect system information, such as cpu, memory etc.
And it offers API that how to use, it follows***


#### Code Snippet Usage
`Collector factory ceate Collector which you want to know information about system, sample cpu`
```go
    collectorFactory := factory.NewCollectorFactory()
    collectorFactory.SetCollectorType("cpu")
    collector := collectorFactory.CreateCollector()
```

`Collector call collect function, show system information`
```go
    data, err:= collector.collect()
    if err!= nil {
    	return err
    }
    fmt.Println(data)
```  

`if you can know more information about system, expanding collector that 
it can impliment collector interface, after, register it.
smaple disk, DiskCollector impliment collector interfact`
```go
    collectorFactory.RegisterCollector("disk", &DiskCollector{})
```

#### API Server
`It offers API and you try to run api server, support getting cpu/memory information in system.
`

```go
    svr := server.NewSDKServer()
    svr.Run()
```
