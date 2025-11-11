DROP TRIGGER IF EXISTS trg_tasks_update_timestamp ON tasks;
DROP TABLE IF EXISTS tasks;

DROP FUNCTION IF EXISTS _update_timestamp()