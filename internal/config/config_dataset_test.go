// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright (c) 2025 The FreeBSD Foundation.
//
// This software was developed by Hayzam Sherif <hayzam@alchemilla.io>
// of Alchemilla Ventures Pvt. Ltd. <hello@alchemilla.io>,
// under sponsorship from the FreeBSD Foundation.

package config

import (
	"testing"

	"github.com/alchemillahq/sylve/internal"
)

func TestGetSylveDatasetRoot(t *testing.T) {
	original := ParsedConfig
	t.Cleanup(func() {
		ParsedConfig = original
	})

	tests := []struct {
		name   string
		config *internal.SylveConfig
		want   string
	}{
		{
			name:   "nil_config_uses_default",
			config: nil,
			want:   "sylve",
		},
		{
			name: "empty_config_uses_default",
			config: &internal.SylveConfig{
				ZFS: internal.ZFSConfig{},
			},
			want: "sylve",
		},
		{
			name: "trims_slashes",
			config: &internal.SylveConfig{
				ZFS: internal.ZFSConfig{
					SylveDataset: "/zroot/Sylve/",
				},
			},
			want: "zroot/Sylve",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ParsedConfig = tt.config
			if got := GetSylveDatasetRoot(); got != tt.want {
				t.Fatalf("GetSylveDatasetRoot() = %q, want %q", got, tt.want)
			}
		})
	}
}
