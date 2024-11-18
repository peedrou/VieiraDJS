CREATE KEYSPACE IF NOT EXISTS my_keyspace 
WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

USE my_keyspace;

CREATE TABLE IF NOT EXISTS task_history (
    job_id uuid,
    execution_time timestamp,
    last_update_time timestamp,
    retry_count int,
    status text,
    PRIMARY KEY (job_id, execution_time)
);

CREATE TABLE IF NOT EXISTS task_schedule (
    next_execution_time bigint,
    job_id uuid,
    PRIMARY KEY (next_execution_time, job_id)
);

CREATE TABLE IF NOT EXISTS jobs (
    job_id uuid PRIMARY KEY,
    created_time timestamp,
    interval text,
    is_recurring boolean,
    max_retries int,
    start_time timestamp
);
