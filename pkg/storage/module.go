package storage

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func NewStore(ctx context.Context, cfg *App) (*Storage, error) {
	if cfg == nil {
		return nil, errors.New("no config error")
	}
	return &Storage{
		Store: make(map[string]Data),
		Cfg:   cfg,
	}, nil
}

func (s *Storage) Do() error {

	for {
		fmt.Println(
			"Chose action: \n" +
				"(1) Get\n" +
				"(2) Set\n" +
				"(3) Delete\n" +
				"Enter number:")
		input := ""
		fmt.Scan(&input)

		switch input {
		case "1":
			fmt.Println("Enter key: ")
			key := ""
			fmt.Scan(&key)
			val, found := s.Get(key)
			if !found {
				fmt.Println("No entity with such a key: ", key)
				continue
			}
			fmt.Println(val)
			continue
		case "2":
			fmt.Println("Enter key: ")
			key := ""
			fmt.Scan(&key)
			fmt.Println("Enter value: ")
			value := ""
			fmt.Scan(&value)
			ttl, _ := strconv.Atoi(s.Cfg.TTL)
			s.Set(key, value, time.Duration(ttl)*time.Second)
			_, found := s.Get(key)
			if !found {
				fmt.Println("Error saving, try again")
				continue
			}
			fmt.Println("Successfully saved")
			continue
		case "3":
			fmt.Println("Enter key: ")
			key := ""
			fmt.Scan(&key)
			s.Delete(key)
			_, found := s.Get(key)
			if !found {
				fmt.Println("Successfully deleted")
				continue
			}
			fmt.Println("Error deleting, try again")
			continue
		default:
			fmt.Println("try again (enter one number: 1, 2 or 3)")
			continue
		}
	}

}
func (s *Storage) Show() error {

	// Adding items to the store with different TTLs
	s.Set("key1", "value1", 5*time.Second)
	s.Set("key2", "value2", 10*time.Second)
	s.Set("key3", "value3", 0)  // No TTL
	s.Set("key4", "value4", -1) // TTL less than 0, ignored

	// Getting items from the store
	val1, found := s.Get("key1")
	if found {
		fmt.Println("Value of key1:", val1)
	} else {
		fmt.Println("Key1 not found")
	}

	val2, found := s.Get("key2")
	if found {
		fmt.Println("Value of key2:", val2)
	} else {
		fmt.Println("Key2 not found")
	}

	val3, found := s.Get("key3")
	if found {
		fmt.Println("Value of key3:", val3)
	} else {
		fmt.Println("Key3 not found")
	}

	// Waiting for some time to let the item with TTL expire
	time.Sleep(6 * time.Second)

	// Checking that the item with expired TTL was deleted
	val1AfterExpire, found := s.Get("key1")
	if found {
		fmt.Println("Value of key1 after expiration:", val1AfterExpire)
	} else {
		fmt.Println("Key1 not found after expiration")
	}

	// Deleting an item from the store
	s.Delete("key2")
	_, found = s.Get("key2")
	if !found {
		fmt.Println("Key2 deleted successfully")
	} else {
		fmt.Println("Key2 deletion failed")
	}
	return nil
}
