package cmd

import (
	"log"

	"github.com/emmett-weisz/paystream-microservice/server"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "paystream",
	Short: "Start the Paystream gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetString("grpc.port")
		if port == "" {
			port = "50051"
		}

		writer := kafka.NewWriter(kafka.WriterConfig{
			Brokers:  []string{"localhost:9092"},
			Topic:    "payments",
			Balancer: &kafka.LeastBytes{},
		})

		log.Printf("Starting gRPC server on port %s...", port)
		if err := server.RunGRPCServer(port, writer); err != nil {
			log.Fatalf("failed to run server: %v", err)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Load config file with viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("No config file found, using defaults: %v", err)
	}

	rootCmd.Flags().String("port", "50051", "gRPC server port")
	viper.BindPFlag("grpc.port", rootCmd.Flags().Lookup("port"))
}
