// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.18.0"
)

type metricK8sDaemonsetCurrentScheduledNodes struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.daemonset.current_scheduled_nodes metric with initial data.
func (m *metricK8sDaemonsetCurrentScheduledNodes) init() {
	m.data.SetName("k8s.daemonset.current_scheduled_nodes")
	m.data.SetDescription("Number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sDaemonsetCurrentScheduledNodes) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sDaemonsetCurrentScheduledNodes) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sDaemonsetCurrentScheduledNodes) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sDaemonsetCurrentScheduledNodes(cfg MetricConfig) metricK8sDaemonsetCurrentScheduledNodes {
	m := metricK8sDaemonsetCurrentScheduledNodes{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sDaemonsetDesiredScheduledNodes struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.daemonset.desired_scheduled_nodes metric with initial data.
func (m *metricK8sDaemonsetDesiredScheduledNodes) init() {
	m.data.SetName("k8s.daemonset.desired_scheduled_nodes")
	m.data.SetDescription("Number of nodes that should be running the daemon pod (including nodes currently running the daemon pod)")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sDaemonsetDesiredScheduledNodes) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sDaemonsetDesiredScheduledNodes) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sDaemonsetDesiredScheduledNodes) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sDaemonsetDesiredScheduledNodes(cfg MetricConfig) metricK8sDaemonsetDesiredScheduledNodes {
	m := metricK8sDaemonsetDesiredScheduledNodes{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sDaemonsetMisscheduledNodes struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.daemonset.misscheduled_nodes metric with initial data.
func (m *metricK8sDaemonsetMisscheduledNodes) init() {
	m.data.SetName("k8s.daemonset.misscheduled_nodes")
	m.data.SetDescription("Number of nodes that are running the daemon pod, but are not supposed to run the daemon pod")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sDaemonsetMisscheduledNodes) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sDaemonsetMisscheduledNodes) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sDaemonsetMisscheduledNodes) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sDaemonsetMisscheduledNodes(cfg MetricConfig) metricK8sDaemonsetMisscheduledNodes {
	m := metricK8sDaemonsetMisscheduledNodes{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sDaemonsetReadyNodes struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.daemonset.ready_nodes metric with initial data.
func (m *metricK8sDaemonsetReadyNodes) init() {
	m.data.SetName("k8s.daemonset.ready_nodes")
	m.data.SetDescription("Number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sDaemonsetReadyNodes) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sDaemonsetReadyNodes) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sDaemonsetReadyNodes) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sDaemonsetReadyNodes(cfg MetricConfig) metricK8sDaemonsetReadyNodes {
	m := metricK8sDaemonsetReadyNodes{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	startTime                               pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                         int                 // maximum observed number of metrics per resource.
	metricsBuffer                           pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                               component.BuildInfo // contains version information
	metricK8sDaemonsetCurrentScheduledNodes metricK8sDaemonsetCurrentScheduledNodes
	metricK8sDaemonsetDesiredScheduledNodes metricK8sDaemonsetDesiredScheduledNodes
	metricK8sDaemonsetMisscheduledNodes     metricK8sDaemonsetMisscheduledNodes
	metricK8sDaemonsetReadyNodes            metricK8sDaemonsetReadyNodes
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                               pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                           pmetric.NewMetrics(),
		buildInfo:                               settings.BuildInfo,
		metricK8sDaemonsetCurrentScheduledNodes: newMetricK8sDaemonsetCurrentScheduledNodes(mbc.Metrics.K8sDaemonsetCurrentScheduledNodes),
		metricK8sDaemonsetDesiredScheduledNodes: newMetricK8sDaemonsetDesiredScheduledNodes(mbc.Metrics.K8sDaemonsetDesiredScheduledNodes),
		metricK8sDaemonsetMisscheduledNodes:     newMetricK8sDaemonsetMisscheduledNodes(mbc.Metrics.K8sDaemonsetMisscheduledNodes),
		metricK8sDaemonsetReadyNodes:            newMetricK8sDaemonsetReadyNodes(mbc.Metrics.K8sDaemonsetReadyNodes),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/k8sclusterreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricK8sDaemonsetCurrentScheduledNodes.emit(ils.Metrics())
	mb.metricK8sDaemonsetDesiredScheduledNodes.emit(ils.Metrics())
	mb.metricK8sDaemonsetMisscheduledNodes.emit(ils.Metrics())
	mb.metricK8sDaemonsetReadyNodes.emit(ils.Metrics())

	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordK8sDaemonsetCurrentScheduledNodesDataPoint adds a data point to k8s.daemonset.current_scheduled_nodes metric.
func (mb *MetricsBuilder) RecordK8sDaemonsetCurrentScheduledNodesDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sDaemonsetCurrentScheduledNodes.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sDaemonsetDesiredScheduledNodesDataPoint adds a data point to k8s.daemonset.desired_scheduled_nodes metric.
func (mb *MetricsBuilder) RecordK8sDaemonsetDesiredScheduledNodesDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sDaemonsetDesiredScheduledNodes.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sDaemonsetMisscheduledNodesDataPoint adds a data point to k8s.daemonset.misscheduled_nodes metric.
func (mb *MetricsBuilder) RecordK8sDaemonsetMisscheduledNodesDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sDaemonsetMisscheduledNodes.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sDaemonsetReadyNodesDataPoint adds a data point to k8s.daemonset.ready_nodes metric.
func (mb *MetricsBuilder) RecordK8sDaemonsetReadyNodesDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sDaemonsetReadyNodes.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
