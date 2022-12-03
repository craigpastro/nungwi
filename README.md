# Nungwi

## Example Usage

```console
# Run the server
$ make run-development

# Write a schema
$ curl -XPOST 'http://localhost:8000/write-schema' \
--header 'Content-Type: application/json' \
--data-raw '{
    "configs": [
        {
            "namespace": "folder",
            "relation": "viewer",
            "rewrite": "self"
        },
        {
            "namespace": "document",
            "relation": "viewer",
            "rewrite": "union(self, tupleToUserset(parent, viewer))"
        }
    ]
}'
{}

# Write some tuples
$ curl -XPOST 'http://localhost:8000/write-tuples' \
--header 'Content-Type: application/json' \
--data-raw '{
    "tuples": [
        {
            "namespace": "folder",
            "id": "x",
            "relation": "viewer",
            "user": "abigail"
        },
        {
            "namespace": "document",
            "id": "1",
            "relation": "parent",
            "user": "object(folder, x)"
        },
        {
            "namespace": "document",
            "id": "2",
            "relation": "parent",
            "user": "object(folder, x)"
        },
        {
            "namespace": "document",
            "id": "1",
            "relation": "viewer",
            "user": "beatrix"
        }
    ]
}'
{}

# Check
$ curl -XPOST 'http://localhost:8000/check' \
--header 'Content-Type: application/json' \
--data-raw '{
    "namespace": "document",
    "id": "1",
    "relation": "viewer",
    "user": "abigail"
}'
{"allowed": true}

# List objects
$ curl -XPOST 'http://localhost:8000/list-objects' \
--header 'Content-Type: application/json' \
--data-raw '{
    "namespace": "document",
    "relation": "viewer",
    "user": "abigail"
}'
{"object_ids":[1,2]}
```

## Modelling

TODO

## This is a WIP!

I am still learning Python, FastAPI, Prolog... and everything else. Updates coming soon. Reach out if you want to know more.

## What is Nungwi?

Nungwi is a village located at the northern end of the island of Zanzibar.
