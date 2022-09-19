package firebase

import (
	"google.golang.org/api/option"
	"path/filepath"
	"runtime"
)

func GetServiceAccount() option.ClientOption {
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "../../")

	return option.WithCredentialsFile(root + "/key/serviceAccount.json")
}
