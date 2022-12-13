:- dynamic(config/3).
:- dynamic(tuple/4).

% checkWR is check with rewrite
checkWR(Namespace, Id, Rel, User, self) :- tuple(Namespace, Id, Rel, User).

checkWR(Namespace, Id, _, User, computedUserset(Rel0)) :- tuple(Namespace, Id, Rel0, User).

checkWR(Namespace, Id, _, User, tupleToUserset(S, T)) :-
    tuple(Namespace, Id, S, object(Namespace0, Id0)),
    config(Namespace0, T, Rewrite),
    checkWR(Namespace0, Id0, T, User, Rewrite).

checkWR(Namespace, Id, _, User, tupleToUserset(S, T)) :-
    tuple(Namespace, Id, S, userset(Namespace0, Id0, T)),
    config(Namespace0, T, Rewrite),
    checkWR(Namespace0, Id0, T, User, Rewrite).

checkWR(Namespace, Id, Rel, User, union(S, _)) :- checkWR(Namespace, Id, Rel, User, S).
checkWR(Namespace, Id, Rel, User, union(_, T)) :- checkWR(Namespace, Id, Rel, User, T).

checkWR(Namespace, Id, Rel, User, intersection(S, T)) :-
    checkWR(Namespace, Id, Rel, User, S),
    checkWR(Namespace, Id, Rel, User, T).

checkWR(Namespace, Id, Rel, User, exclusion(S, T)) :-
    checkWR(Namespace, Id, Rel, User, S),
    \+ checkWR(Namespace, Id, Rel, User, T).

% check add the cut at the end so it just finds the answer and won't backtrack.
check(Namespace, Id, Rel, User) :-
    config(Namespace, Rel, Rewrite),
    checkWR(Namespace, Id, Rel, User, Rewrite), !.

list(Namespace, Id, Rel, User) :-
    config(Namespace, Rel, Rewrite),
    checkWR(Namespace, Id, Rel, User, Rewrite).
