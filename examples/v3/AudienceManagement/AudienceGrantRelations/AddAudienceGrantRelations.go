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

type AudienceGrantRelationsAddExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.AudienceGrantRelationsAddRequest
}

func (e *AudienceGrantRelationsAddExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.AudienceGrantRelationsAddRequest{
		AccountId: int64(0),
		GrantType: model.AudienceGrantType_GRANT_TYPE_BUSINESS,
		GrantSpec: &model.GrantSpec{
			GrantToBusinessSpec: &model.GrantToBusinessSpec{
				GrantBusinessId: int64(0),
				GrantScopeType:  model.AudienceGrantScopeType_BUSINESS,
				GrantBusinessPermission: &model.GrantBusinessPermission{
					GrantPermissionTypeList: &[]string{"GRANT_PERMISSION_TYPE_TARGET"},
				},
			},
		},
		AudienceIdList: &[]int64{int64(0)},
	}
}

func (e *AudienceGrantRelationsAddExample) RunExample() (interface{}, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.AudienceGrantRelations().Add(ctx, e.Data)
}

func main() {
	e := &AudienceGrantRelationsAddExample{}
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
