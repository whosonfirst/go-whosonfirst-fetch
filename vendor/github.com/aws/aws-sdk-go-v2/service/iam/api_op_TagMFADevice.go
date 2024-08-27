// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds one or more tags to an IAM virtual multi-factor authentication (MFA)
// device. If a tag with the same key name already exists, then that tag is
// overwritten with the new value.
//
// A tag consists of a key name and an associated value. By assigning tags to your
// resources, you can do the following:
//
//   - Administrative grouping and discovery - Attach tags to resources to aid in
//     organization and search. For example, you could search for all resources with
//     the key name Project and the value MyImportantProject. Or search for all
//     resources with the key name Cost Center and the value 41200.
//
//   - Access control - Include tags in IAM user-based and resource-based
//     policies. You can use tags to restrict access to only an IAM virtual MFA device
//     that has a specified tag attached. For examples of policies that show how to use
//     tags to control access, see [Control access using IAM tags]in the IAM User Guide.
//
//   - If any one of the tags is invalid or if you exceed the allowed maximum
//     number of tags, then the entire request fails and the resource is not created.
//     For more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
//   - Amazon Web Services always interprets the tag Value as a single string. If
//     you need to store an array, you can store comma-separated values in the string.
//     However, you must interpret the value in your code.
//
// [Control access using IAM tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func (c *Client) TagMFADevice(ctx context.Context, params *TagMFADeviceInput, optFns ...func(*Options)) (*TagMFADeviceOutput, error) {
	if params == nil {
		params = &TagMFADeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagMFADevice", params, optFns, c.addOperationTagMFADeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagMFADeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagMFADeviceInput struct {

	// The unique identifier for the IAM virtual MFA device to which you want to add
	// tags. For virtual MFA devices, the serial number is the same as the ARN.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	SerialNumber *string

	// The list of tags that you want to attach to the IAM virtual MFA device. Each
	// tag consists of a key name and an associated value.
	//
	// This member is required.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type TagMFADeviceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTagMFADeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpTagMFADevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpTagMFADevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TagMFADevice"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpTagMFADeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTagMFADevice(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opTagMFADevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TagMFADevice",
	}
}
