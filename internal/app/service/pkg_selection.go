package service

import "github.com/hiteshpattanayak-tw/courier_service/internal/app/model"

type PackageSelectionService struct{}

func ProvidePackageSelectionService() PackageSelectionService {
	return PackageSelectionService{}
}

func (ps PackageSelectionService) SelectPackages(pkgs []model.Package, limit float32) []model.Package {
	if len(pkgs) == 0 {
		return nil
	}

	if len(pkgs) == 1 {
		return pkgs
	}

	ps.SortPackages(pkgs)

	allCombinations := make([][]model.Package, 0)
	combinations := make([]model.Package, 0)

	copyPkgs := make([]model.Package, len(pkgs))
	copy(copyPkgs, pkgs)

	ps.getAllPkgCombinations(copyPkgs, &allCombinations, combinations, 0, limit)

	return ps.getPkgsWithHighestNumberAndHighestWeight(allCombinations)
}

func (ps PackageSelectionService) SortPackages(pkgs []model.Package) {
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

func (ps PackageSelectionService) getAllPkgCombinations(pkgs []model.Package, allCombinations *[][]model.Package, combination []model.Package, idx int, limit float32) {
	lightestWt := ps.getLightestPkg(ps.getRemainingPkgs(pkgs, combination))
	if limit <= lightestWt {
		if !ps.isPackagesPresent(combination, *allCombinations) {
			c := make([]model.Package, len(combination))
			copy(c, combination)
			*allCombinations = append(*allCombinations, c)
		}
		return
	}

	for i := idx; i < len(pkgs); i++ {
		if float32(pkgs[i].Weight) > limit {
			break
		}
		combination = append(combination, pkgs[i])
		ps.getAllPkgCombinations(pkgs, allCombinations, combination, i+1, limit-float32(pkgs[i].Weight))
		combination = combination[0 : len(combination)-1]
	}
}

func (ps PackageSelectionService) getLightestPkg(pkgs []model.Package) float32 {
	if len(pkgs) == 0 {
		return -1
	}

	lightest := pkgs[0]
	for _, p := range pkgs {
		if p.Weight < lightest.Weight {
			lightest = p
		}
	}
	return float32(lightest.Weight)
}

func (ps PackageSelectionService) getRemainingPkgs(pkgs []model.Package, currPkgs []model.Package) []model.Package {
	remainingPkgs := make([]model.Package, 0)
	for _, p := range pkgs {
		if !ps.isPackagePresent(p, currPkgs) {
			remainingPkgs = append(remainingPkgs, p)
		}
	}

	return remainingPkgs
}

func (ps PackageSelectionService) isPackagePresent(pkg model.Package, currPkgs []model.Package) bool {
	for _, p := range currPkgs {
		if p.Id == pkg.Id {
			return true
		}
	}

	return false
}

func (ps PackageSelectionService) isPackagesPresent(pkgs []model.Package, allPkgCombinations [][]model.Package) bool {
	for _, pcs := range allPkgCombinations {
		if len(ps.getRemainingPkgs(pcs, pkgs)) == 0 {
			return true
		}
	}

	return false
}

func (ps PackageSelectionService) getPkgsWithHighestNumberAndHighestWeight(allPkgCombinations [][]model.Package) []model.Package {
	highest := float32(0)
	longest := 0
	var selectedPkgs []model.Package
	for _, pkgs := range allPkgCombinations {
		if wt := ps.getPkgsCombinedWeight(pkgs); wt > highest && len(pkgs) >= longest {
			longest = len(pkgs)
			highest = wt
			selectedPkgs = pkgs
		}
	}

	return selectedPkgs
}

func (ps PackageSelectionService) getPkgsCombinedWeight(pkgs []model.Package) float32 {
	wt := 0
	for _, p := range pkgs {
		wt += p.Weight
	}
	return float32(wt)
}
