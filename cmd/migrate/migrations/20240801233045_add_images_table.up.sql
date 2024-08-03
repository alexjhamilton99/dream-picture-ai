CREATE TABLE IF NOT EXISTS images (
  id SERIAL PRIMARY KEY,
  user_id UUID REFERENCES auth.users,
  status INT NOT NULL DEFAULT 1,
  prompt TEXT NOT NULL,
  deleted BOOLEAN NOT NULL DEFAULT FALSE,
  image_location TEXT,
  batch_id UUID NOT NULL,
  deleted_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
