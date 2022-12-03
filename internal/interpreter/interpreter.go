package interpreter

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/craigpastro/nungwi/internal/gen/nungwi/v1alpha"
	"github.com/ichiban/prolog"
)

type tuple struct {
	Namespace, Id, Relation, User string
}

type Interpreter struct {
	Interpreter *prolog.Interpreter
}

func NewInterpreter(interpreter *prolog.Interpreter) *Interpreter {
	return &Interpreter{
		Interpreter: interpreter,
	}
}

func (i *Interpreter) AddTuples(ctx context.Context, tuples []*pb.Tuple) error {
	facts := strings.Join(transformTuples(tuples), " ")
	fmt.Println(facts)

	if err := i.Interpreter.ExecContext(ctx, facts); err != nil {
		return fmt.Errorf("interpreter error: %w", err)
	}

	return nil
}

func (i *Interpreter) DeleteTuples(ctx context.Context, tuples []*pb.Tuple) error {
	sols, err := i.Interpreter.QueryContext(ctx, "tuple(A, B, C, X).")
	if err != nil {
		return fmt.Errorf("interpreter error: %w", err)
	}
	defer sols.Close()

	for sols.Next() {
		var t tuple
		if err := sols.Scan(&t); err != nil {
			return fmt.Errorf("scan error: %w", err)
		}
		fmt.Printf(">>> %s %s %s %s\n", t.Namespace, t.Id, t.Relation, t.User)
	}

	if err := sols.Err(); err != nil {
		return fmt.Errorf("scan end error: %w", err)
	}

	return nil
}

// func transformSchema(namespaces []*pb.Namespace) []string {
// 	return nil
// }

func transformTuples(tuples []*pb.Tuple) []string {
	res := make([]string, 0, len(tuples))

	for _, tuple := range tuples {
		var user string
		switch u := tuple.User.Ref.(type) {
		case *pb.User_Id:
			user = u.Id
		case *pb.User_Object:
			user = fmt.Sprintf("object(%s, %s)", u.Object.Namespace, u.Object.Id)
		case *pb.User_Userset:
			user = fmt.Sprintf("userset(%s, %s, %s)", u.Userset.Object.Namespace, u.Userset.Object.Id, u.Userset.Relation)

		}
		res = append(res, fmt.Sprintf("tuple(%s, %s, %s, %s).", tuple.Object.Namespace, tuple.Object.Id, tuple.Relation, user))
	}

	return res
}
