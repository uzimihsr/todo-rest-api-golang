# API design

|API name|method|Path|
|---|---|---|
|Create ToDo|POST|/todo|
|Read ToDo|GET|/todo/{id}|
|Update ToDo|PATCH|/todo/{id}|
|Delete ToDo|DELETE|/todo/{id}|
|List ToDo|GET|/todo|

## Create ToDo

Create ToDo.  

### HTTP request

```
POST /todo
```

### Body parameters

```json
{
    "title": "Buy a new pencil",    // string: required
    "done": false                   // boolean
}
```

### Response

#### code

|code|description|
|---|---|
|200|OK|

#### body

```json
{
    "id": 123,
    "title": "Buy a new pencil",
    "done": false,
    "created_at": "2021-06-15T00:35:07Z",
    "updated_at": "2021-06-15T00:35:07Z"
}
```

## Read ToDo

Read the specified ToDo.  

### HTTP request

```
GET /todo/{id}
```

### Path parameters

|parameter|description|
|---|---|
|id|`number`<br>`required`<br>ID number of the ToDo|

### Response

#### code

|code|description|
|---|---|
|200|OK|

#### body

```json
{
    "id": 123,
    "title": "Buy a new pencil",
    "done": false,
    "created_at": "2021-06-15T00:35:07Z",
    "updated_at": "2021-06-15T00:35:07Z"
}
```

## Update Todo

partially update the specified ToDo.  

### HTTP request

```
PATCH /todo/{id}
```

### Path parameters

|parameter|description|
|---|---|
|id|`number`<br>`required`<br>ID number of the ToDo|

### Body parameters

```json
{
    "title": "Buy a new pencil",    // string: If not specified, the original value is retained.
    "done": true                    // boolean: If not specified, the original value is retained.
}
```

### Response

#### code

|code|description|
|---|---|
|200|OK|

#### body

```json
{
    "id": 123,
    "title": "Buy a new pencil",
    "done": true,
    "created_at": "2021-06-15T00:35:07Z",
    "updated_at": "2021-06-15T00:40:10Z"
}
```

## Delete Todo

delete the specified ToDo  

### HTTP request

```
DELETE /todo/{id}
```

### Path parameters

|parameter|description|
|---|---|
|id|`number`<br>`required`<br>ID number of the ToDo|

### Response

#### code

|code|description|
|---|---|
|200|OK|

#### body

```json
{
    "id": 123,
    "title": "Buy a new pencil",
    "done": true,
    "created_at": "2021-06-15T00:35:07Z",
    "updated_at": "2021-06-15T00:40:10Z"
}
```

## List Todo

list ToDo

### HTTP request

```
GET /todo?done={done}
```

### Query parameters

|parameter|default|description|
|---|---|---|
|done|null|`boolean`<br>filter by true or false|

### Response

#### code

|code|description|
|---|---|
|200|OK|

#### body

```json
[
    {
        "id": 123,
        "title": "Buy a new pencil",
        "done": true,
        "created_at": "2021-06-15T00:35:07Z",
        "updated_at": "2021-06-15T00:40:10Z"
    },
    {
        "id": 456,
        "title": "Go to the cinema to see a movie",
        "done": false,
        "created_at": "2021-06-15T00:35:07Z",
        "updated_at": "2021-06-15T00:40:10Z"
    }
]
```


