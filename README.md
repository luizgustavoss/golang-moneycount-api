# MoneyCount API


A simple API, built for learning purposes, that can be used to define the amount of notes and coins (cash) you need for one or more monetary values.



## Important Concepts

Although this is a simple and straight API some explanations are important.

This API is useful for those who need to deal with lots of cash for paying bills. For example, suppose somebody, or some company, needs to pay its employees some amount of money in cash, with values for each employee being different, for some reason. In most cases it could be paied direcly in bank accounts or even payment in check, but if for some reason it needs to be in cash, withdraw the money in the bank in the exact quantity of notes and coins to sum each payment can be a challenge.

Money Count helps in this: you can use the API to calculate the amount of notes and coins for a monetary event (the whole employee's payment), and for each entry.



## Ubiquitous Language

To better undestant the functionality, bellow there are some common terms that belong to the ubiquitous language of the context of the project:

**Currency** - represents a currency used in one or more countries, for example Dollar (USD), Real (BRL). At the moment the API only consider these two currencies, but more currencies will be available soon.

**Currency Filter** - if for some reason you can't rely on one or more notes or coins for your payments, whatever reason, you have a chance to say witch notes and coins will be considered. A *Currency Filter* is specific for a *Currency*.

**Event Entry** - a single payment you have to do.

**Event** - represents an event in which you have a group of payments (*Event Entry*) to do. You decide what an *Event* means for you, it could be all the payments to do in one day, or the payments of a week. The important thing is that you can group the values to calculate the amount of notes and coins to be used.



## Pre Requisites

To run the project you must have **Golang** installed and configured. Refer to the [installation instructions](https://golang.org/doc/install) for details. 

You should have at least version 1.11.x, but prefer the latest version available on the [download page](https://golang.org/dl/). 


## Cloning and Running

You can clone and run the project through the commands bellow:

```
$ git clone git@github.com:luizgustavoss/golang-moneycount-api.git && cd golang-moneycount-api

$ go run main.go
```


## Endpoints

Once you have run the project, you can access the available endpoints:

 - [GET] http://localhost:5000/currencies

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

 - [GET] http://localhost:5000/currencies/{code}

Response body example:
```json
{
    "code": "BRL",
    "name": "Real",
    "symbol": "R$"
}
```


 - [GET] http://localhost:5000/currency-filters?currency-code={code}

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


 - [POST] http://localhost:5000/event-entries-maps

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
    "map": {
      "0.05": 0,
      "0.10": 2,
      "0.25": 1,
      "0.50": 1,
      "1.00": 0,
      "10.00": 0,
      "100.00": 3,
      "2.00": 1,
      "20.00": 1,
      "5.00": 1,
      "50.00": 1
    },
    "remaining_value": 0.03,
    "total_value": 377.98
  }
}
```

[POST] http://localhost:5000/event-maps

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
    "map": {
      "0.05": 1,
      "0.10": 2,
      "0.25": 2,
      "0.50": 0,
      "1.00": 1,
      "10.00": 1,
      "100.00": 4,
      "2.00": 2,
      "20.00": 2,
      "5.00": 2,
      "50.00": 0
    },
    "remaining_value": 0.03,
    "total_value": 465.78
  },
  "entries": [
    {
      "code": "SUP2021-05-0001",
      "description": "Supplier 0001 May 2021",
      "value": 129.33,
      "currency_map": {
        "map": {
          "0.05": 1,
          "0.10": 0,
          "0.25": 1,
          "0.50": 0,
          "1.00": 0,
          "10.00": 0,
          "100.00": 1,
          "2.00": 2,
          "20.00": 1,
          "5.00": 1,
          "50.00": 0
        },
        "remaining_value": 0.03,
        "total_value": 129.33
      }
    },
    {
      "code": "SUP2021-05-0006",
      "description": "Supplier 0006 May 2021",
      "value": 336.45,
      "currency_map": {
        "map": {
          "0.05": 0,
          "0.10": 2,
          "0.25": 1,
          "0.50": 0,
          "1.00": 1,
          "10.00": 1,
          "100.00": 3,
          "2.00": 0,
          "20.00": 1,
          "5.00": 1,
          "50.00": 0
        },
        "remaining_value": 0,
        "total_value": 336.45
      }
    }
  ]
}
```