from fastapi import FastAPI, HTTPException
from pyswip import Prolog
from structlog import get_logger
from schemas import (
    CheckRequest,
    ListObjectsRequest,
    WriteSchemaRequest,
    WriteTuplesRequest,
    DeleteTuplesRequest,
)
from validate import validate_rewrite, validate_user


app = FastAPI()
app.state.prolog = Prolog()
app.state.prolog.consult("zanzibar.pl")
app.state.prolog.dynamic("schema/3")
app.state.prolog.dynamic("tuple/4")

app.state.log = get_logger()


@app.post("/write-schema")
def write_schema(req: WriteSchemaRequest):
    for config in req.schema:
        if not validate_rewrite(config.rewrite):
            raise HTTPException(
                status_code=400,
                detail=f"rewrite not valid in relation config: {config}",
            )

        if len(list(app.state.prolog.query(config))) == 0:
            app.state.prolog.assertz(config)
            app.state.log.info("assertz", config=config)

    return {}


@app.post("/delete-schema")
def delete_schema():
    app.state.prolog.retractall("schema(_, _, _)")
    return {}


@app.post("/write-tuples")
def write_tuples(req: WriteTuplesRequest):
    for tuple in req.tuples:
        if not validate_user(tuple.user):
            raise HTTPException(
                status_code=400, detail=f"user not valid in tuple: {tuple}"
            )

        if len(list(app.state.prolog.query(tuple))) == 0:
            app.state.prolog.assertz(tuple)
            app.state.log.info("assertz", tuple=tuple)

    return {}


@app.post("/delete-tuples")
def delete_tuples(req: DeleteTuplesRequest):
    for tuple in req.tuples:
        try:
            app.state.prolog.retract(tuple)
        except StopIteration:
            app.state.log.info("try to delete tuple which does not exist", tuple=tuple)

    return {}


@app.post("/delete-all-tuples")
def delete_tuples():
    app.state.prolog.retractall("tuple(_, _, _, _)")
    return {}


@app.post("/check")
def check(req: CheckRequest):
    if not validate_user(req.user):
        raise HTTPException(status_code=400, detail=f"user not valid: {user}")

    query = f"check({req.namespace}, {req.id}, {req.relation}, {req.user})"
    check = list(app.state.prolog.query(query))

    allowed = False
    if len(check) == 1:
        allowed = True

    app.state.log.info("check", query=query, allowed=allowed)

    return {"allowed": allowed}


@app.post("/list-objects")
def check(req: ListObjectsRequest):
    if not validate_user(req.user):
        raise HTTPException(status_code=400, detail=f"user not valid: {user}")

    query = f"list({req.namespace}, X, {req.relation}, {req.user})"
    objs = [obj["X"] for obj in app.state.prolog.query(query)]

    app.state.log.info("check", query=query, objs=objs)

    return {"object_ids": objs}
