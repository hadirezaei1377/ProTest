package unit_test_test

import (
	"ProTest/unit_test"
	"testing"
)

func TestLoadFromTrustedLocationOrNil(t *testing.T) {
	location := unit_test.LoadFromTrustedLocationOrNil("Asia/Iran/Tehran")
	if location == nil {
		t.Fatal("the location is nil for Asia/Iran/Tehran")
	}

	/*
		location := unit_test.LoadFromTrustedLocationOrNil("Tehran")
		if location == nil {
			t.Fatal("the location is nil for Tehran")
		}
	*/
}
