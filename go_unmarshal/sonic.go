package main

import (
	"encoding/json"
	"fmt"

	"github.com/bytedance/sonic"
)

type Rate struct {
	DetailedCharges []struct {
		Charge struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		} `json:"charge"`
		Type string `json:"type"`
	} `json:"detailed_charges"`
}

func main() {
	// Case 1: detailed_charges 为 null 的场景
	// jsonNull := `{
	// 	"detailed_charges": null
	// }`

	// // Case 2: detailed_charges 为数组的场景
	// jsonArray := `{
	// 	"detailed_charges": [{
	// 		"charge": {
	// 			"amount": 298.62,
	// 			"currency": "USD"
	// 		},
	// 		"type": "base"
	// 	}]
	// }`

	// case 3: 空数组的场景
	jsonStr := `{
  "meta" : {
    "code" : 20100,
    "type" : "Created",
    "message" : "The request has been fulfilled and a new resource has been created."
  },
  "data" : {
    "list" : [ {
      "id" : "ada56536add54ad78a1a590494c2aaa9",
      "tracking_numbers" : [ "863T30049491A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T14:05:39.374175Z",
      "updated_at" : "2025-03-12T14:05:41.630953Z",
      "succeed_at" : "2025-03-12T14:05:41.62347794Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : "123",
      "order_number" : "12345",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/ada56536-add5-4ad7-8a1a-590494c2aaa9-1741788340815.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : [ {
        "insured_value" : {
          "amount" : 27,
          "currency" : "USD"
        },
        "type" : "insurance"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "00561c755278452ebd35d8c29385084f",
      "tracking_numbers" : [ "863T30049490A002" ],
      "organization_id" : "3fe7752813cb46b7baa82d0ba9ac5c96",
      "created_at" : "2025-03-12T14:01:23.473359Z",
      "updated_at" : "2025-03-12T14:01:26.181725Z",
      "succeed_at" : "2025-03-12T14:01:26.174310814Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "22303"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "94209"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "Alexandria",
        "company_name" : "weqwe",
        "contact_name" : "Jade 222",
        "country" : "USA",
        "email" : "yj.tang@aftership.com",
        "postal_code" : "22303",
        "state" : "VA",
        "street1" : "2000 Huntington Avenue"
      },
      "ship_to" : {
        "city" : "Sacramento",
        "company_name" : "21233",
        "contact_name" : "Jade",
        "country" : "USA",
        "phone" : "1-140-225-3341",
        "postal_code" : "94209",
        "state" : "CA",
        "street1" : "28292 Daugherty Orchard"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "9087f22e89024213adf77e88cec1a8a6",
      "order_number" : "#1021",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/00561c75-5278-452e-bd35-d8c29385084f-1741788085586.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "qweqwewq9HCULBYZ", "wwring x 1", "eeee$27.00 USD" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "de3b9720dd86405b993eee6e8fd6f433",
      "service_type" : "yodel_returns",
      "rate" : {
        "booking_cut_off" : null,
        "extra_info" : null,
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account New2",
          "id" : "de3b9720dd86405b993eee6e8fd6f433",
          "slug" : "yodel"
        }
      },
      "service_options" : [ {
        "insured_value" : {
          "amount" : 27,
          "currency" : "USD"
        },
        "type" : "insurance"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "internal_custom_fields" : {
          "allocation_id" : "045c2fbd418a4c9c9d8a96bf7b50d288",
          "notify_customer" : "false",
          "operator_account_id" : "663c5c3139644f1481cbe9bfe5bf911c",
          "returns_shipping_type" : "retailer_label",
          "tracking_slug" : "yodel"
        },
        "shipper_account" : {
          "custom_fields" : {
            "label_qrcode_enabled" : true,
            "shipment_label_enabled" : false
          }
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 0.79475
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 2,
        "height" : 2,
        "unit" : "ft",
        "width" : 2
      } ],
      "items" : [ {
        "barcode" : "972019892255497201989225549720198922554",
        "description" : "ring",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0903/5956/2554/files/IMG_3935.jpg?v=1730181618" ],
        "item_id" : "15457596735802",
        "price" : {
          "amount" : 27,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "9720198922554",
        "weight" : {
          "unit" : "lb",
          "value" : 0.16975
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "a6ac321b6ee94634a906c90a9102dc5b",
      "tracking_numbers" : [ "863T30049489A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T13:57:57.920767Z",
      "updated_at" : "2025-03-12T13:57:59.780851Z",
      "succeed_at" : "2025-03-12T13:57:59.76848278Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/a6ac321b-6ee9-4634-a906-c90a9102dc5b-1741787879002.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "ca52eebe01eb45b58f77e5085d90ea0b",
      "tracking_numbers" : [ "863T30049488A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T13:56:56.955007Z",
      "updated_at" : "2025-03-12T13:57:03.32535Z",
      "succeed_at" : "2025-03-12T13:57:03.317261193Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/ca52eebe-01eb-45b5-8f77-e5085d90ea0b-1741787822386.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : [ {
        "insured_value" : {
          "amount" : 27,
          "currency" : "USD"
        },
        "type" : "insurance"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "1477f0f3fee9416fba7e1dfcc202e4d3",
      "tracking_numbers" : [ "863T30049487A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T13:07:56.194975Z",
      "updated_at" : "2025-03-12T13:08:02.608537Z",
      "succeed_at" : "2025-03-12T13:08:02.599405625Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/1477f0f3-fee9-416f-ba7e-1dfcc202e4d3-1741784881548.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "b81876ea69864182aa5b595e9ebaf767",
      "tracking_numbers" : [ "863T30049486A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T12:38:52.348591Z",
      "updated_at" : "2025-03-12T12:38:54.309253Z",
      "succeed_at" : "2025-03-12T12:38:54.30148763Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : "123",
      "order_number" : "12345",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/b81876ea-6986-4182-aa5b-595e9ebaf767-1741783133557.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : [ {
        "insured_value" : {
          "amount" : 27,
          "currency" : "USD"
        },
        "type" : "insurance"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "c07159e09d28481eac0cc5aca8bb87b9",
      "tracking_numbers" : [ "863T30049485A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T12:37:39.716255Z",
      "updated_at" : "2025-03-12T12:37:41.796858Z",
      "succeed_at" : "2025-03-12T12:37:41.788391399Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : "12345",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/c07159e0-9d28-481e-ac0c-c5aca8bb87b9-1741783061026.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : [ {
        "insured_value" : {
          "amount" : 27,
          "currency" : "USD"
        },
        "type" : "insurance"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "d90763ff587b4d0fb2c2713807982f3f",
      "tracking_numbers" : [ "863T30049484A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T12:37:21.935295Z",
      "updated_at" : "2025-03-12T12:37:23.868758Z",
      "succeed_at" : "2025-03-12T12:37:23.86105179Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/d90763ff-587b-4d0f-b2c2-713807982f3f-1741783043080.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : [ {
        "insured_value" : {
          "amount" : 27,
          "currency" : "USD"
        },
        "type" : "insurance"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "d3cb826181ac4be199279f5dedb2283b",
      "tracking_numbers" : [ "863T30049483A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T12:36:07.542575Z",
      "updated_at" : "2025-03-12T12:36:14.043408Z",
      "succeed_at" : "2025-03-12T12:36:14.036241758Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/d3cb8261-81ac-4be1-9927-9f5dedb2283b-1741782973246.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : [ {
        "insured_value" : {
          "amount" : 27,
          "currency" : "USD"
        },
        "type" : "insurance"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "7faaed366712464a84aa6bc5b95ee090",
      "tracking_numbers" : [ "863T30049482A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T12:22:42.074015Z",
      "updated_at" : "2025-03-12T12:22:48.560457Z",
      "succeed_at" : "2025-03-12T12:22:48.552829679Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/7faaed36-6712-464a-84aa-6bc5b95ee090-1741782167757.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "7594237c2dfb41088a835addbe242de4",
      "tracking_numbers" : [ "863T30049481A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T11:58:17.382175Z",
      "updated_at" : "2025-03-12T11:58:23.766232Z",
      "succeed_at" : "2025-03-12T11:58:23.754918071Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/7594237c-2dfb-4108-8a83-5addbe242de4-1741780702938.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "e0c341a9173f4f1fbe904fe2d8d1beca",
      "tracking_numbers" : [ "863T30049480A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T11:45:02.197679Z",
      "updated_at" : "2025-03-12T11:45:04.277314Z",
      "succeed_at" : "2025-03-12T11:45:04.267989294Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/e0c341a9-173f-4f1f-be90-4fe2d8d1beca-1741779903486.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : [ {
        "insured_value" : {
          "amount" : 27,
          "currency" : "USD"
        },
        "type" : "insurance"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "208e889ea1ea4a438fefb8781775a196",
      "tracking_numbers" : [ "863T30049479A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T11:42:32.009567Z",
      "updated_at" : "2025-03-12T11:42:38.743298Z",
      "succeed_at" : "2025-03-12T11:42:38.734425323Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/208e889e-a1ea-4a43-8fef-b8781775a196-1741779758025.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "f999a62fecca4eb79df82141d1567b46",
      "tracking_numbers" : [ "794684898952" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T11:32:28.934463Z",
      "updated_at" : "2025-03-12T11:32:35.732444Z",
      "succeed_at" : "2025-03-12T11:32:35.715702313Z",
      "meta" : {
        "source" : "tts"
      },
      "source" : "tts",
      "ship_from" : {
        "city" : "Centerville",
        "company_name" : "Postmen",
        "contact_name" : "Contact Name",
        "country" : "USA",
        "email" : "usps_discounted@test.com",
        "phone" : "18682306030",
        "postal_code" : "45458",
        "state" : "OH",
        "street1" : "8700 Meadowcreek Dr",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Santa Ana",
        "company_name" : "Maersk",
        "contact_name" : "Mister Jones",
        "country" : "USA",
        "email" : "maersk.test@example.com",
        "phone" : "8777285325",
        "postal_code" : "92705",
        "state" : "CA",
        "street1" : "1600 E. Saint Andrew Pl.",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : "202503120007",
      "return_shipment" : false,
      "files" : {
        "customs_declaration" : null,
        "invoice" : null,
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6"
        },
        "manifest" : null,
        "packing_slip" : null,
        "qr_code" : null
      },
      "box_type" : "custom",
      "references" : [ "order950000002" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "aa45a382-919f-482f-aa81-374f7f1a743a",
      "service_type" : "fedex_ground_economy",
      "rate" : {
        "booking_cut_off" : null,
        "charge_weight" : {
          "unit" : "lb",
          "value" : 3
        },
        "delivery_date" : "2025-03-19",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 18.05,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 1.1,
            "currency" : "USD"
          },
          "type" : "delivery_confirmation"
        }, {
          "charge" : {
            "amount" : 1.15,
            "currency" : "USD"
          },
          "type" : "fuel"
        }, {
          "charge" : {
            "amount" : 1.9,
            "currency" : "USD"
          },
          "type" : "other"
        } ],
        "error_message" : null,
        "info_message" : null,
        "pickup_deadline" : null,
        "service_name" : "FedEx Ground® Economy",
        "service_type" : "fedex_ground_economy",
        "shipper_account" : {
          "description" : "[FedEx] Testing account include fedex ground economy",
          "id" : "aa45a382-919f-482f-aa81-374f7f1a743a",
          "slug" : "fedex"
        },
        "total_charge" : {
          "amount" : 22.2,
          "currency" : "USD"
        },
        "transit_time" : null
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.2
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 3
      },
      "dimension" : [ {
        "depth" : 4,
        "height" : 3,
        "unit" : "in",
        "width" : 7
      } ],
      "items" : [ {
        "description" : "F231",
        "item_id" : "1313131223",
        "price" : {
          "amount" : 50,
          "currency" : "USD"
        },
        "quantity" : 3,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "method" : {
          "account_number" : "123456",
          "country" : "USA",
          "postal_code" : "123456",
          "type" : "account"
        },
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : "JVBERi0xLjQKMSAwIG9iago8PAovVHlwZSAvQ2F0YWxvZwovUGFnZXMgMyAwIFIKPj4KZW5kb2JqCjIgMCBvYmoKPDwKL1R5cGUgL091dGxpbmVzCi9Db3VudCAwCj4+CmVuZG9iagozIDAgb2JqCjw8Ci9UeXBlIC9QYWdlcwovQ291bnQgMQovS2lkcyBbMTggMCBSXQo+PgplbmRvYmoKNCAwIG9iagpbL1BERiAvVGV4dCAvSW1hZ2VCIC9JbWFnZUMgL0ltYWdlSV0KZW5kb2JqCjUgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9IZWx2ZXRpY2EtQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjcgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhLU9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iago4IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0hlbHZldGljYS1Cb2xkT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjkgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvQ291cmllcgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEwIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjExIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEyIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZE9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxMyAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Sb21hbgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjE0IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL1RpbWVzLUJvbGQKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNSAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1JdGFsaWMKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Cb2xkSXRhbGljCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKMTcgMCBvYmogCjw8Ci9DcmVhdGlvbkRhdGUgKEQ6MjAwMykKL1Byb2R1Y2VyIChGZWRFeCBTZXJ2aWNlcykKL1RpdGxlIChGZWRFeCBTaGlwcGluZyBMYWJlbCkNL0NyZWF0b3IgKEZlZEV4IEN1c3RvbWVyIEF1dG9tYXRpb24pDS9BdXRob3IgKENMUyBWZXJzaW9uIDUxMjAxMzUpCj4+CmVuZG9iagoxOCAwIG9iago8PAovVHlwZSAvUGFnZQ0vUGFyZW50IDMgMCBSCi9SZXNvdXJjZXMgPDwgL1Byb2NTZXQgNCAwIFIgCiAvRm9udCA8PCAvRjEgNSAwIFIgCi9GMiA2IDAgUiAKL0YzIDcgMCBSIAovRjQgOCAwIFIgCi9GNSA5IDAgUiAKL0Y2IDEwIDAgUiAKL0Y3IDExIDAgUiAKL0Y4IDEyIDAgUiAKL0Y5IDEzIDAgUiAKL0YxMCAxNCAwIFIgCi9GMTEgMTUgMCBSIAovRjEyIDE2IDAgUiAKID4+Ci9YT2JqZWN0IDw8IC9GZWRFeEdyb3VuZCAyMCAwIFIKPj4KPj4KL01lZGlhQm94IFswIDAgMjg4IDQzMl0KL1RyaW1Cb3hbMCAwIDI4OCA0MzJdCi9Db250ZW50cyAxOSAwIFIKL1JvdGF0ZSAwPj4KZW5kb2JqCjE5IDAgb2JqCjw8IC9MZW5ndGggMzE5MQovRmlsdGVyIFsvQVNDSUk4NURlY29kZSAvRmxhdGVEZWNvZGVdIAo+PgpzdHJlYW0KR2F0PS8/I050LCVYZkVBcyQ4bldnN1l1LDs1LThKYyo9MUtdNWQxYC1AK24iXm4qLWpuSFtZODJOLlAsOTFlL2A0Oz05WSIoVkk/VD9BSUEKJ0xoN1sqNWkvOFsyLDwnb0ZbJFhFTypVND9aYnBLbWFuZVNaVEtCQmE5RzxlcU5yYEBsKmg1SSJQSHNrUEkxQ1hJPSJuPSQnTm9UakFPRkAKPDZEXSxIK1RVRm0zYVUwL0g3XydCVThbZikvIVhCb1o+Pl1WdXBEYGJmb2ZHU1kpTU9ySjBMTEIoTWtJTDtRTSNMTmczPFgsQDVTSVdfdUcKT1haYkA5cUMqZ0leaVA8XjtQbkk6MjNmYjNRUl9sPS1uLDwwKC4/VzpaZ20/QUpYKnNgPj80WltwRDcnaGBlLjFzMUdxVz9ONFw2SWtecysKayY5IXJoRzBuTjcxVGtbLSxBLSs+Rm4iZFMyNF9QJy51RDkzKkVCUDVDXCFJJ3RdSiphamskbVNiX3NIRzVfNSpbWDxwbVtUb1dUYnBUaSIKPy83XTNLQTRUSG9iVlttNGhdJ1doQnBoPmY/PHUhQyMwbENTXEh0IWJnRUQ4SiVASXA/XU5IPV5PRFU8UFRPYHE8ZzxWQGYlI2lEbkEjYC8KKDgmWHVmdEEkPEtyckgtTW1KKnRSX0xyYDMnVD5RY3BOL0FIS2ZLWT8lYGdtTkIyUE43L1BcaVVeUFNhRyRXVVFCT0JMaS40VStPL3VMP1kKbjJVNF5OWTVlTSkrdWVfV3FVbjMyRW9eRmBGRDtVWnM4SDpXWSlXYityNWleYG5UbSQ/IzBmSDRlaWBPcT5PLVojUXUzR2d0cVhiXDd0YFgKVVJxbFZKSUlWUlFCXVY0JTVHJT5rUHRNVnJwRSJbUWZkTExwPWw/MF4pZnVoYWRZPTphRFBgJlUkUW5bZkxTXF1dQ15SVkY7QkgxUV5paisKYHFEVllNJG5yTTdkLTZKZEYnXmNjQE5GJEQ1RCY5MXM+ZEhgTV1jazVLclpGR0dhKUJxVWY6bCtBQUpoKjVHR2Q7Rzc5dSNUQ05cZiQqVT8KRGdsQDFJZkgrVURvWzVaZVwlX0NSNVJlWkYlNWNfZUkzaW5SZDRnQGltN3EhOUNPTmpTc2tJZy4/YjlbTC1WKmMwKEFfZkQiMjkmRyttOV0KZFppSkc+IVQwZSUkW1xNOm8qWUZrOixZOiQ1PSJBUSY+LlNbL2orM2FxOEorUW9JaSQtVnNRbU1jZktUMGdwa1NpQydFZWY6ZiVKTjclSU4KLTpBNkI0Z1pzPTcraVFWMEZgUVQuPjlhbGswZmJsTT8wUCZFTF1RVjotQyojXEpfJUtbZzMtXywpMCQuIjBKRitdYWM9JyhhVUFkO248SiEKPCpJLDNVLzdBRUZMcnJ0Z1MxNjopZUhmX01qS10sQC9zOVAtXF9xczQhUmE7XnUtSWJMPTZkX01Kays9NTI1LzosL2VJblUnLVteNSc9aTEKakNyWklMY0EtSSxzImYvVV1sWzFGdTl0aihmJExTV1V1dGlFIm0hbz82V2k8amMhbTklK1YrMCNgPFUmLkQ3YDo9NmFob11gO0ZJZThlI20KUWYnamJLU2xDLG8qQlVUMU1uck0mXVxVQmlvS0ZjZzttKUM9ZkotOVM8ZXQ/Z3U2Y00sXVVATVIkLEhkOlktJ1leN21jRixTMUU4QzkkJ3EKaFlXRHVQYSYiXChdXFxtazhcVmVGVGBiImQ4UVIzRjtQImNQRzUzPWZBKjM/UWA9L01ucSNoY2VddF1cQDVFdVByQnBJaGFNYFU0LlxJUWsKS3Q5TkMtbCRcKmUlMVFjX2NaLiolNjdAPixOXSg1QzFsVmxvRlRkKzZuWTFcOGxRZS9WWGhEWyM+YU0wPio+V1czKUglKV5AXktnNFtCQiwKTWNuPlZnLkBFND9XVWMvIU42NmNGK09MWlA3Ijlqa2A/MHNnJ2hMPSMxUUQqTmhyajQ1aF5wP3AwbystISw+RlFFRCJJV1NeI28rWCRGI28KO3EkdFtBayphYWBFQ3QjIz0na0BvS0BUcEF0IlU5a1c9TXQnQEdmWFwpPkBUQylFL2FXc2pNXldhcEk5YFoha1dgSyJzPTwpWzNpZ1g1SnEKJFwtIiJNYzE6U0lfY1YlaCRzZExnQkltdFdBPSw0Tkg/IignRyRaUVZePiNhNidMYWhcNi1cNmknOGtvbC44Wmc3TzJhYERNYTVRZVByNyoKRGJTWTVLIlEiVVM9MWolVj43a0s1ZlAoSCI6JEVVVzUzQCQlczoibTk7IzRURkE5XSlDOSNxWFFdQ3FTTVxoSEAiLyh0JmZtXztDNGNWRj0KZz5zUzAlVEhBOEoiRGBHNiJPOSNrUUBLSC5lYT9II3NzTWkubSFgR1dNamxTZiw2JCgoa0M9Zk0yRERjPClgcWFCUEZTWlJ0SjZGJVkuXzQKRGdNdCxFWD1sRlMnMW4wOWJtTjBPNlRFOWhUKEllXzFDPTA/Q3JWJiMicyl1JEEkaUAlWTM4OUwhQCJyOjxxREc6MGJYVVcxP3Iwck8rXFAKSEIlN1BGaUFwYz1lWFhXLCprdHQrSikiXjozaDBoKSNLOjJvdG0sR09OS1cpQnE3Y01AbSsjTDY7O2c/OStkIyhFRV5zX2tRZWo4RVpLXlEKJCxVWytkNzErL208UD8uQEJYSS5XISxlOSIwbCZWQlE4ITNwMjlrN21gRz40VGxEU11TXjNXRWNKUCwxcjs+VUZpUzlzTEosOnAsL0VdTUIKWmBeP1EyWTxHLlltTCpFQDloPnU5V0ReaW9VaCVjVUhxYEltJXRDWkk0SmBLSiRIXkdNdFFjPmlvL1cqVTorYzc1blpUWz5oPk1eUEdxbzAKPVQsS19CPC8pVjRmaExIbXI5dDJnNT4jNz9fNUdYPm5ZcjFRSCpXOFlEU0N1LlZuPy8qdUI6ZkolPjNxby5lWzI/JG5PajUxP1gjazZIZUIKMy5DbjZxOUZBIUVjMXVdSVIuU00wWHNyZ1MoRlVoaG84PWxTUEVZSj0rMikkJTksQkZCNmBqSmxAKjU1X0E7RDpLYUEzLTklTzZ1RGxtaU0KclJdV29EYi5mVDkjaCwlUU5sM0g9NVtMQG08aVUlRCw2ZFdxQ3JUJk0rbHQlWSRYNSpvOydxdGRyOS5EMUNcOnFoZFBfXzlAXHNHTHJnJT0KaiJUU1IjLXVzRUJWYSxBVTtVV04wYklMT1BXVlc+KXRTJDI6KUJWYl00MSwhUSwnXj47THVgTSZQQj0vJVRKc2wyTzZxbUA1K1A4MS5BXToKIkAuOWVNVF9CQl9tRFA/JGNeW0pdNS9lJjBwSXRNOlhJVD1VZGwjUzp1YU4vKTReaUIpRkxNREJULU5IX2RVLiRUcW9xN0JTdCFfIUQ4OSoKX2ojJVtlRixcJTFoX1VlZCxBRy0xcihRLE8vIiYqa21aRUdWY2dZR01ZNyVKU2w9bydsYDVYJ01tVkBORSwpNypqWWtdSj5gKHNSV3Jua0UKKkx1YSIjMF5SJltzMlZsJVdLU2dHJHFyJzVqUFktTGYjUnUiLlloJUdKVnAxMltkbyJCaGZZR0FuI01EXnI1PXAtJ1I0Q3FiTG9dQiZZZC0KJ0Q/ISxuJ1ZESksyPlZTMyZNQXVkS0NtTGM3LDFGWE9manA6Xy9eZzMhOkUwO3REW2Q5ZCZIPFM8QV5JX2NFR1hkXiFzUThsT18vMjpJU0kKLUxnQiVFYHM2VVUnV3FlUTtdT24nbD9YXzc2KXExZSdKNEEwJDVXMlw6Y2JQKUpGLnBWSGpMJitDXjR1UUE1Y1shWUpNcFBZPWI5KVxbM1MKOjxVUTkxUChBSkVDcVJcU3JLLUJFaCZHMD0kNzRRUWwpYV9ARzcsN2A1T1xWZG0scjU/TmduRzMlWWVCYzMsSUhgQDc0YkUmTTtILXVpRTsKVk9JbGM+X3BpYTEiZjYuaV08VG1NNXQ1dVZMXS9SQVJfQyVecjA1NCxubm5TYF82SUg6PFVSJDFQKEJ1RTRTQ1ovXTMlSWJCLHEpY0FDa0QKajBXazcpZWlnUnJJL0khKVc3dD1kcCtPPlIoKi1vUy5qQXBLWSlDU15GPnFyaj4pOkFXR2JEVDBzc141QEgqZClEYk1dJG81TyhraXFaNXEKYEdZO186TTZpc2QiSzViSUc0WVBMYi9RMmZDQTY7fj4KZW5kc3RyZWFtCmVuZG9iagoyMCAwIG9iago8PCAvVHlwZSAvWE9iamVjdAovU3VidHlwZSAvSW1hZ2UKL1dpZHRoIDExOAovSGVpZ2h0IDQ3Ci9Db2xvclNwYWNlIC9EZXZpY2VHcmF5Ci9CaXRzUGVyQ29tcG9uZW50IDgKL0xlbmd0aCA0NTAKL0ZpbHRlciBbL0FTQ0lJODVEZWNvZGUgL0ZsYXRlRGVjb2RlXQo+PnN0cmVhbQpHYiIvZUpJXVI/I1huTGdUNjxkZi9JLiY/Qj1oY0RhZy1yRDRdMCVQZk1UWitgUVNJK0pdQ14mckQzQThsTThubkZGUCFJLmYqallDMlRuWwo8N2koN1w8SkghPHVWdW0zSEpmazw4KSslWk4qTWYtcyVsPlZaVUNzLDlgOW9XJlRNJS9jZ21KVSUuZkxoYkkkJz5CZ3VfaUBRK0MmYldyNgozRENNNi9GVU10LERCXiwqUjAwcyJdV0ZHQzBqOiNXYTRnVzwqUWVUOyw/O0JTIW9ydTV1JF8zQjJqXDJDLidiLkZYcUBYMUEqLGFcU2FHSgo6bk1rY0lRXC5KZjowNTdfLUw0SjwlMlFNbUlhb0BcODhGS0diMXAsbVc7JS5ITkMuO1YuK2BkKXM6ckJoVWpMSEJYMlJDXjlfcTgxblZbSwpMPTJRQkZ0O3JUQ242W2ZwbXByN1NfY0chZFdbLHNYMSsrOFI8KS0lRmNVLmtScFdTKVhFOzkrOzRsUytfNkFzcmNFYlFMY0lFY187JU1LRQozY2pHVDkqTltKLnBOMldtSThXPyRsO2M2XTVkLz9wOHQ8c0VhSjJlcmwrfj4KZW5kc3RyZWFtCmVuZG9iagp4cmVmCjAgMjEKMDAwMDAwMDAwMCA2NTUzNSBmIAowMDAwMDAwMDA5IDAwMDAwIG4gCjAwMDAwMDAwNTggMDAwMDAgbiAKMDAwMDAwMDEwNCAwMDAwMCBuIAowMDAwMDAwMTYyIDAwMDAwIG4gCjAwMDAwMDAyMTQgMDAwMDAgbiAKMDAwMDAwMDMxMiAwMDAwMCBuIAowMDAwMDAwNDE1IDAwMDAwIG4gCjAwMDAwMDA1MjEgMDAwMDAgbiAKMDAwMDAwMDYzMSAwMDAwMCBuIAowMDAwMDAwNzI3IDAwMDAwIG4gCjAwMDAwMDA4MjkgMDAwMDAgbiAKMDAwMDAwMDkzNCAwMDAwMCBuIAowMDAwMDAxMDQzIDAwMDAwIG4gCjAwMDAwMDExNDQgMDAwMDAgbiAKMDAwMDAwMTI0NCAwMDAwMCBuIAowMDAwMDAxMzQ2IDAwMDAwIG4gCjAwMDAwMDE0NTIgMDAwMDAgbiAKMDAwMDAwMTYyMiAwMDAwMCBuIAowMDAwMDAxOTY2IDAwMDAwIG4gCjAwMDAwMDUyNDkgMDAwMDAgbiAKdHJhaWxlcgo8PAovSW5mbyAxNyAwIFIKL1NpemUgMjEKL1Jvb3QgMSAwIFIKPj4Kc3RhcnR4cmVmCjU4ODUKJSVFT0YK"
    }, {
      "id" : "8240c4ad803745e490aa9617477d422b",
      "tracking_numbers" : [ "794684898687" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T11:27:01.029343Z",
      "updated_at" : "2025-03-12T11:27:08.071075Z",
      "succeed_at" : "2025-03-12T11:27:08.03802035Z",
      "meta" : {
        "source" : "tts"
      },
      "source" : "tts",
      "ship_from" : {
        "city" : "Centerville",
        "company_name" : "Postmen",
        "contact_name" : "Contact Name",
        "country" : "USA",
        "email" : "usps_discounted@test.com",
        "phone" : "18682306030",
        "postal_code" : "45458",
        "state" : "OH",
        "street1" : "8700 Meadowcreek Dr",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Santa Ana",
        "company_name" : "Maersk",
        "contact_name" : "Mister Jones",
        "country" : "USA",
        "email" : "maersk.test@example.com",
        "phone" : "8777285325",
        "postal_code" : "92705",
        "state" : "CA",
        "street1" : "1600 E. Saint Andrew Pl.",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : "202503120006",
      "return_shipment" : false,
      "files" : {
        "customs_declaration" : null,
        "invoice" : null,
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6"
        },
        "manifest" : null,
        "packing_slip" : null,
        "qr_code" : null
      },
      "box_type" : "custom",
      "references" : [ "order950000002" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "15eebae3-1646-4349-88f3-e8d0ec00e577",
      "service_type" : "fedex_ground_home_delivery",
      "rate" : {
        "booking_cut_off" : null,
        "charge_weight" : {
          "unit" : "lb",
          "value" : 172.7
        },
        "delivery_date" : "2025-03-18",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 201.07,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 5.95,
            "currency" : "USD"
          },
          "type" : "residential_delivery"
        }, {
          "charge" : {
            "amount" : 305,
            "currency" : "USD"
          },
          "type" : "oversize"
        }, {
          "charge" : {
            "amount" : 28.16,
            "currency" : "USD"
          },
          "type" : "fuel"
        } ],
        "error_message" : null,
        "info_message" : null,
        "pickup_deadline" : null,
        "service_name" : "FedEx Home Delivery®",
        "service_type" : "fedex_ground_home_delivery",
        "shipper_account" : {
          "description" : "fedex compatible opp test",
          "id" : "15eebae3-1646-4349-88f3-e8d0ec00e577",
          "slug" : "fedex"
        },
        "total_charge" : {
          "amount" : 540.18,
          "currency" : "USD"
        },
        "transit_time" : null
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.2
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 172.7
      },
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "in",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "F231",
        "item_id" : "1313131223",
        "price" : {
          "amount" : 50,
          "currency" : "USD"
        },
        "quantity" : 3,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "method" : {
          "account_number" : "123456",
          "country" : "USA",
          "postal_code" : "123456",
          "type" : "account"
        },
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : "JVBERi0xLjQKMSAwIG9iago8PAovVHlwZSAvQ2F0YWxvZwovUGFnZXMgMyAwIFIKPj4KZW5kb2JqCjIgMCBvYmoKPDwKL1R5cGUgL091dGxpbmVzCi9Db3VudCAwCj4+CmVuZG9iagozIDAgb2JqCjw8Ci9UeXBlIC9QYWdlcwovQ291bnQgMQovS2lkcyBbMTggMCBSXQo+PgplbmRvYmoKNCAwIG9iagpbL1BERiAvVGV4dCAvSW1hZ2VCIC9JbWFnZUMgL0ltYWdlSV0KZW5kb2JqCjUgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9IZWx2ZXRpY2EtQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjcgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhLU9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iago4IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0hlbHZldGljYS1Cb2xkT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjkgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvQ291cmllcgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEwIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjExIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEyIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZE9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxMyAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Sb21hbgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjE0IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL1RpbWVzLUJvbGQKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNSAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1JdGFsaWMKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Cb2xkSXRhbGljCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKMTcgMCBvYmogCjw8Ci9DcmVhdGlvbkRhdGUgKEQ6MjAwMykKL1Byb2R1Y2VyIChGZWRFeCBTZXJ2aWNlcykKL1RpdGxlIChGZWRFeCBTaGlwcGluZyBMYWJlbCkNL0NyZWF0b3IgKEZlZEV4IEN1c3RvbWVyIEF1dG9tYXRpb24pDS9BdXRob3IgKENMUyBWZXJzaW9uIDUxMjAxMzUpCj4+CmVuZG9iagoxOCAwIG9iago8PAovVHlwZSAvUGFnZQ0vUGFyZW50IDMgMCBSCi9SZXNvdXJjZXMgPDwgL1Byb2NTZXQgNCAwIFIgCiAvRm9udCA8PCAvRjEgNSAwIFIgCi9GMiA2IDAgUiAKL0YzIDcgMCBSIAovRjQgOCAwIFIgCi9GNSA5IDAgUiAKL0Y2IDEwIDAgUiAKL0Y3IDExIDAgUiAKL0Y4IDEyIDAgUiAKL0Y5IDEzIDAgUiAKL0YxMCAxNCAwIFIgCi9GMTEgMTUgMCBSIAovRjEyIDE2IDAgUiAKID4+Ci9YT2JqZWN0IDw8IC9iYXJjb2RlMCAyMCAwIFIKPj4KPj4KL01lZGlhQm94IFswIDAgMjg4IDQzMl0KL1RyaW1Cb3hbMCAwIDI4OCA0MzJdCi9Db250ZW50cyAxOSAwIFIKL1JvdGF0ZSAwPj4KZW5kb2JqCjE5IDAgb2JqCjw8IC9MZW5ndGggMjYwNwovRmlsdGVyIFsvQVNDSUk4NURlY29kZSAvRmxhdGVEZWNvZGVdIAo+PgpzdHJlYW0KR2F0PS4+QXAiQSZVYklXcyQzNTFSZk45LmI3JCVMWlJyIm5sPD5qSl0rOUAxJyw9J2koWSV0JUQxP3NFYVxUMmZQKTFVa1lUPFQwKjZyTF8KKjBPP1RkMVJQX1BkU1huLV1JNFMkTFFAJEElVFpcQSFKbzxVMDBUMi1WImhNRk9FQEs6MTgsOk5QXklJLEByZXAqQSZUQ24hM0JPbl4hJiUKWjIhM0UtbUEqa0QkKS5MQGNzaHQ6KCxLaSY8Ii9XcHVWSilgUyVsVE1nbkg1TWJYY0Q8dWY2a3BKM2A/VCNDSExoL1lfMFtBKDZbcSJmKkEKXVxlVzlkWWMiKi5FMSEyazQ/TGI6UU5dXG1wYy4nJiMwR0gnRExlW21BZydAazJQUDRjNC0pVEpDdTpQKUdKVzYsMz5yWTdUc05lLUJZR0UKJTVGcTwnSWZILzxebjRJPF9TRlBeXGhmODZIUFI2T0tObChkQFlYaDI/bi5EVlNINjw6c1JTSGtwbFhzJT0mKi0pczwmP0B0TjwwNCQ+U2MKUm5fZTY4YzBvV3IqbEhdM0lfKDNpZG1oZmFbcDlcYGlrQSxjZ29XUXJCJ05eT3MwWipTbWdlTFRMOXUqLFVfYyloaWdIYGtCMWI2K3FAWiwKKVBkXSoiKTUqYDBxRFQmRyQvLGlxITRuQz9KW211KzI8NjddajZNUiYiJEMsXlllQkBrRkRAJ1A9bVZuPmcyJTJWOV1YVDIiL0tbUzpNMDIKKFZmcWVgcDFDRzMnR3VabWItZWtgMFNXPks0IWZWU0pXVzpdZ3QoRUgkIihRLmQiKS9sMFJbPWxfYk0uRE9qU0teM21vJExsVylAYkhIOFUKZms3VVtqbD10TXBTP2JqXDEmdFEwY1MoL2swKl5lMzUjOmhcR1JzM0RoZ0s0YyRGYlVoQ046MUtmRDRYUltvPkZAWDkzQzAsPmA5ZEAzKEEKUDA6J1JeVydtRVkkVlI3ckVEMjdPaztzdWY1P3U1P28qSzMycmFVLyJSWThDLW9MVlMuRm5VVDFOcltKYGJUQjRmPDdbW0QubmBOPDxhWXMKIzQvU1QzVGE+WksyKl1yK20ySW9NJDopYEszR0ljSD9fN00jbGxpPlEnUk0xVD1cNmZWRD8sYk1hQHAqYXUrTEUrLGRcJSU8OkwkPikoSmgKQm1qSV9LRkI0a01QOV5bSEdCQT9xcDNTaVZpSE0uQDlNYS47KkVWbFAnaE1hPkRoXDlwPDByLGorIk9OYlonP2NcVXAkXUJoLUhBKSM6Z00KSihQZyZoJnFORWhhU0ElUVpNZCdiY3RcUShxbC4+NSYqUDwrPS9IYWVLPFxXa3NBckRFbVNNbVBrTCw2LT1DL0EwXDFXME1pbSNWJ0M7JkoKXD9QTW5NMV9RVkJdKF5pLUxVZFswayIkUGdCNjVCNFBFVlJJXykhZ3I4WnIwZSYxQEYnW3FpWj9lOj00UVBDLUpbdWFGUEVeWCciaV5NPHUKNE4xLEJzK0FIckgyOlpUWDlZImFQK0ZmZC4kKEotclJFMV03VlEqXyJyJiIjWEB0S15nS0ByIT8+NENcImk7Uz8zIzkhLC1NKDpVaFtXVTUKLjYiY2ROM1IpUlhiOW5WLSpTaDBNWVc5MSMwKk1AXnNWIW9laycvSGAtJ2RlbHFGcycjX2s0UltqPiMzUlIwdEkvcCs9MDteNTIhRTxhaEIKUlxnRnJkajI1TzhoVGNoMyhyJF5XUGdaSmxDaXVgcDUncE1TQUk+TlE7RS51TiFBVlYvbloiQ11jKC1DVzVOaGlOcFJDUF4tWENsPCVJMFUKOzRTbjMlUCRvZy4ibFA5MShCV2FmL1onPWd0bSFgT3EtazJIVi5QSDdAPCM/XWpmTHNEZEAjXyd0QCM0XDYlaDEocShyaUk0TztMPmQiXEwKLThtJnBIIy9BUFEwQ25HazhcYzhTJGIvMzZkcClbWEIxLFxNSXA9P1RmKj1qJ2FjN2k9My4zSTlWUmlqMmBAWl48OF1ma0skPEN1RDY3b10KOUlsYSRcNl49P1FwZmpRYHAuK2NTWC1jbFRlTDtMLy5KLlU9SnErQEFkSGZXLUdnM2ovIjlSaVNaPjpRUVA2Wl1fLjs3O1swI15MPz1QRk0KJDlrIT9KbjImaDhbU1dkMiUlPD4uS3RjaDxOVSY+ViRhR0poYyFcZzJAVUghQkAqTmRST1oqWlgwZywjUEFoSCRmV2NsKiEuK2ZFcW5YUV8KQU1uTD9DWlVrbD9DUWxsLWtZRDAuXDFeYS83JkFOZyhdVSRrcks7MU5sQm5XLDpnUTlYJid0OlJbNFRhLUE4Nm4iZzZIKDE0dFZpVFFsclwKUXUhLHNpZC9NND5gTit0JD4yMEwvMGxhXT9LXTlYTDwuLklLKDVeQEwwJm89bV9fJUFLZU1LOSsibmpDNjlERGhSPmYwLUM1SjBuMEFJMm0KLnFNSCNdYVA4NDZJVSxeWSQ1UmU2IXUoMl9fMFA3WXBIVlViL2tYOSo5VDpGOjotZmhAKDArTEtTajQiKCZeQzMoWTo3J1k8Jkp1OlQxWTIKXFxqY0w/N0UoIW9dIyltWW9iZ08/QGhEWClYaCVrSGA8OD5jJmE/KV1pIm1DNjFWZCVZW1ArSTtDL1BPTUs4VnVTOz03YWhWRkYrREJZWHEKP2xOM2BjYFMsdS0qVTtUOCJrO2xNUTRdIkVpT0MuPDxCZS9DbFZsIz1rZD5Rb2puPFovIz9kamAsTUxZWzRqUS5eLk1lQideJ1lbUWpLU1UKPF5zN21dOzluOD9yWCZWXWlBYSJIM01xT1YwK3VsMz9LZkc2dUBaSSQxKnN1cSpHMSFgQ2xgRT5GRG8kbSQwPV1hMDo9bGQsJDwuLyY2JiUKbk1MQW5gYCNERWNqY3JVZkljNFZhajJhZFkiW1RLPyFPUUVOPjguMjkjXEtIXFVbPVxcITUtPSlgVUFaUyxjYlAvTEBRJEx1bkUpSmhlLjcKanRtZ19jRV5UWipGKkNrKUZYXkJhLzczOVk0cGNwUSIqQGc8XUZzOlFhM0w4MCE8Zks8YzotXls+cWpGWHNwYChORG5BP0RPaGVgKTVkaDsKMi01KE9XQkEsdGk4ci1FNUZvJkRyYiljWURFJ1ZhYTBiR19DK0E4Wy41XXB1KSJ0ZCUsKl5eKl08UFlDXk1LKGRPVVBtOUQ8bC47OkhpQVoKXzM5YkMjUEM7ZGwhYE44P1A1NDZHXk5jJ3BkcC9OUDxRQ281WkoiUT8/XEVRN101akg/OVxYT3BhTWYzbTVXaSg8TnBhJi9zcCNUOSktdFoKSz5IK1UyIkBnR09KLypEZHMpcXMqNmtFVE0lT1ZkIy5VWCtaKzBkRXIoaCg7SU5BPChiXlUzLT1yOkprWmY9XDVFZkRkczUqKylQVT83NUQKUWAlMFBqVCJCXGZhJX4+CmVuZHN0cmVhbQplbmRvYmoKMjAgMCBvYmoKPDwgL1R5cGUgL1hPYmplY3QKL1N1YnR5cGUgL0ltYWdlCi9XaWR0aCAyOTQKL0hlaWdodCA1NQovQ29sb3JTcGFjZSAvRGV2aWNlR3JheQovQml0c1BlckNvbXBvbmVudCA4Ci9MZW5ndGggMTExOAovRmlsdGVyIFsvQVNDSUk4NURlY29kZSAvRmxhdGVEZWNvZGVdCj4+c3RyZWFtCkdiIi9jXzJkOykkcSdeZjVIVVkhbzklMChkKm9RQlY4K3AtQTtZY2dcOWtFbXA6PlAiXGJjLEcoN3RERXF0SzpCNkxqSlZVT0plNDVFQyJRCmdTL1MnSStuNipYV2Qyb18+UlNtKUBxblJCNyMtby4rKSw+QEpyUWcqM3EtIlBhOWBOO0wxN1wocittXCZeW01yXDlJQShMPi83RTAyWHFhClkrLEwzIU5PTjNFLWJqT1VnJjtqV0FYWEpKXl5JJlBqVkgvNltgKy9UNmIvMSxLSEMrUF9RcidsQTxaalVSIU5gL18vZENeIUFEImkwXUs2Cm5MJFRUO0hjXUUuN0loVFs3czFrMnUjQmJSXVNHUnBzSkxqKUxQUjYlSHQsVnIyWFVRMDVzKyoha20uSTRIV2M9OUJlWjtNVV1cRlFLZDNuCi1xPXVbRz0+LSZVNlhRWjw/NDgsQVJlNkFRJFFAbzlcWCcpREFFKm1bJDZtKXMuYFAua3E9JjVmW0FBKFhFSTFFbEJjQEVgVzNgMHFHajR0CllPQD8kSSdANWAwWC1iIzNNcVkwWyQ1SCxuZWlSVGYhOCxbVSw3IzU6Oi10Kl5wKi1HYW1sTGlcMHFmP0gxYjs4aXEtYD0vXFp1KDhlVl5WCjk1LGhcVUskP2E1aUtxSztoYCYoIS9nP2tqbyYjTUxbP2pdNEc+UFpbUkY1aihsXTgqaW5XZ3JGJVpwYi40cUJzUF5ISFs5M0c8RHFga2s+ClJXZWhfOzpRa1hiaVJEbjBgUiQwPVZ0MUBpYTwqYC5BVUknNHUnSUFbJDVILGRYVGZATjJxJXMjY1YpdVYvV2dqNjA3LWVvLU5oKjRkWy1FClIkN0FGZEBrPUJFXC1DNnBeMVd0VjZsSik7WmI9JUkhZUIsWlZccS1uczozI11AXC4xQlVyQywsTlhZTEombz9TNC5ML1gvNGVxaFU8Y19ACjhzXnVkQHUxPVkyTHJgLFUvKnQjbWRhX2EvXDxSVi5WTlBaMUNiZ2cmY0ZFRDc/PExdRzRkIW9UJGwlYj4iXE47J2RqVUhPY1JnTG4wQlloCmFbcjNzTmxZWE0qZT4+K0NhIUNoVWtNSlJEKzNbNmBtYDhjTVBbXEVER3FJZk1vXnNXTlxlWUJXZSlbUUM4Q1FDXyRGaV8qY2hEJU06LitrClpRT0FrZik2LmteM1RXMWI3bjJSOUJlVmRdY0ByUzlfXikjPTY/SnVWYmBxZWgxPUY9W0dOOG5FJS5YU2RRQCFaRjo5RkwqVG4yYydcPl1TCjhrXEQkVi43RVZNcihRPVooLFtrI2YlYipVaS4jPiFFYltoXXFpYjxjZGwzJD9wXEwjXF9YS3NOMlNrRFxzL04vPnRYWWNiKUF0UidTSipJCkAnSzJCSS5SOVhLKj9MaSg8VTEmNkEuW2JQRDZRM0Y2TitpMzV0VHE8XitRaWRTZnBJOHBnTzZxPyg4MlQtRn4+CmVuZHN0cmVhbQplbmRvYmoKeHJlZgowIDIxCjAwMDAwMDAwMDAgNjU1MzUgZiAKMDAwMDAwMDAwOSAwMDAwMCBuIAowMDAwMDAwMDU4IDAwMDAwIG4gCjAwMDAwMDAxMDQgMDAwMDAgbiAKMDAwMDAwMDE2MiAwMDAwMCBuIAowMDAwMDAwMjE0IDAwMDAwIG4gCjAwMDAwMDAzMTIgMDAwMDAgbiAKMDAwMDAwMDQxNSAwMDAwMCBuIAowMDAwMDAwNTIxIDAwMDAwIG4gCjAwMDAwMDA2MzEgMDAwMDAgbiAKMDAwMDAwMDcyNyAwMDAwMCBuIAowMDAwMDAwODI5IDAwMDAwIG4gCjAwMDAwMDA5MzQgMDAwMDAgbiAKMDAwMDAwMTA0MyAwMDAwMCBuIAowMDAwMDAxMTQ0IDAwMDAwIG4gCjAwMDAwMDEyNDQgMDAwMDAgbiAKMDAwMDAwMTM0NiAwMDAwMCBuIAowMDAwMDAxNDUyIDAwMDAwIG4gCjAwMDAwMDE2MjIgMDAwMDAgbiAKMDAwMDAwMTk2MyAwMDAwMCBuIAowMDAwMDA0NjYyIDAwMDAwIG4gCnRyYWlsZXIKPDwKL0luZm8gMTcgMCBSCi9TaXplIDIxCi9Sb290IDEgMCBSCj4+CnN0YXJ0eHJlZgo1OTY3CiUlRU9GCg=="
    }, {
      "id" : "ea3c2f15f4ab44a8a864dbc7c65959cb",
      "tracking_numbers" : [ "794684898595" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T11:26:11.522895Z",
      "updated_at" : "2025-03-12T11:26:18.686909Z",
      "succeed_at" : "2025-03-12T11:26:18.675965842Z",
      "meta" : {
        "source" : "tts"
      },
      "source" : "tts",
      "ship_from" : {
        "city" : "Centerville",
        "company_name" : "Postmen",
        "contact_name" : "Contact Name",
        "country" : "USA",
        "email" : "usps_discounted@test.com",
        "phone" : "18682306030",
        "postal_code" : "45458",
        "state" : "OH",
        "street1" : "8700 Meadowcreek Dr",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Santa Ana",
        "company_name" : "Maersk",
        "contact_name" : "Mister Jones",
        "country" : "USA",
        "email" : "maersk.test@example.com",
        "phone" : "8777285325",
        "postal_code" : "92705",
        "state" : "CA",
        "street1" : "1600 E. Saint Andrew Pl.",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : "202503120005",
      "return_shipment" : false,
      "files" : {
        "customs_declaration" : null,
        "invoice" : null,
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6"
        },
        "manifest" : null,
        "packing_slip" : null,
        "qr_code" : null
      },
      "box_type" : "custom",
      "references" : [ "order950000002" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "15eebae3-1646-4349-88f3-e8d0ec00e577",
      "service_type" : "fedex_ground",
      "rate" : {
        "booking_cut_off" : null,
        "charge_weight" : {
          "unit" : "lb",
          "value" : 172.7
        },
        "delivery_date" : "2025-03-18",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 201.07,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 260,
            "currency" : "USD"
          },
          "type" : "oversize"
        }, {
          "charge" : {
            "amount" : 25.36,
            "currency" : "USD"
          },
          "type" : "fuel"
        } ],
        "error_message" : null,
        "info_message" : null,
        "pickup_deadline" : null,
        "service_name" : "FedEx Ground®",
        "service_type" : "fedex_ground",
        "shipper_account" : {
          "description" : "fedex compatible opp test",
          "id" : "15eebae3-1646-4349-88f3-e8d0ec00e577",
          "slug" : "fedex"
        },
        "total_charge" : {
          "amount" : 486.43,
          "currency" : "USD"
        },
        "transit_time" : null
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.2
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 172.7
      },
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "in",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "F231",
        "item_id" : "1313131223",
        "price" : {
          "amount" : 50,
          "currency" : "USD"
        },
        "quantity" : 3,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "method" : {
          "account_number" : "123456",
          "country" : "USA",
          "postal_code" : "123456",
          "type" : "account"
        },
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : "JVBERi0xLjQKMSAwIG9iago8PAovVHlwZSAvQ2F0YWxvZwovUGFnZXMgMyAwIFIKPj4KZW5kb2JqCjIgMCBvYmoKPDwKL1R5cGUgL091dGxpbmVzCi9Db3VudCAwCj4+CmVuZG9iagozIDAgb2JqCjw8Ci9UeXBlIC9QYWdlcwovQ291bnQgMQovS2lkcyBbMTggMCBSXQo+PgplbmRvYmoKNCAwIG9iagpbL1BERiAvVGV4dCAvSW1hZ2VCIC9JbWFnZUMgL0ltYWdlSV0KZW5kb2JqCjUgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9IZWx2ZXRpY2EtQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjcgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhLU9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iago4IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0hlbHZldGljYS1Cb2xkT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjkgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvQ291cmllcgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEwIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjExIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEyIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZE9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxMyAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Sb21hbgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjE0IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL1RpbWVzLUJvbGQKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNSAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1JdGFsaWMKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Cb2xkSXRhbGljCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKMTcgMCBvYmogCjw8Ci9DcmVhdGlvbkRhdGUgKEQ6MjAwMykKL1Byb2R1Y2VyIChGZWRFeCBTZXJ2aWNlcykKL1RpdGxlIChGZWRFeCBTaGlwcGluZyBMYWJlbCkNL0NyZWF0b3IgKEZlZEV4IEN1c3RvbWVyIEF1dG9tYXRpb24pDS9BdXRob3IgKENMUyBWZXJzaW9uIDUxMjAxMzUpCj4+CmVuZG9iagoxOCAwIG9iago8PAovVHlwZSAvUGFnZQ0vUGFyZW50IDMgMCBSCi9SZXNvdXJjZXMgPDwgL1Byb2NTZXQgNCAwIFIgCiAvRm9udCA8PCAvRjEgNSAwIFIgCi9GMiA2IDAgUiAKL0YzIDcgMCBSIAovRjQgOCAwIFIgCi9GNSA5IDAgUiAKL0Y2IDEwIDAgUiAKL0Y3IDExIDAgUiAKL0Y4IDEyIDAgUiAKL0Y5IDEzIDAgUiAKL0YxMCAxNCAwIFIgCi9GMTEgMTUgMCBSIAovRjEyIDE2IDAgUiAKID4+Ci9YT2JqZWN0IDw8IC9GZWRFeEdyb3VuZCAyMCAwIFIKL0dyb3VuZEcgMjEgMCBSCi9iYXJjb2RlMCAyMiAwIFIKPj4KPj4KL01lZGlhQm94IFswIDAgMjg4IDQzMl0KL1RyaW1Cb3hbMCAwIDI4OCA0MzJdCi9Db250ZW50cyAxOSAwIFIKL1JvdGF0ZSAwPj4KZW5kb2JqCjE5IDAgb2JqCjw8IC9MZW5ndGggMjU1MgovRmlsdGVyIFsvQVNDSUk4NURlY29kZSAvRmxhdGVEZWNvZGVdIAo+PgpzdHJlYW0KR2F0PS5iQUxORScmPHI6NTdtJFEmZzAsRFomNFYvcih0aGVwc2tbQlxlajRKQExcRmBmWiNvPWNDaHFkJW5UP1lgUVEzTCUzUlhbXGBeMiEKYEY9c28pKF1KYWBaKyoib2o2O29LMyUnWCY1cC1bVmZnIV5WOXVDNChSK2ZTV0V1OzFvMU8pbXByOzUobUZmIi9RaEtVK1wqbzpgQ0FzaV4KKEgpNS5VNVtPcUxVS0soKyxmbT0zX3RWXkZeZmVVOyhwZzZTPVI8PjI/MyxVM3EpLGBARSE1KkdQY0BQa3M6a0hoKDNgRTErVEJyXTw7YCgKTXFhZSkxcDojNWYzNF5yWS0uNi1JQ2BjOklUTT0hXS5oc0sySz5ZSEszUSVnY15GVmFLRDEiLjBZbVIlNW0nJ0ovTWItR1pKJyNuMmkyNDwKQkxsUWNZJVREJ2FdTzVzPjs9QVVycGVIOSREWDY2XykrJWdOVj9TMWtZX3NeKDY3Q10uLSdeaUU7R15rTkpqVjwvbSczX0s+Pj1oVycqNygKM0xMciY2Xi5AQDs+Uk0jKk0jQkhmI1lYQE9QK0ReIjYiJHJmM1oxWV40MyZQNlNiSFdfOHFFQzBsZVZnLyE/W1ZKKmZjbm5Tb1Y/NGM7WyQKWW0rcFpLdWRkYDM9X3RNbzUwZlYjNio/LCNFLj5sYFZsRiReTzAxX1pfLTJXUUVaa1EqY29XbmtHKipAXDJVJW87aGUnQC9rQy5xSSo9aiYKI1AyOmVeZlZNRlc4bWdza29QaUA2WXUjWF1gISFIZ1E9b2tmbmpmXCxrbUBbTzRBJm1IOHJUX2ByJF40MHUpIV0tX0ovai8qYzBObUdkR2wKQUY+VEhDMnM6SHBEbXNZKkd1KTE+dGtQXV1EP3BQVF88OSxtaGFtMnEzUE1bQF87J2NsZHRhbUNwJmZtcDg+c3JOQjAxI1ZlOig5Syc7bW4KKlFgRy9PXkRTJm48Xj5DVnNtMUJkbiZnSlYiSEZjWS04YltyJHVvb1ZdcEViKihnaF0haUw/cEptY1FJX0tecSFqNzpgYGphUnBpRnN1ZVsKSmpObCorW29wX0BBMEdTS0VoaCJsK3M3Yz0+VmQkTilrQTw1dS85SmUsT3AtXlxkJ2I9cmxIb0lTJUpaI0thcFsmIykmRV1GXXBpTnRIRVIKU1ZgMlRfWkZDM0I7UyNkVHMtKFxMWUskZmclUWg5UXI8alosbHEjLEBxUUlNP1EmWyUiTkZwXSlSJUMtSW0wRT5FRmQvXzVoUUh1Sm5NMEMKLiZlbWZlcFIocWByUkQ5KiYtPUAkKDQuWXFWWFpmMTUwaWM1Ii8wTFpYKGszKGJaaVRMKzRdVnE8dEJdOk50Q3Vdalc4WSxYNjUvMXRzT1kKRUVsdG8zVT1YZ2lhb05aPFk1aSEuMFdIKCstaEY3UXUlb3EmbDo9ZFVuTV1bLTdTb1VJVkFxOSZeLSwwMWRlTlIlIWlLb1dMSGEtLDJqdEcKZTxedThmYUlGaiRWRlJFJW1EcUlEa3JMV1YnVCtEUm82XWNHRVQ2YzJLby0nKkx1Pz86b1xOKzJCYSlcPVhCZ1g2TVMhQi5VNW8uLypCNSgKO19MKEU5SkxHdEk0W0M7ISc8bmEyamNfJ18rM3JyPXNoZTBWVTBTTytlVjFMTGFeSyQsLjMoUj5cPmw5MV0hOEM2NW5EPClhWkViRjZMT08KMSRLVmtPM3Q9ZERCJStkQ2UzPnFOYmNORkYnYFs6Xz9kQUlhdEJgOyNzTD87Um49RXU9IiJLQTJuaVRCNERbXGFXNTxaYUgnRXA2YFIsUzsKMGNNJnUyNkpfTFljLTcpOUtGI0siaCg4NiVJNEY+LF5icis5RV1ncT1QK1FSJ3M7VzlZZUcpRS1wcS9eWz1aJDFvYTpMLVszUjExTUpcJjYKTlUlOU4tJFpSc11MTEg7PlVwTSc+ZCJdZFM9OENyb0pQaTU2RVA3QDY/LjZSWm51bTVAQF0vQ0ElSlMzW2c/cis+UFdARGZLImtLcCpEJD4KJ01Qcys7YilAZCRaJTw0RWxqKj05QE9UQzFNSSFFQ0Q9cUc+UlhgLGReIWY6ViVJKkpIOCktIlFUTCo/MW0yXyRWQGQ0PjxARXJPQFVPIyYKZGU5RnJKMk9QaSElQT0/MURpVkdLXCFHLmoqR1JKQiZIQDVhZmlETFhicTYoZlA/XDlpXiQqWyI2YT5DMzxKUVQkOGw6cUg1WStLQHIqJmAKaUhmP01sb0omTTMjcytiLUNbP1ZcREFvUD9fV21eXy9laDBWO3FvZEdPU0VTazMqUCVfJDhYRmAoLD9rX2hbJnQlVVVBMGwyN1kwWmojY10KLjxxcWZOWiJbM0Y1a2NkPzc/U0czVk5rL1sjYEo6WClRbF1iPHBAQClXZHBrYm5PQGFIX25BLV8oM1hKa3JjWys8PEtrMFVob0pnPnFsMHMKVi85Skg6YjdDOiVZIlgmVjdGSXBeZz0hKi0kZ1NUNldha2ZCczVCciRTPi5eMF9Vcm9NcDc/XkpcPXQvajphOThTOkQoOkRGckJvRENJNC4KbFRFVDUyLEEuKj5kbU5YSFRbKEMoQnM/Uk9KbWAkO007PjFuKyc8Sj9yLGlYbzE1a2NHIlwqITRtYmheQEJcOGlNQGo3IlBnU1MwX19YI14KQDsyPF5lXEgmTy5JR2I2NClFJkwvIy8uKl07N2lOP3JYJStdaE48c3E/WDwsTVEjbjpJIyplW1UsS29gIydOWjhaS0ssMSNzKiRIbTJaNiQKOG0sRy8zKyw7KGhDTEJPSkRGKy4yTGBfJ0pSXmxxREZkZENoUDRCcmIuMHUsPzosbkJENWRscC8hWHV1WjQhMTdiWTtnSldXcjA2JlcnZFIKZ1Y/UU5AZ2dwTTBXOSEpTlk5JEo+TFBXcmhvOWZpYnA3MnJXQkNsbyI7SG5rVGwiJjxQPkY8W2UvZz8xa3BXQVdbI09BbDNfK2hHIjxAMXUKI0RrPUAuQz9bWjFTXCRgPCFMbltbTT9sIUljT09CakQrSE9RWEJJVmYsWUo2NTIqSnRAPGQuVyVYQDxWZCNSYCdnWCRTQjxeSTRII2FQKEoKQW5rRTVPJWhDPSxmPTZfV0hIam1rOm1haTBzRjZJI05KJiJrW0JTIz9kQ1BIViMySkYnckRmXUEnNFonbDMhKSZLXVJLR0dARSI+ZlAjW1QKVzciVU1hMUErVi8jcWYqQjcuKTY+MUUkbkdEIV9rMGJCXHBkYV9XWEJhLjhrOG0iXEJVRy03Rlw1PFRQNDVzXmxxL1o7J2dpMyc9XXRgUCsKLkw4Ukoya29MLTtHWFstMSM3cl0zaGc1NS1nL1AhLS0+JTIjYydpfj4KZW5kc3RyZWFtCmVuZG9iagoyMCAwIG9iago8PCAvVHlwZSAvWE9iamVjdAovU3VidHlwZSAvSW1hZ2UKL1dpZHRoIDExOAovSGVpZ2h0IDQ3Ci9Db2xvclNwYWNlIC9EZXZpY2VHcmF5Ci9CaXRzUGVyQ29tcG9uZW50IDgKL0xlbmd0aCA0NTAKL0ZpbHRlciBbL0FTQ0lJODVEZWNvZGUgL0ZsYXRlRGVjb2RlXQo+PnN0cmVhbQpHYiIvZUpJXVI/I1huTGdUNjxkZi9JLiY/Qj1oY0RhZy1yRDRdMCVQZk1UWitgUVNJK0pdQ14mckQzQThsTThubkZGUCFJLmYqallDMlRuWwo8N2koN1w8SkghPHVWdW0zSEpmazw4KSslWk4qTWYtcyVsPlZaVUNzLDlgOW9XJlRNJS9jZ21KVSUuZkxoYkkkJz5CZ3VfaUBRK0MmYldyNgozRENNNi9GVU10LERCXiwqUjAwcyJdV0ZHQzBqOiNXYTRnVzwqUWVUOyw/O0JTIW9ydTV1JF8zQjJqXDJDLidiLkZYcUBYMUEqLGFcU2FHSgo6bk1rY0lRXC5KZjowNTdfLUw0SjwlMlFNbUlhb0BcODhGS0diMXAsbVc7JS5ITkMuO1YuK2BkKXM6ckJoVWpMSEJYMlJDXjlfcTgxblZbSwpMPTJRQkZ0O3JUQ242W2ZwbXByN1NfY0chZFdbLHNYMSsrOFI8KS0lRmNVLmtScFdTKVhFOzkrOzRsUytfNkFzcmNFYlFMY0lFY187JU1LRQozY2pHVDkqTltKLnBOMldtSThXPyRsO2M2XTVkLz9wOHQ8c0VhSjJlcmwrfj4KZW5kc3RyZWFtCmVuZG9iagoyMSAwIG9iago8PCAvVHlwZSAvWE9iamVjdAovU3VidHlwZSAvSW1hZ2UKL1dpZHRoIDU0Ci9IZWlnaHQgNTQKL0NvbG9yU3BhY2UgL0RldmljZUdyYXkKL0JpdHNQZXJDb21wb25lbnQgOAovTGVuZ3RoIDE5NgovRmlsdGVyIFsvQVNDSUk4NURlY29kZSAvRmxhdGVEZWNvZGVdCj4+c3RyZWFtCkdiIjBPM3UzPmgkcS9NMEooJVM3LEt1WTwtcy5sSC0ncUVTKHBST0hDclBuJDdaOSlhNiJMUidTI29DNToiRWxoTCQyMyUyPzpfJ2xjWS90CjkmQzQvTkByT1tbRSdqYmBzPlw6UHBqbTU8QyxNPSZYO28sQVFqbWZha0ptRyQwRHBUPGAvV0hDSUtkaDhwPDBrSTI/YWFVaUlMMVczIkMhCjZEZDY2bCw+V15IW2YuPSQpckxuVGlTZkk0M0A5STxXfj4KZW5kc3RyZWFtCmVuZG9iagoyMiAwIG9iago8PCAvVHlwZSAvWE9iamVjdAovU3VidHlwZSAvSW1hZ2UKL1dpZHRoIDI5NAovSGVpZ2h0IDU4Ci9Db2xvclNwYWNlIC9EZXZpY2VHcmF5Ci9CaXRzUGVyQ29tcG9uZW50IDgKL0xlbmd0aCAxMTYwCi9GaWx0ZXIgWy9BU0NJSTg1RGVjb2RlIC9GbGF0ZURlY29kZV0KPj5zdHJlYW0KR2IiL2M2I1BhSyRqNjImOk1VUSQ9SlMlW3AoWC5bIlU3YDY3P09HNm1fN2tVSGVzLWJUKVNZV0whUkE7SixaSkRwQEovMWU7bmRXUiw8P0IKRjVKU2k4XDFdX0lDZF9NWFJXZDxdQiZAR05ZNmNJPmg2WUheZDU5JEZIbGNuZFxqQ2sqSV9YNFM7UTExPVJYTnVeN3ExMFVoVG9WZzlTYmoKJjVLLV1tQjBIKWxpT2VqKUU5bFA4Kk87Tyo3V0NjN1EkM0tCQiwsTC02VC1uR0FRUy5wcXRDKGtuO1dIJTZlUDtkZktiYTJFRWdITXFMcmcKOSgsXTliVGlwVDVcb2FgLmpcK0BEPUlsMmZwN08jaGBJJixvWFFwWVgpJTcuYCZWL04+Q0ZnMlJBNGtnaFBLOlNsVGRZSTdJTVFpM01OYEAKRzxASlhVOls4aGpTY3E8UFg6NDlWR189aiZDXj1pVDNGYEsjVF1iOm9LIjpSO0czWChiWz0pQV1kaWBkPzFVJ205L0hrQlMjQU9kKVlYJF8KZD1pKkcoJyJIczU9J0VlY3A8XT87MG5AZm1mNXRmLz42Ky9CJkQ2dCgwKS4mQjpqT21uWkNaWC1rMWxVb2loMTFrdVddbjUpLzxtLTY+OjAKZFNZPU07TywiR2IkL04xLFJCJVUhIzQwRFFnOj1KR2hYZ1pbQmtnKENMZTA7ZVFDY1VjRHQ4NDk0Rytxa3BYNTthXUNub1BbdS9TMDE4aycKN19sRFZGSmFxcjlRJzUlbzhmOS1LamYqTSVdMXBQK1deN08oUiUnVTgqT2tfR2plSks8QkAoL2laUyJMXydrNCJhbilZIWpuRUwjVkE5WCcKWy0qUFI/Sj1MQ2M1bEVpYmdLI0cuKHIzTyg1ZSY1Xio7XGpQXU8jKixtTmBoRy9sLTxhZVZFT2pkJkFjYVpiNGc6PCNAREcsKk0lL0RNNUEKVVwkJTRGWTBFbGhabk81TkZCQiMpLCU5Wj8xaT4hZk8mYWFnQDgtR0QwJTdQUGFSLkJISE1YJVgsYlYqSENOPkIlUVBBRm1WP0NicnFWRDMKbHJZZGdlZltFMTVUL0heR2FeaWg4Kk9rX2N0V2xsSz0tK2JxMV8uNlNATSZJXihCdW1CTlhlN0hMQDNsL2RIWjshLjE6VVxEIipaLVBBJ1AKVjUqWiwtJTI8bz8vYDoiQisvPlUiTEFiJTk1XF9ROU1mLk02Y1E0SltPLjQ7NCpHKChlaTFgbCljXiVeNEhAaWVIQXRnQk5obiQ0MyoyJ0sKKUI7Kk1LcjIoJlRpbDE6RTYyISYnOnJqLy5HJkkvMmMvMSlUY0soUWRrXUxMRkhrWFpWIkNMXlY8X1ROWXUmclVsU1xzW1RKbzJTVGI9VjUKMWpXPFw6O3RoJyY5WVo7M3QhJzUiO1FgQG5wP1VdZGJDIkxOTlQmaFB0cVNdYXJCLkIuQUI/SUtpZm0qLWE5Nl43NkA7KmcnOHEnN21uYkIKbWtbLWUqTW9QJ2EtPyNmczc2YkxNRjxtfj4KZW5kc3RyZWFtCmVuZG9iagp4cmVmCjAgMjMKMDAwMDAwMDAwMCA2NTUzNSBmIAowMDAwMDAwMDA5IDAwMDAwIG4gCjAwMDAwMDAwNTggMDAwMDAgbiAKMDAwMDAwMDEwNCAwMDAwMCBuIAowMDAwMDAwMTYyIDAwMDAwIG4gCjAwMDAwMDAyMTQgMDAwMDAgbiAKMDAwMDAwMDMxMiAwMDAwMCBuIAowMDAwMDAwNDE1IDAwMDAwIG4gCjAwMDAwMDA1MjEgMDAwMDAgbiAKMDAwMDAwMDYzMSAwMDAwMCBuIAowMDAwMDAwNzI3IDAwMDAwIG4gCjAwMDAwMDA4MjkgMDAwMDAgbiAKMDAwMDAwMDkzNCAwMDAwMCBuIAowMDAwMDAxMDQzIDAwMDAwIG4gCjAwMDAwMDExNDQgMDAwMDAgbiAKMDAwMDAwMTI0NCAwMDAwMCBuIAowMDAwMDAxMzQ2IDAwMDAwIG4gCjAwMDAwMDE0NTIgMDAwMDAgbiAKMDAwMDAwMTYyMiAwMDAwMCBuIAowMDAwMDAxOTk5IDAwMDAwIG4gCjAwMDAwMDQ2NDMgMDAwMDAgbiAKMDAwMDAwNTI3OSAwMDAwMCBuIAowMDAwMDA1NjYwIDAwMDAwIG4gCnRyYWlsZXIKPDwKL0luZm8gMTcgMCBSCi9TaXplIDIzCi9Sb290IDEgMCBSCj4+CnN0YXJ0eHJlZgo3MDA3CiUlRU9GCg=="
    }, {
      "id" : "06828c6d3e814dd88e4dd2d61df0689e",
      "tracking_numbers" : [ "794684898460" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T11:25:31.407615Z",
      "updated_at" : "2025-03-12T11:25:39.230995Z",
      "succeed_at" : "2025-03-12T11:25:39.187855657Z",
      "meta" : {
        "source" : "tts"
      },
      "source" : "tts",
      "ship_from" : {
        "city" : "Centerville",
        "company_name" : "Postmen",
        "contact_name" : "Contact Name",
        "country" : "USA",
        "email" : "usps_discounted@test.com",
        "phone" : "18682306030",
        "postal_code" : "45458",
        "state" : "OH",
        "street1" : "8700 Meadowcreek Dr",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Santa Ana",
        "company_name" : "Maersk",
        "contact_name" : "Mister Jones",
        "country" : "USA",
        "email" : "maersk.test@example.com",
        "phone" : "8777285325",
        "postal_code" : "92705",
        "state" : "CA",
        "street1" : "1600 E. Saint Andrew Pl.",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : "202503120004",
      "return_shipment" : false,
      "files" : {
        "customs_declaration" : null,
        "invoice" : null,
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6"
        },
        "manifest" : null,
        "packing_slip" : null,
        "qr_code" : null
      },
      "box_type" : "custom",
      "references" : [ "order950000002" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "15eebae3-1646-4349-88f3-e8d0ec00e577",
      "service_type" : "fedex_standard_overnight",
      "rate" : {
        "booking_cut_off" : null,
        "charge_weight" : null,
        "delivery_date" : "2025-03-13",
        "detailed_charges" : null,
        "error_message" : null,
        "info_message" : "No rate quotes returned from carrier.",
        "pickup_deadline" : null,
        "service_name" : "FedEx Standard Overnight®",
        "service_type" : "fedex_standard_overnight",
        "shipper_account" : {
          "description" : "fedex compatible opp test",
          "id" : "15eebae3-1646-4349-88f3-e8d0ec00e577",
          "slug" : "fedex"
        },
        "total_charge" : null,
        "transit_time" : null
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.2
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "in",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "F231",
        "item_id" : "1313131223",
        "price" : {
          "amount" : 50,
          "currency" : "USD"
        },
        "quantity" : 3,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "method" : {
          "account_number" : "123456",
          "country" : "USA",
          "postal_code" : "123456",
          "type" : "account"
        },
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : "JVBERi0xLjQKMSAwIG9iago8PAovVHlwZSAvQ2F0YWxvZwovUGFnZXMgMyAwIFIKPj4KZW5kb2JqCjIgMCBvYmoKPDwKL1R5cGUgL091dGxpbmVzCi9Db3VudCAwCj4+CmVuZG9iagozIDAgb2JqCjw8Ci9UeXBlIC9QYWdlcwovQ291bnQgMQovS2lkcyBbMTggMCBSXQo+PgplbmRvYmoKNCAwIG9iagpbL1BERiAvVGV4dCAvSW1hZ2VCIC9JbWFnZUMgL0ltYWdlSV0KZW5kb2JqCjUgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9IZWx2ZXRpY2EtQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjcgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhLU9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iago4IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0hlbHZldGljYS1Cb2xkT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjkgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvQ291cmllcgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEwIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjExIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEyIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZE9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxMyAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Sb21hbgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjE0IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL1RpbWVzLUJvbGQKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNSAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1JdGFsaWMKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Cb2xkSXRhbGljCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKMTcgMCBvYmogCjw8Ci9DcmVhdGlvbkRhdGUgKEQ6MjAwMykKL1Byb2R1Y2VyIChGZWRFeCBTZXJ2aWNlcykKL1RpdGxlIChGZWRFeCBTaGlwcGluZyBMYWJlbCkNL0NyZWF0b3IgKEZlZEV4IEN1c3RvbWVyIEF1dG9tYXRpb24pDS9BdXRob3IgKENMUyBWZXJzaW9uIDUxMjAxMzUpCj4+CmVuZG9iagoxOCAwIG9iago8PAovVHlwZSAvUGFnZQ0vUGFyZW50IDMgMCBSCi9SZXNvdXJjZXMgPDwgL1Byb2NTZXQgNCAwIFIgCiAvRm9udCA8PCAvRjEgNSAwIFIgCi9GMiA2IDAgUiAKL0YzIDcgMCBSIAovRjQgOCAwIFIgCi9GNSA5IDAgUiAKL0Y2IDEwIDAgUiAKL0Y3IDExIDAgUiAKL0Y4IDEyIDAgUiAKL0Y5IDEzIDAgUiAKL0YxMCAxNCAwIFIgCi9GMTEgMTUgMCBSIAovRjEyIDE2IDAgUiAKID4+Ci9YT2JqZWN0IDw8IC9GZWRFeEV4cHJlc3MgMjAgMCBSCi9FeHByZXNzRSAyMSAwIFIKL2JhcmNvZGUwIDIyIDAgUgo+Pgo+PgovTWVkaWFCb3ggWzAgMCAyODggNDMyXQovVHJpbUJveFswIDAgMjg4IDQzMl0KL0NvbnRlbnRzIDE5IDAgUgovUm90YXRlIDA+PgplbmRvYmoKMTkgMCBvYmoKPDwgL0xlbmd0aCA0MTU3Ci9GaWx0ZXIgWy9BU0NJSTg1RGVjb2RlIC9GbGF0ZURlY29kZV0gCj4+CnN0cmVhbQpHYXQ9L2JBcGpzJmAnZTlzKyUrY0BUQE4oPTYkIjFMYmxJZk1hTjFTW14sQkxaL25EKC9na2ZnOzxAK2IhKmtRKjYrPWhxLzByZjtGOEdycwpJS0NAcSJFM2dNZnFtYXJjNmhWYD1RP19sQ3IoMV5aY2hGKFokOiM0ckdfO2xfRzBMbycpT2I1YW5XbTc2VC5IJG1fPXRkTTBMUElPPTpEVgo5ZG5yYE9JL2VRPTMkNDg/WzVTb25ZXzo9azRsXCFpc1xjWjJ0bXA7LUE+cHUvZjVEWmdXLHA+YlREblRodSlIVmo1Mk01YSNhbVk1Iy1qNQpHOzo2MElcYGc7MllTKz08X0FaMyNaY0lja01UNVksPWotQSZuPCd1SUNzcSxuRztGQGZrdWheUWNBamVjKW1uQThUYWMnaW9daCxycC9UOwpqLU4wcm1BbS5BWmkzX1RwP01NRCZlQDBQcV87Mm9bPTlTWUVybT9rSk1DLjhDSioiTEhJQ0k1bV0yXTomT1BpSXFSVStfVzVhb0FqIyFeOApySEQmJTs8VWNwXToyPEkrTGBpQEd1KDlKcV1gSlxHSzRiSlM7LUBoQGxNNCo6QyxLO0NjSWlmcCIkNWhpIlEmJi1tdFdbV2thbl5BcCxWTwoqbFpgMTteL2c5QiQ8b15ZJFAuKS4lYzUiaz1ZPjZKaHFGJDhfRUtsaS0vZD4oQCRHcSckTUpvbS1pXi5jMGEnNT0qbyJySTxrZWpCKEg5awprLklMZTgtR1wpK1MqNiM9JlNTKS02KT03IlltcGtqP1QlanFdaFdQbG8zUC4tXSwvJVBEcFRDZj1WIjFVbiJPIkE4QkhfRFI3JCVJc0paXApJOktecDpRWkpENFsqVThZRXM6WGN1Kz9naDBzU0BwP2g5RGsybk5mQT5KY1FSSTNpNFwsMzVkREpLX1BgNTcwJnBALSMxRl0vMiJCKUMhYgpBJSImSSs5JDlxVFosNGJvaDUkWWxRJkhPK1phL1M6RHI/WC83L050PUNrJzA3JWpbLy8vVC9iSDcuX0tsMl1BJjJgTFdpaUpHRSJTOmthJgpbdGVtbi1ta0lNZSk8QnA+S2dNc1s2TCwpSVhBXFJEci8wNVNnYDZgYHFTZztmKCpNMkRJMl9lUmNlOjBBK2hpT2RvLCYvQjI5NFo6UlhDXgppMi5AWWBhIUNUcCI7TVFGNE8+Nl8uWVVgRTRWWy9eNSZLRm4lVV40UWZmJ1NSIygzVmFYZC5BM1VISW9oNDsxJ0FRNSNjcWRuOUpMOmYzagpmTHI7NFloMyttRWFJXDtyLWlnKW1sa1lWNDwtJlFgWix1K09JQWpeVl9DRXVua0Y/cGYmbWo5a3FcaF4pKTI0I0VLJXEsa1ZCWGRPQ2w/WQpPTDtXZ2RtJ2xtXlVXbDhycExYPklPdV1TL1VWRCZNXkhLOVM6VyZPVm4jJjk1STM+ISJzYGAjMC5tYEE+ZSJqTm1tTUhDckBwLnQvWSZANwpVKT48aktmLltKP1pKUCRvIW8+dSU7dDYtTydHLF5uMnJMI1U+WmgnZE84L3RPMHFBOCgvSGEmRkQ9L1xUKTBkMSRfWDpEVzMnL0MjLGBgYAohciU5KmVJakE2MT0lYm1GRD9fRi49SV4kZnA1SEwkOGBEaTd0XUlNKzJjRHRHKm80OWVkUDtGUGEyUDQuJSVnWlBBJnFpXT4wVGJuZWc4UApAQ2pSVlsoJlRuO0BYI2EuU09ILy47QSZbXSRjTjAmS2VpYkI6LXU9WlFIKDBXaWA1PU1kVTlVOU0oRzZjbFJPPD5jOk0vM0NpPTU8cFkodAo0amxnRVBwX3RyOGVORDc1L3JETzhDXCoxRiwmb1wxYnVRL1VFbyVXcSdzY2MwXHM9NU1rclZdbSlmJk8ialMhX1VuT1ZBQjBhJ0kmP0NEZQoxYyVmcztocFcpJW1JTkxSRnIhSWA+T0dDVzVtYFBSTzsyZFgiY1QpV0IiMTtEMWlKSldNP1BYTGFKWFRhTFBgXUdYIUBlKFciZkJTPGV0IwpuLXBoQEo7SVU6Mls2MCpub2ZbbjFqKFomOGpbOSZFJnMyKEJKPTFOVGwzQiwpSC5jSUUqQiFLXEAwSWwub0BLZmE6Y21SZG1dai8sI105aAoqI11IaGBnTDhYJ2dnPSJTSTdGY1s1Z2YqMm8xX242ZD00VWZFaEtkWzdXJi5fNmMlLSQhZUsxW1ZoV10xUk4vY1U6ZDshYzwpKi1BV2dkcQojLGJdSkVoZHQqUyJiJGA4T1BAMDkxNkpaNk5uRi1tREIyZEIuLD8nKkxobj1MXzxROThNUFMzLj5ecT5wNTRObkNtSV1ZKTVhNjVHOTg5ZwpOMU1AaG1QRW1mNyFbVSlgUFdERVtcYztIajBWNF1vTls9QyQ5RT9tazheQHFXRk9zQWlOSF1AOFJaX1wuPmBwKCMrR1wuK0hlTTkiT2tmIwo3blFVaG9lJmdpLChpZiw8T1kpZmEoQjlvVSxZaCkrbkknN21CdC5sL1BSK1AlXVBVXV8zbFQ6UXVQLzIxbXM4P1JhN281Qy4mLlVtLTJfUgovUSFFNU1eJHVQX0prKVUwSzpBQFpfT2BuJVo+biJgPDtuUVZeOCktNDs0IlxZdXVuOiRyMEdwTD0sL00+QTc6RmooPDcmLDRHRWsuS1hUWQppdDgpLU9oa2Fnbmk2KW1FdGhpUVBaSmBtT3Nuc2tbT0BkXlNuSC5eQXQ1LzgvZTlraixcKjNdXGNOYCpcJj9lXGQyZXRFT28vbjQuOydASAohTC9QME9qIylGKTVsN2o3UFNTRzIzaj9PTERDTyFEWyosMCZfWVJAPFRrLD1CbysxKEVlUHAkTlNOZXIhTUNFc0JbMSRZW0BuRz1tPE0uYApoX21cRVJXRmlvQnV1RiohQiQxXzAvUW9pOVZVY2EsMDpzWSouUE9iPTdicEdgaSVxNz9AMDpiT1I8VEliamtBOGsnOSNoXUFEPF1IPjJjWQpxWmVca1UlKjVuZzxOUmlnV2YtOmdqSFRDaC1ZVS9xWTJKNFstbEVqcjhoWlheQEAlQTpQKzg8OGdzIUM+Pj5rcGw7SClTMF5fRVw1MTxBNwo+MGpXOjs7RHEtJDNgRidYJ0FEdEoqKGVDcjhBZVJgKzJXJU82YDxOcDo1JFciQi8pQDNwXGQ5OlVOTXU6OjY2Kio9QTtCU1NDNk1xL2g4UQpZZG81V0cnPjhZNTNhMWUzPTRCUjJyIUJZZGhSW2c/J29kOGNzSzFBQ0FbKlZTQF5GQCtnaiVZbSFVRGFeJXBXIlleY0ZmcE1vcnRUPixqTQptNU1CTWpsNFRcPUovWXJzNU9mcS8pJStTMVdwcDNvR0t1OU1yKixhKFVoSz5TSE80clFFXTZDUFYwQmREb1tSW0UiNVByOCs8bCklJ2JlcgoyUjMlbCRpU2dlSixTPm0rJFxDJEZPXzU1TFFhX09qM0VOLEY5SEg5aTciTlkya0hQWWczVStsbWJlJSlTWnArUklTXD1KTkpPSiFDaE1pJQpbYE9zLE5jME4sSGcmVE1FLzI5a0Y1NjgvNCUlYSdfS3MkIS1La08oOjZzcEs/bi5WTmhKLWZZLCxOcmMpKXJOPS1JPjw8N0RuRVo+bT1iYAozMU43UC9lOUdSbzZWV0MldGk+KlcvUz0yJnRnITReMGgwYzo6M0BETmVTRkI/Y14ncFtgLT8jRUIpTC1nL1gnRmBnTV8tWHJTOmE0PlRbOwo1XnFOIkNjK2kocWJFUmlLMl88KVouNDdVO1kqIVo6OFBEV2BqLlU7T1pZdTtKZyMobVBwKGY3aDdEWCcobj9VNyU/ajVEJSQlOSclQzgvUQo5VydqLGskOS4oWyhhaW0kY0toO0ZxRUpPTTFnckEkcjNXVidgJGQoMCw5L0JTMGFhRVZVQVhBVUYvUi82OyhBNWpbbCYhXzVbJmIrJichdAoxSkdcUUlUMG1dXGkuQmQiV0BsS2llYUghcFQ7L2I+LDxAYWlEX3BiMXMrU3RBNVY2XEBxZlBCYjE9cUVvVStgO1liMm5kK1VMNz46LEFGVApMNVZoTDZoPT85T0EkT1pZMS88RzxgK0VaNk1fPTAhc1lXdGFsQlkoUTw+VUFFSyxOYEZJKTtKPE0qKlM4KjohaiMsWzdbM05PIzJFXkdZIgo3NWRwdF9jMmgmbXIzY1JmT1xHP1FFamldVyVqUVxFWy5nNDBGXllnZmheO1VwJW9ZIW4vMkctQEw0OGJLYSV1I1piUz82PDZKWHMrOmdaRgo5M2tqMDFebSgwKk5aX142Vis9OldtWV1xUiM3Y2RwQ11HKCMsXDBjWDNyRylVZDZJNihtcjtFQylTVDw7WSYsLWRNI19wUiM3YCNGL01pYgo6NnBSW2JlUjEtbFNqPlVdUzk3IlAuWk9lLj5fYCZYalAuPyJkVS1gX15VSDVINldIYi4tMnIuU0soRFYra2YhJDI1NGdsWztEb0w2by5UZwpudEtjV2ohRyJHPSxOMThFJlxqSyU1M2UwcEcuRm49czQqSjtdTCxHSV00XC0oOlwzKThPUEAwQipzK1RhV18pXmdsQmhkK3BPc24qLiRiXQpnQWw7c14pIUNmP3NsIzIpT2lTUSRZZ1hLT3JaNTkqRFtRYkYwVzBHME00P1ZKYSFwdFhMbjwzJC8lcSElS2ZZWHBuO2VNbT9TMl0uRFJZdApRLkVoSiJXM2ltKjJUMUxJL1Nic1I7azgoN0QlP0k7XkJsUTdAWG1AO2VbaTNFLGV0bW9GZSFUTVNuWzhMT1FjRCE/ImE6MkUoaTxSbnQiIwpQYzg0Y0MiSiFUa3BBZjlDYyVnUVgwR3FKT1ojTz1TOFE1ZV0kU0BEOUpHUk1iNTEtRChLLjZbOTZ0KWUzbTE5QUE/WSgtY1RZQVk9Pk1xbQpONylmaE0hKDE2V29aZmZDXEYxUEBoLEUsOjBwYG4jYFlcdCIyWUpxL1BodEEtYT5jQjosYTg6Yzk2VFBnbEApaFBsXyM+O184USxrRm4hIwo7QyVGNmgpViFjQj4rVmlvPzkwLyV0c1hJU0QwVE5mVEVPKyRyNGkhJWJZXUJqQUJFRTczW0BfKnFcT2RqO0BIP1EndC9IVlBDLVY6JWZQVgpBWChLJyIjMmQjaFUoKVVFTix1TylDW0JBMGZZXitlYi8qcCk2I0FwWiEzVChdJzg5UWJpVm1aNGJHOlxCPycxMmE2XCxkUT8sZDdvaFghagpnbCwxJ3BcRChqU2UzJ2xpXmZKdF9sNC1lQ2BKRlZZL1BhRmlXMHBMO2VnMmVlKTlCZU5rUSY6MzFOYGJEU0RqYiZodT9sZEpBZEQmdEViIwoiRzFLMzM3U1lQMys/azY4PGwmKi1bbSUxYTBXM2giWDExaSwxOS49U0tbZDlLZHRHMWFITnRwJHI0OHE0cTszWW5GajtNLVZdRTVXViM/JQpmWFdGV1thJTYiW3BXU0ZiUHFVIVA8QW1+PgplbmRzdHJlYW0KZW5kb2JqCjIwIDAgb2JqCjw8IC9UeXBlIC9YT2JqZWN0Ci9TdWJ0eXBlIC9JbWFnZQovV2lkdGggMTE4Ci9IZWlnaHQgNDkKL0NvbG9yU3BhY2UgL0RldmljZUdyYXkKL0JpdHNQZXJDb21wb25lbnQgOAovTGVuZ3RoIDQ2MQovRmlsdGVyIFsvQVNDSUk4NURlY29kZSAvRmxhdGVEZWNvZGVdCj4+c3RyZWFtCkdiIi9lSkldUj8jWG5Ya1Q2QT5YS25EYldCIk1aRWFqUHBkNF0nKF5AUmx0ImtCIlMwJGYlPz1tPTRDOzsuJE4oalJjWVIlc2QmbDkkT3ReCikmNz1MKDctZXFxZUAuKC9TQzk4NC1mTFhlTFo1cWNaNWxBRUI1NT4zQDA9ayI7UmgyaEJjUUwzJko/NW9bY0gwYlNGTl48OztANDFTWTltCmMiWkk9WWwjYzQpYF8iWWxtTjMlLkEubE5TYFZlL1k8JT8uNygoKS8uZEJbO2shSlslJyVESlxJPythNzo1b05rWVxxRVM+MzxJNWJabShJCiVkPEBCKThEYnQzcEtHY2xiUWsxJ2laVnM3Tk0wXG8mUipxVT5nSHUtRnJDa2g8a0hGRypcJE09X3NWLy0kXEpGOmUvPz5mWEErQ0NtcERlCmRTUDZHY3UsXFgnQUw9YTVKQ1s0NkcuaHAyPVU0ZV4kZW9cUSNiKm03QzgkLUpBKGNaP1lgJ0M7NkpwSGVQbFwkSEVDIXVeKlFhKzZnRU1rCltmRVYqYi4jPnByPi1XU2YjPmgmVz91a2tUc2hZYD0vWC42NkEqO2Jpcy9kZ0BRWFhzcjhtO34+CmVuZHN0cmVhbQplbmRvYmoKMjEgMCBvYmoKPDwgL1R5cGUgL1hPYmplY3QKL1N1YnR5cGUgL0ltYWdlCi9XaWR0aCA1NAovSGVpZ2h0IDU0Ci9Db2xvclNwYWNlIC9EZXZpY2VHcmF5Ci9CaXRzUGVyQ29tcG9uZW50IDgKL0xlbmd0aCA3NwovRmlsdGVyIFsvQVNDSUk4NURlY29kZSAvRmxhdGVEZWNvZGVdCj4+c3RyZWFtCkdiIjBKZDBUZHEkajRvRl5VLCJIVHM5RUlFOzBBVCxfRSpMWiVvQDdKbDVWO0gnQ3M9VHJxRGFILjRCZiNjNE9WVDsoZCNmPEdFOX4+CmVuZHN0cmVhbQplbmRvYmoKMjIgMCBvYmoKPDwgL1R5cGUgL1hPYmplY3QKL1N1YnR5cGUgL0ltYWdlCi9XaWR0aCAyOTQKL0hlaWdodCA2NwovQ29sb3JTcGFjZSAvRGV2aWNlR3JheQovQml0c1BlckNvbXBvbmVudCA4Ci9MZW5ndGggMTM0NgovRmlsdGVyIFsvQVNDSUk4NURlY29kZSAvRmxhdGVEZWNvZGVdCj4+c3RyZWFtCkdiIi9jNiNtciIkajY+KjpHXUNiVSZpTVYscHQyWk5aSGAxNyQwX2pAZVcoYGx1VmN0SDJATDI+NEBGb3FSLGhIZCFPKnJUPUNMZklpaiUxCjouKmZeRTNkNyVIYGBSPE1QOz1daGNCUztoVlRlNGwoZDtfcj9EUCxFNDIhc2RkTCleYWwtNF9TWCEoS0ZzPDo5XTcjYUMuJGtGJUQrJ2JECm9pPCtnYXN1Xi5vbEYwdGdzSz9aQkdmWnUuU0ZKKWg1LyVZUSVkTDVRQFslRklSYyRnPGVxJlNhMyohUmwmJU1rLyJJOGFDcGEkTXFGLF5gCjJWPT5EO1MjQ2psNXBYL1FyKDJMWm1AUipGZE5TPFc9PT9jZSsqRGg2Ql06T2Q1Tl1tI3VsMF0jVHAlVz5XNyM6UGE5YFw5JChfYTw5MioiCkM5Y0FtLXNJQ29CMS9VbiJIQlcoUk8xQlk2RkxGUDlpZCNVN00qYmUrZShIOi8+XTFiTktRSHMoVnI3NVRNU2Z0OFBxOjwuY05qTy48KkZJCipIVGxCVXUhU0E9U3N0TCRRWWBsUEIoUE0qNUBgKGRSP2RCbz5HdVxgNDA4ZDlbL1xGTlI+YjEtPStdOCRfay9TO1MxRW0ydDBwMFEwNCwhCiZQRzZ1UHJiNVo7UyNLIl1XNEE+NDhFJ0hmZCY4KGZEbFdFbC0mN2AiRjA3MS42MylJKVRPc041XUlMKD46JyJjcE8iREFWPGFALik1OUNUCkskKFtJRmshISI6UCdNVFJLdDFIIml0ZCpbPUdddVo1PVFuMCs9bTNFS0YuJywnQVRFPV4rS2trcTQ4Ozk9SEorR15uP2FGJFw2VStBa0U3CkNgZFJWaHNwXDBqSCdrJG1VdS86NCEmY0ZWMUk5ZVFJOFtCODE/MywuLFs4Tl8rNDhENmU0JyRFMj04Jk9dKTZDXCJcY1Umbi8uW2dTMjxHClFRaCY9OWpeTGNrcVhRazhrZmdgZFUuRDg1V2tKTVZPVWNeUjI/Sk1RJz1ddWRWPVFvQ3MnU0xfXUU4YXJRcF9MUWRTZEM5JCg1R1xzPVNNCmE9b0xfImBlPDdOOShJXT1kXUItbjZSc1JVNEB1L1pYJzJBWjpvIlg2RFJESmkqJmBHa3J0cmFjJ0MlNDlOWUY3KTxAZV1LSDpXUUpHT1ZYCkRFYSgxMkUhUzQ8OWxJVVwpKzQoNyRjOCFxRUw0TzQ4TTJpMVEhRyM1KVEtZiI9M0hSO2NDaS1YVWQ6L1ZFLEQrMS9BcDxjK15aRyZuVSVRCiVbUE9AXGtHW3M1TCtCaUpYIyloVG07S1c7VDNVVCY0bU4wOjUnL2coRF5PPmpZIUMrRkksbSFPZDxdTy9Ka2k4K3E+ZzFFaltbPDE5YFdcCks5bCltP09ZLidCcDYvaz9CUnVnLj04Xzk2b1pKYlFkU2RDOSQrZzBjJURtQCVYM3FxbiUmNSIxJy9tP2loNzUiNFlcRyEuYis8YF1fOlI/CilGZGtFbEQzbChpNVJDR1YwJixnMWFTJDdoc088QTdtPEdOPilCNydpWTYoVlNNRVxGUlxPbWtWWm1OYmp0by01US9zLmRRSThbQjk4TnMvCnBYazNcZi5LQE4oWkhqKiQqYHRvISJATy5LRmcnKTEpYWUoa0w6KjZkcChrY29UbXAnXGtfTT1aczcoTlhlVGFMKl9JNGpMV3RURzttbXBNCiQ6ckpvRl8maFdpdS8oODlSbF1bcnNcckUvTF87YlY1PGciUGtPcUlvKjpdPzNwSH4+CmVuZHN0cmVhbQplbmRvYmoKeHJlZgowIDIzCjAwMDAwMDAwMDAgNjU1MzUgZiAKMDAwMDAwMDAwOSAwMDAwMCBuIAowMDAwMDAwMDU4IDAwMDAwIG4gCjAwMDAwMDAxMDQgMDAwMDAgbiAKMDAwMDAwMDE2MiAwMDAwMCBuIAowMDAwMDAwMjE0IDAwMDAwIG4gCjAwMDAwMDAzMTIgMDAwMDAgbiAKMDAwMDAwMDQxNSAwMDAwMCBuIAowMDAwMDAwNTIxIDAwMDAwIG4gCjAwMDAwMDA2MzEgMDAwMDAgbiAKMDAwMDAwMDcyNyAwMDAwMCBuIAowMDAwMDAwODI5IDAwMDAwIG4gCjAwMDAwMDA5MzQgMDAwMDAgbiAKMDAwMDAwMTA0MyAwMDAwMCBuIAowMDAwMDAxMTQ0IDAwMDAwIG4gCjAwMDAwMDEyNDQgMDAwMDAgbiAKMDAwMDAwMTM0NiAwMDAwMCBuIAowMDAwMDAxNDUyIDAwMDAwIG4gCjAwMDAwMDE2MjIgMDAwMDAgbiAKMDAwMDAwMjAwMSAwMDAwMCBuIAowMDAwMDA2MjUwIDAwMDAwIG4gCjAwMDAwMDY4OTcgMDAwMDAgbiAKMDAwMDAwNzE1OCAwMDAwMCBuIAp0cmFpbGVyCjw8Ci9JbmZvIDE3IDAgUgovU2l6ZSAyMwovUm9vdCAxIDAgUgo+PgpzdGFydHhyZWYKODY5MQolJUVPRgo="
    }, {
      "id" : "e5f580bd1cc3445b83b56b28a2200e46",
      "tracking_numbers" : [ "794684897588" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T11:21:41.547903Z",
      "updated_at" : "2025-03-12T11:21:49.235766Z",
      "succeed_at" : "2025-03-12T11:21:49.227545001Z",
      "meta" : {
        "source" : "tts"
      },
      "source" : "tts",
      "ship_from" : {
        "city" : "Centerville",
        "company_name" : "Postmen",
        "contact_name" : "Contact Name",
        "country" : "USA",
        "email" : "usps_discounted@test.com",
        "phone" : "18682306030",
        "postal_code" : "45458",
        "state" : "OH",
        "street1" : "8700 Meadowcreek Dr",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Santa Ana",
        "company_name" : "Maersk",
        "contact_name" : "Mister Jones",
        "country" : "USA",
        "email" : "maersk.test@example.com",
        "phone" : "8777285325",
        "postal_code" : "92705",
        "state" : "CA",
        "street1" : "1600 E. Saint Andrew Pl.",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : "202503120002",
      "return_shipment" : false,
      "files" : {
        "customs_declaration" : null,
        "invoice" : null,
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6"
        },
        "manifest" : null,
        "packing_slip" : null,
        "qr_code" : null
      },
      "box_type" : "custom",
      "references" : [ "order950000002" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "15eebae3-1646-4349-88f3-e8d0ec00e577",
      "service_type" : "fedex_express_saver",
      "rate" : {
        "booking_cut_off" : null,
        "charge_weight" : null,
        "delivery_date" : "2025-03-17",
        "detailed_charges" : null,
        "error_message" : null,
        "info_message" : "No rate quotes returned from carrier.",
        "pickup_deadline" : null,
        "service_name" : "FedEx Express Saver®",
        "service_type" : "fedex_express_saver",
        "shipper_account" : {
          "description" : "fedex compatible opp test",
          "id" : "15eebae3-1646-4349-88f3-e8d0ec00e577",
          "slug" : "fedex"
        },
        "total_charge" : null,
        "transit_time" : null
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.2
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "in",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "F231",
        "item_id" : "1313131223",
        "price" : {
          "amount" : 50,
          "currency" : "USD"
        },
        "quantity" : 3,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "method" : {
          "account_number" : "123456",
          "country" : "USA",
          "postal_code" : "123456",
          "type" : "account"
        },
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : "JVBERi0xLjQKMSAwIG9iago8PAovVHlwZSAvQ2F0YWxvZwovUGFnZXMgMyAwIFIKPj4KZW5kb2JqCjIgMCBvYmoKPDwKL1R5cGUgL091dGxpbmVzCi9Db3VudCAwCj4+CmVuZG9iagozIDAgb2JqCjw8Ci9UeXBlIC9QYWdlcwovQ291bnQgMQovS2lkcyBbMTggMCBSXQo+PgplbmRvYmoKNCAwIG9iagpbL1BERiAvVGV4dCAvSW1hZ2VCIC9JbWFnZUMgL0ltYWdlSV0KZW5kb2JqCjUgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9IZWx2ZXRpY2EtQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjcgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhLU9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iago4IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0hlbHZldGljYS1Cb2xkT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjkgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvQ291cmllcgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEwIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjExIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEyIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZE9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxMyAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Sb21hbgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjE0IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL1RpbWVzLUJvbGQKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNSAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1JdGFsaWMKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Cb2xkSXRhbGljCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKMTcgMCBvYmogCjw8Ci9DcmVhdGlvbkRhdGUgKEQ6MjAwMykKL1Byb2R1Y2VyIChGZWRFeCBTZXJ2aWNlcykKL1RpdGxlIChGZWRFeCBTaGlwcGluZyBMYWJlbCkNL0NyZWF0b3IgKEZlZEV4IEN1c3RvbWVyIEF1dG9tYXRpb24pDS9BdXRob3IgKENMUyBWZXJzaW9uIDUxMjAxMzUpCj4+CmVuZG9iagoxOCAwIG9iago8PAovVHlwZSAvUGFnZQ0vUGFyZW50IDMgMCBSCi9SZXNvdXJjZXMgPDwgL1Byb2NTZXQgNCAwIFIgCiAvRm9udCA8PCAvRjEgNSAwIFIgCi9GMiA2IDAgUiAKL0YzIDcgMCBSIAovRjQgOCAwIFIgCi9GNSA5IDAgUiAKL0Y2IDEwIDAgUiAKL0Y3IDExIDAgUiAKL0Y4IDEyIDAgUiAKL0Y5IDEzIDAgUiAKL0YxMCAxNCAwIFIgCi9GMTEgMTUgMCBSIAovRjEyIDE2IDAgUiAKID4+Ci9YT2JqZWN0IDw8IC9GZWRFeEV4cHJlc3MgMjAgMCBSCi9FeHByZXNzRSAyMSAwIFIKL2JhcmNvZGUwIDIyIDAgUgo+Pgo+PgovTWVkaWFCb3ggWzAgMCAyODggNDMyXQovVHJpbUJveFswIDAgMjg4IDQzMl0KL0NvbnRlbnRzIDE5IDAgUgovUm90YXRlIDA+PgplbmRvYmoKMTkgMCBvYmoKPDwgL0xlbmd0aCA0MTY2Ci9GaWx0ZXIgWy9BU0NJSTg1RGVjb2RlIC9GbGF0ZURlY29kZV0gCj4+CnN0cmVhbQpHYXQ9L2JBTGsiJmAhaWE1N2ttXypiPV4uPElBSzMsSU1DQTtHMGxyNF5OQXNmXSxDM19oaiFrSzcmKFdhZj9KXj1saU1bKFA2VydnY0F1Tgo9OloyJS9cdFk9Pl1xUixPTCM4KDAlVGdhKDBnOmBrM09rNkJJb2FEaHFrIT00YGYoNl47VSJHPnInLTZwbE5iSlY4azU4NVUsJlcuOmdUOwo3aFo0aDNbQ1JJbWQ5ISJibm80TkY4NVBeUDwmUjBiVyNcPXJvUUJnZ1NAWFNoP0RxS0hpKjxVUkpHbiVyZ1EtIkllX1tOcEAnM0FnI3BJcgpvZy1TUmhOLCVuZUUxJ15gQ2U1YiMvTFhjbGw5OTFAZyZjYT5bZjZwVTVvaVJuJVpIRjFSRlRLcD1xV1daYW9pJjE1UlZKOERJR2loTDBaLwoyaCI+L1kyP3JlU0ZtRVAwIUIwZGtCUUw0Z0NNLGhYNS1eKlIxbmR0I2BQYS89QXQqc2Y7JEBSW3BQWXBCKmJiOGxLYTBuVURjTE9OOHBTNgpHMkt1MDdcTCYpR0lQbTI2I0tcX286QVFzTyUhO2ttZ2M2XjMiPUtZYSk2SjQ0IUBRMlovS0I5ZW0jVkBebjBiXjVXU1hGOz9hNTFELl0tXwprPj9WJDInbnRlZ1BnRiJmNipHMSxLY28tX1JcKjwsJTM5SjBNMHBjTm81K24/NURnZCw4LzNDViFLMlhycEJbN1FXSDdpRmE4SC1tV1MxMQo0bSdEUVZGV1xBKGEvQ1MwTlYwQl8nVSxmVlplZm1xUkQpL0RLYz4lLC4kJVpDZ1NJMWVyTV90NU1VKCEwOV5GYHA2KU84MnF1VT1tQ1M7UgptSWhLP0gyZU02VSFpUEcoQHVpJDdoJzxaTFtHJWE9NGMnWTVGTDkwaSQnRDJRY1g9WzE7RFhOKk9fO0hJdDAzLVJyQSRTb05PaTk/WWhAQApicU5jLGsyJCU2V3JPYCM+LFlnKTo6OHFlUFNnLF5GU2x0cURHVDdcaS5iNnRlM3QnJFlmRD8oQEJMK2JEUWBcMFNiYCwqXCkiZXEoa25gaApjMVpWZGJtVkxEailSMi5CPTxgT0RtIjZnWmU4WEBEcjZ0cGxmRWhQckdkNGc/PT5EcmsnSFdCSFltKVliP008YjRkWyNVP0cxWllFLXAqcQpJWTJLcGQ7STFJYGAvU3QxRXBJOyJWUWJAX1tpWmNKYSFEVVsvQGkpLU4hQSZfQCRKOVVMIUYrcW5xZDBdP0JUWkcrSSpcMCZjUVJXITxjVwpAMCFSMT1HOi5mWzNpaSlfIWdqN2RCMStaXGRAb1pKMm9MTGkyRjM5QTpjblMjcT9gdUxHcmFsSVdMVyhpbFI+cTFTPjwiR1ssblNLPCMjXQotNislL1MlQ11bPWlMM1VTWl9PNTU8SklQJDM+KShDRSRHNjNWTUFGLzBSMSQiVWInWmBWQy9hUkQ0NXBbZj5ePFNLYEowYiJNL2BUbTpSXQpmaCE8VEU9dSxOX25HWGY3ZygqMClcTFRYN1ZxYVI8RmNLYDFXKElWPz8lLUpGOz06STk+bmowQmRUM0JEK01RY1QtU0ZYVV9GIUMoIUZNUQpQWEpYMzUjQW1xQFlJMjNHcFxuUkQ2WVhnTzBeJWEmbE1HVEFxSzhcaTRKa3JtLjZBXC03WUI/SzJXRFpqYlhFUkFoXyxFWC07bzBIPE0xUgo6PDpaX1JNaHV1bk9yKGBLbVxcNkFtb2dvUXBfVDg6WV0oI2luQD10RV8rZ3NkVVhcbERES2xVO2ckVmtqUGlVcWJSRHMnNWMudGloR1VJIQpPW00iQUZEX1BdLmxEMGduJCdJVSRWSG1zXSVJYyNRRFI6YCFNUiZNXzBhSF1oLyFZOUk6TiMpalwhNloxX1dQcSFpVFkvJWw4UVE1ImErQQpMKy1pQFtVJi1PIXJRJnMhbUxqSEQmXC8vOy5iU2FuPFlWWUMnTFQqTmpIKz5XO2lUNygyZEcmUlcoO3RfYFBhczZBVlNsT14oSUcnMkQtWwpnb3NrbVRBXGppP3VISEs5OCMxajA7T0NcLUAzTWYobj9jMWdZai9NZz07UT04XUVJRVBmQTUvVVw1TiNKT0lbPCJlbD0kLS1ELV1NaT9YLgpNJVlMbiZQWTZVbXFJVyEycW8vaGFxPy9QUjo0NVE2TCowSEVKdDJBM186WVdfM0FcJWsxOjoiI1ksSkFRKyMkQmUnYG1dbzQqZl9BZ2JcLwpBdDUvODA8KHAsPFlwV2Y7TU45TlBqOlxTZTUwW1NGJ2hLX1JXMkxUbEVxYClmS2AkTldAIWpxVWxIWmdKayhZZyFlcFEmXyd0ND0pInJpSwpTOD1LKGxQWlJyPWpKMz5nUGo5JTlBPSshLj5kVkFqR0peQVtfa2dqQlo9YkNCXiVmJ2hcRlouLDBxXTM5MmluJD9jJTwzPGcpOCVOKUw/KAo1N0JGUkhwSyNIXGQ+NzE9VzoiMV0oSVtJVz9Rc24mbGxEJ1ZwXTByPiwhdU0yQmhXPTA7YF8mOU04PShkNmpxcE9gRydVJjglR21hQlFVOQpWWVEvaiE3ZGErTmZba0NRNyozKzdATlksQDV1KDsobj9jMWZobmdvTDtjRTZAS09yOWRoOiZRU2RBJGloXmoiaCMsYGFzX2U9LGIvW19eNApFQUtSI09uXGcxJ29vZiRrXyMqUDEscGpfKmFOXSNOZTV0UE1RNy5cOEpMTUY+OXFRJWNhLi1qMUpVUldRUj46PU90azA/PD86PS5DXT5oRgpjU09VRzgmISVKcWNSZ1krZ2ZzLCxhcltVLC01ZFJCSVYmLT0lRmBBKTRZYChhSCFSLlU4R15yPitsOCJYJihwVEVMTClwJmNjSHAnJFZrIQpXQWBnaFs2Ol50Ti0+JUBkQEEoPDhrWFcnOCg6V1peSkpqbjBTdVljN1pwKDgzO2hRLDQuYGxxSHMzKl1cQDknUmwoS25CREpqW1xbO2ptVApEIWhiJz5DNitTZ1MsOHVPIkZBXDVNUFlxbV5OKSxmbCtVWmhFT1hdJEIpVkokMy5RaUhUNDMvOWROPiRYRyFsO2cyYHFHT2Q7R1pKRlJrYAo7bmpyZSU1KENUMm9WOlI2LS10b1FYXkVNPlxtXzJtMWNJcTY6UW9wbF5WX2BcWGRPPjtpYDBXbWwhc0c+OVBgJ1hnRUJwQ2EnUm5YWUtHOwopVVgiVEYqRztcJkpDXlkpNiJVMCo4WyssX2NPbSomSD1qS1hVRUtfNnI3WHMtNkghbERhOlZgYFk+YFEvYzU7KWEnWlhdcmhvLj1tPiQyaApwQVNHITE3LGJJck9FNEJaalEuISoxXzRAW0ZNSG4pMGhgOjZWKSs/NCQyYEU8XVNwMCtFKCpoZWpsQFxgVig/cz5VWVRWSkxUXmVTcCQpOgpbTHEoO1tONTkoVG4hJGFEYTQ2dGRzRjAwWHI4T242XmRPJDAhNzAjNi1wKUk1Mk03bmc2OGpGPkU6OFRyVkpoIjJKWDosSSJHQzhLSyFgVwpTZTlRLi47LUxaXidROldrdF5YZSpYN3MsRUZaYzwpOztZLWxSYEAvVEY5Nlc9Y1M4Z1dHJ0cjW0FjVFFQW1peZGkrTnNwNVUyZzglQkZ1IwosRGJJP0R1cFdgNEBnX24wOEskWiRyL2FUMT1AZEpKNiZLQ0pmU2VpZUgxIlhsPUNaWStVQD5lI3EmNl1pRF9VSnBGWjRbQi4/SE4+MCEjbApSOjJOZm8nOjonYj5CYkdHTENSTVM7TzUnUE86cjRSY2UqdS5eXUY9UUhjZ2w8IjEsWkw7O2JJTChNNERvbzpEL0Umc15NOUdUb0wzMjFSTwpianJ0M1NgM0o0R0QxXnViISUxKCYtTEQ2NiQxUk0nJkh0TzMwaSdWakgnMk88OGIzKlE/WyMvOG8kL0BPclRTaydqNTYwLiFaajxWcFxkVAo2NkJpK0ApT0UhMWMlaEM2QE5IVlkuP08tTzlITzRWckVRLm0uVypITEtZdU9iMkg7XkguOXBsXy4/Si5VZ24rTlIqKnVWWiNsVyJPZyNySQpKZnRGXyRyL1VMJ1pMbllRQT1KJSdIcmkhME86W2xmPU0rc0AiXjBWMmRnI1NSUU1YPEM2PTQlaTdBZzU3VkNCYTt0VUhCSCN1ZWhAOjFRTApPTWNCMytYX3JhUT1UU0QxSkJab29jaixYUCRKOWhAT1lmRV47YSxVaiRtRkVAMmJGWVo4XklrKlctZ0lmSjxxMGA9YllNJDRSbiVBcT1OSgpXS3Q7cDVvdERrMm5RYFxCR2QsPjNgclQyTDYsU1E8Mko5bDBHUjxRWTllLUUkcXMxTz0vQWoyODshLTYwSkpJaWUyMTJXVXVcKDhWQy9VbAowR1I1JEQ1VTNcOjZsJTdiZVIxLWxkcFlLXVM5NyJQJy46aTtcSUpLPkdJL10kUzYtNyVdOlNwbFwsSXRVNk1eP0dCLFpLS2tUZCtmOC1MUwplOTdjNyZDITE3YWJ1VEpOZittdT5xJkNYTFZyQSRCXkc8QklQYXVaVV9SP1gzbFJlKW1XV3QkQD1kL0ZQIWguZUZKZSFKUk9mV3BjYE5aYAomMj8vMGsibWwuZGZ1MVRvc0cmUyE8XmZVY29AaTteN11aLTAlZzhYZSsmWkIzWXVULypNWDteKG1vMlxfS3I0WE9IP1NEazplUzlOPzpiPwpmLlBgUVBvTGNyO2VbaFJoTEZLZTovVDBVYmRVcihSTWkkWFUqdUIqRixHV3VdU1EvZkUlN3RZKUkiMkkiTUk0JlRdMG9PYyhTSUg2VTZBQgpRSGJFKmtwQkA2NUBuUyw7RUZML0U0UUA8KVs/aG5xZUJxKzdca1tgKTVmUThTYVErPTFfV1BDU1svNV1hYjxdL29USiZKU0JlRigqUlFCXwpuJEEnVEVuMVhvMmtyMzgqOSk1alF0Z29wPiswbSFvVSpORlZqI19NKy8uJ1JVLWs7Vl0maCcqbTNKIVFfPEwqN2IuMzsyUFpxQSU+MC5wQQozN1MpKEZyZ3U+LjdHSEZmVV1xRUdpaWQ8NE84c2JuTl5kXmA0OTBxTClpaG5gOEdgIlVXSURbLEVFZ0M5Ry81NVpJK2xmYzgiQWIuaSIkUAo5TGklNU1zTkgnYiIvKSIlRS4xKEg6KWhFYGdzOCkuJj1jXWA4R19WVzMzbm9iUXNLXU1VZlFQZG5NM05BMl5APm1NVFBUM1xzT18qT0JPMgovYSZSTVguUTFHZ2wsMipxdHJFT2NnWnVERU1RLkdpRF43QSliMidCQ1tuIWpmSD1bQzkkT0pMKVJhcStVK15gbyVdQGRHUnI+NzE2anVnXgpbL1IvNCJRcnNcNWQ1KThWY2s4XExHKDdcVHNBOjBqMl4sIzREVW9zaSQoZkYsMF44PikxbDNnVnVqKVI0QSJrVmksbjtGajROXSVdPWY2Rgo4Sj5RUC5fWVg3OjJXW1xYbiEmIigoPGxFcSoxc140V2F+PgplbmRzdHJlYW0KZW5kb2JqCjIwIDAgb2JqCjw8IC9UeXBlIC9YT2JqZWN0Ci9TdWJ0eXBlIC9JbWFnZQovV2lkdGggMTE4Ci9IZWlnaHQgNDkKL0NvbG9yU3BhY2UgL0RldmljZUdyYXkKL0JpdHNQZXJDb21wb25lbnQgOAovTGVuZ3RoIDQ2MQovRmlsdGVyIFsvQVNDSUk4NURlY29kZSAvRmxhdGVEZWNvZGVdCj4+c3RyZWFtCkdiIi9lSkldUj8jWG5Ya1Q2QT5YS25EYldCIk1aRWFqUHBkNF0nKF5AUmx0ImtCIlMwJGYlPz1tPTRDOzsuJE4oalJjWVIlc2QmbDkkT3ReCikmNz1MKDctZXFxZUAuKC9TQzk4NC1mTFhlTFo1cWNaNWxBRUI1NT4zQDA9ayI7UmgyaEJjUUwzJko/NW9bY0gwYlNGTl48OztANDFTWTltCmMiWkk9WWwjYzQpYF8iWWxtTjMlLkEubE5TYFZlL1k8JT8uNygoKS8uZEJbO2shSlslJyVESlxJPythNzo1b05rWVxxRVM+MzxJNWJabShJCiVkPEBCKThEYnQzcEtHY2xiUWsxJ2laVnM3Tk0wXG8mUipxVT5nSHUtRnJDa2g8a0hGRypcJE09X3NWLy0kXEpGOmUvPz5mWEErQ0NtcERlCmRTUDZHY3UsXFgnQUw9YTVKQ1s0NkcuaHAyPVU0ZV4kZW9cUSNiKm03QzgkLUpBKGNaP1lgJ0M7NkpwSGVQbFwkSEVDIXVeKlFhKzZnRU1rCltmRVYqYi4jPnByPi1XU2YjPmgmVz91a2tUc2hZYD0vWC42NkEqO2Jpcy9kZ0BRWFhzcjhtO34+CmVuZHN0cmVhbQplbmRvYmoKMjEgMCBvYmoKPDwgL1R5cGUgL1hPYmplY3QKL1N1YnR5cGUgL0ltYWdlCi9XaWR0aCA1NAovSGVpZ2h0IDU0Ci9Db2xvclNwYWNlIC9EZXZpY2VHcmF5Ci9CaXRzUGVyQ29tcG9uZW50IDgKL0xlbmd0aCA3NwovRmlsdGVyIFsvQVNDSUk4NURlY29kZSAvRmxhdGVEZWNvZGVdCj4+c3RyZWFtCkdiIjBKZDBUZHEkajRvRl5VLCJIVHM5RUlFOzBBVCxfRSpMWiVvQDdKbDVWO0gnQ3M9VHJxRGFILjRCZiNjNE9WVDsoZCNmPEdFOX4+CmVuZHN0cmVhbQplbmRvYmoKMjIgMCBvYmoKPDwgL1R5cGUgL1hPYmplY3QKL1N1YnR5cGUgL0ltYWdlCi9XaWR0aCAyOTQKL0hlaWdodCA2NwovQ29sb3JTcGFjZSAvRGV2aWNlR3JheQovQml0c1BlckNvbXBvbmVudCA4Ci9MZW5ndGggMTM1MwovRmlsdGVyIFsvQVNDSUk4NURlY29kZSAvRmxhdGVEZWNvZGVdCj4+c3RyZWFtCkdiIi9jZ0pUOk8kcSdeZjVIUUxTbF1SaGdwNS86dUw0XGozWiktSzxhYU5tOWtON0RCcSJqZWlhR1Y4O1tJQ1N0R0wjKzdcRjdbaSsrJjcvCkVTUGhZWkc0T3BOUU1POmtYYSdFUGVOMUw7YWkhXFBXWWJWalkuMCs+PXFRUUc6RGU7RS5bYy5hc19adU9qIT5aXi40TnM1OStpRWVPOmBDCjpFS204Kz9IM0ZjQUU8dTo3QjIsMTxZJTE9QHU2TitzZ1NxZlYqcywpZUZlaixILldiMDFsWD9rdWRiSzZRZHIzXiwxUVNwXWleZHFATktoCkhBMGYiUUk4W0I4NGkoZyZrQT1uOys7Qi9bL0wnTT9sNS5VbTgjKGJmcD4zXixgKjBxW0xBUzRoaD0+alM7S2tubWE0cDk5KHNfa1QrUV51CkRsXW8iNGtdc0g7OnFIUFNPKnArNytBX1wxPzpwR1E3U29UbEVHbW1hTkwvQjBsY2JkOGtcQ2VWMExvJ180MnFTLmc9PDcyJ1tYLyctMGwiCjtgOGJMQEByXXAkZEtAU2NOcUJLSG1WT0tnbD11KDtIY1svLkhTaTo1MyZCV1BySF80KiwzaFo4dDQ4KS8lS0BbJ2kudWFMNzQ4MCpcNC9aClQhOjlnTS5hOyRhJjYpNT0iLDwpLk87S2VPMU5PWUA1QVsnSjFBM1tqaCojW0kuWSQiNGJcaiVOYi05S2BES2BUUHJzUUVvOy9nWVkjRUpSClZ1ajEyPTRVQms7MHViLSNyKD9yJlkoWm45U1JbSCFaNjM9KGQlKGoiO11CRVM1YFpBMitrJFEvWiM6Y1FdYS10W2ssTlloZVw6XytyT0UuCkpLZWBDRGY3XF9IQEhyLSNnNlRPPW9JUUUoWi8vKitcbm8yWyQzMUBsQUpnUXIqSyFRL2M6XFJsXl80aD4pTDFlO1pKVl9XXzMlSSwmIzwyCjxIZltrNzJCU3JJNDZiRWJFVGxEMipEQD41Qm1nRkdcXj9JRW5ZUVplOjllYTtPNk9GN14qNCo4MzVmOSpzcFdPZ2k0QllRSThbQjloP0JfCi09ImojNyNKRzM5OmUiQTVpJGtcW2M0MkUmNmkmMUg/Q0prKSouLyhaTCQoLWtnJGBYWF5pazRWOCZdPGJNJUYjamFjVTRRSmwqLCtuVF5yCk0ibkxmbCZWVDNONyFXTC5JK0FUO2ZDJXEsImxXIXI9KCg/Tl1zNSUpKz9UWjUyYDROIVttMGlhWzwqYzgiXD5dKTZhZ2sxL0FxaVJmLiUrClYzUj4jMyo4ZEdrcVhQanJBcXRXQFQlZzBMPz49JSRLMGxyJydVXCgzdShqPS9oRyhbYj9bXzVrcThnJCwqWTdPUUo2RSJvWUdcXCImSjEiCiU1XUBJX0ZHKnJDc1lyQmUnc11tZGw8VCQ6MDQ0V2BedU5GO1MjRktubDg8NkdiZGZRQlVIcitHY0pma2pFYF11OSQuOUBtbyo+KWRbQWZSCnEjNFRGLUc+ak04c2R0KFxrYFhPTC9qbiY2MitfRSI3S2VlRiR1ODokcFxiLSNnLz5UOVhsKGJGc1RqcFtOX1M7cEA6MFZqTCk0VlFJOVZCClx0WFZILTA5cSpyJUNISUtGY1YpKj8jLSYxNDNdaSNGMUtpZSErNl47IXViMThRZmZuOUdeQE1WNmxLVDtJVloxTCI+cUFlTl4saXFxY0JmClI/TThYLUd0ZzBVWyUyYCVzK10pM2g8NEJiZlEqTktrOEEsbHIlIlpeSXI6NXMkLU1gTzxsK1F+PgplbmRzdHJlYW0KZW5kb2JqCnhyZWYKMCAyMwowMDAwMDAwMDAwIDY1NTM1IGYgCjAwMDAwMDAwMDkgMDAwMDAgbiAKMDAwMDAwMDA1OCAwMDAwMCBuIAowMDAwMDAwMTA0IDAwMDAwIG4gCjAwMDAwMDAxNjIgMDAwMDAgbiAKMDAwMDAwMDIxNCAwMDAwMCBuIAowMDAwMDAwMzEyIDAwMDAwIG4gCjAwMDAwMDA0MTUgMDAwMDAgbiAKMDAwMDAwMDUyMSAwMDAwMCBuIAowMDAwMDAwNjMxIDAwMDAwIG4gCjAwMDAwMDA3MjcgMDAwMDAgbiAKMDAwMDAwMDgyOSAwMDAwMCBuIAowMDAwMDAwOTM0IDAwMDAwIG4gCjAwMDAwMDEwNDMgMDAwMDAgbiAKMDAwMDAwMTE0NCAwMDAwMCBuIAowMDAwMDAxMjQ0IDAwMDAwIG4gCjAwMDAwMDEzNDYgMDAwMDAgbiAKMDAwMDAwMTQ1MiAwMDAwMCBuIAowMDAwMDAxNjIyIDAwMDAwIG4gCjAwMDAwMDIwMDEgMDAwMDAgbiAKMDAwMDAwNjI1OSAwMDAwMCBuIAowMDAwMDA2OTA2IDAwMDAwIG4gCjAwMDAwMDcxNjcgMDAwMDAgbiAKdHJhaWxlcgo8PAovSW5mbyAxNyAwIFIKL1NpemUgMjMKL1Jvb3QgMSAwIFIKPj4Kc3RhcnR4cmVmCjg3MDcKJSVFT0YK"
    }, {
      "id" : "8273531b6c0b463f8986faaf5aed8944",
      "tracking_numbers" : [ "794684897350" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T11:20:46.997439Z",
      "updated_at" : "2025-03-12T11:20:57.383291Z",
      "succeed_at" : "2025-03-12T11:20:57.345190392Z",
      "meta" : {
        "source" : "tts"
      },
      "source" : "tts",
      "ship_from" : {
        "city" : "Centerville",
        "company_name" : "Postmen",
        "contact_name" : "Contact Name",
        "country" : "USA",
        "email" : "usps_discounted@test.com",
        "phone" : "18682306030",
        "postal_code" : "45458",
        "state" : "OH",
        "street1" : "8700 Meadowcreek Dr",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Santa Ana",
        "company_name" : "Maersk",
        "contact_name" : "Mister Jones",
        "country" : "USA",
        "email" : "maersk.test@example.com",
        "phone" : "8777285325",
        "postal_code" : "92705",
        "state" : "CA",
        "street1" : "1600 E. Saint Andrew Pl.",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : "202503120001",
      "return_shipment" : false,
      "files" : {
        "customs_declaration" : null,
        "invoice" : null,
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6"
        },
        "manifest" : null,
        "packing_slip" : null,
        "qr_code" : null
      },
      "box_type" : "custom",
      "references" : [ "order950000002" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "15eebae3-1646-4349-88f3-e8d0ec00e577",
      "service_type" : "fedex_2_day",
      "rate" : {
        "booking_cut_off" : null,
        "charge_weight" : null,
        "delivery_date" : "2025-03-14",
        "detailed_charges" : null,
        "error_message" : null,
        "info_message" : "No rate quotes returned from carrier.",
        "pickup_deadline" : null,
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "fedex compatible opp test",
          "id" : "15eebae3-1646-4349-88f3-e8d0ec00e577",
          "slug" : "fedex"
        },
        "total_charge" : null,
        "transit_time" : null
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.2
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "in",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "F231",
        "item_id" : "1313131223",
        "price" : {
          "amount" : 50,
          "currency" : "USD"
        },
        "quantity" : 3,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "method" : {
          "account_number" : "123456",
          "country" : "USA",
          "postal_code" : "123456",
          "type" : "account"
        },
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : "JVBERi0xLjQKMSAwIG9iago8PAovVHlwZSAvQ2F0YWxvZwovUGFnZXMgMyAwIFIKPj4KZW5kb2JqCjIgMCBvYmoKPDwKL1R5cGUgL091dGxpbmVzCi9Db3VudCAwCj4+CmVuZG9iagozIDAgb2JqCjw8Ci9UeXBlIC9QYWdlcwovQ291bnQgMQovS2lkcyBbMTggMCBSXQo+PgplbmRvYmoKNCAwIG9iagpbL1BERiAvVGV4dCAvSW1hZ2VCIC9JbWFnZUMgL0ltYWdlSV0KZW5kb2JqCjUgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9IZWx2ZXRpY2EtQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjcgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhLU9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iago4IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0hlbHZldGljYS1Cb2xkT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjkgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvQ291cmllcgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEwIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjExIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEyIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZE9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxMyAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Sb21hbgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjE0IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL1RpbWVzLUJvbGQKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNSAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1JdGFsaWMKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Cb2xkSXRhbGljCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKMTcgMCBvYmogCjw8Ci9DcmVhdGlvbkRhdGUgKEQ6MjAwMykKL1Byb2R1Y2VyIChGZWRFeCBTZXJ2aWNlcykKL1RpdGxlIChGZWRFeCBTaGlwcGluZyBMYWJlbCkNL0NyZWF0b3IgKEZlZEV4IEN1c3RvbWVyIEF1dG9tYXRpb24pDS9BdXRob3IgKENMUyBWZXJzaW9uIDUxMjAxMzUpCj4+CmVuZG9iagoxOCAwIG9iago8PAovVHlwZSAvUGFnZQ0vUGFyZW50IDMgMCBSCi9SZXNvdXJjZXMgPDwgL1Byb2NTZXQgNCAwIFIgCiAvRm9udCA8PCAvRjEgNSAwIFIgCi9GMiA2IDAgUiAKL0YzIDcgMCBSIAovRjQgOCAwIFIgCi9GNSA5IDAgUiAKL0Y2IDEwIDAgUiAKL0Y3IDExIDAgUiAKL0Y4IDEyIDAgUiAKL0Y5IDEzIDAgUiAKL0YxMCAxNCAwIFIgCi9GMTEgMTUgMCBSIAovRjEyIDE2IDAgUiAKID4+Ci9YT2JqZWN0IDw8IC9GZWRFeEV4cHJlc3MgMjAgMCBSCi9FeHByZXNzRSAyMSAwIFIKL2JhcmNvZGUwIDIyIDAgUgo+Pgo+PgovTWVkaWFCb3ggWzAgMCAyODggNDMyXQovVHJpbUJveFswIDAgMjg4IDQzMl0KL0NvbnRlbnRzIDE5IDAgUgovUm90YXRlIDA+PgplbmRvYmoKMTkgMCBvYmoKPDwgL0xlbmd0aCA0MTUwCi9GaWx0ZXIgWy9BU0NJSTg1RGVjb2RlIC9GbGF0ZURlY29kZV0gCj4+CnN0cmVhbQpHYXQ9L2JBTFMqJmAiLGk1O1xIci80PFM8V2RcVj9pZiVhLmpBPWZvWFRcLD9aL25DY19oaiFrNkc+WF1dLlpcOEZIc2kpLzBuXExET29QNApwJD9HKS0kWT8vYjpZKFdFXSwpbWA+WUIpQExLb1ozUC5zOylOL09ZWU89LiQ+ZScjOT9YZU1eL2ZERURwbE5iSlYyI1ItNVUpZGwuOmdUOwo3aFo0aDNbQ1I5bWBrOl1lSkknVkY4WThIUDwmWDJfRGhXM3JsdEQkQ20uMT9HU2M8c21ITlpXKFp0SjRdNU4vbXE9L0E+XUFwUjFCOmdOZwplVSYwMEdXc1BSO2Q0VS4nSXRKJkdSN0UkQShUWDoqSC4kM2E8ZWhtLU1odWVyOEZdV1txPDY0bVRwanFEYDJnMlNnVnFJSFxuNjhqMT90KgpZRGxtJVJIbGZYJzUuOVdcRiVoa2lBSkBFaTljOy1jLjJmdV5zISssMV4oZnQ0ZmcpJGo2SzlIR0lvTTdSalNLSDlCdUQvOXFyVW9KckovTwo/MSNDUlZNbWs5MT5ea084SGNTKDNSMnQkZU5GVk9oN1l0QyRmTnBJKEVeNnEoM2kwRSgxUDInTjMwUiwnYGpyalQwUEY0PlooYS4xY01hNgpsMEJiM1ZPQS5GIiNqXjJub2FXMkg6JUVhNm9cODotUERyM0sxV3EvbypwLllqXmssLDMkUTBJNkhdOyRtR0hjK0k8aTRaPTUyPWNtcXQnWwpjc3Q/b0M1UEQ1TT8rbkMkP0lMLispSjchSVRkPFM9cyJkS25vZk05JzpKZ11CKG1pUltYUFRnX3NxP0pKIkRtZ1s/a14zXF8xaS5cKmtzPApHLlZcOnA/anBSNmAkXE4vYHVTJDopX2pSTzFXJG1ZLStzUElrXD89K1tlZExjXWRWXUFVXzgrM3BzSD1FJ3U2PlF1PE1dTipAW1s/Wio8VQpicU5jLGxKOkRwPHJsTSZaciJNWjpVVCc8L2IrXS5nJHJSbFxOUytQSlU9TilVYjhCPEpOWGRdKD1tZGJpb11NK2VYPFVacFVlPSFLaSdmOQpnWys+PGhgT2gyMyxQUzJuJyNaKExYIkJLQksjYlRoPHBoWkhiRE8/WVdvaFlNXGskaFVVZjZvZmoxKTZBcFw9ZVtuIm5GaE1XMnInUS1eOAo0RGByUUdoXDRnZ21JQjVJNyEodFgkRDlhQUFtSSEwSlJ0ZDgnQmU+RFIob0suNktCdTVPYWdwbi4wVG1tQzBdU0dGY3BVMCZhOl9XLVx1IQppPDxbRFltUzxWQ0pEbHFKYlNfamZyW0Y3WjNnJj9ebWw8YUU3QDNoMTtUTnUiSTcwNl9eaSNWXlVkKypuU0hDKUgoYmVpcW1abzsjJDFvVAo5SzMzXVtXWmBnQiwwU0dHSjs5I3JUcUQpLG05cnJZTzBsN2FnITRrP2pZcS8vbDlYOC9DJ20oYVNjIkNNbzAwRG1NdElROy1uO0MpT2dDOQpjOyJHZ04wQig3UmRRXyIzP2ZYQmdrX1UtMmZicmhYcy4ocFNmJm0mbjs3VGZQY0s+VD1zPXJFVmpVcFJbQG4jVEk/Q0gzOmY4Ilc8XCtjNgo7XEVNa3Foc0ldJmg9Iz1jU05DWkM1WlhQckIyWzovWVVaLkJJOSssIy1rZGklNmxOIUBBLzlVMUg2W1g9cnVAaT9qOERcOCFAYSI+anVzOwpjRltuSG9lOEBeLXIqKDZuXihxcT5zXF9xXlloNDE5WSRcKUtXYiNYSURlVzBsNT9KKmwwTzIxXG5XYXFuS0poKSUrQFk4UzxoT1NdZShrZAouJ107O0Y9S1dabC5rLlk4P09pRy9CXEgjZE5HZ2VbdT1FMDxBLDMpJStEMGFnIlNILzpZO2MwWCRlYltGOVs3alFNUEIvOlgySFpHZzd1aAolNSJGUCZtRWs6R1NSNUQxVm87NEZrQS84aTAtZnNVITU4ayNcXG9kRC1yJ1ElTVVdcF9yLFg2ayVrcEQmT0ptcE5tXk5EUlcoa0NrRCREKApxMUtrdCJEX0xkPXNFPFBJaXEvbzEhcEMmYG8qKExoQ2J0c146NSdwLHA9MFwzMVwvTzFgPFxjZCNURFRTTEcrRTMlXCE2PmhgZXJRNy5xRgo4Ni5cMEZITkpkYWpeSXNMU2RKPkYoVWVbTEpbYmdrPGRgZEk4NXNJOWVgRXQ7InNXPDJTc3IyWGNhJmQwSi0iJDZKIlxlWF5gYSMkdGMjcQotbmhSITFLSzlJUkFCXEBfKCMsNmd1dEVQVkopTXAnJSlaZSZsSixxOnJOV1JCLydKNCkrbVAiYSs1RVhVLXQ1PTY1ZHIoUFhIci9IVVU3cgoyRzdXUyUrRVkrLV1qUjA3VU0/KnBSLSJuVTpfUFdpbGAlamRBQi9jajBWNF1vTls9MydtL2dlY1QrWmw7b2RqKF9HQS1PUC8/RkE7XEJnMAokbDBqSytYL0FhOWVcMEQzYF5qXFhgbT5kJ0gqRihXNE0oPkk1WmAsJyNlOmIkKDQjJ2wrQ1NeQmtiXkxHWSIpXiRJMTE/X01scVVIJ0VsbwpCZy85clciRGIyRTM7IldacFxWcS87ZVAxJEghMUledTFHRmNBb2RJaktPSE9RREJfRSk1ZWRUblUyYSMhRmlBSFsyR0wjVyFvSEdmc2ZEPAokSEM/S1BYT0FobzdwUiMpTlEjK01bVz9SUGZGKDVTPjReTFY3YXQvUFhOPXRaVUZuY245cFVdQEUqIzxpcHU7QDllcUdAVl4zZ2ZEXyVjKQo7MGUhKzNlbzAxSFhnO1FVRyFMTjwzamphY1MtcFsyUnBNNjhgJSRGLDB0cEBXZj9CL2ZkW21JSSYtLnRAKFEnTTtdNW5PZUthc2QnVTM6VgpAN21yJCY/UDlyTVFYXTdIcTYqVzBXITNcKyNhOjovPCsiTVtVdG1wZHE5bEgrRityVlxqL2RXR1k3I0ZvJyEwZz5xOEo/NlA4cS9oM0U5PQoqa2c3XikxJC9lZ1A6Q3NOUjJwPVZPSUQuTVxqQk8yaCpgSCNCJDg0Rjs5NComcTRsL08pS0YqayIxWXI9OiZRMFovbkJpVUYla2NvXSJoNgo5ciM7LGNGLF0waCVYbzsiby9lQz5ZYGE0RjFIU2xHWHI1dFIxbldAXHUocEVEVDxbM2VDLD9XcDcnKiRYY1NwTDxqSkpeSko0X01EKTdUQAoyaSVzNVMlJSJdMUVsOVtwUTNOZituKllLcHNuVTlgJzVuSylaXUw4T2NlQHJhYk9zMyVCLE5HNCwqJEwvLGQwV2lGZkMwcylJZEFwWWQ0LwpeQWxnImJrQEJdNTtsciQ8YWtyXjtsJFIwW0pKLSFLPzJNQyNnIlo9QjFIN1suVkVyPGo9VColcVZZO2BmS1MvUCRgMmYlPUNSPEtCZj9sPApIYytSVWlyOlknWUNIQkBMTDVMWERobFw+QSZKailGJ2NrVydoQEhsZz9IRG45K2JOWEVZYSNWTCdkOypzI0VlVWotXF47WnNrOFNCcStMSwpTREZgOCc7Il9nOmlWT1pcSVdtJD5YJDE7SiUlNUBOPVhTSiVQVTBiVy0+S2QlZyxiWFBrQSdVYjgpUEdHL2g+RjBKSiZJM1t1JUtzNTVYRApiSlc5QSwnSjwwN0ByQkVpSFo5MURGJWpgPmoyL2Foa0FVO20iISFRMlRYWSMvPk1NPTNFWG1uMmBXJmxWY2xzMEM9MmM4O2ZVS3VHbytYWwpfX2A9bmY8YDduVUpzQUVVdWE5aSdpLF5mZ15iU2RjYm5bP0MoInBNN1FRSWcpK2tHS0pdIkdxK2JdKzIqRGFXM0ZAOWdgSS80ZyY7KixQKgpIIk5GPUVHT3QvNj44VFQnL2tZdEMkWnRSa240OkVZOE5qJkNHRG5sSGA4LShqWGZyRzBVOjN1KVo7V04oIyNxQC9sMiUjJHIzWS5HcUA0ZwpiVjBKNnEsKkVmbElsOywoMixbdE1DNEZVaG1gRydDNmFCIUxLIWEoZDMqZGtPQEIsJE5oci5QLlUsOVxkVVNTNV8xSEQtTCpaN0IzNmNKQQooYlVndSYnXDwtNjhFK1tcbCdtYz49QV5dJDBiKl0lMmJJbS1CRz5uW0VMcydMaGREbVVVK2RnYjZcSnRII3VlaEA6NilwT01jQikrWGEpLAo8YjFlWTFKRzNEb2FeXkRNSHBGYEBPWWZEXjthLitlNCppNkAyYkZZWjs5LS1wJW9ZIW4vMkctQFtTa1hZVWhTKFtKX0JcbyYrVCRlRSlCTApfbl5CTV8wVGxUS1thMiMuXW08O1dHKFJFLVw4JSQ/b2Y0QC5cLVh0MWU9R0MiNExAOzlrMjQxOSY8XU1IMzI3dTttRmgsRXNgKjdBNFsnTwokciFNJ0lLUip0Q0MmSDthX05sK1ApKVRcKUpxJVhAKipDSlM8PjsoW2gwYT87UWBAQmFsLj1eVTkyPTwsSEAmK2dTJ0tpMXVoT2UwNyliPwpESzI9bVYncFUiTFhZZT9uRkc2Mj9EXDkqNDJnJyFNRjVfRjBEJVpMIy5HZyZNLmpvbjlZbW5xPVtYRilnRixjPUBXM2dBVGtFZWslNT9FIQpDQ3MvPyMsWW4oRUNeMkhUUV02R2ZgP2BTbldcI1MoTWltXTdPXV09MFx0bEVvaS5QKyteOnFsQDoyXEc1RW5wR2g2PFMmIlpRXmFkVFNXRgpyNSxlMVVRZ1wpWUowS18xSlssLUJVQVZqJ0ZGR1cmX0ZxWTwqRzI6PlRCJTRdZ1NpbmljJHMnVGpPcmk1WGxIUmMmS1IoKC90SyZBU1BARwozcV9FTVxjbjM4IzdbWC9DMkRLalVrQCMzVmBTamcwKFxRaCcyNXMsWk1uaiwkYFZZPi0rdSNtU1JqMDFaPC1WUkItNC9nWDtST0c3Ozk9agpgOkhZVmViNWZIMj5iPlYwbyZbWy1hMXVfNkROZWFKMiJsX2dRVWckOEx0QyllI2BrR0YmLS1sXCQ5bls7RiJXL1dYVjtgNDs6W2tNRj5tdAo0LDs7VGJdZl8pU2YrSGNKXj5IcEYxW0pCNC1ETFhfPEgzYU9QT009Kl05ZnMsVmhAQFRJNnEsQSZvMF1DKVJ0V05pJnRnZlE0IVU7KiU+WwpgOEZVQDZbLCQ+S3MiJF1VcjNEU2ksZkguZjBBMDo4PHFTW2g4MFRML0xRVGcyUSw5cykuW19JMUpUa0xaZzJ1amtGLWVhVDs0M3NOKVMlIQpyJkBVI0xsWjRoVWVKaGY8IStscDdqRlI/blhjSTcrQ2ZAX2Mpa2gkZ3E7T2ooRXJwOWRxUitQZkAqaEcjLGBfTXMuWmUla0x1RmtZXGFsSgpnKi4+Jy8xcHQmWj1baDFhNU0sJkNUQ0pJLCJqSTJibGxLNUptVmBcIWouKTs4Yy1FYEZTaytZRTpzP09Eb1srTkpmKCJNRGEnL0VSbzNfXQo0Rzo2SyF1ImY5JWw3NkluR34+CmVuZHN0cmVhbQplbmRvYmoKMjAgMCBvYmoKPDwgL1R5cGUgL1hPYmplY3QKL1N1YnR5cGUgL0ltYWdlCi9XaWR0aCAxMTgKL0hlaWdodCA0OQovQ29sb3JTcGFjZSAvRGV2aWNlR3JheQovQml0c1BlckNvbXBvbmVudCA4Ci9MZW5ndGggNDYxCi9GaWx0ZXIgWy9BU0NJSTg1RGVjb2RlIC9GbGF0ZURlY29kZV0KPj5zdHJlYW0KR2IiL2VKSV1SPyNYblhrVDZBPlhLbkRiV0IiTVpFYWpQcGQ0XScoXkBSbHQia0IiUzAkZiU/PW09NEM7Oy4kTihqUmNZUiVzZCZsOSRPdF4KKSY3PUwoNy1lcXFlQC4oL1NDOTg0LWZMWGVMWjVxY1o1bEFFQjU1PjNAMD1rIjtSaDJoQmNRTDMmSj81b1tjSDBiU0ZOXjw7O0A0MVNZOW0KYyJaST1ZbCNjNClgXyJZbG1OMyUuQS5sTlNgVmUvWTwlPy43KCgpLy5kQls7ayFKWyUnJURKXEk/K2E3OjVvTmtZXHFFUz4zPEk1YlptKEkKJWQ8QEIpOERidDNwS0djbGJRazEnaVpWczdOTTBcbyZSKnFVPmdIdS1GckNraDxrSEZHKlwkTT1fc1YvLSRcSkY6ZS8/PmZYQStDQ21wRGUKZFNQNkdjdSxcWCdBTD1hNUpDWzQ2Ry5ocDI9VTRlXiRlb1xRI2IqbTdDOCQtSkEoY1o/WWAnQzs2SnBIZVBsXCRIRUMhdV4qUWErNmdFTWsKW2ZFVipiLiM+cHI+LVdTZiM+aCZXP3Vra1RzaFlgPS9YLjY2QSo7YmlzL2RnQFFYWHNyOG07fj4KZW5kc3RyZWFtCmVuZG9iagoyMSAwIG9iago8PCAvVHlwZSAvWE9iamVjdAovU3VidHlwZSAvSW1hZ2UKL1dpZHRoIDU0Ci9IZWlnaHQgNTQKL0NvbG9yU3BhY2UgL0RldmljZUdyYXkKL0JpdHNQZXJDb21wb25lbnQgOAovTGVuZ3RoIDc3Ci9GaWx0ZXIgWy9BU0NJSTg1RGVjb2RlIC9GbGF0ZURlY29kZV0KPj5zdHJlYW0KR2IiMEpkMFRkcSRqNG9GXlUsIkhUczlFSUU7MEFULF9FKkxaJW9AN0psNVY7SCdDcz1UcnFEYUguNEJmI2M0T1ZUOyhkI2Y8R0U5fj4KZW5kc3RyZWFtCmVuZG9iagoyMiAwIG9iago8PCAvVHlwZSAvWE9iamVjdAovU3VidHlwZSAvSW1hZ2UKL1dpZHRoIDI5NAovSGVpZ2h0IDY3Ci9Db2xvclNwYWNlIC9EZXZpY2VHcmF5Ci9CaXRzUGVyQ29tcG9uZW50IDgKL0xlbmd0aCAxMzQyCi9GaWx0ZXIgWy9BU0NJSTg1RGVjb2RlIC9GbGF0ZURlY29kZV0KPj5zdHJlYW0KR2IiL2NiQUo6YyRqNjImNUhWXy1aLGlNQjMlKTBnUksxWSJabFdgYVhtVVcyZyNxS3JvQ1ZBK2YqQiY9VTsmbT4qVSgzIm1zIVJdVG1uLnMKb2RXPmFSNWlgKG1KWC5lWiY1RCc2LDBQb1EvOCxvV3AkNT4rb1NNSTMsbDxqUUk+NSI7V19wSm9aaWUjcUptZEhrSVovJ1klWC4zbC1QNT8KI09BS2MxITIhMy9CcSxEPjEkaCxlLzVeTWc+I1MwOHA7Xyg7Nig2MTVcWDdzOWBiW1g8SVUsdDpjYz08WDxBLnIvYVNKSzs7YD1ETkRVY2EKNT1naTNrVWlYLUZIX207NTojTVxbcjIhWUJoTi4zLyIkYEZXPDwwOlFmSiRDS0h1Q1FBU2AzODJYYHNUSSRBKzBaLjUvZztTI0xNcTJDXHUKcUgtYE5jNUkoJU1vLDFIYTphQTwmNj9dcUE7RU9tUTdSQCdSUDtQYGlcUHImLTtHVVRCZG0+X0EuU1cpRTZmOkpQWkNqYGU7Z0piMzcyLkIKUz9jQDUlZjk1VDduIWxGSW4iT1JLLGNacmYtXE4tOTUsaFxVdVhVNV1bb1RSLmg5U0tVYTBaNTJRIWtHNkVXZmljLyJdQlY0aS1hYS1UX00KbVQ/QWhpXTQtalonOlImPEQpQE8yM1FebmY9SzBRSkpsQjFURW4+I2NBKiVqWGUyPEooJzgzR0gtKFRyQklwVkJtUVs8PTtIY11FLj0hXjsKSz5JZV9GZDs6VjRZYmRyMjwpIlUtdSctT2t0JjQ0W18vbV8qOUpkXFgiLzBkPCpIIzxLajdiYjhrXEQkLHE7UXMvRD4/YGZjIm90ZF09UWcKNW43XkxEcCVObklqSStOX2s2KjBvPmxhbmhjIjUzRyw7JUJQVixrVk1dKXFJTyskTDVJYFJXRC8xaypVNzpGW1s1anRpPUBfUDVJbTVfVkIKWTJ1TSQrTy85ImRTZFlNYWZKQ2hfaVhMbFQqLF5kQXFkMzJRQ1Nac1U4IWpPOEIyUnREXjlSRlBqcDpFRDElO0tkLExecFskMzFBVTVhaSIKamdgSz1wbjBSOEAtVC4uNFFiR09nOzJEUCxFRiVFQT9EKmRHYHVeJlsyMz4iO21NWFJkYmRrZi8mdFpLZiMuaVcvdT5xZ0tbQFU7JjphKlwKNlRWZnMjXlNzbldGZEowZ2o0bDU2PV9Nc1Y4Jl8iZSdTQ11SOEc0MW5ZN01kNjl0Jyk0Pmo3Qmclbk1ZM1JzaCVPRXAtTGY6O2BaLXJDYXQKbTloQHBWOT4qRm9jcFpoI2MlU1YlRiVQSVBVX3R0S2guKjpHPDBoalxGO2RyO3NtT2U4cDtgJygyO0dQaGY5VmA5Jzo8P2xgWVVQa0gzOVAKQW4kcXUuJzY+UENCVlpwanAvUzBkImwrOmwnXk8oQmRtPmJGPlNnIj1URlhUJmNRPThKMk1QY2UpTmVKT2ZoWy8oYGhPcClDIkN1ZC1DXi8KSEZnJichYVxKbGtxPUhpLzpwWUJMSydiWkEoY2tmTlssRVg8KiFXJWYtci1ANWlWODVINDMuLWBFYFVAMUV0byQkQzkrPFxrVC5jR2NBN14KSDhrViEuTjorKmYiMVhRJW5GJmIzZSUybVBFKm5RKyxXbVBWTys1c1xtdVtcZTYsQzprcGUiNjwwKSZxUTlwQWtPOjAhZVtSP3AhbzpmMWwKWSM+dTNAbScuJ0xTaW5TOERAaDw4OCU1JlY2bEo5TXBhZmhKLFRKQCJyQGF+PgplbmRzdHJlYW0KZW5kb2JqCnhyZWYKMCAyMwowMDAwMDAwMDAwIDY1NTM1IGYgCjAwMDAwMDAwMDkgMDAwMDAgbiAKMDAwMDAwMDA1OCAwMDAwMCBuIAowMDAwMDAwMTA0IDAwMDAwIG4gCjAwMDAwMDAxNjIgMDAwMDAgbiAKMDAwMDAwMDIxNCAwMDAwMCBuIAowMDAwMDAwMzEyIDAwMDAwIG4gCjAwMDAwMDA0MTUgMDAwMDAgbiAKMDAwMDAwMDUyMSAwMDAwMCBuIAowMDAwMDAwNjMxIDAwMDAwIG4gCjAwMDAwMDA3MjcgMDAwMDAgbiAKMDAwMDAwMDgyOSAwMDAwMCBuIAowMDAwMDAwOTM0IDAwMDAwIG4gCjAwMDAwMDEwNDMgMDAwMDAgbiAKMDAwMDAwMTE0NCAwMDAwMCBuIAowMDAwMDAxMjQ0IDAwMDAwIG4gCjAwMDAwMDEzNDYgMDAwMDAgbiAKMDAwMDAwMTQ1MiAwMDAwMCBuIAowMDAwMDAxNjIyIDAwMDAwIG4gCjAwMDAwMDIwMDEgMDAwMDAgbiAKMDAwMDAwNjI0MyAwMDAwMCBuIAowMDAwMDA2ODkwIDAwMDAwIG4gCjAwMDAwMDcxNTEgMDAwMDAgbiAKdHJhaWxlcgo8PAovSW5mbyAxNyAwIFIKL1NpemUgMjMKL1Jvb3QgMSAwIFIKPj4Kc3RhcnR4cmVmCjg2ODAKJSVFT0YK"
    }, {
      "id" : "fb9ada7904ef4f22a81b020657567bdb",
      "tracking_numbers" : [ "5332787865" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T10:59:48.729759Z",
      "updated_at" : "2025-03-12T10:59:52.067787Z",
      "succeed_at" : "2025-03-12T10:59:52.060707568Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "GBR",
            "postal_code" : "CO3 0PP"
          },
          "ship_to" : {
            "country" : "CAN",
            "postal_code" : "M6A 1P6"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Colchester",
        "contact_name" : "Courtney Elliott",
        "country" : "GBR",
        "email" : "t@t.com",
        "phone" : "416-306-8001",
        "postal_code" : "CO3 0PP",
        "state" : "England",
        "street1" : "19 Spring Sedge Close"
      },
      "ship_to" : {
        "city" : "North York",
        "company_name" : "UK-MDC",
        "contact_name" : "UK-MDC",
        "country" : "CAN",
        "phone" : "416-306-8001",
        "postal_code" : "M6A 1P6",
        "state" : "ON",
        "street1" : "105 Bentworth Avenue",
        "street2" : "Unit 4"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "GBR",
      "ship_to_country" : "CAN",
      "order_id" : null,
      "order_number" : "#nyy1301",
      "return_shipment" : true,
      "files" : {
        "invoice" : {
          "file_type" : "invoice",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-12/fb9ada79-04ef-4f22-a81b-020657567bdb-1741777191142.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x8",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/fb9ada79-04ef-4f22-a81b-020657567bdb-1741777191142.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA REFLSL4PN7U", "ORDER REFM192313764UZ" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "dhl",
      "carrier_account_id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
      "service_type" : "dhl_express_worldwide",
      "rate" : {
        "delivery_date" : "2025-03-17T23:59:00Z",
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "DHL Express Worldwide",
        "service_type" : "dhl_express_worldwide",
        "shipper_account" : {
          "description" : "[mydhl] Demo Account [Automate Test]",
          "id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
          "slug" : "dhl"
        },
        "transit_time" : 6
      },
      "service_options" : [ {
        "enabled" : true,
        "type" : "paperless_invoice"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "description" : "",
            "dimension" : {
              "depth" : 2.1,
              "height" : 2.1,
              "unit" : "cm",
              "width" : 2.1
            },
            "items" : [ {
              "barcode" : "",
              "description" : "Single Mini Hoop - 14k Yellow Gold",
              "description_cn" : "",
              "dimension" : {
                "depth" : 0,
                "height" : 0,
                "unit" : "",
                "width" : 0
              },
              "hs_code" : "7113195090",
              "image_urls" : [ "https://cdn.shopify.com/s/files/1/0837/7489/8461/files/0-SingleMiniHoop-14K-Angled_045_50e57c60-f395-4735-be7c-58865a504acb.jpg?v=1722958121" ],
              "item_id" : "15992277532957",
              "origin_country" : "CAN",
              "price" : {
                "amount" : 31.67,
                "currency" : "GBP"
              },
              "quantity" : 1,
              "return_reason" : "",
              "sku" : "p52704669",
              "weight" : {
                "unit" : "lb",
                "value" : 6.875E-4
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 1.011
            }
          } ]
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.011
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 2.1,
        "height" : 2.1,
        "unit" : "cm",
        "width" : 2.1
      } ],
      "items" : [ {
        "description" : "Single Mini Hoop - 14k Yellow Gold",
        "hs_code" : "7113195090",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0837/7489/8461/files/0-SingleMiniHoop-14K-Angled_045_50e57c60-f395-4735-be7c-58865a504acb.jpg?v=1722958121" ],
        "item_id" : "15992277532957",
        "origin_country" : "CAN",
        "price" : {
          "amount" : 31.67,
          "currency" : "GBP"
        },
        "quantity" : 1,
        "sku" : "p52704669",
        "weight" : {
          "unit" : "lb",
          "value" : 6.875E-4
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "release-nike-sandbox-api.aftership.io",
      "label_base64" : null
    }, {
      "id" : "f943db7da73e48d29331a3e2f89c886c",
      "tracking_numbers" : null,
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T10:42:36.818959Z",
      "updated_at" : "2025-03-12T10:42:40.374238Z",
      "succeed_at" : "2025-03-12T10:42:40.364124921Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Salt Lake City",
        "contact_name" : "test nyy",
        "country" : "USA",
        "email" : "yy.nie@aftership.com",
        "phone" : "0123456789",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "phone" : "0123456789",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "a403c966b49d4561ba590ef8f9b1bb45",
      "order_number" : "#nyy1029",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/f943db7d-a73e-48d2-9331-a3e2f89c886c-1741776160137.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "test20250312001436543" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "gps",
      "carrier_account_id" : "a4f2c3c784424c4da0eb3ffefa582542",
      "service_type" : "gps_ups_ground",
      "rate" : {
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 6.6,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 4.720000000000001,
            "currency" : "USD"
          },
          "type" : "other"
        } ],
        "service_name" : "UPS® Ground",
        "service_type" : "gps_ups_ground",
        "shipper_account" : {
          "description" : "[gps] Demo Account",
          "id" : "a4f2c3c784424c4da0eb3ffefa582542",
          "slug" : "gps"
        },
        "total_charge" : {
          "amount" : 11.32,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "tracking_slug" : "ups"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 3.2
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 3,
        "height" : 5,
        "unit" : "cm",
        "width" : 4
      } ],
      "items" : [ {
        "description" : "苹果X XR玻璃手机壳",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/2407/9670/products/O1CN01PDsTa62MCKM0hpcrR__0-item_pic.jpg_430x430q90.jpg?v=1687770918" ],
        "item_id" : "16181170438454",
        "origin_country" : "USA",
        "price" : {
          "amount" : 14,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "9990098449929707830GHJKI UIYGGGUIIHHBVVVBHHJJJJJJJJJJBB",
        "weight" : {
          "unit" : "kg",
          "value" : 0.2
        }
      } ],
      "billing" : {
        "paid_by" : "recipient"
      },
      "operator_id" : "release-nike-sandbox-api.aftership.io",
      "label_base64" : null
    }, {
      "id" : "79f9ac0dd8fc47429a309673dc6749ba",
      "tracking_numbers" : [ "794684895004" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T10:39:28.163007Z",
      "updated_at" : "2025-03-12T10:39:38.834206Z",
      "succeed_at" : "2025-03-12T10:39:38.823690915Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "CAN",
            "postal_code" : "B3Z 0M6"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "94209"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Upper Tantallon",
        "company_name" : "Jacqui Allain",
        "contact_name" : "fadf tafgsdfa",
        "country" : "CAN",
        "email" : "asjdfla@test.com",
        "phone" : "(902)8262222",
        "postal_code" : "B3Z 0M6",
        "state" : "NS",
        "street1" : "70 Rockfield Drive",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sacramento",
        "company_name" : "Lakin and Sons",
        "contact_name" : "Dr.Joe",
        "country" : "USA",
        "email" : "sample@test.com",
        "fax" : "190-644-2218",
        "phone" : "'1-140-225-3341",
        "postal_code" : "94209",
        "state" : "CA",
        "street1" : "28292 Daugherty Orchard",
        "type" : "business"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "CAN",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : "nyy32434001",
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-12/bd721c66-350a-4368-aea6-b471d248b92e-1741775978441639.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/0c2dd5c1-1f83-413e-adc4-b81467c038f1-1741775975851147.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "P_O_NUMBER:00000", "12345", "DEPARTMENT_NUMBER:CS4/NGST/12345" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
      "service_type" : "fedex_international_priority",
      "rate" : {
        "charge_weight" : {
          "unit" : "lb",
          "value" : 14.7
        },
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 298.62,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 1.5,
            "currency" : "USD"
          },
          "type" : "us_inbound_processing_fee"
        }, {
          "charge" : {
            "amount" : 53,
            "currency" : "USD"
          },
          "type" : "demand_surcharge"
        }, {
          "charge" : {
            "amount" : 17.54,
            "currency" : "USD"
          },
          "type" : "fuel"
        } ],
        "service_name" : "FedEx International Priority®",
        "service_type" : "fedex_international_priority",
        "shipper_account" : {
          "description" : "[FedEx] Postmen Testing Account",
          "id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
          "slug" : "fedex"
        },
        "total_charge" : {
          "amount" : 370.66,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : {
          "amount" : 370.66,
          "currency" : "USD"
        },
        "form_id" : 430,
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "description" : "Food XS",
            "dimension" : {
              "depth" : 40,
              "height" : 40,
              "unit" : "cm",
              "width" : 20
            },
            "items" : [ {
              "description" : "Food Bar",
              "origin_country" : "USA",
              "price" : {
                "amount" : 3,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "imac2014",
              "weight" : {
                "unit" : "kg",
                "value" : 0.6
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 2
            }
          } ]
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 2
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 14.7
      },
      "dimension" : [ {
        "depth" : 40,
        "height" : 40,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "origin_country" : "USA",
        "price" : {
          "amount" : 3,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "release-nike-sandbox-api.aftership.io",
      "label_base64" : null
    }, {
      "id" : "515d06294b6f43bfa50ff962532ae35b",
      "tracking_numbers" : [ "794684894913" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T10:38:23.668079Z",
      "updated_at" : "2025-03-12T10:38:38.764267Z",
      "succeed_at" : "2025-03-12T10:38:38.756678231Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "CAN",
            "postal_code" : "B3Z 0M6"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "94209"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Upper Tantallon",
        "company_name" : "Jacqui Allain",
        "contact_name" : "fadf tafgsdfa",
        "country" : "CAN",
        "email" : "asjdfla@test.com",
        "phone" : "(902)8262222",
        "postal_code" : "B3Z 0M6",
        "state" : "NS",
        "street1" : "70 Rockfield Drive",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sacramento",
        "company_name" : "Lakin and Sons",
        "contact_name" : "Dr.Joe",
        "country" : "USA",
        "email" : "sample@test.com",
        "fax" : "190-644-2218",
        "phone" : "'1-140-225-3341",
        "postal_code" : "94209",
        "state" : "CA",
        "street1" : "28292 Daugherty Orchard",
        "type" : "business"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "CAN",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-12/a7cf7173-28b7-493a-987a-d282d64d0c02-1741775917746005.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/44c582d5-06ed-4d36-a098-ea0b199a6e47-1741775914348199.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "P_O_NUMBER:00000", "12345", "DEPARTMENT_NUMBER:CS4/NGST/12345" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
      "service_type" : "fedex_international_priority",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "FedEx International Priority®",
        "service_type" : "fedex_international_priority",
        "shipper_account" : {
          "description" : "[FedEx] Postmen Testing Account",
          "id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
          "slug" : "fedex"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null,
        "form_id" : 430,
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "description" : "Food XS",
            "dimension" : {
              "depth" : 40,
              "height" : 40,
              "unit" : "cm",
              "width" : 20
            },
            "items" : [ {
              "description" : "Food Bar",
              "origin_country" : "USA",
              "price" : {
                "amount" : 3,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "imac2014",
              "weight" : {
                "unit" : "kg",
                "value" : 0.6
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 2
            }
          } ]
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 2
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 40,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "origin_country" : "USA",
        "price" : {
          "amount" : 3,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "release-nike-sandbox-api.aftership.io",
      "label_base64" : null
    }, {
      "id" : "c2dc691902fc42eaaeaffd235e4cf599",
      "tracking_numbers" : [ "SF7444495652101", "SF7444515516896" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T09:57:41.985471Z",
      "updated_at" : "2025-03-12T09:57:46.477559Z",
      "succeed_at" : "2025-03-12T09:57:46.469695167Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "CHN",
            "postal_code" : "518000"
          },
          "ship_to" : {
            "country" : "TWN",
            "postal_code" : "097612"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "深圳市",
        "company_name" : "aftership(深圳)",
        "contact_name" : "Panda",
        "country" : "CHN",
        "phone" : "13800138000",
        "postal_code" : "518000",
        "state" : "南山区",
        "street1" : "广东省深圳市南山区粤兴三道18号AAA大楼1号区",
        "street2" : "A区19幢39单元",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "香港",
        "company_name" : "aftership(香港)",
        "contact_name" : "大王叫你去巡山",
        "country" : "TWN",
        "email" : "test@test.com",
        "phone" : "(662)123-4567",
        "postal_code" : "097612",
        "street1" : "香港新界荃灣青山公路625號麗城花園第三期麗城薈地下48及49A號舖",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "CHN",
      "ship_to_country" : "TWN",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-12/c2e200be-eeab-4375-997e-6b0ee8a07f9f-1741773465414958.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/08c4791a-d3cf-421c-ad40-f3f8831a22c8-1741773464734902.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "183934-139324-34233" ],
      "carrier_references" : null,
      "carrier_account_slug" : "sf-express",
      "carrier_account_id" : "32146288eba64c809a2c3f4365c5aa5c",
      "service_type" : "sf-express_standard_express",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "SF Express Standard Express",
        "service_type" : "sf-express_standard_express",
        "shipper_account" : {
          "description" : "[sf-express] Demo Account can use",
          "id" : "32146288eba64c809a2c3f4365c5aa5c",
          "slug" : "sf-express"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 19.68
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      }, {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      } ],
      "items" : [ {
        "description" : "D04010150006 - Wild Mugwort Restoring Serumu0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      }, {
        "description" : "D04010150006 - Wild Mugwort Restoring Serumu0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "f8df9ad37d1642aaa564ea0fe350ea6b",
      "tracking_numbers" : [ "SF7444495652068" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T09:55:40.021935Z",
      "updated_at" : "2025-03-12T09:55:54.851987Z",
      "succeed_at" : "2025-03-12T09:55:54.845289675Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "CHN",
            "postal_code" : "518000"
          },
          "ship_to" : {
            "country" : "TWN",
            "postal_code" : "097612"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "深圳市",
        "company_name" : "aftership(深圳)",
        "contact_name" : "Panda",
        "country" : "CHN",
        "phone" : "13800138000",
        "postal_code" : "518000",
        "state" : "南山区",
        "street1" : "广东省深圳市南山区粤兴三道18号AAA大楼1号区",
        "street2" : "A区19幢39单元",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "香港",
        "company_name" : "aftership(香港)",
        "contact_name" : "大王叫你去巡山",
        "country" : "TWN",
        "email" : "test@test.com",
        "phone" : "(662)123-4567",
        "postal_code" : "097612",
        "street1" : "香港新界荃灣青山公路625號麗城花園第三期麗城薈地下48及49A號舖",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "CHN",
      "ship_to_country" : "TWN",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-12/c9fc2bcd-4877-4945-8693-e9d328ef4c92-1741773353706651.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/f1bc8bb9-3da8-4cff-b8dd-de0917b2a653-1741773344661664.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "183934-139324-34233" ],
      "carrier_references" : null,
      "carrier_account_slug" : "sf-express",
      "carrier_account_id" : "32146288eba64c809a2c3f4365c5aa5c",
      "service_type" : "sf-express_standard_express",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "SF Express Standard Express",
        "service_type" : "sf-express_standard_express",
        "shipper_account" : {
          "description" : "[sf-express] Demo Account can use",
          "id" : "32146288eba64c809a2c3f4365c5aa5c",
          "slug" : "sf-express"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 9.84
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      } ],
      "items" : [ {
        "description" : "D04010150006 - Wild Mugwort Restoring Serumu0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "992422c3543a4ae08c0afc7d6a5d39f4",
      "tracking_numbers" : [ "SF7444495648803" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T07:58:52.497071Z",
      "updated_at" : "2025-03-12T07:58:54.52188Z",
      "succeed_at" : "2025-03-12T07:58:54.513161932Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "CHN",
            "postal_code" : "518000"
          },
          "ship_to" : {
            "country" : "TWN",
            "postal_code" : "097612"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "深圳市",
        "company_name" : "aftership(深圳)",
        "contact_name" : "Panda",
        "country" : "CHN",
        "phone" : "13800138000",
        "postal_code" : "518000",
        "state" : "南山区",
        "street1" : "广东省深圳市南山区粤兴三道18号AAA大楼1号区",
        "street2" : "A区19幢39单元",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "香港",
        "company_name" : "aftership(香港)",
        "contact_name" : "大王叫你去巡山",
        "country" : "TWN",
        "email" : "test@test.com",
        "phone" : "(662)123-4567",
        "postal_code" : "097612",
        "street1" : "香港新界荃灣青山公路625號麗城花園第三期麗城薈地下48及49A號舖",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "CHN",
      "ship_to_country" : "TWN",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-12/8859abf4-e815-44b2-94a5-9ebb43164a6c-1741766333575233.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "8x12",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/a67d22c1-426a-4630-bc26-b27f8c561812-1741766333558684.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "183934-139324-34233" ],
      "carrier_references" : null,
      "carrier_account_slug" : "sf-express",
      "carrier_account_id" : "32146288eba64c809a2c3f4365c5aa5c",
      "service_type" : "sf-express_standard_express",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "SF Express Standard Express",
        "service_type" : "sf-express_standard_express",
        "shipper_account" : {
          "description" : "[sf-express] Demo Account can use",
          "id" : "32146288eba64c809a2c3f4365c5aa5c",
          "slug" : "sf-express"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 9.84
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      } ],
      "items" : [ {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "72220c02a0e74bd5adb7da92e297bd50",
      "tracking_numbers" : [ "SF7444495648760", "SF7444515516611", "SF7444515516620", "SF7444515516639" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T07:57:36.266559Z",
      "updated_at" : "2025-03-12T07:57:48.650294Z",
      "succeed_at" : "2025-03-12T07:57:48.642090949Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "CHN",
            "postal_code" : "518000"
          },
          "ship_to" : {
            "country" : "TWN",
            "postal_code" : "097612"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "深圳市",
        "company_name" : "aftership(深圳)",
        "contact_name" : "Panda",
        "country" : "CHN",
        "phone" : "13800138000",
        "postal_code" : "518000",
        "state" : "南山区",
        "street1" : "广东省深圳市南山区粤兴三道18号AAA大楼1号区",
        "street2" : "A区19幢39单元",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "香港",
        "company_name" : "aftership(香港)",
        "contact_name" : "大王叫你去巡山",
        "country" : "TWN",
        "email" : "test@test.com",
        "phone" : "(662)123-4567",
        "postal_code" : "097612",
        "street1" : "香港新界荃灣青山公路625號麗城花園第三期麗城薈地下48及49A號舖",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "CHN",
      "ship_to_country" : "TWN",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-12/4064d0df-d0de-4364-9d3b-b56479b1e065-1741766264256264.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "8x12",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/7270857f-f3a5-443e-a26b-e247a1aa3d27-1741766267388683.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "183934-139324-34233" ],
      "carrier_references" : null,
      "carrier_account_slug" : "sf-express",
      "carrier_account_id" : "32146288eba64c809a2c3f4365c5aa5c",
      "service_type" : "sf-express_standard_express",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "SF Express Standard Express",
        "service_type" : "sf-express_standard_express",
        "shipper_account" : {
          "description" : "[sf-express] Demo Account can use",
          "id" : "32146288eba64c809a2c3f4365c5aa5c",
          "slug" : "sf-express"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 39.36
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      }, {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      }, {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      }, {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      } ],
      "items" : [ {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      }, {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      }, {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      }, {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "9771af2a826842fdbf1b0efd5544cb55",
      "tracking_numbers" : [ "SF7444495645370", "SF7444515516453", "SF7444515516462", "SF7444515516471" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T07:41:26.346463Z",
      "updated_at" : "2025-03-12T07:41:32.554741Z",
      "succeed_at" : "2025-03-12T07:41:32.541697901Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "CHN",
            "postal_code" : "518000"
          },
          "ship_to" : {
            "country" : "TWN",
            "postal_code" : "097612"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "深圳市",
        "company_name" : "aftership(深圳)",
        "contact_name" : "Panda",
        "country" : "CHN",
        "phone" : "13800138000",
        "postal_code" : "518000",
        "state" : "南山区",
        "street1" : "广东省深圳市南山区粤兴三道18号AAA大楼1号区",
        "street2" : "A区19幢39单元",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "香港",
        "company_name" : "aftership(香港)",
        "contact_name" : "大王叫你去巡山",
        "country" : "TWN",
        "email" : "test@test.com",
        "phone" : "(662)123-4567",
        "postal_code" : "097612",
        "street1" : "香港新界荃灣青山公路625號麗城花園第三期麗城薈地下48及49A號舖",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "CHN",
      "ship_to_country" : "TWN",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-12/46a4ce77-4f1c-467f-9534-1d9b2007e5d7-1741765291579845.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/9a24d544-6a5b-47b5-b35d-6c7a5cead3b9-1741765290043552.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "183934-139324-34233" ],
      "carrier_references" : null,
      "carrier_account_slug" : "sf-express",
      "carrier_account_id" : "32146288eba64c809a2c3f4365c5aa5c",
      "service_type" : "sf-express_standard_express",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "SF Express Standard Express",
        "service_type" : "sf-express_standard_express",
        "shipper_account" : {
          "description" : "[sf-express] Demo Account can use",
          "id" : "32146288eba64c809a2c3f4365c5aa5c",
          "slug" : "sf-express"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 39.36
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      }, {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      }, {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      }, {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      } ],
      "items" : [ {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      }, {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      }, {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      }, {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "5432f7b8539c403292a083141b3d9e43",
      "tracking_numbers" : [ "13033378" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T07:31:29.097327Z",
      "updated_at" : "2025-03-12T07:31:31.931797Z",
      "succeed_at" : "2025-03-12T07:31:31.92335295Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP",
            "postal_code" : "08150"
          },
          "ship_to" : {
            "country" : "FRA",
            "postal_code" : "75001"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles(barcelona)",
        "company_name" : "Panda of china",
        "contact_name" : "Javid",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08150",
        "state" : "Mixed",
        "street1" : "ho no 115/flat no b1 street no 7.",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Paris",
        "contact_name" : "Ismael Tejón",
        "country" : "FRA",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "75001",
        "state" : "VIL",
        "street1" : "Calle de prueba, 34.",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "ESP",
      "ship_to_country" : "FRA",
      "order_id" : null,
      "order_number" : "100001",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/5432f7b8-539c-4032-92a0-83141b3d9e43-1741764691145.pdf"
        }
      },
      "box_type" : "nacex_sample",
      "references" : [ "REFERENCE0" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "nacex",
      "carrier_account_id" : "c5573249c14b4f9eb077e36985b2f6e9",
      "service_type" : "nacex_servicio-aereo",
      "rate" : {
        "delivery_date" : "2025-03-14",
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "NACEX SERVICIO AEREO",
        "service_type" : "nacex_servicio-aereo",
        "shipper_account" : {
          "description" : "[nacex] Demo account new",
          "id" : "c5573249c14b4f9eb077e36985b2f6e9",
          "slug" : "nacex"
        },
        "transit_time" : 2
      },
      "service_options" : [ {
        "end_time" : "13:00:00",
        "start_time" : "12:00:00",
        "type" : "pickup"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment_code" : "430032811"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "ESP",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "release-core-sandbox-api.aftership.io",
      "label_base64" : null
    }, {
      "id" : "237e4744ec83487eb909fbce894cd318",
      "tracking_numbers" : [ "SF7444495644675", "SF7444515516250", "SF7444515516269", "SF7444515516278" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T07:19:25.120063Z",
      "updated_at" : "2025-03-12T07:19:33.20348Z",
      "succeed_at" : "2025-03-12T07:19:33.187505618Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "CHN",
            "postal_code" : "518000"
          },
          "ship_to" : {
            "country" : "TWN",
            "postal_code" : "097612"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "深圳市",
        "company_name" : "aftership(深圳)",
        "contact_name" : "Panda",
        "country" : "CHN",
        "phone" : "13800138000",
        "postal_code" : "518000",
        "state" : "南山区",
        "street1" : "广东省深圳市南山区粤兴三道18号AAA大楼1号区",
        "street2" : "A区19幢39单元",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "香港",
        "company_name" : "aftership(香港)",
        "contact_name" : "大王叫你去巡山",
        "country" : "TWN",
        "email" : "test@test.com",
        "phone" : "(662)123-4567",
        "postal_code" : "097612",
        "street1" : "香港新界荃灣青山公路625號麗城花園第三期麗城薈地下48及49A號舖",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "CHN",
      "ship_to_country" : "TWN",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-12/3f93f46d-2a26-494a-9dc0-b53ce2450b59-1741763972164810.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/c58aedb2-1402-408c-ba62-6b76a925f0e1-1741763970140823.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "183934-139324-34233" ],
      "carrier_references" : null,
      "carrier_account_slug" : "sf-express",
      "carrier_account_id" : "32146288eba64c809a2c3f4365c5aa5c",
      "service_type" : "sf-express_standard_express",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "SF Express Standard Express",
        "service_type" : "sf-express_standard_express",
        "shipper_account" : {
          "description" : "[sf-express] Demo Account can use",
          "id" : "32146288eba64c809a2c3f4365c5aa5c",
          "slug" : "sf-express"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 39.36
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      }, {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      }, {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      }, {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      } ],
      "items" : [ {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      }, {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      }, {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      }, {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "f7afdeada064478190e815fda93582f4",
      "tracking_numbers" : [ "SF7444495644611" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T07:16:13.598447Z",
      "updated_at" : "2025-03-12T07:16:21.746122Z",
      "succeed_at" : "2025-03-12T07:16:21.731967209Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "CHN",
            "postal_code" : "518000"
          },
          "ship_to" : {
            "country" : "TWN",
            "postal_code" : "097612"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "深圳市",
        "company_name" : "aftership(深圳)",
        "contact_name" : "Panda",
        "country" : "CHN",
        "phone" : "13800138000",
        "postal_code" : "518000",
        "state" : "南山区",
        "street1" : "广东省深圳市南山区粤兴三道18号AAA大楼1号区",
        "street2" : "A区19幢39单元",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "香港",
        "company_name" : "aftership(香港)",
        "contact_name" : "大王叫你去巡山",
        "country" : "TWN",
        "email" : "test@test.com",
        "phone" : "(662)123-4567",
        "postal_code" : "097612",
        "street1" : "香港新界荃灣青山公路625號麗城花園第三期麗城薈地下48及49A號舖",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "CHN",
      "ship_to_country" : "TWN",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-12/6bf9d0e5-8fd8-4f85-b013-133a7ccec8c5-1741763780497545.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/988f887c-6cea-4438-a6bb-cdc3626fdfcb-1741763776062819.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "183934-139324-34233" ],
      "carrier_references" : null,
      "carrier_account_slug" : "sf-express",
      "carrier_account_id" : "32146288eba64c809a2c3f4365c5aa5c",
      "service_type" : "sf-express_standard_express",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "SF Express Standard Express",
        "service_type" : "sf-express_standard_express",
        "shipper_account" : {
          "description" : "[sf-express] Demo Account can use",
          "id" : "32146288eba64c809a2c3f4365c5aa5c",
          "slug" : "sf-express"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 9.84
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      } ],
      "items" : [ {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "7c0cf5ba4a2d4bf0bf6e0961434d8495",
      "tracking_numbers" : [ "794684857172" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T06:41:57.929679Z",
      "updated_at" : "2025-03-12T06:42:11.270696Z",
      "succeed_at" : "2025-03-12T06:42:11.261122383Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "92705"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Salt Lake City",
        "contact_name" : "test nyy",
        "country" : "USA",
        "email" : "yy.nie@aftership.com",
        "phone" : "0123456789",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY"
      },
      "ship_to" : {
        "city" : "Santa Ana",
        "company_name" : "N/A",
        "contact_name" : "Paul S. Allen",
        "country" : "USA",
        "email" : "sssss@collectors.com",
        "phone" : "123456",
        "postal_code" : "92705",
        "state" : "CA",
        "street1" : "1600 E. Saint Andrew Pl.",
        "street2" : "#100"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-12/4d9b4f76-8163-458c-8137-8a65f0e3b610-1741761730899686.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/1587ad64-90ff-45b4-bfe2-507eb85a13d7-1741761728946650.pdf"
        }
      },
      "box_type" : "fedex_tube",
      "references" : [ "reference1" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "9889e945-aa59-4b6e-ab40-aedb7f491668",
      "service_type" : "fedex_2_day",
      "rate" : {
        "charge_weight" : {
          "unit" : "lb",
          "value" : 9
        },
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 48.51,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 2.43,
            "currency" : "USD"
          },
          "type" : "fuel"
        } ],
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "[fedex] testing account",
          "id" : "9889e945-aa59-4b6e-ab40-aedb7f491668",
          "slug" : "fedex"
        },
        "total_charge" : {
          "amount" : 50.94,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : {
          "amount" : 50.94,
          "currency" : "USD"
        },
        "form_id" : 201,
        "shipment" : {
          "parcels" : [ {
            "box_type" : "fedex_tube",
            "description" : "Phones",
            "dimension" : {
              "depth" : 40,
              "height" : 30,
              "unit" : "cm",
              "width" : 20
            },
            "items" : [ {
              "description" : "Samsung Galaxy Tab wifi 3G 64GB black",
              "hs_code" : "315815",
              "origin_country" : "USA",
              "price" : {
                "amount" : 50,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "Epic_Food_Bar",
              "weight" : {
                "unit" : "kg",
                "value" : 0.6
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 1.5
            }
          } ]
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 9
      },
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Samsung Galaxy Tab wifi 3G 64GB black",
        "hs_code" : "315815",
        "origin_country" : "USA",
        "price" : {
          "amount" : 50,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "bf5babd48bee4ef5a53547caa3c8b718",
      "tracking_numbers" : null,
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T05:50:42.680047Z",
      "updated_at" : "2025-03-12T05:50:46.198188Z",
      "succeed_at" : "2025-03-12T05:50:46.189308003Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Salt Lake City",
        "contact_name" : "test nyy",
        "country" : "USA",
        "email" : "yy.nie@aftership.com",
        "phone" : "0123456789",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "phone" : "0123456789",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "a403c966b49d4561ba590ef8f9b1bb45",
      "order_number" : "#1024",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/bf5babd4-8bee-4ef5-a535-47caa3c8b718-1741758645938.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "test20250312002" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "gps",
      "carrier_account_id" : "a4f2c3c784424c4da0eb3ffefa582542",
      "service_type" : "gps_ups_ground",
      "rate" : {
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 6.6,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 4.720000000000001,
            "currency" : "USD"
          },
          "type" : "other"
        } ],
        "service_name" : "UPS® Ground",
        "service_type" : "gps_ups_ground",
        "shipper_account" : {
          "description" : "[gps] Demo Account",
          "id" : "a4f2c3c784424c4da0eb3ffefa582542",
          "slug" : "gps"
        },
        "total_charge" : {
          "amount" : 11.32,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "tracking_slug" : "ups"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 3.2
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 3,
        "height" : 5,
        "unit" : "cm",
        "width" : 4
      } ],
      "items" : [ {
        "description" : "苹果X XR玻璃手机壳",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/2407/9670/products/O1CN01PDsTa62MCKM0hpcrR__0-item_pic.jpg_430x430q90.jpg?v=1687770918" ],
        "item_id" : "16181170438454",
        "origin_country" : "USA",
        "price" : {
          "amount" : 14,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "9990098449929707830GHJKI UIYGGGUIIHHBVVVBHHJJJJJJJJJJBB",
        "weight" : {
          "unit" : "kg",
          "value" : 0.2
        }
      } ],
      "billing" : {
        "paid_by" : "recipient"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "2a4b7c7a40be455abdad3659793f8099",
      "tracking_numbers" : [ "SF7444495641284" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T05:48:04.104415Z",
      "updated_at" : "2025-03-12T05:48:09.073471Z",
      "succeed_at" : "2025-03-12T05:48:07.489002067Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "CHN",
            "postal_code" : "518000"
          },
          "ship_to" : {
            "country" : "TWN",
            "postal_code" : "097612"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "深圳市",
        "company_name" : "aftership(深圳)",
        "contact_name" : "Panda",
        "country" : "CHN",
        "phone" : "13800138000",
        "postal_code" : "518000",
        "state" : "南山区",
        "street1" : "广东省深圳市南山区粤兴三道18号AAA大楼1号区",
        "street2" : "A区19幢39单元",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "香港",
        "company_name" : "aftership(香港)",
        "contact_name" : "大王叫你去巡山",
        "country" : "TWN",
        "email" : "test@test.com",
        "phone" : "(662)123-4567",
        "postal_code" : "097612",
        "street1" : "香港新界荃灣青山公路625號麗城花園第三期麗城薈地下48及49A號舖",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "CHN",
      "ship_to_country" : "TWN",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-12/0d7fd9e3-e2f4-4e52-a453-ae150e1b8fa1-1741758486610323.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/0e63d407-2778-48c1-92cf-0b7ede94051e-1741758486155778.pdf"
        },
        "packing_slip" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/packing_slip/2025-03-12/541691e6-6f18-4609-bf85-32499ade3780-1741758488748608.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "183934-139324-34233" ],
      "carrier_references" : null,
      "carrier_account_slug" : "sf-express",
      "carrier_account_id" : "32146288eba64c809a2c3f4365c5aa5c",
      "service_type" : "sf-express_standard_express",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "SF Express Standard Express",
        "service_type" : "sf-express_standard_express",
        "shipper_account" : {
          "description" : "[sf-express] Demo Account can use",
          "id" : "32146288eba64c809a2c3f4365c5aa5c",
          "slug" : "sf-express"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 9.84
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      } ],
      "items" : [ {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "cc622b7a77bc4256a5805fbbcdda9c8b",
      "tracking_numbers" : [ "SF7444495641275" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T05:47:39.801263Z",
      "updated_at" : "2025-03-12T05:47:50.093017Z",
      "succeed_at" : "2025-03-12T05:47:49.209926506Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "CHN",
            "postal_code" : "518000"
          },
          "ship_to" : {
            "country" : "TWN",
            "postal_code" : "097612"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "深圳市",
        "company_name" : "aftership(深圳)",
        "contact_name" : "Panda",
        "country" : "CHN",
        "phone" : "13800138000",
        "postal_code" : "518000",
        "state" : "南山区",
        "street1" : "广东省深圳市南山区粤兴三道18号AAA大楼1号区",
        "street2" : "A区19幢39单元",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "香港",
        "company_name" : "aftership(香港)",
        "contact_name" : "大王叫你去巡山",
        "country" : "TWN",
        "email" : "test@test.com",
        "phone" : "(662)123-4567",
        "postal_code" : "097612",
        "street1" : "香港新界荃灣青山公路625號麗城花園第三期麗城薈地下48及49A號舖",
        "type" : "residential"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "CHN",
      "ship_to_country" : "TWN",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-12/06c19f13-f518-4223-97b5-eb3a1c4488a7-1741758468243937.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/e2e9b877-82cb-4a4c-9f33-6b366e74b2da-1741758463067144.pdf"
        },
        "packing_slip" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/packing_slip/2025-03-12/ed92610f-75ee-4563-8ad3-5b6ddb1431bf-1741758469807818.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "183934-139324-34233" ],
      "carrier_references" : null,
      "carrier_account_slug" : "sf-express",
      "carrier_account_id" : "32146288eba64c809a2c3f4365c5aa5c",
      "service_type" : "sf-express_standard_express",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "SF Express Standard Express",
        "service_type" : "sf-express_standard_express",
        "shipper_account" : {
          "description" : "[sf-express] Demo Account can use",
          "id" : "32146288eba64c809a2c3f4365c5aa5c",
          "slug" : "sf-express"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 9.84
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 21,
        "height" : 52,
        "unit" : "cm",
        "width" : 65
      } ],
      "items" : [ {
        "description" : "D04010150006 - Wild Mugwort Restoring Serum\u0002Repairing",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1999,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "imac2014",
        "weight" : {
          "unit" : "kg",
          "value" : 9.54
        }
      }, {
        "description" : "Google Nexus 10",
        "origin_country" : "USA",
        "price" : {
          "amount" : 550,
          "currency" : "CNY"
        },
        "quantity" : 1,
        "sku" : "google man",
        "weight" : {
          "unit" : "kg",
          "value" : 0.3
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "739da68ad58d49a3bfdf56a0b52d5c53",
      "tracking_numbers" : null,
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-12T05:45:09.588239Z",
      "updated_at" : "2025-03-12T05:45:14.490776Z",
      "succeed_at" : "2025-03-12T05:45:13.161743223Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Salt Lake City",
        "contact_name" : "test nyy",
        "country" : "USA",
        "email" : "yy.nie@aftership.com",
        "phone" : "0123456789",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "phone" : "0123456789",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "a403c966b49d4561ba590ef8f9b1bb45",
      "order_number" : "#1024",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-12/739da68a-d58d-49a3-bfdf-56a0b52d5c53-1741758312951.pdf"
        },
        "packing_slip" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/packing_slip/2025-03-12/857539ed-6208-456a-bea8-f7dd963ee658-1741758314138155.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "test20250312001" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "gps",
      "carrier_account_id" : "a4f2c3c784424c4da0eb3ffefa582542",
      "service_type" : "gps_ups_ground",
      "rate" : {
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 6.6,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 4.720000000000001,
            "currency" : "USD"
          },
          "type" : "other"
        } ],
        "service_name" : "UPS® Ground",
        "service_type" : "gps_ups_ground",
        "shipper_account" : {
          "description" : "[gps] Demo Account",
          "id" : "a4f2c3c784424c4da0eb3ffefa582542",
          "slug" : "gps"
        },
        "total_charge" : {
          "amount" : 11.32,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "tracking_slug" : "ups"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 3.2
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 3,
        "height" : 5,
        "unit" : "cm",
        "width" : 4
      } ],
      "items" : [ {
        "description" : "苹果X XR玻璃手机壳",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/2407/9670/products/O1CN01PDsTa62MCKM0hpcrR__0-item_pic.jpg_430x430q90.jpg?v=1687770918" ],
        "item_id" : "16181170438454",
        "origin_country" : "USA",
        "price" : {
          "amount" : 14,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "9990098449929707830GHJKI UIYGGGUIIHHBVVVBHHJJJJJJJJJJBB",
        "weight" : {
          "unit" : "kg",
          "value" : 0.2
        }
      } ],
      "billing" : {
        "paid_by" : "recipient"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "93a4f9753fd74684ac538845e7ac2acd",
      "tracking_numbers" : [ "863T30049478A002" ],
      "organization_id" : "3fe7752813cb46b7baa82d0ba9ac5c96",
      "created_at" : "2025-03-12T05:37:33.258671Z",
      "updated_at" : "2025-03-12T05:37:39.246607Z",
      "succeed_at" : "2025-03-12T05:37:39.240295617Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "22303"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "94209"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "Alexandria",
        "company_name" : "weqwe",
        "contact_name" : "Jade 222",
        "country" : "USA",
        "email" : "yj.tang@aftership.com",
        "postal_code" : "22303",
        "state" : "VA",
        "street1" : "2000 Huntington Avenue"
      },
      "ship_to" : {
        "city" : "Sacramento",
        "company_name" : "21233",
        "contact_name" : "Jade",
        "country" : "USA",
        "phone" : "1-140-225-3341",
        "postal_code" : "94209",
        "state" : "CA",
        "street1" : "28292 Daugherty Orchard"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "9087f22e89024213adf77e88cec1a8a6",
      "order_number" : "#1021",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/93a4f975-3fd7-4684-ac53-8845e7ac2acd-1741757858877.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "qweqwewqVXOK1BXG", "wwring x 1", "eeee$27.00 USD" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "de3b9720dd86405b993eee6e8fd6f433",
      "service_type" : "yodel_returns",
      "rate" : {
        "booking_cut_off" : null,
        "extra_info" : null,
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account New2",
          "id" : "de3b9720dd86405b993eee6e8fd6f433",
          "slug" : "yodel"
        }
      },
      "service_options" : [ {
        "insured_value" : {
          "amount" : 27,
          "currency" : "USD"
        },
        "type" : "insurance"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "internal_custom_fields" : {
          "allocation_id" : "e7bcbd0db55440a6a20db7d4e2075303",
          "notify_customer" : "false",
          "operator_account_id" : "663c5c3139644f1481cbe9bfe5bf911c",
          "returns_shipping_type" : "retailer_label",
          "tracking_slug" : "yodel"
        },
        "shipper_account" : {
          "custom_fields" : {
            "label_qrcode_enabled" : true,
            "shipment_label_enabled" : false
          }
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 0.794375
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 2,
        "height" : 2,
        "unit" : "ft",
        "width" : 2
      } ],
      "items" : [ {
        "barcode" : "972019892255497201989225549720198922554",
        "description" : "ring",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0903/5956/2554/files/IMG_3935.jpg?v=1730181618" ],
        "item_id" : "15457596735802",
        "price" : {
          "amount" : 27,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "9720198922554",
        "weight" : {
          "unit" : "lb",
          "value" : 0.16975
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "159437e540e04c54a5fa2d2c16b25746",
      "tracking_numbers" : [ "863T30049475A002" ],
      "organization_id" : "3fe7752813cb46b7baa82d0ba9ac5c96",
      "created_at" : "2025-03-12T03:51:52.672143Z",
      "updated_at" : "2025-03-12T03:52:00.659962Z",
      "succeed_at" : "2025-03-12T03:52:00.60493349Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "AUS",
            "postal_code" : "6053"
          },
          "ship_to" : {
            "country" : "AUS",
            "postal_code" : "6053"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "Bayswater",
        "company_name" : "weqwe",
        "contact_name" : "Jade 222",
        "country" : "AUS",
        "email" : "yj.tang@aftership.com",
        "phone" : "1-800-818-971",
        "postal_code" : "6053",
        "state" : "WA",
        "street1" : "39 Bassendean Road"
      },
      "ship_to" : {
        "city" : "Bayswater",
        "company_name" : "21233",
        "contact_name" : "Jade",
        "country" : "AUS",
        "phone" : "1-140-225-3341",
        "postal_code" : "6053",
        "state" : "WA",
        "street1" : "39 Bassendean Road"
      },
      "ship_date" : "2025-03-12",
      "ship_from_country" : "AUS",
      "ship_to_country" : "AUS",
      "order_id" : "9087f22e89024213adf77e88cec1a8a6",
      "order_number" : "#1021",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-12/159437e5-40e0-4c54-a5fa-2d2c16b25746-1741751520197.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "qweqwewq0EEG2HBN", "wwring x 1", "eeee$27.00 USD" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "28ce7d875e6041ae883753a6f5578fcc",
      "service_type" : "yodel_returns",
      "rate" : {
        "booking_cut_off" : null,
        "extra_info" : null,
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Yodel Returns (print in store)",
        "service_type" : "yodel_returns",
        "shipper_account" : {
          "description" : "[yodel] Demo Account New",
          "id" : "28ce7d875e6041ae883753a6f5578fcc",
          "slug" : "yodel"
        }
      },
      "service_options" : [ {
        "insured_value" : {
          "amount" : 27,
          "currency" : "USD"
        },
        "type" : "insurance"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "internal_custom_fields" : {
          "allocation_id" : "7a4237f308ae41ad945a70bb2f4db44f",
          "notify_customer" : "false",
          "operator_account_id" : "663c5c3139644f1481cbe9bfe5bf911c",
          "returns_shipping_type" : "retailer_label",
          "tracking_slug" : "yodel"
        },
        "shipper_account" : {
          "custom_fields" : {
            "shipment_label_enabled" : true
          }
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 0.79475
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 24,
        "height" : 24,
        "unit" : "in",
        "width" : 24
      } ],
      "items" : [ {
        "barcode" : "972019892255497201989225549720198922554",
        "description" : "ring",
        "hs_code" : "hs_code",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0903/5956/2554/files/IMG_3935.jpg?v=1730181618" ],
        "item_id" : "15457596735802",
        "origin_country" : "AUS",
        "price" : {
          "amount" : 27,
          "currency" : "AUD"
        },
        "quantity" : 1,
        "sku" : "9720198922554",
        "weight" : {
          "unit" : "lb",
          "value" : 0.17
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "50627d779ac943c09d0b6b89210616ac",
      "tracking_numbers" : [ "5332741772" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T11:37:26.927679Z",
      "updated_at" : "2025-03-11T11:37:32.406125Z",
      "succeed_at" : "2025-03-11T11:37:31.000336226Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "GBR",
            "postal_code" : "CO3 0PP"
          },
          "ship_to" : {
            "country" : "CAN",
            "postal_code" : "M6A 1P6"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Colchester",
        "contact_name" : "Courtney Elliott",
        "country" : "GBR",
        "email" : "t@t.com",
        "phone" : "416-306-8001",
        "postal_code" : "CO3 0PP",
        "state" : "England",
        "street1" : "19 Spring Sedge Close"
      },
      "ship_to" : {
        "city" : "North York",
        "company_name" : "UK-MDC",
        "contact_name" : "UK-MDC",
        "country" : "CAN",
        "phone" : "416-306-8001",
        "postal_code" : "M6A 1P6",
        "state" : "ON",
        "street1" : "105 Bentworth Avenue",
        "street2" : "Unit 4"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "GBR",
      "ship_to_country" : "CAN",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "invoice" : {
          "file_type" : "invoice",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-11/50627d77-9ac9-43c0-9d0b-6b89210616ac-1741693050790.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x8",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/50627d77-9ac9-43c0-9d0b-6b89210616ac-1741693050790.pdf"
        },
        "packing_slip" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/packing_slip/2025-03-11/a3137d9e-fde1-4056-bc92-512a9b011c5d-1741693052138063.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA REFLSL4PN7U", "ORDER REFM192313764UZ" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "dhl",
      "carrier_account_id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
      "service_type" : "dhl_express_worldwide",
      "rate" : {
        "booking_cut_off" : "2025-03-12T22:00:00+08:00",
        "charge_weight" : {
          "unit" : "kg",
          "value" : 1.5
        },
        "delivery_date" : "2025-03-15T07:59:00+08:00",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 205.21,
            "currency" : "EUR"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 61.05,
            "currency" : "EUR"
          },
          "type" : "fuel_surcharge"
        } ],
        "pickup_deadline" : "2025-03-13T00:00:00+08:00",
        "service_name" : "DHL Express Worldwide",
        "service_type" : "dhl_express_worldwide",
        "shipper_account" : {
          "description" : "[mydhl] Demo Account [Automate Test]",
          "id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
          "slug" : "dhl"
        },
        "total_charge" : {
          "amount" : 266.26,
          "currency" : "EUR"
        },
        "transit_time" : 2
      },
      "service_options" : [ {
        "enabled" : true,
        "type" : "paperless_invoice"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "description" : "",
            "dimension" : {
              "depth" : 2.1,
              "height" : 2.1,
              "unit" : "cm",
              "width" : 2.1
            },
            "items" : [ {
              "barcode" : "",
              "description" : "Single Mini Hoop - 14k Yellow Gold",
              "description_cn" : "",
              "dimension" : {
                "depth" : 0,
                "height" : 0,
                "unit" : "",
                "width" : 0
              },
              "hs_code" : "7113195090",
              "image_urls" : [ "https://cdn.shopify.com/s/files/1/0837/7489/8461/files/0-SingleMiniHoop-14K-Angled_045_50e57c60-f395-4735-be7c-58865a504acb.jpg?v=1722958121" ],
              "item_id" : "15992277532957",
              "origin_country" : "CAN",
              "price" : {
                "amount" : 31.67,
                "currency" : "GBP"
              },
              "quantity" : 1,
              "return_reason" : "",
              "sku" : "p52704669",
              "weight" : {
                "unit" : "lb",
                "value" : 6.875E-4
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 1.011
            }
          } ]
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.011
      },
      "chargeable_weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "dimension" : [ {
        "depth" : 2.1,
        "height" : 2.1,
        "unit" : "cm",
        "width" : 2.1
      } ],
      "items" : [ {
        "description" : "Single Mini Hoop - 14k Yellow Gold",
        "hs_code" : "7113195090",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0837/7489/8461/files/0-SingleMiniHoop-14K-Angled_045_50e57c60-f395-4735-be7c-58865a504acb.jpg?v=1722958121" ],
        "item_id" : "15992277532957",
        "origin_country" : "CAN",
        "price" : {
          "amount" : 31.67,
          "currency" : "GBP"
        },
        "quantity" : 1,
        "sku" : "p52704669",
        "weight" : {
          "unit" : "lb",
          "value" : 6.875E-4
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "62feace16df14fd2ab6d525c7d9750ed",
      "tracking_numbers" : [ "5332741761" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T11:37:09.425503Z",
      "updated_at" : "2025-03-11T11:37:15.052105Z",
      "succeed_at" : "2025-03-11T11:37:14.09091991Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "GBR",
            "postal_code" : "CO3 0PP"
          },
          "ship_to" : {
            "country" : "CAN",
            "postal_code" : "M6A 1P6"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Colchester",
        "contact_name" : "Courtney Elliott",
        "country" : "GBR",
        "email" : "t@t.com",
        "phone" : "416-306-8001",
        "postal_code" : "CO3 0PP",
        "state" : "England",
        "street1" : "19 Spring Sedge Close"
      },
      "ship_to" : {
        "city" : "North York",
        "company_name" : "UK-MDC",
        "contact_name" : "UK-MDC",
        "country" : "CAN",
        "phone" : "416-306-8001",
        "postal_code" : "M6A 1P6",
        "state" : "ON",
        "street1" : "105 Bentworth Avenue",
        "street2" : "Unit 4"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "GBR",
      "ship_to_country" : "CAN",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "invoice" : {
          "file_type" : "invoice",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-11/62feace1-6df1-4fd2-ab6d-525c7d9750ed-1741693033868.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x8",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/62feace1-6df1-4fd2-ab6d-525c7d9750ed-1741693033916.pdf"
        },
        "packing_slip" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/packing_slip/2025-03-11/e2f01466-60c5-4379-b855-f13498c8ab26-1741693034751177.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA REFLSL4PN7U", "ORDER REFM192313764UZ" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "dhl",
      "carrier_account_id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
      "service_type" : "dhl_express_worldwide",
      "rate" : {
        "booking_cut_off" : "2025-03-12T22:00:00+08:00",
        "charge_weight" : {
          "unit" : "kg",
          "value" : 1.5
        },
        "delivery_date" : "2025-03-15T07:59:00+08:00",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 205.21,
            "currency" : "EUR"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 61.05,
            "currency" : "EUR"
          },
          "type" : "fuel_surcharge"
        } ],
        "pickup_deadline" : "2025-03-13T00:00:00+08:00",
        "service_name" : "DHL Express Worldwide",
        "service_type" : "dhl_express_worldwide",
        "shipper_account" : {
          "description" : "[mydhl] Demo Account [Automate Test]",
          "id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
          "slug" : "dhl"
        },
        "total_charge" : {
          "amount" : 266.26,
          "currency" : "EUR"
        },
        "transit_time" : 2
      },
      "service_options" : [ {
        "enabled" : true,
        "type" : "paperless_invoice"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "description" : "",
            "dimension" : {
              "depth" : 0.1,
              "height" : 0.1,
              "unit" : "cm",
              "width" : 0.1
            },
            "items" : [ {
              "barcode" : "",
              "description" : "Single Mini Hoop - 14k Yellow Gold",
              "description_cn" : "",
              "dimension" : {
                "depth" : 0,
                "height" : 0,
                "unit" : "",
                "width" : 0
              },
              "hs_code" : "7113195090",
              "image_urls" : [ "https://cdn.shopify.com/s/files/1/0837/7489/8461/files/0-SingleMiniHoop-14K-Angled_045_50e57c60-f395-4735-be7c-58865a504acb.jpg?v=1722958121" ],
              "item_id" : "15992277532957",
              "origin_country" : "CAN",
              "price" : {
                "amount" : 31.67,
                "currency" : "GBP"
              },
              "quantity" : 1,
              "return_reason" : "",
              "sku" : "p52704669",
              "weight" : {
                "unit" : "lb",
                "value" : 6.875E-4
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 1.011
            }
          } ]
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.011
      },
      "chargeable_weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "dimension" : [ {
        "depth" : 0.1,
        "height" : 0.1,
        "unit" : "cm",
        "width" : 0.1
      } ],
      "items" : [ {
        "description" : "Single Mini Hoop - 14k Yellow Gold",
        "hs_code" : "7113195090",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0837/7489/8461/files/0-SingleMiniHoop-14K-Angled_045_50e57c60-f395-4735-be7c-58865a504acb.jpg?v=1722958121" ],
        "item_id" : "15992277532957",
        "origin_country" : "CAN",
        "price" : {
          "amount" : 31.67,
          "currency" : "GBP"
        },
        "quantity" : 1,
        "sku" : "p52704669",
        "weight" : {
          "unit" : "lb",
          "value" : 6.875E-4
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "cb711e15af6f4362b44cbdae1098b779",
      "tracking_numbers" : [ "5332741724" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T11:36:22.583407Z",
      "updated_at" : "2025-03-11T11:36:28.965894Z",
      "succeed_at" : "2025-03-11T11:36:27.469727207Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "GBR",
            "postal_code" : "CO3 0PP"
          },
          "ship_to" : {
            "country" : "CAN",
            "postal_code" : "M6A 1P6"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Colchester",
        "contact_name" : "Courtney Elliott",
        "country" : "GBR",
        "email" : "t@t.com",
        "phone" : "416-306-8001",
        "postal_code" : "CO3 0PP",
        "state" : "England",
        "street1" : "19 Spring Sedge Close"
      },
      "ship_to" : {
        "city" : "North York",
        "company_name" : "UK-MDC",
        "contact_name" : "UK-MDC",
        "country" : "CAN",
        "phone" : "416-306-8001",
        "postal_code" : "M6A 1P6",
        "state" : "ON",
        "street1" : "105 Bentworth Avenue",
        "street2" : "Unit 4"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "GBR",
      "ship_to_country" : "CAN",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "invoice" : {
          "file_type" : "invoice",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-11/cb711e15-af6f-4362-b44c-bdae1098b779-1741692987208.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x8",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/cb711e15-af6f-4362-b44c-bdae1098b779-1741692987213.pdf"
        },
        "packing_slip" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/packing_slip/2025-03-11/8a4592e3-62bf-4b61-b009-ecedbbca6039-1741692988675352.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA REFLSL4PN7U", "ORDER REFM192313764UZ" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "dhl",
      "carrier_account_id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
      "service_type" : "dhl_express_worldwide",
      "rate" : {
        "booking_cut_off" : "2025-03-12T22:00:00+08:00",
        "charge_weight" : {
          "unit" : "kg",
          "value" : 1.5
        },
        "delivery_date" : "2025-03-15T07:59:00+08:00",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 205.21,
            "currency" : "EUR"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 61.05,
            "currency" : "EUR"
          },
          "type" : "fuel_surcharge"
        } ],
        "pickup_deadline" : "2025-03-13T00:00:00+08:00",
        "service_name" : "DHL Express Worldwide",
        "service_type" : "dhl_express_worldwide",
        "shipper_account" : {
          "description" : "[mydhl] Demo Account [Automate Test]",
          "id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
          "slug" : "dhl"
        },
        "total_charge" : {
          "amount" : 266.26,
          "currency" : "EUR"
        },
        "transit_time" : 2
      },
      "service_options" : [ {
        "enabled" : true,
        "type" : "paperless_invoice"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "description" : "",
            "dimension" : {
              "depth" : 0.1,
              "height" : 0.1,
              "unit" : "cm",
              "width" : 0.1
            },
            "items" : [ {
              "barcode" : "",
              "description" : "Single Mini Hoop - 14k Yellow Gold",
              "description_cn" : "",
              "dimension" : {
                "depth" : 0,
                "height" : 0,
                "unit" : "",
                "width" : 0
              },
              "hs_code" : "7113195090",
              "image_urls" : [ "https://cdn.shopify.com/s/files/1/0837/7489/8461/files/0-SingleMiniHoop-14K-Angled_045_50e57c60-f395-4735-be7c-58865a504acb.jpg?v=1722958121" ],
              "item_id" : "15992277532957",
              "origin_country" : "CAN",
              "price" : {
                "amount" : 31.67,
                "currency" : "GBP"
              },
              "quantity" : 1,
              "return_reason" : "",
              "sku" : "p52704669",
              "weight" : {
                "unit" : "lb",
                "value" : 6.875E-4
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 1.011
            }
          } ]
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.011
      },
      "chargeable_weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "dimension" : [ {
        "depth" : 0.1,
        "height" : 0.1,
        "unit" : "cm",
        "width" : 0.1
      } ],
      "items" : [ {
        "description" : "Single Mini Hoop - 14k Yellow Gold",
        "hs_code" : "7113195090",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0837/7489/8461/files/0-SingleMiniHoop-14K-Angled_045_50e57c60-f395-4735-be7c-58865a504acb.jpg?v=1722958121" ],
        "item_id" : "15992277532957",
        "origin_country" : "CAN",
        "price" : {
          "amount" : 31.67,
          "currency" : "GBP"
        },
        "quantity" : 1,
        "sku" : "p52704669",
        "weight" : {
          "unit" : "lb",
          "value" : 6.875E-4
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "0f09cd20b0ea443896e5e75624828c30",
      "tracking_numbers" : [ "5332741702" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T11:35:15.991599Z",
      "updated_at" : "2025-03-11T11:35:23.586671Z",
      "succeed_at" : "2025-03-11T11:35:21.877676825Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "GBR",
            "postal_code" : "CO3 0PP"
          },
          "ship_to" : {
            "country" : "CAN",
            "postal_code" : "M6A 1P6"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Colchester",
        "contact_name" : "Courtney Elliott",
        "country" : "GBR",
        "email" : "t@t.com",
        "phone" : "416-306-8001",
        "postal_code" : "CO3 0PP",
        "state" : "England",
        "street1" : "19 Spring Sedge Close"
      },
      "ship_to" : {
        "city" : "North York",
        "company_name" : "UK-MDC",
        "contact_name" : "UK-MDC",
        "country" : "CAN",
        "phone" : "416-306-8001",
        "postal_code" : "M6A 1P6",
        "state" : "ON",
        "street1" : "105 Bentworth Avenue",
        "street2" : "Unit 4"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "GBR",
      "ship_to_country" : "CAN",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "invoice" : {
          "file_type" : "invoice",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-11/0f09cd20-b0ea-4438-96e5-e75624828c30-1741692921653.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x8",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/0f09cd20-b0ea-4438-96e5-e75624828c30-1741692921653.pdf"
        },
        "packing_slip" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/packing_slip/2025-03-11/60f88f6b-044c-4194-a242-972edc42e9bc-1741692923231169.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA REFLSL4PN7U", "ORDER REFM192313764UZ" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "dhl",
      "carrier_account_id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
      "service_type" : "dhl_express_worldwide",
      "rate" : {
        "booking_cut_off" : "2025-03-12T22:00:00+08:00",
        "charge_weight" : {
          "unit" : "kg",
          "value" : 1.5
        },
        "delivery_date" : "2025-03-15T07:59:00+08:00",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 205.21,
            "currency" : "EUR"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 61.05,
            "currency" : "EUR"
          },
          "type" : "fuel_surcharge"
        } ],
        "pickup_deadline" : "2025-03-13T00:00:00+08:00",
        "service_name" : "DHL Express Worldwide",
        "service_type" : "dhl_express_worldwide",
        "shipper_account" : {
          "description" : "[mydhl] Demo Account [Automate Test]",
          "id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
          "slug" : "dhl"
        },
        "total_charge" : {
          "amount" : 266.26,
          "currency" : "EUR"
        },
        "transit_time" : 2
      },
      "service_options" : [ {
        "enabled" : true,
        "type" : "paperless_invoice"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "description" : "",
            "dimension" : {
              "depth" : 0.1,
              "height" : 0.1,
              "unit" : "cm",
              "width" : 0.1
            },
            "items" : [ {
              "barcode" : "",
              "description" : "Single Mini Hoop - 14k Yellow Gold",
              "description_cn" : "",
              "dimension" : {
                "depth" : 0,
                "height" : 0,
                "unit" : "",
                "width" : 0
              },
              "hs_code" : "7113195090",
              "image_urls" : [ "https://cdn.shopify.com/s/files/1/0837/7489/8461/files/0-SingleMiniHoop-14K-Angled_045_50e57c60-f395-4735-be7c-58865a504acb.jpg?v=1722958121" ],
              "item_id" : "15992277532957",
              "origin_country" : "CAN",
              "price" : {
                "amount" : 31.67,
                "currency" : "GBP"
              },
              "quantity" : 1,
              "return_reason" : "",
              "sku" : "p52704669",
              "weight" : {
                "unit" : "lb",
                "value" : 6.875E-4
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 1.011
            }
          } ]
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.011
      },
      "chargeable_weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "dimension" : [ {
        "depth" : 0.1,
        "height" : 0.1,
        "unit" : "cm",
        "width" : 0.1
      } ],
      "items" : [ {
        "description" : "Single Mini Hoop - 14k Yellow Gold",
        "hs_code" : "7113195090",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0837/7489/8461/files/0-SingleMiniHoop-14K-Angled_045_50e57c60-f395-4735-be7c-58865a504acb.jpg?v=1722958121" ],
        "item_id" : "15992277532957",
        "origin_country" : "CAN",
        "price" : {
          "amount" : 31.67,
          "currency" : "GBP"
        },
        "quantity" : 1,
        "sku" : "p52704669",
        "weight" : {
          "unit" : "lb",
          "value" : 6.875E-4
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "35cdfef1dd6e429b9f3dc553a257b78a",
      "tracking_numbers" : [ "111JD790798601000931504" ],
      "organization_id" : "f64e2c3030584d84ba88fea2b1151c3b",
      "created_at" : "2025-03-11T10:02:24.171615Z",
      "updated_at" : "2025-03-11T10:02:30.987903Z",
      "succeed_at" : "2025-03-11T10:02:30.978406456Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "AUS",
            "postal_code" : "2770"
          },
          "ship_to" : {
            "country" : "AUS",
            "postal_code" : "2770"
          }
        },
        "retention_days" : "365",
        "source" : "admin"
      },
      "source" : "admin",
      "ship_from" : {
        "city" : "Whalan",
        "company_name" : "c",
        "contact_name" : "l",
        "country" : "AUS",
        "email" : "test@test.com",
        "id" : "f9a0f7c5fdf74414bc568e9133be3b21",
        "phone" : "430412345678",
        "postal_code" : "2770",
        "state" : "NSW",
        "street1" : "10 Ellengowan Crescent",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Whalan",
        "company_name" : "c",
        "contact_name" : "l",
        "country" : "AUS",
        "email" : "test@test.com",
        "phone" : "430412345678",
        "postal_code" : "2770",
        "state" : "NSW",
        "street1" : "10 Ellengowan Crescent",
        "type" : "business"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "AUS",
      "ship_to_country" : "AUS",
      "order_id" : "5b4d271139d34144bc0e70ed4296f27c",
      "order_number" : "000146",
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-11/cd6be800-17c9-499f-bfda-135f15c55ebc-1741687348641614.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/b2db7b1e-1886-47dd-938d-5cb02a2b9540-1741687348181137.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "000146" ],
      "carrier_references" : null,
      "carrier_account_slug" : "australia-post",
      "carrier_account_id" : "0154af62-4aba-42a2-8381-902b7c23ed01",
      "service_type" : "australia-post_parcel_post_sign",
      "rate" : {
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 21.150000000000002,
            "currency" : "AUD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 2.11,
            "currency" : "AUD"
          },
          "type" : "tax"
        } ],
        "service_name" : "Parcel Post (Signature)",
        "service_type" : "australia-post_parcel_post_sign",
        "shipper_account" : {
          "description" : "[Australia Post] Postmen Testing Account",
          "id" : "0154af62-4aba-42a2-8381-902b7c23ed01",
          "slug" : "australia-post"
        },
        "total_charge" : {
          "amount" : 23.26,
          "currency" : "AUD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : {
          "amount" : 23.26,
          "currency" : "AUD"
        },
        "internal_custom_fields" : {
          "account_id" : "5bee4d23e2bb4da9bd6fa82cb7f29768",
          "ship_from" : {
            "id" : "f9a0f7c5fdf74414bc568e9133be3b21"
          }
        },
        "shipment_id" : "gmoK0EhGusMAAAGVYM8CD4Sn"
      },
      "custom_fields" : {
        "location_name" : "AUS"
      },
      "weight" : {
        "unit" : "lb",
        "value" : 5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 5,
        "height" : 5,
        "unit" : "cm",
        "width" : 5
      } ],
      "items" : [ {
        "description" : "Auto-generated by Postmen",
        "price" : {
          "amount" : 1,
          "currency" : "USD"
        },
        "quantity" : 1,
        "weight" : {
          "unit" : "lb",
          "value" : 5
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "5bee4d23e2bb4da9bd6fa82cb7f29768",
      "label_base64" : null
    }, {
      "id" : "694803389a89472f9fce4bb9b25ff29b",
      "tracking_numbers" : [ "111JD790798501000931507" ],
      "organization_id" : "f64e2c3030584d84ba88fea2b1151c3b",
      "created_at" : "2025-03-11T10:01:44.438207Z",
      "updated_at" : "2025-03-11T10:01:51.415702Z",
      "succeed_at" : "2025-03-11T10:01:51.404539941Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "AUS",
            "postal_code" : "2770"
          },
          "ship_to" : {
            "country" : "AUS",
            "postal_code" : "2770"
          }
        },
        "retention_days" : "365",
        "source" : "admin"
      },
      "source" : "admin",
      "ship_from" : {
        "city" : "Whalan",
        "company_name" : "c",
        "contact_name" : "l",
        "country" : "AUS",
        "email" : "test@test.com",
        "id" : "f9a0f7c5fdf74414bc568e9133be3b21",
        "phone" : "430412345678",
        "postal_code" : "2770",
        "state" : "NSW",
        "street1" : "10 Ellengowan Crescent",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Whalan",
        "company_name" : "c",
        "contact_name" : "l",
        "country" : "AUS",
        "email" : "test@test.com",
        "phone" : "430412345678",
        "postal_code" : "2770",
        "state" : "NSW",
        "street1" : "10 Ellengowan Crescent",
        "type" : "business"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "AUS",
      "ship_to_country" : "AUS",
      "order_id" : "7e0a54aa69ef411b95f4870568acb31a",
      "order_number" : "000145",
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-11/1e2f515c-c110-4619-9f6e-566642e9edcf-1741687308925119.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/13a10f67-6167-411f-911c-a049fb2bb187-1741687308447612.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "000145" ],
      "carrier_references" : null,
      "carrier_account_slug" : "australia-post",
      "carrier_account_id" : "0154af62-4aba-42a2-8381-902b7c23ed01",
      "service_type" : "australia-post_parcel_post_sign",
      "rate" : {
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 21.150000000000002,
            "currency" : "AUD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 2.11,
            "currency" : "AUD"
          },
          "type" : "tax"
        } ],
        "service_name" : "Parcel Post (Signature)",
        "service_type" : "australia-post_parcel_post_sign",
        "shipper_account" : {
          "description" : "[Australia Post] Postmen Testing Account",
          "id" : "0154af62-4aba-42a2-8381-902b7c23ed01",
          "slug" : "australia-post"
        },
        "total_charge" : {
          "amount" : 23.26,
          "currency" : "AUD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : {
          "amount" : 23.26,
          "currency" : "AUD"
        },
        "internal_custom_fields" : {
          "account_id" : "5bee4d23e2bb4da9bd6fa82cb7f29768",
          "ship_from" : {
            "id" : "f9a0f7c5fdf74414bc568e9133be3b21"
          }
        },
        "shipment_id" : "etAK0EVqSwsAAAGVbDQCEISn"
      },
      "custom_fields" : {
        "location_name" : "AUS"
      },
      "weight" : {
        "unit" : "lb",
        "value" : 5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 5,
        "height" : 5,
        "unit" : "cm",
        "width" : 5
      } ],
      "items" : [ {
        "description" : "Auto-generated by Postmen",
        "price" : {
          "amount" : 1,
          "currency" : "USD"
        },
        "quantity" : 1,
        "weight" : {
          "unit" : "lb",
          "value" : 5
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "5bee4d23e2bb4da9bd6fa82cb7f29768",
      "label_base64" : null
    }, {
      "id" : "96fc586965ea4eb38f07d1960259dfa7",
      "tracking_numbers" : [ "111JD790798201000931506" ],
      "organization_id" : "f64e2c3030584d84ba88fea2b1151c3b",
      "created_at" : "2025-03-11T09:58:02.968351Z",
      "updated_at" : "2025-03-11T09:58:09.705504Z",
      "succeed_at" : "2025-03-11T09:58:09.696496108Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "AUS",
            "postal_code" : "2770"
          },
          "ship_to" : {
            "country" : "AUS",
            "postal_code" : "2770"
          }
        },
        "retention_days" : "365",
        "source" : "admin"
      },
      "source" : "admin",
      "ship_from" : {
        "city" : "Whalan",
        "company_name" : "c",
        "contact_name" : "l",
        "country" : "AUS",
        "email" : "test@test.com",
        "id" : "f9a0f7c5fdf74414bc568e9133be3b21",
        "phone" : "430412345678",
        "postal_code" : "2770",
        "state" : "NSW",
        "street1" : "10 Ellengowan Crescent",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Whalan",
        "company_name" : "c",
        "contact_name" : "l",
        "country" : "AUS",
        "email" : "test@test.com",
        "phone" : "430412345678",
        "postal_code" : "2770",
        "state" : "NSW",
        "street1" : "10 Ellengowan Crescent",
        "type" : "business"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "AUS",
      "ship_to_country" : "AUS",
      "order_id" : "5bd046534f9d46519aab43148271b67e",
      "order_number" : "000144",
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-11/132e9434-955d-4b92-91d8-40813c9053c4-1741687087611533.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/1167f6db-19ba-4a15-a791-0db60fc461ce-1741687087006188.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "000144" ],
      "carrier_references" : null,
      "carrier_account_slug" : "australia-post",
      "carrier_account_id" : "0154af62-4aba-42a2-8381-902b7c23ed01",
      "service_type" : "australia-post_parcel_post_sign",
      "rate" : {
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 21.150000000000002,
            "currency" : "AUD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 2.11,
            "currency" : "AUD"
          },
          "type" : "tax"
        } ],
        "service_name" : "Parcel Post (Signature)",
        "service_type" : "australia-post_parcel_post_sign",
        "shipper_account" : {
          "description" : "[Australia Post] Postmen Testing Account",
          "id" : "0154af62-4aba-42a2-8381-902b7c23ed01",
          "slug" : "australia-post"
        },
        "total_charge" : {
          "amount" : 23.26,
          "currency" : "AUD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : {
          "amount" : 23.26,
          "currency" : "AUD"
        },
        "internal_custom_fields" : {
          "account_id" : "5bee4d23e2bb4da9bd6fa82cb7f29768",
          "ship_from" : {
            "id" : "f9a0f7c5fdf74414bc568e9133be3b21"
          }
        },
        "shipment_id" : "0awK0EhGKW4AAAGV3tICD4Sj"
      },
      "custom_fields" : {
        "location_name" : "AUS"
      },
      "weight" : {
        "unit" : "lb",
        "value" : 5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 5,
        "height" : 5,
        "unit" : "cm",
        "width" : 5
      } ],
      "items" : [ {
        "description" : "Auto-generated by Postmen",
        "price" : {
          "amount" : 1,
          "currency" : "USD"
        },
        "quantity" : 1,
        "weight" : {
          "unit" : "lb",
          "value" : 5
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "5bee4d23e2bb4da9bd6fa82cb7f29768",
      "label_base64" : null
    }, {
      "id" : "349bcd274b5240bfbadd0d59ead52242",
      "tracking_numbers" : [ "9234690110067000000017" ],
      "organization_id" : "e430e307533f44f58cf8ef3d3b0001b7",
      "created_at" : "2025-03-11T09:20:57.325231Z",
      "updated_at" : "2025-03-11T09:20:59.382927Z",
      "succeed_at" : "2025-03-11T09:20:59.372430085Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84116"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "10007"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Salt Lake City",
        "company_name" : "Test from company",
        "contact_name" : "test shipper",
        "country" : "USA",
        "email" : "tx.test@test.com",
        "phone" : "6285201311234",
        "postal_code" : "84116",
        "state" : "UT",
        "street1" : "5220 Wiley11 Post Way",
        "street2" : "Building 3",
        "street3" : "Reserved for international addresses",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "New York",
        "company_name" : "Test to company",
        "contact_name" : "Test consignee",
        "country" : "USA",
        "email" : "tx.test@test.com",
        "phone" : "6285201315678",
        "postal_code" : "10007",
        "state" : "UT",
        "street1" : "5220 Wiley11 Post Way",
        "street2" : "Building 3",
        "street3" : "Reserved for international addresses",
        "type" : "business"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : "test250311001",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/349bcd27-4b52-40bf-badd-0d59ead52242-1741684859158.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "testorder1030", "billingRef1", "billingRef2" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "usps",
      "carrier_account_id" : "6ece87ac67254b42a46a4d4f1e8aee4f",
      "service_type" : "usps_ground_advantage",
      "rate" : {
        "booking_cut_off" : null,
        "charge_weight" : {
          "unit" : "lb",
          "value" : 1.7636980974790206
        },
        "delivery_date" : "2025-03-15",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 10.85,
            "currency" : "USD"
          },
          "type" : "base"
        } ],
        "extra_info" : null,
        "service_name" : "USPS Ground Advantage",
        "service_type" : "usps_ground_advantage",
        "shipper_account" : {
          "description" : "usps 20241225",
          "id" : "6ece87ac67254b42a46a4d4f1e8aee4f",
          "slug" : "usps"
        },
        "total_charge" : {
          "amount" : 10.85,
          "currency" : "USD"
        },
        "transit_time" : 4
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "is_domestic" : true
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 0.8
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 1.7636980974790206
      },
      "dimension" : [ {
        "depth" : 10,
        "height" : 10,
        "unit" : "cm",
        "width" : 10
      } ],
      "items" : [ {
        "description" : "T-shirt  - 40",
        "price" : {
          "amount" : 184.79,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "GW2871SP40",
        "weight" : {
          "unit" : "kg",
          "value" : 0.8
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "release-incy-sandbox-api.aftership.io",
      "label_base64" : null
    }, {
      "id" : "2f38f69fbef74b1281696181d6f8773a",
      "tracking_numbers" : [ "420841019205590383191000000397" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T08:15:35.688399Z",
      "updated_at" : "2025-03-11T08:15:43.930503Z",
      "succeed_at" : "2025-03-11T08:15:42.809033883Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Salt Lake City",
        "contact_name" : "test nyy",
        "country" : "USA",
        "email" : "yy.nie@aftership.com",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "phone" : "0123456789",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "a403c966b49d4561ba590ef8f9b1bb45",
      "order_number" : "#1024",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/2f38f69f-bef7-4b12-8169-6181d6f8773a-1741680942597.pdf"
        },
        "packing_slip" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/packing_slip/2025-03-11/abb7ca89-a7bb-4350-aea2-a011feadba0f-1741680943573209.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA#7MRPIEW1 Order##1024", "苹果X XR玻璃手机壳 x 1 Look", "13123万231212#W" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "gps",
      "carrier_account_id" : "a4f2c3c784424c4da0eb3ffefa582542",
      "service_type" : "gps_usps_priority_mail",
      "rate" : {
        "detailed_charges" : {
          "charge" : {
            "amount" : 12.93,
            "currency" : "USD"
          },
          "type" : "base"
        },
        "service_name" : "USPS Priority Mail",
        "service_type" : "gps_usps_priority_mail",
        "shipper_account" : {
          "description" : "[gps] Demo Account",
          "id" : "a4f2c3c784424c4da0eb3ffefa582542",
          "slug" : "gps"
        },
        "total_charge" : {
          "amount" : 12.93,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "tracking_slug" : "usps"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 3.2
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 3,
        "height" : 5,
        "unit" : "cm",
        "width" : 4
      } ],
      "items" : [ {
        "description" : "苹果X XR玻璃手机壳",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/2407/9670/products/O1CN01PDsTa62MCKM0hpcrR__0-item_pic.jpg_430x430q90.jpg?v=1687770918" ],
        "item_id" : "16181170438454",
        "origin_country" : "USA",
        "price" : {
          "amount" : 14,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "9990098449929707830GHJKI UIYGGGUIIHHBVVVBHHJJJJJJJJJJBB",
        "weight" : {
          "unit" : "kg",
          "value" : 0.2
        }
      } ],
      "billing" : {
        "paid_by" : "recipient"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "65746333153c467f8809c5bc68fadfde",
      "tracking_numbers" : [ "795495355726" ],
      "organization_id" : "251aabe48d9f45a3975d8413eeb77001",
      "created_at" : "2025-03-11T07:58:37.010703Z",
      "updated_at" : "2025-03-11T07:58:48.487073Z",
      "succeed_at" : "2025-03-11T07:58:48.476244458Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "11238"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "94209"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "Brooklyn",
        "contact_name" : "rikki rikki",
        "country" : "USA",
        "email" : "lj.chen@aftership.com",
        "phone" : "+498912661098",
        "postal_code" : "11238",
        "state" : "NY",
        "street1" : "200 Eastern Parkway"
      },
      "ship_to" : {
        "city" : "Sacramento",
        "company_name" : "Zhou",
        "contact_name" : "HanFu",
        "country" : "USA",
        "phone" : "+498912661098",
        "postal_code" : "94209",
        "state" : "CA",
        "street1" : "28292 Daugherty Orchard"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "c84362b267f9444e8eb507896d3cee11",
      "order_number" : "#1175",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/dbbebcfc-fed0-4944-b9a3-824261d4578a-1741679927510641.pdf"
        }
      },
      "box_type" : "fedex_large_box",
      "references" : [ "RMA#HKQ7JH3D Order##1175", "Bundle- partial A x 2 fedex" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
      "service_type" : "fedex_2_day",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "[FedEx] Postmen Testing Account",
          "id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
          "slug" : "fedex"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null,
        "form_id" : "0221",
        "internal_custom_fields" : {
          "allocation_id" : "d9ef5b96f9b946cebcb2c5403ad0166c",
          "notify_customer" : "true",
          "operator_account_id" : "76df5a8994b44af09d175abcae8a7bf2",
          "returns_shipping_type" : "retailer_label",
          "tracking_slug" : "fedex"
        },
        "shipment" : {
          "parcels" : [ {
            "box_type" : "fedex_large_box",
            "dimension" : {
              "depth" : 10,
              "height" : 10,
              "unit" : "cm",
              "width" : 10
            },
            "items" : [ {
              "description" : "Bundle- partial A",
              "image_urls" : [ "https://cdn.shopify.com/s/files/1/0636/0046/0006/products/d9991903a19ff565705e83dc645e665a.jpg?v=1677059765" ],
              "item_id" : "15389495034086",
              "origin_country" : null,
              "price" : {
                "amount" : 200,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "Bundle- partial A",
              "weight" : {
                "unit" : "lb",
                "value" : 2.205
              }
            } ],
            "weight" : {
              "unit" : "lb",
              "value" : 14.41
            }
          } ]
        },
        "shipper_account" : {
          "custom_fields" : {
            "shipment_label_enabled" : true
          }
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 14.41
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 10,
        "height" : 10,
        "unit" : "cm",
        "width" : 10
      } ],
      "items" : [ {
        "description" : "Bundle- partial A",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0636/0046/0006/products/d9991903a19ff565705e83dc645e665a.jpg?v=1677059765" ],
        "item_id" : "15389495034086",
        "price" : {
          "amount" : 200,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "Bundle- partial A",
        "weight" : {
          "unit" : "lb",
          "value" : 2.205
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "ee37eb77625f4af1b9eaddca63c15971",
      "tracking_numbers" : null,
      "organization_id" : "6c70bde5fe4f4230ae133c68fb2330b3",
      "created_at" : "2025-03-11T07:48:22.003359Z",
      "updated_at" : "2025-03-11T07:48:23.243657Z",
      "succeed_at" : "2025-03-11T07:48:23.23451193Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "12231"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "12231"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "New York",
        "contact_name" : "huanglynne",
        "country" : "USA",
        "email" : "jl.huang+us@aftership.com",
        "postal_code" : "12231",
        "state" : "NY",
        "street1" : "New York"
      },
      "ship_to" : {
        "city" : "New York",
        "contact_name" : "huanglynne",
        "country" : "USA",
        "email" : "jl.huang@aftership.com",
        "postal_code" : "12231",
        "state" : "NY",
        "street1" : "New York"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "37c414d77194430dbc909955b1e716f8",
      "order_number" : "9P66F496",
      "return_shipment" : true,
      "files" : {
        "qr_code" : {
          "file_type" : "png",
          "paper_size" : "2x2",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-11/ee37eb77-625f-4af1-b9ea-ddca63c15971-1741679302967.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA#9P66F496 Order##tax-exclude1239", "exclude tax = 10% bundle Partial e" ],
      "carrier_references" : [ {
        "dropoff_number" : "HRXVPK2Q"
      } ],
      "carrier_account_slug" : "happy-returns",
      "carrier_account_id" : "7bf84d35-2cfa-4141-b5e8-b7d4762d99b6",
      "service_type" : "happy-returns_standard",
      "rate" : {
        "booking_cut_off" : null,
        "extra_info" : null,
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Happy Returns Standard",
        "service_type" : "happy-returns_standard",
        "shipper_account" : {
          "description" : "Happy returns",
          "id" : "7bf84d35-2cfa-4141-b5e8-b7d4762d99b6",
          "slug" : "happy-returns"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "internal_custom_fields" : {
          "returns_shipping_type" : "happy_returns",
          "tracking_slug" : "happy-returns"
        },
        "shipper_account" : {
          "custom_fields" : {
            "shipment_label_enabled" : true
          }
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 0.001
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 10,
        "height" : 10,
        "unit" : "cm",
        "width" : 10
      } ],
      "items" : [ {
        "description" : "exclude tax = 10% bundle Partial equal prices - 21ky0303-white / D",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0840/1243/3705/products/Hdea0a98d9c68431882adfe75df66c82dD.jpg?v=1699423437" ],
        "item_id" : "15762407096617",
        "price" : {
          "amount" : 90,
          "currency" : "USD"
        },
        "quantity" : 1,
        "return_reason" : "dhl-germany",
        "sku" : "66371698-21ky0303-white-d",
        "weight" : {
          "unit" : "kg",
          "value" : 1.0E-6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "3c725b19318248d9a5bc88d3cdea143c",
      "tracking_numbers" : [ "863T30049461A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T06:34:29.081135Z",
      "updated_at" : "2025-03-11T06:34:36.699007Z",
      "succeed_at" : "2025-03-11T06:34:35.91950921Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-11/3c725b19-3182-48d9-a5bc-88d3cdea143c-1741674875114.png"
        },
        "packing_slip" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/packing_slip/2025-03-11/ed0bee8c-6905-4dec-8efd-3b9ece1fa811-1741674876403833.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : [ {
        "enabled" : true,
        "type" : "paperless_invoice"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "b77b90458bf74e6e97d76bf5ccc94711",
      "tracking_numbers" : [ "863T30049460A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T06:14:35.983039Z",
      "updated_at" : "2025-03-11T06:14:43.638899Z",
      "succeed_at" : "2025-03-11T06:14:42.804755118Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP",
            "postal_code" : "08150"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08150",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-11/b77b9045-8bf7-4e6e-97d7-6bf5ccc94711-1741673681996.png"
        },
        "packing_slip" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/packing_slip/2025-03-11/c5ad9375-585b-4205-bbf7-c798a60c409d-1741673683273201.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : [ {
        "enabled" : true,
        "type" : "paperless_invoice"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "02b7ccc7fecd4c86b063504a4c912748",
      "tracking_numbers" : [ "9009551371456" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T06:05:57.940831Z",
      "updated_at" : "2025-03-11T06:06:01.768022Z",
      "succeed_at" : "2025-03-11T06:06:00.810249054Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "GBR",
            "postal_code" : "43228"
          },
          "ship_to" : {
            "country" : "GBR",
            "postal_code" : "98109"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Columbus",
        "company_name" : "Test",
        "contact_name" : "Customer WH",
        "country" : "GBR",
        "email" : "test@test.com",
        "phone" : "+96567606920",
        "postal_code" : "43228",
        "state" : "OH",
        "street1" : "4401 Equity Dr"
      },
      "ship_to" : {
        "city" : "Seattle",
        "contact_name" : "Jarrod Christman",
        "country" : "GBR",
        "email" : "test@test2.com",
        "phone" : "+96567606921",
        "postal_code" : "98109",
        "state" : "WA",
        "street1" : "400 9th Ave N"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "GBR",
      "ship_to_country" : "GBR",
      "order_id" : "b6f2a29437024faa82d1fe8feedc01a0",
      "order_number" : "test_order_number_100030",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/df6416fe-3328-4b10-a82b-d546f5835b57-1741673160271472.pdf"
        },
        "packing_slip" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/packing_slip/2025-03-11/9ceaf6fb-05f3-43fe-8a5a-f8ba9f7bee9d-1741673161416797.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "REFERENCE0", "REFERENCE1" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "amazon-shipping",
      "carrier_account_id" : "5998064f2bd4400b8b89340f28eef504",
      "service_type" : "amazon-shipping_ground_us_on-amazon",
      "rate" : {
        "charge_weight" : {
          "unit" : "kg",
          "value" : 2
        },
        "service_name" : "Amazon Shipping Ground (US On Amazon)",
        "service_type" : "amazon-shipping_ground_us_on-amazon",
        "shipper_account" : {
          "description" : "[amazon-shipping] Demo Account [2025-03-11]",
          "id" : "5998064f2bd4400b8b89340f28eef504",
          "slug" : "amazon-shipping"
        },
        "total_charge" : {
          "amount" : 4.76,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment_id" : "amzn1.49392bb7-709a-427e-9310-ce15de6c0424"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 2
      },
      "chargeable_weight" : {
        "unit" : "kg",
        "value" : 2
      },
      "dimension" : [ {
        "depth" : 100,
        "height" : 100,
        "unit" : "cm",
        "width" : 100
      } ],
      "items" : [ {
        "description" : "Samsung Galaxy Tab wifi 3G 64GB black",
        "hs_code" : "1234.12",
        "origin_country" : "USA",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "2e41040fa79f4b3da9b3541dfbeddd2c",
      "tracking_numbers" : [ "794684692683" ],
      "organization_id" : "3fe7752813cb46b7baa82d0ba9ac5c96",
      "created_at" : "2025-03-11T06:03:25.894127Z",
      "updated_at" : "2025-03-11T06:03:32.926888Z",
      "succeed_at" : "2025-03-11T06:03:32.914831549Z",
      "meta" : {
        "source" : "tts"
      },
      "source" : "tts",
      "ship_from" : {
        "city" : "Centerville",
        "company_name" : "Postmen",
        "contact_name" : "Contact Name",
        "country" : "USA",
        "email" : "usps_discounted@test.com",
        "phone" : "18682306030",
        "postal_code" : "45458",
        "state" : "OH",
        "street1" : "8700 Meadowcreek Dr",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Tolland",
        "company_name" : "Maersk",
        "contact_name" : "Mister Jones",
        "country" : "USA",
        "email" : "maersk.test@example.com",
        "phone" : "8777285325",
        "postal_code" : "06084",
        "state" : "CT",
        "street1" : "102 Timber Trail",
        "type" : "residential"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : "20150309_fedex_ground_economy",
      "return_shipment" : false,
      "files" : {
        "customs_declaration" : null,
        "invoice" : null,
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6"
        },
        "manifest" : null,
        "packing_slip" : null,
        "qr_code" : null
      },
      "box_type" : "custom",
      "references" : [ "order950000002" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "454b6226-60db-4c5f-9c77-eb2236abc7d9",
      "service_type" : "fedex_ground_economy",
      "rate" : {
        "booking_cut_off" : null,
        "charge_weight" : {
          "unit" : "lb",
          "value" : 3
        },
        "delivery_date" : "2025-03-14",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 16.58,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 1.1,
            "currency" : "USD"
          },
          "type" : "delivery_confirmation"
        }, {
          "charge" : {
            "amount" : 6.2,
            "currency" : "USD"
          },
          "type" : "delivery_area"
        }, {
          "charge" : {
            "amount" : 1.41,
            "currency" : "USD"
          },
          "type" : "fuel"
        }, {
          "charge" : {
            "amount" : 1.9,
            "currency" : "USD"
          },
          "type" : "other"
        } ],
        "error_message" : null,
        "info_message" : "shipTimestamp is invalid",
        "pickup_deadline" : null,
        "service_name" : "FedEx Ground® Economy",
        "service_type" : "fedex_ground_economy",
        "shipper_account" : {
          "description" : "[fedex] Demo Account",
          "id" : "454b6226-60db-4c5f-9c77-eb2236abc7d9",
          "slug" : "fedex"
        },
        "total_charge" : {
          "amount" : 27.19,
          "currency" : "USD"
        },
        "transit_time" : null
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.2
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 3
      },
      "dimension" : [ {
        "depth" : 4,
        "height" : 3,
        "unit" : "in",
        "width" : 7
      } ],
      "items" : [ {
        "description" : "F231",
        "item_id" : "1313131223",
        "price" : {
          "amount" : 50,
          "currency" : "USD"
        },
        "quantity" : 3,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "method" : {
          "account_number" : "123456",
          "country" : "USA",
          "postal_code" : "123456",
          "type" : "account"
        },
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : "JVBERi0xLjQKMSAwIG9iago8PAovVHlwZSAvQ2F0YWxvZwovUGFnZXMgMyAwIFIKPj4KZW5kb2JqCjIgMCBvYmoKPDwKL1R5cGUgL091dGxpbmVzCi9Db3VudCAwCj4+CmVuZG9iagozIDAgb2JqCjw8Ci9UeXBlIC9QYWdlcwovQ291bnQgMQovS2lkcyBbMTggMCBSXQo+PgplbmRvYmoKNCAwIG9iagpbL1BERiAvVGV4dCAvSW1hZ2VCIC9JbWFnZUMgL0ltYWdlSV0KZW5kb2JqCjUgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9IZWx2ZXRpY2EtQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjcgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvSGVsdmV0aWNhLU9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iago4IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0hlbHZldGljYS1Cb2xkT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjkgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL1N1YnR5cGUgL1R5cGUxCi9CYXNlRm9udCAvQ291cmllcgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEwIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZAovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjExIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItT2JsaXF1ZQovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjEyIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL0NvdXJpZXItQm9sZE9ibGlxdWUKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxMyAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Sb21hbgovRW5jb2RpbmcgL01hY1JvbWFuRW5jb2RpbmcKPj4KZW5kb2JqCjE0IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9TdWJ0eXBlIC9UeXBlMQovQmFzZUZvbnQgL1RpbWVzLUJvbGQKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNSAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1JdGFsaWMKL0VuY29kaW5nIC9NYWNSb21hbkVuY29kaW5nCj4+CmVuZG9iagoxNiAwIG9iago8PAovVHlwZSAvRm9udAovU3VidHlwZSAvVHlwZTEKL0Jhc2VGb250IC9UaW1lcy1Cb2xkSXRhbGljCi9FbmNvZGluZyAvTWFjUm9tYW5FbmNvZGluZwo+PgplbmRvYmoKMTcgMCBvYmogCjw8Ci9DcmVhdGlvbkRhdGUgKEQ6MjAwMykKL1Byb2R1Y2VyIChGZWRFeCBTZXJ2aWNlcykKL1RpdGxlIChGZWRFeCBTaGlwcGluZyBMYWJlbCkNL0NyZWF0b3IgKEZlZEV4IEN1c3RvbWVyIEF1dG9tYXRpb24pDS9BdXRob3IgKENMUyBWZXJzaW9uIDUxMjAxMzUpCj4+CmVuZG9iagoxOCAwIG9iago8PAovVHlwZSAvUGFnZQ0vUGFyZW50IDMgMCBSCi9SZXNvdXJjZXMgPDwgL1Byb2NTZXQgNCAwIFIgCiAvRm9udCA8PCAvRjEgNSAwIFIgCi9GMiA2IDAgUiAKL0YzIDcgMCBSIAovRjQgOCAwIFIgCi9GNSA5IDAgUiAKL0Y2IDEwIDAgUiAKL0Y3IDExIDAgUiAKL0Y4IDEyIDAgUiAKL0Y5IDEzIDAgUiAKL0YxMCAxNCAwIFIgCi9GMTEgMTUgMCBSIAovRjEyIDE2IDAgUiAKID4+Ci9YT2JqZWN0IDw8IC9GZWRFeEdyb3VuZCAyMCAwIFIKPj4KPj4KL01lZGlhQm94IFswIDAgMjg4IDQzMl0KL1RyaW1Cb3hbMCAwIDI4OCA0MzJdCi9Db250ZW50cyAxOSAwIFIKL1JvdGF0ZSAwPj4KZW5kb2JqCjE5IDAgb2JqCjw8IC9MZW5ndGggMzE4MwovRmlsdGVyIFsvQVNDSUk4NURlY29kZSAvRmxhdGVEZWNvZGVdIAo+PgpzdHJlYW0KR2F0PS8/I3M0LyVYbHImclc/ZzBkVW5yWyxZN0U3aFwnLjFoMlBwcjtZJjxSWGZfISNuXTAuIVMiOU5OJF4uZTgkVWM0ITFKdTBEa01kJFUKI19HMHJLL2MwJGV0VmJsNShgO1I6PWBQRE1pTCFGaDBTZCxRNnVlOmZta00/KyZkazpyPiZrMStDIi9KWjUqR0hEZjg2JDYxWGw9Plk9aloKOSk3MSZEaklfWnF1QSgzamFVa0tjMTY/P2lzW3FDVEEzXCUzcTI4KGJeQjwqZnVgYkteS15cJFJrTCRiKidXIjdcQG9GW1Nfb29cWDZdR2IKOVxLJkdsTGFmIXJVcCRMM0VDTnNVaDFxJSJJXDxDbEpoTy9xOyRBcWQyXyIrY18+Y2pwb04qPUcqV1ppMnR1WE1eRmxebWYmKXNcVDFtJU4KcTxLcjBbbzk4LGQ9WGVaJDFNTmlRNyNIMGs9alJYS0Y4OztOYFlQWTpFUTNBYEdYXmNJLy9XMzFJJ2RSWD9Ta2tRN3FkRVFFc3RiUWUyNmEKSGNzJU9qaml0aGpfWz50SSg+YiNRS0lUIzJ0T05AUHJVcl1tblQwXG5eR25NYkVwUF1TYW5aPVtvWlkzUnU/YVtfOiR0TEYqSlUjQkQ1IjMKKFVhUCowMjw+R0U7b1JwX3UuYF9uLitoT2wzNUxiRlYuNFMyITRgL18hR2hFZFhIYyNmOklvc18hR2s2MGQzQy9lPjtiXiJzLC4hSzI9bGMKLSJEYzk4dEBiRlNdLUFRXHAwI1VOZzsjUzpHbjBANylYWmdkITMwLlJZcVg2KyNAKTtQT0l0WU0qVSEtYjNfPjZNNDRdZjRTbCYoJz9DLzYKQDdlUzI/UEUmLUBmRzcsbXImMDxcTS9mSlpYNGNOajMsaiJPaEBcXUsnXT9tKkwqaFMicmlDSkRqVTpHSi8jYFVeYTguZV45W15gOjplPFgKYm02ST4rXixlXWJyTy0vWyZvNCZMYyQzdTVAMzRFQXJbQTkxIkdeUGB1WVRpQDElLzwhMSdAQEBQWjpzY2o3QGlBcSddJW8wW2xvbFpgI00KNDZEMTtgVENVKzoiJSVPZDxwQnQiM1FaX29lbydUa1FDVm5JI0Zfc0wqOiFKMV5EVyJUWl88a0drR09TR1ZPOU86aFdnQjkncFVNKVYtOW8KJmV0WS88NSohLVYwbCRoaVc2I1Y7Kk1OSyEzL2pVVTxOYFFhIWxILVBlWG5SJ29ybm5QX14vTyxSU3I/PiJac1RTUFJLUj4kMTBjTlBSIWIKanNZcHQhKiRORGlvSmdYayM+NGUuKFI5OSwxL0NDVzVSazJKMjVuKCdUZStmOjcyUXVSTmhBYk84cmRUTGNBLWlGVnElKSQ0W2YpaWhXQykKNzRHLGRnKnNGLCJ0WzAsPVhiJ2k1YyRcIjlcPzN1O0lwJUA6MUBabC9KNSxbMkM2SVhTIipUNjMpMyYwQCdYb00/KlBdJGcwPWIpTmAvRSgKTl9WMD9fLicqRCU3TUlxO15RXTw0X1hwOjdHPD9bTF9ybk8mYWMqXmM2cy1SQFUpL2pEaVc7M1x0IisjITEzaHFSOE1aYTMpOF43bldSYEUKNkYyTjoyQiUuJmowVjNvbyhAIT5BMTNlRzwpR1lOOTExTiZFOV9WMFw2S25GOGBRYklDQmppSGQya1hmTF4kKyU8VU8jWUk3U0goYiMxLCsKSkBgKy9OazJlXyFmc19QLkYoZ1AlIk1OWlc5X1xjR0gnPm4uVFc/bWkmRztVRjA4M2glJj81aWNvcyQ6MHFhZz5XLmYvYC9lQVtnImU8RTkKO2IrTjBYUnEndVM+aj06bzc3Ik8qZEtlMGA6YyU4UGdGKGkuTG9FVnEtS1NEVmZOTSc4VEZTZDptcEZab10zLDs1bV0wJT1qRFMlVm9aJEUKLmwqbiNlOlloJl8xPjRWNkwjOkhqOHRvMiIuPXVIWzgsLFNabyYlWV08S1VmbSknJm5ENW5nVVdFR24jSCtSMm9MaiEwQFJhVE9LODdTR1QKTXFkLmlXMjVCY0xUZlEkO2lxW1ROb0FKVVsuckM4R0khJkRgWyg0LCIvKHIwZmo6R1lgXS1KblVEczlhPDdvTj9vWGpKLGhfPUhXZTNwLT8KRSghZDpLN2Aqa18xPmRpJHBBcV4qSlAtRV9ePEYnTSxJXi0yInFiai83TTRqVFUpVk1DRjomR091M1RVNylhXmxXbUw+ci0+Rj5faERaOUgKOXBEUihrLmAkKGBbJURhIzpjXk9fLz4sP0ZaZVNUVE1xRyg5dU1UdFduP3AwTnInQkcoUVtuYlRVKmFSSWFGaVFOXShoMGpXbDtjPClbNFQKU3JiLV5BUXUxdSIvKSk0cEMnZSc2QWZlNjNyaiNwVyxsZSwpKENtaEpELEMtTT9EQy48ISk+dUJydHA4Sz1MWGNKNkZGQEhPYVxXJnBjL2cKUjAwK04uZUkzbWVYOStFTnRddW8iMTopJDUjOG0lUUtmQSchalRvZyE0Py1lYypeKTtjSUxgRz5mJUpab2RdJE8+RlxAImBxQW82aV1zVkkKQksoTyFWJVhNdUIyT1tmaEF1XjtJXFRCR0UyRl9RWkhkSURhZVM6LypcRDYmVlMxaV9MZFxqN21GNFNnI0pVbWwqSiY7PV9KcGRJN11FKksKSChZPUxZQTk1dSJcTm4+WiEvTWZuYlclUW1KTnM9UFA6OCpqJlckPFJrTlprXWViVEJmblo+VFxOLS9abnBcO1RjNDtqRz8hJEsya1ZcRTEKJ0MoXnRcPm4vSjw2XEQ6VDhJaWY1dUAuPVZAbChBIig/PXJeOmxAXHAhcnA6cnAoJWtnVSpxJiQpMmRIayJeaVZlMVdZW0lUcjtqZWBGYWsKZXU1ODkzMHVtSTdAUEIrVCN0blUnJjpCamEhWGBmTU4/WEVDWzUlRCg0NTdwXGIrLDsxIidyRkhvPHVLOzYtJHJZN0xWVG9eNTZATlZVSFAKaD8kc1wxPGEtN0UxcS8uWT5NQDhOTDl1M1NdJW49ciYsZj1XLHJadW5EJUEuclBqa04vXnNlLiVMbUMiVz0sbGQ+MWw9QFlIQDNMT0dMZj0KJyxSbFspSmxsaEpWVihLcm04Ky5TaiFNU0BqNnBIUkUiSnNeTDVYalV0LlFURSw/Uk4lPDBKI19wSGt1V1lfMzxQOEVBQV9pdTJPQUAzVC4KOD5mQkFxMG5RSzg7RypNUCFtNXI+K1hNUjkqLyVObFctIi1ANSk4R1A0Kzo1Lm9jWXM6LlBOdWUhNVhgPyk3PXBoQmFOJi9IZzhXOVkjdGIKbF43bFBZVSZMbWBdUDY8am1jVjBWSW5MbS8nU3Uia3BEQ2pCayVvMmNLVWlKZCJNW3NQSS4rMTlebWVTRlJHQyNxI1g5MVI5Z0NqJVdJZmAKKkNucCwxZC1JbDtbPnJ1M3MyIW5uTWs6UlJgYVAvb3BOY1wrW1E6JGcqXlY4Zydeb3FvZ1JVJSJ1bko9UjtzNl5FJV0rbVBKa0NZXG9fbzoKcFxRVilPZiQqbTcpWClmPjNgbkhNMUx0JCNuYHEvWmovUUpQT0VPc1ZkOi4vTWlwalVWZTBDQVJfWmtcM0AwbUhbLyxodDExb04qMiYjW3AKN2RqTzdlcS1fUyJ1bklSO3BrZHEwJFBpNWtiUWZVTC88cW48JmIiKSMjaytJKSwyNzVcdCJxP2c7Nm1gZEYsPUIyRyxUQVM9NC5eN01DI2kKMiohNXJlJ0o0PTAkNVcyUiJkTEdpZl1ndExGSG08JlhWUkM5TydoJTFDWSRZL15wa1EtNGhtcDo8VVE5MVAoQUpFQ3FRKVQ4ZV07KVEyXCoKaWdFJ01lTW5pZ2Y/VkIwIiMzUW9bJl1gKDZxaUtlbmhCMCErUT86aE5iS1crTkEja2Y4dDszKGwiMHUxKFpMWlE+VypOVyU8T1FdJzMhLVkKLTgqMk1SQCVZNzBRQGgmYWkkJ0MxIi9DQGRrV1oiOU9RaXVTPEFeWkFuLl1oIkssZz9Kc2pFbSdmYjppMCVcOEwtOCooX0lIIzsjUXRgNS4Ka21rLmRaQksuWjc/MWJdYFpkLGxTQ1QnYlEzZk07OVJLY2JgNUZFXkxNOlxMQ2deRzdecilzLGpoKj1JQ3M5ZV4yaTZeOWVKNVxuPSQzJWEKWjlvOTBzNydndEBqQThWKzNoTGAtaX4+CmVuZHN0cmVhbQplbmRvYmoKMjAgMCBvYmoKPDwgL1R5cGUgL1hPYmplY3QKL1N1YnR5cGUgL0ltYWdlCi9XaWR0aCAxMTgKL0hlaWdodCA0NwovQ29sb3JTcGFjZSAvRGV2aWNlR3JheQovQml0c1BlckNvbXBvbmVudCA4Ci9MZW5ndGggNDUwCi9GaWx0ZXIgWy9BU0NJSTg1RGVjb2RlIC9GbGF0ZURlY29kZV0KPj5zdHJlYW0KR2IiL2VKSV1SPyNYbkxnVDY8ZGYvSS4mP0I9aGNEYWctckQ0XTAlUGZNVForYFFTSStKXUNeJnJEM0E4bE04bm5GRlAhSS5mKmpZQzJUblsKPDdpKDdcPEpIITx1VnVtM0hKZms8OCkrJVpOKk1mLXMlbD5WWlVDcyw5YDlvVyZUTSUvY2dtSlUlLmZMaGJJJCc+Qmd1X2lAUStDJmJXcjYKM0RDTTYvRlVNdCxEQl4sKlIwMHMiXVdGR0MwajojV2E0Z1c8KlFlVDssPztCUyFvcnU1dSRfM0IyalwyQy4nYi5GWHFAWDFBKixhXFNhR0oKOm5Na2NJUVwuSmY6MDU3Xy1MNEo8JTJRTW1JYW9AXDg4RktHYjFwLG1XOyUuSE5DLjtWLitgZClzOnJCaFVqTEhCWDJSQ145X3E4MW5WW0sKTD0yUUJGdDtyVENuNltmcG1wcjdTX2NHIWRXWyxzWDErKzhSPCktJUZjVS5rUnBXUylYRTs5Kzs0bFMrXzZBc3JjRWJRTGNJRWNfOyVNS0UKM2NqR1Q5Kk5bSi5wTjJXbUk4Vz8kbDtjNl01ZC8/cDh0PHNFYUoyZXJsK34+CmVuZHN0cmVhbQplbmRvYmoKeHJlZgowIDIxCjAwMDAwMDAwMDAgNjU1MzUgZiAKMDAwMDAwMDAwOSAwMDAwMCBuIAowMDAwMDAwMDU4IDAwMDAwIG4gCjAwMDAwMDAxMDQgMDAwMDAgbiAKMDAwMDAwMDE2MiAwMDAwMCBuIAowMDAwMDAwMjE0IDAwMDAwIG4gCjAwMDAwMDAzMTIgMDAwMDAgbiAKMDAwMDAwMDQxNSAwMDAwMCBuIAowMDAwMDAwNTIxIDAwMDAwIG4gCjAwMDAwMDA2MzEgMDAwMDAgbiAKMDAwMDAwMDcyNyAwMDAwMCBuIAowMDAwMDAwODI5IDAwMDAwIG4gCjAwMDAwMDA5MzQgMDAwMDAgbiAKMDAwMDAwMTA0MyAwMDAwMCBuIAowMDAwMDAxMTQ0IDAwMDAwIG4gCjAwMDAwMDEyNDQgMDAwMDAgbiAKMDAwMDAwMTM0NiAwMDAwMCBuIAowMDAwMDAxNDUyIDAwMDAwIG4gCjAwMDAwMDE2MjIgMDAwMDAgbiAKMDAwMDAwMTk2NiAwMDAwMCBuIAowMDAwMDA1MjQxIDAwMDAwIG4gCnRyYWlsZXIKPDwKL0luZm8gMTcgMCBSCi9TaXplIDIxCi9Sb290IDEgMCBSCj4+CnN0YXJ0eHJlZgo1ODc3CiUlRU9GCg=="
    }, {
      "id" : "f80f831dba23482baa90e7fa554cc5ea",
      "tracking_numbers" : [ "3446116796" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T05:40:50.856031Z",
      "updated_at" : "2025-03-11T05:40:57.37361Z",
      "succeed_at" : "2025-03-11T05:40:54.610831708Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "GBR",
            "postal_code" : "M24 2RW"
          },
          "ship_to" : {
            "country" : "GBR",
            "postal_code" : "HA1 4TR"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Manchester",
        "company_name" : "CLARKE TELECOM LTD - Stores",
        "contact_name" : "kieron.slack",
        "country" : "GBR",
        "email" : "kieron.slack@clarke-telecom.com",
        "phone" : "+447391016558",
        "postal_code" : "M24 2RW",
        "street1" : "Unit 12-15,Stakehill",
        "street3" : "MIddleton ",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Harrow",
        "company_name" : "Electro Rent UK Ltd",
        "contact_name" : "GBR Har Dispatch Manager",
        "country" : "GBR",
        "email" : "despatch@electrorent.com",
        "eori_number" : "GB541333672000",
        "phone" : "+442084200200",
        "postal_code" : "HA1 4TR",
        "state" : "GB-HRW",
        "street1" : " Unit 1, Waverley",
        "tax_id" : "GB 541 333 672",
        "type" : "business"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "GBR",
      "ship_to_country" : "GBR",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x8",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/f80f831d-ba23-482b-aa90-e7fa554cc5ea-1741671653748.pdf"
        },
        "packing_slip" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/packing_slip/2025-03-11/af1cf1b4-6b57-4e84-b045-1e4a6c223766-1741671657010594.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "SON040000609" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "dhl",
      "carrier_account_id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
      "service_type" : "dhl_domestic_express",
      "rate" : {
        "delivery_date" : "2025-03-12T23:59:00Z",
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "DHL Express Domestic",
        "service_type" : "dhl_domestic_express",
        "shipper_account" : {
          "description" : "[mydhl] Demo Account [Automate Test]",
          "id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
          "slug" : "dhl"
        },
        "transit_time" : 2
      },
      "service_options" : [ {
        "enabled" : false,
        "type" : "paperless_invoice"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "description" : "A Default Box",
            "dimension" : {
              "depth" : 2.1,
              "height" : 2.1,
              "unit" : "cm",
              "width" : 2.1
            },
            "items" : [ {
              "barcode" : "",
              "description" : "NOT RESTRICTEDNarda Nardalert S3 Personal Monitor Mainframe & 2271/31 Sensor Module",
              "description_cn" : "",
              "dimension" : {
                "depth" : 0,
                "height" : 0,
                "unit" : "",
                "width" : 0
              },
              "hs_code" : "9030310000",
              "image_urls" : null,
              "item_id" : "",
              "origin_country" : "CRI",
              "price" : {
                "amount" : 37.49,
                "currency" : "GBP"
              },
              "quantity" : 1,
              "return_reason" : "",
              "sku" : "",
              "weight" : {
                "unit" : "kg",
                "value" : 1
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 1
            }
          } ]
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 2.1,
        "height" : 2.1,
        "unit" : "cm",
        "width" : 2.1
      } ],
      "items" : [ {
        "description" : "NOT RESTRICTEDNarda Nardalert S3 Personal Monitor Mainframe & 2271/31 Sensor Module",
        "hs_code" : "9030310000",
        "origin_country" : "CRI",
        "price" : {
          "amount" : 37.49,
          "currency" : "GBP"
        },
        "quantity" : 1,
        "weight" : {
          "unit" : "kg",
          "value" : 1
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "0737b037c24842d8a85897ed5f0eeb3e",
      "tracking_numbers" : [ "863T30049458A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T03:56:23.979599Z",
      "updated_at" : "2025-03-11T03:56:26.089832Z",
      "succeed_at" : "2025-03-11T03:56:26.079161031Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP",
            "postal_code" : "08150"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08150",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-11/0737b037-c248-42d8-a858-97ed5f0eeb3e-1741665385274.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : [ {
        "enabled" : true,
        "type" : "paperless_invoice"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "b3cfc9fd32794e4b84fb75cf6b7b6b9b",
      "tracking_numbers" : [ "9018530126264" ],
      "organization_id" : "e430e307533f44f58cf8ef3d3b0001b7",
      "created_at" : "2025-03-11T03:52:04.389375Z",
      "updated_at" : "2025-03-11T03:52:06.788384Z",
      "succeed_at" : "2025-03-11T03:52:06.778317974Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "admin"
      },
      "source" : "admin",
      "ship_from" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "id" : "e4cd7fcc9eb54bc180b75be7d318708a",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "a98e99da959d4350b5aed0fa13dcda32",
      "order_number" : "1234475",
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-11/689069be-7f93-4ba0-bb84-1a0cd7f9a75d-1741665126460781.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/2dec98ef-9798-4345-a4d0-431acf3df2a0-1741665125832935.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "1234475" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "amazon-shipping",
      "carrier_account_id" : "db5fc5c4bbac441cbb73ab71cfef41a3",
      "service_type" : "amazon-shipping_ground_us_on-amazon",
      "rate" : {
        "booking_cut_off" : null,
        "charge_weight" : {
          "unit" : "kg",
          "value" : 0
        },
        "extra_info" : null,
        "service_name" : "Amazon Shipping Ground (US On Amazon)",
        "service_type" : "amazon-shipping_ground_us_on-amazon",
        "shipper_account" : {
          "description" : "Amazon test US",
          "id" : "db5fc5c4bbac441cbb73ab71cfef41a3",
          "slug" : "amazon-shipping"
        },
        "total_charge" : {
          "amount" : 4.76,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "internal_custom_fields" : {
          "account_id" : "a5d1dbd350bf4653a5ccace388f2cb48",
          "ship_from" : {
            "id" : "e4cd7fcc9eb54bc180b75be7d318708a"
          }
        },
        "shipment_id" : "amzn1.846b8172-e89e-44d0-9999-3784f4fb2fe1"
      },
      "custom_fields" : {
        "location_name" : "Test US location"
      },
      "weight" : {
        "unit" : "lb",
        "value" : 0.001
      },
      "chargeable_weight" : {
        "unit" : "kg",
        "value" : 0
      },
      "dimension" : [ {
        "depth" : 10,
        "height" : 5,
        "unit" : "cm",
        "width" : 10
      } ],
      "items" : [ {
        "description" : "Auto-generated by Postmen",
        "price" : {
          "amount" : 1,
          "currency" : "USD"
        },
        "quantity" : 1,
        "weight" : {
          "unit" : "kg",
          "value" : 0
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "a5d1dbd350bf4653a5ccace388f2cb48",
      "label_base64" : null
    }, {
      "id" : "57018c5ac38e46b9b6ee345d2834077b",
      "tracking_numbers" : [ "863T30049457A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T03:47:33.018655Z",
      "updated_at" : "2025-03-11T03:47:35.122073Z",
      "succeed_at" : "2025-03-11T03:47:35.114128866Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP",
            "postal_code" : "08150"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08150",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-11/57018c5a-c38e-46b9-b6ee-345d2834077b-1741664854311.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "a9901d0e510b4c929f0a002e258c5f74",
      "tracking_numbers" : [ "863T30049456A002" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T03:42:15.371039Z",
      "updated_at" : "2025-03-11T03:42:22.171429Z",
      "succeed_at" : "2025-03-11T03:42:22.162729205Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP",
            "postal_code" : "08150"
          },
          "ship_to" : {
            "country" : "ESP",
            "postal_code" : "08830"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08150",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_to" : {
        "city" : "Sant Boi De Llobregat",
        "contact_name" : "Ismael Tejón",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08830",
        "street1" : "Calle de prueba, 34",
        "type" : "residential"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "ESP",
      "ship_to_country" : "ESP",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "png",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-11/a9901d0e-510b-4c92-9f0a-002e258c5f74-1741664541315.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "xxx-123456", "xxx-654321" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "yodel",
      "carrier_account_id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
      "service_type" : "yodel_returns",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "shipper_account" : {
          "description" : "[yodel] Demo Account",
          "id" : "f4e7e34b-735c-4b84-b163-9197bdac3f54",
          "slug" : "yodel"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "DEU",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "1f2b5dce8fe040e790e5a22ea147b7db",
      "tracking_numbers" : [ "13031743" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T03:11:31.153263Z",
      "updated_at" : "2025-03-11T03:11:33.067254Z",
      "succeed_at" : "2025-03-11T03:11:33.057218688Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP",
            "postal_code" : "08150"
          },
          "ship_to" : {
            "country" : "FRA",
            "postal_code" : "75001"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles(barcelona)",
        "company_name" : "Panda of china",
        "contact_name" : "Javid",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08150",
        "state" : "Mixed",
        "street1" : "ho no 115/flat no b1 street no 7.",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Paris",
        "contact_name" : "Ismael Tejón",
        "country" : "FRA",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "75001",
        "state" : "VIL",
        "street1" : "Calle de prueba, 34.",
        "type" : "residential"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "ESP",
      "ship_to_country" : "FRA",
      "order_id" : null,
      "order_number" : "100001",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/1f2b5dce-8fe0-40e7-90e5-a22ea147b7db-1741662692296.pdf"
        }
      },
      "box_type" : "nacex_sample",
      "references" : [ "SCF – IKEA GRAN CANARIA | SCF." ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "nacex",
      "carrier_account_id" : "c5573249c14b4f9eb077e36985b2f6e9",
      "service_type" : "nacex_servicio-aereo",
      "rate" : {
        "delivery_date" : "2025-03-13",
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "NACEX SERVICIO AEREO",
        "service_type" : "nacex_servicio-aereo",
        "shipper_account" : {
          "description" : "[nacex] Demo account new",
          "id" : "c5573249c14b4f9eb077e36985b2f6e9",
          "slug" : "nacex"
        },
        "transit_time" : 2
      },
      "service_options" : [ {
        "end_time" : "13:00:00",
        "start_time" : "12:00:00",
        "type" : "pickup"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment_code" : "429862185"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "ESP",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "e7e7a49328564a0da620aa3a1340c8a1",
      "tracking_numbers" : [ "9013420180341" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T02:54:02.743215Z",
      "updated_at" : "2025-03-11T02:54:05.510869Z",
      "succeed_at" : "2025-03-11T02:54:05.503657861Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "GBR",
            "postal_code" : "43228"
          },
          "ship_to" : {
            "country" : "GBR",
            "postal_code" : "98109"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Columbus",
        "company_name" : "Test",
        "contact_name" : "Customer WH",
        "country" : "GBR",
        "email" : "test@test.com",
        "phone" : "+96567606920",
        "postal_code" : "43228",
        "state" : "OH",
        "street1" : "4401 Equity Dr"
      },
      "ship_to" : {
        "city" : "Seattle",
        "contact_name" : "Jarrod Christman",
        "country" : "GBR",
        "email" : "test@test2.com",
        "phone" : "+96567606921",
        "postal_code" : "98109",
        "state" : "WA",
        "street1" : "400 9th Ave N"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "GBR",
      "ship_to_country" : "GBR",
      "order_id" : "b6f2a29437024faa82d1fe8feedc01a0",
      "order_number" : "test_order_number_100030",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-11/5dbe64c0-1a60-4a7a-85ca-e3692f2c336b-1741661645145365.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "REFERENCE0", "REFERENCE1" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "amazon-shipping",
      "carrier_account_id" : "5998064f2bd4400b8b89340f28eef504",
      "service_type" : "amazon-shipping_ground_us_on-amazon",
      "rate" : {
        "charge_weight" : {
          "unit" : "kg",
          "value" : 2
        },
        "service_name" : "Amazon Shipping Ground (US On Amazon)",
        "service_type" : "amazon-shipping_ground_us_on-amazon",
        "shipper_account" : {
          "description" : "[amazon-shipping] Demo Account [2025-03-11]",
          "id" : "5998064f2bd4400b8b89340f28eef504",
          "slug" : "amazon-shipping"
        },
        "total_charge" : {
          "amount" : 4.76,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment_id" : "amzn1.851dcbc0-8b7c-454d-877b-d42f13b91d11"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 2
      },
      "chargeable_weight" : {
        "unit" : "kg",
        "value" : 2
      },
      "dimension" : [ {
        "depth" : 100,
        "height" : 100,
        "unit" : "cm",
        "width" : 100
      } ],
      "items" : [ {
        "description" : "Samsung Galaxy Tab wifi 3G 64GB black",
        "hs_code" : "1234.12",
        "origin_country" : "USA",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "91e683ae7d8e4e7cabe2fe283e893175",
      "tracking_numbers" : [ "450068309482" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-11T02:31:51.328015Z",
      "updated_at" : "2025-03-11T02:31:52.552579Z",
      "succeed_at" : "2025-03-11T02:31:52.545514225Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "JPN",
            "postal_code" : "1860003"
          },
          "ship_to" : {
            "country" : "JPN",
            "postal_code" : "1350061"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "国立市",
        "contact_name" : "田中(タナカ) 太郎(タロウ)",
        "country" : "JPN",
        "phone" : "01234567890",
        "postal_code" : "1860003",
        "state" : "東京都",
        "street1" : "富士見台２-４３-４"
      },
      "ship_to" : {
        "city" : "江東区",
        "company_name" : "ラルフローレン合同会社 DC name",
        "contact_name" : "person name",
        "country" : "JPN",
        "phone" : "01234567890",
        "postal_code" : "1350061",
        "state" : "東京都",
        "street1" : "豊洲５－４－９"
      },
      "ship_date" : "2025-03-11",
      "ship_from_country" : "JPN",
      "ship_to_country" : "JPN",
      "order_id" : null,
      "order_number" : "20250311a01",
      "return_shipment" : true,
      "files" : null,
      "box_type" : "custom",
      "references" : [ "Order #20250311a01" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "taqbin-jp",
      "carrier_account_id" : "aef6e8b4-8c1e-4ac8-bf1d-55b0fc465754",
      "service_type" : "taqbin-jp_standard",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "TA-Q-BIN",
        "service_type" : "taqbin-jp_standard",
        "shipper_account" : {
          "description" : "[taqbin-jp] Return RL Account",
          "id" : "aef6e8b4-8c1e-4ac8-bf1d-55b0fc465754",
          "slug" : "taqbin-jp"
        }
      },
      "service_options" : [ {
        "end_time" : "17:00:00",
        "start_time" : "00:00:00",
        "type" : "pickup"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : { },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 1,
        "height" : 1,
        "unit" : "cm",
        "width" : 1
      } ],
      "items" : [ {
        "description" : "衣類その他",
        "origin_country" : "USA",
        "price" : {
          "amount" : 72000,
          "currency" : "JPY"
        },
        "quantity" : 1,
        "weight" : {
          "unit" : "kg",
          "value" : 1
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "2febe5f631b34e30b23d950bcbfdcecc",
      "tracking_numbers" : [ "6690756192" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T11:20:59.348943Z",
      "updated_at" : "2025-03-10T11:21:02.841491Z",
      "succeed_at" : "2025-03-10T11:21:02.831527251Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "GBR",
            "postal_code" : "M24 2RW"
          },
          "ship_to" : {
            "country" : "GBR",
            "postal_code" : "HA1 4TR"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Manchester",
        "company_name" : "CLARKE TELECOM LTD - Stores",
        "contact_name" : "kieron.slack",
        "country" : "GBR",
        "email" : "kieron.slack@clarke-telecom.com",
        "phone" : "+447391016558",
        "postal_code" : "M24 2RW",
        "street1" : "Unit 12-15,Stakehill Industrial Estate,",
        "street3" : "MIddleton ",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Harrow",
        "company_name" : "Electro Rent UK Ltd",
        "contact_name" : "GBR Har Dispatch Manager",
        "country" : "GBR",
        "email" : "despatch@electrorent.com",
        "eori_number" : "GB541333672000",
        "phone" : "+442084200200",
        "postal_code" : "HA1 4TR",
        "state" : "GB-HRW",
        "street1" : " Unit 1, Waverley Industrial Park\nHailsham Drive",
        "tax_id" : "GB 541 333 672",
        "type" : "business"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "GBR",
      "ship_to_country" : "GBR",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x8",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/2febe5f6-31b3-4e30-b23d-950bcbfdcecc-1741605662688.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "SON040000609" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "dhl",
      "carrier_account_id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
      "service_type" : "dhl_domestic_express",
      "rate" : {
        "booking_cut_off" : "2025-03-11T22:30:00+08:00",
        "charge_weight" : {
          "unit" : "kg",
          "value" : 1
        },
        "delivery_date" : "2025-03-13T07:59:00+08:00",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 125.1,
            "currency" : "EUR"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 37.22,
            "currency" : "EUR"
          },
          "type" : "fuel_surcharge"
        } ],
        "pickup_deadline" : "2025-03-12T00:30:00+08:00",
        "service_name" : "DHL Express Domestic",
        "service_type" : "dhl_domestic_express",
        "shipper_account" : {
          "description" : "[mydhl] Demo Account [Automate Test]",
          "id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
          "slug" : "dhl"
        },
        "total_charge" : {
          "amount" : 162.32,
          "currency" : "EUR"
        },
        "transit_time" : 1
      },
      "service_options" : [ {
        "enabled" : false,
        "type" : "paperless_invoice"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "description" : "A Default Box",
            "dimension" : {
              "depth" : 2.1,
              "height" : 2.1,
              "unit" : "cm",
              "width" : 2.1
            },
            "items" : [ {
              "barcode" : "",
              "description" : "NOT RESTRICTEDNarda Nardalert S3 Personal Monitor Mainframe & 2271/31 Sensor Module",
              "description_cn" : "",
              "dimension" : {
                "depth" : 0,
                "height" : 0,
                "unit" : "",
                "width" : 0
              },
              "hs_code" : "9030310000",
              "image_urls" : null,
              "item_id" : "",
              "origin_country" : "CRI",
              "price" : {
                "amount" : 37.49,
                "currency" : "GBP"
              },
              "quantity" : 1,
              "return_reason" : "",
              "sku" : "",
              "weight" : {
                "unit" : "kg",
                "value" : 1
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 1
            }
          } ]
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1
      },
      "chargeable_weight" : {
        "unit" : "kg",
        "value" : 1
      },
      "dimension" : [ {
        "depth" : 2.1,
        "height" : 2.1,
        "unit" : "cm",
        "width" : 2.1
      } ],
      "items" : [ {
        "description" : "NOT RESTRICTEDNarda Nardalert S3 Personal Monitor Mainframe & 2271/31 Sensor Module",
        "hs_code" : "9030310000",
        "origin_country" : "CRI",
        "price" : {
          "amount" : 37.49,
          "currency" : "GBP"
        },
        "quantity" : 1,
        "weight" : {
          "unit" : "kg",
          "value" : 1
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "da772a01917b4775957367569a3b99d7",
      "tracking_numbers" : [ "6690755945" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T11:20:10.578687Z",
      "updated_at" : "2025-03-10T11:20:14.970283Z",
      "succeed_at" : "2025-03-10T11:20:14.960380784Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "GBR",
            "postal_code" : "M24 2RW"
          },
          "ship_to" : {
            "country" : "GBR",
            "postal_code" : "HA1 4TR"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Manchester",
        "company_name" : "CLARKE TELECOM LTD - Stores",
        "contact_name" : "kieron.slack",
        "country" : "GBR",
        "email" : "kieron.slack@clarke-telecom.com",
        "phone" : "+447391016558",
        "postal_code" : "M24 2RW",
        "street1" : "Unit 12-15,Stakehill Industrial Estate,",
        "street3" : "MIddleton ",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Harrow",
        "company_name" : "Electro Rent UK Ltd",
        "contact_name" : "GBR Har Dispatch Manager",
        "country" : "GBR",
        "email" : "despatch@electrorent.com",
        "eori_number" : "GB541333672000",
        "phone" : "+442084200200",
        "postal_code" : "HA1 4TR",
        "state" : "GB-HRW",
        "street1" : " Unit 1, Waverley Industrial Park\nHailsham Drive",
        "tax_id" : "GB 541 333 672",
        "type" : "business"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "GBR",
      "ship_to_country" : "GBR",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x8",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/da772a01-917b-4775-9573-67569a3b99d7-1741605614770.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "SON040000609" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "dhl",
      "carrier_account_id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
      "service_type" : "dhl_domestic_express",
      "rate" : {
        "booking_cut_off" : "2025-03-11T22:30:00+08:00",
        "charge_weight" : {
          "unit" : "kg",
          "value" : 1
        },
        "delivery_date" : "2025-03-13T07:59:00+08:00",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 125.1,
            "currency" : "EUR"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 37.22,
            "currency" : "EUR"
          },
          "type" : "fuel_surcharge"
        } ],
        "pickup_deadline" : "2025-03-12T00:30:00+08:00",
        "service_name" : "DHL Express Domestic",
        "service_type" : "dhl_domestic_express",
        "shipper_account" : {
          "description" : "[mydhl] Demo Account [Automate Test]",
          "id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
          "slug" : "dhl"
        },
        "total_charge" : {
          "amount" : 162.32,
          "currency" : "EUR"
        },
        "transit_time" : 1
      },
      "service_options" : [ {
        "enabled" : false,
        "type" : "paperless_invoice"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "description" : "A Default Box",
            "dimension" : {
              "depth" : 0.1,
              "height" : 0.1,
              "unit" : "cm",
              "width" : 0.1
            },
            "items" : [ {
              "barcode" : "",
              "description" : "NOT RESTRICTEDNarda Nardalert S3 Personal Monitor Mainframe & 2271/31 Sensor Module",
              "description_cn" : "",
              "dimension" : {
                "depth" : 0,
                "height" : 0,
                "unit" : "",
                "width" : 0
              },
              "hs_code" : "9030310000",
              "image_urls" : null,
              "item_id" : "",
              "origin_country" : "CRI",
              "price" : {
                "amount" : 37.49,
                "currency" : "GBP"
              },
              "quantity" : 1,
              "return_reason" : "",
              "sku" : "",
              "weight" : {
                "unit" : "kg",
                "value" : 1
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 1
            }
          } ]
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1
      },
      "chargeable_weight" : {
        "unit" : "kg",
        "value" : 1
      },
      "dimension" : [ {
        "depth" : 0.1,
        "height" : 0.1,
        "unit" : "cm",
        "width" : 0.1
      } ],
      "items" : [ {
        "description" : "NOT RESTRICTEDNarda Nardalert S3 Personal Monitor Mainframe & 2271/31 Sensor Module",
        "hs_code" : "9030310000",
        "origin_country" : "CRI",
        "price" : {
          "amount" : 37.49,
          "currency" : "GBP"
        },
        "quantity" : 1,
        "weight" : {
          "unit" : "kg",
          "value" : 1
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "ba9bacab51b7426daf3ccf0727895c47",
      "tracking_numbers" : [ "795495345265" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T10:17:32.091951Z",
      "updated_at" : "2025-03-10T10:17:41.970536Z",
      "succeed_at" : "2025-03-10T10:17:41.96209665Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "admin"
      },
      "source" : "admin",
      "ship_from" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "a2c65660cdd440f7a821b82b1c9c5fc2",
      "order_number" : "20250310b004",
      "return_shipment" : true,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-10/3344554a-eadb-4474-bade-b7f89a5f25fe-1741601861086865.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/3cd7e89c-a8cc-450b-9185-cbd6e235470e-1741601860580563.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "20250310b004" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
      "service_type" : "fedex_2_day",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "[FedEx] Postmen Testing Account",
          "id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
          "slug" : "fedex"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null,
        "form_id" : 221,
        "internal_custom_fields" : {
          "account_id" : "db910fb96c37438aa6015ff3efda6064"
        },
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 2,
              "height" : 2,
              "unit" : "in",
              "width" : 2
            },
            "items" : [ {
              "description" : "20250310a002",
              "hs_code" : "3132",
              "origin_country" : "USA",
              "price" : {
                "amount" : 2,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "323",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            }, {
              "description" : "20250310b002",
              "hs_code" : "2212",
              "origin_country" : "USA",
              "price" : {
                "amount" : 1,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "221",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 8
            }
          } ]
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 8
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 2,
        "height" : 2,
        "unit" : "in",
        "width" : 2
      } ],
      "items" : [ {
        "description" : "20250310a002",
        "hs_code" : "3132",
        "origin_country" : "USA",
        "price" : {
          "amount" : 2,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "323",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      }, {
        "description" : "20250310b002",
        "hs_code" : "2212",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "221",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "db910fb96c37438aa6015ff3efda6064",
      "label_base64" : null
    }, {
      "id" : "6e839717f0794fc7b01c6d251da060b5",
      "tracking_numbers" : [ "794684521714" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T10:17:03.243151Z",
      "updated_at" : "2025-03-10T10:17:11.131843Z",
      "succeed_at" : "2025-03-10T10:17:11.123353498Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : "20250310b004",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/fd8ebbed-620a-4d4d-878d-0d46e87eb2f5-1741601830125914.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "20250310b004" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
      "service_type" : "fedex_2_day",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "[FedEx] Postmen Testing Account",
          "id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
          "slug" : "fedex"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null,
        "form_id" : 201,
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 2,
              "height" : 2,
              "unit" : "in",
              "width" : 2
            },
            "items" : [ {
              "description" : "20250310a002",
              "hs_code" : "3132",
              "origin_country" : "USA",
              "price" : {
                "amount" : 2,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "323",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            }, {
              "description" : "20250310b002",
              "hs_code" : "2212",
              "origin_country" : "USA",
              "price" : {
                "amount" : 1,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "221",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 8
            }
          } ]
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 8
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 2,
        "height" : 2,
        "unit" : "in",
        "width" : 2
      } ],
      "items" : [ {
        "description" : "20250310a002",
        "hs_code" : "3132",
        "origin_country" : "USA",
        "price" : {
          "amount" : 2,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "323",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      }, {
        "description" : "20250310b002",
        "hs_code" : "2212",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "221",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "5655c069cc7346f2acbc06908d24c63e",
      "tracking_numbers" : [ "794684521931" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T10:16:53.824527Z",
      "updated_at" : "2025-03-10T10:17:02.468813Z",
      "succeed_at" : "2025-03-10T10:17:02.458352836Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : "20250310b003",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/028a3727-c6da-4a23-bc32-dde0c4f3bff1-1741601821269249.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "20250310b003" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
      "service_type" : "fedex_2_day",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "[FedEx] Postmen Testing Account",
          "id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
          "slug" : "fedex"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null,
        "form_id" : 201,
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 2,
              "height" : 2,
              "unit" : "in",
              "width" : 2
            },
            "items" : [ {
              "description" : "20250310a002",
              "hs_code" : "3132",
              "origin_country" : "USA",
              "price" : {
                "amount" : 2,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "323",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            }, {
              "description" : "20250310b002",
              "hs_code" : "2212",
              "origin_country" : "USA",
              "price" : {
                "amount" : 1,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "221",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 8
            }
          } ]
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 8
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 2,
        "height" : 2,
        "unit" : "in",
        "width" : 2
      } ],
      "items" : [ {
        "description" : "20250310a002",
        "hs_code" : "3132",
        "origin_country" : "USA",
        "price" : {
          "amount" : 2,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "323",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      }, {
        "description" : "20250310b002",
        "hs_code" : "2212",
        "origin_country" : "USA",
        "price" : {
          "amount" : 1,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "221",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "17c0b22f6a684d3d815e9a8c48cd13de",
      "tracking_numbers" : [ "794684521585" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T10:12:42.177679Z",
      "updated_at" : "2025-03-10T10:12:50.354765Z",
      "succeed_at" : "2025-03-10T10:12:50.318434179Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : "20250310b003",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/0ede835b-1fcf-4394-8ac2-764cd3cf1cad-1741601569438088.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "20250310b003" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
      "service_type" : "fedex_2_day",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "[FedEx] Postmen Testing Account",
          "id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
          "slug" : "fedex"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null,
        "form_id" : 201,
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 2,
              "height" : 2,
              "unit" : "in",
              "width" : 2
            },
            "items" : [ {
              "description" : "20250310a002",
              "hs_code" : "3132",
              "origin_country" : null,
              "price" : {
                "amount" : 2,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "323",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            }, {
              "description" : "20250310b002",
              "hs_code" : "2212",
              "origin_country" : null,
              "price" : {
                "amount" : 1,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "221",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 8
            }
          } ]
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 8
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 2,
        "height" : 2,
        "unit" : "in",
        "width" : 2
      } ],
      "items" : [ {
        "description" : "20250310a002",
        "hs_code" : "3132",
        "price" : {
          "amount" : 2,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "323",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      }, {
        "description" : "20250310b002",
        "hs_code" : "2212",
        "price" : {
          "amount" : 1,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "221",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "e67f1b0ff5c34b4c88cae2da5cb07bde",
      "tracking_numbers" : [ "794684521405" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T10:10:41.790031Z",
      "updated_at" : "2025-03-10T10:10:50.184403Z",
      "succeed_at" : "2025-03-10T10:10:50.175683637Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "admin"
      },
      "source" : "admin",
      "ship_from" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "71dc0c26c05e44ffaf189a4d9377c559",
      "order_number" : "20250310b001",
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-10/538575bf-32ac-4b1f-8b3d-216193c55cd2-1741601449241157.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/7283b555-8261-4f0f-b318-ec10573d6e19-1741601448770334.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "20250310b001" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
      "service_type" : "fedex_2_day",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "[FedEx] Postmen Testing Account",
          "id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
          "slug" : "fedex"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null,
        "form_id" : 201,
        "internal_custom_fields" : {
          "account_id" : "db910fb96c37438aa6015ff3efda6064"
        },
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 2,
              "height" : 2,
              "unit" : "in",
              "width" : 2
            },
            "items" : [ {
              "description" : "20250310a002",
              "hs_code" : "3132",
              "origin_country" : null,
              "price" : {
                "amount" : 2,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "323",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            }, {
              "description" : "20250310b002",
              "hs_code" : "2212",
              "origin_country" : null,
              "price" : {
                "amount" : 1,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "221",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 8
            }
          } ]
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 8
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 2,
        "height" : 2,
        "unit" : "in",
        "width" : 2
      } ],
      "items" : [ {
        "description" : "20250310a002",
        "hs_code" : "3132",
        "price" : {
          "amount" : 2,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "323",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      }, {
        "description" : "20250310b002",
        "hs_code" : "2212",
        "price" : {
          "amount" : 1,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "221",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "db910fb96c37438aa6015ff3efda6064",
      "label_base64" : null
    }, {
      "id" : "b76eda18187d4a0e8254e7a42cdaaa17",
      "tracking_numbers" : [ "795495344990" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T09:57:50.821103Z",
      "updated_at" : "2025-03-10T09:58:02.151009Z",
      "succeed_at" : "2025-03-10T09:58:02.142778548Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "admin"
      },
      "source" : "admin",
      "ship_from" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "1bc837b2be844bf8a7897b4e510a59f8",
      "order_number" : "20250310a002",
      "return_shipment" : true,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-10/668e2ada-9450-472c-a4a9-99e81fe66590-1741600681070948.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/d3919c75-8412-4fb1-893b-b55521b1d82f-1741600680674347.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "20250310a002" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
      "service_type" : "fedex_2_day",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "[FedEx] Postmen Testing Account",
          "id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
          "slug" : "fedex"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null,
        "form_id" : 221,
        "internal_custom_fields" : {
          "account_id" : "db910fb96c37438aa6015ff3efda6064"
        },
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 2,
              "height" : 2,
              "unit" : "in",
              "width" : 2
            },
            "items" : [ {
              "description" : "20250310a002",
              "hs_code" : "3132",
              "origin_country" : null,
              "price" : {
                "amount" : 2,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "323",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            }, {
              "description" : "20250310b002",
              "hs_code" : "2212",
              "origin_country" : null,
              "price" : {
                "amount" : 1,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "221",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 8
            }
          } ]
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 8
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 2,
        "height" : 2,
        "unit" : "in",
        "width" : 2
      } ],
      "items" : [ {
        "description" : "20250310a002",
        "hs_code" : "3132",
        "price" : {
          "amount" : 2,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "323",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      }, {
        "description" : "20250310b002",
        "hs_code" : "2212",
        "price" : {
          "amount" : 1,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "221",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "db910fb96c37438aa6015ff3efda6064",
      "label_base64" : null
    }, {
      "id" : "229690ece99e40dea201c35478a41885",
      "tracking_numbers" : [ "9405500106032109512377" ],
      "organization_id" : "b001d64c4d6343afb6e5bb69580d7661",
      "created_at" : "2025-03-10T09:56:10.646271Z",
      "updated_at" : "2025-03-10T09:56:16.09919Z",
      "succeed_at" : "2025-03-10T09:56:14.429136381Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "78752"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "warranty",
        "tags" : "automatic"
      },
      "source" : "warranty",
      "ship_from" : {
        "city" : "Austin",
        "contact_name" : "USA nyy nie",
        "country" : "USA",
        "email" : "yy.nie+8345834511@aftership.com",
        "phone" : "0123456789",
        "postal_code" : "78752",
        "state" : "TX",
        "street1" : "6929 Airport Boulevard",
        "street2" : "110"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "phone" : "0123456789",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/229690ec-e99e-40de-a201-c35478a41885-1741600574248.pdf"
        },
        "packing_slip" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/packing_slip/2025-03-10/a300c204-6a03-4603-b358-84d013c7e41a-1741600575827405.pdf"
        },
        "qr_code" : {
          "file_type" : "png",
          "paper_size" : "1x1",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-10/229690ec-e99e-40de-a201-c35478a41885-1741600574249.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA#7CGFPYOE" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "vesyl",
      "carrier_account_id" : "f05425c9c542400daa98e02e38bfcced",
      "service_type" : "vesyl_priority_mail",
      "rate" : {
        "booking_cut_off" : null,
        "charge_weight" : {
          "unit" : "lb",
          "value" : 0.5
        },
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 9.92,
            "currency" : "USD"
          },
          "type" : "base"
        } ],
        "extra_info" : null,
        "service_name" : "USPS Priority Mail",
        "service_type" : "vesyl_priority_mail",
        "shipper_account" : {
          "description" : "[VESYL] Testing account",
          "id" : "f05425c9c542400daa98e02e38bfcced",
          "slug" : "vesyl"
        },
        "total_charge" : {
          "amount" : 9.92,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "internal_custom_fields" : {
          "tracking_slug" : "usps"
        },
        "shipment_id" : "shp_880b7b256d8db03ed641e36d701be99c",
        "shipper_account" : {
          "custom_fields" : {
            "label_qrcode_enabled" : true,
            "shipment_label_enabled" : true
          }
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 0.001
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 0.5
      },
      "dimension" : [ {
        "depth" : 10,
        "height" : 10,
        "unit" : "cm",
        "width" : 10
      } ],
      "items" : [ {
        "description" : "computer 3 all - white",
        "hs_code" : "560229",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0579/9401/8890/files/5ea530e9e2bf89fb03ce8b45439262a1_ef563b7a-340b-41eb-a8d5-3bd399b853f8.png?v=1721108200" ],
        "item_id" : "054103bb2efb4168a1585d0db2121cce",
        "origin_country" : "USA",
        "price" : {
          "amount" : 78,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "8888422001",
        "weight" : {
          "unit" : "kg",
          "value" : 0.001
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "d9b455fbd80b495d80652711b75e702b",
      "tracking_numbers" : [ "9405500106032109512360" ],
      "organization_id" : "b001d64c4d6343afb6e5bb69580d7661",
      "created_at" : "2025-03-10T09:53:08.038591Z",
      "updated_at" : "2025-03-10T09:53:14.678613Z",
      "succeed_at" : "2025-03-10T09:53:14.667704748Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "78752"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "94104"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "Austin",
        "company_name" : "test company",
        "contact_name" : "USA test test",
        "country" : "USA",
        "email" : "yy.nie+8345834511@aftership.com",
        "postal_code" : "78752",
        "state" : "TX",
        "street1" : "6929 Airport Boulevard",
        "street2" : "110"
      },
      "ship_to" : {
        "city" : "San Francisco",
        "company_name" : "Test company",
        "contact_name" : "test",
        "country" : "USA",
        "phone" : "1112223333",
        "postal_code" : "94104",
        "state" : "CA",
        "street1" : "345 California St"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "3b260d44be8442ed83e98965b025365d",
      "order_number" : "#nyy2171",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/d9b455fb-d80b-495d-8065-2711b75e702b-1741600394408.pdf"
        },
        "qr_code" : {
          "file_type" : "png",
          "paper_size" : "1x1",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-10/d9b455fb-d80b-495d-8065-2711b75e702b-1741600394408.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "test 0306N2O7AFGW" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "vesyl",
      "carrier_account_id" : "f05425c9c542400daa98e02e38bfcced",
      "service_type" : "vesyl_priority_mail",
      "rate" : {
        "booking_cut_off" : null,
        "charge_weight" : {
          "unit" : "lb",
          "value" : 6
        },
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 27.16,
            "currency" : "USD"
          },
          "type" : "base"
        } ],
        "extra_info" : null,
        "service_name" : "USPS Priority Mail",
        "service_type" : "vesyl_priority_mail",
        "shipper_account" : {
          "description" : "[VESYL] Testing account",
          "id" : "f05425c9c542400daa98e02e38bfcced",
          "slug" : "vesyl"
        },
        "total_charge" : {
          "amount" : 27.16,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "internal_custom_fields" : {
          "allocation_id" : "d919cc42cd114a03a9c8a1a85dba64f6",
          "notify_customer" : "true",
          "operator_account_id" : "b120d6eca0764f9cb9f76ff3bf6c1e1d",
          "returns_shipping_type" : "retailer_label",
          "tracking_slug" : "usps"
        },
        "shipment_id" : "shp_7f293563b959f1addc763d23029a4080",
        "shipper_account" : {
          "custom_fields" : {
            "label_qrcode_enabled" : true,
            "shipment_label_enabled" : true
          }
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 5.001
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 6
      },
      "dimension" : [ {
        "depth" : 10,
        "height" : 13,
        "unit" : "in",
        "width" : 12
      } ],
      "items" : [ {
        "description" : "computer 3 all - white",
        "hs_code" : "560229",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0579/9401/8890/files/5ea530e9e2bf89fb03ce8b45439262a1_ef563b7a-340b-41eb-a8d5-3bd399b853f8.png?v=1721108200" ],
        "item_id" : "14959955247178",
        "origin_country" : "USA",
        "price" : {
          "amount" : 78,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "8888422001",
        "weight" : {
          "unit" : "lb",
          "value" : 0.001
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "ce3d077a42284555983e20ab5fb1e7b8",
      "tracking_numbers" : [ "794684517711" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T09:51:13.410959Z",
      "updated_at" : "2025-03-10T09:51:22.213409Z",
      "succeed_at" : "2025-03-10T09:51:22.203868299Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "admin"
      },
      "source" : "admin",
      "ship_from" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "id" : "439c069587b24f8e82969c0c18bbd5a4",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "1bc837b2be844bf8a7897b4e510a59f8",
      "order_number" : "20250310a002",
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-10/43958c8f-97e6-47c8-b048-a6551d010f46-1741600281261998.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/773cea98-d008-4083-9c5f-c6f2bd7c54c0-1741600280833786.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "20250310a002" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
      "service_type" : "fedex_2_day",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "[FedEx] Postmen Testing Account",
          "id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
          "slug" : "fedex"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null,
        "form_id" : 201,
        "internal_custom_fields" : {
          "account_id" : "db910fb96c37438aa6015ff3efda6064",
          "ship_from" : {
            "id" : "439c069587b24f8e82969c0c18bbd5a4"
          }
        },
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 2,
              "height" : 2,
              "unit" : "in",
              "width" : 2
            },
            "items" : [ {
              "description" : "20250310a002",
              "hs_code" : "3132",
              "origin_country" : null,
              "price" : {
                "amount" : 2,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "323",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            }, {
              "description" : "20250310b002",
              "hs_code" : "2212",
              "origin_country" : null,
              "price" : {
                "amount" : 1,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "221",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 8
            }
          } ]
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : {
        "location_name" : "Test US location"
      },
      "weight" : {
        "unit" : "kg",
        "value" : 8
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 2,
        "height" : 2,
        "unit" : "in",
        "width" : 2
      } ],
      "items" : [ {
        "description" : "20250310a002",
        "hs_code" : "3132",
        "price" : {
          "amount" : 2,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "323",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      }, {
        "description" : "20250310b002",
        "hs_code" : "2212",
        "price" : {
          "amount" : 1,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "221",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "db910fb96c37438aa6015ff3efda6064",
      "label_base64" : null
    }, {
      "id" : "0616b29897be43be85b82bd4921baa23",
      "tracking_numbers" : [ "IN-SB-2-C7AYWFVGJBTWDE-AZ" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T09:46:23.255695Z",
      "updated_at" : "2025-03-10T09:46:24.40677Z",
      "succeed_at" : "2025-03-10T09:46:24.39879617Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "IDN",
            "postal_code" : "12150"
          },
          "ship_to" : {
            "country" : "IDN",
            "postal_code" : "12630"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Solo",
        "company_name" : "Watsons Indonesia",
        "contact_name" : "Watsons Indonesia",
        "country" : "IDN",
        "location" : {
          "lat" : -6.261,
          "lng" : 106.8106
        },
        "phone" : "02129528445",
        "postal_code" : "12150",
        "state" : "Jakarta Selatan",
        "street1" : "Jl. Pangeran Antasari No.36",
        "street2" : "NO.36, RT.12/RW.5 RT.12/RW.5"
      },
      "ship_to" : {
        "city" : "Luwu",
        "contact_name" : "Ragil Rika rismaya",
        "country" : "IDN",
        "email" : "rikarismaya830@gmail.com",
        "location" : {
          "lat" : -6.351687999999999,
          "lng" : 106.7995192
        },
        "phone" : "81296977330",
        "postal_code" : "12630",
        "state" : "Dki Jakarta",
        "street1" : "Jl. Aselih No.53 C. RT.7/RW.1, Cipedak Jagakarsa Cipedak    ",
        "street3" : "Jakarta Selatan"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "IDN",
      "ship_to_country" : "IDN",
      "order_id" : null,
      "order_number" : "11264205",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/a120ca11-dfd6-4ea7-9de4-04285c3bf0d2-1741599984041365.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "398299" ],
      "carrier_references" : null,
      "carrier_account_slug" : "grab",
      "carrier_account_id" : "ace77747ebdc4d03ba47e88b275ed5bb",
      "service_type" : "grab_instant",
      "rate" : {
        "charge_weight" : {
          "unit" : "kg",
          "value" : 10
        },
        "delivery_date" : "2025-03-10T10:10:46Z",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 100000,
            "currency" : "IDR"
          },
          "type" : "base"
        } ],
        "pickup_deadline" : "2025-03-10T09:40:46Z",
        "service_name" : "Grab Instant",
        "service_type" : "grab_instant",
        "shipper_account" : {
          "description" : "[grab] watsons sandbox account",
          "id" : "ace77747ebdc4d03ba47e88b275ed5bb",
          "slug" : "grab"
        },
        "total_charge" : {
          "amount" : 100000,
          "currency" : "IDR"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : {
          "amount" : 100000,
          "currency" : "IDR"
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 10
      },
      "chargeable_weight" : {
        "unit" : "kg",
        "value" : 10
      },
      "dimension" : [ {
        "depth" : 30,
        "height" : 30,
        "unit" : "cm",
        "width" : 30
      } ],
      "items" : [ {
        "description" : "Infallible 24H Matte Cover Foundation (Hasil Natural  Tahan Lama) - 123 Natural Vanilla",
        "item_id" : "BP_24561",
        "origin_country" : "IDN",
        "price" : {
          "amount" : 170500,
          "currency" : "IDR"
        },
        "quantity" : 1,
        "sku" : "24561",
        "weight" : {
          "unit" : "kg",
          "value" : 0.001
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "4034feb70fa645a2af0ef6391fcfa00d",
      "tracking_numbers" : [ "794684517413" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T09:43:40.794607Z",
      "updated_at" : "2025-03-10T09:43:51.177222Z",
      "succeed_at" : "2025-03-10T09:43:51.169632741Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "84101"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "84101"
          }
        },
        "retention_days" : "365",
        "source" : "admin"
      },
      "source" : "admin",
      "ship_from" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "id" : "439c069587b24f8e82969c0c18bbd5a4",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Salt Lake City",
        "company_name" : "Test company",
        "contact_name" : "Contact name",
        "country" : "USA",
        "email" : "test@test.com",
        "phone" : "18015325501",
        "postal_code" : "84101",
        "state" : "UT",
        "street1" : "230 W 200 S LBBY",
        "type" : "business"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "b011bdbf82094a20a94dfec6c747d7f4",
      "order_number" : "20250310a001",
      "return_shipment" : false,
      "files" : {
        "invoice" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/invoice/2025-03-10/ed454968-1a22-4f4f-bb97-fbf1fd89a05d-1741599830197264.pdf"
        },
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "a4",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/5eddbca9-994f-4bb9-af40-8dc8c5afabe5-1741599825214451.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "20250310a001" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
      "service_type" : "fedex_2_day",
      "rate" : {
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "[FedEx] Postmen Testing Account",
          "id" : "ed05300d-91c1-4dae-ba3b-fedbcd270b71",
          "slug" : "fedex"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : null,
        "form_id" : 201,
        "internal_custom_fields" : {
          "account_id" : "db910fb96c37438aa6015ff3efda6064",
          "ship_from" : {
            "id" : "439c069587b24f8e82969c0c18bbd5a4"
          }
        },
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 2,
              "height" : 2,
              "unit" : "in",
              "width" : 2
            },
            "items" : [ {
              "description" : "20250310a001",
              "hs_code" : "323",
              "origin_country" : null,
              "price" : {
                "amount" : 2,
                "currency" : "USD"
              },
              "quantity" : 2,
              "sku" : "31231",
              "weight" : {
                "unit" : "kg",
                "value" : 2
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 4
            }
          } ]
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : {
        "location_name" : "Test US location"
      },
      "weight" : {
        "unit" : "kg",
        "value" : 4
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 2,
        "height" : 2,
        "unit" : "in",
        "width" : 2
      } ],
      "items" : [ {
        "description" : "20250310a001",
        "hs_code" : "323",
        "price" : {
          "amount" : 2,
          "currency" : "USD"
        },
        "quantity" : 2,
        "sku" : "31231",
        "weight" : {
          "unit" : "kg",
          "value" : 2
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "db910fb96c37438aa6015ff3efda6064",
      "label_base64" : null
    }, {
      "id" : "4011e90ffd894338a47b00d6e8d323c0",
      "tracking_numbers" : [ "IN-SB-2-C7AYWFVGJBTWDE-AZ" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T09:40:45.314559Z",
      "updated_at" : "2025-03-10T09:40:47.811937Z",
      "succeed_at" : "2025-03-10T09:40:47.803722871Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "IDN",
            "postal_code" : "12150"
          },
          "ship_to" : {
            "country" : "IDN",
            "postal_code" : "12630"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Solo",
        "company_name" : "Watsons Indonesia",
        "contact_name" : "Watsons Indonesia",
        "country" : "IDN",
        "location" : {
          "lat" : -6.261,
          "lng" : 106.8106
        },
        "phone" : "02129528445",
        "postal_code" : "12150",
        "state" : "Jakarta Selatan",
        "street1" : "Jl. Pangeran Antasari No.36",
        "street2" : "NO.36, RT.12/RW.5 RT.12/RW.5"
      },
      "ship_to" : {
        "city" : "Luwu",
        "contact_name" : "Ragil Rika rismaya",
        "country" : "IDN",
        "email" : "rikarismaya830@gmail.com",
        "location" : {
          "lat" : -6.351687999999999,
          "lng" : 106.7995192
        },
        "phone" : "81296977330",
        "postal_code" : "12630",
        "state" : "Dki Jakarta",
        "street1" : "Jl. Aselih No.53 C. RT.7/RW.1, Cipedak Jagakarsa Cipedak    ",
        "street3" : "Jakarta Selatan"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "IDN",
      "ship_to_country" : "IDN",
      "order_id" : null,
      "order_number" : "11264205",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "default",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/a56927b0-1fae-4000-a141-8f4dd97e2f10-1741599647296758.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "398299" ],
      "carrier_references" : null,
      "carrier_account_slug" : "grab",
      "carrier_account_id" : "ace77747ebdc4d03ba47e88b275ed5bb",
      "service_type" : "grab_instant",
      "rate" : {
        "charge_weight" : {
          "unit" : "kg",
          "value" : 10
        },
        "delivery_date" : "2025-03-10T10:10:46Z",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 100000,
            "currency" : "IDR"
          },
          "type" : "base"
        } ],
        "pickup_deadline" : "2025-03-10T09:40:46Z",
        "service_name" : "Grab Instant",
        "service_type" : "grab_instant",
        "shipper_account" : {
          "description" : "[grab] watsons sandbox account",
          "id" : "ace77747ebdc4d03ba47e88b275ed5bb",
          "slug" : "grab"
        },
        "total_charge" : {
          "amount" : 100000,
          "currency" : "IDR"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : {
          "amount" : 100000,
          "currency" : "IDR"
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 10
      },
      "chargeable_weight" : {
        "unit" : "kg",
        "value" : 10
      },
      "dimension" : [ {
        "depth" : 30,
        "height" : 30,
        "unit" : "cm",
        "width" : 30
      } ],
      "items" : [ {
        "description" : "Infallible 24H Matte Cover Foundation (Hasil Natural  Tahan Lama) - 123 Natural Vanilla",
        "item_id" : "BP_24561",
        "origin_country" : "IDN",
        "price" : {
          "amount" : 170500,
          "currency" : "IDR"
        },
        "quantity" : 1,
        "sku" : "24561",
        "weight" : {
          "unit" : "kg",
          "value" : 0.001
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "5ad2218a33d4443b9fbef62824febe80",
      "tracking_numbers" : [ "13030653" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T09:09:53.002079Z",
      "updated_at" : "2025-03-10T09:09:54.812685Z",
      "succeed_at" : "2025-03-10T09:09:54.804371799Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP",
            "postal_code" : "08150"
          },
          "ship_to" : {
            "country" : "FRA",
            "postal_code" : "75001"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles(barcelona)",
        "company_name" : "Panda of china",
        "contact_name" : "Javid",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08150",
        "state" : "Mixed",
        "street1" : "ho no 115/flat no b1 street no 7.",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Paris",
        "contact_name" : "Ismael Tejón",
        "country" : "FRA",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "75001",
        "state" : "VIL",
        "street1" : "Calle de prueba, 34.",
        "type" : "residential"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "ESP",
      "ship_to_country" : "FRA",
      "order_id" : null,
      "order_number" : "100001",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/5ad2218a-33d4-443b-9fbe-f62824febe80-1741597794058.pdf"
        }
      },
      "box_type" : "nacex_sample",
      "references" : [ "REFERENCE0" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "nacex",
      "carrier_account_id" : "c5573249c14b4f9eb077e36985b2f6e9",
      "service_type" : "nacex_servicio-aereo",
      "rate" : {
        "delivery_date" : "2025-03-12",
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "NACEX SERVICIO AEREO",
        "service_type" : "nacex_servicio-aereo",
        "shipper_account" : {
          "description" : "[nacex] Demo account new",
          "id" : "c5573249c14b4f9eb077e36985b2f6e9",
          "slug" : "nacex"
        },
        "transit_time" : 2
      },
      "service_options" : [ {
        "end_time" : "13:00:00",
        "start_time" : "12:00:00",
        "type" : "pickup"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment_code" : "429687879"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "ESP",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "a2e45ef6641c4e7b8d6d43df323dc21b",
      "tracking_numbers" : [ "13030596" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T08:36:26.027375Z",
      "updated_at" : "2025-03-10T08:36:27.783043Z",
      "succeed_at" : "2025-03-10T08:36:27.77321553Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP",
            "postal_code" : "08150"
          },
          "ship_to" : {
            "country" : "FRA",
            "postal_code" : "75001"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles(barcelona)",
        "company_name" : "Panda of china",
        "contact_name" : "Javid",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08150",
        "state" : "Mixed",
        "street1" : "ho no 115/flat no b1 street no 7.",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Paris",
        "contact_name" : "Ismael Tejón",
        "country" : "FRA",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "75001",
        "state" : "VIL",
        "street1" : "Calle de prueba, 34.",
        "type" : "residential"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "ESP",
      "ship_to_country" : "FRA",
      "order_id" : null,
      "order_number" : "100001",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/a2e45ef6-641c-4e7b-8d6d-43df323dc21b-1741595787042.pdf"
        }
      },
      "box_type" : "nacex_sample",
      "references" : [ "REFERENCE0" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "nacex",
      "carrier_account_id" : "c5573249c14b4f9eb077e36985b2f6e9",
      "service_type" : "nacex_servicio-aereo",
      "rate" : {
        "delivery_date" : "2025-03-12",
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "NACEX SERVICIO AEREO",
        "service_type" : "nacex_servicio-aereo",
        "shipper_account" : {
          "description" : "[nacex] Demo account new",
          "id" : "c5573249c14b4f9eb077e36985b2f6e9",
          "slug" : "nacex"
        },
        "transit_time" : 2
      },
      "service_options" : [ {
        "end_time" : "13:00:00",
        "start_time" : "12:00:00",
        "type" : "pickup"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment_code" : "429680939"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "ESP",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "release-core-sandbox-api.aftership.io",
      "label_base64" : null
    }, {
      "id" : "548d1528701d4c298ab8e75ffbd863eb",
      "tracking_numbers" : [ "795495344291" ],
      "organization_id" : "4e5b6e3b1027486b8afac661f92e7718",
      "created_at" : "2025-03-10T08:31:28.377103Z",
      "updated_at" : "2025-03-10T08:31:40.375262Z",
      "succeed_at" : "2025-03-10T08:31:40.365972295Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "12231"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "75050"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "New York",
        "contact_name" : "Lynne huang",
        "country" : "USA",
        "email" : "jl.huang@aftership.com",
        "phone" : "18444744726",
        "postal_code" : "12231",
        "state" : "NY",
        "street1" : "New York"
      },
      "ship_to" : {
        "city" : "Grand Prairie",
        "contact_name" : "fate-test-store",
        "country" : "USA",
        "phone" : "18444744726",
        "postal_code" : "75050",
        "state" : "TX",
        "street1" : "202 N Great SW Pkwy",
        "street2" : "Doors 10-11 Door 1"
      },
      "ship_date" : "2025-03-09",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "3e37a5e74d6e481b985b135b4daab003",
      "order_number" : "#1164#",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/87a26491-f1b5-4c34-a350-db1e415360ea-1741595499535931.pdf"
        },
        "qr_code" : {
          "file_type" : "png",
          "paper_size" : "1x1",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-10/548d1528-701d-4c29-8ab8-e75ffbd863eb-1741595499175820.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA#0QFG218W", "Order##1164#", "Source:AfterShip Returns" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "bd64045f-f5a6-4ee1-9b1e-e039bdf5e19b",
      "service_type" : "fedex_2_day",
      "rate" : {
        "charge_weight" : {
          "unit" : "lb",
          "value" : 51
        },
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 422.06,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 1.05,
            "currency" : "USD"
          },
          "type" : "return_printed_label"
        }, {
          "charge" : {
            "amount" : 55,
            "currency" : "USD"
          },
          "type" : "additional_handling_surcharge_-_weight"
        }, {
          "charge" : {
            "amount" : 23.85,
            "currency" : "USD"
          },
          "type" : "fuel"
        } ],
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "TEST",
          "id" : "bd64045f-f5a6-4ee1-9b1e-e039bdf5e19b",
          "slug" : "fedex"
        },
        "total_charge" : {
          "amount" : 501.96,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : {
          "amount" : 501.96,
          "currency" : "USD"
        },
        "form_id" : "0221",
        "internal_custom_fields" : {
          "allocation_id" : "90a3eac059bb4d289a710fedd0c47bec",
          "notify_customer" : "true",
          "returns_shipping_type" : "retailer_label",
          "tracking_slug" : "fedex"
        },
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 1,
              "height" : 1,
              "unit" : "in",
              "width" : 1
            },
            "items" : [ {
              "description" : "One Men Sofa Chair - Ship in 3 boxes",
              "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/9446/5320/files/chair.jpg?v=1691120900" ],
              "item_id" : "16100360487208",
              "origin_country" : null,
              "price" : {
                "amount" : 85,
                "currency" : "USD"
              },
              "quantity" : 1,
              "sku" : "45878652010792",
              "weight" : {
                "unit" : "lb",
                "value" : 50
              }
            } ],
            "weight" : {
              "unit" : "lb",
              "value" : 50.5
            }
          } ]
        },
        "shipper_account" : {
          "custom_fields" : {
            "label_qrcode_enabled" : true,
            "shipment_label_enabled" : true
          }
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 50.5
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 51
      },
      "dimension" : [ {
        "depth" : 1,
        "height" : 1,
        "unit" : "in",
        "width" : 1
      } ],
      "items" : [ {
        "description" : "One Men Sofa Chair - Ship in 3 boxes",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/9446/5320/files/chair.jpg?v=1691120900" ],
        "item_id" : "16100360487208",
        "price" : {
          "amount" : 85,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "45878652010792",
        "weight" : {
          "unit" : "lb",
          "value" : 50
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "564c879c5be34f418f5ca9825e0e16a5",
      "tracking_numbers" : [ "795495344280" ],
      "organization_id" : "4e5b6e3b1027486b8afac661f92e7718",
      "created_at" : "2025-03-10T08:31:28.373343Z",
      "updated_at" : "2025-03-10T08:31:40.398958Z",
      "succeed_at" : "2025-03-10T08:31:40.388634315Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "12231"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "75050"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "New York",
        "contact_name" : "Lynne huang",
        "country" : "USA",
        "email" : "jl.huang@aftership.com",
        "phone" : "18444744726",
        "postal_code" : "12231",
        "state" : "NY",
        "street1" : "New York"
      },
      "ship_to" : {
        "city" : "Grand Prairie",
        "contact_name" : "fate-test-store",
        "country" : "USA",
        "phone" : "18444744726",
        "postal_code" : "75050",
        "state" : "TX",
        "street1" : "202 N Great SW Pkwy",
        "street2" : "Doors 10-11 Door 1"
      },
      "ship_date" : "2025-03-09",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "3e37a5e74d6e481b985b135b4daab003",
      "order_number" : "#1164#",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/f0953169-11a4-40b7-9e09-d728c86dcfa2-1741595499682491.pdf"
        },
        "qr_code" : {
          "file_type" : "png",
          "paper_size" : "1x1",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-10/564c879c-5be3-4f41-8f5c-a9825e0e16a5-1741595499154584.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA#0QFG218W", "Order##1164#", "Source:AfterShip Returns" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "bd64045f-f5a6-4ee1-9b1e-e039bdf5e19b",
      "service_type" : "fedex_2_day",
      "rate" : {
        "charge_weight" : {
          "unit" : "lb",
          "value" : 51
        },
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 422.06,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 1.05,
            "currency" : "USD"
          },
          "type" : "return_printed_label"
        }, {
          "charge" : {
            "amount" : 55,
            "currency" : "USD"
          },
          "type" : "additional_handling_surcharge_-_weight"
        }, {
          "charge" : {
            "amount" : 23.85,
            "currency" : "USD"
          },
          "type" : "fuel"
        } ],
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "TEST",
          "id" : "bd64045f-f5a6-4ee1-9b1e-e039bdf5e19b",
          "slug" : "fedex"
        },
        "total_charge" : {
          "amount" : 501.96,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : {
          "amount" : 501.96,
          "currency" : "USD"
        },
        "form_id" : "0221",
        "internal_custom_fields" : {
          "allocation_id" : "b62376136ab34e4ab1341ac6916019c4",
          "notify_customer" : "true",
          "returns_shipping_type" : "retailer_label",
          "tracking_slug" : "fedex"
        },
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 1,
              "height" : 1,
              "unit" : "in",
              "width" : 1
            },
            "items" : [ {
              "description" : "One Men Sofa Chair - Ship in 3 boxes",
              "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/9446/5320/files/chair.jpg?v=1691120900" ],
              "item_id" : "16100360487208",
              "origin_country" : null,
              "price" : {
                "amount" : 85,
                "currency" : "USD"
              },
              "quantity" : 1,
              "sku" : "45878652010792",
              "weight" : {
                "unit" : "lb",
                "value" : 50
              }
            } ],
            "weight" : {
              "unit" : "lb",
              "value" : 50.5
            }
          } ]
        },
        "shipper_account" : {
          "custom_fields" : {
            "label_qrcode_enabled" : true,
            "shipment_label_enabled" : true
          }
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 50.5
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 51
      },
      "dimension" : [ {
        "depth" : 1,
        "height" : 1,
        "unit" : "in",
        "width" : 1
      } ],
      "items" : [ {
        "description" : "One Men Sofa Chair - Ship in 3 boxes",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/9446/5320/files/chair.jpg?v=1691120900" ],
        "item_id" : "16100360487208",
        "price" : {
          "amount" : 85,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "45878652010792",
        "weight" : {
          "unit" : "lb",
          "value" : 50
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "1921b2b7165d4f5c97c60be2ec4be890",
      "tracking_numbers" : [ "795495344269" ],
      "organization_id" : "4e5b6e3b1027486b8afac661f92e7718",
      "created_at" : "2025-03-10T08:29:13.426079Z",
      "updated_at" : "2025-03-10T08:29:25.270804Z",
      "succeed_at" : "2025-03-10T08:29:25.257649674Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "12231"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "75050"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "New York",
        "contact_name" : "Lynne huang",
        "country" : "USA",
        "email" : "jl.huang@aftership.com",
        "phone" : "18444744726",
        "postal_code" : "12231",
        "state" : "NY",
        "street1" : "New York"
      },
      "ship_to" : {
        "city" : "Grand Prairie",
        "contact_name" : "fate-test-store",
        "country" : "USA",
        "phone" : "18444744726",
        "postal_code" : "75050",
        "state" : "TX",
        "street1" : "202 N Great SW Pkwy",
        "street2" : "Doors 10-11 Door 1"
      },
      "ship_date" : "2025-03-09",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "3e37a5e74d6e481b985b135b4daab003",
      "order_number" : "#1164#",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/0bcf804e-041d-499b-8683-ac0831d8090a-1741595364559279.pdf"
        },
        "qr_code" : {
          "file_type" : "png",
          "paper_size" : "1x1",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-10/1921b2b7-165d-4f5c-97c6-0be2ec4be890-1741595364143634.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA#VM859XQL", "Order##1164#", "Source:AfterShip Returns" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "bd64045f-f5a6-4ee1-9b1e-e039bdf5e19b",
      "service_type" : "fedex_2_day",
      "rate" : {
        "charge_weight" : {
          "unit" : "lb",
          "value" : 51
        },
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 422.06,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 1.05,
            "currency" : "USD"
          },
          "type" : "return_printed_label"
        }, {
          "charge" : {
            "amount" : 55,
            "currency" : "USD"
          },
          "type" : "additional_handling_surcharge_-_weight"
        }, {
          "charge" : {
            "amount" : 23.85,
            "currency" : "USD"
          },
          "type" : "fuel"
        } ],
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "TEST",
          "id" : "bd64045f-f5a6-4ee1-9b1e-e039bdf5e19b",
          "slug" : "fedex"
        },
        "total_charge" : {
          "amount" : 501.96,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : {
          "amount" : 501.96,
          "currency" : "USD"
        },
        "form_id" : "0221",
        "internal_custom_fields" : {
          "allocation_id" : "2f19f3b80b22404cb70e6557427e30c8",
          "notify_customer" : "true",
          "operator_account_id" : "302b5d7c3c7d4db5b1801007fda3fcd2",
          "returns_shipping_type" : "retailer_label",
          "tracking_slug" : "fedex"
        },
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 1,
              "height" : 1,
              "unit" : "in",
              "width" : 1
            },
            "items" : [ {
              "description" : "One Men Sofa Chair - Ship in 3 boxes",
              "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/9446/5320/files/chair.jpg?v=1691120900" ],
              "item_id" : "16100360487208",
              "origin_country" : null,
              "price" : {
                "amount" : 85,
                "currency" : "USD"
              },
              "quantity" : 1,
              "sku" : "45878652010792",
              "weight" : {
                "unit" : "lb",
                "value" : 50
              }
            } ],
            "weight" : {
              "unit" : "lb",
              "value" : 50.5
            }
          } ]
        },
        "shipper_account" : {
          "custom_fields" : {
            "label_qrcode_enabled" : true,
            "shipment_label_enabled" : true
          }
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 50.5
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 51
      },
      "dimension" : [ {
        "depth" : 1,
        "height" : 1,
        "unit" : "in",
        "width" : 1
      } ],
      "items" : [ {
        "description" : "One Men Sofa Chair - Ship in 3 boxes",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/9446/5320/files/chair.jpg?v=1691120900" ],
        "item_id" : "16100360487208",
        "price" : {
          "amount" : 85,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "45878652010792",
        "weight" : {
          "unit" : "lb",
          "value" : 50
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "e9e46449a63f456395c856693f4e49dd",
      "tracking_numbers" : [ "795495344270" ],
      "organization_id" : "4e5b6e3b1027486b8afac661f92e7718",
      "created_at" : "2025-03-10T08:29:13.410255Z",
      "updated_at" : "2025-03-10T08:29:24.998157Z",
      "succeed_at" : "2025-03-10T08:29:24.99060032Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "12231"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "75050"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "New York",
        "contact_name" : "Lynne huang",
        "country" : "USA",
        "email" : "jl.huang@aftership.com",
        "phone" : "18444744726",
        "postal_code" : "12231",
        "state" : "NY",
        "street1" : "New York"
      },
      "ship_to" : {
        "city" : "Grand Prairie",
        "contact_name" : "fate-test-store",
        "country" : "USA",
        "phone" : "18444744726",
        "postal_code" : "75050",
        "state" : "TX",
        "street1" : "202 N Great SW Pkwy",
        "street2" : "Doors 10-11 Door 1"
      },
      "ship_date" : "2025-03-09",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "3e37a5e74d6e481b985b135b4daab003",
      "order_number" : "#1164#",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/abe05b30-fa6b-49f8-8df9-d006aeb6bff5-1741595364119273.pdf"
        },
        "qr_code" : {
          "file_type" : "png",
          "paper_size" : "1x1",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-10/e9e46449-a63f-4563-95c8-56693f4e49dd-1741595363752472.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA#VM859XQL", "Order##1164#", "Source:AfterShip Returns" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "bd64045f-f5a6-4ee1-9b1e-e039bdf5e19b",
      "service_type" : "fedex_2_day",
      "rate" : {
        "charge_weight" : {
          "unit" : "lb",
          "value" : 51
        },
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 422.06,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 1.05,
            "currency" : "USD"
          },
          "type" : "return_printed_label"
        }, {
          "charge" : {
            "amount" : 55,
            "currency" : "USD"
          },
          "type" : "additional_handling_surcharge_-_weight"
        }, {
          "charge" : {
            "amount" : 23.85,
            "currency" : "USD"
          },
          "type" : "fuel"
        } ],
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "TEST",
          "id" : "bd64045f-f5a6-4ee1-9b1e-e039bdf5e19b",
          "slug" : "fedex"
        },
        "total_charge" : {
          "amount" : 501.96,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : {
          "amount" : 501.96,
          "currency" : "USD"
        },
        "form_id" : "0221",
        "internal_custom_fields" : {
          "allocation_id" : "963786b41dc349b9a67a1ee000bae3eb",
          "notify_customer" : "true",
          "operator_account_id" : "302b5d7c3c7d4db5b1801007fda3fcd2",
          "returns_shipping_type" : "retailer_label",
          "tracking_slug" : "fedex"
        },
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 1,
              "height" : 1,
              "unit" : "in",
              "width" : 1
            },
            "items" : [ {
              "description" : "One Men Sofa Chair - Ship in 3 boxes",
              "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/9446/5320/files/chair.jpg?v=1691120900" ],
              "item_id" : "16100360487208",
              "origin_country" : null,
              "price" : {
                "amount" : 85,
                "currency" : "USD"
              },
              "quantity" : 1,
              "sku" : "45878652010792",
              "weight" : {
                "unit" : "lb",
                "value" : 50
              }
            } ],
            "weight" : {
              "unit" : "lb",
              "value" : 50.5
            }
          } ]
        },
        "shipper_account" : {
          "custom_fields" : {
            "label_qrcode_enabled" : true,
            "shipment_label_enabled" : true
          }
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 50.5
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 51
      },
      "dimension" : [ {
        "depth" : 1,
        "height" : 1,
        "unit" : "in",
        "width" : 1
      } ],
      "items" : [ {
        "description" : "One Men Sofa Chair - Ship in 3 boxes",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/9446/5320/files/chair.jpg?v=1691120900" ],
        "item_id" : "16100360487208",
        "price" : {
          "amount" : 85,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "45878652010792",
        "weight" : {
          "unit" : "lb",
          "value" : 50
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "fbae18b3034b4280be181951997293c0",
      "tracking_numbers" : [ "795495344247" ],
      "organization_id" : "4e5b6e3b1027486b8afac661f92e7718",
      "created_at" : "2025-03-10T08:27:57.908911Z",
      "updated_at" : "2025-03-10T08:28:16.494926Z",
      "succeed_at" : "2025-03-10T08:28:16.487040469Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "12231"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "75050"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "New York",
        "contact_name" : "Lynne huang",
        "country" : "USA",
        "email" : "jl.huang@aftership.com",
        "phone" : "18444744726",
        "postal_code" : "12231",
        "state" : "NY",
        "street1" : "New York"
      },
      "ship_to" : {
        "city" : "Grand Prairie",
        "contact_name" : "fate-test-store",
        "country" : "USA",
        "phone" : "18444744726",
        "postal_code" : "75050",
        "state" : "TX",
        "street1" : "202 N Great SW Pkwy",
        "street2" : "Doors 10-11 Door 1"
      },
      "ship_date" : "2025-03-09",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "3e37a5e74d6e481b985b135b4daab003",
      "order_number" : "#1164#",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/d4f556df-5f6b-493b-a33b-a773b9f6d370-1741595295302046.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA#O1549A8S", "Order##1164#", "Source:AfterShip Returns" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "bd64045f-f5a6-4ee1-9b1e-e039bdf5e19b",
      "service_type" : "fedex_2_day",
      "rate" : {
        "charge_weight" : {
          "unit" : "lb",
          "value" : 51
        },
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 422.06,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 1.05,
            "currency" : "USD"
          },
          "type" : "return_printed_label"
        }, {
          "charge" : {
            "amount" : 55,
            "currency" : "USD"
          },
          "type" : "additional_handling_surcharge_-_weight"
        }, {
          "charge" : {
            "amount" : 23.85,
            "currency" : "USD"
          },
          "type" : "fuel"
        } ],
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "TEST",
          "id" : "bd64045f-f5a6-4ee1-9b1e-e039bdf5e19b",
          "slug" : "fedex"
        },
        "total_charge" : {
          "amount" : 501.96,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : {
          "amount" : 501.96,
          "currency" : "USD"
        },
        "form_id" : "0221",
        "internal_custom_fields" : {
          "allocation_id" : "78acea1f51d6481995f6ab312e4c25d0",
          "notify_customer" : "true",
          "returns_shipping_type" : "retailer_label",
          "tracking_slug" : "fedex"
        },
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 1,
              "height" : 1,
              "unit" : "in",
              "width" : 1
            },
            "items" : [ {
              "description" : "One Men Sofa Chair - Ship in 3 boxes",
              "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/9446/5320/files/chair.jpg?v=1691120900" ],
              "item_id" : "16100360487208",
              "origin_country" : null,
              "price" : {
                "amount" : 85,
                "currency" : "USD"
              },
              "quantity" : 1,
              "sku" : "45878652010792",
              "weight" : {
                "unit" : "lb",
                "value" : 50
              }
            } ],
            "weight" : {
              "unit" : "lb",
              "value" : 50.5
            }
          } ]
        },
        "shipper_account" : {
          "custom_fields" : {
            "label_qrcode_enabled" : true,
            "shipment_label_enabled" : true
          }
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 50.5
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 51
      },
      "dimension" : [ {
        "depth" : 1,
        "height" : 1,
        "unit" : "in",
        "width" : 1
      } ],
      "items" : [ {
        "description" : "One Men Sofa Chair - Ship in 3 boxes",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/9446/5320/files/chair.jpg?v=1691120900" ],
        "item_id" : "16100360487208",
        "price" : {
          "amount" : 85,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "45878652010792",
        "weight" : {
          "unit" : "lb",
          "value" : 50
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "f426a31e955e4c538636d81b5dfcc764",
      "tracking_numbers" : [ "795495344258" ],
      "organization_id" : "4e5b6e3b1027486b8afac661f92e7718",
      "created_at" : "2025-03-10T08:27:57.900847Z",
      "updated_at" : "2025-03-10T08:28:17.667743Z",
      "succeed_at" : "2025-03-10T08:28:17.657351871Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "12231"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "75050"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "New York",
        "contact_name" : "Lynne huang",
        "country" : "USA",
        "email" : "jl.huang@aftership.com",
        "phone" : "18444744726",
        "postal_code" : "12231",
        "state" : "NY",
        "street1" : "New York"
      },
      "ship_to" : {
        "city" : "Grand Prairie",
        "contact_name" : "fate-test-store",
        "country" : "USA",
        "phone" : "18444744726",
        "postal_code" : "75050",
        "state" : "TX",
        "street1" : "202 N Great SW Pkwy",
        "street2" : "Doors 10-11 Door 1"
      },
      "ship_date" : "2025-03-09",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "3e37a5e74d6e481b985b135b4daab003",
      "order_number" : "#1164#",
      "return_shipment" : true,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/1046984d-a0ce-4bf2-b2bd-32955263c9f8-1741595296923437.pdf"
        },
        "qr_code" : {
          "file_type" : "png",
          "paper_size" : "1x1",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-10/f426a31e-955e-4c53-8636-d81b5dfcc764-1741595296531749.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA#O1549A8S", "Order##1164#", "Source:AfterShip Returns" ],
      "carrier_references" : null,
      "carrier_account_slug" : "fedex",
      "carrier_account_id" : "bd64045f-f5a6-4ee1-9b1e-e039bdf5e19b",
      "service_type" : "fedex_2_day",
      "rate" : {
        "charge_weight" : {
          "unit" : "lb",
          "value" : 51
        },
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 422.06,
            "currency" : "USD"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 1.05,
            "currency" : "USD"
          },
          "type" : "return_printed_label"
        }, {
          "charge" : {
            "amount" : 55,
            "currency" : "USD"
          },
          "type" : "additional_handling_surcharge_-_weight"
        }, {
          "charge" : {
            "amount" : 23.85,
            "currency" : "USD"
          },
          "type" : "fuel"
        } ],
        "service_name" : "FedEx 2Day®",
        "service_type" : "fedex_2_day",
        "shipper_account" : {
          "description" : "TEST",
          "id" : "bd64045f-f5a6-4ee1-9b1e-e039bdf5e19b",
          "slug" : "fedex"
        },
        "total_charge" : {
          "amount" : 501.96,
          "currency" : "USD"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "courier_charge" : {
          "amount" : 501.96,
          "currency" : "USD"
        },
        "form_id" : "0221",
        "internal_custom_fields" : {
          "allocation_id" : "45c63533f6664624ac9d76b6e7e5babc",
          "notify_customer" : "true",
          "returns_shipping_type" : "retailer_label",
          "tracking_slug" : "fedex"
        },
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "dimension" : {
              "depth" : 1,
              "height" : 1,
              "unit" : "in",
              "width" : 1
            },
            "items" : [ {
              "description" : "One Men Sofa Chair - Ship in 3 boxes",
              "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/9446/5320/files/chair.jpg?v=1691120900" ],
              "item_id" : "16100360487208",
              "origin_country" : null,
              "price" : {
                "amount" : 85,
                "currency" : "USD"
              },
              "quantity" : 1,
              "sku" : "45878652010792",
              "weight" : {
                "unit" : "lb",
                "value" : 50
              }
            } ],
            "weight" : {
              "unit" : "lb",
              "value" : 50.5
            }
          } ]
        },
        "shipper_account" : {
          "custom_fields" : {
            "label_qrcode_enabled" : true,
            "shipment_label_enabled" : true
          }
        },
        "tracking_id_type" : "FEDEX"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 50.5
      },
      "chargeable_weight" : {
        "unit" : "lb",
        "value" : 51
      },
      "dimension" : [ {
        "depth" : 1,
        "height" : 1,
        "unit" : "in",
        "width" : 1
      } ],
      "items" : [ {
        "description" : "One Men Sofa Chair - Ship in 3 boxes",
        "image_urls" : [ "https://cdn.shopify.com/s/files/1/0788/9446/5320/files/chair.jpg?v=1691120900" ],
        "item_id" : "16100360487208",
        "price" : {
          "amount" : 85,
          "currency" : "USD"
        },
        "quantity" : 1,
        "sku" : "45878652010792",
        "weight" : {
          "unit" : "lb",
          "value" : 50
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "feaf5615350944acbbc718303b854327",
      "tracking_numbers" : null,
      "organization_id" : "c21ce0173f854533a929f800ed495b6a",
      "created_at" : "2025-03-10T08:02:28.132959Z",
      "updated_at" : "2025-03-10T08:02:28.878573Z",
      "succeed_at" : "2025-03-10T08:02:28.871490762Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "35075"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "35075"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "Austin",
        "company_name" : "AfterShip",
        "contact_name" : "longlongleo",
        "country" : "USA",
        "email" : "y.yang@aftership.com",
        "phone" : "19876543213",
        "postal_code" : "35075",
        "state" : "CA",
        "street1" : "Marktplatz 1"
      },
      "ship_to" : {
        "city" : "Austin",
        "contact_name" : "longlongleo",
        "country" : "USA",
        "email" : "sh.huang+returns01@aftership.com",
        "phone" : "19876543213",
        "postal_code" : "35075",
        "state" : "CA",
        "street1" : "Marktplatz 1"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "f3e1a2563bc64b419243730db500d808",
      "order_number" : "471",
      "return_shipment" : true,
      "files" : {
        "qr_code" : {
          "file_type" : "png",
          "paper_size" : "2x2",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-10/feaf5615-3509-44ac-bbc7-18303b854327-1741593748761.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA#3OS3EAUS Order#471", "demo for US - 3 x 1 Looks differen" ],
      "carrier_references" : [ {
        "dropoff_number" : "HRY2YV2F"
      } ],
      "carrier_account_slug" : "happy-returns",
      "carrier_account_id" : "15ed0573-3e6d-4753-a39a-49f973f9b526",
      "service_type" : "happy-returns_standard",
      "rate" : {
        "booking_cut_off" : null,
        "extra_info" : null,
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Happy Returns Standard",
        "service_type" : "happy-returns_standard",
        "shipper_account" : {
          "description" : "test",
          "id" : "15ed0573-3e6d-4753-a39a-49f973f9b526",
          "slug" : "happy-returns"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "internal_custom_fields" : {
          "returns_shipping_type" : "happy_returns",
          "tracking_slug" : "happy-returns"
        },
        "shipper_account" : {
          "custom_fields" : {
            "shipment_label_enabled" : true
          }
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 0.001
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 10,
        "height" : 10,
        "unit" : "cm",
        "width" : 10
      } ],
      "items" : [ {
        "description" : "demo for US - 3",
        "image_urls" : [ "https://wordpress11.automizelydemo.com/wp-content/uploads/2023/06/蜘蛛侠.jpeg" ],
        "item_id" : "680",
        "price" : {
          "amount" : 9.15,
          "currency" : "GBP"
        },
        "quantity" : 1,
        "return_reason" : "Looks different to image on site",
        "sku" : "testnullvariant004v2-1-7-1-1-1-1-1-1",
        "weight" : {
          "unit" : "kg",
          "value" : 1.0E-6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "8977d66212b440c788d5d6d43ad1f9de",
      "tracking_numbers" : null,
      "organization_id" : "c21ce0173f854533a929f800ed495b6a",
      "created_at" : "2025-03-10T08:02:04.881023Z",
      "updated_at" : "2025-03-10T08:02:06.193208Z",
      "succeed_at" : "2025-03-10T08:02:06.186111269Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "USA",
            "postal_code" : "35075"
          },
          "ship_to" : {
            "country" : "USA",
            "postal_code" : "35075"
          }
        },
        "retention_days" : "365",
        "source" : "returns",
        "tags" : "automatic"
      },
      "source" : "returns",
      "ship_from" : {
        "city" : "Austin",
        "company_name" : "AfterShip",
        "contact_name" : "longlongleo",
        "country" : "USA",
        "email" : "y.yang@aftership.com",
        "phone" : "19876543213",
        "postal_code" : "35075",
        "state" : "CA",
        "street1" : "Marktplatz 1"
      },
      "ship_to" : {
        "city" : "Austin",
        "contact_name" : "longlongleo",
        "country" : "USA",
        "email" : "sh.huang+returns01@aftership.com",
        "phone" : "19876543213",
        "postal_code" : "35075",
        "state" : "CA",
        "street1" : "Marktplatz 1"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "USA",
      "ship_to_country" : "USA",
      "order_id" : "f3e1a2563bc64b419243730db500d808",
      "order_number" : "471",
      "return_shipment" : true,
      "files" : {
        "qr_code" : {
          "file_type" : "png",
          "paper_size" : "2x2",
          "url" : "https://testing-sandbox-download.postmen.io/qr_code/2025-03-10/8977d662-12b4-40c7-88d5-d6d43ad1f9de-1741593725955.png"
        }
      },
      "box_type" : "custom",
      "references" : [ "RMA#AEH19L2V Order#471", "demo for US - 3 x 1 Looks differen" ],
      "carrier_references" : [ {
        "dropoff_number" : "HRSQQVLA"
      } ],
      "carrier_account_slug" : "happy-returns",
      "carrier_account_id" : "15ed0573-3e6d-4753-a39a-49f973f9b526",
      "service_type" : "happy-returns_standard",
      "rate" : {
        "booking_cut_off" : null,
        "extra_info" : null,
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "Happy Returns Standard",
        "service_type" : "happy-returns_standard",
        "shipper_account" : {
          "description" : "test",
          "id" : "15ed0573-3e6d-4753-a39a-49f973f9b526",
          "slug" : "happy-returns"
        }
      },
      "service_options" : null,
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "internal_custom_fields" : {
          "returns_shipping_type" : "happy_returns",
          "tracking_slug" : "happy-returns"
        },
        "shipper_account" : {
          "custom_fields" : {
            "shipment_label_enabled" : true
          }
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "lb",
        "value" : 0.001
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 10,
        "height" : 10,
        "unit" : "cm",
        "width" : 10
      } ],
      "items" : [ {
        "description" : "demo for US - 3",
        "image_urls" : [ "https://wordpress11.automizelydemo.com/wp-content/uploads/2023/06/蜘蛛侠.jpeg" ],
        "item_id" : "680",
        "price" : {
          "amount" : 9.15,
          "currency" : "GBP"
        },
        "quantity" : 1,
        "return_reason" : "Looks different to image on site",
        "sku" : "testnullvariant004v2-1-7-1-1-1-1-1-1",
        "weight" : {
          "unit" : "kg",
          "value" : 1.0E-6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : null,
      "label_base64" : null
    }, {
      "id" : "9b8a10db075744c7a30cc0e2951ce32f",
      "tracking_numbers" : [ "13030540" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T07:58:28.884991Z",
      "updated_at" : "2025-03-10T07:58:30.724803Z",
      "succeed_at" : "2025-03-10T07:58:30.71710857Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP",
            "postal_code" : "08150"
          },
          "ship_to" : {
            "country" : "FRA",
            "postal_code" : "75001"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles(barcelona)",
        "company_name" : "Panda of china",
        "contact_name" : "Javid",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08150",
        "state" : "Mixed",
        "street1" : "ho no 115/flat no b1 street no 7.",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Paris",
        "contact_name" : "Ismael Tejón",
        "country" : "FRA",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "75001",
        "state" : "VIL",
        "street1" : "Calle de prueba, 34.",
        "type" : "residential"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "ESP",
      "ship_to_country" : "FRA",
      "order_id" : null,
      "order_number" : "100001",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/9b8a10db-0757-44c7-a30c-c0e2951ce32f-1741593510044.pdf"
        }
      },
      "box_type" : "nacex_sample",
      "references" : [ "REFERENCE0" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "nacex",
      "carrier_account_id" : "c5573249c14b4f9eb077e36985b2f6e9",
      "service_type" : "nacex_servicio-aereo",
      "rate" : {
        "delivery_date" : "2025-03-12",
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "NACEX SERVICIO AEREO",
        "service_type" : "nacex_servicio-aereo",
        "shipper_account" : {
          "description" : "[nacex] Demo account new",
          "id" : "c5573249c14b4f9eb077e36985b2f6e9",
          "slug" : "nacex"
        },
        "transit_time" : 2
      },
      "service_options" : [ {
        "end_time" : "13:00:00",
        "start_time" : "12:00:00",
        "type" : "pickup"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment_code" : "429673338"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "ESP",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "release-core-sandbox-api.aftership.io",
      "label_base64" : null
    }, {
      "id" : "b71c45088e9e4a59b263b71560310cff",
      "tracking_numbers" : [ "13030532" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T07:53:09.660783Z",
      "updated_at" : "2025-03-10T07:53:11.499262Z",
      "succeed_at" : "2025-03-10T07:53:11.492402446Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP",
            "postal_code" : "08150"
          },
          "ship_to" : {
            "country" : "FRA",
            "postal_code" : "75001"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles(barcelona)",
        "company_name" : "Panda of china",
        "contact_name" : "Javid",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08150",
        "state" : "Mixed",
        "street1" : "ho no 115/flat no b1 street no 7.",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Paris",
        "contact_name" : "Ismael Tejón",
        "country" : "FRA",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "75001",
        "state" : "VIL",
        "street1" : "Calle de prueba, 34.",
        "type" : "residential"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "ESP",
      "ship_to_country" : "FRA",
      "order_id" : null,
      "order_number" : "100001",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/b71c4508-8e9e-4a59-b263-b71560310cff-1741593190608.pdf"
        }
      },
      "box_type" : "nacex_sample",
      "references" : [ "REFERENCE0" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "nacex",
      "carrier_account_id" : "c5573249c14b4f9eb077e36985b2f6e9",
      "service_type" : "nacex_servicio-aereo",
      "rate" : {
        "delivery_date" : "2025-03-12",
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "NACEX SERVICIO AEREO",
        "service_type" : "nacex_servicio-aereo",
        "shipper_account" : {
          "description" : "[nacex] Demo account new",
          "id" : "c5573249c14b4f9eb077e36985b2f6e9",
          "slug" : "nacex"
        },
        "transit_time" : 2
      },
      "service_options" : [ {
        "end_time" : "13:00:00",
        "start_time" : "12:00:00",
        "type" : "pickup"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment_code" : "429672707"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "ESP",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "release-core-sandbox-api.aftership.io",
      "label_base64" : null
    }, {
      "id" : "13eb494ca5b841d6a42038298f969538",
      "tracking_numbers" : [ "13030529" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T07:50:50.697679Z",
      "updated_at" : "2025-03-10T07:50:53.529648Z",
      "succeed_at" : "2025-03-10T07:50:53.521217668Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "ESP",
            "postal_code" : "08150"
          },
          "ship_to" : {
            "country" : "FRA",
            "postal_code" : "75001"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Parets Del Valles(barcelona)",
        "company_name" : "Panda of china",
        "contact_name" : "Javid",
        "country" : "ESP",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "08150",
        "state" : "Mixed",
        "street1" : "ho no 115/flat no b1 street no 7.",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Paris",
        "contact_name" : "Ismael Tejón",
        "country" : "FRA",
        "email" : "test@test.com",
        "phone" : "931234567",
        "postal_code" : "75001",
        "state" : "VIL",
        "street1" : "Calle de prueba, 34.",
        "type" : "residential"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "ESP",
      "ship_to_country" : "FRA",
      "order_id" : null,
      "order_number" : "100001",
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x6",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/13eb494c-a5b8-41d6-a420-38298f969538-1741593052680.pdf"
        }
      },
      "box_type" : "nacex_sample",
      "references" : [ "REFERENCE0" ],
      "carrier_references" : [ ],
      "carrier_account_slug" : "nacex",
      "carrier_account_id" : "c5573249c14b4f9eb077e36985b2f6e9",
      "service_type" : "nacex_servicio-aereo",
      "rate" : {
        "delivery_date" : "2025-03-12",
        "info_message" : "No rate quotes returned from carrier.",
        "service_name" : "NACEX SERVICIO AEREO",
        "service_type" : "nacex_servicio-aereo",
        "shipper_account" : {
          "description" : "[nacex] Demo account new",
          "id" : "c5573249c14b4f9eb077e36985b2f6e9",
          "slug" : "nacex"
        },
        "transit_time" : 2
      },
      "service_options" : [ {
        "end_time" : "13:00:00",
        "start_time" : "12:00:00",
        "type" : "pickup"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment_code" : "429672401"
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1.5
      },
      "chargeable_weight" : null,
      "dimension" : [ {
        "depth" : 40,
        "height" : 30,
        "unit" : "cm",
        "width" : 20
      } ],
      "items" : [ {
        "description" : "Food Bar",
        "hs_code" : "11111111",
        "origin_country" : "ESP",
        "price" : {
          "amount" : 50,
          "currency" : "EUR"
        },
        "quantity" : 2,
        "sku" : "Epic_Food_Bar",
        "weight" : {
          "unit" : "kg",
          "value" : 0.6
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    }, {
      "id" : "924f4fd4c94542e18030ac9e3301a188",
      "tracking_numbers" : [ "7211841346" ],
      "organization_id" : "55c3427cb9e04ea38f61a90705fe1f56",
      "created_at" : "2025-03-10T03:21:19.734639Z",
      "updated_at" : "2025-03-10T03:21:26.268789Z",
      "succeed_at" : "2025-03-10T03:21:26.26024244Z",
      "meta" : {
        "carrier_rate_info" : {
          "ship_from" : {
            "country" : "GBR",
            "postal_code" : "M24 2RW"
          },
          "ship_to" : {
            "country" : "GBR",
            "postal_code" : "HA1 4TR"
          }
        },
        "retention_days" : "365",
        "source" : "api"
      },
      "source" : "api",
      "ship_from" : {
        "city" : "Manchester",
        "company_name" : "CLARKE TELECOM LTD - Stores",
        "contact_name" : "kieron.slack",
        "country" : "GBR",
        "email" : "kieron.slack@clarke-telecom.com",
        "phone" : "+447391016558",
        "postal_code" : "M24 2RW",
        "street1" : "Unit 12-15,Stakehill Industrial Estate,",
        "street3" : "MIddleton ",
        "type" : "business"
      },
      "ship_to" : {
        "city" : "Harrow",
        "company_name" : "Electro Rent UK Ltd",
        "contact_name" : "GBR Har Dispatch Manager",
        "country" : "GBR",
        "email" : "despatch@electrorent.com",
        "eori_number" : "GB541333672000",
        "phone" : "+442084200200",
        "postal_code" : "HA1 4TR",
        "state" : "GB-HRW",
        "street1" : " Unit 1, Waverley Industrial Park\nHailsham Drive",
        "tax_id" : "GB 541 333 672",
        "type" : "business"
      },
      "ship_date" : "2025-03-10",
      "ship_from_country" : "GBR",
      "ship_to_country" : "GBR",
      "order_id" : null,
      "order_number" : null,
      "return_shipment" : false,
      "files" : {
        "label" : {
          "file_type" : "pdf",
          "paper_size" : "4x8",
          "url" : "https://testing-sandbox-download.postmen.io/label/2025-03-10/924f4fd4-c945-42e1-8030-ac9e3301a188-1741576886076.pdf"
        }
      },
      "box_type" : "custom",
      "references" : [ "SON040000609" ],
      "carrier_references" : [ {
        "pickup_reference" : "PRG250310002508"
      } ],
      "carrier_account_slug" : "dhl",
      "carrier_account_id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
      "service_type" : "dhl_domestic_express",
      "rate" : {
        "booking_cut_off" : "2025-03-10T22:30:00+08:00",
        "charge_weight" : {
          "unit" : "kg",
          "value" : 41
        },
        "delivery_date" : "2025-03-12T07:59:00+08:00",
        "detailed_charges" : [ {
          "charge" : {
            "amount" : 960.22,
            "currency" : "EUR"
          },
          "type" : "base"
        }, {
          "charge" : {
            "amount" : 285.67,
            "currency" : "EUR"
          },
          "type" : "fuel_surcharge"
        } ],
        "pickup_deadline" : "2025-03-11T00:30:00+08:00",
        "service_name" : "DHL Express Domestic",
        "service_type" : "dhl_domestic_express",
        "shipper_account" : {
          "description" : "[mydhl] Demo Account [Automate Test]",
          "id" : "0efa2bc603044aa691ac8f2a2d3c8cfd",
          "slug" : "dhl"
        },
        "total_charge" : {
          "amount" : 1245.89,
          "currency" : "EUR"
        },
        "transit_time" : 1
      },
      "service_options" : [ {
        "end_time" : "17:00:00",
        "start_time" : "09:00:00",
        "type" : "pickup"
      }, {
        "enabled" : false,
        "type" : "paperless_invoice"
      } ],
      "err" : null,
      "livemode" : false,
      "status" : "created",
      "extra_info" : {
        "shipment" : {
          "parcels" : [ {
            "box_type" : "custom",
            "description" : "A Default Box",
            "dimension" : {
              "depth" : 70,
              "height" : 53,
              "unit" : "cm",
              "width" : 55
            },
            "items" : [ {
              "barcode" : "",
              "description" : "NOT RESTRICTEDNarda Nardalert S3 Personal Monitor Mainframe & 2271/31 Sensor Module",
              "description_cn" : "",
              "dimension" : {
                "depth" : 0,
                "height" : 0,
                "unit" : "",
                "width" : 0
              },
              "hs_code" : "9030310000",
              "image_urls" : null,
              "item_id" : "",
              "origin_country" : "CRI",
              "price" : {
                "amount" : 37.49,
                "currency" : "GBP"
              },
              "quantity" : 1,
              "return_reason" : "",
              "sku" : "",
              "weight" : {
                "unit" : "kg",
                "value" : 1
              }
            } ],
            "weight" : {
              "unit" : "kg",
              "value" : 1
            }
          } ]
        }
      },
      "custom_fields" : null,
      "weight" : {
        "unit" : "kg",
        "value" : 1
      },
      "chargeable_weight" : {
        "unit" : "kg",
        "value" : 41
      },
      "dimension" : [ {
        "depth" : 70,
        "height" : 53,
        "unit" : "cm",
        "width" : 55
      } ],
      "items" : [ {
        "description" : "NOT RESTRICTEDNarda Nardalert S3 Personal Monitor Mainframe & 2271/31 Sensor Module",
        "hs_code" : "9030310000",
        "origin_country" : "CRI",
        "price" : {
          "amount" : 37.49,
          "currency" : "GBP"
        },
        "quantity" : 1,
        "weight" : {
          "unit" : "kg",
          "value" : 1
        }
      } ],
      "billing" : {
        "paid_by" : "shipper"
      },
      "operator_id" : "postmen-anonymous",
      "label_base64" : null
    } ],
    "pagination" : {
      "total" : 89,
      "next_token" : ""
    }
  }
}`
	// 使用 encoding/json
	// fmt.Println("=== 使用 encoding/json ===")
	// testEncodingJSON(jsonNull, "null case")
	// testEncodingJSON(jsonArray, "array case")

	// // 使用 sonic
	// fmt.Println("\n=== 使用 sonic ===")
	// testSonic(jsonNull, "null case")
	// testSonic(jsonArray, "array case")

	fmt.Println("\n=== 使用 sonic ===")
	testEncodingJSON(jsonStr, "actual encoding case")
	testSonic(jsonStr, "actul sonic case")
}

func testEncodingJSON(jsonStr, testCase string) {
	labelResps := &LabelResps{}
	err := json.Unmarshal([]byte(jsonStr), &labelResps)
	if err != nil {
		fmt.Printf("encoding/json %s 解析失败: %v\n", testCase, err)
	} else {
		fmt.Printf("encoding/json %s 解析成功，detailed_charges: %+v\n", testCase, labelResps.Data.List[0].Rate.DetailedCharges)
	}
}

func testSonic(jsonStr, testCase string) {
	labelResps := &LabelResps{}
	err := sonic.Unmarshal([]byte(jsonStr), labelResps)
	if err != nil {
		fmt.Printf("sonic %s 解析失败: %v\n", testCase, err)
	} else {
		fmt.Printf("sonic %s 解析成功，detailed_charges: %+v\n", testCase, labelResps.Data.List[0].Rate.DetailedCharges)
	}
}

type LabelResp struct {
	Data LabelAPIRespData `json:"data"`
}

type LabelResps struct {
	Data LabelAPIRespsData `json:"data"`
}

type LabelAPIRespsData struct {
	List       []LabelAPIRespData `json:"list"`
	Pagination Pagination         `json:"pagination"`
}

type Pagination struct {
	Total     int64  `json:"total"`
	NextToken string `json:"next_token"`
}

type LabelAPIRespData struct {
	Id              string   `json:"id"`
	TrackingNumbers []string `json:"tracking_numbers"`
	OrganizationId  string   `json:"organization_id"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	SucceedAt       string   `json:"succeed_at"`
	Meta            struct {
		Source string `json:"source"`
	} `json:"meta"`
	Source   string `json:"source"`
	ShipFrom struct {
		City       string `json:"city"`
		Country    string `json:"country"`
		PostalCode string `json:"postal_code"`
		State      string `json:"state"`
	} `json:"ship_from"`
	ShipTo struct {
		City       string `json:"city"`
		Country    string `json:"country"`
		PostalCode string `json:"postal_code"`
		State      string `json:"state"`
	} `json:"ship_to"`
	ShipDate        string      `json:"ship_date"`
	ShipFromCountry string      `json:"ship_from_country"`
	ShipToCountry   string      `json:"ship_to_country"`
	OrderId         interface{} `json:"order_id"`
	OrderNumber     interface{} `json:"order_number"`
	ReturnShipment  bool        `json:"return_shipment"`
	Files           struct {
		Label struct {
			FileType  string `json:"file_type"`
			PaperSize string `json:"paper_size"`
			Url       string `json:"url"`
		} `json:"label"`
	} `json:"files"`
	BoxType            string      `json:"box_type"`
	References         []string    `json:"references"`
	CarrierReferences  interface{} `json:"carrier_references"`
	CarrierAccountSlug string      `json:"carrier_account_slug"`
	CarrierAccountId   string      `json:"carrier_account_id"`
	ServiceType        string      `json:"service_type"`
	Rate               struct {
		DetailedCharges *[]struct {
			Charge struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"charge"`
			Type string `json:"type"`
		} `json:"detailed_charges,omitempty"`
		ServiceName    string `json:"service_name"`
		ServiceType    string `json:"service_type"`
		ShipperAccount struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Slug        string `json:"slug"`
		} `json:"shipper_account"`
		TotalCharge struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		} `json:"total_charge"`
		TransitTime int `json:"transit_time"`
	} `json:"rate"`
	ServiceOptions interface{} `json:"service_options"`
	Err            interface{} `json:"err"`
	Livemode       bool        `json:"livemode"`
	Status         string      `json:"status"`
	ExtraInfo      struct {
		CourierCharge struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		} `json:"courier_charge"`
		Manifest struct {
			TransactionId string `json:"transaction_id"`
		} `json:"manifest"`
	} `json:"extra_info"`
	CustomFields interface{} `json:"custom_fields"`
	LabelBase64  string      `json:"label_base64"`
}
