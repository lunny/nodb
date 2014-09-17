package nodb

import "testing"

func testTxRollback(t *testing.T, db *DB) {
	var err error
	key1 := []byte("tx_key1")
	key2 := []byte("tx_key2")
	field2 := []byte("tx_field2")

	err = db.Set(key1, []byte("value"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.HSet(key2, field2, []byte("value"))
	if err != nil {
		t.Fatal(err)
	}

	var tx *Tx
	tx, err = db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	defer tx.Rollback()

	err = tx.Set(key1, []byte("1"))

	if err != nil {
		t.Fatal(err)
	}

	_, err = tx.HSet(key2, field2, []byte("2"))

	if err != nil {
		t.Fatal(err)
	}

	_, err = tx.HSet([]byte("no_key"), field2, []byte("2"))

	if err != nil {
		t.Fatal(err)
	}

	if v, err := tx.Get(key1); err != nil {
		t.Fatal(err)
	} else if string(v) != "1" {
		t.Fatal(string(v))
	}

	if v, err := tx.HGet(key2, field2); err != nil {
		t.Fatal(err)
	} else if string(v) != "2" {
		t.Fatal(string(v))
	}

	err = tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}

	if v, err := db.Get(key1); err != nil {
		t.Fatal(err)
	} else if string(v) != "value" {
		t.Fatal(string(v))
	}

	if v, err := db.HGet(key2, field2); err != nil {
		t.Fatal(err)
	} else if string(v) != "value" {
		t.Fatal(string(v))
	}
}

func testTxCommit(t *testing.T, db *DB) {
	var err error
	key1 := []byte("tx_key1")
	key2 := []byte("tx_key2")
	field2 := []byte("tx_field2")

	err = db.Set(key1, []byte("value"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.HSet(key2, field2, []byte("value"))
	if err != nil {
		t.Fatal(err)
	}

	var tx *Tx
	tx, err = db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	defer tx.Rollback()

	err = tx.Set(key1, []byte("1"))

	if err != nil {
		t.Fatal(err)
	}

	_, err = tx.HSet(key2, field2, []byte("2"))

	if err != nil {
		t.Fatal(err)
	}

	if v, err := tx.Get(key1); err != nil {
		t.Fatal(err)
	} else if string(v) != "1" {
		t.Fatal(string(v))
	}

	if v, err := tx.HGet(key2, field2); err != nil {
		t.Fatal(err)
	} else if string(v) != "2" {
		t.Fatal(string(v))
	}

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}

	if v, err := db.Get(key1); err != nil {
		t.Fatal(err)
	} else if string(v) != "1" {
		t.Fatal(string(v))
	}

	if v, err := db.HGet(key2, field2); err != nil {
		t.Fatal(err)
	} else if string(v) != "2" {
		t.Fatal(string(v))
	}
}

func testTxSelect(t *testing.T, db *DB) {
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	defer tx.Rollback()

	tx.Set([]byte("tx_select_1"), []byte("a"))

	tx.Select(1)

	tx.Set([]byte("tx_select_2"), []byte("b"))

	if err = tx.Commit(); err != nil {
		t.Fatal(err)
	}

	if v, err := db.Get([]byte("tx_select_1")); err != nil {
		t.Fatal(err)
	} else if string(v) != "a" {
		t.Fatal(string(v))
	}

	if v, err := db.Get([]byte("tx_select_2")); err != nil {
		t.Fatal(err)
	} else if v != nil {
		t.Fatal("must nil")
	}

	db, _ = db.l.Select(1)

	if v, err := db.Get([]byte("tx_select_2")); err != nil {
		t.Fatal(err)
	} else if string(v) != "b" {
		t.Fatal(string(v))
	}

	if v, err := db.Get([]byte("tx_select_1")); err != nil {
		t.Fatal(err)
	} else if v != nil {
		t.Fatal("must nil")
	}
}
