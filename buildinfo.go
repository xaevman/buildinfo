package buildinfo

import (
    "fmt"
    "runtime"
)

// branch enum
const (
    LOCAL = iota
    MAIN
    RELEASE
)

// name mappings
var (
    branchMap = map[string]int{
        "Local":   LOCAL,
        "Main":    MAIN,
        "Release": RELEASE,
    }

    branchNameMap = map[int]string{
        LOCAL:   "Local",
        MAIN:    "Main",
        RELEASE: "Release",
    }

    archNameMap = map[string]string{
        "386":   "86",
        "amd64": "64",
    }

    osNameMap = map[string]string{
        "darwin":  "darwin",
        "linux":   "linux",
        "windows": "win",
    }
)

func GetBranch(val int) string {
    name, ok := branchNameMap[val]
    if !ok {
        return "LOCAL"
    }

    return name
}

func GetBranchName(val string) int {
    id, ok := branchMap[val]
    if !ok {
        return LOCAL
    }

    return id
}

func GetPlatform() (string, error) {
    os, ok := osNameMap[runtime.GOOS]
    if !ok {
        return "", fmt.Errorf("OS Unsupported: %s", runtime.GOOS)
    }

    arch, ok := archNameMap[runtime.GOARCH]
    if !ok {
        return "", fmt.Errorf("Architecture Unsupported: %s", runtime.GOARCH)
    }

    return fmt.Sprintf("%s%s", os, arch), nil
}
