-- Create users table
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name TEXT,
  company_name TEXT,
  email TEXT,
  password TEXT,
  office TEXT,
  linkedin_link TEXT,
  interest TEXT,
  image TEXT
);

-- Create projects table
CREATE TABLE IF NOT EXISTS projects (
  id SERIAL PRIMARY KEY,
  name TEXT,
  description TEXT,
  macro_setor TEXT,
  micro_setor TEXT,
  image_link TEXT,
  user_id INT,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Create connections table
CREATE TABLE IF NOT EXISTS connections (
  id SERIAL PRIMARY KEY,
  feedback TEXT,
  status BOOLEAN,
  project_id INT,
  user_id INT,
  FOREIGN KEY (project_id) REFERENCES projects(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Create user_connections table
CREATE TABLE IF NOT EXISTS user_connections (
  id SERIAL PRIMARY KEY,
  user_id INT,
  connection_id INT,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (connection_id) REFERENCES connections(id)
);

-- Create ratings table
CREATE TABLE IF NOT EXISTS ratings (
  id SERIAL PRIMARY KEY,
  rating INT,
  user_id INT,
  project_id INT,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (project_id) REFERENCES projects(id)
);

