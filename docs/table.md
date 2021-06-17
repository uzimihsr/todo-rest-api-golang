# Table design

|table_name|description|
|---|---|
|todo|ToDo table|

## ToDo table

|column|type|option|
|---|---|---|
|id|INT|AUTO_INCREMENT<br>PRIMARY_KEY|
|title|VARCHAR(100)|NOT NULL|
|done|BOOLEAN|NOT NULL<br>DEFAULT false|
|created_at|DATETIME|NOT NULL<br>DEFAULT CURRENT_TIMESTAMP|
|updated_at|DATETIME|NOT NULL<br>DEFAULT CURRENT_TIMESTAMP|