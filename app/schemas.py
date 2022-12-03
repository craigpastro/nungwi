from pydantic import BaseModel, Field


class RelationConfig(BaseModel):
    namespace: str
    relation: str
    rewrite: str


class Tuple(BaseModel):
    namespace: str
    id: str
    relation: str
    user: str


class WriteSchemaRequest(BaseModel):
    schema: list[RelationConfig]


class WriteTuplesRequest(BaseModel):
    tuples: list[Tuple]


class DeleteTuplesRequest(BaseModel):
    tuples: list[Tuple]


class CheckRequest(BaseModel):
    namespace: str
    id: str
    relation: str
    user: str


class ListObjectsRequest(BaseModel):
    namespace: str
    relation: str
    user: str
