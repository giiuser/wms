// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Uploads a part in a multipart upload. In this operation, you provide part data
// in your request. However, you have an option to specify your existing Amazon S3
// object as a data source for the part you are uploading. To upload a part from an
// existing object, you use the UploadPartCopy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html)
// operation. You must initiate a multipart upload (see CreateMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html))
// before you can upload any part. In response to your initiate request, Amazon S3
// returns an upload ID, a unique identifier, that you must include in your upload
// part request. Part numbers can be any number from 1 to 10,000, inclusive. A part
// number uniquely identifies a part and also defines its position within the
// object being created. If you upload a new part using the same part number that
// was used with a previous part, the previously uploaded part is overwritten. Each
// part must be at least 5 MB in size, except the last part. There is no size limit
// on the last part of your multipart upload. To ensure that data is not corrupted
// when traversing the network, specify the Content-MD5 header in the upload part
// request. Amazon S3 checks the part data against the provided MD5 value. If they
// do not match, Amazon S3 returns an error. If the upload request is signed with
// Signature Version 4, then AWS S3 uses the x-amz-content-sha256 header as a
// checksum instead of Content-MD5. For more information see Authenticating
// Requests: Using the Authorization Header (AWS Signature Version 4)
// (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-auth-using-authorization-header.html).
// Note: After you initiate multipart upload and upload one or more parts, you must
// either complete or abort multipart upload in order to stop getting charged for
// storage of the uploaded parts. Only after you either complete or abort multipart
// upload, Amazon S3 frees up the parts storage and stops charging you for the
// parts storage. For more information on multipart uploads, go to Multipart Upload
// Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html) in
// the Amazon S3 User Guide . For information on the permissions required to use
// the multipart upload API, go to Multipart Upload and Permissions
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html) in the
// Amazon S3 User Guide. You can optionally request server-side encryption where
// Amazon S3 encrypts your data as it writes it to disks in its data centers and
// decrypts it for you when you access it. You have the option of providing your
// own encryption key, or you can use the AWS managed encryption keys. If you
// choose to provide your own encryption key, the request headers you provide in
// the request must match the headers you used in the request to initiate the
// upload by using CreateMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html).
// For more information, go to Using Server-Side Encryption
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html)
// in the Amazon S3 User Guide. Server-side encryption is supported by the S3
// Multipart Upload actions. Unless you are using a customer-provided encryption
// key, you don't need to specify the encryption parameters in each UploadPart
// request. Instead, you only need to specify the server-side encryption parameters
// in the initial Initiate Multipart request. For more information, see
// CreateMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html).
// If you requested server-side encryption using a customer-provided encryption key
// in your initiate multipart upload request, you must provide identical encryption
// information in each part upload using the following headers.
//
// *
// x-amz-server-side-encryption-customer-algorithm
//
// *
// x-amz-server-side-encryption-customer-key
//
// *
// x-amz-server-side-encryption-customer-key-MD5
//
// Special Errors
//
// * Code:
// NoSuchUpload
//
// * Cause: The specified multipart upload does not exist. The upload
// ID might be invalid, or the multipart upload might have been aborted or
// completed.
//
// * HTTP Status Code: 404 Not Found
//
// * SOAP Fault Code Prefix:
// Client
//
// Related Resources
//
// * CreateMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)
//
// *
// CompleteMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html)
//
// *
// AbortMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html)
//
// *
// ListParts
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html)
//
// *
// ListMultipartUploads
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html)
func (c *Client) UploadPart(ctx context.Context, params *UploadPartInput, optFns ...func(*Options)) (*UploadPartOutput, error) {
	if params == nil {
		params = &UploadPartInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UploadPart", params, optFns, addOperationUploadPartMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UploadPartOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UploadPartInput struct {

	// The name of the bucket to which the multipart upload was initiated. When using
	// this action with an access point, you must direct requests to the access point
	// hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// action with an access point through the AWS SDKs, you provide the access point
	// ARN in place of the bucket name. For more information about access point ARNs,
	// see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide. When using this action with Amazon S3 on Outposts,
	// you must direct requests to the S3 on Outposts hostname. The S3 on Outposts
	// hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com. When using
	// this action using S3 on Outposts through the AWS SDKs, you provide the Outposts
	// bucket ARN in place of the bucket name. For more information about S3 on
	// Outposts ARNs, see Using S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in the
	// Amazon S3 User Guide.
	//
	// This member is required.
	Bucket *string

	// Object key for which the multipart upload was initiated.
	//
	// This member is required.
	Key *string

	// Part number of part being uploaded. This is a positive integer between 1 and
	// 10,000.
	//
	// This member is required.
	PartNumber int32

	// Upload ID identifying the multipart upload whose part is being uploaded.
	//
	// This member is required.
	UploadId *string

	// Object data.
	Body io.Reader

	// Size of the body in bytes. This parameter is useful when the size of the body
	// cannot be determined automatically.
	ContentLength int64

	// The base64-encoded 128-bit MD5 digest of the part data. This parameter is
	// auto-populated when using the command from the CLI. This parameter is required
	// if object lock parameters are specified.
	ContentMD5 *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header. This must be the same
	// encryption key specified in the initiate multipart upload request.
	SSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string
}

type UploadPartOutput struct {

	// Indicates whether the multipart upload uses an S3 Bucket Key for server-side
	// encryption with AWS KMS (SSE-KMS).
	BucketKeyEnabled bool

	// Entity tag for the uploaded object.
	ETag *string

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm used.
	SSECustomerAlgorithm *string

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) was used for the object.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256, aws:kms).
	ServerSideEncryption types.ServerSideEncryption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUploadPartMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpUploadPart{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUploadPart{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUploadPartValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUploadPart(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addUploadPartUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUploadPart(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "UploadPart",
	}
}

// getUploadPartBucketMember returns a pointer to string denoting a provided bucket
// member valueand a boolean indicating if the input has a modeled bucket name,
func getUploadPartBucketMember(input interface{}) (*string, bool) {
	in := input.(*UploadPartInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addUploadPartUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getUploadPartBucketMember,
		},
		UsePathStyle:            options.UsePathStyle,
		UseAccelerate:           options.UseAccelerate,
		SupportsAccelerate:      true,
		TargetS3ObjectLambda:    false,
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
