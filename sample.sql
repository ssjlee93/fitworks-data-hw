-- Insert into exercises table
INSERT INTO exercises (exercise, description)
VALUES
('Bench Press', 'A strength exercise for chest muscles.'),
('Running', 'Cardio exercise typically done on a treadmill or outside.'),
('Deadlift', 'Strength training exercise targeting the lower back and legs.');

-- Insert into routine table
INSERT INTO routine (name, description)
VALUES
('Full Body Strength', 'A full-body strength training routine.'),
('Morning Cardio', 'Light cardio session to start the day.'),
('Leg Day', 'Focused on lower body strength.');

-- Insert into routine_exercise table
-- Strength exercises
INSERT INTO routine_exercise (routine_id, exercise_id, ordinal, exercise_type, sets, reps, weight, unit)
VALUES
(1, 1, 1, 'strength', 4, 10, 80.0, 'kg'),
(3, 3, 1, 'strength', 3, 8, 100.0, 'lb'),
(3, 1, 2, 'strength', 5, 5, 90.0, 'kg');

-- Cardio exercises
INSERT INTO routine_exercise (routine_id, exercise_id, ordinal, exercise_type, intensity, time, distance, unit)
VALUES
(3, 2, 1, 'cardio', 7, 30, 5.0, 'km'),
(1, 2, 2, 'cardio', 8, 45, 3.0, 'mi'),
(1, 3, 3, 'cardio', 6, 20, 2.0, 'km');

-- Insert into homework table
INSERT INTO homework (due_date, completed, assigner, assignee, routine_id, deleted)
VALUES
('2024-03-20 10:00:00', NULL, 1, 2, 1, NULL),
('2024-03-21 12:00:00', '2024-03-21 11:50:00', 1, 3, 2, NULL),
('2024-03-22 09:30:00', NULL, 2, 1, 3, NULL);

