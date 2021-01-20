
package cmd

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"bytes"
)

var(
	artistName string
)

//type Artist struct {
//
//}

// artistCmd represents the artist command
var artistCmd = &cobra.Command{
	Use:   "artist",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getToken()
	},
}

func getToken(){
	siteHost := "https://accounts.spotify.com/api"
	CLIENTID, isCLIENTID := os.LookupEnv("CLIENT_ID")
	CLIENTSECRET, isCLIENTSECRET :=os.LookupEnv("CLIENT_SECRET")
	credentials := CLIENTID + ":" + CLIENTSECRET
	credentialsEncoded := b64.StdEncoding.EncodeToString([]byte(credentials))
	fmt.Println(credentialsEncoded)
	var missing []string
	if !isCLIENTID {
		missing = append(missing, "CLIENT_ID")
	}

	if !isCLIENTSECRET {
		missing = append(missing, "CLIENT_SECRET")
	}

	if len(missing) > 0 {
		fmt.Errorf("Missing environment variables: %s", missing)
	}


	data := url.Values{}
	data.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/token", siteHost), bytes.NewBufferString(data.Encode()))
	req.Header.Set("authorization", "basic " + credentialsEncoded)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	fmt.Println("req", req)
	if err != nil {
		log.Println(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err 1", err)
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err 2", err)
	}
	resp.Body.Close()
	if err != nil {
		log.Fatal("err 3", err)
	}
	fmt.Println("response...", string(f))
}

func init() {
	rootCmd.AddCommand(artistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// artistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	artistCmd.Flags().StringVarP(&artistName, "name", "n", "", "Returns artist detail given artist name")
	//artistCmd.MarkPersistentFlagRequired("name")
}
