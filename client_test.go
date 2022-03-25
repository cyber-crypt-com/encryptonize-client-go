package client

import (
	"testing"

	"context"
	"log"
	"os"
)

var uid string
var password string
var certPath = ""
var endpoint = "localhost:9000"

var scopes = []Scope{
	ScopeRead,
	ScopeCreate,
	ScopeUpdate,
	ScopeDelete,
	ScopeIndex,
	ScopeObjectPermissions,
	ScopeUserManagement,
}

func failOnError(message string, err error, t *testing.T) {
	if err != nil {
		t.Fatalf("%s: %v", message, err)
	}
}

func TestMain(m *testing.M) {
	var ok bool
	uid, ok = os.LookupEnv("E2E_TEST_UID")
	if !ok {
		log.Fatal("E2E_TEST_UID must be set")
	}
	password, ok = os.LookupEnv("E2E_TEST_PASS")
	if !ok {
		log.Fatal("E2E_TEST_PASS must be set")
	}
	value, ok := os.LookupEnv("E2E_TEST_CERT")
	if ok {
		certPath = value
	}
	value, ok = os.LookupEnv("E2E_TEST_URL")
	if ok {
		endpoint = value
	}

	os.Exit(m.Run())
}

func TestUtility(t *testing.T) {
	c, err := NewClient(context.Background(), endpoint, certPath)
	failOnError("NewClient failed", err, t)
	defer c.Close()

	err = c.LoginUser(uid, password)
	failOnError("LoginUser failed", err, t)
	_, err = c.Health()
	failOnError("Health check failed", err, t)
	_, err = c.Version()
	failOnError("Version request failed", err, t)
}

func TestUserManagement(t *testing.T) {
	c, err := NewClient(context.Background(), endpoint, certPath)
	failOnError("NewClient failed", err, t)
	defer c.Close()

	err = c.LoginUser(uid, password)
	failOnError("LoginUser failed", err, t)

	createUserResponse, err := c.CreateUser(scopes)
	failOnError("CreateUser failed", err, t)

	createGroupResponse, err := c.CreateGroup(scopes)
	failOnError("CreateGroup failed", err, t)

	err = c.AddUserToGroup(createUserResponse.UserID, createGroupResponse.GroupID)
	failOnError("AddUserToGroup failed", err, t)

	err = c.RemoveUserFromGroup(createUserResponse.UserID, createGroupResponse.GroupID)
	failOnError("RemoveUserFromGroup failed", err, t)

	err = c.RemoveUser(createUserResponse.UserID)
	failOnError("RemoveUser failed", err, t)
}

func TestEncrypt(t *testing.T) {
	c, err := NewClient(context.Background(), endpoint, certPath)
	failOnError("NewClient failed", err, t)
	defer c.Close()

	err = c.LoginUser(uid, password)
	failOnError("LoginUser failed", err, t)

	createUserResponse, err := c.CreateUser(scopes)
	failOnError("CreateUser failed", err, t)
	err = c.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("LoginUser failed", err, t)

	plaintext := []byte("foo")
	associatedData := []byte("bar")
	encryptResponse, err := c.Encrypt(plaintext, associatedData)
	failOnError("Encrypt failed", err, t)

	decryptResponse, err := c.Decrypt(encryptResponse.ObjectID, encryptResponse.Ciphertext, encryptResponse.AssociatedData)
	failOnError("Decrypt failed", err, t)
	if string(decryptResponse.Plaintext) != string(plaintext) {
		t.Fatal("Decryption returned wrong plaintext")
	}
	if string(decryptResponse.AssociatedData) != string(associatedData) {
		t.Fatal("Decryption returned wrong data")
	}
}

func TestStore(t *testing.T) {
	c, err := NewClient(context.Background(), endpoint, certPath)
	failOnError("NewClient failed", err, t)
	defer c.Close()

	err = c.LoginUser(uid, password)
	failOnError("LoginUser failed", err, t)

	createUserResponse, err := c.CreateUser(scopes)
	failOnError("CreateUser failed", err, t)
	err = c.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("LoginUser failed", err, t)

	plaintext := []byte("foo")
	associatedData := []byte("bar")
	storeResponse, err := c.Store(plaintext, associatedData)
	failOnError("Store failed", err, t)

	retrieveResponse, err := c.Retrieve(storeResponse.ObjectID)
	failOnError("Retrieve failed", err, t)
	if string(retrieveResponse.Plaintext) != string(plaintext) {
		t.Fatal("Decryption returned wrong plaintext")
	}
	if string(retrieveResponse.AssociatedData) != string(associatedData) {
		t.Fatal("Decryption returned wrong data")
	}

	plaintext = []byte("foobar")
	associatedData = []byte("barbaz")
	err = c.Update(storeResponse.ObjectID, plaintext, associatedData)
	failOnError("Update failed", err, t)

	err = c.Delete(storeResponse.ObjectID)
	failOnError("Delete failed", err, t)
}

func TestPermissions(t *testing.T) {
	c, err := NewClient(context.Background(), endpoint, certPath)
	failOnError("NewClient failed", err, t)
	defer c.Close()

	err = c.LoginUser(uid, password)
	failOnError("LoginUser failed", err, t)

	createUserResponse, err := c.CreateUser(scopes)
	failOnError("CreateUser failed", err, t)
	err = c.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("LoginUser failed", err, t)

	plaintext := []byte("foo")
	associatedData := []byte("bar")
	storeResponse, err := c.Store(plaintext, associatedData)
	failOnError("Store failed", err, t)

	err = c.AddPermission(storeResponse.ObjectID, createUserResponse.UserID)
	failOnError("AddPermission failed", err, t)

	_, err = c.GetPermissions(storeResponse.ObjectID)
	failOnError("GetPermissions failed", err, t)

	err = c.RemovePermission(storeResponse.ObjectID, createUserResponse.UserID)
	failOnError("RemovePermission failed", err, t)
}
