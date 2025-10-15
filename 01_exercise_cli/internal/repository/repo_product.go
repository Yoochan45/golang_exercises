package repository

import (
	"01_exercise/internal/entity"
	"context" // ✅ digunakan untuk mengontrol lifecycle query (timeout, cancel, dsb)
	"database/sql"
	"fmt"
	"log"
)

// ProductRepository berfungsi sebagai layer untuk berinteraksi dengan tabel "products" di database.
// Layer ini memisahkan logika database dari logika bisnis (service/controller).
type ProductRepository struct {
	DB *sql.DB
}

// NewProductRepository membuat instance baru dari ProductRepository.
// Ini bagian dari dependency injection: kita "menyuntikkan" *sql.DB dari luar.
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// GetAll mengambil semua data produk dari tabel products.
func (r *ProductRepository) GetAll(ctx context.Context) ([]entity.Products, error) {
	// ✅ gunakan QueryContext agar bisa dikontrol dengan context (timeout, cancel)
	rows, err := r.DB.QueryContext(ctx, `
		SELECT product_id, name, description, price, stock
		FROM products
		ORDER BY product_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []entity.Products{}

	for rows.Next() {
		var p entity.Products
		// &p.Field → memberi alamat agar Scan() bisa isi data langsung ke variabel
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock); err != nil {
			fmt.Println("Scan error:", err)
			continue
		}
		products = append(products, p)
	}
	return products, nil
}

// AddProduct menambahkan produk baru ke database.
func (r *ProductRepository) AddProduct(ctx context.Context, name, description string, price float64, stock int) error {
	// ✅ gunakan ExecContext agar query bisa dibatalkan lewat context
	_, err := r.DB.ExecContext(ctx, `
		INSERT INTO products (name, description, price, stock)
		VALUES ($1, $2, $3, $4)
	`, name, description, price, stock)

	if err != nil {
		log.Println("Failed to add product:", err)
	}
	return err
}

// DeleteProduct menghapus produk berdasarkan ID.
func (r *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, `DELETE FROM products WHERE product_id = $1`, id)
	if err != nil {
		log.Println("Failed to delete product:", err)
	}
	return err
}

// UpdateProduct memperbarui data produk berdasarkan ID.
func (r *ProductRepository) UpdateProduct(ctx context.Context, name, description string, price float64, stock, id int) error {
	_, err := r.DB.ExecContext(ctx, `
		UPDATE products
		SET name = $1, description = $2, price = $3, stock = $4
		WHERE product_id = $5
	`, name, description, price, stock, id)

	if err != nil {
		log.Println("Failed to update product:", err)
	}
	return err
}

// ShowProductById mengambil satu produk berdasarkan ID-nya.
func (r *ProductRepository) ShowProductById(ctx context.Context, id int) (entity.Products, error) {
	var product entity.Products

	// ✅ QueryRowContext lebih efisien untuk ambil satu data
	err := r.DB.QueryRowContext(ctx, `
		SELECT product_id, name, description, price, stock
		FROM products
		WHERE product_id = $1
	`, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock)

	if err != nil {
		if err == sql.ErrNoRows {
			return product, fmt.Errorf("product with id %d not found", id)
		}
		log.Println("Error fetching product:", err)
		return product, err
	}

	return product, nil
}

// ReduceStock mengurangi stok produk berdasarkan jumlah pembelian.
func (r *ProductRepository) ReduceStock(ctx context.Context, productID, quantity int) error {
	var stock int

	// Ambil stok sekarang dari database
	err := r.DB.QueryRowContext(ctx, `
		SELECT stock FROM products WHERE product_id = $1
	`, productID).Scan(&stock)
	if err != nil {
		return fmt.Errorf("failed to fetch stock for product ID %d: %w", productID, err)
	}

	if stock < quantity {
		return fmt.Errorf("insufficient stock for product ID %d: current stock %d, requested %d",
			productID, stock, quantity)
	}

	// Update stok setelah pembelian
	_, err = r.DB.ExecContext(ctx, `
		UPDATE products SET stock = stock - $1 WHERE product_id = $2
	`, quantity, productID)
	if err != nil {
		return fmt.Errorf("failed to reduce stock for product ID %d: %w", productID, err)
	}

	return nil
}
