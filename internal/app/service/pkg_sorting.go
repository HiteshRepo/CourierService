package service

import "github.com/hiteshpattanayak-tw/courier_service/internal/app/model"

type PackageSorting struct{}

func ProvidePackageSortingService() PackageSorting {
	return PackageSorting{}
}

func (ps PackageSorting) SortPackages(pkgs []model.Package) {
	for i := 0; i < len(pkgs); i++ {
		maxWeightPkg := i

		for j := i + 1; j < len(pkgs); j++ {
			if pkgs[maxWeightPkg].Weight < pkgs[j].Weight {
				maxWeightPkg = j
			}
		}

		pkgs[i], pkgs[maxWeightPkg] = pkgs[maxWeightPkg], pkgs[i]
	}

	for i := 0; i < len(pkgs)-1; i++ {
		if pkgs[i].Weight == pkgs[i+1].Weight && pkgs[i].DistanceInKm > pkgs[i+1].DistanceInKm {
			pkgs[i], pkgs[i+1] = pkgs[i+1], pkgs[i]
		}
	}
}
