#!/bin/bash

grep 'json:' ../rest-client/arrays_performance.go | sed 's/^[ \t]*//' | grep Usec | grep -v QosRateLimit| sed 's/`json:"//' | sed 's/"`//' | awk '{print "\tch <- prometheus.MustNewConstMetric(\n\t\tc.LatencyDesc,\n\t\tprometheus.GaugeValue,\n\t\tfloat64(ap." $1 "),\n\t\t\"" $3 "\",\n\t)"}'

grep 'json:' ../rest-client/arrays_performance.go | sed 's/^[ \t]*//' | grep -v Usec | sed 's/`json:"//' | sed 's/"`//' | grep BytesPerSec | awk '{print "\tch <- prometheus.MustNewConstMetric(\n\t\tc.BandwidthDesc,\n\t\tprometheus.GaugeValue,\n\t\tfloat64(ap." $1 "),\n\t\t\"" $3 "\",\n\t)"}'

grep 'json:' ../rest-client/arrays_performance.go | sed 's/^[ \t]*//' | grep -v Usec | sed 's/`json:"//' | sed 's/"`//' | grep -v BytesPerSec | grep PerSec | awk '{print "\tch <- prometheus.MustNewConstMetric(\n\t\tc.ThroughputDesc,\n\t\tprometheus.GaugeValue,\n\t\tfloat64(ap." $1 "),\n\t\t\"" $3 "\",\n\t)"}'

grep 'json:' ../rest-client/arrays_performance.go | sed 's/^[ \t]*//' | grep -v Usec | sed 's/`json:"//' | sed 's/"`//' | grep -v BytesPerSec | grep -v PerSec | grep BytesPer | awk '{print "\tch <- prometheus.MustNewConstMetric(\n\t\tc.AverageSizeDesc,\n\t\tprometheus.GaugeValue,\n\t\tfloat64(ap." $1 "),\n\t\t\"" $3 "\",\n\t)"}'

