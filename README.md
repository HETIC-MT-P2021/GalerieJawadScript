# Back API

## Install

`go get -u github.com/gin-gonic/gin`

## Usage

`go run main.go`

# API Documentation

The API is made with [Goland](https://golang.org/) and a front-end interface also exists to present gallery content. [View more about the front-end.]()

## Summary
- [Get Started](#get-started)
- [API Reference](#api-reference)
  - [Images](#images)
  - [Categories](#categories)
  - [Tags](#tags)

## Get started

The first thing you can do when calling the API is to check if the service is running. To do so you can make a **GET** call on the `/ping` route.

If you get a 200 status code with a small `.json` response everything should be allright.

---

## API Reference

### Images

---

Data Structure:
Field | Type
-|-
id              | integer
name            | string
description     | string
category        | string
uuid_file       | string
date_created    | timestamp
date_updated    | timestamp
tags            | array[string]

JSON representation :

```json
[
    {
        "id": 1,
        "name": "Name of the image",
        "description": "A little description of the image",
        "category": "Category name",
        "uuid_file": "path/uuid_file",
        "date_created": 123546789,
        "date_updated": 123456789,
        "tags":[
            "restaurant",
            "museum"
        ]
    }
]
```

* Get all images

Subject | Information
-|-
URL                 | api/images
Method              | `GET`
Statue Expected     | `200`

* Get one image

Subject | Information
-|-
URL                 | `api/images/:id`
Method              | `GET`
Statue Expected     | `200`

* Create an image

Subject | Information
-|-
URL         | `api/images`
Method      | `POST`
Statue Expected      | `201`

Input parameter:

Field | Details | required
-|-|-
name            | string        | [x]
description     | string        | [x]
category        | string        | [x]
uuid_file       | string        | [x]
tags            | array[string] | []

* Delete image

Subject | Information
-|-
URL                 | `api/images/:id`
Method              | `DELETE`
Statue Expected     | `200`

* Update image

Subject | Information
-|-
URL                 | `api/images/:id`
Method              | `PUT`
Statue Expected     | `200, 201, 204`

Input parameter:

Field | Type | required
-|-|-
name            | string        | [x]
description     | string        | [x]
category        | string        | [x]
uuid_file       | string        | [x]
tags            | array[string] | []

### Categories

---

Data Structure:
Field | Type
-|-
id              | integer
name            | string
date_created    | timestamp

JSON representation :

```json
[
    {
        "id": 1,
        "name": "Category Name",
        "date_created": 123546789
    }
]
```

* Get all categories

Subject | Information
-|-
URL                 | `api/categories`
Method              | `GET`
Statue Expected     | `200`

* Get one category

Subject | Information
-|-
URL                 | `api/categories/:id`
Method              | `GET`
Statue Expected     | `200`

* Create an category

Subject | Information
-|-
URL         | `api/categories`
Method      | `POST`
Statue Expected      | `201`

Input parameter:

Field | Details | required
-|-|-
name            | string        | [x]

* Delete category

Subject | Information
-|-
URL                 | `api/categories/:id`
Method              | `DELETE`
Statue Expected     | `200`

* Update category

Subject | Information
-|-
URL                 | `api/categories/:id`
Method              | `PUT`
Statue Expected     | `200, 201, 204`

Input parameter:

Field | Details | required
-|-|-
name            | string        | [x]

### Tags

---


Data Structure:
Field | Type
-|-
id              | integer
name            | string
date_created    | timestamp

JSON representation :

```json
[
    {
        "id": 1,
        "name": "Tag Name",
        "date_created": 123546789
    }
]
```

* Get all tags

Subject | Information
-|-
URL                 | `api/tags`
Method              | `GET`
Statue Expected     | `200`

* Get one tag

Subject | Information
-|-
URL                 | `api/tags/:id`
Method              | `GET`
Statue Expected     | `200`

* Create an tag

Subject | Information
-|-
URL         | `api/tags`
Method      | `POST`
Statue Expected      | `201`

Input parameter:

Field | Details | required
-|-|-
name            | string        | [x]

* Delete tag

Subject | Information
-|-
URL                 | `api/tags/:id`
Method              | `DELETE`
Statue Expected     | `200`

* Update tag

Subject | Information
-|-
URL                 | `api/tags/:id`
Method              | `PUT`
Statue Expected     | `200, 201, 204`

Input parameter:

Field | Details | required
-|-|-
name            | string        | [x]