package read

import (
	"testing"
)

const testFile = "/.testconfig.json"


func TestGetFilePath(t *testing.T) {
	actual := "/Users/martinpare/blog_aggregator/internal/read/.testconfig.json"
	testCase, err := getConfigFilePath(testFile)
	if err != nil {
		t.Fatalf("func: GetConfigFilePath does not work")
	}
	
	if actual != testCase {
		t.Fatalf("filepaths do not match. test failed\n actual: %s | test: %s", actual, testCase)
	}

}

func TestJsonFile(t *testing.T) {
	actual := "dbURL: testing"
	test, err := Read(testFile)
	if err != nil {
		t.Fatalf("could not read testfile")
	}

	if actual !=  test.DbURL {
		t.Fatalf("data does not match")
	}
}