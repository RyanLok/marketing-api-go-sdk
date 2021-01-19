/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type SplitTestsAddRequest struct {
	AccountId          int64    `json:"account_id,omitempty"`
	SplitTestName      string   `json:"split_test_name,omitempty"`
	BeginTime          int64    `json:"begin_time,omitempty"`
	EndTime            int64    `json:"end_time,omitempty"`
	AdgroupIdList      *[]int64 `json:"adgroup_id_list,omitempty"`
	SmartExpandEnabled *bool    `json:"smart_expand_enabled,omitempty"`
}
