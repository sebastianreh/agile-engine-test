# Agile-engine-test

To run the project please run the following command

`./agile-engine-test`

-------
#### Endpoints

Retrieves the history of transactions

`/api/user/history`


-------

Stores a transaction

`/api/user/transaction/commit`

JSON request model:

```
{
    "type": string,
    "amount": float
}
```

Type could be either `credit` or `debit`.

-------

Get the transaction object by ID

`/api/user/history/:transactionID`

TransactionID should be of uuid type such as:

`59029dec-d179-44b2-b1db-198d93d806fc`
