package cmd

import (
	"github.com/replicatedcom/gin-example/api"
	"github.com/replicatedcom/gin-example/db"
	"github.com/replicatedcom/gin-example/inject"
	"github.com/replicatedcom/gin-example/services"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// apiCmd runs the gin example api server
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Runs the gin example api server",
	RunE: func(cmd *cobra.Command, args []string) error {
		dataSourceName := viper.GetString("data-source-name")
		db, err := db.Init(dataSourceName)
		if err != nil {
			log.WithField("err", err).Error("failed to initialize database")
			return err
		}

		env := &inject.Env{
			UserService: services.User(db),
		}
		return api.Run(env)
	},
}

func init() {
	apiCmd.Flags().String("secret", "unicornsAreAwesome", "JWT signing key")
	apiCmd.Flags().String("data-source-name", "user=postgres password=password host=postgres dbname=postgres sslmode=disable", "Database connection string")

	viper.BindPFlags(apiCmd.PersistentFlags()) // bind all pflags to viper
	viper.BindPFlags(apiCmd.Flags())

	RootCmd.AddCommand(apiCmd)
}
