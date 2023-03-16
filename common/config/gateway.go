package config

import (
	"github.com/spf13/viper"
)

func GetGatewayMaxTcpNum() int32 {
	return viper.GetInt32("gateway.tcp_max_num")
}

func GetGatewayEpollerChanNum() int {
	return viper.GetInt("gateway.epoll_channel_size")
}
func GetGatewayEpollerNum() int {
	return viper.GetInt("gateway.epoll_num")
}
func GetGatewayEpollWaitQueueSize() int {
	return viper.GetInt("gateway.epoll_wait_queue_size")
}
func GetGatewayTCPServerPort() int {
	return viper.GetInt("gateway.tcp_server_port")
}
func GetGatewayRPCServerPort() int {
	return viper.GetInt("gateway.rpc_server_port")
}
func GetGatewayWorkerPoolNum() int {
	return viper.GetInt("gateway.worker_pool_num")
}
func GetGatewayCmdChannelNum() int {
	return viper.GetInt("gateway.cmd_channel_num")
}

func GetGatewayServiceAddr() string {
	return viper.GetString("gateway.service_addr")
}

func GetGatewayServiceName() string {
	return viper.GetString("gateway.service_name")
}

func GetGatewayRPCWeight() int {
	return viper.GetInt("gateway.weight")
}

func GetGatewayStateServerEndPoint() string {
	return viper.GetString("gateway.state_server_endpoint")
}
