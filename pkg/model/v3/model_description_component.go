/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 文本组件
type DescriptionComponent struct {
	ComponentId *int64             `json:"component_id,omitempty"`
	Value       *DescriptionStruct `json:"value,omitempty"`
}