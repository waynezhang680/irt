package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateUsersTable, downCreateUsersTable)
}

func upCreateUsersTable(tx *sql.Tx) error {
	// 首先检查表是否存在
	var exists bool
	err := tx.QueryRow(`
        SELECT EXISTS (
            SELECT FROM information_schema.tables 
            WHERE table_schema = 'public' 
            AND table_name = 'users'
        );
    `).Scan(&exists)
	if err != nil {
		return err
	}

	// 如果表不存在，创建表
	if !exists {
		_, err = tx.Exec(`
            CREATE TABLE users (
                id SERIAL PRIMARY KEY,
                username VARCHAR(50) NOT NULL,
                email VARCHAR(100) NOT NULL,
                password_hash VARCHAR(255) NOT NULL,
                role VARCHAR(20) NOT NULL,
                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                deleted_at TIMESTAMP WITH TIME ZONE
            );
        `)
		if err != nil {
			return err
		}

		// 添加唯一索引
		_, err = tx.Exec(`
            CREATE UNIQUE INDEX IF NOT EXISTS idx_users_username ON users(username) WHERE deleted_at IS NULL;
            CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email) WHERE deleted_at IS NULL;
        `)
		if err != nil {
			return err
		}
	}

	return nil
}

func downCreateUsersTable(tx *sql.Tx) error {
	// 删除表（如果存在）
	_, err := tx.Exec(`DROP TABLE IF EXISTS users;`)
	return err
}
