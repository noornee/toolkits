package redis

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/noornee/toolkits/cmd"
	"github.com/noornee/toolkits/internal/environment"
	redisclient "github.com/noornee/toolkits/internal/redis"
	"github.com/spf13/cobra"
)

var (
	key      string
	filePath string
	envPath  string
)

// Define the parent redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Interact with Redis",
}

// Define the set subcommand
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value in Redis",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the environment
		env, err := environment.New(envPath)
		if err != nil {
			log.Fatalf("Error initializing environment: %v", err)
		}

		// Initialize Redis client
		client, err := redisclient.NewRedisClient(env)
		if err != nil {
			log.Fatalf("Error initializing Redis: %v", err)
		}

		// Read file content (if you want to store file data as the value)
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		// Set value in Redis
		ctx := context.Background()
		if err := client.SetValue(ctx, key, string(fileData), 0); err != nil {
			log.Fatalf("Error setting value in Redis: %v", err)
		}

		fmt.Println("Data stored successfully in Redis!")
	},
}

// Define the get subcommand
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value from Redis",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the environment
		env, err := environment.New(envPath)
		if err != nil {
			log.Fatalf("Error initializing environment: %v", err)
		}

		// Initialize Redis client
		client, err := redisclient.NewRedisClient(env)
		if err != nil {
			log.Fatalf("Error initializing Redis: %v", err)
		}

		// Get value from Redis
		ctx := context.Background()
		value, err := client.GetValue(ctx, key)
		if err != nil {
			log.Fatalf("Error getting value from Redis: %v", err)
		}

		fmt.Printf("%s", value)
	},
}

func init() {
	// Add the `redis` command as a subcommand of the root command
	cmd.RootCmd.AddCommand(redisCmd)
	cmd.RootCmd.PersistentFlags().StringVarP(&envPath, "env", "e", ".env", "Path to the .env file")

	// Add the `set` subcommand
	redisCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&key, "key", "k", "", "The Redis key where data will be stored")
	setCmd.Flags().StringVarP(&filePath, "file", "f", "", "The path to the file whose content will be stored in Redis")
	setCmd.MarkFlagRequired("key")  //nolint:errcheck
	setCmd.MarkFlagRequired("file") //nolint:errcheck

	// Add the `get` subcommand
	redisCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&key, "key", "k", "", "The Redis key to retrieve data from")
	getCmd.MarkFlagRequired("key") //nolint:errcheck
}
