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

package httphandler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus-community/windows_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/collectors/version"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Interface guard.
var _ http.Handler = (*MetricsHTTPHandler)(nil)

const defaultScrapeTimeout = 10.0

type MetricsHTTPHandler struct {
	metricCollectors *collector.Collection
	// exporterMetricsRegistry is a separate registry for the metrics about
	// the exporter itself.
	exporterMetricsRegistry *prometheus.Registry

	logger  *slog.Logger
	options Options
}

type Options struct {
	DisableExporterMetrics bool
	TimeoutMargin          float64
}

func New(logger *slog.Logger, metricCollectors *collector.Collection, options *Options) *MetricsHTTPHandler {
	if options == nil {
		options = &Options{
			DisableExporterMetrics: false,
			TimeoutMargin:          0.5,
		}
	}

	handler := &MetricsHTTPHandler{
		metricCollectors: metricCollectors,
		logger:           logger,
		options:          *options,
	}

	if !options.DisableExporterMetrics {
		handler.exporterMetricsRegistry = prometheus.NewRegistry()
		handler.exporterMetricsRegistry.MustRegister(
			collectors.NewBuildInfoCollector(),
			collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
			collectors.NewGoCollector(),
		)
	}

	return handler
}

func (c *MetricsHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := c.logger.With(
		slog.String("remote", r.RemoteAddr),
	)

	scrapeTimeout := c.getScrapeTimeout(logger, r)

	handler, err := c.handlerFactory(logger, scrapeTimeout, r.URL.Query()["collect[]"])
	if err != nil {
		logger.Warn("Couldn't create filtered metrics handler",
			slog.Any("err", err),
		)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "Couldn't create filtered metrics handler: %s", err)

		return
	}

	handler.ServeHTTP(w, r)
}

func (c *MetricsHTTPHandler) getScrapeTimeout(logger *slog.Logger, r *http.Request) time.Duration {
	var timeoutSeconds float64

	if v := r.Header.Get("X-Prometheus-Scrape-Timeout-Seconds"); v != "" {
		var err error

		timeoutSeconds, err = strconv.ParseFloat(v, 64)
		if err != nil {
			logger.Warn(fmt.Sprintf("Couldn't parse X-Prometheus-Scrape-Timeout-Seconds: %q. Defaulting timeout to %f", v, defaultScrapeTimeout))
		}
	}

	if timeoutSeconds == 0 {
		timeoutSeconds = defaultScrapeTimeout
	}

	timeoutSeconds -= c.options.TimeoutMargin

	return time.Duration(timeoutSeconds) * time.Second
}

func (c *MetricsHTTPHandler) handlerFactory(logger *slog.Logger, scrapeTimeout time.Duration, requestedCollectors []string) (http.Handler, error) {
	reg := prometheus.NewRegistry()
	reg.MustRegister(version.NewCollector("windows_exporter"))

	collectionHandler, err := c.metricCollectors.NewHandler(scrapeTimeout, c.logger, requestedCollectors)
	if err != nil {
		return nil, fmt.Errorf("couldn't create collector handler: %w", err)
	}

	if err := reg.Register(collectionHandler); err != nil {
		return nil, fmt.Errorf("couldn't register Prometheus collector: %w", err)
	}

	var regHandler http.Handler
	if c.exporterMetricsRegistry != nil {
		regHandler = promhttp.HandlerFor(
			prometheus.Gatherers{c.exporterMetricsRegistry, reg},
			promhttp.HandlerOpts{
				ErrorLog:            slog.NewLogLogger(logger.Handler(), slog.LevelError),
				ErrorHandling:       promhttp.ContinueOnError,
				MaxRequestsInFlight: 1,
				Registry:            c.exporterMetricsRegistry,
				EnableOpenMetrics:   true,
				ProcessStartTime:    c.metricCollectors.GetStartTime(),
			},
		)

		// Note that we have to use h.exporterMetricsRegistry here to
		// use the same promhttp metrics for all expositions.
		regHandler = promhttp.InstrumentMetricHandler(
			c.exporterMetricsRegistry, regHandler,
		)
	} else {
		regHandler = promhttp.HandlerFor(
			reg,
			promhttp.HandlerOpts{
				ErrorLog:            slog.NewLogLogger(logger.Handler(), slog.LevelError),
				ErrorHandling:       promhttp.ContinueOnError,
				MaxRequestsInFlight: 1,
				EnableOpenMetrics:   true,
				ProcessStartTime:    c.metricCollectors.GetStartTime(),
			},
		)
	}

	return regHandler, nil
}
