## API Endpoints

The [OpenApi specification](./api.yml) of the API contains details of available endpoints, and 
bellow you can find some examples of payloads you can use for each one.

You can also import the [Postman workspace example](./go-moneycount-api.postman_collection.json) available.

### List currencies

- [GET] http://localhost:5000/v1/currencies

Response body example:
```json
[
    {
        "code": "BRL",
        "name": "Real",
        "symbol": "R$"
    },
    {
        "code": "USD",
        "name": "Dollar",
        "symbol": "$"
    }
]
```

### Get a specific currency

- [GET] http://localhost:5000/v1/currencies/{code}

Response body example:
```json
{
    "code": "BRL",
    "name": "Real",
    "symbol": "R$"
}
```

### Get a filter for a given currency (necessary for requesting event and event entry maps)

- [GET] http://localhost:5000/v1/currency-filters?currency-code={code}

Response body example:
```json
{
  "currency_code": "BRL",
  "values": [
    {
      "currency_value": "0.10",
      "should_use": true
    },
    {
      "currency_value": "0.50",
      "should_use": true
    },
    {
      "currency_value": "1.00",
      "should_use": true
    },
    {
      "currency_value": "2.00",
      "should_use": true
    },
    {
      "currency_value": "10.00",
      "should_use": true
    },
    {
      "currency_value": "20.00",
      "should_use": true
    },
    {
      "currency_value": "100.00",
      "should_use": true
    },
    {
      "currency_value": "0.05",
      "should_use": true
    },
    {
      "currency_value": "0.25",
      "should_use": true
    },
    {
      "currency_value": "5.00",
      "should_use": true
    },
    {
      "currency_value": "50.00",
      "should_use": true
    }
  ]
}
```

### Request an event entry currency map 

- [POST] http://localhost:5000/v1/event-entries-maps

Request body example:
```json
{
  "event_entry" : {
    "code" : "12345678",
    "description" : "Supplier Payment May 2021",
    "value" : 377.98
  },
  "filter" : {
    "currency_code" : "BRL",
    "values": [
      {
        "currency_value": "1.00",
        "should_use": true
      },
      {
        "currency_value": "2.00",
        "should_use": true
      },
      {
        "currency_value": "20.00",
        "should_use": true
      },
      {
        "currency_value": "50.00",
        "should_use": true
      },
      {
        "currency_value": "0.05",
        "should_use": true
      },
      {
        "currency_value": "0.10",
        "should_use": true
      },
      {
        "currency_value": "0.25",
        "should_use": true
      },
      {
        "currency_value": "0.50",
        "should_use": true
      },
      {
        "currency_value": "5.00",
        "should_use": true
      },
      {
        "currency_value": "10.00",
        "should_use": true
      },
      {
        "currency_value": "100.00",
        "should_use": true
      }
    ]
  }
}
```

Response body example:
```json
{
  "currency_code": "BRL",
  "code": "12345678",
  "description": "12345678",
  "value": 377.98,
  "currency_map": {
    "currencies": [
      {
        "currency_value": "5.00",
        "count": 1
      },
      {
        "currency_value": "0.50",
        "count": 1
      },
      {
        "currency_value": "0.25",
        "count": 1
      },
      {
        "currency_value": "0.10",
        "count": 2
      },
      {
        "currency_value": "100.00",
        "count": 3
      },
      {
        "currency_value": "50.00",
        "count": 1
      },
      {
        "currency_value": "10.00",
        "count": 0
      },
      {
        "currency_value": "0.05",
        "count": 0
      },
      {
        "currency_value": "20.00",
        "count": 1
      },
      {
        "currency_value": "2.00",
        "count": 1
      },
      {
        "currency_value": "1.00",
        "count": 0
      }
    ],
    "remaining_value": 0.03,
    "total_value": 377.98
  }
}
```

### Request an event currency map

[POST] http://localhost:5000/v1/event-maps

Request body example:
```json
{
  "event" : {
    "code" : "SUP2021-05",
    "description" : "Suppliers Payment May 2021",
    "entries" : [
      {
        "code" : "SUP2021-05-0001",
        "description" : "Supplier 0001 May 2021",
        "value" : 129.33
      },
      {
        "code" : "SUP2021-05-0006",
        "description" : "Supplier 0006 May 2021",
        "value" : 336.45
      }
    ]
  },
  "filter" : {
    "currency_code" : "BRL",
    "values": [
      {
        "currency_value": "1.00",
        "should_use": true
      },
      {
        "currency_value": "2.00",
        "should_use": true
      },
      {
        "currency_value": "20.00",
        "should_use": true
      },
      {
        "currency_value": "50.00",
        "should_use": true
      },
      {
        "currency_value": "0.05",
        "should_use": true
      },
      {
        "currency_value": "0.10",
        "should_use": true
      },
      {
        "currency_value": "0.25",
        "should_use": true
      },
      {
        "currency_value": "0.50",
        "should_use": true
      },
      {
        "currency_value": "5.00",
        "should_use": true
      },
      {
        "currency_value": "10.00",
        "should_use": true
      },
      {
        "currency_value": "100.00",
        "should_use": true
      }
    ]
  }
}

```

Response body example:
```json
{
  "currency_code": "BRL",
  "code": "SUP2021-05",
  "description": "Suppliers Payment May 2021",
  "currency_map": {
    "currencies": [
      {
        "currency_value": "20.00",
        "count": 2
      },
      {
        "currency_value": "1.00",
        "count": 1
      },
      {
        "currency_value": "0.50",
        "count": 0
      },
      {
        "currency_value": "0.10",
        "count": 2
      },
      {
        "currency_value": "100.00",
        "count": 4
      },
      {
        "currency_value": "0.25",
        "count": 2
      },
      {
        "currency_value": "10.00",
        "count": 1
      },
      {
        "currency_value": "50.00",
        "count": 0
      },
      {
        "currency_value": "5.00",
        "count": 2
      },
      {
        "currency_value": "2.00",
        "count": 2
      },
      {
        "currency_value": "0.05",
        "count": 1
      }
    ],
    "remaining_value": 0.03,
    "total_value": 465.78
  },
  "entries": [
    {
      "code": "SUP2021-05-0001",
      "description": "Supplier 0001 May 2021",
      "value": 129.33,
      "currency_map": {
        "currencies": [
          {
            "currency_value": "1.00",
            "count": 0
          },
          {
            "currency_value": "0.50",
            "count": 0
          },
          {
            "currency_value": "0.10",
            "count": 0
          },
          {
            "currency_value": "20.00",
            "count": 1
          },
          {
            "currency_value": "10.00",
            "count": 0
          },
          {
            "currency_value": "5.00",
            "count": 1
          },
          {
            "currency_value": "2.00",
            "count": 2
          },
          {
            "currency_value": "0.25",
            "count": 1
          },
          {
            "currency_value": "0.05",
            "count": 1
          },
          {
            "currency_value": "100.00",
            "count": 1
          },
          {
            "currency_value": "50.00",
            "count": 0
          }
        ],
        "remaining_value": 0.03,
        "total_value": 129.33
      }
    },
    {
      "code": "SUP2021-05-0006",
      "description": "Supplier 0006 May 2021",
      "value": 336.45,
      "currency_map": {
        "currencies": [
          {
            "currency_value": "1.00",
            "count": 1
          },
          {
            "currency_value": "0.25",
            "count": 1
          },
          {
            "currency_value": "10.00",
            "count": 1
          },
          {
            "currency_value": "2.00",
            "count": 0
          },
          {
            "currency_value": "20.00",
            "count": 1
          },
          {
            "currency_value": "5.00",
            "count": 1
          },
          {
            "currency_value": "0.50",
            "count": 0
          },
          {
            "currency_value": "0.10",
            "count": 2
          },
          {
            "currency_value": "0.05",
            "count": 0
          },
          {
            "currency_value": "100.00",
            "count": 3
          },
          {
            "currency_value": "50.00",
            "count": 0
          }
        ],
        "remaining_value": 0,
        "total_value": 336.45
      }
    }
  ]
}
```