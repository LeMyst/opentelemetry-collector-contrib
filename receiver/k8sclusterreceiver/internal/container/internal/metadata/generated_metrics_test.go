// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testConfigCollection int

const (
	testSetDefault testConfigCollection = iota
	testSetAll
	testSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		configSet testConfigCollection
	}{
		{
			name:      "default",
			configSet: testSetDefault,
		},
		{
			name:      "all_set",
			configSet: testSetAll,
		},
		{
			name:      "none_set",
			configSet: testSetNone,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sContainerCPULimitDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sContainerCPURequestDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sContainerEphemeralstorageLimitDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sContainerEphemeralstorageRequestDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sContainerMemoryLimitDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sContainerMemoryRequestDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sContainerReadyDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sContainerRestartsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sContainerStorageLimitDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sContainerStorageRequestDataPoint(ts, 1)

			metrics := mb.Emit(WithContainerID("container.id-val"), WithContainerImageName("container.image.name-val"), WithContainerImageTag("container.image.tag-val"), WithK8sContainerName("k8s.container.name-val"), WithK8sNamespaceName("k8s.namespace.name-val"), WithK8sNodeName("k8s.node.name-val"), WithK8sPodName("k8s.pod.name-val"), WithK8sPodUID("k8s.pod.uid-val"), WithOpencensusResourcetype("opencensus.resourcetype-val"))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			attrVal, ok := rm.Resource().Attributes().Get("container.id")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.ContainerID.Enabled, ok)
			if mb.resourceAttributesConfig.ContainerID.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "container.id-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("container.image.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.ContainerImageName.Enabled, ok)
			if mb.resourceAttributesConfig.ContainerImageName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "container.image.name-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("container.image.tag")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.ContainerImageTag.Enabled, ok)
			if mb.resourceAttributesConfig.ContainerImageTag.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "container.image.tag-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("k8s.container.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.K8sContainerName.Enabled, ok)
			if mb.resourceAttributesConfig.K8sContainerName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "k8s.container.name-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("k8s.namespace.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.K8sNamespaceName.Enabled, ok)
			if mb.resourceAttributesConfig.K8sNamespaceName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "k8s.namespace.name-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("k8s.node.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.K8sNodeName.Enabled, ok)
			if mb.resourceAttributesConfig.K8sNodeName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "k8s.node.name-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("k8s.pod.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.K8sPodName.Enabled, ok)
			if mb.resourceAttributesConfig.K8sPodName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "k8s.pod.name-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("k8s.pod.uid")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.K8sPodUID.Enabled, ok)
			if mb.resourceAttributesConfig.K8sPodUID.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "k8s.pod.uid-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("opencensus.resourcetype")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.OpencensusResourcetype.Enabled, ok)
			if mb.resourceAttributesConfig.OpencensusResourcetype.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "opencensus.resourcetype-val", attrVal.Str())
			}
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 9)

			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "k8s.container.cpu_limit":
					assert.False(t, validatedMetrics["k8s.container.cpu_limit"], "Found a duplicate in the metrics slice: k8s.container.cpu_limit")
					validatedMetrics["k8s.container.cpu_limit"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Maximum resource limit set for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details", ms.At(i).Description())
					assert.Equal(t, "{cpu}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
				case "k8s.container.cpu_request":
					assert.False(t, validatedMetrics["k8s.container.cpu_request"], "Found a duplicate in the metrics slice: k8s.container.cpu_request")
					validatedMetrics["k8s.container.cpu_request"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Resource requested for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details", ms.At(i).Description())
					assert.Equal(t, "{cpu}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
				case "k8s.container.ephemeralstorage_limit":
					assert.False(t, validatedMetrics["k8s.container.ephemeralstorage_limit"], "Found a duplicate in the metrics slice: k8s.container.ephemeralstorage_limit")
					validatedMetrics["k8s.container.ephemeralstorage_limit"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Maximum resource limit set for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.container.ephemeralstorage_request":
					assert.False(t, validatedMetrics["k8s.container.ephemeralstorage_request"], "Found a duplicate in the metrics slice: k8s.container.ephemeralstorage_request")
					validatedMetrics["k8s.container.ephemeralstorage_request"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Resource requested for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.container.memory_limit":
					assert.False(t, validatedMetrics["k8s.container.memory_limit"], "Found a duplicate in the metrics slice: k8s.container.memory_limit")
					validatedMetrics["k8s.container.memory_limit"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Maximum resource limit set for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.container.memory_request":
					assert.False(t, validatedMetrics["k8s.container.memory_request"], "Found a duplicate in the metrics slice: k8s.container.memory_request")
					validatedMetrics["k8s.container.memory_request"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Resource requested for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.container.ready":
					assert.False(t, validatedMetrics["k8s.container.ready"], "Found a duplicate in the metrics slice: k8s.container.ready")
					validatedMetrics["k8s.container.ready"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Whether a container has passed its readiness probe (0 for no, 1 for yes)", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.container.restarts":
					assert.False(t, validatedMetrics["k8s.container.restarts"], "Found a duplicate in the metrics slice: k8s.container.restarts")
					validatedMetrics["k8s.container.restarts"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "How many times the container has restarted in the recent past. This value is pulled directly from the K8s API and the value can go indefinitely high and be reset to 0 at any time depending on how your kubelet is configured to prune dead containers. It is best to not depend too much on the exact value but rather look at it as either == 0, in which case you can conclude there were no restarts in the recent past, or > 0, in which case you can conclude there were restarts in the recent past, and not try and analyze the value beyond that.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.container.storage_limit":
					assert.False(t, validatedMetrics["k8s.container.storage_limit"], "Found a duplicate in the metrics slice: k8s.container.storage_limit")
					validatedMetrics["k8s.container.storage_limit"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Maximum resource limit set for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.container.storage_request":
					assert.False(t, validatedMetrics["k8s.container.storage_request"], "Found a duplicate in the metrics slice: k8s.container.storage_request")
					validatedMetrics["k8s.container.storage_request"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Resource requested for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				}
			}
		})
	}
}
