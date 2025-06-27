/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"path/filepath"
	"runtime"
	"testing"

	"kmodules.xyz/image-packer/pkg/lib"
)

func Test_CheckImageArchitectures(t *testing.T) {
	dir, err := rootDir()
	if err != nil {
		t.Error(err)
	}

	if err := lib.CheckImageArchitectures([]string{
		filepath.Join(dir, "catalog", "imagelist.yaml"),
	}, []string{
		"ghcr.io/stashed/kubedump:0.2.0-v4",
		"ghcr.io/stashed/stash-elasticsearch:5.6.4-v36",
		"ghcr.io/stashed/stash-elasticsearch:6.2.4-v36",
		"ghcr.io/stashed/stash-elasticsearch:6.3.0-v36",
		"ghcr.io/stashed/stash-elasticsearch:6.4.0-v36",
		"ghcr.io/stashed/stash-elasticsearch:6.5.3-v36",
		"ghcr.io/stashed/stash-elasticsearch:6.8.0-v36",
		"ghcr.io/stashed/stash-elasticsearch:7.14.0-v22",
		"ghcr.io/stashed/stash-elasticsearch:7.2.0-v36",
		"ghcr.io/stashed/stash-elasticsearch:7.3.2-v36",
		"ghcr.io/stashed/stash-elasticsearch:8.2.0-v19",
		"ghcr.io/stashed/stash-mariadb:10.5.8-v30",
		"ghcr.io/stashed/stash-mongodb:3.4.17-v37",
		"ghcr.io/stashed/stash-mongodb:3.6.8-v37",
		"ghcr.io/stashed/stash-mysql:5.7.25-v37",
		"ghcr.io/stashed/stash-mysql:8.0.14-v36",
		"ghcr.io/stashed/stash-mysql:8.0.21-v30",
		"ghcr.io/stashed/stash-mysql:8.0.3-v36",
		"ghcr.io/stashed/stash-percona-xtradb:5.7-v26",
		"ghcr.io/stashed/stash-percona-xtradb:8.0",
		"ghcr.io/stashed/stash-percona-xtradb:8.4",
		"ghcr.io/stashed/stash-vault:1.10.3-v16",
	}); err != nil {
		t.Errorf("CheckImageArchitectures() error = %v", err)
	}
}

func rootDir() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("failed to locate root dir")
	}

	return filepath.Clean(filepath.Join(filepath.Dir(file), "..")), nil
}
