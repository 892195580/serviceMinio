package config

const (
	testEndpoint = "play.min.io"
	testAccessKeyID = "Q3AM3UQ867SPQQA43P2F"
	testSecretAccessKey = "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	selfEndpoint ="192.168.36.131:9000"
	selfAccessKeyID = "nKNB72xEmN7A01Wz"
	selfSecretAccessKey = "6mTdpUnPqHp5Hwc2ZO8LDLApsVKeBWJ8"
	useSSL = true
)

func GetCredentials() (ep, id, key string, ssl bool) {
	//endpoint, accessKeyID, secretAccessKey := testEndpoint,testAccessKeyID ,testSecretAccessKey
	endpoint, accessKeyID, secretAccessKey := selfEndpoint,selfAccessKeyID,selfSecretAccessKey
	return endpoint, accessKeyID, secretAccessKey, useSSL
}



