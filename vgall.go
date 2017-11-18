// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !minimal

package plot // import "gonum.org/v1/plot"

import (
	_ "gonum.org/v1/plot/vg/vgeps"
	_ "gonum.org/v1/plot/vg/vgimg"
	_ "gonum.org/v1/plot/vg/vgpdf"
	_ "gonum.org/v1/plot/vg/vgsvg"
)
