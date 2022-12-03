import re

pattern = "^[a-z][a-zA-Z0-9_-]*$"


def validate_rewrite(rewrite: str) -> bool:
    rewrite = re.sub("\s+", "", rewrite)

    if rewrite == "self":
        return True
    elif rewrite.startswith("computedUserset(") and rewrite.endswith(")"):
        return re.match(pattern, rewrite[16:-1])
    elif rewrite.startswith("tupleToUserset(") and rewrite.endswith(")"):
        tupleset, computed_userset = rewrite[15:-1].split(",")
        return re.match(pattern, tupleset) and re.match(pattern, computed_userset)
    elif rewrite.startswith("union(") and rewrite.endswith(")"):
        p1, p2 = rewrite[6:-1].split(",", 1)
        return validate_rewrite(p1) and validate_rewrite(p2)
    elif rewrite.startswith("intersection(") and rewrite.endswith(")"):
        p1, p2 = rewrite[13:-1].split(",", 1)
        return validate_rewrite(p1) and validate_rewrite(p2)
    elif rewrite.startswith("exclusion(") and rewrite.endswith(")"):
        p1, p2 = rewrite[10:-1].split(",", 1)
        return validate_rewrite(p1) and validate_rewrite(p2)

    return False


def validate_user(user: str) -> bool:
    user = re.sub("\s+", "", user)

    if user.startswith("object(") and user.endswith(")"):
        try:
            namespace, id = user[7:-1].split(",")
        except ValueError:
            return False
        return re.match(pattern, namespace) and re.match(pattern, id)
    elif user.startswith("userset(") and user.endswith(")"):
        try:
            namespace, id, relation = user[8:-1].split(",")
        except ValueError:
            return False
        return (
            re.match(pattern, namespace)
            and re.match(pattern, id)
            and re.match(pattern, relation)
        )
    else:
        return re.match(pattern, user)
