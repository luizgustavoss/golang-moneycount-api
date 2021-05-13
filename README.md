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

Once you have run the project, you can access the available endpoints:

http://localhost:5000/currencies

http://localhost:5000/currencies/{code}