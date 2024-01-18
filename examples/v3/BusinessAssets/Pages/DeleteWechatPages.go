/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model/v3"
)

type WechatPagesDeleteExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.WechatPagesDeleteRequest
}

func (e *WechatPagesDeleteExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.WechatPagesDeleteRequest{
		PageId: int64(0),
	}
}

func (e *WechatPagesDeleteExample) RunExample() (model.WechatPagesDeleteResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.WechatPages().Delete(ctx, e.Data)
}

func main() {
	e := &WechatPagesDeleteExample{}
	e.Init()
	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Headers:", headers)
}
