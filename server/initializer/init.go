package initializer

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/ChocolateAceCream/blog/global"
	"go.uber.org/zap"
)

const (
	InitOrderInternal = 10
	InitOrderExternal = 100
)

type Initializer interface {
	Name() string
	Initialize(ctx context.Context) (next context.Context, err error)
	InitDataVerify(ctx context.Context) bool
}

type orderedInitializer struct {
	order int
	Initializer
}

type initializerGroup []*orderedInitializer //

var (
	initializers initializerGroup
	cache        map[string]*orderedInitializer
)

type InitService struct{}

func Register(order int, i Initializer) {
	name := i.Name()

	if cache == nil {
		cache = make(map[string]*orderedInitializer)
	}
	if initializers == nil {
		initializers = []*orderedInitializer{}
	}

	if _, existed := cache[name]; existed {
		global.LOGGER.Error("Fail to register initializer", zap.Error(fmt.Errorf("name conflict on %s", name)))
		panic(fmt.Sprintf("Initializer name conflict on %s", name))
	}
	ni := orderedInitializer{order, i}
	initializers = append(initializers, &ni)
	cache[name] = &ni
}

func (initService *InitService) InitDB() (err error) {
	ctx := context.TODO()
	if len(initializers) == 0 {
		return errors.New("no initializer registered, please check again")
	}

	//sort from lower order to higher order, so initiliaze with dependency will execute later
	sort.Sort(&initializers)

	// execute Initialize function for each initializer in initializers
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) {
		c()
	}(cancel)
	for _, init := range initializers {
		if init.InitDataVerify(next) {
			// if init data already inserted, jump to next init
			global.LOGGER.Info(fmt.Sprintf("%s has been initialized", init.Name()))
			continue
		}
		if n, err := init.Initialize(next); err != nil {
			global.LOGGER.Error(fmt.Sprintf("failed to initialize %s", init.Name()))
			return err
		} else {
			next = n
		}
	}
	global.LOGGER.Info("done data initialization")
	return nil
}

/* -- sortable interface -- */
func (ig initializerGroup) Len() int {
	return len(ig)
}

func (ig initializerGroup) Less(i, j int) bool {
	return ig[i].order < ig[j].order
}

func (ig initializerGroup) Swap(i, j int) {
	ig[i], ig[j] = ig[j], ig[i]
}
