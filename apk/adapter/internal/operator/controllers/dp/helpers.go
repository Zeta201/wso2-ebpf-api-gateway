package dp

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func addressOf[T any](v T) *T {
	return &v
}

// splits the host and port
func splitHostNamspace(hostNamespace string) []string {
	split := strings.Split(hostNamespace, ".")
	if len(split) == 2 {
		return split
	}
	if len(split) == 1 {
		return []string{split[0], "default"}
	}
	return []string{"dummy-service", "default"}
}

func toAny(message proto.Message) *anypb.Any {
	a, err := anypb.New(message)
	if err != nil {
		return nil
	}
	return a
}

func getClusterName(ns, name string) string {
	// the name is having the format of "namespace:name:port"
	// -> slash would prevent ParseResources from rewriting with CEC namespace and name!
	return fmt.Sprintf("%s/%s", ns, name)
}
