package repo

type WBCalculation16122022 struct {
	Category                                  string `db:"category"`
	Item                                      string `db:"item"`
	WarehouseWB                               int    `db:"warehouse_wb"`
	SuppliersWarehouseWBToWarehouseWBPercent  int    `db:"suppliers_warehouse_taking_it_to_warehouse_wb_percent"`
	SuppliersWarehouseWBToClientMyselfPercent int    `db:"suppliers_warehouse_taking_it_to_client_myself_percent"`
}
