import httpx
import uuid


def write_schema():
    r = httpx.post(
        "http://localhost:8000/write-schema",
        json={
            "configs": [
                {"namespace": "folder", "relation": "viewer", "rewrite": "self"},
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
                    "id": f"{i}-{j}",
                    "relation": "viewer",
                    "user": "abigail",
                }
            )

        r = httpx.post("http://localhost:8000/write-tuples", json={"tuples": tuples})
        assert r.status_code == httpx.codes.OK


write_schema()
write_tuples(100, 100)
