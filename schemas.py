from __future__ import annotations
from typing import ForwardRef, Optional, Union
from pydantic import BaseModel, Field


class Empty(BaseModel):
    class Config:
        extra = "forbid"


class ComputedUserset(BaseModel):
    relation: str


class TupleToUserset(BaseModel):
    tupleset: str
    computed_userset: str


class Union(BaseModel):
    elem1: Rewrite
    elem2: Rewrite


class Intersection(BaseModel):
    elem1: Rewrite
    elem2: Rewrite


class Exclusion(BaseModel):
    minuend: Rewrite
    subtrahend: Rewrite


class Rewrite(BaseModel):
    this: Optional[Empty] = None
    computed_userset: Optional[ComputedUserset] = None
    tuple_to_userset: Optional[TupleToUserset] = None
    union: Optional[Union] = None
    intersection: Optional[Intersection] = None
    exclusion: Optional[Exclusion] = None


Union.update_forward_refs()
Intersection.update_forward_refs()
Exclusion.update_forward_refs()


class Namespace(BaseModel):
    name: str
    relations: dict[str, Rewrite]


class Object(BaseModel):
    namespace: str
    id: str


class Userset(BaseModel):
    namespace: str
    id: str
    relation: str


class User(BaseModel):
    user_id: Optional[str] = None
    object: Optional[Object] = None
    userset: Optional[Userset] = None


class Tuple(BaseModel):
    namespace: str
    id: str
    relation: str
    user: User


class WriteSchemaRequest(BaseModel):
    namespaces: list[Namespace]


class WriteTuplesRequest(BaseModel):
    tuples: list[Tuple]


class DeleteTuplesRequest(BaseModel):
    tuples: list[Tuple]


class CheckRequest(BaseModel):
    namespace: str
    id: str
    relation: str
    user: User


class ListObjectsRequest(BaseModel):
    namespace: str
    relation: str
    user: User
