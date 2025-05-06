-- Создание таблицы users
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                     deleted_at TIMESTAMP WITH TIME ZONE,
                                     username VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role VARCHAR(50) DEFAULT 'USER'
    );

-- Создание таблицы shoes
CREATE TABLE IF NOT EXISTS shoes (
                                     id SERIAL PRIMARY KEY,
                                     brand VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    size INTEGER NOT NULL CHECK (size > 0),
    price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    description TEXT
    );
