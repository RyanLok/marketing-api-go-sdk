/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type BidwordDeleteRequest struct {
	AccountId *int64   `json:"account_id,omitempty"`
	List      *[]int64 `json:"list,omitempty"`
}