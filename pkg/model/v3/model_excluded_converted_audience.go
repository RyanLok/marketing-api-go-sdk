/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 排除已转化人群行为定向
type ExcludedConvertedAudience struct {
	ExcludedDimension      ExcludedDimension `json:"excluded_dimension,omitempty"`
	ConversionBehaviorList *[]string         `json:"conversion_behavior_list,omitempty"`
}
