package cmd

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/reaper47/ind-appointment-checker/internal/jobs"
	"github.com/spf13/cobra"
)

var (
	currentBiometricsAppointment string
	currentResidenceAppointment  string
)

var rootCmd = &cobra.Command{
	Use:   "ind",
	Short: "An IND appointement availability checker and notifier",
	Long: `An IND appointement availability checker and notifier.

The program fetches IND availabilities for every city every 10m. The user 
will be notified via Telegram when there is an appointment earlier than 
the currently booked date.`,
	Run: func(cmd *cobra.Command, args []string) {
		tBio, err := time.Parse("02/01/2006", currentBiometricsAppointment)
		if err != nil {
			tBio = time.Now().Add(1460 * time.Hour)
			log.Println("set current date to two months from now (" + tBio.Format("02/01/2006") + ") for biometrics")
		}

		tResidence, err := time.Parse("02/01/2006", currentResidenceAppointment)
		if err != nil {
			tBio = time.Now().Add(1460 * time.Hour)
			log.Println("set current date to two months from now (" + tBio.Format("02/01/2006") + ") for residence sticker")
		}

		jobs.ScheduleCronJobs(tBio, tResidence)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Flags().StringVarP(&currentBiometricsAppointment, "biometrics-appointment", "b", "", "Current biometrics appointment date [dd/mm/yy]")
	rootCmd.Flags().StringVarP(&currentResidenceAppointment, "residence-appointment", "r", "", "Current residence sticker appointment date [dd/mm/yy]")
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
