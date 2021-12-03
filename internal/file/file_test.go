package file

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	path := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png"
	fmt.Println(Exists(path))
}

func TestIsDir(t *testing.T) {
	path := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv"
	fmt.Println(IsDir(path))
}

func TestIsFile(t *testing.T) {
	path := "/Users/apple/workspace_stariverpool/go-image/testdata/opencv-logo.png"
	fmt.Println(IsFile(path))

}
