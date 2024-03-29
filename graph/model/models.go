package model

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalMapSlice(val map[string][]string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalMapSlice(v interface{}) (map[string][]string, error) {
	if m, ok := v.(map[string][]string); ok {
		return m, nil
	}

	return nil, fmt.Errorf("%T is not a map slice", v)
}

func MarshalTimeDuration(val time.Duration) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalTimeDuration(v interface{}) (time.Duration, error) {
	if m, ok := v.(time.Duration); ok {
		return m, nil
	}

	return 0, fmt.Errorf("%T is not a time duration", v)
}

func MarshalMapString(val map[string]string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalMapString(v interface{}) (map[string]string, error) {
	if m, ok := v.(map[string]string); ok {
		return m, nil
	}

	return nil, fmt.Errorf("%T is not a map string", v)
}
