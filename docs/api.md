# API design

|API name|method|Path|
|---|---|---|
|Create ToDo|POST|/todo|
|Read ToDo|GET|/todo/{id}|
|Update ToDo|PATCH|/todo/{id}|
|Delete ToDo|DELETE|/todo/{id}|
|List ToDo|GET|/todo|

- [API design](#api-design)
  - [Create ToDo](#create-todo)
    - [HTTP request](#http-request)
    - [Body parameters](#body-parameters)
    - [Response](#response)
      - [code](#code)
      - [body](#body)
  - [Read ToDo](#read-todo)
    - [HTTP request](#http-request-1)
    - [Path parameters](#path-parameters)
    - [Response](#response-1)
      - [code](#code-1)
      - [body](#body-1)
  - [Update Todo](#update-todo)
    - [HTTP request](#http-request-2)
    - [Path parameters](#path-parameters-1)
    - [Body parameters](#body-parameters-1)
    - [Response](#response-2)
      - [code](#code-2)
      - [body](#body-2)
  - [Delete Todo](#delete-todo)
    - [HTTP request](#http-request-3)
    - [Path parameters](#path-parameters-2)
    - [Response](#response-3)
      - [code](#code-3)
      - [body](#body-3)
  - [List Todo](#list-todo)
    - [HTTP request](#http-request-4)
    - [Query parameters](#query-parameters)
    - [Response](#response-4)
      - [code](#code-4)
      - [body](#body-4)

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
