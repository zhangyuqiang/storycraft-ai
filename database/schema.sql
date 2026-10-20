CREATE DATABASE IF NOT EXISTS storycraft_ai CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE storycraft_ai;

CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    nickname VARCHAR(150) NULL,
    avatar_url VARCHAR(500) NULL,
    status ENUM('active', 'disabled', 'deleted') DEFAULT 'active',
    last_login_at DATETIME NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL
);

CREATE TABLE IF NOT EXISTS projects (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    owner_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NULL,
    cover_url VARCHAR(500) NULL,
    genre VARCHAR(100) NULL,
    visibility ENUM('private', 'team', 'public') DEFAULT 'private',
    status ENUM('draft', 'in_progress', 'published') DEFAULT 'draft',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    INDEX idx_projects_owner (owner_id),
    INDEX idx_projects_status (status)
);

CREATE TABLE IF NOT EXISTS project_members (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    role ENUM('owner', 'editor', 'reviewer', 'viewer') DEFAULT 'viewer',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_project_members_project (project_id),
    INDEX idx_project_members_user (user_id),
    UNIQUE KEY uq_project_member (project_id, user_id)
);

CREATE TABLE IF NOT EXISTS chapters (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    summary TEXT NULL,
    sort_index INT NOT NULL DEFAULT 0,
    status ENUM('draft', 'ready', 'published') DEFAULT 'draft',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    INDEX idx_chapters_project (project_id),
    INDEX idx_chapters_sort (project_id, sort_index)
);

CREATE TABLE IF NOT EXISTS scenes (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    chapter_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    scene_number INT NOT NULL DEFAULT 1,
    location VARCHAR(255) NULL,
    mood VARCHAR(100) NULL,
    summary TEXT NULL,
    status ENUM('draft', 'ready', 'published') DEFAULT 'draft',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_scenes_project (project_id),
    INDEX idx_scenes_chapter (chapter_id)
);

CREATE TABLE IF NOT EXISTS characters (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    alias VARCHAR(255) NULL,
    role_type VARCHAR(100) NULL,
    age VARCHAR(50) NULL,
    gender VARCHAR(50) NULL,
    description TEXT NULL,
    personality TEXT NULL,
    portrait_url VARCHAR(500) NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_characters_project (project_id)
);

CREATE TABLE IF NOT EXISTS scene_characters (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    scene_id BIGINT NOT NULL,
    character_id BIGINT NOT NULL,
    role VARCHAR(100) NULL,
    notes TEXT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uq_scene_character (scene_id, character_id)
);

CREATE TABLE IF NOT EXISTS shots (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    chapter_id BIGINT NOT NULL,
    scene_id BIGINT NOT NULL,
    shot_number INT NOT NULL DEFAULT 1,
    title VARCHAR(255) NULL,
    description TEXT NULL,
    camera_info VARCHAR(255) NULL,
    action TEXT NULL,
    dialogue TEXT NULL,
    visual_prompt TEXT NULL,
    source_type ENUM('manual', 'ai') DEFAULT 'manual',
    ai_task_id BIGINT NULL,
    cover_url VARCHAR(500) NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_shots_scene (scene_id),
    INDEX idx_shots_project (project_id)
);

CREATE TABLE IF NOT EXISTS assets (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    owner_id BIGINT NOT NULL,
    asset_type ENUM('image', 'video', 'audio', 'document') NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_url VARCHAR(500) NOT NULL,
    mime_type VARCHAR(200) NULL,
    size_bytes BIGINT NULL,
    source ENUM('upload', 'ai_generated') DEFAULT 'upload',
    related_table VARCHAR(100) NULL,
    related_id BIGINT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_assets_project (project_id),
    INDEX idx_assets_owner (owner_id)
);

CREATE TABLE IF NOT EXISTS ai_tasks (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    task_type ENUM('story_outline', 'scene_outline', 'shot_prompt', 'image', 'video') NOT NULL,
    provider VARCHAR(100) NOT NULL,
    model_name VARCHAR(255) NULL,
    status ENUM('pending', 'processing', 'success', 'failed') DEFAULT 'pending',
    prompt TEXT NOT NULL,
    input_json JSON NULL,
    output_json JSON NULL,
    result_url VARCHAR(500) NULL,
    error_message TEXT NULL,
    progress INT DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    completed_at DATETIME NULL,
    INDEX idx_ai_tasks_project (project_id),
    INDEX idx_ai_tasks_status (status)
);

CREATE TABLE IF NOT EXISTS comments (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    chapter_id BIGINT NULL,
    scene_id BIGINT NULL,
    shot_id BIGINT NULL,
    user_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    parent_id BIGINT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_comments_project (project_id),
    INDEX idx_comments_shot (shot_id)
);

CREATE TABLE IF NOT EXISTS reviews (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    chapter_id BIGINT NULL,
    scene_id BIGINT NULL,
    shot_id BIGINT NULL,
    review_type ENUM('story', 'art', 'final') DEFAULT 'story',
    reviewer_id BIGINT NOT NULL,
    status ENUM('pending', 'approved', 'rejected') DEFAULT 'pending',
    comment TEXT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_reviews_project (project_id)
);

CREATE TABLE IF NOT EXISTS releases (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    chapter_id BIGINT NULL,
    version VARCHAR(50) NOT NULL,
    title VARCHAR(255) NOT NULL,
    summary TEXT NULL,
    status ENUM('draft', 'published') DEFAULT 'draft',
    published_at DATETIME NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_releases_project (project_id)
);

CREATE TABLE IF NOT EXISTS tags (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL UNIQUE,
    category VARCHAR(100) NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS project_tags (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    tag_id BIGINT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uq_project_tag (project_id, tag_id)
);

CREATE TABLE IF NOT EXISTS notifications (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    project_id BIGINT NULL,
    type VARCHAR(50) NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    is_read TINYINT(1) DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_notifications_user (user_id)
);

INSERT INTO tags (name, category) VALUES
('fantasy', 'genre'),
('romance', 'genre'),
('sci-fi', 'genre'),
('drama', 'genre'),
('character-driven', 'style');
