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


def test_validate_user_id():
    assert validate_user("abigail")
    assert not validate_user("1")


def test_validate_user_object():
    assert validate_user("object(document, 1)")
    assert not validate_user("object(1, 1)")


def test_validate_user_userset():
    assert validate_user("userset(document, 1, viewer)")
    assert not validate_user("object(1, 1, viewer)")
    assert not validate_user("object(document, 1, 1)")
