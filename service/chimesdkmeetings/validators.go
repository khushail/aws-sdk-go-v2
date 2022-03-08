// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmeetings

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBatchCreateAttendee struct {
}

func (*validateOpBatchCreateAttendee) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchCreateAttendee) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchCreateAttendeeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchCreateAttendeeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateAttendee struct {
}

func (*validateOpCreateAttendee) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAttendee) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAttendeeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAttendeeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateMeeting struct {
}

func (*validateOpCreateMeeting) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateMeeting) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateMeetingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateMeetingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateMeetingWithAttendees struct {
}

func (*validateOpCreateMeetingWithAttendees) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateMeetingWithAttendees) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateMeetingWithAttendeesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateMeetingWithAttendeesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAttendee struct {
}

func (*validateOpDeleteAttendee) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAttendee) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAttendeeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAttendeeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMeeting struct {
}

func (*validateOpDeleteMeeting) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMeeting) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMeetingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMeetingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAttendee struct {
}

func (*validateOpGetAttendee) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAttendee) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAttendeeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAttendeeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMeeting struct {
}

func (*validateOpGetMeeting) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMeeting) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMeetingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMeetingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAttendees struct {
}

func (*validateOpListAttendees) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAttendees) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAttendeesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAttendeesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartMeetingTranscription struct {
}

func (*validateOpStartMeetingTranscription) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartMeetingTranscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartMeetingTranscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartMeetingTranscriptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopMeetingTranscription struct {
}

func (*validateOpStopMeetingTranscription) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopMeetingTranscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopMeetingTranscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopMeetingTranscriptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchCreateAttendeeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchCreateAttendee{}, middleware.After)
}

func addOpCreateAttendeeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAttendee{}, middleware.After)
}

func addOpCreateMeetingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateMeeting{}, middleware.After)
}

func addOpCreateMeetingWithAttendeesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateMeetingWithAttendees{}, middleware.After)
}

func addOpDeleteAttendeeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAttendee{}, middleware.After)
}

func addOpDeleteMeetingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMeeting{}, middleware.After)
}

func addOpGetAttendeeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAttendee{}, middleware.After)
}

func addOpGetMeetingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMeeting{}, middleware.After)
}

func addOpListAttendeesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAttendees{}, middleware.After)
}

func addOpStartMeetingTranscriptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartMeetingTranscription{}, middleware.After)
}

func addOpStopMeetingTranscriptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopMeetingTranscription{}, middleware.After)
}

func validateCreateAttendeeRequestItem(v *types.CreateAttendeeRequestItem) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAttendeeRequestItem"}
	if v.ExternalUserId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExternalUserId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateAttendeeRequestItemList(v []types.CreateAttendeeRequestItem) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAttendeeRequestItemList"}
	for i := range v {
		if err := validateCreateAttendeeRequestItem(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateMeetingWithAttendeesRequestItemList(v []types.CreateAttendeeRequestItem) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateMeetingWithAttendeesRequestItemList"}
	for i := range v {
		if err := validateCreateAttendeeRequestItem(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEngineTranscribeMedicalSettings(v *types.EngineTranscribeMedicalSettings) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EngineTranscribeMedicalSettings"}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if len(v.Specialty) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Specialty"))
	}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTranscriptionConfiguration(v *types.TranscriptionConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TranscriptionConfiguration"}
	if v.EngineTranscribeMedicalSettings != nil {
		if err := validateEngineTranscribeMedicalSettings(v.EngineTranscribeMedicalSettings); err != nil {
			invalidParams.AddNested("EngineTranscribeMedicalSettings", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchCreateAttendeeInput(v *BatchCreateAttendeeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchCreateAttendeeInput"}
	if v.MeetingId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MeetingId"))
	}
	if v.Attendees == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Attendees"))
	} else if v.Attendees != nil {
		if err := validateCreateAttendeeRequestItemList(v.Attendees); err != nil {
			invalidParams.AddNested("Attendees", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAttendeeInput(v *CreateAttendeeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAttendeeInput"}
	if v.MeetingId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MeetingId"))
	}
	if v.ExternalUserId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExternalUserId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateMeetingInput(v *CreateMeetingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateMeetingInput"}
	if v.ClientRequestToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientRequestToken"))
	}
	if v.MediaRegion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MediaRegion"))
	}
	if v.ExternalMeetingId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExternalMeetingId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateMeetingWithAttendeesInput(v *CreateMeetingWithAttendeesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateMeetingWithAttendeesInput"}
	if v.ClientRequestToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientRequestToken"))
	}
	if v.MediaRegion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MediaRegion"))
	}
	if v.ExternalMeetingId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExternalMeetingId"))
	}
	if v.Attendees == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Attendees"))
	} else if v.Attendees != nil {
		if err := validateCreateMeetingWithAttendeesRequestItemList(v.Attendees); err != nil {
			invalidParams.AddNested("Attendees", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAttendeeInput(v *DeleteAttendeeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAttendeeInput"}
	if v.MeetingId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MeetingId"))
	}
	if v.AttendeeId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttendeeId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMeetingInput(v *DeleteMeetingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMeetingInput"}
	if v.MeetingId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MeetingId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAttendeeInput(v *GetAttendeeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAttendeeInput"}
	if v.MeetingId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MeetingId"))
	}
	if v.AttendeeId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttendeeId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMeetingInput(v *GetMeetingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMeetingInput"}
	if v.MeetingId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MeetingId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAttendeesInput(v *ListAttendeesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAttendeesInput"}
	if v.MeetingId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MeetingId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartMeetingTranscriptionInput(v *StartMeetingTranscriptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartMeetingTranscriptionInput"}
	if v.MeetingId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MeetingId"))
	}
	if v.TranscriptionConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TranscriptionConfiguration"))
	} else if v.TranscriptionConfiguration != nil {
		if err := validateTranscriptionConfiguration(v.TranscriptionConfiguration); err != nil {
			invalidParams.AddNested("TranscriptionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopMeetingTranscriptionInput(v *StopMeetingTranscriptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopMeetingTranscriptionInput"}
	if v.MeetingId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MeetingId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}