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

func TestBytesMarshalJSON(t *testing.T) {
	tests := []struct {
		bytes Bytes
		want  []byte
	}{
		{Bytes(1), []byte(`"1B"`)},
		{Bytes(1000), []byte(`"1000B"`)},
		{Bytes(1024), []byte(`"1.0kB"`)},
		{Bytes(1024 * 1024), []byte(`"1.0MB"`)},
		{Bytes(1024 * 1024 * 1024), []byte(`"1.0GB"`)},
		{Bytes(1024 * 1024 * 1024 * 1024), []byte(`"1.0TB"`)},
		{Bytes(1024 * 1024 * 1024 * 1024 * 1024), []byte(`"1.0PB"`)},
		{Bytes(1024 * 1024 * 1024 * 1024 * 1024 * 1024), []byte(`"1.0EB"`)},
		{Bytes(12 * 1024), []byte(`"12kB"`)},
		{Bytes(99 * 1024), []byte(`"99kB"`)},
		{Bytes(250 * 1024), []byte(`"250kB"`)},
		{Bytes(1000 * 1024 * 1024), []byte(`"1000MB"`)},
	}

	for _, test := range tests {
		result, err := json.Marshal(test.bytes)
		if err != nil {
			t.Errorf("%#v: error: %s", test.bytes, err)
		}
		if !bytes.Equal(result, test.want) {
			t.Errorf("%#v: got %s, want %s", test.bytes, result, test.want)
		}
	}
}

func TestBytesMarshalText(t *testing.T) {
	tests := []struct {
		bytes Bytes
		want  []byte
	}{
		{Bytes(1), []byte(`1B`)},
		{Bytes(1000), []byte(`1000B`)},
		{Bytes(1024), []byte(`1.0kB`)},
		{Bytes(1024 * 1024), []byte(`1.0MB`)},
		{Bytes(1024 * 1024 * 1024), []byte(`1.0GB`)},
		{Bytes(1024 * 1024 * 1024 * 1024), []byte(`1.0TB`)},
		{Bytes(1024 * 1024 * 1024 * 1024 * 1024), []byte(`1.0PB`)},
		{Bytes(1024 * 1024 * 1024 * 1024 * 1024 * 1024), []byte(`1.0EB`)},
		{Bytes(12 * 1024), []byte(`12kB`)},
		{Bytes(99 * 1024), []byte(`99kB`)},
		{Bytes(250 * 1024), []byte(`250kB`)},
		{Bytes(1000 * 1024 * 1024), []byte(`1000MB`)},
	}

	for _, test := range tests {
		result, err := test.bytes.MarshalText()
		if err != nil {
			t.Errorf("%#v: error: %s", test.bytes, err)
		}
		if !bytes.Equal(result, test.want) {
			t.Errorf("%#v: got %s, want %s", test.bytes, result, test.want)
		}
	}
}

func TestBytesUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input []byte
		want  Bytes
	}{
		{[]byte(`"1B"`), Bytes(1)},
		{[]byte(`"1000B"`), Bytes(1000)},
		{[]byte(`"1.0kB"`), Bytes(1024)},
		{[]byte(`"1.0MB"`), Bytes(1024 * 1024)},
		{[]byte(`"1.0GB"`), Bytes(1024 * 1024 * 1024)},
		{[]byte(`"1.0TB"`), Bytes(1024 * 1024 * 1024 * 1024)},
		{[]byte(`"1.0PB"`), Bytes(1024 * 1024 * 1024 * 1024 * 1024)},
		{[]byte(`"1.0EB"`), Bytes(1024 * 1024 * 1024 * 1024 * 1024 * 1024)},
		{[]byte(`"12kB"`), Bytes(12 * 1024)},
		{[]byte(`"99kB"`), Bytes(99 * 1024)},
		{[]byte(`"250kB"`), Bytes(250 * 1024)},
		{[]byte(`"1000MB"`), Bytes(1000 * 1024 * 1024)},
	}

	for _, test := range tests {
		b := Bytes(0)
		err := json.Unmarshal(test.input, &b)
		if err != nil {
			t.Errorf("%s: error: %s", test.input, err)
		}
		if b != test.want {
			t.Errorf("%s: got %v, want %v", test.input, b, test.want)
		}
	}
}

func TestBytesUnmarshalText(t *testing.T) {
	tests := []struct {
		input []byte
		want  Bytes
	}{
		{[]byte(`1B`), Bytes(1)},
		{[]byte(`1000B`), Bytes(1000)},
		{[]byte(`1.0kB`), Bytes(1024)},
		{[]byte(`1.0MB`), Bytes(1024 * 1024)},
		{[]byte(`1.0GB`), Bytes(1024 * 1024 * 1024)},
		{[]byte(`1.0TB`), Bytes(1024 * 1024 * 1024 * 1024)},
		{[]byte(`1.0PB`), Bytes(1024 * 1024 * 1024 * 1024 * 1024)},
		{[]byte(`1.0EB`), Bytes(1024 * 1024 * 1024 * 1024 * 1024 * 1024)},
		{[]byte(`12kB`), Bytes(12 * 1024)},
		{[]byte(`99kB`), Bytes(99 * 1024)},
		{[]byte(`250kB`), Bytes(250 * 1024)},
		{[]byte(`1000MB`), Bytes(1000 * 1024 * 1024)},
	}

	for _, test := range tests {
		b := Bytes(0)
		err := b.UnmarshalText(test.input)
		if err != nil {
			t.Errorf("%s: error: %s", test.input, err)
		}
		if b != test.want {
			t.Errorf("%s: got %v, want %v", test.input, b, test.want)
		}
	}
}

func TestBytesBytes(t *testing.T) {
	tests := []struct {
		bytes Bytes
		want  uint64
	}{
		{Bytes(1), 1},
		{Bytes(1024), 1024},
	}

	for _, test := range tests {
		if response := test.bytes.Bytes(); response != test.want {
			t.Errorf("%q: got %q, want %q", test.bytes, response, test.want)
		}
	}
}
