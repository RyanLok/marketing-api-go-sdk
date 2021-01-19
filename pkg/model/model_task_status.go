/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// TaskStatus : 任务状态
type TaskStatus string

// List of TaskStatus
const (
	TaskStatus_PENDING    TaskStatus = "TASK_STATUS_PENDING"
	TaskStatus_PROCESSING TaskStatus = "TASK_STATUS_PROCESSING"
	TaskStatus_EXPIRED    TaskStatus = "TASK_STATUS_EXPIRED"
	TaskStatus_COMPLETED  TaskStatus = "TASK_STATUS_COMPLETED"
	TaskStatus_CANCELLED  TaskStatus = "TASK_STATUS_CANCELLED"
	TaskStatus_FAIL       TaskStatus = "TASK_STATUS_FAIL"
	TaskStatus_DELETED    TaskStatus = "TASK_STATUS_DELETED"
	TaskStatus_DRAFT      TaskStatus = "TASK_STATUS_DRAFT"
)
