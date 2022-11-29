package client

import (
	core "k8s.io/api/core/v1"
	"polaris-slo-cloud.github.io/polaris-scheduler/v2/framework/util"
)

var (
	_ error = (*PolarisErrorDto)(nil)
)

// Contains the scheduling decision for a pod within a cluster.
type ClusterSchedulingDecision struct {
	// The Pod to be scheduled.
	Pod *core.Pod `json:"pod" yaml:"pod"`

	// The name of the node, to which the pod has been assigned.
	NodeName string `json:"nodeName" yaml:"nodeName"`
}

// Augments a node with information computed by the Polaris Scheduling framework.
type ClusterNode struct {
	*core.Node `json:",inline" yaml:",inline"`

	// The pods that are already scheduled on (bound to) this node.
	Pods []*ClusterPod

	// The pods that are queued to be bound to this node.
	// These pods are currently in the binding pipeline, but their resources are already accounted for in
	// the node's AvailableResources, because committing a scheduling decision may take some time.
	QueuedPods []*ClusterPod

	// The resources that are currently available for allocation on the node.
	//
	// Unlike the Kubernetes Allocatable field, these AvailableResources already accounts for
	// resources consumed by other pods.
	AvailableResources *util.Resources `json:"availableResources" yaml:"availableResources"`

	// The total amount of resources that are available on the node.
	TotalResources *util.Resources `json:"totalResources" yaml:"totalResources"`
}

// Creates a new cluster node, based on the specified node object, assuming that no pods are scheduled on it yet.
func NewClusterNode(node *core.Node) *ClusterNode {
	cn := &ClusterNode{
		Node: node,
		// A new node does not have any pods yet, so both resources are set to the total amount that is allocatable.
		AvailableResources: util.NewResourcesFromList(node.Status.Allocatable),
		TotalResources:     util.NewResourcesFromList(node.Status.Allocatable),
		Pods:               make([]*ClusterPod, 0),
		QueuedPods:         make([]*ClusterPod, 0),
	}
	return cn
}

// Creates a new cluster node, based on the specified node object and the pods that are already scheduled and queued on it.
func NewClusterNodeWithPods(node *core.Node, pods []*ClusterPod, queuedPods []*ClusterPod) *ClusterNode {
	cn := &ClusterNode{
		Node:               node,
		AvailableResources: util.NewResourcesFromList(node.Status.Allocatable),
		TotalResources:     util.NewResourcesFromList(node.Status.Allocatable),
		Pods:               pods,
		QueuedPods:         queuedPods,
	}

	for _, pod := range pods {
		cn.AvailableResources.Subtract(pod.TotalResources)
	}
	for _, pod := range queuedPods {
		cn.AvailableResources.Subtract(pod.TotalResources)
	}

	return cn
}

// Condensed information about an existing pod in a cluster.
type ClusterPod struct {

	// The namespace of the pod.
	Namespace string

	// The name of the pod.
	Name string

	// The total resources consumed by this pod.
	TotalResources *util.Resources

	// Affinity/anti-affinity information.
	// This is nil, if not present on the pod.
	Affinity *core.Affinity
}

// Creates a new ClusterPod, based on the specified pod object.
func NewClusterPod(pod *core.Pod) *ClusterPod {
	cp := &ClusterPod{
		Namespace:      pod.Namespace,
		Name:           pod.Name,
		TotalResources: util.CalculateTotalPodResources(pod),
		Affinity:       pod.Spec.Affinity,
	}
	return cp
}

// A generic DTO for transmitting error information.
type PolarisErrorDto struct {
	Message string `json:"message" yaml:"message"`
}

func NewPolarisErrorDto(err error) *PolarisErrorDto {
	polarisErr := &PolarisErrorDto{
		Message: err.Error(),
	}
	return polarisErr
}

// Error implements error
func (e *PolarisErrorDto) Error() string {
	return e.Message
}
