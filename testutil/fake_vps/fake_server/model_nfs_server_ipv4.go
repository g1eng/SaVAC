/*
 * さくらのVPS APIドキュメント
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.5.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package helper

type NfsServerIpv4 struct {

	// アドレス
	Address string `json:"address"`

	// サブネットマスク
	Netmask string `json:"netmask"`
}
