/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 业务点
type BusinessPointDataStructs struct {
	Level    *int64    `json:"level,omitempty"`
	ParentId *string   `json:"parent_id,omitempty"`
	Value    *string   `json:"value,omitempty"`
	Desc     *string   `json:"desc,omitempty"`
	Options  *[]string `json:"options,omitempty"`
}
