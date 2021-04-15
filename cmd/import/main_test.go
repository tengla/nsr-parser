package main

import (
	"os"
	"testing"
)

func TestKtx(t *testing.T) {
	t.Run("import", func(t *testing.T) {
		d, _ := os.Getwd()
		file := d + "/../../data/nsr.test.xml"
		importRecords(file, false)
	})
}
