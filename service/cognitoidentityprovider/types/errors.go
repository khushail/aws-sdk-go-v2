// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// This exception is thrown when a user tries to confirm the account with an email
// or phone number that has already been supplied as an alias from a different
// account. This exception tells user that an account with this email or phone
// already exists.
type AliasExistsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *AliasExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AliasExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AliasExistsException) ErrorCode() string             { return "AliasExistsException" }
func (e *AliasExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when a verification code fails to deliver successfully.
type CodeDeliveryFailureException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *CodeDeliveryFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CodeDeliveryFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CodeDeliveryFailureException) ErrorCode() string             { return "CodeDeliveryFailureException" }
func (e *CodeDeliveryFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown if the provided code doesn't match what the server was
// expecting.
type CodeMismatchException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *CodeMismatchException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CodeMismatchException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CodeMismatchException) ErrorCode() string             { return "CodeMismatchException" }
func (e *CodeMismatchException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown if two or more modifications are happening
// concurrently.
type ConcurrentModificationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the provider is already supported by the user
// pool.
type DuplicateProviderException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DuplicateProviderException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateProviderException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateProviderException) ErrorCode() string             { return "DuplicateProviderException" }
func (e *DuplicateProviderException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when there is a code mismatch and the service fails to
// configure the software token TOTP multi-factor authentication (MFA).
type EnableSoftwareTokenMFAException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *EnableSoftwareTokenMFAException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EnableSoftwareTokenMFAException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EnableSoftwareTokenMFAException) ErrorCode() string {
	return "EnableSoftwareTokenMFAException"
}
func (e *EnableSoftwareTokenMFAException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown if a code has expired.
type ExpiredCodeException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ExpiredCodeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ExpiredCodeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ExpiredCodeException) ErrorCode() string             { return "ExpiredCodeException" }
func (e *ExpiredCodeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when Amazon Cognito encounters a group that already
// exists in the user pool.
type GroupExistsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *GroupExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GroupExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GroupExistsException) ErrorCode() string             { return "GroupExistsException" }
func (e *GroupExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when Amazon Cognito encounters an internal error.
type InternalErrorException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalErrorException) ErrorCode() string             { return "InternalErrorException" }
func (e *InternalErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// This exception is thrown when Amazon Cognito isn't allowed to use your email
// identity. HTTP status code: 400.
type InvalidEmailRoleAccessPolicyException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidEmailRoleAccessPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidEmailRoleAccessPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidEmailRoleAccessPolicyException) ErrorCode() string {
	return "InvalidEmailRoleAccessPolicyException"
}
func (e *InvalidEmailRoleAccessPolicyException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when Amazon Cognito encounters an invalid Lambda
// response.
type InvalidLambdaResponseException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidLambdaResponseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidLambdaResponseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidLambdaResponseException) ErrorCode() string             { return "InvalidLambdaResponseException" }
func (e *InvalidLambdaResponseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the specified OAuth flow is not valid.
type InvalidOAuthFlowException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidOAuthFlowException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidOAuthFlowException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidOAuthFlowException) ErrorCode() string             { return "InvalidOAuthFlowException" }
func (e *InvalidOAuthFlowException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the Amazon Cognito service encounters an invalid
// parameter.
type InvalidParameterException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when Amazon Cognito encounters an invalid password.
type InvalidPasswordException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidPasswordException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPasswordException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPasswordException) ErrorCode() string             { return "InvalidPasswordException" }
func (e *InvalidPasswordException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is returned when the role provided for SMS configuration doesn't
// have permission to publish using Amazon SNS.
type InvalidSmsRoleAccessPolicyException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidSmsRoleAccessPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSmsRoleAccessPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSmsRoleAccessPolicyException) ErrorCode() string {
	return "InvalidSmsRoleAccessPolicyException"
}
func (e *InvalidSmsRoleAccessPolicyException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when the trust relationship is not valid for the role
// provided for SMS configuration. This can happen if you don't trust
// cognito-idp.amazonaws.com or the external ID provided in the role does not match
// what is provided in the SMS configuration for the user pool.
type InvalidSmsRoleTrustRelationshipException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidSmsRoleTrustRelationshipException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSmsRoleTrustRelationshipException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSmsRoleTrustRelationshipException) ErrorCode() string {
	return "InvalidSmsRoleTrustRelationshipException"
}
func (e *InvalidSmsRoleTrustRelationshipException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when the user pool configuration is not valid.
type InvalidUserPoolConfigurationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidUserPoolConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidUserPoolConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidUserPoolConfigurationException) ErrorCode() string {
	return "InvalidUserPoolConfigurationException"
}
func (e *InvalidUserPoolConfigurationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when a user exceeds the limit for a requested Amazon
// Web Services resource.
type LimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when Amazon Cognito can't find a multi-factor
// authentication (MFA) method.
type MFAMethodNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *MFAMethodNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MFAMethodNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MFAMethodNotFoundException) ErrorCode() string             { return "MFAMethodNotFoundException" }
func (e *MFAMethodNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when a user isn't authorized.
type NotAuthorizedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *NotAuthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotAuthorizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotAuthorizedException) ErrorCode() string             { return "NotAuthorizedException" }
func (e *NotAuthorizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when a password reset is required.
type PasswordResetRequiredException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *PasswordResetRequiredException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PasswordResetRequiredException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PasswordResetRequiredException) ErrorCode() string             { return "PasswordResetRequiredException" }
func (e *PasswordResetRequiredException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when a precondition is not met.
type PreconditionNotMetException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *PreconditionNotMetException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PreconditionNotMetException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PreconditionNotMetException) ErrorCode() string             { return "PreconditionNotMetException" }
func (e *PreconditionNotMetException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the Amazon Cognito service can't find the
// requested resource.
type ResourceNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the specified scope doesn't exist.
type ScopeDoesNotExistException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ScopeDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ScopeDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ScopeDoesNotExistException) ErrorCode() string             { return "ScopeDoesNotExistException" }
func (e *ScopeDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the software token time-based one-time password
// (TOTP) multi-factor authentication (MFA) isn't activated for the user pool.
type SoftwareTokenMFANotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SoftwareTokenMFANotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SoftwareTokenMFANotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SoftwareTokenMFANotFoundException) ErrorCode() string {
	return "SoftwareTokenMFANotFoundException"
}
func (e *SoftwareTokenMFANotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the user has made too many failed attempts for a
// given action, such as sign-in.
type TooManyFailedAttemptsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *TooManyFailedAttemptsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyFailedAttemptsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyFailedAttemptsException) ErrorCode() string             { return "TooManyFailedAttemptsException" }
func (e *TooManyFailedAttemptsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the user has made too many requests for a given
// operation.
type TooManyRequestsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *TooManyRequestsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRequestsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRequestsException) ErrorCode() string             { return "TooManyRequestsException" }
func (e *TooManyRequestsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exception that is thrown when the request isn't authorized. This can happen due
// to an invalid access token in the request.
type UnauthorizedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UnauthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthorizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthorizedException) ErrorCode() string             { return "UnauthorizedException" }
func (e *UnauthorizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when Amazon Cognito encounters an unexpected exception
// with Lambda.
type UnexpectedLambdaException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UnexpectedLambdaException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnexpectedLambdaException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnexpectedLambdaException) ErrorCode() string             { return "UnexpectedLambdaException" }
func (e *UnexpectedLambdaException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the specified identifier isn't supported.
type UnsupportedIdentityProviderException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UnsupportedIdentityProviderException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedIdentityProviderException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedIdentityProviderException) ErrorCode() string {
	return "UnsupportedIdentityProviderException"
}
func (e *UnsupportedIdentityProviderException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Exception that is thrown when you attempt to perform an operation that isn't
// enabled for the user pool client.
type UnsupportedOperationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UnsupportedOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedOperationException) ErrorCode() string             { return "UnsupportedOperationException" }
func (e *UnsupportedOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exception that is thrown when an unsupported token is passed to an operation.
type UnsupportedTokenTypeException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UnsupportedTokenTypeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedTokenTypeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedTokenTypeException) ErrorCode() string             { return "UnsupportedTokenTypeException" }
func (e *UnsupportedTokenTypeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request failed because the user is in an unsupported state.
type UnsupportedUserStateException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UnsupportedUserStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedUserStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedUserStateException) ErrorCode() string             { return "UnsupportedUserStateException" }
func (e *UnsupportedUserStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when you're trying to modify a user pool while a user
// import job is in progress for that pool.
type UserImportInProgressException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UserImportInProgressException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserImportInProgressException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserImportInProgressException) ErrorCode() string             { return "UserImportInProgressException" }
func (e *UserImportInProgressException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the Amazon Cognito service encounters a user
// validation exception with the Lambda service.
type UserLambdaValidationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UserLambdaValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserLambdaValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserLambdaValidationException) ErrorCode() string             { return "UserLambdaValidationException" }
func (e *UserLambdaValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when Amazon Cognito encounters a user name that already
// exists in the user pool.
type UsernameExistsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UsernameExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UsernameExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UsernameExistsException) ErrorCode() string             { return "UsernameExistsException" }
func (e *UsernameExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when a user isn't confirmed successfully.
type UserNotConfirmedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UserNotConfirmedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserNotConfirmedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserNotConfirmedException) ErrorCode() string             { return "UserNotConfirmedException" }
func (e *UserNotConfirmedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when a user isn't found.
type UserNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UserNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserNotFoundException) ErrorCode() string             { return "UserNotFoundException" }
func (e *UserNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when user pool add-ons aren't enabled.
type UserPoolAddOnNotEnabledException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UserPoolAddOnNotEnabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserPoolAddOnNotEnabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserPoolAddOnNotEnabledException) ErrorCode() string {
	return "UserPoolAddOnNotEnabledException"
}
func (e *UserPoolAddOnNotEnabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when a user pool tag can't be set or updated.
type UserPoolTaggingException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UserPoolTaggingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserPoolTaggingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserPoolTaggingException) ErrorCode() string             { return "UserPoolTaggingException" }
func (e *UserPoolTaggingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
