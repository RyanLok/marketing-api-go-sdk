/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type TargetingsShareGetListStruct struct {
	TargetingId        int64 `json:"targeting_id,omitempty"`
	ShareToAccountId   int64 `json:"share_to_account_id,omitempty"`
	ShareToTargetingId int64 `json:"share_to_targeting_id,omitempty"`
	SharedTime         int64 `json:"shared_time,omitempty"`
}
