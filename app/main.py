from fastapi import FastAPI, HTTPException
from pyswip import Prolog
from structlog import get_logger
from .schema import (
    CheckRequest,
    ListObjectsRequest,
    WriteSchemaRequest,
    WriteTuplesRequest,
    DeleteTuplesRequest,
)
from .validate import validate_rewrite, validate_user


app = FastAPI()

p = Prolog()
p.consult("app/zanzibar.pl")
p.dynamic("config/3")
p.dynamic("tuple/4")
app.state.prolog = p

app.state.log = get_logger()


@app.post("/write-schema")
def write_schema(req: WriteSchemaRequest):
    for config in req.configs:
        if not validate_rewrite(config.rewrite):
            raise HTTPException(
                status_code=400,
                detail=f"rewrite not valid in relation config: {config}",
            )

        s = config.to_prolog()
        app.state.log.info("assertz", config=s)

        if len(list(app.state.prolog.query(s))) == 0:
            app.state.log.info("assertz", config=s)
            app.state.prolog.assertz(s)

    return {}


@app.post("/delete-schema")
def delete_schema():
    app.state.prolog.retractall("config(_, _, _)")
    return {}


@app.post("/write-tuples")
def write_tuples(req: WriteTuplesRequest):
    for tuple in req.tuples:
        if not validate_user(tuple.user):
            raise HTTPException(
                status_code=400, detail=f"user not valid in tuple: {tuple}"
            )

        t = tuple.to_prolog()
        if len(list(app.state.prolog.query(t))) == 0:
            app.state.prolog.assertz(t)
            app.state.log.info("assertz", tuple=t)

    return {}


@app.post("/delete-tuples")
def delete_tuples(req: DeleteTuplesRequest):
    for tuple in req.tuples:
        try:
            app.state.prolog.retract(tuple.to_prolog())
        except StopIteration:
            app.state.log.info(
                "tried to delete tuple which does not exist", tuple=tuple
            )

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
