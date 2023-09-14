package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/boltdb/bolt"
)

const (
	dbDir     = "mydatabase"
	bucket    = "mybucket"
	dataKey   = "mykey"
	dataValue = "myvalue"
	logFile   = "logs/log.txt"
	logBackup = 10
)

var logger *log.Logger

func init() {
	// Creates the logs directory
	err := os.MkdirAll(filepath.Dir(logFile), os.ModePerm)
	if err != nil {
		log.Fatal("Failed to create logs directory:", err)
	}

	// Opens the log file
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	logger = log.New(file, "", log.Ldate|log.Ltime)
}

func main() {
	dbPath, err := getDBPath("mydatabase.db")
	if err != nil {
		logger.Fatal("Failed to get database path:", err)
	}

	// Opens the BoltDB database
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		logger.Fatal("Failed to open the database:", err)
	}
	defer db.Close()

	// Creates a bucket (if it doesn't exist)
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
	if err != nil {
		logger.Fatal("Failed to create bucket:", err)
	}

	err = db.Update(func(tx *bolt.Tx) error { // Adds data to bucket
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(dataKey), []byte(dataValue))
		return err
	})
	if err != nil {
		logger.Fatal("Failed to add data to the bucket:", err)
	}

	// Retrieves data from the bucket
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(dataKey))
		logger.Printf("Data retrieved: %s\n", v)
		return nil
	})
	if err != nil {
		logger.Fatal("Failed to retrieve data from the bucket:", err)
	}

	logger.Println("BoltDB setup, data added, and maintenance completed successfully")
}

func getDBPath(dbName string) (string, error) {
	// Creates the database directory if it doesn't exist
	err := os.MkdirAll(dbDir, os.ModePerm)
	if err != nil {
		log.Println("Failed to create database directory:", err)
		return "", err
	}

	dbPath := filepath.Join(dbDir, dbName)

	return dbPath, nil
}
