/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type BidwordDeleteResponseData struct {
	SuccessList *[]BidwordRespStruct `json:"success_list,omitempty"`
	ErrorList   *[]BidwordRespStruct `json:"error_list,omitempty"`
}
