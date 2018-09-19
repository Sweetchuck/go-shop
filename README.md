# Dummy shop written in Go


## Environment variables


### SHOP_DATA_SOURCE_URL

Example values:
* `mysql://me:my-pass@tcp(127.0.0.1:3306)/my_db_name?charset=utf8&parseTime=True&loc=Local`
* `sqlite3://somewhere.sqlite`
* `memory://this-part-is-ignored`


### SHOP_ADDRESS

Example value: `:8080`


### SHOP_API_PATH_PREFIX

Example value: `/api/v1`


## Paths

* GET `/api/v1/status`
* GET `/api/v1/person`
* POST `/api/v1/person`
* GET `/api/v1/person/:id`
* PATCH `/api/v1/person/:id`
* DELETE `/api/v1/person/:id`
* GET `/api/v1/shoe`
* POST `/api/v1/shoe`
* GET `/api/v1/shoe/:id`
* PATCH `/api/v1/shoe/:id`
* DELETE `/api/v1/shoe/:id`


## Person

```json
{
  "Name": "Foo",
  "Pass": "bar"
}
```


## Shoe

```json
{
  "Brand": "Tisza",
  "Type": "X"
}
```
