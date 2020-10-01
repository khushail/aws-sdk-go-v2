// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Information about the batch policy.
type BatchPolicy struct {

	// The amount of time, in seconds, to wait for the batch to complete.  </p> <p>If a
	// batch times out, and there are pending requests that were failing due to an
	// internal failure (like <code>InternalServiceError</code>), they will be moved to
	// the failed list and the batch status will be <code>Failed</code>. If the pending
	// requests were failing for any other reason, the failed pending requests will be
	// moved to the failed list and the batch status will be <code>TimedOut</code>.
	// </p>
	TimeoutInSeconds *int64

	// The number of active simulation jobs create as part of the batch that can be in
	// an active state at the same time. Active states include: Pending,Preparing,
	// Running, Restarting, RunningFailed and Terminating. All other states are
	// terminal states.
	MaxConcurrency *int32
}

// Compute information for the simulation job.
type Compute struct {

	// The simulation unit limit. Your simulation is allocated CPU and memory
	// proportional to the supplied simulation unit limit. A simulation unit is 1 vcpu
	// and 2GB of memory. You are only billed for the SU utilization you consume up to
	// the maximim value provided.
	SimulationUnitLimit *int32
}

// Compute information for the simulation job
type ComputeResponse struct {

	// The simulation unit limit. Your simulation is allocated CPU and memory
	// proportional to the supplied simulation unit limit. A simulation unit is 1 vcpu
	// and 2GB of memory. You are only billed for the SU utilization you consume up to
	// the maximim value provided.
	SimulationUnitLimit *int32
}

// Information about a data source.
type DataSource struct {

	// The name of the data source.
	Name *string

	// The S3 bucket where the data files are located.
	S3Bucket *string

	// The list of S3 keys identifying the data source files.
	S3Keys []*S3KeyOutput
}

// Information about a data source.
type DataSourceConfig struct {

	// The name of the data source.
	//
	// This member is required.
	Name *string

	// The S3 bucket where the data files are located.
	//
	// This member is required.
	S3Bucket *string

	// The list of S3 keys identifying the data source files.
	//
	// This member is required.
	S3Keys []*string
}

// Information about a deployment application configuration.
type DeploymentApplicationConfig struct {

	// The Amazon Resource Name (ARN) of the robot application.
	//
	// This member is required.
	Application *string

	// The launch configuration.
	//
	// This member is required.
	LaunchConfig *DeploymentLaunchConfig

	// The version of the application.
	//
	// This member is required.
	ApplicationVersion *string
}

// Information about a deployment configuration.
type DeploymentConfig struct {

	// The percentage of deployments that need to fail before stopping deployment.
	FailureThresholdPercentage *int32

	// The download condition file.
	DownloadConditionFile *S3Object

	// The amount of time, in seconds, to wait for deployment to a single robot to
	// complete. Choose a time between 1 minute and 7 days. The default is 5 hours.
	RobotDeploymentTimeoutInSeconds *int64

	// The percentage of robots receiving the deployment at the same time.
	ConcurrentDeploymentPercentage *int32
}

// Information about a deployment job.
type DeploymentJob struct {

	// A short description of the reason why the deployment job failed.
	FailureReason *string

	// The status of the deployment job.
	Status DeploymentStatus

	// The deployment job failure code.
	FailureCode DeploymentJobErrorCode

	// The time, in milliseconds since the epoch, when the deployment job was created.
	CreatedAt *time.Time

	// The deployment configuration.
	DeploymentConfig *DeploymentConfig

	// The deployment application configuration.
	DeploymentApplicationConfigs []*DeploymentApplicationConfig

	// The Amazon Resource Name (ARN) of the deployment job.
	Arn *string

	// The Amazon Resource Name (ARN) of the fleet.
	Fleet *string
}

// Configuration information for a deployment launch.
type DeploymentLaunchConfig struct {

	// The deployment pre-launch file. This file will be executed prior to the launch
	// file.
	PreLaunchFile *string

	// The deployment post-launch file. This file will be executed after the launch
	// file.
	PostLaunchFile *string

	// The package name.
	//
	// This member is required.
	PackageName *string

	// The launch file name.
	//
	// This member is required.
	LaunchFile *string

	// An array of key/value pairs specifying environment variables for the robot
	// application
	EnvironmentVariables map[string]*string
}

// Information about a failed create simulation job request.
type FailedCreateSimulationJobRequest struct {

	// The time, in milliseconds since the epoch, when the simulation job batch failed.
	FailedAt *time.Time

	// The simulation job request.
	Request *SimulationJobRequest

	// The failure code.
	FailureCode SimulationJobErrorCode

	// The failure reason of the simulation job request.
	FailureReason *string
}

// Information about a filter.
type Filter struct {

	// The name of the filter.
	Name *string

	// A list of values.
	Values []*string
}

// Information about a fleet.
type Fleet struct {

	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time

	// The status of the last fleet deployment.
	LastDeploymentStatus DeploymentStatus

	// The Amazon Resource Name (ARN) of the fleet.
	Arn *string

	// The name of the fleet.
	Name *string

	// The time of the last deployment.
	LastDeploymentTime *time.Time

	// The Amazon Resource Name (ARN) of the last deployment job.
	LastDeploymentJob *string
}

// Information about a launch configuration.
type LaunchConfig struct {

	// The package name.
	//
	// This member is required.
	PackageName *string

	// The environment variables for the application launch.
	EnvironmentVariables map[string]*string

	// The launch file name.
	//
	// This member is required.
	LaunchFile *string

	// Boolean indicating whether a streaming session will be configured for the
	// application. If True, AWS RoboMaker will configure a connection so you can
	// interact with your application as it is running in the simulation. You must
	// configure and luanch the component. It must have a graphical user interface.
	StreamUI *bool

	// The port forwarding configuration.
	PortForwardingConfig *PortForwardingConfig
}

// The logging configuration.
type LoggingConfig struct {

	// A boolean indicating whether to record all ROS topics.
	//
	// This member is required.
	RecordAllRosTopics *bool
}

// Describes a network interface.
type NetworkInterface struct {

	// The IPv4 address of the network interface within the subnet.
	PrivateIpAddress *string

	// The ID of the network interface.
	NetworkInterfaceId *string

	// The IPv4 public address of the network interface.
	PublicIpAddress *string
}

// The output location.
type OutputLocation struct {

	// The S3 bucket for output.
	S3Bucket *string

	// The S3 folder in the s3Bucket where output files will be placed.
	S3Prefix *string
}

// Configuration information for port forwarding.
type PortForwardingConfig struct {

	// The port mappings for the configuration.
	PortMappings []*PortMapping
}

// An object representing a port mapping.
type PortMapping struct {

	// The port number on the simulation job instance to use as a remote connection
	// point.
	//
	// This member is required.
	JobPort *int32

	// The port number on the application.
	//
	// This member is required.
	ApplicationPort *int32

	// A Boolean indicating whether to enable this port mapping on public IP.
	EnableOnPublicIp *bool
}

// Information about the progress of a deployment job.
type ProgressDetail struct {

	// Precentage of the step that is done. This currently only applies to the
	// Downloading/Extracting step of the deployment. It is empty for other steps.
	PercentDone *float32

	// The current progress status. Validating Validating the deployment.
	// DownloadingExtracting Downloading and extracting the bundle on the robot.
	// ExecutingPreLaunch Executing pre-launch script(s) if provided. Launching
	// Launching the robot application. ExecutingPostLaunch Executing post-launch
	// script(s) if provided. Finished Deployment is complete.
	CurrentProgress RobotDeploymentStep

	// Estimated amount of time in seconds remaining in the step. This currently only
	// applies to the Downloading/Extracting step of the deployment. It is empty for
	// other steps.
	EstimatedTimeRemainingSeconds *int32

	// The Amazon Resource Name (ARN) of the deployment job.
	TargetResource *string
}

// Information about a rendering engine.
type RenderingEngine struct {

	// The version of the rendering engine.
	Version *string

	// The name of the rendering engine.
	Name RenderingEngineType
}

// Information about a robot.
type Robot struct {

	// The Amazon Resource Name (ARN) of the last deployment job.
	LastDeploymentJob *string

	// The architecture of the robot.
	Architecture Architecture

	// The time of the last deployment.
	LastDeploymentTime *time.Time

	// The status of the robot.
	Status RobotStatus

	// The time, in milliseconds since the epoch, when the robot was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the fleet.
	FleetArn *string

	// The Amazon Resource Name (ARN) of the robot.
	Arn *string

	// The Greengrass group associated with the robot.
	GreenGrassGroupId *string

	// The name of the robot.
	Name *string
}

// Application configuration information for a robot.
type RobotApplicationConfig struct {

	// The version of the robot application.
	ApplicationVersion *string

	// The launch configuration for the robot application.
	//
	// This member is required.
	LaunchConfig *LaunchConfig

	// The application information for the robot application.
	//
	// This member is required.
	Application *string
}

// Summary information for a robot application.
type RobotApplicationSummary struct {

	// The version of the robot application.
	Version *string

	// The Amazon Resource Name (ARN) of the robot.
	Arn *string

	// The time, in milliseconds since the epoch, when the robot application was last
	// updated.
	LastUpdatedAt *time.Time

	// The name of the robot application.
	Name *string

	// Information about a robot software suite (ROS distribution).
	RobotSoftwareSuite *RobotSoftwareSuite
}

// Information about a robot deployment.
type RobotDeployment struct {

	// The status of the robot deployment.
	Status RobotStatus

	// A short description of the reason why the robot deployment failed.
	FailureReason *string

	// The time, in milliseconds since the epoch, when the deployment was started.
	DeploymentStartTime *time.Time

	// The robot deployment failure code.
	FailureCode DeploymentJobErrorCode

	// The robot deployment Amazon Resource Name (ARN).
	Arn *string

	// Information about how the deployment is progressing.
	ProgressDetail *ProgressDetail

	// The time, in milliseconds since the epoch, when the deployment finished.
	DeploymentFinishTime *time.Time
}

// Information about a robot software suite (ROS distribution).
type RobotSoftwareSuite struct {

	// The version of the robot software suite (ROS distribution).
	Version RobotSoftwareSuiteVersionType

	// The name of the robot software suite (ROS distribution).
	Name RobotSoftwareSuiteType
}

// Information about S3 keys.
type S3KeyOutput struct {

	// The S3 key.
	S3Key *string

	// The etag for the object.
	Etag *string
}

// Information about an S3 object.
type S3Object struct {

	// The etag of the object.
	Etag *string

	// The key of the object.
	//
	// This member is required.
	Key *string

	// The bucket containing the object.
	//
	// This member is required.
	Bucket *string
}

// Information about a simulation application configuration.
type SimulationApplicationConfig struct {

	// The launch configuration for the simulation application.
	//
	// This member is required.
	LaunchConfig *LaunchConfig

	// The version of the simulation application.
	ApplicationVersion *string

	// The application information for the simulation application.
	//
	// This member is required.
	Application *string
}

// Summary information for a simulation application.
type SimulationApplicationSummary struct {

	// The name of the simulation application.
	Name *string

	// The Amazon Resource Name (ARN) of the simulation application.
	Arn *string

	// Information about a simulation software suite.
	SimulationSoftwareSuite *SimulationSoftwareSuite

	// The version of the simulation application.
	Version *string

	// Information about a robot software suite (ROS distribution).
	RobotSoftwareSuite *RobotSoftwareSuite

	// The time, in milliseconds since the epoch, when the simulation application was
	// last updated.
	LastUpdatedAt *time.Time
}

// Information about a simulation job.
type SimulationJob struct {

	// The logging configuration.
	LoggingConfig *LoggingConfig

	// Information about a network interface.
	NetworkInterface *NetworkInterface

	// The maximum simulation job duration in seconds. The value must be 8 days
	// (691,200 seconds) or less.
	MaxJobDurationInSeconds *int64

	// The simulation job execution duration in milliseconds.
	SimulationTimeMillis *int64

	// The IAM role that allows the simulation instance to call the AWS APIs that are
	// specified in its associated policies on your behalf. This is how credentials are
	// passed in to your simulation job.
	IamRole *string

	// VPC configuration information.
	VpcConfig *VPCConfigResponse

	// Status of the simulation job.
	Status SimulationJobStatus

	// The failure behavior the simulation job. Continue Restart the simulation job in
	// the same host instance. Fail Stop the simulation job and terminate the instance.
	FailureBehavior FailureBehavior

	// A unique identifier for this SimulationJob request.
	ClientRequestToken *string

	// A list of simulation applications.
	SimulationApplications []*SimulationApplicationConfig

	// Compute information for the simulation job
	Compute *ComputeResponse

	// The name of the simulation job.
	Name *string

	// The Amazon Resource Name (ARN) of the simulation job.
	Arn *string

	// The data sources for the simulation job.
	DataSources []*DataSource

	// A list of robot applications.
	RobotApplications []*RobotApplicationConfig

	// The time, in milliseconds since the epoch, when the simulation job was last
	// updated.
	LastUpdatedAt *time.Time

	// The reason why the simulation job failed.
	FailureReason *string

	// A map that contains tag keys and tag values that are attached to the simulation
	// job.
	Tags map[string]*string

	// The time, in milliseconds since the epoch, when the simulation job was last
	// started.
	LastStartedAt *time.Time

	// The failure code of the simulation job if it failed.
	FailureCode SimulationJobErrorCode

	// Location for output files generated by the simulation job.
	OutputLocation *OutputLocation
}

// Information about a simulation job batch.
type SimulationJobBatchSummary struct {

	// The time, in milliseconds since the epoch, when the simulation job batch was
	// created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the batch.
	Arn *string

	// The number of failed simulation job requests.
	FailedRequestCount *int32

	// The number of pending simulation job requests.
	PendingRequestCount *int32

	// The status of the simulation job batch. Pending The simulation job batch request
	// is pending. InProgress The simulation job batch is in progress. Failed The
	// simulation job batch failed. One or more simulation job requests could not be
	// completed due to an internal failure (like InternalServiceError). See
	// failureCode and failureReason for more information. Completed The simulation
	// batch job completed. A batch is complete when (1) there are no pending
	// simulation job requests in the batch and none of the failed simulation job
	// requests are due to InternalServiceError and (2) when all created simulation
	// jobs have reached a terminal state (for example, Completed or Failed). Canceled
	// The simulation batch job was cancelled. Canceling The simulation batch job is
	// being cancelled. Completing The simulation batch job is completing. TimingOut
	// The simulation job batch is timing out. If a batch timing out, and there are
	// pending requests that were failing due to an internal failure (like
	// InternalServiceError), the batch status will be Failed. If there are no such
	// failing request, the batch status will be TimedOut. TimedOut The simulation
	// batch job timed out.
	Status SimulationJobBatchStatus

	// The time, in milliseconds since the epoch, when the simulation job batch was
	// last updated.
	LastUpdatedAt *time.Time

	// The number of created simulation job requests.
	CreatedRequestCount *int32
}

// Information about a simulation job request.
type SimulationJobRequest struct {

	// The maximum simulation job duration in seconds. The value must be 8 days
	// (691,200 seconds) or less.
	//
	// This member is required.
	MaxJobDurationInSeconds *int64

	// Compute information for the simulation job
	Compute *Compute

	// Boolean indicating whether to use default simulation tool applications.
	UseDefaultApplications *bool

	// The logging configuration.
	LoggingConfig *LoggingConfig

	// Specify data sources to mount read-only files from S3 into your simulation.
	// These files are available under /opt/robomaker/datasources/data_source_name.
	// There is a limit of 100 files and a combined size of 25GB for all
	// DataSourceConfig objects.
	DataSources []*DataSourceConfig

	// The robot applications to use in the simulation job.
	RobotApplications []*RobotApplicationConfig

	// The output location.
	OutputLocation *OutputLocation

	// A map that contains tag keys and tag values that are attached to the simulation
	// job request.
	Tags map[string]*string

	// The simulation applications to use in the simulation job.
	SimulationApplications []*SimulationApplicationConfig

	// The failure behavior the simulation job. Continue Restart the simulation job in
	// the same host instance. Fail Stop the simulation job and terminate the instance.
	FailureBehavior FailureBehavior

	// If your simulation job accesses resources in a VPC, you provide this parameter
	// identifying the list of security group IDs and subnet IDs. These must belong to
	// the same VPC. You must provide at least one security group and two subnet IDs.
	VpcConfig *VPCConfig

	// The IAM role name that allows the simulation instance to call the AWS APIs that
	// are specified in its associated policies on your behalf. This is how credentials
	// are passed in to your simulation job.
	IamRole *string
}

// Summary information for a simulation job.
type SimulationJobSummary struct {

	// The name of the simulation job.
	Name *string

	// A list of simulation job simulation application names.
	SimulationApplicationNames []*string

	// A list of simulation job robot application names.
	RobotApplicationNames []*string

	// The time, in milliseconds since the epoch, when the simulation job was last
	// updated.
	LastUpdatedAt *time.Time

	// The names of the data sources.
	DataSourceNames []*string

	// The status of the simulation job.
	Status SimulationJobStatus

	// The Amazon Resource Name (ARN) of the simulation job.
	Arn *string
}

// Information about a simulation software suite.
type SimulationSoftwareSuite struct {

	// The name of the simulation software suite.
	Name SimulationSoftwareSuiteType

	// The version of the simulation software suite.
	Version *string
}

// Information about a source.
type Source struct {

	// The s3 object key.
	S3Key *string

	// A hash of the object specified by s3Bucket and s3Key.
	Etag *string

	// The s3 bucket name.
	S3Bucket *string

	// The taget processor architecture for the application.
	Architecture Architecture
}

// Information about a source configuration.
type SourceConfig struct {

	// The s3 object key.
	S3Key *string

	// The Amazon S3 bucket name.
	S3Bucket *string

	// The target processor architecture for the application.
	Architecture Architecture
}

// If your simulation job accesses resources in a VPC, you provide this parameter
// identifying the list of security group IDs and subnet IDs. These must belong to
// the same VPC. You must provide at least one security group and two subnet IDs.
type VPCConfig struct {

	// A list of one or more subnet IDs in your VPC.
	//
	// This member is required.
	Subnets []*string

	// A boolean indicating whether to assign a public IP address.
	AssignPublicIp *bool

	// A list of one or more security groups IDs in your VPC.
	SecurityGroups []*string
}

// VPC configuration associated with your simulation job.
type VPCConfigResponse struct {

	// The VPC ID associated with your simulation job.
	VpcId *string

	// A list of security group IDs associated with the simulation job.
	SecurityGroups []*string

	// A list of subnet IDs associated with the simulation job.
	Subnets []*string

	// A boolean indicating if a public IP was assigned.
	AssignPublicIp *bool
}