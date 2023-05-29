# Fishery API

## Function Documentation

### `func GetAllByRange(harga string, size string, tanggal string) ([]model.Commodity, error)`

This function retrieves records from the storage based on the specified range of `harga` (price), `size`, and `tanggal` (date).

- Parameters:
  - `harga` (string): The price range to filter the records.
  - `size` (string): The size range to filter the records.
  - `tanggal` (string): The date range to filter the records.

- Returns:
  - `[]model.Commodity`: A slice of `model.Commodity` objects that match the specified range.
  - `error`: An error if the retrieval process fails.

### `func GetAllByCommodity(komoditas string) ([]model.Commodity, error)`

This function retrieves records from the storage based on the specified `komoditas` (commodity).

- Parameters:
  - `komoditas` (string): The commodity to filter the records.

- Returns:
  - `[]model.Commodity`: A slice of `model.Commodity` objects that match the specified commodity.
  - `error`: An error if the retrieval process fails.

### `func GetById(uuid string) (model.Commodity, error)`

This function retrieves a single record from the storage based on the specified `uuid` (unique identifier).

- Parameters:
  - `uuid` (string): The unique identifier of the record.

- Returns:
  - `model.Commodity`: The `model.Commodity` object that matches the specified `uuid`.
  - `error`: An error if the retrieval process fails.

### `func GetByArea(areaProvinsi string, areaKota string) ([]model.Commodity, error)`

This function retrieves records from the storage based on the specified `areaProvinsi` (province) and `areaKota` (city).

- Parameters:
  - `areaProvinsi` (string): The province to filter the records.
  - `areaKota` (string): The city to filter the records.

- Returns:
  - `[]model.Commodity`: A slice of `model.Commodity` objects that match the specified province and city.
  - `error`: An error if the retrieval process fails.

### `func AddRecords(records []model.Commodity) error`

This function adds new records to the storage.

- Parameters:
  - `records` ([]model.Commodity): A slice of `model.Commodity` objects to be added.

- Returns:
  - `error`: An error if the addition process fails.

### `func UpdateRecords(uuid string, updatedRec model.Commodity) error`

This function updates an existing record in the storage based on the specified `uuid` (unique identifier).

- Parameters:
  - `uuid` (string): The unique identifier of the record to be updated.
  - `updatedRec` (model.Commodity): The updated `model.Commodity` object.

- Returns:
  - `error`: An error if the update process fails.

### `func DeleteRecords(uuid string) error`

This function deletes a record from the storage based on the specified `uuid` (unique identifier).

- Parameters:
  - `uuid` (string): The unique identifier of the record to be deleted.

- Returns:
  - `error`: An error if the deletion process fails.

### `func GetMaxPrice(week int, commodity string) ([]model.Commodity, error)`

This function retrieves records from the storage based on the specified `week` and `commodity` to get the maximum price.

- Parameters:
  - `week` (int): The week number to filter the records.
  - `commodity` (string): The commodity to filter the records.

- Returns:
  - `[]model.Commodity`: A slice of `model.Commodity` objects that match the specified week and commodity, containing the records with the maximum price.
  - `error`: An error if the retrieval process fails.

### `func GetMostRecords(commodity string) ([]model.Commodity, error)`

This function retrieves records from the storage based on the specified `commodity` to get the most records.

- Parameters:
  - `commodity` (string): The commodity to filter the records.

- Returns:
  - `[]model.Commodity`: A slice of `model.Commodity` objects that match the specified commodity, containing the records with the most occurrences.
  - `error`: An error if the retrieval process fails.
