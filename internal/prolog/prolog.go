package prolog

import (
	"context"
	"fmt"

	pb "github.com/craigpastro/nungwi/internal/gen/nungwi/v1alpha"
	"github.com/ichiban/prolog"
)

type Prolog struct {
	interpreter *prolog.Interpreter
}

const zanzibar string = `
:- dynamic(config/3).
:- dynamic(tuple/4).

checkWR(Namespace, Id, Rel, User, self) :- tuple(Namespace, Id, Rel, User).

checkWR(Namespace, Id, _, User, computedUserset(Rel0)) :-
    config(Namespace, Rel0, Rewrite),
    checkWR(Namespace, Id, Rel0, User, Rewrite).

checkWR(Namespace, Id, _, User, tupleToUserset(Tupleset, ComputedUserset)) :-
    tuple(Namespace, Id, Tupleset, object(Namespace0, Id0)),
    config(Namespace0, ComputedUserset, Rewrite),
    checkWR(Namespace0, Id0, ComputedUserset, User, Rewrite).

checkWR(Namespace, Id, _, User, tupleToUserset(Tupleset, ComputedUserset)) :-
    tuple(Namespace, Id, Tupleset, userset(Namespace0, Id0, ComputedUserset)),
    config(Namespace0, ComputedUserset, Rewrite),
    checkWR(Namespace0, Id0, ComputedUserset, User, Rewrite).

checkWR(Namespace, Id, Rel, User, union(L, _)) :- checkWR(Namespace, Id, Rel, User, L).
checkWR(Namespace, Id, Rel, User, union(_, R)) :- checkWR(Namespace, Id, Rel, User, R).

checkWR(Namespace, Id, Rel, User, intersection(L, R)) :-
    checkWR(Namespace, Id, Rel, User, L),
    checkWR(Namespace, Id, Rel, User, R).

checkWR(Namespace, Id, Rel, User, exclusion(M, S)) :-
    checkWR(Namespace, Id, Rel, User, M),
    \+ checkWR(Namespace, Id, Rel, User, S).

check(Namespace, Id, Rel, User) :-
    config(Namespace, Rel, Rewrite),
    checkWR(Namespace, Id, Rel, User, Rewrite), !.

list(Namespace, Id, Rel, User) :-
    config(Namespace, Rel, Rewrite),
    checkWR(Namespace, Id, Rel, User, Rewrite).
`

func MustNew() *Prolog {
	interpreter := prolog.New(nil, nil)
	if err := interpreter.Exec(zanzibar); err != nil {
		panic(err)
	}

	return &Prolog{
		interpreter: interpreter,
	}
}

func (p *Prolog) WriteSchema(ctx context.Context, configs []*pb.RelationConfig) error {
	for _, config := range configs {
		f := fmt.Sprintf("config(%s, %s, %s)", config.Namespace, config.Relation, config.Rewrite)
		ff := fmt.Sprintf(":- assertz(%s).", f)
		if err := p.interpreter.ExecContext(ctx, ff); err != nil {
			return fmt.Errorf("prolog exec error: %w", err)
		}
	}

	return nil
}

func (p *Prolog) GetSchema(ctx context.Context) ([]*pb.RelationConfig, error) {
	sols, err := p.interpreter.QueryContext(ctx, "config(Namespace, Relation, Rewrite).")
	if err != nil {
		return nil, fmt.Errorf("prolog query error: %w", err)
	}
	defer sols.Close()

	var configs []*pb.RelationConfig
	for sols.Next() {
		var config struct {
			Namespace, Relation, Rewrite prolog.TermString
		}
		if err := sols.Scan(&config); err != nil {
			return nil, fmt.Errorf("prolog scan error: %w", err)
		}
		configs = append(configs, &pb.RelationConfig{
			Namespace: string(config.Namespace),
			Relation:  string(config.Relation),
			Rewrite:   string(config.Rewrite),
		})
	}

	if err := sols.Err(); err != nil {
		return nil, fmt.Errorf("prolog error: %w", err)
	}

	return configs, nil
}

func (p *Prolog) DeleteSchema(ctx context.Context) error {
	if err := p.interpreter.ExecContext(ctx, ":- retractall(config(_, _, _))."); err != nil {
		return fmt.Errorf("prolog exec error: %w", err)
	}

	return nil
}

func (p *Prolog) WriteTuples(ctx context.Context, tuples []*pb.Tuple) error {
	for _, tuple := range tuples {
		f := fmt.Sprintf("tuple(%s, %s, %s, %s)", tuple.Namespace, tuple.Id, tuple.Relation, tuple.User)

		sols, err := p.interpreter.QueryContext(ctx, fmt.Sprintf("%s.", f))
		if err != nil {
			return fmt.Errorf("prolog query error: %w", err)
		}

		if sols.Next() {
			// Don't insert the same tuple more than once.
			continue
		}

		ff := fmt.Sprintf(":- assertz(%s).", f)
		if err := p.interpreter.ExecContext(ctx, ff); err != nil {
			return fmt.Errorf("prolog exec error: %w", err)
		}
	}

	return nil
}

func (p *Prolog) GetTuples(ctx context.Context) ([]*pb.Tuple, error) {
	sols, err := p.interpreter.QueryContext(ctx, "tuple(Namespace, ID, Relation, User).")
	if err != nil {
		return nil, fmt.Errorf("prolog query error: %w", err)
	}
	defer sols.Close()

	var tuples []*pb.Tuple
	for sols.Next() {
		var tuple struct {
			Namespace, ID, Relation, User prolog.TermString
		}
		if err := sols.Scan(&tuple); err != nil {
			return nil, fmt.Errorf("prolog scan error: %w", err)
		}
		tuples = append(tuples, &pb.Tuple{
			Namespace: string(tuple.Namespace),
			Id:        string(tuple.ID),
			Relation:  string(tuple.Relation),
			User:      string(tuple.User),
		})
	}

	if err := sols.Err(); err != nil {
		return nil, fmt.Errorf("prolog sols error: %w", err)
	}

	return tuples, nil
}

func (p *Prolog) DeleteTuples(ctx context.Context, tuples []*pb.Tuple) error {
	for _, tuple := range tuples {
		f := fmt.Sprintf("tuple(%s, %s, %s, %s)", tuple.Namespace, tuple.Id, tuple.Relation, tuple.User)
		ff := fmt.Sprintf(":- retract(%s).", f)
		if err := p.interpreter.ExecContext(ctx, ff); err != nil {
			return fmt.Errorf("prolog exec error: %w", err)
		}
	}

	return nil
}

func (p *Prolog) Check(ctx context.Context, tuple *pb.Tuple) (bool, error) {
	f := fmt.Sprintf("check(%s, %s, %s, %s).", tuple.Namespace, tuple.Id, tuple.Relation, tuple.User)
	sols, err := p.interpreter.QueryContext(ctx, f)
	if err != nil {
		return false, fmt.Errorf("prolog query error: %w", err)
	}
	defer sols.Close()

	allowed := false
	if sols.Next() {
		allowed = true
	}

	if err := sols.Err(); err != nil {
		return false, fmt.Errorf("prolog sols error: %w", err)
	}

	return allowed, nil
}

func (p *Prolog) ListObjects(ctx context.Context, namespace, relation, user string) ([]string, error) {
	f := fmt.Sprintf("list(%s, ID, %s, %s).", namespace, relation, user)
	sols, err := p.interpreter.QueryContext(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("prolog query error: %w", err)
	}
	defer sols.Close()

	var ids []string
	for sols.Next() {
		var s struct {
			ID prolog.TermString
		}
		if err := sols.Scan(&s); err != nil {
			return nil, fmt.Errorf("prolog scan error: %w", err)
		}
		ids = append(ids, string(s.ID))
	}

	if err := sols.Err(); err != nil {
		return nil, fmt.Errorf("prolog sols error: %w", err)
	}

	return ids, nil
}
