// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ChangeStatus string

// Enum values for ChangeStatus
const (
	ChangeStatusPreparing ChangeStatus = "PREPARING"
	ChangeStatusApplying  ChangeStatus = "APPLYING"
	ChangeStatusSucceeded ChangeStatus = "SUCCEEDED"
	ChangeStatusCancelled ChangeStatus = "CANCELLED"
	ChangeStatusFailed    ChangeStatus = "FAILED"
)

// Values returns all known values for ChangeStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChangeStatus) Values() []ChangeStatus {
	return []ChangeStatus{
		"PREPARING",
		"APPLYING",
		"SUCCEEDED",
		"CANCELLED",
		"FAILED",
	}
}

type FailureCode string

// Enum values for FailureCode
const (
	FailureCodeClienterror FailureCode = "CLIENT_ERROR"
	FailureCodeServerfault FailureCode = "SERVER_FAULT"
)

// Values returns all known values for FailureCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FailureCode) Values() []FailureCode {
	return []FailureCode{
		"CLIENT_ERROR",
		"SERVER_FAULT",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

// Values returns all known values for SortOrder. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASCENDING",
		"DESCENDING",
	}
}