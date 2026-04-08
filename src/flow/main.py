from typing import Annotated
import dagger
from dagger import dag, function, object_type, check,field, Doc


@object_type
class Flow:

    op_token: Annotated[
        dagger.Secret | None,
        Doc("1Password service account token. Required by some functions."),
    ]

    @function
    @check
    def test(self):
        if self.op_token is None:
            print("No 1Password token provided")
        else:
            print("1Password token provided")
        return
