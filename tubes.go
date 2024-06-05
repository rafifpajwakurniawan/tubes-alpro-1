package main

import (
	"fmt"
	"strings"
)

type Product struct {
	Name  string
	Brand string
	Type  string
	Price float64
	Stock int
}

type Inventory struct {
	Products []Product
}

func (inv *Inventory) AddProduct(p Product) {
	inv.Products = append(inv.Products, p)
}

func (inv *Inventory) FindProduct(name string) *Product {
	for i, p := range inv.Products {
		if p.Name == name {
			return &inv.Products[i]
		}
	}
	return nil
}

func (inv *Inventory) UpdateProduct(name string, newProduct Product) bool {
	for i, p := range inv.Products {
		if p.Name == name {
			inv.Products[i] = newProduct
			return true
		}
	}
	return false
}

func (inv *Inventory) DeleteProduct(name string) bool {
	for i, p := range inv.Products {
		if p.Name == name {
			inv.Products = append(inv.Products[:i], inv.Products[i+1:]...)
			return true
		}
	}
	return false
}

func (inv *Inventory) SortProductsBy(criteria string) {
	BubbleSort(inv.Products, criteria)
}

func BubbleSort(products []Product, criteria string) {
	n := len(products)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			swap := false
			switch criteria {
			case "price":
				if products[j].Price > products[j+1].Price {
					swap = true
				}
			case "name":
				if products[j].Name > products[j+1].Name {
					swap = true
				}
			case "brand":
				if products[j].Brand > products[j+1].Brand {
					swap = true
				}
			}
			if swap {
				products[j], products[j+1] = products[j+1], products[j]
			}
		}
	}
}

func (inv *Inventory) DisplayProducts() {
	headers := []string{"Name", "Brand", "Type", "Price", "Stock"}
	rows := [][]string{}

	for _, p := range inv.Products {
		row := []string{p.Name, p.Brand, p.Type, fmt.Sprintf("%.2f", p.Price), fmt.Sprintf("%d", p.Stock)}
		rows = append(rows, row)
	}

	printTable(headers, rows)
}

func printTable(headers []string, rows [][]string) {
	// Calculate column widths
	columnWidths := make([]int, len(headers))
	for i, header := range headers {
		columnWidths[i] = len(header)
	}
	for _, row := range rows {
		for i, cell := range row {
			if len(cell) > columnWidths[i] {
				columnWidths[i] = len(cell)
			}
		}
	}

	// Print headers with lines
	printLine(columnWidths, '┌', '┬', '┐')
	printRow(headers, columnWidths)
	printLine(columnWidths, '├', '┼', '┤')

	// Print rows
	for _, row := range rows {
		printRow(row, columnWidths)
	}
	printLine(columnWidths, '└', '┴', '┘')
}

func printRow(row []string, columnWidths []int) {
	for i, cell := range row {
		fmt.Printf("│ %-*s ", columnWidths[i], cell)
	}
	fmt.Println("│")
}

func printLine(columnWidths []int, left, mid, right rune) {
	fmt.Printf("%c", left)
	for i, width := range columnWidths {
		fmt.Printf("%s", strings.Repeat("─", width+2))
		if i < len(columnWidths)-1 {
			fmt.Printf("%c", mid)
		}
	}
	fmt.Printf("%c\n", right)
}

func main() {
	inv := Inventory{}

	// Menambahkan produk
	inv.AddProduct(Product{Name: "TV", Brand: "Samsung", Type: "Electronics", Price: 499.99, Stock: 10})
	inv.AddProduct(Product{Name: "Laptop", Brand: "Apple", Type: "Electronics", Price: 999.99, Stock: 5})
	inv.AddProduct(Product{Name: "Phone", Brand: "OnePlus", Type: "Electronics", Price: 299.99, Stock: 15})

	// Menampilkan semua produk
	fmt.Println("Semua produk:")
	inv.DisplayProducts()

	// Mencari produk
	fmt.Println("\nMencari produk 'Laptop':")
	if p := inv.FindProduct("Laptop"); p != nil {
		fmt.Printf("Ditemukan: %v\n", *p)
	} else {
		fmt.Println("Produk tidak ditemukan")
	}

	// Mengubah produk
	fmt.Println("\nMengubah produk 'Laptop':")
	if inv.UpdateProduct("Laptop", Product{Name: "Laptop", Brand: "Apple", Type: "Electronics", Price: 899.99, Stock: 7}) {
		fmt.Println("Produk berhasil diubah")
	} else {
		fmt.Println("Produk tidak ditemukan")
	}

	// Menampilkan semua produk setelah perubahan
	fmt.Println("\nSemua produk setelah perubahan:")
	inv.DisplayProducts()

	// Menghapus produk
	fmt.Println("\nMenghapus produk 'Phone':")
	if inv.DeleteProduct("Phone") {
		fmt.Println("Produk berhasil dihapus")
	} else {
		fmt.Println("Produk tidak ditemukan")
	}

	// Menampilkan semua produk setelah penghapusan
	fmt.Println("\nSemua produk setelah penghapusan:")
	inv.DisplayProducts()

	// Mengurutkan produk berdasarkan harga
	fmt.Println("\nMengurutkan produk berdasarkan harga:")
	inv.SortProductsBy("price")
	inv.DisplayProducts()
}
