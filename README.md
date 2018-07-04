KSUID
=====

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