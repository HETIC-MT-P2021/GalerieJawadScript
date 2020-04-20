# Front end documentation

The following routes are available on our front-end.

## Requirement 
Require ELM 0.19.1.
## Installation
To install ELM dependencies, do in root folder, next command line:
```bash
elm make
```

That permit to build your elm-stuff according to your elm.json.

## URL
the project is available at the following url: 
[no-domain-chosen-yet.fr](https://no-domain-chosen-yet.fr)

## Routes

| Url | Description |
|---|---|
|/|Home gallery with the latest added photos|
|/images/`:id`|watch full details of an image|
|/new|Add a new image to the gallery|
|/edit/`:id`|Edit an existing image|
|/categories|View all different categories and acces to their images|
|/categories/`:slug`|View images from a category|
|/tags|View all different tags and acces to their images|
|/tags/`:slug`|View images related to a tag|