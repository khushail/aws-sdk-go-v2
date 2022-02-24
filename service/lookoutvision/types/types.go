// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The description for a dataset. For more information, see DescribeDataset.
type DatasetDescription struct {

	// The Unix timestamp for the time and date that the dataset was created.
	CreationTimestamp *time.Time

	// The type of the dataset. The value train represents a training dataset or single
	// dataset project. The value test represents a test dataset.
	DatasetType *string

	// Statistics about the images in a dataset.
	ImageStats *DatasetImageStats

	// The Unix timestamp for the date and time that the dataset was last updated.
	LastUpdatedTimestamp *time.Time

	// The name of the project that contains the dataset.
	ProjectName *string

	// The status of the dataset.
	Status DatasetStatus

	// The status message for the dataset.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Location information about a manifest file. You can use a manifest file to
// create a dataset.
type DatasetGroundTruthManifest struct {

	// The S3 bucket location for the manifest file.
	S3Object *InputS3Object

	noSmithyDocumentSerde
}

// Statistics about the images in a dataset.
type DatasetImageStats struct {

	// the total number of images labeled as an anomaly.
	Anomaly *int32

	// The total number of labeled images.
	Labeled *int32

	// The total number of images labeled as normal.
	Normal *int32

	// The total number of images in the dataset.
	Total *int32

	noSmithyDocumentSerde
}

// Summary information for an Amazon Lookout for Vision dataset. For more
// information, see DescribeDataset and ProjectDescription.
type DatasetMetadata struct {

	// The Unix timestamp for the date and time that the dataset was created.
	CreationTimestamp *time.Time

	// The type of the dataset.
	DatasetType *string

	// The status for the dataset.
	Status DatasetStatus

	// The status message for the dataset.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Information about the location of a manifest file that Amazon Lookout for Vision
// uses to to create a dataset.
type DatasetSource struct {

	// Location information for the manifest file.
	GroundTruthManifest *DatasetGroundTruthManifest

	noSmithyDocumentSerde
}

// The prediction results from a call to DetectAnomalies.
type DetectAnomalyResult struct {

	// The confidence that Amazon Lookout for Vision has in the accuracy of the
	// prediction.
	Confidence *float32

	// True if the image contains an anomaly, otherwise false.
	IsAnomalous bool

	// The source of the image that was analyzed. direct means that the images was
	// supplied from the local computer. No other values are supported.
	Source *ImageSource

	noSmithyDocumentSerde
}

// Configuration information for the AWS IoT Greengrass component created in a
// model packaging job. For more information, see StartModelPackagingJob. You can't
// specify a component with the same ComponentName and Componentversion as an
// existing component with the same component name and component version.
type GreengrassConfiguration struct {

	// A name for the AWS IoT Greengrass component.
	//
	// This member is required.
	ComponentName *string

	// An S3 location in which Lookout for Vision stores the component artifacts.
	//
	// This member is required.
	S3OutputLocation *S3Location

	// Additional compiler options for the Greengrass component. Currently, only NVIDIA
	// Graphics Processing Units (GPU) are supported. If you specify TargetPlatform,
	// you must specify CompilerOptions. If you specify TargetDevice, don't specify
	// CompilerOptions. For more information, see Compiler options in the Amazon
	// Lookout for Vision Developer Guide.
	CompilerOptions *string

	// A description for the AWS IoT Greengrass component.
	ComponentDescription *string

	// A Version for the AWS IoT Greengrass component. If you don't provide a value, a
	// default value of  Model Version.0.0 is used.
	ComponentVersion *string

	// A set of tags (key-value pairs) that you want to attach to the AWS IoT
	// Greengrass component.
	Tags []Tag

	// The target device for the model. Currently the only supported value is
	// jetson_xavier. If you specify TargetDevice, you can't specify TargetPlatform.
	TargetDevice TargetDevice

	// The target platform for the model. If you specify TargetPlatform, you can't
	// specify TargetDevice.
	TargetPlatform *TargetPlatform

	noSmithyDocumentSerde
}

// Information about the AWS IoT Greengrass component created by a model packaging
// job.
type GreengrassOutputDetails struct {

	// The name of the component.
	ComponentName *string

	// The version of the component.
	ComponentVersion *string

	// The Amazon Resource Name (ARN) of the component.
	ComponentVersionArn *string

	noSmithyDocumentSerde
}

// The source for an image.
type ImageSource struct {

	// The type of the image.
	Type *string

	noSmithyDocumentSerde
}

// Amazon S3 Location information for an input manifest file.
type InputS3Object struct {

	// The Amazon S3 bucket that contains the manifest.
	//
	// This member is required.
	Bucket *string

	// The name and location of the manifest file withiin the bucket.
	//
	// This member is required.
	Key *string

	// The version ID of the bucket.
	VersionId *string

	noSmithyDocumentSerde
}

// Describes an Amazon Lookout for Vision model.
type ModelDescription struct {

	// The unix timestamp for the date and time that the model was created.
	CreationTimestamp *time.Time

	// The description for the model.
	Description *string

	// The unix timestamp for the date and time that the evaluation ended.
	EvaluationEndTimestamp *time.Time

	// The S3 location where Amazon Lookout for Vision saves the manifest file that was
	// used to test the trained model and generate the performance scores.
	EvaluationManifest *OutputS3Object

	// The S3 location where Amazon Lookout for Vision saves the performance metrics.
	EvaluationResult *OutputS3Object

	// The identifer for the AWS Key Management Service (AWS KMS) key that was used to
	// encrypt the model during training.
	KmsKeyId *string

	// The Amazon Resource Name (ARN) of the model.
	ModelArn *string

	// The version of the model
	ModelVersion *string

	// The S3 location where Amazon Lookout for Vision saves model training files.
	OutputConfig *OutputConfig

	// Performance metrics for the model. Created during training.
	Performance *ModelPerformance

	// The status of the model.
	Status ModelStatus

	// The status message for the model.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Describes an Amazon Lookout for Vision model.
type ModelMetadata struct {

	// The unix timestamp for the date and time that the model was created.
	CreationTimestamp *time.Time

	// The description for the model.
	Description *string

	// The Amazon Resource Name (ARN) of the model.
	ModelArn *string

	// The version of the model.
	ModelVersion *string

	// Performance metrics for the model. Not available until training has successfully
	// completed.
	Performance *ModelPerformance

	// The status of the model.
	Status ModelStatus

	// The status message for the model.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Configuration information for a Amazon Lookout for Vision model packaging job.
// For more information, see StartModelPackagingJob.
type ModelPackagingConfiguration struct {

	// Configuration information for the AWS IoT Greengrass component in a model
	// packaging job.
	//
	// This member is required.
	Greengrass *GreengrassConfiguration

	noSmithyDocumentSerde
}

// Information about a model packaging job. For more information, see
// DescribeModelPackagingJob.
type ModelPackagingDescription struct {

	// The Unix timestamp for the time and date that the model packaging job was
	// created.
	CreationTimestamp *time.Time

	// The name of the model packaging job.
	JobName *string

	// The Unix timestamp for the time and date that the model packaging job was last
	// updated.
	LastUpdatedTimestamp *time.Time

	// The configuration information used in the model packaging job.
	ModelPackagingConfiguration *ModelPackagingConfiguration

	// The description for the model packaging job.
	ModelPackagingJobDescription *string

	// The AWS service used to package the job. Currently Lookout for Vision can
	// package jobs with AWS IoT Greengrass.
	ModelPackagingMethod *string

	// Information about the output of the model packaging job. For more information,
	// see DescribeModelPackagingJob.
	ModelPackagingOutputDetails *ModelPackagingOutputDetails

	// The version of the model used in the model packaging job.
	ModelVersion *string

	// The name of the project that's associated with a model that's in the model
	// package.
	ProjectName *string

	// The status of the model packaging job.
	Status ModelPackagingJobStatus

	// The status message for the model packaging job.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Metadata for a model packaging job. For more information, see
// ListModelPackagingJobs.
type ModelPackagingJobMetadata struct {

	// The Unix timestamp for the time and date that the model packaging job was
	// created.
	CreationTimestamp *time.Time

	// The name of the model packaging job.
	JobName *string

	// The Unix timestamp for the time and date that the model packaging job was last
	// updated.
	LastUpdatedTimestamp *time.Time

	// The description for the model packaging job.
	ModelPackagingJobDescription *string

	// The AWS service used to package the job. Currently Lookout for Vision can
	// package jobs with AWS IoT Greengrass.
	ModelPackagingMethod *string

	// The version of the model that is in the model package.
	ModelVersion *string

	// The project that contains the model that is in the model package.
	ProjectName *string

	// The status of the model packaging job.
	Status ModelPackagingJobStatus

	// The status message for the model packaging job.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Information about the output from a model packaging job.
type ModelPackagingOutputDetails struct {

	// Information about the AWS IoT Greengrass component in a model packaging job.
	Greengrass *GreengrassOutputDetails

	noSmithyDocumentSerde
}

// Information about the evaluation performance of a trained model.
type ModelPerformance struct {

	// The overall F1 score metric for the trained model.
	F1Score *float32

	// The overall precision metric value for the trained model.
	Precision *float32

	// The overall recall metric value for the trained model.
	Recall *float32

	noSmithyDocumentSerde
}

// The S3 location where Amazon Lookout for Vision saves model training files.
type OutputConfig struct {

	// The S3 location for the output.
	//
	// This member is required.
	S3Location *S3Location

	noSmithyDocumentSerde
}

// The S3 location where Amazon Lookout for Vision saves training output.
type OutputS3Object struct {

	// The bucket that contains the training output.
	//
	// This member is required.
	Bucket *string

	// The location of the training output in the bucket.
	//
	// This member is required.
	Key *string

	noSmithyDocumentSerde
}

// Describe an Amazon Lookout for Vision project. For more information, see
// DescribeProject.
type ProjectDescription struct {

	// The unix timestamp for the date and time that the project was created.
	CreationTimestamp *time.Time

	// A list of datasets in the project.
	Datasets []DatasetMetadata

	// The Amazon Resource Name (ARN) of the project.
	ProjectArn *string

	// The name of the project.
	ProjectName *string

	noSmithyDocumentSerde
}

// Metadata about an Amazon Lookout for Vision project.
type ProjectMetadata struct {

	// The unix timestamp for the date and time that the project was created.
	CreationTimestamp *time.Time

	// The Amazon Resource Name (ARN) of the project.
	ProjectArn *string

	// The name of the project.
	ProjectName *string

	noSmithyDocumentSerde
}

// Information about the location of training output or the output of a model
// packaging job.
type S3Location struct {

	// The S3 bucket that contains the training or model packaging job output. If you
	// are training a model, the bucket must in your AWS account. If you use an S3
	// bucket for a model packaging job, the S3 bucket must be in the same AWS Region
	// and AWS account in which you use AWS IoT Greengrass.
	//
	// This member is required.
	Bucket *string

	// The path of the folder, within the S3 bucket, that contains the output.
	Prefix *string

	noSmithyDocumentSerde
}

// A key and value pair that is attached to the specified Amazon Lookout for Vision
// model.
type Tag struct {

	// The key of the tag that is attached to the specified model.
	//
	// This member is required.
	Key *string

	// The value of the tag that is attached to the specified model.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The platform on which a model runs on an AWS IoT Greengrass core device.
type TargetPlatform struct {

	// The target accelerator for the model. NVIDIA (Nvidia graphics processing unit)
	// is the only accelerator that is currently supported. You must also specify the
	// gpu-code, trt-ver, and cuda-ver compiler options.
	//
	// This member is required.
	Accelerator TargetPlatformAccelerator

	// The target architecture for the model. The currently supported architectures are
	// X86_64 (64-bit version of the x86 instruction set) and ARM_64 (ARMv8 64-bit
	// CPU).
	//
	// This member is required.
	Arch TargetPlatformArch

	// The target operating system for the model. Linux is the only operating system
	// that is currently supported.
	//
	// This member is required.
	Os TargetPlatformOs

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
