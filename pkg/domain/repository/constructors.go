// Code generated by tools/update-providers.go; DO NOT EDIT.
package repository

import "github.com/gnulinuxindia/internet-chowkidar/ent"

func ProvideBlocksRepository(
	db *ent.Client,
) BlocksRepository {
	return &blocksRepositoryImpl{
		db: db,
	}
}

func ProvideCounterRepository(
	db *ent.Client,
) CounterRepository {
	return &counterRepositoryImpl{
		db: db,
	}
}

func ProvideIspRepository(
	db *ent.Client,
) IspRepository {
	return &ispRepositoryImpl{
		db: db,
	}
}

func ProvideReportsRepository(
	db *ent.Client,
) ReportsRepository {
	return &reportsRepositoryImpl{
		db: db,
	}
}

func ProvideSitesRepository(
	db *ent.Client,
) SitesRepository {
	return &sitesRepositoryImpl{
		db: db,
	}
}
