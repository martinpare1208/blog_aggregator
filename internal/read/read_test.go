package read

import (
	"testing"
)


func TestGetFilePath(t *testing.T) {
	actual := "/Users/martinpare/blog_aggregator/internal/config/.gatorconfig.json"
	testCase, err := getConfigFilePath()
	if err != nil {
		t.Fatalf("func: GetConfigFilePath does not work")
	}
	if actual != testCase {
		t.Fatalf("filepaths do not match. test failed")
	}

}