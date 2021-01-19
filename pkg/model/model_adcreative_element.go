/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告创意元素
type AdcreativeElement struct {
	Name           string                           `json:"name,omitempty"`
	ElementType    ElementType                      `json:"element_type,omitempty"`
	FieldType      FieldType                        `json:"field_type,omitempty"`
	Required       *bool                            `json:"required,omitempty"`
	Description    string                           `json:"description,omitempty"`
	ParentName     string                           `json:"parent_name,omitempty"`
	EnumProperty   *AdcreativeElementEnumProperty   `json:"enum_property,omitempty"`
	ArrayProperty  *AdcreativeElementArrayProperty  `json:"array_property,omitempty"`
	StructProperty *AdcreativeElementStructProperty `json:"struct_property,omitempty"`
	Restriction    *AdcreativeElementRestriction    `json:"restriction,omitempty"`
}
