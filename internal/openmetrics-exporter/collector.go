package collectors

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

func Collector(ctx context.Context, metrics string, registry *prometheus.Registry, faclient *client.FAClient) bool {

	arrayscoll := NewArraysCollector(faclient)
	registry.MustRegister(
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(),
		arrayscoll,
	)
	if metrics == "all" || metrics == "array" {
		alertscoll := NewAlertsCollector(faclient)
		arrayperfcoll := NewArraysPerformanceCollector(faclient)
		arrayspacecoll := NewArraySpaceCollector(faclient)
		hwcoll := NewHardwareCollector(faclient)
		drcoll := NewDriveCollector(faclient)
		nicsperfcoll := NewNetworkInterfacesPerformanceCollector(faclient)
		registry.MustRegister(
		           alertscoll,
		           arrayperfcoll,
		           arrayspacecoll,
		           hwcoll,
		           drcoll,
		           nicsperfcoll,
		         )
	}
	if metrics == "all" || metrics == "directories" {
		dirperfcoll := NewDirectoriesPerformanceCollector(faclient)
		dirspacecoll := NewDirectoriesSpaceCollector(faclient)
		registry.MustRegister(
		           dirperfcoll,
		           dirspacecoll,
		         )
	}
	if metrics == "all" || metrics == "hosts" {
		hostperfcoll := NewHostsPerformanceCollector(faclient)
		hostspacecoll := NewHostsSpaceCollector(faclient)
		hostconncoll := NewHostConnectionsCollector(faclient)
		registry.MustRegister(
		           hostconncoll,
		           hostperfcoll,
		           hostspacecoll,
		         )
	}
	if metrics == "all" || metrics == "pods" {
		podsperfcoll := NewPodsPerformanceCollector(faclient)
		podsspacecoll := NewPodsSpaceCollector(faclient)
		podsperfrepl := NewPodsPerformanceReplicationCollector(faclient)
		podreplinkperfcoll := NewPodReplicaLinksPerformanceCollector(faclient)
		podreplinklagcoll := NewPodReplicaLinksLagCollector(faclient)
		registry.MustRegister(
		           podsperfcoll,
		           podsspacecoll,
		           podsperfrepl,
		           podreplinkperfcoll,
		           podreplinklagcoll,
		         )
	}
	if metrics == "all" || metrics == "volumes" {
		vols := faclient.GetVolumes()
		volperfcoll := NewVolumesPerformanceCollector(faclient, vols)
		volspacecoll := NewVolumesSpaceCollector(vols)
		registry.MustRegister(
		           volperfcoll,
		           volspacecoll,
		         )
	}
	return true
}
