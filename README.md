# Fishery API CLI

This repository contains a command-line interface (CLI) tool for interacting with the Fishery API.

## Prerequisites

- Go (version X.X.X)
- Git

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/fidellr/fishery_api.git
   ```

2. Change to the project directory:

   ```shell
   cd fishery_api/examples
   ```

3. Build the CLI binary:

   ```shell
   go build -o fishery-cli
   ```

## Usage

The Fishery API CLI allows you to perform various operations on the fishery data. Here are the available commands:

### Create New Records
To create new records of commodity
```shell
./fishery-cli addCommodityRecords '[{"komoditas": "bandeng", "area_provinsi": "jawa barat", "area_kota": "cirebon", "size": "30", "price": "20000" }]'
```

### Get All By Commodity
To get all existing commodity records by commodity's name
```shell
./fishery-cli getAllByCommodity 'bandeng'
```

### Get Commodity By ID
To get all existing commodity records by its ID/UUID
```shell
./fishery-cli getCommodityByID '2d7be273-e36d-479e-a0a2-d45d73941090'
```

### Update Records
To update records more than one at once
```shell
./fishery-cli updateCommodityRecords '[{ "uuid": "11da9b73-3765-4c19-a30d-d234e3c8abf0", "komoditas": "GURAME", "area_provinsi": "JAWA BARAT", "area_kota": "DEPOK", "size": "500", "price": "87000", "tgl_parsed": "2022-01-03T15:40:09Z", "timestamp": "1641224409052" }, { "uuid": "eb3b2547-1b7c-4304-8281-dfc7d0b8991e", "komoditas": "BANDENG", "area_provinsi": "JAWA BARAT", "area_kota": "CILILIN", "size": "13", "price": "8400", "tgl_parsed": "2022-01-03T16:13:30Z", "timestamp": "1641226410400" }]'
```

### Delete Records
To delete existing record more than one at once

```shell
./fishery-cli deleteCommodityRecords '["11da9b73-3765-4c19-a30d-d234e3c8abf0","eb3b2547-1b7c-4304-8281-dfc7d0b8991e"]'
```

### Get All Records by Range

To retrieve records based on a specific price range, size range, and date range

```shell
./fishery-cli getAllByRange '{"price": "48000", "size": "120", "date": "2022-01-04"}'
```

### Get Records by Commodity Name and Area

To retrieve records based on a specific province and city

```shell
./fishery-cli getAllByCommodityAndArea '{"komoditas": "bandeng", "area_provinsi": "JAWA BARAT", "area_kota": "DEPOK"}'
```

### Get Records by Price Range
To retrieve records based on a specific price range

```shell
./fishery-cli getAllByPriceRange '{"price": "50000-100000"}'
```

### Get Latest 10 Records
To retrieve latest 10 existing records
```shell
./fishery-cli getLatestTenCommodities
```
