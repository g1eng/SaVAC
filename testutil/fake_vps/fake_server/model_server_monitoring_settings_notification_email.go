/*
 * さくらのVPS APIドキュメント
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.5.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package helper

// ServerMonitoringSettingsNotificationEmail - emailでの通知の設定。会員情報に登録されているメールアドレス宛に送信されます。
type ServerMonitoringSettingsNotificationEmail struct {

	// 通知のON/OFF * true 通知ON * false 通知OFF
	Enabled bool `json:"enabled"`
}
