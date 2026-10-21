<template>
  <div>
    <header class="header-row">
      <h1>Projects</h1>
      <button @click="createNewProject">+ Create project</button>
    </header>

    <div v-if="loading" class="empty">Loading projects...</div>
    <div v-else-if="error" class="empty error">{{ error }}</div>

    <div v-else class="list">
      <div v-for="project in projects" :key="project.id" class="project-card">
        <div>
          <div class="tag">{{ project.genre }}</div>
          <h3>{{ project.title }}</h3>
          <p>{{ project.summary || project.description || 'No description yet.' }}</p>
        </div>
        <div class="meta">
          <span>{{ project.status }}</span>
          <span>{{ project.chapters || 0 }} chapters</span>
          <RouterLink :to="`/projects/${project.id}`" class="detail-link">Open</RouterLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { RouterLink } from 'vue-router'
import { createProject, fetchProjects, type Project } from '../api/client'

const projects = ref<Project[]>([])
const loading = ref(true)
const error = ref('')

async function loadProjects() {
  try {
    loading.value = true
    error.value = ''
    projects.value = await fetchProjects()
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Unknown error'
  } finally {
    loading.value = false
  }
}

async function createNewProject() {
  try {
    await createProject({
      title: 'New Project',
      description: 'Created from the frontend workflow',
      genre: 'fantasy',
      visibility: 'private',
    })
    await loadProjects()
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Project creation failed'
  }
}

onMounted(() => {
  loadProjects()
})
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

.empty {
  padding: 18px;
  color: #cbd5e1;
  background: rgba(15, 23, 42, 0.6);
  border-radius: 12px;
}

.error {
  color: #fca5a5;
}
</style>
