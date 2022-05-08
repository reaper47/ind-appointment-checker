package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/reaper47/ind-appointment-checker/internal/jobs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ind",
	Short: "An IND appointement availability checker and notifier",
	Long: `An IND appointement availability checker and notifier.

The program fetches IND availabilities for every city every 10m. The user 
will be notified via Telegram when there is an appointment earlier than 
the currently booked date.`,
	Run: func(cmd *cobra.Command, args []string) {
		jobs.ScheduleCronJobs()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal("could not get working dir:", err)
	}

	err = godotenv.Load(filepath.Dir(exe) + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}
