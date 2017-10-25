package main

import (
	"context"

	"fmt"

	"github.com/fzerorubigd/k8s-netdata-proxy/sets"
)

func routes(ctx context.Context, change <-chan sets.String) {
	for {
		select {
		case <-ctx.Done():
			return
		case rec := <-change:
			for i := range rec {
				fmt.Println(i, "=>", rec[i])
			}
		}
	}
}
