CREATE TYPE unit_enum AS ENUM
('kg', 'lb', 'mi', 'km');

CREATE TYPE exercise_type_enum AS ENUM
('strength', 'cardio');

CREATE TABLE exercises (
    exercise_id SERIAL PRIMARY KEY,
    exercise VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE routine (
    routine_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE routine_exercise (
    -- routine-exercise is a many-to-many relationship table
    routine_id INT NOT NULL,
    exercise_id INT NOT NULL,
    ordinal INT NOT NULL,
    exercise_type exercise_type_enum NOT NULL,
    -- strength
    sets INT,
    reps INT,
    weight FLOAT,
    -- cardio
    intensity INT,
    time INT,
    distance INT,
    unit unit_enum NOT NULL,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CHECK (
        (exercise_type = 'strength'::exercise_type_enum AND sets IS NOT NULL AND reps IS NOT NULL AND weight IS NOT NULL) OR
        (exercise_type = 'cardio'::exercise_type_enum AND intensity IS NOT NULL AND time IS NOT NULL AND distance IS NOT NULL)
    ),
    PRIMARY KEY (routine_id, exercise_id),
    FOREIGN KEY (routine_id) REFERENCES routine(routine_id),
    FOREIGN KEY (exercise_id) REFERENCES exercises(exercise_id)

);

CREATE TABLE homework (
    homework_id SERIAL PRIMARY KEY,
    due_date TIMESTAMP,
    completed TIMESTAMP,
    assigner INT NOT NULL,
    assignee INT NOT NULL,
    routine_id INT NOT NULL,
    deleted TIMESTAMP,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (assigner) REFERENCES users(user_id),
    FOREIGN KEY (assignee) REFERENCES users(user_id),
    FOREIGN KEY (routine_id) REFERENCES routine(routine_id)
);
