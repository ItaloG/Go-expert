package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
	// Products []Product // has many
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	// CategoryID   int // belong to
	// Category     Category // belong to
	Categories   []Category `gorm:"many2many:products_categories;"`
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	/**
	lockando uma linha com transaction


	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Eletronicos"
	tx.Debug().Save(&c)
	tx.Commit()
	*/

	/**
	relacionamento MANY TO MANY

	category := Category{Name: "Cozinha"}
	db.Create(&category)

	category2 := Category{Name: "Eletronicos"}
	db.Create(&category2)

	db.Create(&Product{
		Name:       "Panela",
		Price:      99.00,
		Categories: []Category{category, category2},
	})


	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name)
		}
	}
	*/

	/**
	relacionamento HAS MANY pegadinha com SerialNumber
	o Model Category não tem relacionamento com SerialNumber,
	por conta disso não podemos chamar o SerialNumber diretamente no Proloader

	category := Category{Name: "Cozinha"}
	db.Create(&category)

	db.Create(&Product{
		Name:       "Panela",
		Price:      99.00,
		CategoryID: 1,
	})

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name, "Number:", product.SerialNumber.Number)
		}
	}
	*/

	/**
	relacionamento HAS MANY

	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	db.Create(&Product{
		Name:       "Notebook",
		Price:      1000.00,
		CategoryID: 1,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name)
		}
	}
	*/

	/**
	relacionamento HAS ONE


	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	db.Create(&Product{
		Name:       "Notebook",
		Price:      1000.00,
		CategoryID: 1,
	})

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
	*/

	/**
	relacionamento BELONGS TO

	create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	db.Create(&Product{
		Name:       "Notebook",
		Price:      1000.00,
		CategoryID: category.ID,
	})

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}
	*/

	// create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.00,
	// })

	// create batch
	// products := []Product{
	// 	{Name: "Notebook", Price: 1000.00},
	// 	{Name: "Mouse", Price: 50.00},
	// 	{Name: "Keyboard", Price: 100.00},
	// }
	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 2)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// pagination
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where
	// var products []Product
	// db.Where("price > ?", 100).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// like
	// var products []Product
	// db.Where("name LIKE ?", "%mou%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// update
	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)
	// db.Delete(&p2)

}
