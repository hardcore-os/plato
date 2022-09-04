package config

import "github.com/spf13/viper"

// GetDiscovName 获取discov用哪种方式实现
func GetDiscovName() string {
	return viper.GetString("prpc.discov.name")
}

// GetDiscovEndpoints 获取discov的 endpoints
func GetDiscovEndpoints() []string {
	return viper.GetStringSlice("discovery.endpoints")
}

// GetTraceEnable 是否开启trace
func GetTraceEnable() bool {
	return viper.GetBool("prpc.trace.enable")
}

// GetTraceCollectionUrl 获取trace collection url
func GetTraceCollectionUrl() string {
	return viper.GetString("prpc.trace.url")
}

// GetTraceServiceName 获取服务名
func GetTraceServiceName() string {
	return viper.GetString("prpc.trace.service_name")
}

// GetTraceSampler 获取trace采样率
func GetTraceSampler() float64 {
	return viper.GetFloat64("prpc.trace.sampler")
}
