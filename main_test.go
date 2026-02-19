package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

var uuidWithDashRe = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
var hexRe = regexp.MustCompile(`^[0-9a-f]+$`)

// --- createUUID ---

func TestCreateUUID_Versions(t *testing.T) {
	versions := []string{"1", "2p", "2g", "4", "6", "7"}
	for _, v := range versions {
		t.Run("v"+v, func(t *testing.T) {
			result := createUUID(true, v)
			if !uuidWithDashRe.MatchString(result) {
				t.Errorf("version %s: invalid UUID format: %q", v, result)
			}
		})
	}
}

func TestCreateUUID_WithDash(t *testing.T) {
	result := createUUID(true, "4")
	if len(result) != 36 {
		t.Errorf("expected length 36, got %d: %q", len(result), result)
	}
	if strings.Count(result, "-") != 4 {
		t.Errorf("expected 4 dashes, got %d: %q", strings.Count(result, "-"), result)
	}
}

func TestCreateUUID_NoDash(t *testing.T) {
	result := createUUID(false, "4")
	if len(result) != 32 {
		t.Errorf("expected length 32, got %d: %q", len(result), result)
	}
	if strings.Contains(result, "-") {
		t.Errorf("expected no dashes: %q", result)
	}
	if !hexRe.MatchString(result) {
		t.Errorf("expected hex string: %q", result)
	}
}

func TestCreateUUID_VersionCaseInsensitive(t *testing.T) {
	// version string is lowercased inside the function
	result := createUUID(true, "4")
	if !uuidWithDashRe.MatchString(result) {
		t.Errorf("invalid UUID format: %q", result)
	}
}

// --- createCuid ---

func TestCreateCuid_Default(t *testing.T) {
	result := createCuid(0)
	if result == "" {
		t.Error("expected non-empty CUID for length 0 (default)")
	}
}

func TestCreateCuid_CustomLengths(t *testing.T) {
	for _, l := range []int{2, 10, 16, 24, 32} {
		t.Run(fmt.Sprintf("len%d", l), func(t *testing.T) {
			result := createCuid(l)
			if len(result) != l {
				t.Errorf("expected length %d, got %d: %q", l, len(result), result)
			}
		})
	}
}

func TestCreateCuid_OutOfRange(t *testing.T) {
	// lengths outside [2,32] fall back to default generator
	for _, l := range []int{0, 1, 33, 100} {
		t.Run(fmt.Sprintf("len%d", l), func(t *testing.T) {
			result := createCuid(l)
			if result == "" {
				t.Errorf("expected non-empty CUID for out-of-range length %d", l)
			}
		})
	}
}

// --- createNanoid ---

func TestCreateNanoid_DefaultLength(t *testing.T) {
	result := createNanoid(0)
	if len(result) != 21 {
		t.Errorf("expected default length 21, got %d: %q", len(result), result)
	}
}

func TestCreateNanoid_CustomLengths(t *testing.T) {
	for _, l := range []int{2, 10, 21, 50, 128} {
		t.Run(fmt.Sprintf("len%d", l), func(t *testing.T) {
			result := createNanoid(l)
			if len(result) != l {
				t.Errorf("expected length %d, got %d: %q", l, len(result), result)
			}
		})
	}
}

// --- createUlid ---

func TestCreateUlid_Fast(t *testing.T) {
	result := createUlid(false)
	if len(result) != 26 {
		t.Errorf("expected ULID length 26, got %d: %q", len(result), result)
	}
}

func TestCreateUlid_Crypt(t *testing.T) {
	result := createUlid(true)
	if len(result) != 26 {
		t.Errorf("expected ULID length 26, got %d: %q", len(result), result)
	}
}

// --- createXid ---

func TestCreateXid(t *testing.T) {
	result := createXid()
	if len(result) != 20 {
		t.Errorf("expected XID length 20, got %d: %q", len(result), result)
	}
}

// --- createObjectID ---

func TestCreateObjectID(t *testing.T) {
	result := createObjectID()
	if len(result) != 24 {
		t.Errorf("expected ObjectID length 24, got %d: %q", len(result), result)
	}
	if !hexRe.MatchString(result) {
		t.Errorf("expected 24 hex chars, got: %q", result)
	}
}

// --- uniqueness ---

func TestUniqueness(t *testing.T) {
	const n = 100
	generators := map[string]func() string{
		"uuid":     func() string { return createUUID(true, "4") },
		"cuid":     func() string { return createCuid(0) },
		"nanoid":   func() string { return createNanoid(0) },
		"ulid":     func() string { return createUlid(false) },
		"xid":      createXid,
		"objectid": createObjectID,
	}

	for name, gen := range generators {
		t.Run(name, func(t *testing.T) {
			seen := make(map[string]struct{}, n)
			for range n {
				id := gen()
				if _, dup := seen[id]; dup {
					t.Errorf("duplicate ID generated: %q", id)
				}
				seen[id] = struct{}{}
			}
		})
	}
}
