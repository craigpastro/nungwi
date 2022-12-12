package validate

import (
	"fmt"

	pb "github.com/craigpastro/nungwi/internal/gen/nungwi/v1alpha"
	"github.com/craigpastro/nungwi/internal/parser"
)

func ValidateTuples(tuples []*pb.Tuple) error {
	for _, tuple := range tuples {
		if ok := parser.ParseUser(tuple.GetUser()); !ok {
			return fmt.Errorf("invalid user in: %s", tuple.String())
		}
	}

	return nil
}

func ValidateConfigs(configs []*pb.RelationConfig) error {
	for _, config := range configs {
		if ok := parser.ParseRewrite(config.GetRewrite()); !ok {
			return fmt.Errorf("invalid rewrite in: %s", config.String())
		}
	}

	return nil
}
