
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	b64 "encoding/base64"
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
		CLIENTID, isCLIENTID := os.LookupEnv("CLIENT_ID")
		CLIENTSECRET, isCLIENTSECRET :=os.LookupEnv("CLIENT_SECRET")
		credentials := CLIENTID + ":" + CLIENTSECRET
		credEncoded := b64.StdEncoding.EncodeToString([]byte(credentials))
		var missing = []string{}
		if !isCLIENTID {
			missing = append(missing, "CLIENT_ID")
		}

		if !isCLIENTSECRET {
			missing = append(missing, "CLIENT_SECRET")
		}

		if len(missing) > 0 {
			fmt.Errorf("Missing environment variables: %s", missing)
		}
		//resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
		//if err != nil {
		//	log.Fatalln(err)
		//}
		////We Read the response body on the line below.
		//body, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		////Convert the body to type string
		//sb := string(body)
		//log.Printf(sb)
		fmt.Println("artist " + artistName + " called")
		fmt.Println("base64 " + credEncoded)
	},
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
