<template>
  <div>
    <header class="header-row">
      <h1>项目</h1>
      <button @click="openCreateModal">+ 新建项目</button>
    </header>

    <div v-if="loading" class="empty">加载中...</div>
    <div v-else-if="error" class="empty error">{{ error }}</div>

    <div v-else class="list">
      <div v-for="project in projects" :key="project.id" class="project-card">
        <div>
          <div class="tag">{{ project.genre }}</div>
          <h3>{{ project.title }}</h3>
          <p>{{ project.summary || project.description || '暂无描述' }}</p>
        </div>
        <div class="meta">
          <span>{{ project.status }}</span>
          <span>{{ project.chapters || 0 }} 章</span>
          <div class="actions">
            <RouterLink :to="`/projects/${project.id}`" class="detail-link">打开</RouterLink>
            <button class="small-button" @click="openEditModal(project)">编辑</button>
          </div>
        </div>
      </div>
    </div>

    <div v-if="showModal" class="modal-backdrop" @click.self="closeModal">
      <div class="modal">
        <h3>{{ editingProject ? '编辑项目' : '新建项目' }}</h3>

        <label>
          项目名称
          <input v-model="form.title" placeholder="请输入项目名称" />
        </label>

        <label>
          描述
          <textarea v-model="form.description" placeholder="请输入项目简介" />
        </label>

        <label>
          类型
          <select v-model="form.genre">
            <option value="fantasy">奇幻</option>
            <option value="sci-fi">科幻</option>
            <option value="romance">爱情</option>
            <option value="drama">剧情</option>
          </select>
        </label>

        <div class="modal-actions">
          <button class="secondary" @click="closeModal">取消</button>
          <button @click="submitProject" :disabled="!form.title.trim()">
            {{ editingProject ? '保存修改' : '创建项目' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { RouterLink } from 'vue-router'
import { createProject, fetchProjects, updateProject, type Project } from '../api/client'

const projects = ref<Project[]>([])
const loading = ref(true)
const error = ref('')
const showModal = ref(false)
const editingProject = ref<Project | null>(null)
const form = ref({ title: '', description: '', genre: 'fantasy' })

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

function openCreateModal() {
  editingProject.value = null
  form.value = { title: '', description: '', genre: 'fantasy' }
  showModal.value = true
}

function openEditModal(project: Project) {
  editingProject.value = project
  form.value = {
    title: project.title,
    description: project.description || '',
    genre: project.genre || 'fantasy',
  }
  showModal.value = true
}

function closeModal() {
  showModal.value = false
  editingProject.value = null
  form.value = { title: '', description: '', genre: 'fantasy' }
}

async function submitProject() {
  if (!form.value.title.trim()) {
    error.value = '项目名称不能为空'
    return
  }

  try {
    if (editingProject.value) {
      await updateProject(editingProject.value.id, {
        title: form.value.title.trim(),
        description: form.value.description,
        genre: form.value.genre,
      })
    } else {
      await createProject({
        title: form.value.title.trim(),
        description: form.value.description,
        genre: form.value.genre,
        visibility: 'private',
      })
    }

    closeModal()
    await loadProjects()
  } catch (err) {
    error.value = err instanceof Error ? err.message : '保存失败'
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

button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
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

.actions {
  display: flex;
  align-items: center;
  gap: 10px;
}

.detail-link {
  color: #93c5fd;
  text-decoration: none;
  font-weight: 600;
}

.small-button {
  padding: 6px 10px;
  font-size: 0.8rem;
  border-radius: 8px;
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

.modal-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(2, 6, 23, 0.65);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 50;
}

.modal {
  width: min(520px, 90vw);
  background: #0f172a;
  border: 1px solid rgba(148, 163, 184, 0.2);
  border-radius: 18px;
  padding: 24px;
  display: grid;
  gap: 16px;
}

label {
  display: grid;
  gap: 8px;
  color: #dbeafe;
}

input,
textarea,
select {
  width: 100%;
  border-radius: 10px;
  border: 1px solid rgba(148, 163, 184, 0.2);
  background: rgba(15, 23, 42, 0.9);
  color: #f8fafc;
  padding: 10px 12px;
}

textarea {
  min-height: 90px;
  resize: vertical;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

.secondary {
  background: rgba(148, 163, 184, 0.12);
  color: #e2e8f0;
}
</style>
