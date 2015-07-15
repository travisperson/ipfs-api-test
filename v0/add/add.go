package add


import (
	config "github.com/travisperson/ipfs-api-test/config"
	files "github.com/travisperson/ipfs-api-test/files"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
	mh "github.com/jbenet/go-multihash"
	"encoding/json"
)


var uri string = config.U("/api/v0/add")

func Get (t *testing.T) {
	assert := assert.New(t)

	resp, err := http.Get(uri)

	assert.NoError(err)
	assert.Equal(http.StatusMethodNotAllowed, resp.StatusCode)
}

func PostEmpty (t *testing.T) {
	assert := assert.New(t)

	resp, err := http.Post(uri, "text/plain", strings.NewReader(""))

	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, resp.StatusCode)
}

func PostFile (t *testing.T) {
	assert := assert.New(t)

	fs := []files.File{
		{"file.txt", "Hello", "QmNULL"},
		{"file2.txt", "Hello2", "QmNULL"},
	}

	request := AddFiles(t, fs)

	// Perform the request
	resp, err := http.DefaultClient.Do(request)

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// Verify our return values
	var file files.File
	dec := json.NewDecoder(resp.Body)

	// TODO: The API does not have to return files in order
	// go-ipfs just happens to do so as of right now
	for _, f := range fs {
		err = dec.Decode(&file)
		assert.NoError(err)

		// Verify the name
		assert.Equal(f.Name, file.Name)

		// Verify hash
		_, err = mh.FromB58String(file.Hash)
		assert.NoError(err)
	}
}

func AddFiles(t *testing.T, fs []files.File) *http.Request {
	assert := assert.New(t)

	// Add a file to the request
	mfs := files.NewMultiFile()

	for _, f := range fs {
		err := mfs.AddFile(f.Name, strings.NewReader(f.Body))
		assert.NoError(err)
	}

	// Close
	err := mfs.Writer.Close()
	assert.NoError(err)

	// Create the request
	request, err := http.NewRequest("POST", uri, mfs.Body)
	assert.NoError(err)

	request.Header.Add("Content-Type", mfs.Writer.FormDataContentType())

	return request
}
