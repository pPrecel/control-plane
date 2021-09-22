// Package mothership provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package mothership

import (
	"time"
)

// Defines values for Status.
const (
	StatusError Status = "error"

	StatusReady Status = "ready"

	StatusReconcilePending Status = "reconcile_pending"

	StatusReconciling Status = "reconciling"
)

// HTTPClusterResponse defines model for HTTPClusterResponse.
type HTTPClusterResponse struct {
	Cluster              string `json:"cluster"`
	ClusterVersion       int64  `json:"clusterVersion"`
	ConfigurationVersion int64  `json:"configurationVersion"`
	Status               Status `json:"status"`
	StatusURL            string `json:"statusURL"`
}

// HTTPClusterStatusResponse defines model for HTTPClusterStatusResponse.
type HTTPClusterStatusResponse struct {
	StatusChanges []StatusChange `json:"statusChanges"`
}

// HTTPErrorResponse defines model for HTTPErrorResponse.
type HTTPErrorResponse struct {
	Error string `json:"error"`
}

// HTTPReconcilerStatus defines model for HTTPReconcilerStatus.
type HTTPReconcilerStatus []ReconcilerStatus

// Cluster defines model for cluster.
type Cluster struct {
	// valid kubeconfig to cluster
	Kubeconfig   string       `json:"kubeconfig"`
	KymaConfig   KymaConfig   `json:"kymaConfig"`
	Metadata     Metadata     `json:"metadata"`
	RuntimeID    string       `json:"runtimeID"`
	RuntimeInput RuntimeInput `json:"runtimeInput"`
}

// Component defines model for component.
type Component struct {
	URL           string          `json:"URL"`
	Component     string          `json:"component"`
	Configuration []Configuration `json:"configuration"`
	Namespace     string          `json:"namespace"`
	Version       string          `json:"version"`
}

// Configuration defines model for configuration.
type Configuration struct {
	Key    string      `json:"key"`
	Secret bool        `json:"secret"`
	Value  interface{} `json:"value"`
}

// KymaConfig defines model for kymaConfig.
type KymaConfig struct {
	Administrators []string    `json:"administrators"`
	Components     []Component `json:"components"`
	Profile        string      `json:"profile"`
	Version        string      `json:"version"`
}

// Metadata defines model for metadata.
type Metadata struct {
	GlobalAccountID string `json:"globalAccountID"`
	InstanceID      string `json:"instanceID"`
	ServiceID       string `json:"serviceID"`
	ServicePlanID   string `json:"servicePlanID"`
	ShootName       string `json:"shootName"`
	SubAccountID    string `json:"subAccountID"`
}

// ReconcilerStatus defines model for reconcilerStatus.
type ReconcilerStatus struct {
	Cluster  string    `json:"cluster"`
	Created  time.Time `json:"created"`
	Metadata Metadata  `json:"metadata"`
	Status   string    `json:"status"`
}

// RuntimeInput defines model for runtimeInput.
type RuntimeInput struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Status defines model for status.
type Status string

// StatusChange defines model for statusChange.
type StatusChange struct {
	Duration int64     `json:"duration"`
	Started  time.Time `json:"started"`
	Status   Status    `json:"status"`
}

// BadRequest defines model for BadRequest.
type BadRequest HTTPErrorResponse

// ClusterNotFound defines model for ClusterNotFound.
type ClusterNotFound HTTPErrorResponse

// InternalError defines model for InternalError.
type InternalError HTTPErrorResponse

// Ok defines model for Ok.
type Ok HTTPClusterResponse

// PostClustersJSONBody defines parameters for PostClusters.
type PostClustersJSONBody Cluster

// PutClustersJSONBody defines parameters for PutClusters.
type PutClustersJSONBody Cluster

// GetReconcilesParams defines parameters for GetReconciles.
type GetReconcilesParams struct {
	RuntimeIDs *[]string `json:"runtimeIDs,omitempty"`
	Statuses   *[]Status `json:"statuses,omitempty"`
	Shoots     *[]string `json:"shoots,omitempty"`
}

// PostClustersJSONRequestBody defines body for PostClusters for application/json ContentType.
type PostClustersJSONRequestBody PostClustersJSONBody

// PutClustersJSONRequestBody defines body for PutClusters for application/json ContentType.
type PutClustersJSONRequestBody PutClustersJSONBody
