// Copyright (c) 2017, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package marshal

import (
	"bytes"
	"encoding/json"
	"net/url"
	"testing"
	"time"
)

func mustURL(s string) URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return URL(*u)
}

func TestURLMarshalJSON(t *testing.T) {
	tests := []struct {
		url  URL
		want []byte
	}{
		{mustURL("https://resenje.org/"), []byte(`"https://resenje.org/"`)},
		{mustURL("http://localhost:8080/status"), []byte(`"http://localhost:8080/status"`)},
	}

	for _, test := range tests {
		result, err := json.Marshal(test.url)
		if err != nil {
			t.Errorf("%#v: error: %s", test.url, err)
		}
		if !bytes.Equal(result, test.want) {
			t.Errorf("%#v: got %s, want %s", test.url, result, test.want)
		}
	}
}

func TestURLMarshalText(t *testing.T) {
	tests := []struct {
		url  URL
		want []byte
	}{
		{mustURL("https://resenje.org/"), []byte(`https://resenje.org/`)},
		{mustURL("http://localhost:8080/status"), []byte(`http://localhost:8080/status`)},
	}

	for _, test := range tests {
		result, err := test.url.MarshalText()
		if err != nil {
			t.Errorf("%#v: error: %s", test.url, err)
		}
		if !bytes.Equal(result, test.want) {
			t.Errorf("%#v: got %s, want %s", test.url, result, test.want)
		}
	}
}

func TestURLUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input []byte
		want  URL
	}{
		{[]byte(`"https://resenje.org/"`), mustURL("https://resenje.org/")},
		{[]byte(`"http://localhost:8080/status"`), mustURL("http://localhost:8080/status")},
	}

	for _, test := range tests {
		u := URL{}
		if err := json.Unmarshal(test.input, &u); err != nil {
			t.Errorf("%s: error: %s", test.input, err)
		}
		if u != test.want {
			t.Errorf("%s: got %v, want %v", test.input, u, test.want)
		}
	}
}

func TestURLUnmarshalJSONError(t *testing.T) {
	tests := []struct {
		input []byte
		err   error
	}{
		{[]byte(`""`), ErrInvalidURL},
		{[]byte(`0`), ErrInvalidURL},
		{[]byte(`"https://resenje.org/"`), nil},
		{[]byte(`"http://localhost:8080/status"`), nil},
	}

	for _, test := range tests {
		u := URL{}
		if err := u.UnmarshalJSON(test.input); err != test.err {
			t.Errorf("%s: got: %q, want error %q", test.input, err, test.err)
		}
	}
}

func TestURLUnmarshalText(t *testing.T) {
	tests := []struct {
		input []byte
		want  URL
	}{
		{[]byte(`https://resenje.org/`), mustURL("https://resenje.org/")},
		{[]byte(`http://localhost:8080/status`), mustURL("http://localhost:8080/status")},
	}

	for _, test := range tests {
		u := URL{}
		if err := u.UnmarshalText(test.input); err != nil {
			t.Errorf("%s: error: %s", test.input, err)
		}
		if u != test.want {
			t.Errorf("%s: got %v, want %v", test.input, u, test.want)
		}
	}
}

func TestURLUnmarshalTextError(t *testing.T) {
	tests := []struct {
		input []byte
		err   string
	}{
		{[]byte(``), "time: invalid duration "},
		{[]byte(`a`), "time: invalid duration a"},
	}

	for _, test := range tests {
		d := Duration(0)
		if err := d.UnmarshalText(test.input); err.Error() != test.err {
			t.Errorf("%s: got: %q, want error %q", test.input, err, test.err)
		}
	}
}

func TestURLDuration(t *testing.T) {
	tests := []struct {
		duration Duration
		want     time.Duration
	}{
		{Duration(0), time.Duration(0)},
		{Duration(time.Second), time.Duration(time.Second)},
		{Duration(-20 * time.Hour), time.Duration(-20 * time.Hour)},
	}

	for _, test := range tests {
		if response := test.duration.Duration(); response != test.want {
			t.Errorf("%q: got %q, want %q", test.duration, response, test.want)
		}
	}
}
