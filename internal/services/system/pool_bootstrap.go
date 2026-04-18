// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright (c) 2025 The FreeBSD Foundation.
//
// This software was developed by Hayzam Sherif <hayzam@alchemilla.io>
// of Alchemilla Ventures Pvt. Ltd. <hello@alchemilla.io>,
// under sponsorship from the FreeBSD Foundation.

package system

import (
	"context"
	"fmt"
	"strings"

	"github.com/alchemillahq/gzfs"
	"github.com/alchemillahq/sylve/internal/config"
)

func requiredSylveDatasets(rootDataset string) []string {
	return []string{
		rootDataset,
		fmt.Sprintf("%s/virtual-machines", rootDataset),
		fmt.Sprintf("%s/jails", rootDataset),
		fmt.Sprintf("%s/bootstraps", rootDataset),
	}
}

func requiredSylveDatasetMountpoint(rootDataset, rootMountpoint, dataset string) string {
	suffix := strings.TrimPrefix(dataset, rootDataset)
	suffix = strings.TrimPrefix(suffix, "/")
	if suffix == "" {
		return rootMountpoint
	}

	return fmt.Sprintf("%s/%s", strings.TrimRight(rootMountpoint, "/"), suffix)
}

func (s *Service) ensureSylveDatasetsOnPool(ctx context.Context, poolName string) ([]*gzfs.Dataset, error) {
	var created []*gzfs.Dataset
	rootDataset := config.GetSylveDatasetRoot()
	rootMountpoint := config.GetSylveMountpointRoot(poolName)

	for _, dataset := range requiredSylveDatasets(rootDataset) {
		fullDatasetName := fmt.Sprintf("%s/%s", poolName, dataset)
		mountpoint := requiredSylveDatasetMountpoint(rootDataset, rootMountpoint, dataset)

		found, err := s.GZFS.ZFS.Get(ctx, fullDatasetName, false)
		if err != nil && !strings.Contains(strings.ToLower(err.Error()), "does not exist") {
			return nil, fmt.Errorf("error_checking_dataset_%s: %w", fullDatasetName, err)
		}

		if found != nil {
			if found.Mountpoint != mountpoint {
				if err := s.GZFS.ZFS.EditFilesystem(ctx, fullDatasetName, map[string]string{
					"mountpoint": mountpoint,
				}); err != nil {
					return nil, fmt.Errorf("error_fixing_mountpoint_%s: %w", fullDatasetName, err)
				}
			}
			continue
		}

		newDataset, err := s.GZFS.ZFS.CreateFilesystem(ctx, fullDatasetName, map[string]string{
			"mountpoint": mountpoint,
		})
		if err != nil {
			return nil, fmt.Errorf("error_creating_dataset_%s: %w", fullDatasetName, err)
		}

		created = append(created, newDataset)
	}

	return created, nil
}
