import pytest

from app.schema import RelationConfig, Tuple


def test_relation_config_to_prolog():
    RelationConfig(
        namespace="document", relation="viewer", rewrite="self"
    ).to_prolog() == "config(document,viewer,self)"


def test_tuple_to_prolog():
    Tuple(
        namespace="document", id="foo", relation="viewer", user="abigail"
    ).to_prolog() == "tuple(document,foo,viewer,abigail)"
