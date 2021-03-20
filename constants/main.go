package main

import "fmt"

// create a constant that's our global limit size
const GlobalLimit int = 100

// create a MaxCacheSize that is 10 times the global limit size.
const MaxCacheSize int = 10 * GlobalLimit

// create our cache prefixes
const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

// declare a `map` that has a `string` for a key and `string` for values.
var cache map[string]string

// create a function to get items from the cache
func cacheGet(key string) string {
	return cache[key]
}

// create a function that sets items in the cache.
func cacheSet(key, val string) {
	// check out the MaxCacheSize to stop the cache going over that size
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

// create a function to get a book from the cache
func GetBook(isbn string) string {
	// use the book cache prefix to create a unique key
	return cacheGet(CacheKeyBook + isbn)
}

// create a function to add a book to the cache
func SetBook(isbn, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

// create a function to get CD data from the cache
func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

// create a function to add CD's to the shared cache
func SetCD(sku, title string) {
	cacheSet(CacheKeyCD+sku, title)
}

func main() {
	// initialize our cache by creating a `map`
	cache = make(map[string]string)

	// add a book to the cache
	SetBook("1234-5678", "Get Ready To Go")

	// add a cd to the cache
	SetCD("1234-5678", "Get Ready To Go Audio Book")

	// get and print the book from the cache
	fmt.Println("Book	:", GetBook("1234-5678"))

	// get and print the CD from the cache
	fmt.Println("CD	:", GetCD("1234-5678"))
}
