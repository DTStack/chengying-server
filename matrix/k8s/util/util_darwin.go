package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var randmu sync.Mutex
var rand uint32

func CreateTempScript(content string, prefix string) (path string, err error) {
	f, err := ioutil.TempFile("", prefix)
	if err != nil {
		return "", err
	}

	if _, err = f.WriteString(content); err != nil {
		f.Close()
		os.Remove(f.Name())
		return "", err
	}
	if err = os.Chmod(f.Name(), 0500); err != nil {
		f.Close()
		os.Remove(f.Name())
		return "", err
	}

	f.Close()
	return f.Name(), nil
}
func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}
func NewFile(dir, fname string) (f *os.File, err error) {
	if dir == "" {
		dir = os.TempDir()
	}
	nconflict := 0
	for i := 0; i < 10000; i++ {
		name := filepath.Join(dir, fname)
		f, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
		if os.IsExist(err) {
			os.Rename(name, name+nextRandom())
			if nconflict++; nconflict > 10 {
				randmu.Lock()
				rand = reseed()
				randmu.Unlock()
			}
			continue
		}
		break
	}
	return
}

func nextRandom() string {
	randmu.Lock()
	r := rand
	if r == 0 {
		r = reseed()
	}
	r = r*1664525 + 1013904223 // constants from Numerical Recipes
	rand = r
	randmu.Unlock()
	return strconv.Itoa(int(1e9 + r%1e9))[1:]
}
