/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// PlayableType : 蹊径落地页试玩类型
type PlayableType string

// List of PlayableType
const (
	PlayableType_PLAYABLE_TYPE_DEFAULT              PlayableType = "PLAYABLE_TYPE_DEFAULT"
	PlayableType_PLAYABLE_TYPE_HOMEMADE_INTERACTION PlayableType = "PLAYABLE_TYPE_HOMEMADE_INTERACTION"
	PlayableType_PLAYABLE_TYPE_MINIGAME_INTERACTION PlayableType = "PLAYABLE_TYPE_MINIGAME_INTERACTION"
	PlayableType_PLAYABLE_TYPE_VIDEO_INTERACTION    PlayableType = "PLAYABLE_TYPE_VIDEO_INTERACTION"
	PlayableType_PLAYABLE_TYPE_WEBSITE_INTERACTION  PlayableType = "PLAYABLE_TYPE_WEBSITE_INTERACTION"
	PlayableType_PLAYABLE_TYPE_ZIP_INTERACTION      PlayableType = "PLAYABLE_TYPE_ZIP_INTERACTION"
	PlayableType_NOT_INTERACT                       PlayableType = "NOT_INTERACT"
	PlayableType_INLINE                             PlayableType = "INLINE"
	PlayableType_TEMPLATE_GAME                      PlayableType = "TEMPLATE_GAME"
	PlayableType_TEMPLATE_VIDEO                     PlayableType = "TEMPLATE_VIDEO"
	PlayableType_TEMPLATE_WEB                       PlayableType = "TEMPLATE_WEB"
	PlayableType_COMPRESSED_PACKAGE                 PlayableType = "COMPRESSED_PACKAGE"
)