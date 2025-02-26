// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/filter"
)

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for nsxt metrics.
type MetricsConfig struct {
	NsxtNodeCPUUtilization        MetricConfig `mapstructure:"nsxt.node.cpu.utilization"`
	NsxtNodeFilesystemUsage       MetricConfig `mapstructure:"nsxt.node.filesystem.usage"`
	NsxtNodeFilesystemUtilization MetricConfig `mapstructure:"nsxt.node.filesystem.utilization"`
	NsxtNodeMemoryCacheUsage      MetricConfig `mapstructure:"nsxt.node.memory.cache.usage"`
	NsxtNodeMemoryUsage           MetricConfig `mapstructure:"nsxt.node.memory.usage"`
	NsxtNodeNetworkIo             MetricConfig `mapstructure:"nsxt.node.network.io"`
	NsxtNodeNetworkPacketCount    MetricConfig `mapstructure:"nsxt.node.network.packet.count"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		NsxtNodeCPUUtilization: MetricConfig{
			Enabled: true,
		},
		NsxtNodeFilesystemUsage: MetricConfig{
			Enabled: true,
		},
		NsxtNodeFilesystemUtilization: MetricConfig{
			Enabled: true,
		},
		NsxtNodeMemoryCacheUsage: MetricConfig{
			Enabled: true,
		},
		NsxtNodeMemoryUsage: MetricConfig{
			Enabled: true,
		},
		NsxtNodeNetworkIo: MetricConfig{
			Enabled: true,
		},
		NsxtNodeNetworkPacketCount: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
	// Experimental: MetricsInclude defines a list of filters for attribute values.
	// If the list is not empty, only metrics with matching resource attribute values will be emitted.
	MetricsInclude []filter.Config `mapstructure:"metrics_include"`
	// Experimental: MetricsExclude defines a list of filters for attribute values.
	// If the list is not empty, metrics with matching resource attribute values will not be emitted.
	// MetricsInclude has higher priority than MetricsExclude.
	MetricsExclude []filter.Config `mapstructure:"metrics_exclude"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for nsxt resource attributes.
type ResourceAttributesConfig struct {
	DeviceID     ResourceAttributeConfig `mapstructure:"device.id"`
	NsxtNodeID   ResourceAttributeConfig `mapstructure:"nsxt.node.id"`
	NsxtNodeName ResourceAttributeConfig `mapstructure:"nsxt.node.name"`
	NsxtNodeType ResourceAttributeConfig `mapstructure:"nsxt.node.type"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		DeviceID: ResourceAttributeConfig{
			Enabled: true,
		},
		NsxtNodeID: ResourceAttributeConfig{
			Enabled: true,
		},
		NsxtNodeName: ResourceAttributeConfig{
			Enabled: true,
		},
		NsxtNodeType: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for nsxt metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
