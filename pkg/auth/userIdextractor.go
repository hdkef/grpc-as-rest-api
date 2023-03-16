package auth

import (
	"context"
	"errors"

	"google.golang.org/grpc/metadata"
)

func ExtractUserID(ctx context.Context) (string, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return "", errors.New("unauthorized no metadata provided")
	}
	userIdSlice := md.Get("x-user-id")
	if len(userIdSlice) < 1 {
		return "", errors.New("unauthorized no userId provided")
	}

	return userIdSlice[0], nil
}
