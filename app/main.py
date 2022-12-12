from fastapi import FastAPI, HTTPException
from pyswip import Prolog
from structlog import get_logger
from threading import Lock

from .schema import (
    CheckRequest,
    ListObjectsRequest,
    WriteSchemaRequest,
    WriteTuplesRequest,
    DeleteTuplesRequest,
)
from .validate import validate_rewrite, validate_user


app = FastAPI()

prolog = Prolog()
prolog.consult("app/zanzibar.pl")
prolog.dynamic("config/3")
prolog.dynamic("tuple/4")

log = get_logger()
mutex = Lock()


@app.post("/write-schema")
def write_schema(req: WriteSchemaRequest):
    for config in req.configs:
        if not validate_rewrite(config.rewrite):
            raise HTTPException(
                status_code=400,
                detail=f"rewrite not valid in relation config: {config}",
            )

        s = config.to_prolog()
        log.info("assertz", config=s)

        mutex.acquire()
        try:
            if len(list(prolog.query(s))) == 0:
                log.info("assertz", config=s)
                prolog.assertz(s)
        finally:
            mutex.release()

    return {}


@app.post("/delete-schema")
def delete_schema():
    mutex.acquire()
    try:
        prolog.retractall("config(_, _, _)")
    finally:
        mutex.release()

    return {}


@app.post("/write-tuples")
def write_tuples(req: WriteTuplesRequest):
    for tuple in req.tuples:
        if not validate_user(tuple.user):
            raise HTTPException(
                status_code=400, detail=f"user not valid in tuple: {tuple}"
            )

        t = tuple.to_prolog()

        mutex.acquire()
        try:
            if len(list(prolog.query(t))) == 0:
                prolog.assertz(t)
                log.info("assertz", tuple=t)
        finally:
            mutex.release()

    return {}


@app.post("/delete-tuples")
def delete_tuples(req: DeleteTuplesRequest):
    for tuple in req.tuples:
        mutex.acquire()
        try:
            prolog.retract(tuple.to_prolog())
        except StopIteration:
            log.info("tried to delete tuple which does not exist", tuple=tuple)
        finally:
            mutex.release()

    return {}


@app.post("/delete-all-tuples")
def delete_tuples():
    mutex.acquire()
    try:
        prolog.retractall("tuple(_, _, _, _)")
    finally:
        mutex.release()

    return {}


@app.post("/check")
def check(req: CheckRequest):
    if not validate_user(req.user):
        raise HTTPException(status_code=400, detail=f"user not valid: {user}")

    query = f"check({req.namespace}, {req.id}, {req.relation}, {req.user})"

    mutex.acquire()
    try:
        check = list(prolog.query(query))
    finally:
        mutex.release()

    allowed = False
    if len(check) == 1:
        allowed = True

    log.info("check", query=query, allowed=allowed)

    return {"allowed": allowed}


@app.post("/list-objects")
def check(req: ListObjectsRequest):
    if not validate_user(req.user):
        raise HTTPException(status_code=400, detail=f"user not valid: {user}")

    query = f"list({req.namespace}, X, {req.relation}, {req.user})"

    mutex.acquire()
    try:
        res = prolog.query(query)
    finally:
        mutex.release()

    objs = [obj["X"] for obj in res]

    log.info("check", query=query, objs=objs)

    return {"object_ids": objs}


@app.on_event("startup")
def shutdown_event():
    log.info("🙌 starting nungwi")


@app.on_event("shutdown")
def shutdown_event():
    log.info("shutting down nungwi 👋")
