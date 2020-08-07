package metrics

import (
	marketplacev1alpha1 "github.com/redhat-marketplace/redhat-marketplace-operator/pkg/apis/marketplace/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	kbsm "k8s.io/kube-state-metrics/pkg/metric"
)

var (
	descPersistentVolumeClaimLabelsDefaultLabels = []string{"namespace", "persistentvolumeclaim"}
)

var pvcMetricsFamilies = []FamilyGenerator{
	{
		FamilyGenerator: kbsm.FamilyGenerator{
			Name: "meterdef_persistentvolumeclaim_info",
			Type: kbsm.Gauge,
			Help: "Metering info for persistentvolumeclaim",
		},
		GenerateMeterFunc: wrapPersistentVolumeClaimFunc(func(pvc *corev1.PersistentVolumeClaim, meterDefinitions []*marketplacev1alpha1.MeterDefinition) *kbsm.Family {
			metrics := []*kbsm.Metric{}

			phase := pvc.Status.Phase

			metrics = append(metrics, &kbsm.Metric{
				LabelKeys:   []string{"phase"},
				LabelValues: []string{string(phase)},
				Value:       1,
			})

			return &kbsm.Family{
				Metrics: metrics,
			}
		}),
	},
}

// wrapPersistentVolumeClaimFunc is a helper function for generating pvc-based metrics
func wrapPersistentVolumeClaimFunc(f func(*v1.PersistentVolumeClaim, []*marketplacev1alpha1.MeterDefinition) *kbsm.Family) func(obj interface{}, meterDefinitions []*marketplacev1alpha1.MeterDefinition) *kbsm.Family {
	return func(obj interface{}, meterDefinitions []*marketplacev1alpha1.MeterDefinition) *kbsm.Family {
		pvc := obj.(*v1.PersistentVolumeClaim)

		metricFamily := f(pvc, meterDefinitions)

		for _, m := range metricFamily.Metrics {
			m.LabelKeys = append(descPersistentVolumeClaimLabelsDefaultLabels, m.LabelKeys...)
			m.LabelValues = append([]string{pvc.Namespace, pvc.Name}, m.LabelValues...)
		}

		metricFamily.Metrics = MapMeterDefinitions(metricFamily.Metrics, meterDefinitions)

		return metricFamily
	}
}
