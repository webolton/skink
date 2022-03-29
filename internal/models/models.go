package models

type ConfigFileData struct {
	ConfigFile string
	SyncedDir  string
	AMI        string
	AwsRegion  string
	Bucket     string
}
