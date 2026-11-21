package sirius

import (
	"os"
	"path/filepath"
	"testing"
)

func openTestJSON(t *testing.T, content string) *SJson {
	t.Helper()
	path := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write test file: %v", err)
	}
	document, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	return document
}

func TestGettersReturnZeroForMalformedIndexKeys(t *testing.T) {
	document := openTestJSON(t, `{"a":"x","list":[{"v":1}]}`)
	keys := []string{"a[", "nope[", "x]y[", "list[-2]", "list[-2].v", "list[0", "list[]"}

	for _, key := range keys {
		t.Run(key, func(t *testing.T) {
			if got := document.Get(key); got != nil {
				t.Errorf("Get(%q) = %v, want nil", key, got)
			}
			if got := document.GetString(key); got != "" {
				t.Errorf("GetString(%q) = %q, want empty", key, got)
			}
			if got := document.GetInt(key); got != 0 {
				t.Errorf("GetInt(%q) = %d, want 0", key, got)
			}
			if got := document.GetInt64(key); got != 0 {
				t.Errorf("GetInt64(%q) = %d, want 0", key, got)
			}
			if got := document.GetFloat64(key); got != 0 {
				t.Errorf("GetFloat64(%q) = %v, want 0", key, got)
			}
			if got := document.GetBool(key); got {
				t.Errorf("GetBool(%q) = true, want false", key)
			}
			if got := document.GetArrayMaps(key); got != nil {
				t.Errorf("GetArrayMaps(%q) = %v, want nil", key, got)
			}
		})
	}
}

func TestGetIntReadsValidIndexKey(t *testing.T) {
	document := openTestJSON(t, `{"list":[{"v":1}]}`)

	got := document.GetInt("list[0].v")

	if got != 1 {
		t.Errorf("GetInt(list[0].v) = %d, want 1", got)
	}
}

func TestGetStringReturnsEmptyForMissingMiddleSegment(t *testing.T) {
	document := openTestJSON(t, `{"databases":{"password":"PARENT","prod":{"password":"p"}},"a":"x"}`)
	keys := []string{"databases.staging.password", "nothere.databases.prod.password", "a.b"}

	for _, key := range keys {
		t.Run(key, func(t *testing.T) {
			got := document.GetString(key)

			if got != "" {
				t.Errorf("GetString(%q) = %q, want empty", key, got)
			}
		})
	}
}

func TestGetBoolReturnsFalseForMissingMiddleSegment(t *testing.T) {
	document := openTestJSON(t, `{"auth":{"require_mfa":true,"admin":{}}}`)

	got := document.GetBool("auth.missing.require_mfa")

	if got {
		t.Error("GetBool(auth.missing.require_mfa) = true, want false")
	}
}

func TestGetReturnsNilForMissingMiddleSegment(t *testing.T) {
	document := openTestJSON(t, `{"databases":{"password":"PARENT","prod":{"password":"p"}}}`)

	got := document.Get("databases.staging.password")

	if got != nil {
		t.Errorf("Get(databases.staging.password) = %v, want nil", got)
	}
}

func TestGetStringResolvesPathBesideParentValue(t *testing.T) {
	document := openTestJSON(t, `{"databases":{"password":"PARENT","prod":{"password":"p"}}}`)

	got := document.GetString("databases.prod.password")

	if got != "p" {
		t.Errorf("GetString(databases.prod.password) = %q, want p", got)
	}
}

func TestGetBoolResolvesExistingPath(t *testing.T) {
	document := openTestJSON(t, `{"auth":{"require_mfa":true,"admin":{}}}`)

	got := document.GetBool("auth.require_mfa")

	if !got {
		t.Error("GetBool(auth.require_mfa) = false, want true")
	}
}

func TestGetReturnsObjectForExistingPath(t *testing.T) {
	document := openTestJSON(t, `{"auth":{"require_mfa":true,"admin":{}}}`)

	got, ok := document.Get("auth.admin").(jsonNode)

	if !ok || len(got) != 0 {
		t.Errorf("Get(auth.admin) = %v, want an empty object", got)
	}
}
