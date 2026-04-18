// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright (c) 2025 The FreeBSD Foundation.
//
// This software was developed by Hayzam Sherif <hayzam@alchemilla.io>
// of Alchemilla Ventures Pvt. Ltd. <hello@alchemilla.io>,
// under sponsorship from the FreeBSD Foundation.

package system

import (
	"reflect"
	"testing"
)

func TestRequiredSylveDatasets(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		root string
		want []string
	}{
		{
			name: "default_like_root",
			root: "sylve",
			want: []string{
				"sylve",
				"sylve/virtual-machines",
				"sylve/jails",
				"sylve/bootstraps",
			},
		},
		{
			name: "custom_root",
			root: "Sylve",
			want: []string{
				"Sylve",
				"Sylve/virtual-machines",
				"Sylve/jails",
				"Sylve/bootstraps",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := requiredSylveDatasets(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("requiredSylveDatasets(%q) = %#v, want %#v", tt.root, got, tt.want)
			}
		})
	}
}
