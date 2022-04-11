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
	UploadPath  string
	development bool
}

var rootCmd = &cobra.Command{
	Use:   "upload-tool",
	Short: "uploda-tool",
	Long:  "A tool to backup mysql file",
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
	fs.StringVar(&QSconfig.AccessKey, "accessKey", "", "access key")
	fs.StringVar(&QSconfig.SecretKey, "secretKey", "", "secret-key")
	fs.StringVar(&QSconfig.BucketName, "bucketName", "", "Specify the bucket for file upload")
	fs.StringVar(&QSconfig.Zone, "zone", "", "Specify the zone where the bucket is located")
	fs.StringVar(&QSconfig.LocalPath, "localPath", "", "The path to the sql file mount in the container")
	fs.StringVar(&QSconfig.UploadPath, "uploadPath", "", "The path to which the sql file needs to be uploaded")
	fs.BoolVar(&QSconfig.development, "development", false, "Use development logger config")
	_ = rootCmd.MarkFlagRequired("accessKey")
	_ = rootCmd.MarkFlagRequired("secretKey")
	_ = rootCmd.MarkFlagRequired("bucketName")
	_ = rootCmd.MarkFlagRequired("localPath")
	_ = rootCmd.MarkFlagRequired("uploadPath")
	_ = rootCmd.MarkFlagRequired("zone")
}
