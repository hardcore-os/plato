package config

import "github.com/spf13/viper"

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
func GetGatewayServerPort() int {
	return viper.GetInt("gateway.server_port")
}
func GetGatewayWorkerPoolNum() int {
	return viper.GetInt("gateway.worker_pool_num")
}
