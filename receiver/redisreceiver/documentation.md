[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# redis

## Default Metrics

The following metrics are emitted by default. Each of them can be disabled by applying the following configuration:

```yaml
metrics:
  <metric_name>:
    enabled: false
```

### redis.clients.blocked

Number of clients pending on a blocking call

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
|  | Sum | Int | Cumulative | false |

### redis.clients.connected

Number of client connections (excluding connections from replicas)

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
|  | Sum | Int | Cumulative | false |

### redis.clients.max_input_buffer

Biggest input buffer among current client connections

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### redis.clients.max_output_buffer

Longest output list among current client connections

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### redis.commands

Number of commands processed per second

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {ops}/s | Gauge | Int |

### redis.commands.processed

Total number of commands processed by the server

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
|  | Sum | Int | Cumulative | true |

### redis.connections.received

Total number of connections accepted by the server

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
|  | Sum | Int | Cumulative | true |

### redis.connections.rejected

Number of connections rejected because of maxclients limit

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
|  | Sum | Int | Cumulative | true |

### redis.cpu.time

System CPU consumed by the Redis server in seconds since server start

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| s | Sum | Double | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| state | Redis CPU usage state | Str: ``sys``, ``sys_children``, ``sys_main_thread``, ``user``, ``user_children``, ``user_main_thread`` |

### redis.db.avg_ttl

Average keyspace keys TTL

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| ms | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| db | Redis database identifier | Any Str |

### redis.db.expires

Number of keyspace keys with an expiration

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
|  | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| db | Redis database identifier | Any Str |

### redis.db.keys

Number of keyspace keys

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
|  | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| db | Redis database identifier | Any Str |

### redis.keys.evicted

Number of evicted keys due to maxmemory limit

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
|  | Sum | Int | Cumulative | true |

### redis.keys.expired

Total number of key expiration events

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
|  | Sum | Int | Cumulative | true |

### redis.keyspace.hits

Number of successful lookup of keys in the main dictionary

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
|  | Sum | Int | Cumulative | true |

### redis.keyspace.misses

Number of failed lookup of keys in the main dictionary

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
|  | Sum | Int | Cumulative | true |

### redis.latest_fork

Duration of the latest fork operation in microseconds

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| us | Gauge | Int |

### redis.memory.fragmentation_ratio

Ratio between used_memory_rss and used_memory

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
|  | Gauge | Double |

### redis.memory.lua

Number of bytes used by the Lua engine

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### redis.memory.peak

Peak memory consumed by Redis (in bytes)

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### redis.memory.rss

Number of bytes that Redis allocated as seen by the operating system

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### redis.memory.used

Total number of bytes allocated by Redis using its allocator

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### redis.net.input

The total number of bytes read from the network

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| By | Sum | Int | Cumulative | true |

### redis.net.output

The total number of bytes written to the network

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| By | Sum | Int | Cumulative | true |

### redis.rdb.changes_since_last_save

Number of changes since the last dump

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
|  | Sum | Int | Cumulative | false |

### redis.replication.backlog_first_byte_offset

The master offset of the replication backlog buffer

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### redis.replication.offset

The server's current replication offset

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### redis.slaves.connected

Number of connected replicas

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
|  | Sum | Int | Cumulative | false |

### redis.uptime

Number of seconds since Redis server start

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| s | Sum | Int | Cumulative | true |

## Optional Metrics

The following metrics are not emitted by default. Each of them can be enabled by applying the following configuration:

```yaml
metrics:
  <metric_name>:
    enabled: true
```

### redis.cmd.calls

Total number of calls for a command

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
|  | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| cmd | Redis command name | Any Str |

### redis.cmd.usec

Total time for all executions of this command

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| us | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| cmd | Redis command name | Any Str |

### redis.maxmemory

The value of the maxmemory configuration directive

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### redis.role

Redis node's role

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
|  | Sum | Int | Cumulative | false |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| role | Redis node's role | Str: ``replica``, ``primary`` |

## Resource Attributes

| Name | Description | Values | Enabled |
| ---- | ----------- | ------ | ------- |
| redis.version | Redis server's version. | Any Str | true |