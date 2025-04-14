/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2024 Daniel Ericsson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 */

package version

import (
	"testing"
)

func TestVersionAlloc(t *testing.T) {
	expect := 7
	got := int(testing.AllocsPerRun(5, func() {
		Version()
	}))
	if got != expect {
		t.Errorf("got %d allocations, expected %d", got, expect)
	}
}

func TestEmptyVersion(t *testing.T) {
	v := Version()
	if v == "" {
		t.Error("Empty Version")
	}
}

func TestDevel(t *testing.T) {
	v := Version()
	if v != development {
		t.Errorf("expected '%s' got '%s'", development, v)
	}
}

func TestOverride(t *testing.T) {
	expect := "1.2.3"
	versionOverride = expect

	v := Version()
	if v != expect {
		t.Errorf("expected '%s' got '%s'", expect, v)
	}
	versionOverride = ""
}
