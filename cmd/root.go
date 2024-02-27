package cmd

import (
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/jsnfwlr/bcrypt-cli/internal/feedback"
	"github.com/jsnfwlr/bcrypt-cli/internal/prompt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var buildVer = "0.0.1-alpha"

var rootCmd = &cobra.Command{
	Version: buildVer,
	Use:     "bcrypt",
	Short:   "bcrypt is a command to produce a bcrypt hash of the input string",
	Long:    "bcrypt is a simple command that ports the `htpasswd -nB` command to Go and can be installed without the other apache-utils",
	Args:    cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		batch, err := cmd.Flags().GetBool("batch")
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveError
		}

		if batch {
			return args, cobra.ShellCompDirectiveError
		}

		return []string{}, cobra.ShellCompDirectiveError
	},

	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		batch, err := cmd.Flags().GetBool("batch")
		feedback.HandleFatalErr(err)

		cost, err := cmd.Flags().GetInt("cost")
		feedback.HandleFatalErr(err)

		if cost > 17 || cost < 4 {
			feedback.HandleFatalErr(fmt.Errorf("cost must be between 4 and 17 - not %d", cost))
		}

		prompted := ""
		switch len(args) {
		case 0:
			if batch {
				feedback.HandleFatalErr(fmt.Errorf("batch-mode requires an input argument"))
			}
			prompted = prompt.Password("Password", false)
		case 1:
			if !batch {
				feedback.HandleFatalErr(fmt.Errorf("interactive-mode does not accept an input argument"))
			}
			prompted = args[0]
		default:
			feedback.HandleFatalErr(fmt.Errorf("too many (%d) input arguments", len(args)))
		}
		runCoreCmd(ctx, cost, prompted)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initBcrypt)

	rootCmd.Flags().BoolP("batch", "b", false, "Use batch mode; i.e., get the password from the command line rather than prompting for it. This option should be used with extreme care, since the password is clearly visible on the command line")
	rootCmd.Flags().IntP("cost", "C", 5, "Set the computing time used for the bcrypt algorithm (higher is more secure but slower, default: 5, valid: 4 to 17)")
}

// initBcrypt reads in config file and ENV variables if set.
func initBcrypt() {
	regVer := regexp.MustCompile(`.*(dev|alpha|test).*`)
	matches := regVer.FindSubmatch([]byte(buildVer))
	if len(matches) == 2 {
		feedback.EnableCaller()
	}

	feedback.SuppressNoise(feedback.NoiseLevel(initQuiet()))
}

func initQuiet() int {
	quietness, err := rootCmd.PersistentFlags().GetCount("quiet")
	if err != nil {
		return 0
	}

	return quietness
}

// runCoreCmd uses bcrypt algorithm to encrypt the string p with a cost of c and outputs the result to the terminal
func runCoreCmd(ctx context.Context, c int, p string) {
	o, err := bcrypt.GenerateFromPassword([]byte(p), c)
	feedback.HandleFatalErr(err)

	feedback.Print(feedback.Required, "%s", string(o))
}
