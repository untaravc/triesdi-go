CREATE INDEX idx_activities_done_at_activity_type_calories_burned 
ON activities(done_at, activity_type, calories_burned);