package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}

func main() {

	//url API

	url := "https://fakestoreapi.com/products"

	// membuat channel untuk mengirim hasil request dari goroutine
	ch := make(chan []Product)

	// melakukan request ke endpoint menggunakan goroutine
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var products []Product
		err = json.Unmarshal(body, &products)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// mengirim hasil request ke channel
		ch <- products
	}()

	// menerima hasil request dari channel
	products := <-ch

	// menampilkan data produk
	for _, product := range products {
		fmt.Println("ID:", product.ID)
		fmt.Println("Title:", product.Title)
		fmt.Println("Price:", product.Price)
		fmt.Println("Description:", product.Description)
		fmt.Println("Category:", product.Category)
		fmt.Println("Image:", product.Image)
		fmt.Println()
	}
}
