// Copyright (c) Facebook, Inc. and its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		pkgStr string
		pkg    Package
		fail   bool
	}{
		{pkgStr: "", fail: true},
		{pkgStr: "name", fail: true},
		{pkgStr: "name-version", fail: true},
		{pkgStr: "name-version-release", fail: true},
		{
			pkgStr: "name-version-release.arch",
			pkg: Package{
				Name: "name",
				Label: Label{
					Epoch:   "",
					Version: "version",
					Release: "release",
				},
				Arch: "arch",
			},
		},
		{
			pkgStr: "name-version-release.arch.rpm",
			pkg: Package{
				Name: "name",
				Label: Label{
					Epoch:   "",
					Version: "version",
					Release: "release",
				},
				Arch: "arch",
			},
		},
		{
			pkgStr: "name-epoch:version-release.arch",
			pkg: Package{
				Name: "name",
				Label: Label{
					Epoch:   "epoch",
					Version: "version",
					Release: "release",
				},
				Arch: "arch",
			},
		},
		{
			pkgStr: "name-epoch:version-release.arch.rpm",
			pkg: Package{
				Name: "name",
				Label: Label{
					Epoch:   "epoch",
					Version: "version",
					Release: "release",
				},
				Arch: "arch",
			},
		},
		{
			pkgStr: "name-version-release.src",
			pkg: Package{
				Name: "name",
				Label: Label{
					Epoch:   "",
					Version: "version",
					Release: "release",
				},
				Arch: "",
			},
		},
		{
			pkgStr: "name-version-release.noarch",
			pkg: Package{
				Name: "name",
				Label: Label{
					Epoch:   "",
					Version: "version",
					Release: "release",
				},
				Arch: "",
			},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			pkg, err := Parse(c.pkgStr)
			if err != nil {
				if !c.fail {
					t.Fatalf("%q: unexpected failure: %v", c.pkgStr, err)
				}
				return
			}
			if c.fail {
				t.Fatalf("%q: unexpected success", c.pkgStr)
			}

			if !reflect.DeepEqual(c.pkg, *pkg) {
				t.Errorf("wrong result:\n\thave: %+v\n\twant: %+v", c.pkg, *pkg)
			}
		})

	}
}

func BenchmarkParse(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Parse("NaMe-1.0-1.i386.rpm")
	}
}
