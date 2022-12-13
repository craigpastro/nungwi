#!/usr/bin/env python3

import httpx


def write_schema():
    r = httpx.post(
        "http://localhost:8080/nungwi.v1alpha.NungwiService/WriteSchema",
        json={
            "configs": [
                {"namespace": "folder", "relation": "viewer", "rewrite": "self"},
                {"namespace": "document", "relation": "parent", "rewrite": "self"},
                {
                    "namespace": "document",
                    "relation": "viewer",
                    "rewrite": "union(self, tupleToUserset(parent, viewer))",
                },
            ]
        },
    )
    assert r.status_code == httpx.codes.OK


def write_tuples(reqs, num_tuples_per_req):
    for i in range(reqs):
        tuples = []
        for j in range(num_tuples_per_req):
            tuples.append(
                {
                    "namespace": "document",
                    "id": f"a{i}b{j}",
                    "relation": "viewer",
                    "user": "abigail",
                }
            )

        r = httpx.post(
            "http://localhost:8080/nungwi.v1alpha.NungwiService/WriteTuples",
            json={"tuples": tuples},
        )
        assert r.status_code == httpx.codes.OK


try:
    write_schema()
except AssertionError:
    print("failed to write_schema")
    quit(1)

try:
    write_tuples(100, 100)
except AssertionError:
    print("failed to write_tuples")
    quit(1)
