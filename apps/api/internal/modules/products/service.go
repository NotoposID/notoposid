package products

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	CreateProduct(product *Product) error
}

type productService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product *Product) error {
	// 1. Validasi Harga (Pricing Logic)
	if product.Price < product.CostPrice {
		return errors.New("harga jual (price) tidak boleh lebih kecil dari harga modal (cost_price)")
	}

	// 2. Pengecekan Duplikasi SKU / Barcode
	if product.SKU != "" {
		existing, err := s.repo.GetBySKU(product.SKU, product.TenantID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if existing != nil && existing.ID != uuid.Nil {
			return errors.New("SKU sudah terdaftar")
		}
	}

	if product.Barcode != "" {
		existing, err := s.repo.GetByBarcode(product.Barcode, product.TenantID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if existing != nil && existing.ID != uuid.Nil {
			return errors.New("Barcode sudah terdaftar")
		}
	}

	// 3. Auto-Generate SKU jika kosong
	if product.SKU == "" {
		prefix := "PRD"
		if product.CategoryID != nil {
			category, err := s.repo.GetCategoryByID(*product.CategoryID)
			if err == nil && len(category.Name) >= 3 {
				prefix = strings.ToUpper(category.Name[:3])
			}
		}
		// Generate simple unique sequence using timestamp
		product.SKU = fmt.Sprintf("%s-%d", prefix, time.Now().Unix()%100000)
	}

	// Simpan ke database
	if err := s.repo.Create(product); err != nil {
		return err
	}

	// 4. Integrasi dengan Modul Lain (Event Publishing - Mock)
	log.Printf("[EVENT PUBLISHED] Produk baru berhasil ditambahkan: %s (SKU: %s)\n", product.Name, product.SKU)

	return nil
}
