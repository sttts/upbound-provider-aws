// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CloudwatchConfigurationInitParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A block the specifies how an opsworks logs look like. See Log Streams.
	LogStreams []LogStreamsInitParameters `json:"logStreams,omitempty" tf:"log_streams,omitempty"`
}

type CloudwatchConfigurationObservation struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A block the specifies how an opsworks logs look like. See Log Streams.
	LogStreams []LogStreamsObservation `json:"logStreams,omitempty" tf:"log_streams,omitempty"`
}

type CloudwatchConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A block the specifies how an opsworks logs look like. See Log Streams.
	// +kubebuilder:validation:Optional
	LogStreams []LogStreamsParameters `json:"logStreams,omitempty" tf:"log_streams,omitempty"`
}

type CustomLayerInitParameters struct {

	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `json:"autoAssignElasticIps,omitempty" tf:"auto_assign_elastic_ips,omitempty"`

	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `json:"autoAssignPublicIps,omitempty" tf:"auto_assign_public_ips,omitempty"`

	// Whether to enable auto-healing for the layer.
	AutoHealing *bool `json:"autoHealing,omitempty" tf:"auto_healing,omitempty"`

	// Will create an EBS volume and connect it to the layer's instances. See Cloudwatch Configuration.
	CloudwatchConfiguration []CloudwatchConfigurationInitParameters `json:"cloudwatchConfiguration,omitempty" tf:"cloudwatch_configuration,omitempty"`

	CustomConfigureRecipes []*string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes,omitempty"`

	CustomDeployRecipes []*string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes,omitempty"`

	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn,omitempty"`

	// Custom JSON attributes to apply to the layer.
	CustomJSON *string `json:"customJson,omitempty" tf:"custom_json,omitempty"`

	CustomSetupRecipes []*string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes,omitempty"`

	CustomShutdownRecipes []*string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes,omitempty"`

	CustomUndeployRecipes []*string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes,omitempty"`

	// Whether to enable Elastic Load Balancing connection draining.
	DrainELBOnShutdown *bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown,omitempty"`

	// Will create an EBS volume and connect it to the layer's instances. See EBS Volume.
	EBSVolume []EBSVolumeInitParameters `json:"ebsVolume,omitempty" tf:"ebs_volume,omitempty"`

	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer,omitempty"`

	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot,omitempty"`

	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int64 `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout,omitempty"`

	// Load-based auto scaling configuration. See Load Based AutoScaling
	LoadBasedAutoScaling []LoadBasedAutoScalingInitParameters `json:"loadBasedAutoScaling,omitempty" tf:"load_based_auto_scaling,omitempty"`

	// A human-readable name for the layer.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A short, machine-readable name for the layer, which will be used to identify it in the Chef node JSON.
	ShortName *string `json:"shortName,omitempty" tf:"short_name,omitempty"`

	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []*string `json:"systemPackages,omitempty" tf:"system_packages,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether to use EBS-optimized instances.
	UseEBSOptimizedInstances *bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances,omitempty"`
}

type CustomLayerObservation struct {

	// The Amazon Resource Name(ARN) of the layer.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `json:"autoAssignElasticIps,omitempty" tf:"auto_assign_elastic_ips,omitempty"`

	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `json:"autoAssignPublicIps,omitempty" tf:"auto_assign_public_ips,omitempty"`

	// Whether to enable auto-healing for the layer.
	AutoHealing *bool `json:"autoHealing,omitempty" tf:"auto_healing,omitempty"`

	// Will create an EBS volume and connect it to the layer's instances. See Cloudwatch Configuration.
	CloudwatchConfiguration []CloudwatchConfigurationObservation `json:"cloudwatchConfiguration,omitempty" tf:"cloudwatch_configuration,omitempty"`

	CustomConfigureRecipes []*string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes,omitempty"`

	CustomDeployRecipes []*string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes,omitempty"`

	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn,omitempty"`

	// Custom JSON attributes to apply to the layer.
	CustomJSON *string `json:"customJson,omitempty" tf:"custom_json,omitempty"`

	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []*string `json:"customSecurityGroupIds,omitempty" tf:"custom_security_group_ids,omitempty"`

	CustomSetupRecipes []*string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes,omitempty"`

	CustomShutdownRecipes []*string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes,omitempty"`

	CustomUndeployRecipes []*string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes,omitempty"`

	// Whether to enable Elastic Load Balancing connection draining.
	DrainELBOnShutdown *bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown,omitempty"`

	// Will create an EBS volume and connect it to the layer's instances. See EBS Volume.
	EBSVolume []EBSVolumeObservation `json:"ebsVolume,omitempty" tf:"ebs_volume,omitempty"`

	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer,omitempty"`

	// The id of the layer.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot,omitempty"`

	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int64 `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout,omitempty"`

	// Load-based auto scaling configuration. See Load Based AutoScaling
	LoadBasedAutoScaling []LoadBasedAutoScalingObservation `json:"loadBasedAutoScaling,omitempty" tf:"load_based_auto_scaling,omitempty"`

	// A human-readable name for the layer.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A short, machine-readable name for the layer, which will be used to identify it in the Chef node JSON.
	ShortName *string `json:"shortName,omitempty" tf:"short_name,omitempty"`

	// ID of the stack the layer will belong to.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []*string `json:"systemPackages,omitempty" tf:"system_packages,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Whether to use EBS-optimized instances.
	UseEBSOptimizedInstances *bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances,omitempty"`
}

type CustomLayerParameters struct {

	// Whether to automatically assign an elastic IP address to the layer's instances.
	// +kubebuilder:validation:Optional
	AutoAssignElasticIps *bool `json:"autoAssignElasticIps,omitempty" tf:"auto_assign_elastic_ips,omitempty"`

	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	// +kubebuilder:validation:Optional
	AutoAssignPublicIps *bool `json:"autoAssignPublicIps,omitempty" tf:"auto_assign_public_ips,omitempty"`

	// Whether to enable auto-healing for the layer.
	// +kubebuilder:validation:Optional
	AutoHealing *bool `json:"autoHealing,omitempty" tf:"auto_healing,omitempty"`

	// Will create an EBS volume and connect it to the layer's instances. See Cloudwatch Configuration.
	// +kubebuilder:validation:Optional
	CloudwatchConfiguration []CloudwatchConfigurationParameters `json:"cloudwatchConfiguration,omitempty" tf:"cloudwatch_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	CustomConfigureRecipes []*string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes,omitempty"`

	// +kubebuilder:validation:Optional
	CustomDeployRecipes []*string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes,omitempty"`

	// The ARN of an IAM profile that will be used for the layer's instances.
	// +kubebuilder:validation:Optional
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn,omitempty"`

	// Custom JSON attributes to apply to the layer.
	// +kubebuilder:validation:Optional
	CustomJSON *string `json:"customJson,omitempty" tf:"custom_json,omitempty"`

	// References to SecurityGroup in ec2 to populate customSecurityGroupIds.
	// +kubebuilder:validation:Optional
	CustomSecurityGroupIDRefs []v1.Reference `json:"customSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate customSecurityGroupIds.
	// +kubebuilder:validation:Optional
	CustomSecurityGroupIDSelector *v1.Selector `json:"customSecurityGroupIdSelector,omitempty" tf:"-"`

	// Ids for a set of security groups to apply to the layer's instances.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=CustomSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=CustomSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	CustomSecurityGroupIds []*string `json:"customSecurityGroupIds,omitempty" tf:"custom_security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	CustomSetupRecipes []*string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes,omitempty"`

	// +kubebuilder:validation:Optional
	CustomShutdownRecipes []*string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes,omitempty"`

	// +kubebuilder:validation:Optional
	CustomUndeployRecipes []*string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes,omitempty"`

	// Whether to enable Elastic Load Balancing connection draining.
	// +kubebuilder:validation:Optional
	DrainELBOnShutdown *bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown,omitempty"`

	// Will create an EBS volume and connect it to the layer's instances. See EBS Volume.
	// +kubebuilder:validation:Optional
	EBSVolume []EBSVolumeParameters `json:"ebsVolume,omitempty" tf:"ebs_volume,omitempty"`

	// Name of an Elastic Load Balancer to attach to this layer
	// +kubebuilder:validation:Optional
	ElasticLoadBalancer *string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer,omitempty"`

	// Whether to install OS and package updates on each instance when it boots.
	// +kubebuilder:validation:Optional
	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot,omitempty"`

	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	// +kubebuilder:validation:Optional
	InstanceShutdownTimeout *int64 `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout,omitempty"`

	// Load-based auto scaling configuration. See Load Based AutoScaling
	// +kubebuilder:validation:Optional
	LoadBasedAutoScaling []LoadBasedAutoScalingParameters `json:"loadBasedAutoScaling,omitempty" tf:"load_based_auto_scaling,omitempty"`

	// A human-readable name for the layer.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A short, machine-readable name for the layer, which will be used to identify it in the Chef node JSON.
	// +kubebuilder:validation:Optional
	ShortName *string `json:"shortName,omitempty" tf:"short_name,omitempty"`

	// ID of the stack the layer will belong to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/opsworks/v1beta1.Stack
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// Reference to a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDRef *v1.Reference `json:"stackIdRef,omitempty" tf:"-"`

	// Selector for a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDSelector *v1.Selector `json:"stackIdSelector,omitempty" tf:"-"`

	// Names of a set of system packages to install on the layer's instances.
	// +kubebuilder:validation:Optional
	SystemPackages []*string `json:"systemPackages,omitempty" tf:"system_packages,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether to use EBS-optimized instances.
	// +kubebuilder:validation:Optional
	UseEBSOptimizedInstances *bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances,omitempty"`
}

type DownscalingInitParameters struct {

	// Custom Cloudwatch auto scaling alarms, to be used as thresholds. This parameter takes a list of up to five alarm names, which are case sensitive and must be in the same region as the stack.
	Alarms []*string `json:"alarms,omitempty" tf:"alarms,omitempty"`

	// The CPU utilization threshold, as a percent of the available CPU. A value of -1 disables the threshold.
	CPUThreshold *float64 `json:"cpuThreshold,omitempty" tf:"cpu_threshold,omitempty"`

	// The amount of time (in minutes) after a scaling event occurs that AWS OpsWorks Stacks should ignore metrics and suppress additional scaling events.
	IgnoreMetricsTime *int64 `json:"ignoreMetricsTime,omitempty" tf:"ignore_metrics_time,omitempty"`

	// The number of instances to add or remove when the load exceeds a threshold.
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The load threshold. A value of -1 disables the threshold.
	LoadThreshold *float64 `json:"loadThreshold,omitempty" tf:"load_threshold,omitempty"`

	// The memory utilization threshold, as a percent of the available memory. A value of -1 disables the threshold.
	MemoryThreshold *float64 `json:"memoryThreshold,omitempty" tf:"memory_threshold,omitempty"`

	// The amount of time, in minutes, that the load must exceed a threshold before more instances are added or removed.
	ThresholdsWaitTime *int64 `json:"thresholdsWaitTime,omitempty" tf:"thresholds_wait_time,omitempty"`
}

type DownscalingObservation struct {

	// Custom Cloudwatch auto scaling alarms, to be used as thresholds. This parameter takes a list of up to five alarm names, which are case sensitive and must be in the same region as the stack.
	Alarms []*string `json:"alarms,omitempty" tf:"alarms,omitempty"`

	// The CPU utilization threshold, as a percent of the available CPU. A value of -1 disables the threshold.
	CPUThreshold *float64 `json:"cpuThreshold,omitempty" tf:"cpu_threshold,omitempty"`

	// The amount of time (in minutes) after a scaling event occurs that AWS OpsWorks Stacks should ignore metrics and suppress additional scaling events.
	IgnoreMetricsTime *int64 `json:"ignoreMetricsTime,omitempty" tf:"ignore_metrics_time,omitempty"`

	// The number of instances to add or remove when the load exceeds a threshold.
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The load threshold. A value of -1 disables the threshold.
	LoadThreshold *float64 `json:"loadThreshold,omitempty" tf:"load_threshold,omitempty"`

	// The memory utilization threshold, as a percent of the available memory. A value of -1 disables the threshold.
	MemoryThreshold *float64 `json:"memoryThreshold,omitempty" tf:"memory_threshold,omitempty"`

	// The amount of time, in minutes, that the load must exceed a threshold before more instances are added or removed.
	ThresholdsWaitTime *int64 `json:"thresholdsWaitTime,omitempty" tf:"thresholds_wait_time,omitempty"`
}

type DownscalingParameters struct {

	// Custom Cloudwatch auto scaling alarms, to be used as thresholds. This parameter takes a list of up to five alarm names, which are case sensitive and must be in the same region as the stack.
	// +kubebuilder:validation:Optional
	Alarms []*string `json:"alarms,omitempty" tf:"alarms,omitempty"`

	// The CPU utilization threshold, as a percent of the available CPU. A value of -1 disables the threshold.
	// +kubebuilder:validation:Optional
	CPUThreshold *float64 `json:"cpuThreshold,omitempty" tf:"cpu_threshold,omitempty"`

	// The amount of time (in minutes) after a scaling event occurs that AWS OpsWorks Stacks should ignore metrics and suppress additional scaling events.
	// +kubebuilder:validation:Optional
	IgnoreMetricsTime *int64 `json:"ignoreMetricsTime,omitempty" tf:"ignore_metrics_time,omitempty"`

	// The number of instances to add or remove when the load exceeds a threshold.
	// +kubebuilder:validation:Optional
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The load threshold. A value of -1 disables the threshold.
	// +kubebuilder:validation:Optional
	LoadThreshold *float64 `json:"loadThreshold,omitempty" tf:"load_threshold,omitempty"`

	// The memory utilization threshold, as a percent of the available memory. A value of -1 disables the threshold.
	// +kubebuilder:validation:Optional
	MemoryThreshold *float64 `json:"memoryThreshold,omitempty" tf:"memory_threshold,omitempty"`

	// The amount of time, in minutes, that the load must exceed a threshold before more instances are added or removed.
	// +kubebuilder:validation:Optional
	ThresholdsWaitTime *int64 `json:"thresholdsWaitTime,omitempty" tf:"thresholds_wait_time,omitempty"`
}

type EBSVolumeInitParameters struct {

	// Encrypt the volume.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// For PIOPS volumes, the IOPS per disk.
	Iops *int64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The path to mount the EBS volume on the layer's instances.
	MountPoint *string `json:"mountPoint,omitempty" tf:"mount_point,omitempty"`

	// The number of disks to use for the EBS volume.
	NumberOfDisks *int64 `json:"numberOfDisks,omitempty" tf:"number_of_disks,omitempty"`

	// The RAID level to use for the volume.
	RaidLevel *string `json:"raidLevel,omitempty" tf:"raid_level,omitempty"`

	// The size of the volume in gigabytes.
	Size *int64 `json:"size,omitempty" tf:"size,omitempty"`

	// The type of volume to create. This may be standard (the default), io1 or gp2.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EBSVolumeObservation struct {

	// Encrypt the volume.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// For PIOPS volumes, the IOPS per disk.
	Iops *int64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The path to mount the EBS volume on the layer's instances.
	MountPoint *string `json:"mountPoint,omitempty" tf:"mount_point,omitempty"`

	// The number of disks to use for the EBS volume.
	NumberOfDisks *int64 `json:"numberOfDisks,omitempty" tf:"number_of_disks,omitempty"`

	// The RAID level to use for the volume.
	RaidLevel *string `json:"raidLevel,omitempty" tf:"raid_level,omitempty"`

	// The size of the volume in gigabytes.
	Size *int64 `json:"size,omitempty" tf:"size,omitempty"`

	// The type of volume to create. This may be standard (the default), io1 or gp2.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EBSVolumeParameters struct {

	// Encrypt the volume.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// For PIOPS volumes, the IOPS per disk.
	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The path to mount the EBS volume on the layer's instances.
	// +kubebuilder:validation:Optional
	MountPoint *string `json:"mountPoint" tf:"mount_point,omitempty"`

	// The number of disks to use for the EBS volume.
	// +kubebuilder:validation:Optional
	NumberOfDisks *int64 `json:"numberOfDisks" tf:"number_of_disks,omitempty"`

	// The RAID level to use for the volume.
	// +kubebuilder:validation:Optional
	RaidLevel *string `json:"raidLevel,omitempty" tf:"raid_level,omitempty"`

	// The size of the volume in gigabytes.
	// +kubebuilder:validation:Optional
	Size *int64 `json:"size" tf:"size,omitempty"`

	// The type of volume to create. This may be standard (the default), io1 or gp2.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LoadBasedAutoScalingInitParameters struct {

	// The downscaling settings, as defined below, used for load-based autoscaling
	Downscaling []DownscalingInitParameters `json:"downscaling,omitempty" tf:"downscaling,omitempty"`

	// Whether load-based auto scaling is enabled for the layer.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// The upscaling settings, as defined below, used for load-based autoscaling
	Upscaling []UpscalingInitParameters `json:"upscaling,omitempty" tf:"upscaling,omitempty"`
}

type LoadBasedAutoScalingObservation struct {

	// The downscaling settings, as defined below, used for load-based autoscaling
	Downscaling []DownscalingObservation `json:"downscaling,omitempty" tf:"downscaling,omitempty"`

	// Whether load-based auto scaling is enabled for the layer.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// The upscaling settings, as defined below, used for load-based autoscaling
	Upscaling []UpscalingObservation `json:"upscaling,omitempty" tf:"upscaling,omitempty"`
}

type LoadBasedAutoScalingParameters struct {

	// The downscaling settings, as defined below, used for load-based autoscaling
	// +kubebuilder:validation:Optional
	Downscaling []DownscalingParameters `json:"downscaling,omitempty" tf:"downscaling,omitempty"`

	// Whether load-based auto scaling is enabled for the layer.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// The upscaling settings, as defined below, used for load-based autoscaling
	// +kubebuilder:validation:Optional
	Upscaling []UpscalingParameters `json:"upscaling,omitempty" tf:"upscaling,omitempty"`
}

type LogStreamsInitParameters struct {

	// Specifies the max number of log events in a batch, up to 10000. The default value is 1000.
	BatchCount *int64 `json:"batchCount,omitempty" tf:"batch_count,omitempty"`

	// Specifies the maximum size of log events in a batch, in bytes, up to 1048576 bytes. The default value is 32768 bytes.
	BatchSize *int64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// Specifies the time duration for the batching of log events. The minimum value is 5000 and default value is 5000.
	BufferDuration *int64 `json:"bufferDuration,omitempty" tf:"buffer_duration,omitempty"`

	// Specifies how the timestamp is extracted from logs. For more information, see the CloudWatch Logs Agent Reference (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AgentReference.html).
	DatetimeFormat *string `json:"datetimeFormat,omitempty" tf:"datetime_format,omitempty"`

	// Specifies the encoding of the log file so that the file can be read correctly. The default is utf_8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// Specifies log files that you want to push to CloudWatch Logs. File can point to a specific file or multiple files (by using wild card characters such as /var/log/system.log*).
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// Specifies the range of lines for identifying a file. The valid values are one number, or two dash-delimited numbers, such as 1, 2-5. The default value is 1.
	FileFingerprintLines *string `json:"fileFingerprintLines,omitempty" tf:"file_fingerprint_lines,omitempty"`

	// Specifies where to start to read data (start_of_file or end_of_file). The default is start_of_file.
	InitialPosition *string `json:"initialPosition,omitempty" tf:"initial_position,omitempty"`

	// Specifies the destination log group. A log group is created automatically if it doesn't already exist.
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// Specifies the pattern for identifying the start of a log message.
	MultilineStartPattern *string `json:"multilineStartPattern,omitempty" tf:"multiline_start_pattern,omitempty"`

	// Specifies the time zone of log event time stamps.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type LogStreamsObservation struct {

	// Specifies the max number of log events in a batch, up to 10000. The default value is 1000.
	BatchCount *int64 `json:"batchCount,omitempty" tf:"batch_count,omitempty"`

	// Specifies the maximum size of log events in a batch, in bytes, up to 1048576 bytes. The default value is 32768 bytes.
	BatchSize *int64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// Specifies the time duration for the batching of log events. The minimum value is 5000 and default value is 5000.
	BufferDuration *int64 `json:"bufferDuration,omitempty" tf:"buffer_duration,omitempty"`

	// Specifies how the timestamp is extracted from logs. For more information, see the CloudWatch Logs Agent Reference (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AgentReference.html).
	DatetimeFormat *string `json:"datetimeFormat,omitempty" tf:"datetime_format,omitempty"`

	// Specifies the encoding of the log file so that the file can be read correctly. The default is utf_8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// Specifies log files that you want to push to CloudWatch Logs. File can point to a specific file or multiple files (by using wild card characters such as /var/log/system.log*).
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// Specifies the range of lines for identifying a file. The valid values are one number, or two dash-delimited numbers, such as 1, 2-5. The default value is 1.
	FileFingerprintLines *string `json:"fileFingerprintLines,omitempty" tf:"file_fingerprint_lines,omitempty"`

	// Specifies where to start to read data (start_of_file or end_of_file). The default is start_of_file.
	InitialPosition *string `json:"initialPosition,omitempty" tf:"initial_position,omitempty"`

	// Specifies the destination log group. A log group is created automatically if it doesn't already exist.
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// Specifies the pattern for identifying the start of a log message.
	MultilineStartPattern *string `json:"multilineStartPattern,omitempty" tf:"multiline_start_pattern,omitempty"`

	// Specifies the time zone of log event time stamps.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type LogStreamsParameters struct {

	// Specifies the max number of log events in a batch, up to 10000. The default value is 1000.
	// +kubebuilder:validation:Optional
	BatchCount *int64 `json:"batchCount,omitempty" tf:"batch_count,omitempty"`

	// Specifies the maximum size of log events in a batch, in bytes, up to 1048576 bytes. The default value is 32768 bytes.
	// +kubebuilder:validation:Optional
	BatchSize *int64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// Specifies the time duration for the batching of log events. The minimum value is 5000 and default value is 5000.
	// +kubebuilder:validation:Optional
	BufferDuration *int64 `json:"bufferDuration,omitempty" tf:"buffer_duration,omitempty"`

	// Specifies how the timestamp is extracted from logs. For more information, see the CloudWatch Logs Agent Reference (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AgentReference.html).
	// +kubebuilder:validation:Optional
	DatetimeFormat *string `json:"datetimeFormat,omitempty" tf:"datetime_format,omitempty"`

	// Specifies the encoding of the log file so that the file can be read correctly. The default is utf_8.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// Specifies log files that you want to push to CloudWatch Logs. File can point to a specific file or multiple files (by using wild card characters such as /var/log/system.log*).
	// +kubebuilder:validation:Optional
	File *string `json:"file" tf:"file,omitempty"`

	// Specifies the range of lines for identifying a file. The valid values are one number, or two dash-delimited numbers, such as 1, 2-5. The default value is 1.
	// +kubebuilder:validation:Optional
	FileFingerprintLines *string `json:"fileFingerprintLines,omitempty" tf:"file_fingerprint_lines,omitempty"`

	// Specifies where to start to read data (start_of_file or end_of_file). The default is start_of_file.
	// +kubebuilder:validation:Optional
	InitialPosition *string `json:"initialPosition,omitempty" tf:"initial_position,omitempty"`

	// Specifies the destination log group. A log group is created automatically if it doesn't already exist.
	// +kubebuilder:validation:Optional
	LogGroupName *string `json:"logGroupName" tf:"log_group_name,omitempty"`

	// Specifies the pattern for identifying the start of a log message.
	// +kubebuilder:validation:Optional
	MultilineStartPattern *string `json:"multilineStartPattern,omitempty" tf:"multiline_start_pattern,omitempty"`

	// Specifies the time zone of log event time stamps.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type UpscalingInitParameters struct {

	// Custom Cloudwatch auto scaling alarms, to be used as thresholds. This parameter takes a list of up to five alarm names, which are case sensitive and must be in the same region as the stack.
	Alarms []*string `json:"alarms,omitempty" tf:"alarms,omitempty"`

	// The CPU utilization threshold, as a percent of the available CPU. A value of -1 disables the threshold.
	CPUThreshold *float64 `json:"cpuThreshold,omitempty" tf:"cpu_threshold,omitempty"`

	// The amount of time (in minutes) after a scaling event occurs that AWS OpsWorks Stacks should ignore metrics and suppress additional scaling events.
	IgnoreMetricsTime *int64 `json:"ignoreMetricsTime,omitempty" tf:"ignore_metrics_time,omitempty"`

	// The number of instances to add or remove when the load exceeds a threshold.
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The load threshold. A value of -1 disables the threshold.
	LoadThreshold *float64 `json:"loadThreshold,omitempty" tf:"load_threshold,omitempty"`

	// The memory utilization threshold, as a percent of the available memory. A value of -1 disables the threshold.
	MemoryThreshold *float64 `json:"memoryThreshold,omitempty" tf:"memory_threshold,omitempty"`

	// The amount of time, in minutes, that the load must exceed a threshold before more instances are added or removed.
	ThresholdsWaitTime *int64 `json:"thresholdsWaitTime,omitempty" tf:"thresholds_wait_time,omitempty"`
}

type UpscalingObservation struct {

	// Custom Cloudwatch auto scaling alarms, to be used as thresholds. This parameter takes a list of up to five alarm names, which are case sensitive and must be in the same region as the stack.
	Alarms []*string `json:"alarms,omitempty" tf:"alarms,omitempty"`

	// The CPU utilization threshold, as a percent of the available CPU. A value of -1 disables the threshold.
	CPUThreshold *float64 `json:"cpuThreshold,omitempty" tf:"cpu_threshold,omitempty"`

	// The amount of time (in minutes) after a scaling event occurs that AWS OpsWorks Stacks should ignore metrics and suppress additional scaling events.
	IgnoreMetricsTime *int64 `json:"ignoreMetricsTime,omitempty" tf:"ignore_metrics_time,omitempty"`

	// The number of instances to add or remove when the load exceeds a threshold.
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The load threshold. A value of -1 disables the threshold.
	LoadThreshold *float64 `json:"loadThreshold,omitempty" tf:"load_threshold,omitempty"`

	// The memory utilization threshold, as a percent of the available memory. A value of -1 disables the threshold.
	MemoryThreshold *float64 `json:"memoryThreshold,omitempty" tf:"memory_threshold,omitempty"`

	// The amount of time, in minutes, that the load must exceed a threshold before more instances are added or removed.
	ThresholdsWaitTime *int64 `json:"thresholdsWaitTime,omitempty" tf:"thresholds_wait_time,omitempty"`
}

type UpscalingParameters struct {

	// Custom Cloudwatch auto scaling alarms, to be used as thresholds. This parameter takes a list of up to five alarm names, which are case sensitive and must be in the same region as the stack.
	// +kubebuilder:validation:Optional
	Alarms []*string `json:"alarms,omitempty" tf:"alarms,omitempty"`

	// The CPU utilization threshold, as a percent of the available CPU. A value of -1 disables the threshold.
	// +kubebuilder:validation:Optional
	CPUThreshold *float64 `json:"cpuThreshold,omitempty" tf:"cpu_threshold,omitempty"`

	// The amount of time (in minutes) after a scaling event occurs that AWS OpsWorks Stacks should ignore metrics and suppress additional scaling events.
	// +kubebuilder:validation:Optional
	IgnoreMetricsTime *int64 `json:"ignoreMetricsTime,omitempty" tf:"ignore_metrics_time,omitempty"`

	// The number of instances to add or remove when the load exceeds a threshold.
	// +kubebuilder:validation:Optional
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The load threshold. A value of -1 disables the threshold.
	// +kubebuilder:validation:Optional
	LoadThreshold *float64 `json:"loadThreshold,omitempty" tf:"load_threshold,omitempty"`

	// The memory utilization threshold, as a percent of the available memory. A value of -1 disables the threshold.
	// +kubebuilder:validation:Optional
	MemoryThreshold *float64 `json:"memoryThreshold,omitempty" tf:"memory_threshold,omitempty"`

	// The amount of time, in minutes, that the load must exceed a threshold before more instances are added or removed.
	// +kubebuilder:validation:Optional
	ThresholdsWaitTime *int64 `json:"thresholdsWaitTime,omitempty" tf:"thresholds_wait_time,omitempty"`
}

// CustomLayerSpec defines the desired state of CustomLayer
type CustomLayerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomLayerParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider CustomLayerInitParameters `json:"initProvider,omitempty"`
}

// CustomLayerStatus defines the observed state of CustomLayer.
type CustomLayerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomLayerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CustomLayer is the Schema for the CustomLayers API. Provides an OpsWorks custom layer resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CustomLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.shortName) || (has(self.initProvider) && has(self.initProvider.shortName))",message="spec.forProvider.shortName is a required parameter"
	Spec   CustomLayerSpec   `json:"spec"`
	Status CustomLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomLayerList contains a list of CustomLayers
type CustomLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomLayer `json:"items"`
}

// Repository type metadata.
var (
	CustomLayer_Kind             = "CustomLayer"
	CustomLayer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomLayer_Kind}.String()
	CustomLayer_KindAPIVersion   = CustomLayer_Kind + "." + CRDGroupVersion.String()
	CustomLayer_GroupVersionKind = CRDGroupVersion.WithKind(CustomLayer_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomLayer{}, &CustomLayerList{})
}
