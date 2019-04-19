// Copyright © 2019 kPherox <admin@mail.kr-kp.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import "testing"

func TestAliasFlags(t *testing.T) {
    m := AliasFlags(map[string]string{
        "hoge": "fuga",
    })

    if m.Get("abc") != "abc" {
        t.Fatal("Failed return non alias")
    }
    if m.Get("hoge") != "fuga" {
        t.Fatal("Failed return original flag")
    }
}

func TestFlagAliasNormalization(t *testing.T) {
    if rootCmd.GlobalNormalizationFunc()(nil, "config") != "config-file" ||
      rootCmd.GlobalNormalizationFunc()(nil, "config-path") != "config-path" {
        t.Fatal("Failed flag alias normalization")
    }
}
