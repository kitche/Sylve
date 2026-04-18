// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright (c) 2025 The FreeBSD Foundation.
//
// This software was developed by Hayzam Sherif <hayzam@alchemilla.io>
// of Alchemilla Ventures Pvt. Ltd. <hello@alchemilla.io>,
// under sponsorship from the FreeBSD Foundation.

package config

import (
	"reflect"
	"testing"

	"github.com/alchemillahq/sylve/internal"
)

func TestSylveVMPathHelpers(t *testing.T) {
	original := ParsedConfig
	t.Cleanup(func() {
		ParsedConfig = original
	})

	ParsedConfig = &internal.SylveConfig{
		ZFS: internal.ZFSConfig{
			SylveDataset:    "SylveRoot",
			SylveMountpoint: "/zroot/Sylve",
		},
	}

	if got := SylveVMDatasetRoot(); got != "SylveRoot/virtual-machines" {
		t.Fatalf("SylveVMDatasetRoot() = %q", got)
	}
	if got := SylveVMDatasetRootForPool("tank"); got != "tank/SylveRoot/virtual-machines" {
		t.Fatalf("SylveVMDatasetRootForPool() = %q", got)
	}
	if got := SylveVMRootDataset("tank", 101); got != "tank/SylveRoot/virtual-machines/101" {
		t.Fatalf("SylveVMRootDataset() = %q", got)
	}
	if got := SylveVMDatasetPath("tank", 101, "raw-1"); got != "tank/SylveRoot/virtual-machines/101/raw-1" {
		t.Fatalf("SylveVMDatasetPath() = %q", got)
	}
	if got := SylveVMRootMountpoint("tank", 101); got != "/zroot/Sylve/virtual-machines/101" {
		t.Fatalf("SylveVMRootMountpoint() = %q", got)
	}
	if got := SylveVMMountpointPath("tank", 101, "raw-1/1.img"); got != "/zroot/Sylve/virtual-machines/101/raw-1/1.img" {
		t.Fatalf("SylveVMMountpointPath() = %q", got)
	}
	if got := SylveVMZvolPath("tank", 101, "zvol-2"); got != "/dev/zvol/tank/SylveRoot/virtual-machines/101/zvol-2" {
		t.Fatalf("SylveVMZvolPath() = %q", got)
	}
	if got := SylveVMDatasetNeedle(); got != "/SylveRoot/virtual-machines/" {
		t.Fatalf("SylveVMDatasetNeedle() = %q", got)
	}

	gotPatterns := SylveVMDatasetLikePatterns(101)
	wantPatterns := []string{
		"%/SylveRoot/virtual-machines/101",
		"%/SylveRoot/virtual-machines/101/%",
		"%/SylveRoot/virtual-machines/101.%",
		"%/SylveRoot/virtual-machines/101_%",
	}
	if !reflect.DeepEqual(gotPatterns, wantPatterns) {
		t.Fatalf("SylveVMDatasetLikePatterns() = %#v, want %#v", gotPatterns, wantPatterns)
	}
}
