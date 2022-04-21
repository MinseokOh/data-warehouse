package transformer

import (
	"database/sql"
	"fmt"
	"github.com/MinseokOh/data-warehouse/types"
	"github.com/MinseokOh/data-warehouse/util/log"
	"sync"
	"time"
)

type Sync struct {
	products   []*types.Sdk
	db         *sql.DB
	logger     *log.Logger
	insertChan chan *types.Sdk
	mux        sync.RWMutex

	syncCount int64
}

func NewSync(config Config) *Sync {
	db, err := sql.Open(
		config.DB.GetDriverName(),
		config.DB.GetDBSource(),
	)

	if err != nil {
		panic(err)
	}

	return &Sync{
		products:   make([]*types.Sdk, 0),
		insertChan: make(chan *types.Sdk),
		logger:     log.NewLoggerConfig("transformer/sync", config.Logger),
		db:         db,
		mux:        sync.RWMutex{},
		syncCount:  0,
	}
}

func (sync *Sync) GetSyncCount() int64 { return sync.syncCount }

func (sync *Sync) pop() []*types.Sdk {
	sync.logger.Debug("pop products", len(sync.products))

	sync.mux.Lock()
	slice := make([]*types.Sdk, len(sync.products))
	copy(slice, sync.products[0:len(sync.products)])
	sync.products = make([]*types.Sdk, 0)
	sync.mux.Unlock()

	return slice
}

func (sync *Sync) Insert(Sdk *types.Sdk) {
	sync.insertChan <- Sdk
}

func (sync *Sync) insert(Sdk *types.Sdk) {
	sync.mux.Lock()
	sync.products = append(sync.products, Sdk)
	sync.mux.Unlock()
}

func (sync *Sync) Run() {
	go sync.sync()

	for {
		select {
		case Sdk := <-sync.insertChan:
			sync.insert(Sdk)
		}
	}
}

func (sync *Sync) sync() {
	for range time.Tick(time.Second) {
		sync.logger.Debug("sync products to database")

		products := sync.pop()
		if len(products) == 0 {
			continue
		}

		query := sync.GetQuery(products)
		_, err := sync.db.Exec(query)
		if err != nil {
			sync.logger.Error(err)
		}

		sync.syncCount += int64(len(products))
	}
}

func (sync *Sync) GetQuery(products []*types.Sdk) string {
	query := `INSERT INTO SDK(ID, NAME, VERSION) VALUES`
	for _, Sdk := range products {
		query += fmt.Sprintf(" ('%s', '%s', '%s'),",
			Sdk.GetId(),
			Sdk.GetName(),
			Sdk.GetVersion(),
		)
	}
	return query[:len(query)-1]
}
