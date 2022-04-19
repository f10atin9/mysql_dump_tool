package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var QSconfig struct {
	BucketName  string
	AccessKey   string
	SecretKey   string
	Zone        string
	LocalPath   string
	BucketPath  string
	development bool
}

var rootCmd = &cobra.Command{
	Use:   "dump-tool",
	Short: "dump-tool",
	Long:  "A tool to dump sql",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return subMain()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	fs := rootCmd.Flags()
	fs.StringVar(&QSconfig.AccessKey, "accessKey", os.Getenv("QingStor_AccessorKey"), "access key")
	fs.StringVar(&QSconfig.SecretKey, "secretKey", os.Getenv("QingStor_SecretKey"), "secret key")
	fs.StringVar(&QSconfig.BucketName, "bucketName", os.Getenv("QingStor_BucketName"), "Specify the bucket for file upload")
	fs.StringVar(&QSconfig.Zone, "zone", os.Getenv("QingStor_Zone"), "Specify the zone where the bucket is located")
	fs.StringVar(&QSconfig.LocalPath, "localPath", os.Getenv("VolumeDump_Path"), "The path to the sql file mount in the container")
	fs.StringVar(&QSconfig.BucketPath, "bucketPath", os.Getenv("Bucket_Path"), "")
	fs.BoolVar(&QSconfig.development, "development", false, "Use development logger config")
}
