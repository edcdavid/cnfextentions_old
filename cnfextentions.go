// Copyright (C) 2020-2021 Red Hat, Inc.
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

package cnfextentions

import (
	"fmt"

	"github.com/onsi/ginkgo/v2"
	"github.com/sirupsen/logrus"
	"github.com/test-network-function/cnf-certification-test/cnf-certification-test/identifiers"
	"github.com/test-network-function/cnf-certification-test/cnf-certification-test/results"
	"github.com/test-network-function/cnf-certification-test/pkg/provider"
	"github.com/test-network-function/test-network-function-claim/pkg/claim"
)

// catalog - can be here or in a subdirectory under cnfextentions
const (
	ExampleTestKey = "myorg"
	ExampleTestID  = "example"
	versionTwo     = "v2.0.0"
	url            = "http://myorg.com/testcases"
)

// Change this function as needed
func formTestURL(suite, name string) string {
	return fmt.Sprintf("%s/%s/%s", url, suite, name)
}

// Simple example test suite. More complex scenarios should use subdirectories
var _ = ginkgo.Describe(ExampleTestKey, func() {
	TestExampleIdentifier := claim.Identifier{
		Url:     formTestURL(ExampleTestKey, ExampleTestID),
		Version: versionTwo,
	}
	var env provider.TestEnvironment
	ginkgo.BeforeEach(func() {
		env = provider.GetTestEnvironment()
	})
	ginkgo.ReportAfterEach(results.RecordResult)

	testID := identifiers.XformToGinkgoItIdentifier(TestExampleIdentifier)
	ginkgo.It(testID, ginkgo.Label(testID), func() {
		testExample(&env)
	})
})

func testExample(env *provider.TestEnvironment) {
	logrus.Infof("Hello!")
}
