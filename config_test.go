// Copyright 2023 oyogames2023
//
// Licensed under the MIT License, you may not use this file except
// in compliance with the License. You may obtain a copy of the
// License at
//
//     https://opensource.org/license/mit
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// This code implementation is based on Google's Open Match framework
// with modifications inspired by the Zeus framework.
//
// reference to: https://github.com/googleforgames/open-match/interal/config

package zeus_config

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"testing"
	"time"
)

func TestReadConfigIgnoreRace(t *testing.T) {
	const defaultCfgName = "config.yaml"
	yaml := []byte(`
database.host: localhost
database.port: 3306
`)

	if err := os.WriteFile(defaultCfgName, yaml, 0666); err != nil {
		t.Fatalf("could not create config file: %s", err)
	}
	defer os.Remove(defaultCfgName)

	cfg, err := New(&ConfigParams{
		ConfigType: "yaml",
		ConfigName: "config",
		ConfigPath: []string{
			".",
		},
		WatchFunc: func(event fsnotify.Event) {
			log.Printf("Server configuration changed, operation: %v, filename: %s", event.Op, event.Name)
		},
	})

	if err != nil {
		t.Fatalf("cannot load config, %s", err)
	}

	if cfg.GetString("database.host") != "localhost" {
		t.Errorf("cfg.GetString('database.host') = %s, expected 'localhost'", cfg.GetString("database.host"))
	}

	if cfg.GetInt32("database.port") != 3306 {
		t.Errorf("cfg.GetString('database.port') = %s, expected '3306'", cfg.GetString("database.port"))
	}

	yaml = []byte(`
database.host: 127.0.0.1
database.port: 3306
`)
	if err := os.WriteFile(defaultCfgName, yaml, 0666); err != nil {
		t.Fatalf("could not update config file: %s", err)
	}

	time.Sleep(time.Second)

	if cfg.GetString("database.host") != "127.0.0.1" {
		t.Errorf("cfg.GetString('database.host') = %s, expected '127.0.0.1'", cfg.GetString("database.host"))
	}
}
