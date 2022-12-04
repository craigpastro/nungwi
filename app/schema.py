from pydantic import BaseModel, Field


class RelationConfig(BaseModel):
    namespace: str
    relation: str
    rewrite: str

    def to_prolog(self) -> str:
        return f"config({self.namespace},{self.relation},{self.rewrite})"


class Tuple(BaseModel):
    namespace: str
    id: str
    relation: str
    user: str

    def to_prolog(self) -> str:
        return f"tuple({self.namespace},{self.id},{self.relation},{self.user})"


class WriteSchemaRequest(BaseModel):
    configs: list[RelationConfig]


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
