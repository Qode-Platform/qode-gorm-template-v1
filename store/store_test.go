package store

import "testing"


func TestCreateAndFind(t *testing.T) {
	db, err := Open("file::memory:")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := CreateUser(db, "Ada", "ada@example.com"); err != nil {
		t.Fatal(err)
	}
	got, err := FindByEmail(db, "ada@example.com")
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != "Ada" {
		t.Fatalf("name = %q", got.Name)
	}
}

func TestUniqueEmailIsEnforced(t *testing.T) {
	db, err := Open("file::memory:")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := CreateUser(db, "Ada", "dup@example.com"); err != nil {
		t.Fatal(err)
	}
	if _, err := CreateUser(db, "Grace", "dup@example.com"); err == nil {
		t.Fatal("expected the unique index to reject a duplicate email")
	}
}
