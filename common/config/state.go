package config

import "github.com/spf13/viper"

func GetSateCmdChannelNum() int {
	return viper.GetInt("state.cmd_channel_num")
}
func GetSateServiceAddr() string {
	return viper.GetString("state.servide_addr")
}
func GetStateServiceName() string {
	return viper.GetString("state.service_name")
}
func GetSateServerPort() int {
	return viper.GetInt("state.server_port")
}
func GetSateRPCWeight() int {
	return viper.GetInt("state.weight")
}
