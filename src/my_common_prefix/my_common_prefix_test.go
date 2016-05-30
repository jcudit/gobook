package main

import (
	"testing"
)

func TestCommonPrefix(t *testing.T) {
	testPaths := []string{
		"/var/cache/salt/file",
		"/var/cache/salt/minion/file",
		"/var/cache/salt/master/file",
		"/var/cache/salt/master/1/file",
		"/var/cache/salt/syndic/file"}
	common := CommonPrefix(testPaths)
	commonPath := CommonPathPrefix(testPaths)
	if common != "/var/cache/salt/" || commonPath != "/var/cache/salt" {
		t.Fatalf("Failed to find common path: %s or %s", common, commonPath)
	}
	testPaths = []string{"/root", "/var/cache/salt", "/var/lib/whatever"}
	common = CommonPrefix(testPaths)
	if common != "/" {
		t.Fatalf("Failed to find common path: %s", common)
	}
	testPaths = []string{"uncommon", "/var/cache/salt", "/var/lib/whatever"}
	common = CommonPrefix(testPaths)
	if common != "" {
		t.Fatalf("Failed to find common path: %s", common)
	}
}
