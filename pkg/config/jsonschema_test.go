/*
Copyright © 2020-2022 The k3d Author(s)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package config

import (
	"testing"

	conf "github.com/rancher/k3d/v5/pkg/config/v1alpha4"
)

func TestValidateSchema(t *testing.T) {

	cfgPath := "./test_assets/config_test_simple.yaml"

	if err := ValidateSchemaFile(cfgPath, []byte(conf.JSONSchema)); err != nil {
		t.Errorf("Validation of config file %s against the default schema failed: %+v", cfgPath, err)
	}

}

func TestValidateSchemaFail(t *testing.T) {

	cfgPath := "./test_assets/config_test_simple_invalid_servers.yaml"

	var err error
	if err = ValidateSchemaFile(cfgPath, []byte(conf.JSONSchema)); err == nil {
		t.Errorf("Validation of config file %s against the default schema passed where we expected a failure", cfgPath)
	}

	expectedErrorText := `- metadata.name: Invalid type. Expected: string, given: integer
`

	if err.Error() != expectedErrorText {
		t.Errorf("Actual validation error\n%s\ndoes not match expected error\n%s\n", err.Error(), expectedErrorText)
	}

}
