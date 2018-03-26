// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.9

package gcimporter

import "github.com/glycerine/gi/pkg/types"

func isAlias(obj *types.TypeName) bool {
	return false // there are no type aliases before Go 1.9
}
