package main

import "fmt"

var versionMajor = "0"
var versionMinor = "0"
var versionRevision = "0"
var versionGitCommitHash string
var versionCompileTime string

func getVersionNumberString() string {
	return fmt.Sprintf("%s.%s.%s", versionMajor, versionMinor, versionRevision)
}

func getVersionFullString() string {
	if len(versionGitCommitHash) == 0 {
		versionGitCommitHash = "UNKNOWN"
	}

	return fmt.Sprintf("AlwaysOnline/%s (+https://github.com/Jamesits/alwaysonline; Commit %s/%s)", getVersionNumberString(), versionGitCommitHash, versionCompileTime)
}
