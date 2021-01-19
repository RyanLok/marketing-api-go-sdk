/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type VideomakerTasksGetResponseData struct {
	Code      int64      `json:"code,omitempty"`
	TaskId    string     `json:"task_id,omitempty"`
	Status    Status     `json:"status,omitempty"`
	VideoId   string     `json:"video_id,omitempty"`
	OtherData *OtherData `json:"other_data,omitempty"`
}
