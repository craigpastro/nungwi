from fastapi.testclient import TestClient
from app.main import app

client = TestClient(app)


def test_check_tupleToUserset():
    resp = client.post(
        "/write-schema",
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
    assert resp.status_code == 200

    resp = client.post(
        "/write-tuples",
        json={
            "tuples": [
                {
                    "namespace": "folder",
                    "id": "x",
                    "relation": "viewer",
                    "user": "abigail",
                },
                {
                    "namespace": "document",
                    "id": "1",
                    "relation": "parent",
                    "user": "object(folder, x)",
                },
            ]
        },
    )
    assert resp.status_code == 200

    resp = client.post(
        "/check",
        json={
            "namespace": "document",
            "id": "1",
            "relation": "viewer",
            "user": "abigail",
        },
    )
    assert resp.status_code == 200
    assert resp.json() == {"allowed": True}
