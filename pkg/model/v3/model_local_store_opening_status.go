/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// LocalStoreOpeningStatus : 门店经营状态
type LocalStoreOpeningStatus string

// List of LocalStoreOpeningStatus
const (
	LocalStoreOpeningStatus_UNKNOWN            LocalStoreOpeningStatus = "OPENING_STATUS_UNKNOWN"
	LocalStoreOpeningStatus_OPENING            LocalStoreOpeningStatus = "OPENING_STATUS_OPENING"
	LocalStoreOpeningStatus_TEMPORARILY_CLOSED LocalStoreOpeningStatus = "OPENING_STATUS_TEMPORARILY_CLOSED"
	LocalStoreOpeningStatus_PERMANENTLY_CLOSED LocalStoreOpeningStatus = "OPENING_STATUS_PERMANENTLY_CLOSED"
)
