package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	_ "expvar"

	_ "net/http/pprof"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func readJsonFromFile() string {
	return `{
  "meta": {
    "code": 20000,
    "type": "OK",
    "message": "OK"
  },
  "data": {
    "trackings": [
      {
        "tracking_number": "j9b8-41ed41",
        "all_tracking_number_types": null,
        "origin_address": {
          "contact_name": "Pandago",
          "company_name": null,
          "street_1": null,
          "street_2": null,
          "street_3": null,
          "city": null,
          "state": null,
          "postal_code": null,
          "country_iso3": null,
          "raw_location": "1 2nd Street #08-01",
          "coordinates": {
            "longitude": 103.8486029,
            "latitude": 1.2923742
          },
          "raw_coordinates": "1.2923742, 103.8486029",
          "type": null
        },
        "destination_address": {
          "contact_name": "Merlion",
          "company_name": null,
          "street_1": null,
          "street_2": null,
          "street_3": null,
          "city": null,
          "state": null,
          "postal_code": null,
          "country_iso3": null,
          "raw_location": "20 Esplanade Drive",
          "coordinates": {
            "longitude": 103.8548608,
            "latitude": 1.2857488
          },
          "raw_coordinates": "1.2857488, 103.8548608",
          "type": null
        },
        "package_count": null,
        "raw_package_count": null,
        "insurance": null,
        "raw_insurance": null,
        "weight": null,
        "raw_weight": null,
        "service_type": null,
        "service_type_name": null,
        "references": null,
        "checkpoints": [
          {
            "message": "Order has been cancelled - Incorrect order details",
            "date_time": {
              "date": "2018-09-13",
              "time": "09:26:40",
              "utc_offset": 8.0
            },
            "raw_date_time": "1536802000",
            "address": null,
            "raw_tag": "CANCELLED",
            "raw_tag_description": null,
            "proof_of_delivery": null,
            "is_stopped": false,
            "hash": "5867d4257360da66975e14e3eac35d4d"
          }
        ],
        "courier_pickup_location": null,
        "pickup_date_time": null,
        "raw_pickup_date_time": null,
        "scheduled_delivery_date_time": null,
        "raw_scheduled_delivery_date_time": null,
        "rescheduled_delivery_date_time": null,
        "raw_rescheduled_delivery_date_time": null,
        "delivered_date_time": null,
        "raw_delivered_date_time": null,
        "signed_by": null,
        "supports_last_mile_tracking": true,
        "next_courier": null,
        "delivery_type": null,
        "proof_of_delivery": "https://pandago-api-sandbox.deliveryhero.io/api/v1/orders/proof_of_delivery/x-1234",
        "slug": "pandago-api",
        "tracking_number_type": null,
        "additional_fields": null,
        "status": {
          "code": 200,
          "subcode": 20000,
          "message": "OK"
        }
      }
    ]
  }
}`
}

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	go func() {
		for {
			log.Println("当前routine数量:", runtime.NumGoroutine())
			time.Sleep(time.Second)
		}
	}()

	router := routing.New()

	router.Get("/", func(c *routing.Context) error {
		fmt.Fprintf(c, "Hello, world!")
		return nil
	})
	router.Any("/pandago", func(c *routing.Context) error {
		fmt.Fprintf(c, readJsonFromFile())
		return nil
	})

	panic(fasthttp.ListenAndServe(":6000", router.HandleRequest))
}
