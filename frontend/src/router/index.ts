import { createRouter, createWebHistory } from 'vue-router'
import DashboardPage from '../pages/DashboardPage.vue'
import ProjectsPage from '../pages/ProjectsPage.vue'
import StudioPage from '../pages/StudioPage.vue'
import AIPage from '../pages/AIPage.vue'
import AIWorkspacePage from '../pages/AIWorkspacePage.vue'
import LoginPage from '../pages/LoginPage.vue'

const routes = [
  { path: '/', name: 'dashboard', component: DashboardPage },
  { path: '/projects', name: 'projects', component: ProjectsPage },
  { path: '/studio', name: 'studio', component: StudioPage },
  { path: '/ai', name: 'ai', component: AIPage },
  { path: '/ai/workspace', name: 'ai-workspace', component: AIWorkspacePage },
  { path: '/login', name: 'login', component: LoginPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
