from schemas import Namespace, Rewrite, Tuple, User


def transform_namespaces(namespaces: list[Namespace]) -> list[str]:
    res = []
    for namespace in namespaces:
        res.extend(transform_namespace(namespace))

    return res


def transform_namespace(namespace: Namespace) -> list[str]:
    res = []
    for relation, rewrite in namespace.relations.items():
        res.append(
            f"schema({namespace.name}, {relation}, {transform_rewrite(rewrite)})"
        )

    return res


def transform_rewrite(rewrite: Rewrite) -> str:
    if rewrite.this:
        return "self"
    elif rewrite.computed_userset:
        return f"computedUserset({rewrite.computed_userset.relation})"
    elif rewrite.tuple_to_userset:
        return f"tupleToUserset({rewrite.tuple_to_userset.tupleset}, {rewrite.tuple_to_userset.computed_userset})"
    elif rewrite.union:
        return f"union({transform_rewrite(rewrite.union.elem1)}, {transform_rewrite(rewrite.union.elem2)})"
    elif rewrite.intersection:
        return f"intersection({transform_rewrite(rewrite.intersection.elem1)}, {transform_rewrite(rewrite.intersection.elem2)})"
    else:  # exclusion
        return f"exclusion({transform_rewrite(rewrite.exclusion.minuend)}, {transform_rewrite(rewrite.exclusion.subtrahend)})"


def transform_tuples(tuples: list[Tuple]) -> list[str]:
    res = []
    for tuple in tuples:
        res.append(transform_tuple(tuple))

    return res


def transform_tuple(tuple: Tuple) -> str:
    return f"tuple({tuple.namespace}, {tuple.id}, {tuple.relation}, {transform_user(tuple.user)})"


def transform_user(user: User) -> str:
    if user.user_id:
        return user.user_id
    elif user.object:
        return f"object({user.object.namespace}, {user.object.id})"
    else:
        return f"userset({user.userset.namespace}, {user.userset.id}, {user.userset.relation})"
