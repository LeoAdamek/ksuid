KSUID 
=====
[![CircleCI](https://circleci.com/gh/LeoAdamek/ksuid.svg?style=svg)](https://circleci.com/gh/LeoAdamek/ksuid)

This is a variant of [segmentio/ksuid](https://github.com/segmentio/ksuid) but with customizable partitioning support
pluggable encoding.

Usage
-----

````go
import "github.com/leoadamek/ksuid"

// Use the package-wide StandardGenerator
k := ksuid.Next()

// Create and use an Async generator
g := ksuid.NewAsyncGenerator()
go g.Run()
k = g.Next()

// Change the package-wide PartitionFunc
ksuid.SetPartitioner(ksuid.StringPartitioner("test"))
k = ksuid.Next()

// Change the custom generator's PartitionFunc
g.Partitioner = ksuid.StringPartitioner("asyn")
k = g.Next()
````

Partitioners
-------------

Partitioners can be used to split KSUIDs into partitions. 
This is especially useful in distributed systems so each node can write to a different partition.

Any parameterless function which returns a `uint32` can be used as a partitioner.

The following partitioners are included:

|Name|Description|
|:--:|:----------|
|`ksuid.NilPartitioner`|Sets everything to partition `0`|
|`ksuid.StringPartitioner(s string)`|Sets all partitions to `s` as a BigEndian `uint32`|
|`ksuid.MacPartitioner()()`|Sets the partition based on the first network interface's hardware address|

Bechmarks
---------

    BenchmarkEncodeBinary-8                   10000000               242    ns/op
    BenchmarkDecodeBinary-8                   30000000                44.2  ns/op
    BenchmarkDecodeHex-8                      10000000               132    ns/op
    BenchmarkEncodeHex-8                       5000000               301    ns/op
    BenchmarkStandardGenerator_Next-8         20000000                81.9  ns/op
    BenchmarkAsyncGenerator_Next-8           200000000                 9.27 ns/op
    BenchmarkStringPartitioner-8            2000000000                 1.69 ns/op
    BenchmarkMacPartitioner-8               2000000000                 1.91 ns/op
    

Roadmap
-------

* Improve Encode/Decode speed for Binary and Hex