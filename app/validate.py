# {
#     "schema": {
#         {
#             "namespace": "document",
#             "relation": "viewer",
#             "rewrite": "self"
#         }
#     }
# }

# {
#     "tuples": {
#         {
#             "namespace": "document",
#             "id": "1",
#             "relation": "viewer",
#             "user": "abigail"
#         }
#     }
# }


def validate_rewrite(rewrite: str) -> bool:
    return True


def validate_user(user: str) -> bool:
    return True
