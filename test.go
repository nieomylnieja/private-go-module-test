package privategomoduletest

import (
	"github.com/pkg/errors"
	_ "golang.org/x/sync/errgroup"
)

func Test() error {
	return errors.New("test")
}
