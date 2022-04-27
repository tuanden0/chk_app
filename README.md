# CHK APP

A simple app that receive csv file, read the content and add it to db.

## How to run

```bash
docker-compose up -d
```

## Upload CSV API
```bash
curl --location --request POST 'localhost:8080' \
--form 'data.csv"'
```

## List Data API
```bash
curl --location --request GET 'localhost:8080' \
--header 'Content-Type: application/json' \
--data-raw '{
    "limit": 10,
    "page": 1,
    "filters": [
        {
            "key": "unix",
            "method": "=",
            "value": "1644719640000"
        }
    ],
    "sort": {
        "order_by": "id",
        "is_acs": true
    }
}'
```