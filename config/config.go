package config

const (
	//endpoint = "play.min.io"
	//accessKeyID = "Q3AM3UQ867SPQQA43P2F"
	//secretAccessKey = "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	endpoint ="192.168.36.131:9000"
	accessKeyID = "nKNB72xEmN7A01Wz"
	secretAccessKey = "6mTdpUnPqHp5Hwc2ZO8LDLApsVKeBWJ8"
	useSSL = true
)

func GetCredentials() (ep, id, key string, ssl bool) {
	return endpoint, accessKeyID, secretAccessKey, useSSL
}



