package main

import (
	"github.com/gubtos/gosubmoduleexample"
	"github.com/gubtos/gosubmoduleexample/integration/storezerolog"
)

func main() {
	logger := storezerolog.Log
	store := gosubmoduleexample.NewStore(logger)

	store.StoreData("my str a", "my str b")
}
