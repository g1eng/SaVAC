package core

const (
	VERSION                             = "0.7.0"
	USER_AGENT                          = "savac-" + VERSION
	CANDIDATE_RELEASE_NUMBER            = "5"
	OBJECT_STORAGE_API_KEY_ENV_NAME     = "SAKURASTORAGE_ACCESS_KEY"
	OBJECT_STORAGE_API_SECRET_ENV_NAME  = "SAKURASTORAGE_ACCESS_SECRET"
	OBJECT_STORAGE_AWS_PROFILE_ENV_NAME = "SAKURASTORAGE_PROFILE"
	OBJECT_STORAGE_DEFAULT_PROFILE_NAME = "sakurastorage"
	OBJECT_STORAGE_DEFAULT_S3_ENDPOINT  = "https://s3.isk01.sakurastorage.jp"
	OBJECT_STORAGE_DEFAULT_S3_REGION    = "jp-north-1"
	OutputTypeText                      = 0
	OutputTypeJson                      = 1
	OutputTypeYaml                      = 2
)

var OutputTypes = []int{OutputTypeText, OutputTypeJson, OutputTypeYaml}
