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
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ConfigParams struct {
	ConfigType string
	ConfigPath []string
	ConfigName string
	WatchFunc  func(fsnotify.Event)
}

// New creates a new Viper object with ConfigParams and parses the config file.
func New(params *ConfigParams) (*viper.Viper, error) {
	cfg := viper.New()
	cfg.SetConfigType(params.ConfigType)
	for _, path := range params.ConfigPath {
		cfg.AddConfigPath(path)
	}
	cfg.SetConfigName(params.ConfigName)
	err := cfg.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("fatal error reading default config file, desc: %s", err.Error())
	}
	if params.WatchFunc != nil {
		cfg.WatchConfig()
		cfg.OnConfigChange(params.WatchFunc)
	}
	return cfg, nil
}
