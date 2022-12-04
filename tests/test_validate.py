import pytest

from app.validate import validate_rewrite, validate_user


def test_validate_self_rewrite():
    assert validate_rewrite("self")


def test_validate_computedUserset_rewrite():
    assert validate_rewrite("computedUserset(viewer)")


def test_validate_tupleToUserset_rewrite():
    assert validate_rewrite("tupleToUserset(parent, viewer)")


def test_validate_union_rewrite():
    assert validate_rewrite("union(self, computedUserset(writer))")


def test_validate_intersection_rewrite():
    assert validate_rewrite("intersection(self, computedUserset(writer))")


def test_validate_exclusion_rewrite():
    assert validate_rewrite(
        "exclusion(computedUserset(writer), computedUserset(banned))"
    )
