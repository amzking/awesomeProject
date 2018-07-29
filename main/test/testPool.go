package testPool

import (
	"log"
	"io"
	"sync/atomic"
	"sync"
	"time"
	"math/rand"
)

const (
	maxGorutines = 25

	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

func (dbConn * dbConnection) Close() (error) {
	log.Println("CLose : Connection", dbConn.ID)
	return nil
}

var idCounter int32


func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create : New COnnection", id)

	return &dbConnection{id}, nil
}


func main() {
	var wg sync.WaitGroup
	wg.Add(maxGorutines)

	p, err := New1(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGorutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("Shutdown Program")
	p.Close();
}
func performQueries(query int, pool *Pool) {

	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer pool.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000))  * time.Millisecond)
	log.Println("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)

}