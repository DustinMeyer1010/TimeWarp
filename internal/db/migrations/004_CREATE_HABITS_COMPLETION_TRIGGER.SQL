CREATE OR REPLACE FUNCTION recalc_habit_completions()
RETURNS TRIGGER AS $$
DECLARE
    total_completions INT;
    habit_completion_time INTERVAL;
    i INT;
BEGIN
    -- Get the completion_time from the related habit
    SELECT completion_time INTO habit_completion_time
    FROM habits
    WHERE id = NEW.task_id;

    -- Only proceed if habit has a valid completion_time
    IF habit_completion_time IS NOT NULL AND habit_completion_time > INTERVAL '0 seconds' THEN

        -- Delete existing completions for this time log
        DELETE FROM habits_completed WHERE time_log_id = NEW.id;

        -- Calculate how many full completions occurred
        total_completions := FLOOR(
            EXTRACT(EPOCH FROM NEW.time_spent) / EXTRACT(EPOCH FROM habit_completion_time)
        );

        -- Insert a row for each completion
        FOR i IN 1..total_completions LOOP
            INSERT INTO habits_completed(task_id, time_log_id, completion_date)
            VALUES (NEW.task_id, NEW.id, NOW());
        END LOOP;

    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER habits_time_logs_completion_trigger
AFTER INSERT OR UPDATE ON habits_time_logs
FOR EACH ROW
EXECUTE FUNCTION recalc_habit_completions();