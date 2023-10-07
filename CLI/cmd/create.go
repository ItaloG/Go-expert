/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ItaloG/Go-expert/CLI/internal/database"
	"github.com/spf13/cobra"
)

func newCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  "Create a new category",
		RunE:  runCreate(categoryDb),
	}
}

// createCmd represents the create command
// var createCmd = &cobra.Command{
// 	Use:   "create",
// 	Short: "A brief description of your command",
// 	Long:  `A longer description that spans multiple lines and likely contains examples`,
// 	RunE:  runCreate(GetCategoryDB(GetDb())),
// Run: func(cmd *cobra.Command, args []string) {
// 	db := GetDb()
// 	category := GetCategoryDB(db)
// 	name, _ := cmd.Flags().GetString("name")
// 	description, _ := cmd.Flags().GetString("description")
// 	category.Create(name, description)

// 	cmd.Help()
// },
// }

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.Flags().StringP("description", "d", "", "Description of the category")
	// createCmd.MarkFlagRequired("name")
	// createCmd.MarkFlagRequired("description")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
