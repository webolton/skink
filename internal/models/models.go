package models

type Config struct {
	DefaultAWSCreds bool
	ConfigFilePath  string
	SyncedDirPath   string
	AMI             string
	AwsRegion       string
	Bucket          string
}
