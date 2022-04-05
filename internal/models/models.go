package models

type Config struct {
	ConfigFilePath string
	SyncedDirPath  string
	AMI            string
	AwsRegion      string
	Bucket         string
}
