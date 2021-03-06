/*
© Copyright IBM Corporation 2017

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
)

const mqmUID uint32 = 999
const mqmGID uint32 = 999

func createVolume(path string) error {
	dataPath := filepath.Join(path, "data")
	fi, err := os.Stat(dataPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dataPath, 0755)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	fi, err = os.Stat(dataPath)
	if err != nil {
		return err
	}
	sys := fi.Sys()
	if sys != nil && runtime.GOOS == "linux" {
		stat := sys.(*syscall.Stat_t)
		if stat.Uid != mqmUID || stat.Gid != mqmGID {
			err = os.Chown(dataPath, int(mqmUID), int(mqmGID))
			if err != nil {
				log.Printf("Error: Unable to change ownership of %v", dataPath)
				return err
			}
		}
	}
	return nil
}
