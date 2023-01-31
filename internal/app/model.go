package app

type WBCalculation16122022 struct {
	Category                                  string `csv:"category"`
	Item                                      string `csv:"item"`
	WarehouseWB                               int    `csv:"warehouse_wb"`
	SuppliersWarehouseWBToWarehouseWBPercent  int    `csv:"suppliers_warehouse_taking_it_to_warehouse_wb_percent"`
	SuppliersWarehouseWBToClientMyselfPercent int    `csv:"suppliers_warehouse_taking_it_to_client_myself_percent"`
}
