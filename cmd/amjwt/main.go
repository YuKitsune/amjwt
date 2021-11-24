package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitsune/amjwt"
	"io"
	"io/ioutil"
	"os"
)

const privateKeyFilePathName = "private-key-file"

var privateKeyFilePath string

const keyIdName = "key-id"

var keyId string

const teamIdName = "team-id"

var teamId string

const expiryDaysName = "expiry"

var expiryDays int

// Apple Music has a maximum of 6 months
const maxExpiryDays = 6 * 30

func main() {

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints the version of the binary",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(amjwt.Version)
			return nil
		},
	}

	var rootCmd = &cobra.Command{
		Use:   "amjwt",
		Short: "Generates the JWT for Apple Music",
		Args:  cobra.OnlyValidArgs,
		RunE:  exec,
	}

	rootCmd.Flags().StringVarP(&keyId, keyIdName, "k", "", "A 10-character key identifier, obtained from your Apple Developer account (required)")
	if err := rootCmd.MarkFlagRequired(keyIdName); err != nil {
		exitFromError(err)
	}

	rootCmd.Flags().StringVarP(&teamId, teamIdName, "t", "", "A 10-character Team ID, obtained this from your Apple developer account (required)")
	if err := rootCmd.MarkFlagRequired(teamIdName); err != nil {
		exitFromError(err)
	}

	rootCmd.Flags().IntVarP(&expiryDays, expiryDaysName, "e", maxExpiryDays, "number of days before the token expires")

	rootCmd.Flags().StringVarP(&privateKeyFilePath, privateKeyFilePathName, "f", "", "The path to the private key file")

	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		// Cobra will handle printing the error if this is the case
		os.Exit(1)
	}
}

func exec(cmd *cobra.Command, args []string) error {

	var privateKeyBytes []byte
	inBytes, err := readStdin()
	if err != nil {
		return err
	}

	if expiryDays > maxExpiryDays {
		return fmt.Errorf("expiry must not be greater than %d days", maxExpiryDays)
	}

	if len(inBytes) == 0 && len(privateKeyFilePath) == 0 {
		return fmt.Errorf("no private key was provided, use the --%s flag or sdtin", privateKeyFilePathName)
	}

	if len(inBytes) > 0 {
		privateKeyBytes = inBytes
	} else {
		privateKeyBytes, err = ioutil.ReadFile(privateKeyFilePath)
		if err != nil {
			return err
		}
	}

	tokenString, err := amjwt.CreateJwt(keyId, teamId, expiryDays, privateKeyBytes)
	if err != nil {
		return err
	}

	fmt.Println(tokenString)
	return nil
}

func readStdin() (bytes []byte, err error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return bytes, err
	}

	// Nothing from stdin, no worries
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		return bytes, nil
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}

			return bytes, err
		}

		bytes = append(bytes, b)
	}

	return bytes, nil
}

func exitFromError(err error) {
	format := "error: %v\n"
	if _, printErr := fmt.Fprintf(os.Stderr, format, err); printErr != nil {
		// couldn't print to stderr, just print normally i guess
		fmt.Printf(format, err)
	}

	os.Exit(1)
}
