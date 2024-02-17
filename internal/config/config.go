package config

import (
	"github.com/spf13/pflag"
)

// var configPath string
// var showHelp bool

type Config struct {
	Server bool
	Client bool
}

func InitConfig() *Config {
	// command := &cobra.Command{
	// 	Run: func(c *cobra.Command, args []string) {
	// 		fmt.Println(viper.GetString("Flag"))
	// 	},
	// }

	// viper.SetDefault("Flag", "Flag Value")

	s := pflag.Bool("server", false, "Run server")
	c := pflag.Bool("client", false, "Run client")

	// pflag.StringVarP(&configPath, "config", "c", "", "Config file path")
	pflag.Parse()
	// if showHelp {
	// 	pflag.Usage()
	// 	return
	// }
	return &Config{Server: *s, Client: *c}
}
