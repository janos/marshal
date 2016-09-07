// Copyright (c) 2016, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package marshal

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestModeMarshalJSON(t *testing.T) {
	tests := []struct {
		mode Mode
		want []byte
	}{
		{Mode(0), []byte(`"000"`)},
		{Mode(0644), []byte(`"644"`)},
		{Mode(0777), []byte(`"777"`)},
	}

	for _, test := range tests {
		result, err := json.Marshal(test.mode)
		if err != nil {
			t.Errorf("%#v: error: %s", test.mode, err)
		}
		if !bytes.Equal(result, test.want) {
			t.Errorf("%#v: got %s, want %s", test.mode, result, test.want)
		}
	}
}

func TestModeMarshalJSONError(t *testing.T) {
	tests := []struct {
		mode Mode
		err  error
	}{
		{Mode(0111), nil},
		{Mode(-1), ErrInvalidOctalMode},
		{Mode(01000), ErrInvalidOctalMode},
	}

	for _, test := range tests {
		_, err := test.mode.MarshalJSON()
		if err != test.err {
			t.Errorf("%#v: got: %q, want error %q", test.mode, err, test.err)
		}
	}
}

func TestModeMarshalText(t *testing.T) {
	tests := []struct {
		mode Mode
		want []byte
	}{
		{Mode(0), []byte(`000`)},
		{Mode(0644), []byte(`644`)},
		{Mode(0777), []byte(`777`)},
	}

	for _, test := range tests {
		result, err := test.mode.MarshalText()
		if err != nil {
			t.Errorf("%#v: error: %s", test.mode, err)
		}
		if !bytes.Equal(result, test.want) {
			t.Errorf("%#v: got %s, want %s", test.mode, result, test.want)
		}
	}
}

func TestModeMarshalTextError(t *testing.T) {
	tests := []struct {
		mode Mode
		err  error
	}{
		{Mode(0111), nil},
		{Mode(-1), ErrInvalidOctalMode},
		{Mode(01000), ErrInvalidOctalMode},
	}

	for _, test := range tests {
		_, err := test.mode.MarshalText()
		if err != test.err {
			t.Errorf("%#v: got: %q, want error %q", test.mode, err, test.err)
		}
	}
}

func TestModeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input []byte
		want  Mode
	}{
		{[]byte(`"000"`), Mode(0)},
		{[]byte(`"644"`), Mode(0644)},
		{[]byte(`"777"`), Mode(0777)},
	}

	for _, test := range tests {
		m := Mode(0)
		if err := json.Unmarshal(test.input, &m); err != nil {
			t.Errorf("%s: error: %s", test.input, err)
		}
		if m != test.want {
			t.Errorf("%s: got %q, want %q", test.input, m, test.want)
		}
	}
}

func TestModeUnmarshalJSONError(t *testing.T) {
	tests := []struct {
		input []byte
		err   error
	}{
		{[]byte(``), ErrInvalidOctalMode},
		{[]byte(`""`), ErrInvalidOctalMode},
		{[]byte(`0`), ErrInvalidOctalMode},
		{[]byte(`"0"`), nil},
		{[]byte(`"555"`), nil},
	}

	for _, test := range tests {
		m := Mode(0)
		if err := m.UnmarshalJSON(test.input); err != test.err {
			t.Errorf("%s: got: %q, want error %q", test.input, err, test.err)
		}
	}
}

func TestModeUnmarshalText(t *testing.T) {
	tests := []struct {
		input []byte
		want  Mode
	}{
		{[]byte(`000`), Mode(0)},
		{[]byte(`644`), Mode(0644)},
		{[]byte(`777`), Mode(0777)},
	}

	for _, test := range tests {
		m := Mode(0)
		if err := m.UnmarshalText(test.input); err != nil {
			t.Errorf("%s: error: %s", test.input, err)
		}
		if m != test.want {
			t.Errorf("%s: got %q, want %q", test.input, m, test.want)
		}
	}
}

func TestModeUnmarshalTextError(t *testing.T) {
	tests := []struct {
		input []byte
		err   error
	}{
		{[]byte(``), ErrInvalidOctalMode},
		{[]byte(`""`), ErrInvalidOctalMode},
		{[]byte(`a`), ErrInvalidOctalMode},
		{[]byte(`0`), nil},
		{[]byte(`555`), nil},
	}

	for _, test := range tests {
		m := Mode(0)
		if err := m.UnmarshalText(test.input); err != test.err {
			t.Errorf("%s: got: %q, want error %q", test.input, err, test.err)
		}
	}
}

func TestModeFileMode(t *testing.T) {
	tests := []struct {
		mode Mode
		want os.FileMode
	}{
		{Mode(0), os.FileMode(0)},
		{Mode(0644), os.FileMode(0644)},
		{Mode(0777), os.FileMode(0777)},
	}

	for _, test := range tests {
		if response := test.mode.FileMode(); response != test.want {
			t.Errorf("%q: got %q, want %q", test.mode, response, test.want)
		}
	}
}
