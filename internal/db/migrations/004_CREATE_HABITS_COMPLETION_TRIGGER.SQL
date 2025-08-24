CREATE OR REPLACE FUNCTION recalc_habit_completions_for_update_spent_time_change()
RETURNS TRIGGER AS $$
DECLARE
    total_completions INT;
    habit_completion_time INTERVAL;
    i INT;
BEGIN
    -- Get the completion_time from the related habit
    SELECT completion_time INTO habit_completion_time
    FROM habits
    WHERE id = NEW.habit_id;

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
            INSERT INTO habits_completed(habit_id, time_log_id, completion_date)
            VALUES (NEW.habit_id, NEW.id, NOW());
        END LOOP;

    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER habits_time_logs_completion_trigger
AFTER INSERT OR UPDATE ON habits_time_logs
FOR EACH ROW
EXECUTE FUNCTION recalc_habit_completions_for_update_spent_time_change();


CREATE OR REPLACE FUNCTION recalc_habit_completion_for_completion_time_change()
RETURNS TRIGGER AS $$
DECLARE 
    total_completions INT;
    log_record RECORD;
    i INT;

BEGIN
    IF NEW.completion_time IS NOT NULL AND NEW.completion_time > INTERVAL '0 seconds' THEN
        FOR log_record IN
            SELECT * FROM habits_time_logs WHERE habit_id = NEW.id
            LOOP
                DELETE FROM habits_completed WHERE time_log_id = log_record.id;

                total_completions := FLOOR(
                    EXTRACT(EPOCH FROM log_record.time_spent) / EXTRACT(EPOCH FROM NEW.completion_time)
                );

            FOR i IN 1..total_completions LOOP
                    INSERT INTO habits_completed(habit_id, time_log_id, completion_date)
                    VALUES (log_record.id, NEW.id, NOW());
                END LOOP;
            END LOOP;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


DROP TRIGGER IF EXISTS recalc_habit_completion_for_completion_time_change ON habits;
CREATE TRIGGER recalc_habit_completion_for_completion_time_change
AFTER UPDATE OF completion_time ON habits
FOR EACH ROW
EXECUTE FUNCTION recalc_habit_completion_for_completion_time_change();