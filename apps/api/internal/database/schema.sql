-- NOTOPOS AI Database Schema
-- Multi-tenant scalable architecture

-- Enable pgvector for AI features
CREATE EXTENSION IF NOT EXISTS vector;

-- Tenants Table
CREATE TABLE tenants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    domain VARCHAR(255) UNIQUE,
    logo_url TEXT,
    subscription_plan VARCHAR(50) DEFAULT 'free',
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Users Table (Multi-tenant)
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash TEXT,
    role VARCHAR(50) DEFAULT 'staff', -- owner, admin, manager, staff
    avatar_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(tenant_id, email)
);

-- Outlets Table
CREATE TABLE outlets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    address TEXT,
    phone VARCHAR(50),
    is_main BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Product Units (Satuan)
CREATE TABLE units (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    unit_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Product Categories
CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Products Table
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    category_id UUID REFERENCES categories(id) ON DELETE SET NULL,
    unit_id UUID REFERENCES units(id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    sku VARCHAR(100),
    barcode VARCHAR(100),
    description TEXT,
    price DECIMAL(15, 2) NOT NULL,
    cost_price DECIMAL(15, 2),
    stock_quantity INT DEFAULT 0,
    min_stock_level INT DEFAULT 10,
    expired_date DATE,
    image_url TEXT,
    embedding vector(1536),
    hpp_price DECIMAL(15, 2),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Transactions Table
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    outlet_id UUID NOT NULL REFERENCES outlets(id),
    user_id UUID NOT NULL REFERENCES users(id),
    invoice_number VARCHAR(100) UNIQUE NOT NULL,
    subtotal DECIMAL(15, 2) NOT NULL,
    tax DECIMAL(15, 2) DEFAULT 0,
    discount DECIMAL(15, 2) DEFAULT 0,
    total DECIMAL(15, 2) NOT NULL,
    payment_method VARCHAR(50) NOT NULL, -- cash, qris, debit, credit
    payment_status VARCHAR(50) DEFAULT 'paid', -- paid, pending, failed, refunded
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Transaction Items
CREATE TABLE transaction_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    transaction_id UUID NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
    product_id UUID NOT NULL REFERENCES products(id),
    quantity INT NOT NULL,
    price DECIMAL(15, 2) NOT NULL,
    total DECIMAL(15, 2) NOT NULL
);

-- Products Histories
CREATE TABLE product_histories (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    product_id BIGINT UNSIGNED NOT NULL,
    type ENUM(
        'sale',
        'purchase',
        'adjustment',
        'return_sale',
        'return_purchase',
        'expired',
        'damaged'
    ) NOT NULL,
    reference_id BIGINT UNSIGNED NULL,
    reference_type VARCHAR(50) NULL,
    qty DECIMAL(15,3) NOT NULL DEFAULT 0,
    price DECIMAL(15,2) NOT NULL DEFAULT 0,
    subtotal DECIMAL(15,2) NOT NULL DEFAULT 0,
    stock_before DECIMAL(15,3) NOT NULL DEFAULT 0,
    stock_after DECIMAL(15,3) NOT NULL DEFAULT 0,
    note VARCHAR(255) NULL,
    created_by BIGINT UNSIGNED NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_product_history_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    CONSTRAINT fk_product_history_user FOREIGN KEY (created_by) REFERENCES users(id)ON DELETE SET NULL
);

-- Inventory Movement Logs
CREATE TABLE inventory_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    product_id UUID NOT NULL REFERENCES products(id),
    outlet_id UUID NOT NULL REFERENCES outlets(id),
    quantity_change INT NOT NULL,
    type VARCHAR(50) NOT NULL, -- sale, restock, adjustment, return
    reference_id UUID, -- transaction_id or purchase_id
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- AI Insights & Logs
CREATE TABLE ai_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL, -- sales_forecast, recommendation, chatbot
    prompt TEXT,
    response TEXT,
    tokens_used INT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for performance
CREATE INDEX idx_products_tenant ON products(tenant_id);
CREATE INDEX idx_transactions_tenant ON transactions(tenant_id);
CREATE INDEX idx_inventory_logs_product ON inventory_logs(product_id);





🛒 Kasir / transaksi penjualan √
📦 Manajemen produk √
🏷️ Barcode scanner
📊 Stok barang √
⚠️ Notifikasi stok minimum √
🚚 Supplier
📥 Pembelian barang
📤 Penjualan barang 
👥 Pelanggan √
👤 Admin & kasir √
💰 Hutang/piutang jika diperlukan
💸 Pengeluaran toko
📈 Laporan penjualan √
📊 Laporan laba/rugi √
📦 Kartu stok
🧾 Cetak struk
🔎 Pencarian berdasarkan barcode/SKU

| Kendala                               | Contoh masalah                                   | Dampak                                |
| ------------------------------------- | ------------------------------------------------ | ------------------------------------- |
| 📦 **Stok tidak akurat**              | Tidak tahu jumlah barang yang tersisa            | Barang habis tanpa diketahui          |
| 📝 **Pencatatan masih manual**        | Transaksi dicatat di buku                        | Mudah salah/hilang                    |
| 💰 **Sulit menghitung keuntungan**    | Harga beli dan harga jual tidak terdokumentasi   | Tidak tahu laba sebenarnya            |
| 🛒 **Transaksi tidak tercatat semua** | Penjualan kecil sering tidak dicatat             | Omzet sulit dihitung                  |
| 📥 **Pembelian tidak terkontrol**     | Tidak ada riwayat pembelian dari supplier        | Sulit mengecek pengeluaran            |
| ⚠️ **Barang kedaluwarsa**             | Tidak mengetahui tanggal kedaluwarsa             | Barang rusak dan menyebabkan kerugian |
| 💳 **Hutang pelanggan**               | Bon pelanggan dicatat di kertas                  | Lupa siapa yang berhutang             |
| 🚚 **Pengelolaan supplier**           | Tidak ada catatan harga dan pembelian sebelumnya | Sulit membandingkan pembelian         |
| 💵 **Uang toko tercampur**            | Uang pribadi dan uang usaha bercampur            | Sulit mengetahui kondisi keuangan     |
| 📊 **Tidak ada laporan**              | Tidak ada laporan harian/bulanan                 | Pemilik sulit mengevaluasi usaha      |

