//go:build darwin
// +build darwin

package memory

import (
	gosigar "github.com/cloudfoundry/gosigar"
)

// GetRAMAppBytes returns the app ram usage in bytes
func GetRAMAppBytes() uint64 {
	mem := &gosigar.ProcMem{}
	if err := mem.Get(0); err != nil {
		return 0
	}
	return mem.Resident
}

// GetRAMSystemBytes returns the system total ram in bytes
func GetRAMSystemBytes() uint64 {
	mem := &gosigar.Mem{}
	if err := mem.Get(); err != nil {
		return 0
	}
	return mem.Total
}
