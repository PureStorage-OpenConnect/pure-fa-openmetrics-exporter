package collectors

import (
	"context"
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/collectors"
)

func Collector(ctx context.Context, metrics string, registry *prometheus.Registry, faclient *client.FAClient) bool {

	arrayscoll := NewArraysCollector(faclient)
	registry.MustRegister(
		//collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		//collectors.NewGoCollector(),
		arrayscoll,
	)
	if metrics == "all" || metrics == "array" {
		alertscoll := NewAlertsCollector(faclient)
		arrayperfcoll := NewArraysPerformanceCollector(faclient)
		arrayspacecoll := NewArraySpaceCollector(faclient)
		hwcoll := NewHardwareCollector(faclient)
		controllercol := NewControllersCollector(faclient)
		drcoll := NewDriveCollector(faclient)
		nicsperfcoll := NewNetworkInterfacesPerformanceCollector(faclient)
		portscoll := NewPortsCollector(faclient)
		interfacecoll := NewNetworkInterfacesCollector(faclient)
		registry.MustRegister(
			controllercol,
			alertscoll,
			arrayperfcoll,
			arrayspacecoll,
			hwcoll,
			drcoll,
			nicsperfcoll,
			portscoll,
			interfacecoll,
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
		volspacecoll := NewVolumesCollector(vols)
		registry.MustRegister(
			volperfcoll,
			volspacecoll,
		)
	}
	return true
}
