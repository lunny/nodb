package nodb

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/lunny/nodb/config"
)

func TestBinLog(t *testing.T) {
	cfg := new(config.Config)

	cfg.BinLog.MaxFileNum = 1
	cfg.BinLog.MaxFileSize = 1024
	cfg.DataDir = "/tmp/ledis_binlog"

	os.RemoveAll(cfg.DataDir)

	b, err := NewBinLog(cfg)
	if err != nil {
		t.Fatal(err)
	}

	if err := b.Log(make([]byte, 1024)); err != nil {
		t.Fatal(err)
	}

	if err := b.Log(make([]byte, 1024)); err != nil {
		t.Fatal(err)
	}

	if fs, err := ioutil.ReadDir(b.LogPath()); err != nil {
		t.Fatal(err)
	} else if len(fs) != 2 {
		t.Fatal(len(fs))
	}
}
