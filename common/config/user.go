package config

import "github.com/spf13/viper"

func GetDomainUserServerName() string {
	return viper.GetString("user_domain.service_name")
}

func GetDomainUserServerAddr() string {
	return viper.GetString("user_domain.service_addr")
}

func GetDomainUserServerPoint() int {
	return viper.GetInt("user_dimain.service_port")
}
func GetDomainUserRPCWeight() int {
	return viper.GetInt("user_dimain.weight")
}
func GetDomainUserDBDNS() string {
	return viper.GetString("user_dimain.db_dns")
}
