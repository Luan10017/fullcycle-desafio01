-- Script SQL para criar a tabela cotacao no SQLite
-- Armazena cada campo de CurrencyRate em colunas individuais

CREATE TABLE IF NOT EXISTS cotacaos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    code TEXT NOT NULL,
    codein TEXT NOT NULL,
    name TEXT,
    high TEXT,
    low TEXT,
    var_bid TEXT,
    pct_change TEXT,
    bid TEXT,
    ask TEXT,
    timestamp TEXT,
    create_date TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME
);

-- Índices para melhor performance
CREATE INDEX IF NOT EXISTS idx_cotacaos_code ON cotacaos(code);
CREATE INDEX IF NOT EXISTS idx_cotacaos_created_at ON cotacaos(created_at);
CREATE INDEX IF NOT EXISTS idx_cotacaos_deleted_at ON cotacaos(deleted_at);
