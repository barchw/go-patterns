package main

import (
	"fmt"
	"log"

	"claude-test/internal/creational/builder"
)

func main() {
	fmt.Println("=== Builder Pattern ===")

	query, err := builder.NewQueryBuilder().
		Select("id", "name", "email").
		From("users").
		Where("active = true").
		Where("age > 18").
		OrderBy("name").
		Limit(10).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(query)

	simple, err := builder.NewQueryBuilder().
		Select("*").
		From("products").
		Build()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(simple)
}
