// Copyright 2020-2022 CYBERCRYPT
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpce2e

import (
	"testing"

	"context"

	"google.golang.org/grpc/codes"

	coreclient "github.com/cyber-crypt-com/encryptonize-core/client"
)

type LoginDetails struct {
	userid, password string
	expectedCode     codes.Code
}

func TestAuthenticated(t *testing.T) {
	client, err := coreclient.NewClient(context.Background(), endpoint, certPath)
	failOnError("Could not create client", err, t)
	defer client.Close()

	err = client.LoginUser(uid, pwd)
	failOnError("Could not log in user", err, t)

	newUser, err := client.CreateUser(protoUserScopes)
	failOnError("Create user request failed", err, t)

	err = client.LoginUser(newUser.UserID, newUser.Password)
	failOnError("Could not log in user", err, t)
}

func TestWrongCredentials(t *testing.T) {
	client, err := coreclient.NewClient(context.Background(), endpoint, certPath)
	failOnError("Could not create client", err, t)
	defer client.Close()

	err = client.LoginUser(uid, pwd)
	failOnError("Could not log in user", err, t)

	newUser, err := client.CreateUser(protoUserScopes)
	failOnError("Create user request failed", err, t)

	wrongCredentials := []LoginDetails{
		{userid: "", password: "", expectedCode: codes.InvalidArgument},
		{userid: "", password: pwd, expectedCode: codes.InvalidArgument},
		{userid: uid, password: "", expectedCode: codes.Unauthenticated},
		{userid: uid, password: "wrong password", expectedCode: codes.Unauthenticated},
		{userid: newUser.UserID, password: pwd, expectedCode: codes.Unauthenticated},
		{userid: newUser.UserID, password: "wrong password", expectedCode: codes.Unauthenticated},
	}

	for _, cred := range wrongCredentials {
		err = client.LoginUser(cred.userid, cred.password)
		failOnSuccess("Should not be able to log in with wrong credentials", err, t)
		checkStatusCode(err, cred.expectedCode, t)
	}
}

func TestWrongToken(t *testing.T) {
	client, err := coreclient.NewClient(context.Background(), endpoint, certPath)
	failOnError("Could not create client", err, t)
	defer client.Close()

	// No token
	_, err = client.Version()
	failOnSuccess("Should not be able to get version with a wrong token", err, t)
	checkStatusCode(err, codes.InvalidArgument, t)

	// Wrong format
	client.SetToken("bad__bad__token!")
	_, err = client.Version()
	failOnSuccess("Should not be able to get version with a wrong token", err, t)
	checkStatusCode(err, codes.Unauthenticated, t)

	// Wrong contents
	client.SetToken("QW4gaW52YWxpZCB0b2tlbg")
	_, err = client.Version()
	failOnSuccess("Should not be able to get version with a wrong token", err, t)
	checkStatusCode(err, codes.Unauthenticated, t)
}
