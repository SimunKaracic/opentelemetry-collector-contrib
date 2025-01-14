[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# k8s/daemonset

**Parent Component:** k8s_cluster

## Default Metrics

The following metrics are emitted by default. Each of them can be disabled by applying the following configuration:

```yaml
metrics:
  <metric_name>:
    enabled: false
```

### k8s.daemonset.current_scheduled_nodes

Number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| 1 | Gauge | Int |

### k8s.daemonset.desired_scheduled_nodes

Number of nodes that should be running the daemon pod (including nodes currently running the daemon pod)

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| 1 | Gauge | Int |

### k8s.daemonset.misscheduled_nodes

Number of nodes that are running the daemon pod, but are not supposed to run the daemon pod

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| 1 | Gauge | Int |

### k8s.daemonset.ready_nodes

Number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| 1 | Gauge | Int |

## Resource Attributes

| Name | Description | Values | Enabled |
| ---- | ----------- | ------ | ------- |
| k8s.daemonset.name | The k8s daemonset name. | Any Str | true |
| k8s.daemonset.uid | The k8s daemonset uid. | Any Str | true |
| k8s.namespace.name | The k8s namespace name. | Any Str | true |
| opencensus.resourcetype | The OpenCensus resource type. | Any Str | true |
