// Copyright (c) 2016, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package marshal

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestCheckboxMarshalJSON(t *testing.T) {
	tests := []struct {
		checkbox Checkbox
		want     []byte
	}{
		{Checkbox(true), []byte(`"on"`)},
		{Checkbox(false), []byte(`false`)},
	}

	for _, test := range tests {
		result, err := json.Marshal(test.checkbox)
		if err != nil {
			t.Errorf("%#v: error: %s", test.checkbox, err)
		}
		if !bytes.Equal(result, test.want) {
			t.Errorf("%#v: got %s, want %s", test.checkbox, result, test.want)
		}
	}
}

func TestCheckboxUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input []byte
		want  Checkbox
	}{
		{[]byte(`"on"`), Checkbox(true)},
		{[]byte(`false`), Checkbox(false)},
		{[]byte(`true`), Checkbox(true)},
	}

	for _, test := range tests {
		d := Checkbox(false)
		if err := json.Unmarshal(test.input, &d); err != nil {
			t.Errorf("%s: error: %s", test.input, err)
		}
		if d != test.want {
			t.Errorf("%s: got %v, want %v", test.input, d, test.want)
		}
	}
}

func TestCheckboxBool(t *testing.T) {
	tests := []struct {
		checkbox Checkbox
		want     bool
	}{
		{Checkbox(true), true},
		{Checkbox(false), false},
	}

	for _, test := range tests {
		if result := test.checkbox.Bool(); result != test.want {
			t.Errorf("%#v: got %v, want %v", test.checkbox, result, test.want)
		}
	}
}
