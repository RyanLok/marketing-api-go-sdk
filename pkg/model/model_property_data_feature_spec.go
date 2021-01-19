/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 属性数据特征规则
type PropertyDataFeatureSpec struct {
	UserPropertySetId  int64                `json:"user_property_set_id,omitempty"`
	PropertyDataKey    string               `json:"property_data_key,omitempty"`
	DataType           FeatureValueDataType `json:"data_type,omitempty"`
	IsMultiValued      *bool                `json:"is_multi_valued,omitempty"`
	PossibleValuesSize int64                `json:"possible_values_size,omitempty"`
}
