/*
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-01 20:54:04
 * @Description:
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
 */
package math_test

import (
	"stl-go/math"
	"testing"
)

func TestMin(t *testing.T) {
	if ans := math.Min(12); ans != 12 {
		t.Errorf("min(12) should be 12, not %d", ans)
	}
	if ans := math.Min(12, 23, 44, 4, 6); ans != 4 {
		t.Errorf("min(12, 23, 44, 4, 6) should be 4, not %d", ans)
	}
	if ans := math.Min[float32](12, 23, 0.5, 88, -1); ans != -1 {
		t.Errorf("min(12, 23, 0.5, 88, -1) should be -1, not %f", ans)
	}
}

func TestMax(t *testing.T) {
	if ans := math.Max(12); ans != 12 {
		t.Errorf("max(12) should be 12, not %d", ans)
	}
	if ans := math.Max(12, 23, 44, 4, 6); ans != 44 {
		t.Errorf("max(12, 23, 44, 4, 6) should be 44, not %d", ans)
	}
	if ans := math.Max[float32](12, 23, 0.5, 88, -1); ans != 88 {
		t.Errorf("max(12, 23, 0.5, 88, -1) should be 88, not %f", ans)
	}
}
