package main

import "fmt"

var versionMajor = 1
var versionMinor = 0
var versionRevision = 0
var versionGitCommitHash string
var versionCompileTime string
var versionCompileHost string
var versionGitStatus string

func getVersionNumberString() string {
	return fmt.Sprintf("%d.%d.%d", versionMajor, versionMinor, versionRevision)
}

func getVersionFullString() string {
	if len(versionCompileHost) == 0 {
		versionCompileHost = "localhost"
	}

	if len(versionGitCommitHash) == 0 {
		versionGitCommitHash = "UNKNOWN"
	}

	if len(versionCompileTime) == 0 {
		versionCompileTime = "UNKNOWN TIME"
	}

	if len(versionGitStatus) == 0 {
		versionGitStatus = "dirty"
	}

	return fmt.Sprintf("AlwaysOnline/%s (+https://github.com/Jamesits/alwaysonline; Compiled on %s for commit %s (%s) at %s)", getVersionNumberString(), versionCompileHost, versionGitCommitHash, versionGitStatus, versionCompileTime)
}
