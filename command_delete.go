package main

func commandDelete(cfg *config, args []string) error {
	return DeleteData(cfg)
}
