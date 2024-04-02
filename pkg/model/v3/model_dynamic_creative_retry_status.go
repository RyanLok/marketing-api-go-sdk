/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// DynamicCreativeRetryStatus : 重试状态
type DynamicCreativeRetryStatus string

// List of DynamicCreativeRetryStatus
const (
	DynamicCreativeRetryStatus_UNKNOWN        DynamicCreativeRetryStatus = "DYNAMIC_CREATIVE_RETRY_STATUS_UNKNOWN"
	DynamicCreativeRetryStatus_WAIT_FOR_RETRY DynamicCreativeRetryStatus = "DYNAMIC_CREATIVE_RETRY_STATUS_WAIT_FOR_RETRY"
	DynamicCreativeRetryStatus_REJECTED       DynamicCreativeRetryStatus = "DYNAMIC_CREATIVE_RETRY_STATUS_REJECTED"
	DynamicCreativeRetryStatus_RETRYING       DynamicCreativeRetryStatus = "DYNAMIC_CREATIVE_RETRY_STATUS_RETRYING"
)