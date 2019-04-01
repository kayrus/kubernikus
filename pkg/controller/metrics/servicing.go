package metrics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(
		ServicingStatusNodes,
	)
}

var ServicingStatusNodes = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: "kubernikus",
		Subsystem: "servicing",
		Name:      "status_nodes",
		Help:      "Update Status of Nodes per Kluster"},
	[]string{"kluster_id", "status", "action"})
