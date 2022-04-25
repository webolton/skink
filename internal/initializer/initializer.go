package initializer

import (
	"github.com/namsral/flag"
)

type Config struct {
	DefaultAWSCreds bool
	ConfigFilePath  string
	SyncedDirPath   string
	AMI             string
	AwsRegion       string
	Bucket          string
}

func (c Config) Initialize(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.String(flag.DefaultConfigFlagname, "/etc/skink.conf", "Path to config file")

	var (
		defaultAWSCreds = flags.Bool("default_aws_creds", true, "Use default AWS credentials file?")
		syncedDirPath   = flags.String("synced_directory_path", "", "Directory to sync to S3")
		ami             = flags.String("ami", "", "AMI owner of S3 bucket")
		awsRegion       = flags.String("aws_region", "", "Region of bucket")
		bucket          = flags.String("bucket", "", "Name of S3 bucket")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	c.DefaultAWSCreds = *defaultAWSCreds
	c.SyncedDirPath = *syncedDirPath
	c.AMI = *ami
	c.AwsRegion = *awsRegion
	c.Bucket = *bucket

	return nil
}
