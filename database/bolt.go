package database

import (
	"encoding/json"
	"fmt"

	"com.github/Fszta/gows/pkg/dag"
	"github.com/boltdb/bolt"
)

var (
	rootBucket = "DB"
	dagBucket  = "DAGS"
)

type BoltClient struct {
	boltDb *bolt.DB
}

// Init bolt db, create buckets if not exists
func (client *BoltClient) init() error {
	var err error
	client.boltDb, err = bolt.Open("/var/lib/gows/gows.db", 384, nil)
	if err != nil {
		fmt.Printf("Fail to open db %v\n", err)
	}

	err = client.boltDb.Update(func(tx *bolt.Tx) error {
		// Create root bucket
		root, err := tx.CreateBucketIfNotExists([]byte(rootBucket))
		if err != nil {
			return fmt.Errorf("Fail to create root bucket: %v\n", err)
		}
		// Create dags bucket
		_, err = root.CreateBucketIfNotExists([]byte(dagBucket))
		if err != nil {
			fmt.Printf("Fail to create dag bucket %v\n", err)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Fail to setup buckets, %v\n", err)
	}
	fmt.Println("Database successfully initialized")

	return nil
}

// SaveDagConfig insert dag configuration into database
func (client *BoltClient) SaveDagConfig(dagConfig *dag.DagConfig) error {
	encoded, err := json.Marshal(dagConfig)
	if err != nil {
		fmt.Println("Fail to marshal dag config")
	}

	err = client.boltDb.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte(rootBucket)).Bucket([]byte(dagBucket)).Put([]byte(dagConfig.DagName), encoded)
		if err != nil {
			fmt.Println("Could not add dag configuration to database")
		}
		return nil
	})
	return err
}

// GetAllDags return all dags configuration from dag bucket
func (client *BoltClient) GetAllDags() []dag.DagConfig {
	var storedDags []dag.DagConfig
	var data dag.DagConfig

	client.boltDb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(rootBucket)).Bucket([]byte(dagBucket))
		b.ForEach(func(k, v []byte) error {
			json.Unmarshal(v, &data)
			storedDags = append(storedDags, data)
			return nil
		})
		return nil
	})
	return storedDags
}

// GetDag return dag configuration based on UUID key
func (client *BoltClient) GetDag(dagUUID string) (*dag.DagConfig, error) {

	var dagConfig *dag.DagConfig

	err := client.boltDb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(rootBucket)).Bucket([]byte(dagBucket))
		if b == nil {
			return fmt.Errorf("Bucket %s not found", dagBucket)
		}
		dagValue := b.Get([]byte(dagUUID))
		json.Unmarshal(dagValue, &dagConfig)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return dagConfig, err
}

// RemoveDag drop dag from database using its UUID
func (client *BoltClient) RemoveDag(dagUUID string) error {
	return client.boltDb.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(rootBucket)).Bucket([]byte(dagBucket))
		return bucket.Delete([]byte(dagUUID))
	})
}

func (client *BoltClient) Close() {
	client.boltDb.Close()
}

func UpdateDagConfig() {

}
