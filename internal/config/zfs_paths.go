// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright (c) 2025 The FreeBSD Foundation.
//
// This software was developed by Hayzam Sherif <hayzam@alchemilla.io>
// of Alchemilla Ventures Pvt. Ltd. <hello@alchemilla.io>,
// under sponsorship from the FreeBSD Foundation.

package config

import (
	"fmt"
	"strings"
)

func SylveVMDatasetRoot() string {
	return fmt.Sprintf("%s/virtual-machines", GetSylveDatasetRoot())
}

func SylveVMDatasetRootForPool(pool string) string {
	return fmt.Sprintf("%s/%s", strings.TrimSpace(pool), SylveVMDatasetRoot())
}

func SylveVMRootDataset(pool string, rid uint) string {
	return fmt.Sprintf("%s/%d", SylveVMDatasetRootForPool(pool), rid)
}

func SylveVMDatasetPath(pool string, rid uint, suffix string) string {
	suffix = strings.TrimPrefix(strings.TrimSpace(suffix), "/")
	if suffix == "" {
		return SylveVMRootDataset(pool, rid)
	}

	return fmt.Sprintf("%s/%s", SylveVMRootDataset(pool, rid), suffix)
}

func SylveVMMountpointRoot(pool string) string {
	return fmt.Sprintf("%s/virtual-machines", strings.TrimRight(GetSylveMountpointRoot(pool), "/"))
}

func SylveVMRootMountpoint(pool string, rid uint) string {
	return fmt.Sprintf("%s/%d", SylveVMMountpointRoot(pool), rid)
}

func SylveVMMountpointPath(pool string, rid uint, suffix string) string {
	suffix = strings.TrimPrefix(strings.TrimSpace(suffix), "/")
	if suffix == "" {
		return SylveVMRootMountpoint(pool, rid)
	}

	return fmt.Sprintf("%s/%s", SylveVMRootMountpoint(pool, rid), suffix)
}

func SylveVMZvolPath(pool string, rid uint, zvolName string) string {
	return fmt.Sprintf("/dev/zvol/%s", SylveVMDatasetPath(pool, rid, zvolName))
}

func SylveVMDatasetNeedle() string {
	return fmt.Sprintf("/%s/", SylveVMDatasetRoot())
}

func SylveVMDatasetLikePatterns(rid uint) []string {
	root := SylveVMDatasetRoot()
	return []string{
		fmt.Sprintf("%%/%s/%d", root, rid),
		fmt.Sprintf("%%/%s/%d/%%", root, rid),
		fmt.Sprintf("%%/%s/%d.%%", root, rid),
		fmt.Sprintf("%%/%s/%d_%%", root, rid),
	}
}
