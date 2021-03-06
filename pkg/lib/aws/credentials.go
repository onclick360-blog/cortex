/*
Copyright 2020 Cortex Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package aws

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func GetCredentialsFromCLIConfigFile() (string, string, error) {
	creds := credentials.NewSharedCredentials("", "")
	if creds == nil {
		return "", "", ErrorReadCredentials()
	}

	value, err := creds.Get()
	if err != nil {
		return "", "", err
	}

	if value.AccessKeyID == "" || value.SecretAccessKey == "" {
		return "", "", ErrorReadCredentials()
	}

	return value.AccessKeyID, value.SecretAccessKey, nil
}
