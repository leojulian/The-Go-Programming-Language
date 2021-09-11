package library

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed.")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty.")
	}

	m0 := &MusicEntry{
		"1",
		"Wake me up",
		"Avicci",
		"https://music.163.com/#/song?id=27713920",
		"MP3"}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("NewMusicManager.Add() failed.")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("NewMusicManager.Find() failed.")
	}
	if m.Id != m0.Id || m.Name != m0.Name || m.Artist != m0.Artist || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("NewMusicManager.Find() failed, Found item mismatch.")
	}

	m, err := mm.Get(0)
	if err != nil {
		t.Error("NewMusicManager.Get() failed. ", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("NewMusicManager.Remove() failed.")
	}
}
