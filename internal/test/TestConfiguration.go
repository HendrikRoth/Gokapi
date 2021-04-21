//  +build test

package testconfiguration

import (
	"os"
	"testing"
)

// Create creates a configuration for unit testing
func Create(initFiles bool) {
	os.Setenv("GOKAPI_CONFIG_DIR", "test")
	os.Setenv("GOKAPI_DATA_DIR", "test")
	os.Mkdir("test", 0777)
	os.WriteFile("test/config.json", configTestFile, 0777)
	if initFiles {
		os.Mkdir("test/data", 0777)
		os.WriteFile("test/data/a8fdc205a9f19cc1c7507a60c4f01b13d11d7fd0", []byte("123"), 0777)
		os.WriteFile("test/data/c4f9375f9834b4e7f0a528cc65c055702bf5f24a", []byte("456"), 0777)
		os.WriteFile("test/data/e017693e4a04a59d0b0f400fe98177fe7ee13cf7", []byte("789"), 0777)
		os.WriteFile("test/data/2341354656543213246465465465432456898794", []byte("abc"), 0777)
		os.WriteFile("test/fileupload.jpg", []byte("abc"), 0777)
	}
}

// Delete deletes the configuration for unit testing
func Delete() {
	os.RemoveAll("test")
}

// StartMockInputStdin simulates a user input on stdin. Call StopMockInputStdin afterwards!
func StartMockInputStdin(t *testing.T, input string) *os.File {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write([]byte(input))
	if err != nil {
		t.Error(err)
	}
	w.Close()

	stdin := os.Stdin
	os.Stdin = r
	return stdin
}

// StopMockInputStdin needs to be called after StartMockInputStdin
func StopMockInputStdin(stdin *os.File) {
	os.Stdin = stdin
}

var configTestFile = []byte(`{
   "Port":"127.0.0.1:53843",
   "AdminName":"test",
   "AdminPassword":"10340aece68aa4fb14507ae45b05506026f276cf",
   "ServerUrl":"http://127.0.0.1:53843/",
   "DefaultDownloads":3,
   "DefaultExpiry":20,
   "DefaultPassword":"123",
   "RedirectUrl":"https://test.com/",
   "Sessions":{
      "validsession":{
         "RenewAt":2147483645,
         "ValidUntil":2147483646
      },
      "logoutsession":{
         "RenewAt":2147483645,
         "ValidUntil":2147483646
      },
      "needsRenewal":{
         "RenewAt":0,
         "ValidUntil":2147483646
      },
      "expiredsession":{
         "RenewAt":0,
         "ValidUntil":0
      }
   },
   "Files":{
      "Wzol7LyY2QVczXynJtVo":{
         "Id":"Wzol7LyY2QVczXynJtVo",
         "Name":"smallfile2",
         "Size":"8 B",
         "SHA256":"e017693e4a04a59d0b0f400fe98177fe7ee13cf7",
         "ExpireAt":2147483646,
         "ExpireAtString":"2021-05-04 15:19",
         "DownloadsRemaining":1,
         "PasswordHash":"",
         "ContentType":"text/html",
         "HotlinkId":""
      },
      "e4TjE7CokWK0giiLNxDL":{
         "Id":"e4TjE7CokWK0giiLNxDL",
         "Name":"smallfile2",
         "Size":"8 B",
         "SHA256":"e017693e4a04a59d0b0f400fe98177fe7ee13cf7",
         "ExpireAt":2147483645,
         "ExpireAtString":"2021-05-04 15:19",
         "DownloadsRemaining":2,
         "PasswordHash":"",
         "ContentType":"text/html",
         "HotlinkId":""
      },
      "wefffewhtrhhtrhtrhtr":{
         "Id":"wefffewhtrhhtrhtrhtr",
         "Name":"smallfile3",
         "Size":"8 B",
         "SHA256":"e017693e4a04a59d0b0f400fe98177fe7ee13cf7",
         "ExpireAt":2147483645,
         "ExpireAtString":"2021-05-04 15:19",
         "DownloadsRemaining":1,
         "PasswordHash":"",
         "ContentType":"text/html",
         "HotlinkId":""
      },
      "deletedfile123456789":{
         "Id":"deletedfile123456789",
         "Name":"DeletedFile",
         "Size":"8 B",
         "SHA256":"invalid",
         "ExpireAt":2147483645,
         "ExpireAtString":"2021-05-04 15:19",
         "DownloadsRemaining":2,
         "PasswordHash":"",
         "ContentType":"text/html",
         "HotlinkId":""
      },
      "jpLXGJKigM4hjtA6T6sN":{
         "Id":"jpLXGJKigM4hjtA6T6sN",
         "Name":"smallfile",
         "Size":"7 B",
         "SHA256":"c4f9375f9834b4e7f0a528cc65c055702bf5f24a",
         "ExpireAt":2147483646,
         "ExpireAtString":"2021-05-04 15:18",
         "DownloadsRemaining":1,
         "ContentType":"text/html",
         "PasswordHash":"7b30508aa9b233ab4b8a11b2af5816bdb58ca3e7",
         "HotlinkId":""
      },
      "jpLXGJKigM4hjtA6T6sN2":{
         "Id":"jpLXGJKigM4hjtA6T6sN2",
         "Name":"smallfile",
         "Size":"7 B",
         "SHA256":"c4f9375f9834b4e7f0a528cc65c055702bf5f24a",
         "ExpireAt":2147483646,
         "ExpireAtString":"2021-05-04 15:18",
         "DownloadsRemaining":1,
         "ContentType":"text/html",
         "PasswordHash":"7b30508aa9b233ab4b8a11b2af5816bdb58ca3e7",
         "HotlinkId":""
      },
      "n1tSTAGj8zan9KaT4u6p":{
         "Id":"n1tSTAGj8zan9KaT4u6p",
         "Name":"picture.jpg",
         "Size":"4 B",
         "SHA256":"a8fdc205a9f19cc1c7507a60c4f01b13d11d7fd0",
         "ExpireAt":2147483646,
         "ExpireAtString":"2021-05-04 15:19",
         "DownloadsRemaining":1,
         "PasswordHash":"",
         "ContentType":"text/html",
         "HotlinkId":"PhSs6mFtf8O5YGlLMfNw9rYXx9XRNkzCnJZpQBi7inunv3Z4A.jpg"
      },
      "cleanuptest123456789":{
         "Id":"cleanuptest123456789",
         "Name":"cleanup",
         "Size":"4 B",
         "SHA256":"2341354656543213246465465465432456898794",
         "ExpireAt":2147483646,
         "ExpireAtString":"2021-05-04 15:19",
         "DownloadsRemaining":0,
         "PasswordHash":"",
         "ContentType":"text/html",
         "HotlinkId":""
      }
   },
   "Hotlinks":{
      "PhSs6mFtf8O5YGlLMfNw9rYXx9XRNkzCnJZpQBi7inunv3Z4A.jpg":{
         "Id":"PhSs6mFtf8O5YGlLMfNw9rYXx9XRNkzCnJZpQBi7inunv3Z4A.jpg",
         "FileId":"n1tSTAGj8zan9KaT4u6p"
      }
   },
   "DownloadStatus":{
      "69JCbLVxx2KxfvB6FYkrDn3oCU7BWT":{
         "Id":"69JCbLVxx2KxfvB6FYkrDn3oCU7BWT",
         "FileId":"cleanuptest123456789",
         "ExpireAt":2147483646
      }
   },
   "ConfigVersion":5,
   "SaltAdmin":"LW6fW4Pjv8GtdWVLSZD66gYEev6NAaXxOVBw7C",
   "SaltFiles":"lL5wMTtnVCn5TPbpRaSe4vAQodWW0hgk00WCZE",
   "LengthId":20,
   "DataDir":"test/data"
}`)
