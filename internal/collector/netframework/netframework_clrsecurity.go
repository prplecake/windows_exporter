// SPDX-License-Identifier: Apache-2.0
//
// Copyright The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build windows

package netframework

import (
	"fmt"

	"github.com/prometheus-community/windows_exporter/internal/mi"
	"github.com/prometheus-community/windows_exporter/internal/types"
	"github.com/prometheus-community/windows_exporter/internal/utils"
	"github.com/prometheus/client_golang/prometheus"
)

func (c *Collector) buildClrSecurity() {
	c.numberLinkTimeChecks = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, collectorClrSecurity+"_link_time_checks_total"),
		"Displays the total number of link-time code access security checks since the application started.",
		[]string{"process"},
		nil,
	)
	c.timeInRTChecks = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, collectorClrSecurity+"_rt_checks_time_percent"),
		"Displays the percentage of time spent performing runtime code access security checks in the last sample.",
		[]string{"process"},
		nil,
	)
	c.stackWalkDepth = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, collectorClrSecurity+"_stack_walk_depth"),
		"Displays the depth of the stack during that last runtime code access security check.",
		[]string{"process"},
		nil,
	)
	c.totalRuntimeChecks = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, collectorClrSecurity+"_runtime_checks_total"),
		"Displays the total number of runtime code access security checks performed since the application started.",
		[]string{"process"},
		nil,
	)
}

type Win32_PerfRawData_NETFramework_NETCLRSecurity struct {
	Name string `mi:"Name"`

	Frequency_PerfTime           uint32 `mi:"Frequency_PerfTime"`
	NumberLinkTimeChecks         uint32 `mi:"NumberLinkTimeChecks"`
	PercentTimeinRTchecks        uint32 `mi:"PercentTimeinRTchecks"`
	PercentTimeSigAuthenticating uint64 `mi:"PercentTimeSigAuthenticating"`
	StackWalkDepth               uint32 `mi:"StackWalkDepth"`
	TotalRuntimeChecks           uint32 `mi:"TotalRuntimeChecks"`
}

func (c *Collector) collectClrSecurity(ch chan<- prometheus.Metric) error {
	var dst []Win32_PerfRawData_NETFramework_NETCLRSecurity
	if err := c.miSession.Query(&dst, mi.NamespaceRootCIMv2, utils.Must(mi.NewQuery("SELECT * FROM Win32_PerfRawData_NETFramework_NETCLRSecurity"))); err != nil {
		return fmt.Errorf("WMI query failed: %w", err)
	}

	for _, process := range dst {
		if process.Name == "_Global_" {
			continue
		}

		ch <- prometheus.MustNewConstMetric(
			c.numberLinkTimeChecks,
			prometheus.CounterValue,
			float64(process.NumberLinkTimeChecks),
			process.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.timeInRTChecks,
			prometheus.GaugeValue,
			float64(process.PercentTimeinRTchecks)/float64(process.Frequency_PerfTime),
			process.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.stackWalkDepth,
			prometheus.GaugeValue,
			float64(process.StackWalkDepth),
			process.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.totalRuntimeChecks,
			prometheus.CounterValue,
			float64(process.TotalRuntimeChecks),
			process.Name,
		)
	}

	return nil
}
