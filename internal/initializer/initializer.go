package initializer

import (
	"errors"

	"github.com/namsral/flag"
	"github.com/webolton/skink/internal/models"
)

func (c *models.Config) Initialize(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.String(flag.DefaultConfigFlagname, "/etc/skink.conf", "Path to config file")

	var (
		defaultAWSCreds = flags.Bool("default_aws_creds", true, "Use default AWS credentials file?")
		syncedDirPath   = flags.String("synced_directory_path", "", "Directory to sync to S3")
		ami             = flags.String("ami", "", "AMI owner of S3 bucket")
		awsRegion       = flags.String("aws_region", "", "Region of bucket")
		bucket          = flags.String("bucket", "", "Name of S3 bucket")
	)
	return errors.New("Placeholder")
}
