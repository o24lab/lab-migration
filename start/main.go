package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"migration"
	"os"
	"pkg/logger"
)

func main() {
	logger.Init()
	var rootCmd = &cobra.Command{
		Use:   "root",
		Short: "root",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Info("start success")
		},
	}
	{
		var userCmd = &cobra.Command{
			Use:   "user",
			Short: "user",
			Run: func(cmd *cobra.Command, args []string) {
				err := migration.User()
				if err != nil {
					logrus.Error(err)
				}
				logrus.Info("user migrated successfully")
			},
		}
		rootCmd.AddCommand(userCmd)
	}
	{
		var kycCmd = &cobra.Command{
			Use:   "kyc",
			Short: "kyc",
			Run: func(cmd *cobra.Command, args []string) {
				err := migration.Kyc()
				if err != nil {
					logrus.Error(err)
				}
				logrus.Info("kyc migrated successfully")
			},
		}
		rootCmd.AddCommand(kycCmd)
	}
	{
		var coinCmd = &cobra.Command{
			Use:   "coin",
			Short: "coin",
			Run: func(cmd *cobra.Command, args []string) {
				err := migration.Coin()
				if err != nil {
					logrus.Error(err)
				}
				logrus.Info("coin migrated successfully")
			},
		}
		rootCmd.AddCommand(coinCmd)
	}
	{
		var allCmd = &cobra.Command{
			Use:   "all",
			Short: "all",
			Run: func(cmd *cobra.Command, args []string) {
				{
					err := migration.User()
					if err != nil {
						logrus.Error(err)
					}
					logrus.Info("user migrated successfully")
				}
				{
					err := migration.Kyc()
					if err != nil {
						logrus.Error(err)
					}
					logrus.Info("kyc migrated successfully")
				}
				{
					err := migration.Coin()
					if err != nil {
						logrus.Error(err)
					}
					logrus.Info("coin migrated successfully")
				}
				logrus.Info("migrate success")
			},
		}
		rootCmd.AddCommand(allCmd)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
