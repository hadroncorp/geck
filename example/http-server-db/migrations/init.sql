CREATE TABLE IF NOT EXISTS tasks (
    task_id VARCHAR(96) PRIMARY KEY,
    task_name VARCHAR(128) NOT NULL,
    status VARCHAR(64),
    create_time TIMESTAMP DEFAULT NOW(),
    create_by VARCHAR(96) DEFAULT 'system',
    last_update_time TIMESTAMP DEFAULT NOW(),
    last_update_by VARCHAR(96) DEFAULT 'system',
    is_active BOOL DEFAULT true,
    version BIGINT DEFAULT 0
);
