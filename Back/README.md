# API Documentation

The API is made with [Goland](https://golang.org/) and a front-end interface also exists to present gallery content. [View more about the front-end.]()

## Summary
- [API Documentation](#api-documentation)
  - [Summary](#summary)
  - [Get started](#get-started)
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

URL | Method | Expected Response | Description
-|-|-|-
images/`:id` | `GET` | `200` | Get an image ressource
images | `GET` | `200` | Get all images ressource
images | `POST` | `201` | Create an image ressource
images/`:id` | `PUT` | `201` | Update an image ressource
images/`:id` | `DELETE` | `200` | Delete an image ressource


JSON representation :

```json
[
    {
        "id": 1, // integer
        "name": "Name of the image", // string
        "description": "A little description of the image", // string
        "category": "Category name", // string
        "uuid_file": "path/uuid_file", // string
        "date_created": 123546789, //datetime
        "date_updated": 123456789, //datetime
        "tags":[
            "restaurant",
            "museum"
        ] // array of string
    }
]
```

### Categories

---

URL | Method | Expected Response | Description
-|-|-|-
categories/`:id` | `GET` | `200` | Get an category ressource
categories | `GET` | `200` | Get all categories ressource
categories | `POST` | `201` | Create an category ressource
categories/`:id` | `PUT` | `201` | Update an category ressource
categories/`:id` | `DELETE` | `200` | Delete an category ressource


JSON representation :

```json
[
    {
        "id": 1, // integer
        "name": "Category Name", // string
        "date_created": 123546789 // datetime
    }
]
```

### Tags

---



URL | Method | Expected Response | Description
-|-|-|-
tags/`:id` | `GET` | `200` | Get an tag ressource
tags | `GET` | `200` | Get all tags ressource
tags | `POST` | `201` | Create an tag ressource
tags/`:id` | `PUT` | `201` | Update an tag ressource
tags/`:id` | `DELETE` | `200` | Delete an tag ressource


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
