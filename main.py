from fastapi import FastAPI, Response
from schemas import (
    CheckRequest,
    ListObjectsRequest,
    WriteSchemaRequest,
    WriteTuplesRequest,
    DeleteTuplesRequest,
)
from pyswip import Prolog
from transform import transform_namespaces, transform_tuples, transform_user
from structlog import get_logger


app = FastAPI()
app.state.prolog = Prolog()
app.state.prolog.consult("zanzibar.pl")
app.state.prolog.dynamic("schema/3")
app.state.prolog.dynamic("tuple/4")

app.state.log = get_logger()


@app.post("/write-schema")
def write_schema(req: WriteSchemaRequest):
    for namespace in transform_namespaces(req.namespaces):
        if len(list(app.state.prolog.query(namespace))) == 0:
            app.state.prolog.assertz(namespace)
            app.state.log.info("assertz", namespace=namespace)

    return {}


@app.post("/delete-schema")
def delete_schema():
    app.state.prolog.retractall("schema(_, _, _)")
    return {}


@app.post("/write-tuples")
def write_tuples(req: WriteTuplesRequest):
    for tuple in transform_tuples(req.tuples):
        if len(list(app.state.prolog.query(tuple))) == 0:
            app.state.prolog.assertz(tuple)
            app.state.log.info("assertz", tuple=tuple)

    return {}


@app.post("/delete-tuples")
def delete_tuples(req: DeleteTuplesRequest):
    for tuple in transform_tuples(req.tuples):
        try:
            app.state.prolog.retract(tuple)
        except StopIteration:
            print(f"tuple does not exist: {tuple}")
    return {}


@app.post("/delete-all-tuples")
def delete_tuples():
    app.state.prolog.retractall("tuple(_, _, _, _)")
    return {}


@app.post("/check")
def check(req: CheckRequest):
    query = (
        f"check({req.namespace}, {req.id}, {req.relation}, {transform_user(req.user)})"
    )
    check = list(app.state.prolog.query(query))

    allowed = False
    if len(check) == 1:
        allowed = True

    app.state.log.info("check", query=query, allowed=allowed)

    return {"allowed": allowed}


@app.post("/list-objects")
def check(req: ListObjectsRequest):
    query = f"list({req.namespace}, X, {req.relation}, {transform_user(req.user)})"
    objs = [obj["X"] for obj in app.state.prolog.query(query)]

    app.state.log.info("check", query=query, objs=objs)

    return {"object_ids": objs}
