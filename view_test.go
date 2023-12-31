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
	"testing"

	"github.com/spf13/viper"
)

func TestSubFromViper(t *testing.T) {
	v := viper.New()
	v.Set("a", "a")
	v.Set("b", "b")
	v.Set("c", "c")
	v.Set("a.a", "a.a")
	v.Set("a.b", "a.b")
	av := Sub(v, "a")
	if av == nil {
		t.Fatalf("Sub(%v, 'a') => %v", v, av)
	}

	if av.GetString("a") != "a.a" {
		t.Errorf("av.GetString('a') = %s, expected 'a.a'", av.GetString("a"))
	}
	if av.GetString("a.a") != "" {
		t.Errorf("av.GetString('a.a') = %s, expected ''", av.GetString("a.a"))
	}
	if av.GetString("b") != "a.b" {
		t.Errorf("av.GetString('b') = %s, expected 'a.b'", av.GetString("b"))
	}
	if av.GetString("c") != "" {
		t.Errorf("av.GetString('c') = %s, expected ''", av.GetString(""))
	}
}
