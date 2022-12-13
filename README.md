# Nungwi

Nungwi is a PoC authorization service inspired by [Google Zanzibar](https://research.google/pubs/pub48190/), written in [Ichiban Prolog](https://github.com/ichiban/prolog). Nungwi is also is a village located at the northern end of the island of Zanzibar.

This is a PoC WIP. Don't use in production. I welcome all contributions!

## Example Usage

```console
# Run the server
$ make run-development

# Write a schema
$ curl -XPOST 'http://localhost:8080/nungwi.v1alpha.NungwiService/WriteSchema' \
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
            "relation": "parent",
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
$ curl -XPOST 'http://localhost:8080/nungwi.v1alpha.NungwiService/WriteTuples' \
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
$ curl -XPOST 'http://localhost:8080/nungwi.v1alpha.NungwiService/Check' \
--header 'Content-Type: application/json' \
--data-raw '{
    "namespace": "document",
    "id": "1",
    "relation": "viewer",
    "user": "abigail"
}'
{"allowed": true}

# List objects
$ curl -XPOST 'http://localhost:8080/nungwi.v1alpha.NungwiService/ListObjects' \
--header 'Content-Type: application/json' \
--data-raw '{
    "namespace": "document",
    "relation": "viewer",
    "user": "abigail"
}'
{"ids": [1,2]}
```

## Modelling

### Schema

A rewrite can be any of the following:
- `self`.
- `computedUserset(relation: str)`
- `tupleToUserset(tupleset: str, computedUserset: str)`
- `union(p1: rewrite, p2: rewrite)`
- `intersection(p1: rewrite, p2: rewrite)`
- `exclusion(minuend: rewrite, subtrahend: rewrite)`

The `str` type indicates that the argument must be a string, and the `rewrite` type indicates that the argument is a rewrite itself.

A relation config consists of a namespace, relation, and rewrite. These are written as:
```python
config(namespace, relation, rewrite)
```

A schema consists of one or more relation configs.

### Tuples

A user can be any of the following:
- `userId: str`
- `object(namespace: str, id: str)`
- `userset(namespace: str, id: str, relation: str)`

A tuple consists of a namespace, id, relation, and user. These are written as:
```python
tuple(namespace, id, relation, user)
```

## Acknowledgements

- Thank you to @ichiban for helping me with [Ichiban Prolog](https://github.com/ichiban/prolog).
