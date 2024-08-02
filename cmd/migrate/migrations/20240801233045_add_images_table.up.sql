CREATE TABLE IF NOT EXISTS images (
  id SERIAL PRIMARY KEY,
  user_id UUID REFERENCES auth.users,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  status INT NOT NULL DEFAULT 1,
  image_location TEXT,
  prompt TEXT NOT NULL,
  deleted BOOLEAN NOT NULL DEFAULT FALSE,
  deleted_at TIMESTAMP
);
