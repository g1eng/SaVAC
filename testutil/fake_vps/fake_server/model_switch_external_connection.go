/*
 * さくらのVPS APIドキュメント
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.5.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package helper

// SwitchExternalConnection - 接続されている外部接続の情報
type SwitchExternalConnection struct {

	// サービスコード
	ServiceCode string `json:"service_code"`

	// 外部接続方式
	Type string `json:"type"`

	Services []SwitchExternalConnectionServicesInner `json:"services"`
}
