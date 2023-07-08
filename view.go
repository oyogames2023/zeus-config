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
	"github.com/spf13/viper"
	"time"
)

// View is a read-only view of the configuration.
// New accessors from Viper should be added here.
type View interface {
	IsSet(string) bool
	GetString(string) string
	GetInt(string) int
	GetInt64(string) int64
	GetFloat64(string) float64
	GetStringSlice(string) []string
	GetBool(string) bool
	GetDuration(string) time.Duration
}

// Mutable is a read-write view of the configuration.
type Mutable interface {
	Set(string, interface{})
	View
}

// Sub returns a subset of configuration filtered by the key.
func Sub(v View, key string) View {
	cfg, ok := v.(*viper.Viper)
	if ok {
		return cfg.Sub(key)
	}
	return nil
}
