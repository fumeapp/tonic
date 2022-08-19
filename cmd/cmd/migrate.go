/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/fumeapp/tonic/database"
	"github.com/fumeapp/tonic/setting"
	"github.com/golang-migrate/migrate"
	"github.com/octoper/go-ray"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")

		if _, err := os.Stat("./database/migrations"); os.IsNotExist(err) {
			fmt.Println("No migration directory found")
		}

		setting.Setup()
		database.Setup()
		ray.Ray(database.DSN())

		m, err := migrate.New("file://database/migrations", database.DSN())
		if err != nil {
			log.Fatal(err)
		}
		m.Up()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
