package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"goblog/config"
	"net/http"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  `App will be served on host and port specified in config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	configs := configSet()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"app":     viper.Get("app.name"),
		})
	})
	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}

func configSet() config.Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}
	var configs config.Config

	if err := viper.Unmarshal(&configs); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return configs
}
