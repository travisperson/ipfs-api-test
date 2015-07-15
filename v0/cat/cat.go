package cat

import (
	config "github.com/travisperson/ipfs-api-test/config"
	files "github.com/travisperson/ipfs-api-test/files"
	add "github.com/travisperson/ipfs-api-test/v0/add"
	"net/http"
	"io/ioutil"
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

var uri string = config.U("/api/v0/cat")

const fileName string = "file2.txt"
const fileBody string = "Hello2"

func before(t *testing.T) string {
	assert := assert.New(t)

	fs := []files.File{
		{fileName, fileBody, "QmNULL"},
	}

	request := add.AddFiles(t, fs)

	// Perform the request
	resp, err := http.DefaultClient.Do(request)

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode, "Add file for before")

	// Verify our return values
	var file files.File
	dec := json.NewDecoder(resp.Body)

	err = dec.Decode(&file)
	assert.NoError(err)

	return file.Hash
}

func Get (t *testing.T) {
	assert := assert.New(t)

	// Setup the file
	hash := before(t)

	// Setup the request
	request, err := http.NewRequest("GET", uri, nil)

	values := request.URL.Query()
	values.Add("arg", hash)
	request.URL.RawQuery = values.Encode()

	assert.NoError(err)

	// Perform the request
	resp, err := http.DefaultClient.Do(request)

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	buf, err := ioutil.ReadAll(resp.Body)

	assert.NoError(err)
	assert.Equal(fileBody, string(buf))
}
