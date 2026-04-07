package tenjin

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go-admin/config"
)

// ReportEvent 上报事件
func ReportEvent(platform, eventName, analyticsInstallationId, advertisingId, developerDeviceId, osVersion, appVersion, ipAddress, value string) error {
	platform = normalizePlatform(platform)
	var cfg config.TenjinPlatform
	if platform == "ios" {
		cfg = config.ExtConfig.Tenjin.Ios
	} else {
		cfg = config.ExtConfig.Tenjin.Android
	}

	if !cfg.Enable {
		return nil
	}

	params := url.Values{}
	params.Set("api_key", cfg.ApiKey)
	params.Set("bundle_id", cfg.BundleId)
	params.Set("platform", platform)
	params.Set("analytics_installation_id", analyticsInstallationId)
	params.Set("event", eventName)
	params.Set("value", value)
	params.Set("sdk_version", "server")

	if advertisingId != "" {
		params.Set("advertising_id", advertisingId)
	}
	if developerDeviceId != "" {
		params.Set("developer_device_id", developerDeviceId)
	}
	if osVersion != "" {
		params.Set("os_version", osVersion)
	} else {
		params.Set("os_version", cfg.DefaultOsVersion)
	}
	if appVersion != "" {
		params.Set("app_version", appVersion)
	} else {
		params.Set("app_version", cfg.DefaultAppVersion)
	}
	if ipAddress != "" {
		params.Set("ip_address", ipAddress)
	}
	params.Set("limit_ad_tracking", cfg.LimitAdTracking)
	params.Set("ad_user_data", cfg.AdUserData)
	params.Set("ad_personalization", cfg.AdPersonalization)

	reqUrl := cfg.ApiUrl + "?" + params.Encode()
	client := &http.Client{Timeout: time.Duration(cfg.TimeoutMs) * time.Millisecond}
	_, err := client.Get(reqUrl)
	return err
}

func normalizePlatform(platform string) string {
	return strings.ToLower(strings.TrimSpace(platform))
}

// ReportOrderSuccess 上报订单成功事件
func ReportOrderSuccess(platform, analyticsInstallationId, advertisingId, developerDeviceId, osVersion, appVersion, ipAddress string, value int) error {
	platform = normalizePlatform(platform)
	eventName := "Cardpartner_sell_success"
	if platform == "ios" {
		eventName = "CardpartnerIOS_SuccessfulTransaction"
	}
	return ReportEvent(platform, eventName, analyticsInstallationId, advertisingId, developerDeviceId, osVersion, appVersion, ipAddress, fmt.Sprintf("%d", value))
}

// ReportWithdrawalSuccess 上报提现成功事件
func ReportWithdrawalSuccess(platform, analyticsInstallationId, advertisingId, developerDeviceId, osVersion, appVersion, ipAddress, amount string) error {
	platform = normalizePlatform(platform)
	eventName := "Cardpartner_withdraw_success"
	if platform == "ios" {
		eventName = "CardpartnerIOS_SuccessfulWithdrawal"
	}
	return ReportEvent(platform, eventName, analyticsInstallationId, advertisingId, developerDeviceId, osVersion, appVersion, ipAddress, amount)
}
