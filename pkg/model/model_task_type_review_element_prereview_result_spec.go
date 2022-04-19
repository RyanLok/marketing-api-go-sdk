/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 获取元素的预审结果条件
type TaskTypeReviewElementPrereviewResultSpec struct {
	AdgroupId  *int64                                    `json:"adgroup_id,omitempty"`
	Elements   *[]PrereviewElementStruct                 `json:"elements,omitempty"`
	Supplement *[]ReviewElementPrereviewSupplementStruct `json:"supplement,omitempty"`
}