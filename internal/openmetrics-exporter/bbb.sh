#!/bin/bash

grep 'json:' ../rest-client/arrays_performance.go | sed 's/^[ \t]*//' | grep -v QosRateLimit | grep Usec | sed 's/`json:"//' | sed 's/"`//' | awk '{ print "\twant[fmt.Sprintf(\"label:<name:\\\"dimension\\\" value:\\\"" $3 "\\\" > gauge:<value:%g > \", p." $1 ")] = true" }'

grep 'json:' ../rest-client/arrays_performance.go | sed 's/^[ \t]*//' | grep -v QosRateLimit | grep -v Usec | grep BytesPerSec | sed 's/`json:"//' | sed 's/"`//' | awk '{ print "\twant[fmt.Sprintf(\"label:<name:\\\"dimension\\\" value:\\\"" $3 "\\\" > gauge:<value:%g > \", p." $1 ")] = true" }'

grep 'json:' ../rest-client/arrays_performance.go | sed 's/^[ \t]*//' | grep -v QosRateLimit | grep -v Usec | grep -v BytesPerSec | grep PerSec | sed 's/`json:"//' | sed 's/"`//' | awk '{ print "\twant[fmt.Sprintf(\"label:<name:\\\"dimension\\\" value:\\\"" $3 "\\\" > gauge:<value:%g > \", p." $1 ")] = true" }'

grep 'json:' ../rest-client/arrays_performance.go | sed 's/^[ \t]*//' | grep -v QosRateLimit | grep -v Usec | grep -v BytesPerSec | grep -v PerSec | grep BytesPer | sed 's/`json:"//' | sed 's/"`//' | awk '{ print "\twant[fmt.Sprintf(\"label:<name:\\\"dimension\\\" value:\\\"" $3 "\\\" > gauge:<value:%g > \", p." $1 ")] = true" }'

grep 'json:' ../rest-client/arrays_performance.go | sed 's/^[ \t]*//' | grep -v QosRateLimit | grep -v Usec | grep -v BytesPerSec | grep -v PerSec | grep -v BytesPer | grep QueueDepth | sed 's/`json:"//' | sed 's/"`//' | awk '{ print "\twant[fmt.Sprintf(\"gauge:<value:%g > \", p." $1 ")] = true" }'

