package config

const (
	DefaultListenAddress             = "0.0.0.0:8080"
	DefaultNodesCacheUpdateInterval  = 200
	DefaultNodesCacheUpdateQueueSize = 1000
)

// Represents the configuration of a polaris-node-sampler instance.
type NodeSamplerConfig struct {

	// The list of addresses and ports to listen on in
	// the format "<IP>:<PORT>"
	//
	// Default: [ "0.0.0.0:8080" ]
	ListenOn []string `json:"listenOn" yaml:"listenOn"`

	// The update interval for the nodes cache in milliseconds.
	//
	// Default: 200
	NodesCacheUpdateIntervalMs uint32

	// The size of the update queue in the nodes cache.
	// The update queue caches watch events that arrive between the update intervals.
	//
	// Default: 1000
	NodesCacheUpdateQueueSize uint32
}

// Sets the default values in the NodeSamplerConfig for fields that are not set properly.
func SetDefaultsNodeSamplerConfig(config *NodeSamplerConfig) {
	if config.ListenOn == nil || len(config.ListenOn) == 0 {
		config.ListenOn = []string{DefaultListenAddress}
	}
	if config.NodesCacheUpdateIntervalMs == 0 {
		config.NodesCacheUpdateIntervalMs = DefaultNodesCacheUpdateInterval
	}
	if config.NodesCacheUpdateQueueSize == 0 {
		config.NodesCacheUpdateQueueSize = DefaultNodesCacheUpdateQueueSize
	}
}
