// Copyright (c) 2016, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package marshal

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

func TestDurationMarshalJSON(t *testing.T) {
	tests := []struct {
		duration Duration
		want     []byte
	}{
		{Duration(-1), []byte(`"-1ns"`)},
		{Duration(0), []byte(`"0s"`)},
		{Duration(-time.Second), []byte(`"-1s"`)},
		{Duration(20 * time.Hour), []byte(`"20h0m0s"`)},
	}

	for _, test := range tests {
		result, err := json.Marshal(test.duration)
		if err != nil {
			t.Errorf("%#v: error: %s", test.duration, err)
		}
		if !bytes.Equal(result, test.want) {
			t.Errorf("%#v: got %s, want %s", test.duration, result, test.want)
		}
	}
}

func TestDurationMarshalText(t *testing.T) {
	tests := []struct {
		duration Duration
		want     []byte
	}{
		{Duration(-1), []byte(`-1ns`)},
		{Duration(0), []byte(`0s`)},
		{Duration(-time.Second), []byte(`-1s`)},
		{Duration(20 * time.Hour), []byte(`20h0m0s`)},
	}

	for _, test := range tests {
		result, err := test.duration.MarshalText()
		if err != nil {
			t.Errorf("%#v: error: %s", test.duration, err)
		}
		if !bytes.Equal(result, test.want) {
			t.Errorf("%#v: got %s, want %s", test.duration, result, test.want)
		}
	}
}

func TestDurationUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input []byte
		want  Duration
	}{
		{[]byte(`"1ns"`), Duration(1)},
		{[]byte(`"0"`), Duration(0)},
		{[]byte(`"20h"`), Duration(20 * time.Hour)},
	}

	for _, test := range tests {
		d := Duration(0)
		if err := json.Unmarshal(test.input, &d); err != nil {
			t.Errorf("%s: error: %s", test.input, err)
		}
		if d != test.want {
			t.Errorf("%s: got %q, want %q", test.input, d, test.want)
		}
	}
}

func TestDurationUnmarshalJSONError(t *testing.T) {
	tests := []struct {
		input []byte
		err   error
	}{
		{[]byte(`""`), ErrInvalidDuration},
		{[]byte(`0`), ErrInvalidDuration},
		{[]byte(`"10s"`), nil},
		{[]byte(`"5m"`), nil},
	}

	for _, test := range tests {
		d := Duration(0)
		if err := d.UnmarshalJSON(test.input); err != test.err {
			t.Errorf("%s: got: %q, want error %q", test.input, err, test.err)
		}
	}
}

func TestDurationUnmarshalText(t *testing.T) {
	tests := []struct {
		input []byte
		want  Duration
	}{
		{[]byte(`1ns`), Duration(1)},
		{[]byte(`0`), Duration(0)},
		{[]byte(`20h`), Duration(20 * time.Hour)},
	}

	for _, test := range tests {
		d := Duration(0)
		if err := d.UnmarshalText(test.input); err != nil {
			t.Errorf("%s: error: %s", test.input, err)
		}
		if d != test.want {
			t.Errorf("%s: got %q, want %q", test.input, d, test.want)
		}
	}
}

func TestDurationUnmarshalTextError(t *testing.T) {
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

func TestDurationDuration(t *testing.T) {
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
