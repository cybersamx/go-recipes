package custom

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var (
	sampleJSON = []byte(`
		{
			"id": "5be20aa1-f6aa-4b5f-abb0-d5bfa1a76ed2",
			"email": "nobody@example.com",
			"password": "MySecurePassword",
			"note": "Hello",
			"tags": "tag1,tag2"
		}
	`)
	sampleUserAlias = UserAliasType{
		ID:       "a0927e9f-dfad-4fab-84cd-2e3ceead1781",
		Email:    "somebody@example.com",
		Password: "DontExposeThis",
		Note:     "",
		Tags:     "",
	}
	sampleUserCustom = UserCustomType{
		ID:       "a0927e9f-dfad-4fab-84cd-2e3ceead1781",
		Email:    "somebody@example.com",
		Password: SensitiveString{"DontExposeThis"},
		Note:     "",
		Tags:     "",
	}
	sampleUserReflect = UserReflect{
		ID:       "a0927e9f-dfad-4fab-84cd-2e3ceead1781",
		Email:    "somebody@example.com",
		Password: "DontExposeThis",
		Note:     "",
		Tags:     "",
	}
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestUserAliasType_UnmarshalJSON(t *testing.T) {
	var user UserAliasType
	err := json.Unmarshal(sampleJSON, &user)
	require.NoError(t, err)

	assert.Equal(t, "5be20aa1-f6aa-4b5f-abb0-d5bfa1a76ed2", user.ID)
	assert.Equal(t, "nobody@example.com", user.Email)
	assert.Equal(t, "MySecurePassword", user.Password)
	assert.Equal(t, "Hello", user.Note)
	assert.Empty(t, user.Tags)
}

func TestUserCustomType_UnmarshalJSON(t *testing.T) {
	var user UserCustomType
	err := json.Unmarshal(sampleJSON, &user)
	require.NoError(t, err)

	assert.Equal(t, "5be20aa1-f6aa-4b5f-abb0-d5bfa1a76ed2", user.ID)
	assert.Equal(t, "nobody@example.com", user.Email)
	assert.Equal(t, "MySecurePassword", user.Password.string)
	assert.Equal(t, "Hello", user.Note)
	assert.Empty(t, user.Tags)
}

func TestUserReflectType_UnmarshalJSON(t *testing.T) {
	var user UserReflect
	err := json.Unmarshal(sampleJSON, &user)
	require.NoError(t, err)

	assert.Equal(t, "5be20aa1-f6aa-4b5f-abb0-d5bfa1a76ed2", user.ID)
	assert.Equal(t, "nobody@example.com", user.Email)
	assert.Equal(t, "MySecurePassword", user.Password)
	assert.Equal(t, "Hello", user.Note)
	assert.Empty(t, user.Tags)
}

func TestUserAliasType_MarshalJSON(t *testing.T) {
	data, err := json.Marshal(sampleUserAlias)
	require.NoError(t, err)

	require.NotNil(t, data)

	var userMap map[string]interface{}
	err = json.Unmarshal(data, &userMap)
	require.NoError(t, err)

	assert.Equal(t, "a0927e9f-dfad-4fab-84cd-2e3ceead1781", userMap["id"])
	assert.Equal(t, "somebody@example.com", userMap["email"])
	assert.Empty(t, userMap["password"]) // JSON has an empty password, ie. "password": "".
	assert.Nil(t, userMap["note"])
}

func TestUserCustomType_MarshalJSON(t *testing.T) {
	data, err := json.Marshal(sampleUserCustom)
	require.NoError(t, err)

	require.NotNil(t, data)

	var userMap map[string]interface{}
	err = json.Unmarshal(data, &userMap)
	require.NoError(t, err)

	assert.Equal(t, "a0927e9f-dfad-4fab-84cd-2e3ceead1781", userMap["id"])
	assert.Equal(t, "somebody@example.com", userMap["email"])
	assert.Empty(t, userMap["password"]) // JSON has an empty password, ie. "password": "".
	assert.Nil(t, userMap["note"])
}

func TestUserReflectType_MarshalJSON(t *testing.T) {
	data, err := json.Marshal(sampleUserReflect)
	require.NoError(t, err)

	require.NotNil(t, data)

	var userMap map[string]interface{}
	err = json.Unmarshal(data, &userMap)
	require.NoError(t, err)

	assert.Equal(t, "a0927e9f-dfad-4fab-84cd-2e3ceead1781", userMap["id"])
	assert.Equal(t, "somebody@example.com", userMap["email"])
	assert.Nil(t, userMap["password"]) // JSON does not have the "password" field.
	assert.Nil(t, userMap["note"])
}
