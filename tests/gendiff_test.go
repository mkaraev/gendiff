package tests

import (
	"fmt"
	"github.com/mkaraev/gendiff/pkg/gendiff"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	TestdataDirectory = "testdata/"
	ExpectedDirectory = "expected/"
)

func TestGenerateDiff(t *testing.T) {
	tests := []struct {
		file1              string
		file2              string
		expectedResultFile string
		format             string
	}{
		{
			file1:              filepath.Join(TestdataDirectory, "file1.yaml"),
			file2:              filepath.Join(TestdataDirectory, "file2.yaml"),
			expectedResultFile: filepath.Join(ExpectedDirectory, "file1_file2_plain.txt"),
			format:             "plain",
		},
		{
			file1:              filepath.Join(TestdataDirectory, "file1.yaml"),
			file2:              filepath.Join(TestdataDirectory, "file2.yaml"),
			expectedResultFile: filepath.Join(ExpectedDirectory, "file1_file2_stylish.txt"),
			format:             "stylish",
		},
		{
			file1: filepath.Join(TestdataDirectory, "file1.yaml"),
			file2: filepath.Join(TestdataDirectory, "empty.yaml"),
			expectedResultFile: filepath.Join(ExpectedDirectory, "file1_empty_json.txt"),
			format: "json",
		},
		{
			file1:              filepath.Join(TestdataDirectory, "file1.yaml"),
			file2:              filepath.Join(TestdataDirectory, "file2.yaml"),
			expectedResultFile: filepath.Join(ExpectedDirectory, "file1_file2_json.txt"),
			format:             "json",
		},
		{
			file1:              filepath.Join(TestdataDirectory, "file1.json"),
			file2:              filepath.Join(TestdataDirectory, "file2.json"),
			expectedResultFile: filepath.Join(ExpectedDirectory, "file1_file2_json.txt"),
			format:             "json",
		},
		{
			file1:              filepath.Join(TestdataDirectory, "file1.json"),
			file2:              filepath.Join(TestdataDirectory, "file2.json"),
			expectedResultFile: filepath.Join(ExpectedDirectory, "file1_file2_plain.txt"),
			format:             "plain",
		},
		{
			file1:              filepath.Join(TestdataDirectory, "file1.json"),
			file2:              filepath.Join(TestdataDirectory, "file2.json"),
			expectedResultFile: filepath.Join(ExpectedDirectory, "file1_file2_stylish.txt"),
			format:             "stylish",
		},
	}
	fmt.Println(os.Getwd())
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %s %s", test.file1, test.file2, test.format), func(t *testing.T) {

			diff, err := gendiff.GenerateDiff(test.file1, test.file2, test.format)
			require.NoError(t, err)
			exptected, err := os.ReadFile(test.expectedResultFile)
			require.NoError(t, err)
			assert.Equal(t, diff, string(exptected), fmt.Sprintf("Expected %s but got %s", string(exptected), diff))
		})
	}
}
