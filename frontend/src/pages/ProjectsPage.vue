<template>
  <div>
    <header class="header-row">
      <h1>Projects</h1>
      <button>+ Create project</button>
    </header>

    <div class="list">
      <div v-for="project in projects" :key="project.id" class="project-card">
        <div>
          <div class="tag">{{ project.genre }}</div>
          <h3>{{ project.title }}</h3>
          <p>{{ project.summary }}</p>
        </div>
        <div class="meta">
          <span>{{ project.status }}</span>
          <span>{{ project.chapters }} chapters</span>
          <RouterLink :to="`/projects/${project.id}`" class="detail-link">Open</RouterLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'

const projects = [
  { id: 1, title: 'Moonfall City', genre: 'Fantasy', summary: 'An exiled guardian is drawn into a forgotten eclipse prophecy.', status: 'In progress', chapters: 6 },
  { id: 2, title: 'Glass Horizon', genre: 'Sci-fi', summary: 'A drifting deep-space pilot returns to a collapsing colony world.', status: 'Draft', chapters: 3 },
]
</script>

<style scoped>
.header-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

h1 { margin: 0; }

button {
  border: 0;
  padding: 10px 18px;
  border-radius: 10px;
  background: linear-gradient(135deg, #8b5cf6, #2563eb);
  color: white;
  font-weight: 600;
  cursor: pointer;
}

.list {
  display: grid;
  gap: 16px;
}

.project-card {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  padding: 20px;
  border-radius: 16px;
  background: rgba(15, 23, 42, 0.75);
  border: 1px solid rgba(148, 163, 184, 0.14);
}

.tag {
  display: inline-block;
  background: rgba(96, 165, 250, 0.15);
  color: #bfdbfe;
  border-radius: 999px;
  padding: 6px 10px;
  font-size: 0.72rem;
  margin-bottom: 8px;
}

p {
  margin: 0;
  color: #cbd5e1;
}

.meta {
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 8px;
  color: #cbd5e1;
}

.detail-link {
  color: #93c5fd;
  text-decoration: none;
  font-weight: 600;
}
</style>
