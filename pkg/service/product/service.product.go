package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alisyahbana/efishery-test/pkg/common/helper"
	"github.com/alisyahbana/efishery-test/pkg/service/product/data"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ProductService struct {
}

func New() ProductService {
	return ProductService{}
}

func (s ProductService) GetAllProducts() ([]data.Product, error) {
	resp, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var listProduct []data.Product

	err = json.NewDecoder(resp.Body).Decode(&listProduct)
	if err != nil {
		return nil, err
	}

	rateUsd, err := helper.GetRatioUSD()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(listProduct); i++ {
		if listProduct[i].Price != nil {
			priceFloat, err := strconv.ParseFloat(*listProduct[i].Price, 64)
			if err != nil {
				log.Println(err, "Failed to convert price to float for uuid:", listProduct[i].ID)
			}
			usdPriceFloat := priceFloat * rateUsd
			listProduct[i].PriceUsd = fmt.Sprintf("$%.2f", usdPriceFloat)
		}
	}

	return listProduct, nil
}

type tempData struct {
	provinsi string
	amount   int
	tahun    string
	minggu   string
}

func (s ProductService) GetCompiledProduct() (res []data.ProductData, err error) {
	listKomoditas, err := s.GetAllProducts()
	if err != nil || listKomoditas == nil {
		log.Println(err)
		err = errors.New("Failed to get commodities list data")
		return res, err
	}

	// Parse the only needed data
	var tempDatas []tempData

	for _, val := range listKomoditas {
		if val.TglParsed == nil || val.Price == nil || val.Size == nil || val.AreaProvinsi == nil {
			continue
		}
		dateTime, err := time.Parse(time.RFC3339, *val.TglParsed)
		if err != nil {
			log.Println(err)
			continue
		}
		year, week := dateTime.ISOWeek()

		price, _ := strconv.Atoi(*val.Price)
		size, _ := strconv.Atoi(*val.Size)
		amount := price * size

		temp := tempData{
			provinsi: *val.AreaProvinsi,
			amount:   amount,
			tahun:    fmt.Sprintf("Tahun %s", strconv.Itoa(year)),
			minggu:   fmt.Sprintf("Minggu ke %s", strconv.Itoa(week)),
		}

		tempDatas = append(tempDatas, temp)
	}

	// Compiling data
	tempMap := make(map[string]map[string]map[string]int)

	for _, val := range tempDatas {
		// Create data set if data provinsi is not exist yet
		if prov, ok := tempMap[val.provinsi]; !ok {
			minggu := map[string]int{val.minggu: val.amount}
			tahun := map[string]map[string]int{val.tahun: minggu}
			tempMap[val.provinsi] = tahun
		} else {
			// Create data set if data provinsi is not exist yet
			if year, ok := prov[val.tahun]; !ok {
				minggu := map[string]int{val.minggu: val.amount}
				prov[val.tahun] = minggu
			} else {
				// Create data set if data profit is not exist yet
				// Add the amount if data profit is exist already
				if profit, ok := year[val.minggu]; !ok {
					year[val.minggu] = val.amount
				} else {
					year[val.minggu] += profit
				}
			}
		}
	}

	var result []data.ProductData

	for key, val := range tempMap {

		productData := data.ProductData{
			Provinsi: key,
			Profit:   val,
			Max:      s.FindMaxProfit(val),
			Min:      s.FindMinProfit(val),
			Avg:      s.FindAvgProfit(val),
			Median:   s.FindMedianProfit(val),
		}

		result = append(result, productData)
	}

	return result, nil
}

func (s ProductService) FindMaxProfit(data map[string]map[string]int) float64 {
	var max int
	for _, val := range data {
		for _, amount := range val {
			if amount >= max {
				max = amount
			}
		}
	}

	return float64(max)
}

func (s ProductService) FindMinProfit(data map[string]map[string]int) float64 {
	min := int(^uint(0) >> 1)
	for _, val := range data {
		for _, amount := range val {
			if amount <= min {
				min = amount
			}
		}
	}

	return float64(min)
}

func (s ProductService) FindAvgProfit(data map[string]map[string]int) float64 {
	var sum, counter int
	for _, val := range data {
		for _, amount := range val {
			sum += amount
			counter++
		}
	}

	return float64(sum / counter)
}

func (s ProductService) FindMedianProfit(data map[string]map[string]int) float64 {
	var arr []int
	for _, val := range data {
		for _, amount := range val {
			arr = append(arr, amount)
		}
	}

	counter := len(arr)

	if counter+1%2 == 0 {
		a := arr[(counter / 2)]
		b := arr[(counter/2)+1]
		return float64((a + b) / 2)
	} else {
		return float64(arr[counter/2])
	}
}
