<template>
  <div class="app-shell">
    <aside class="sidebar">
      <div class="brand">StoryCraft AI</div>
      <nav>
        <RouterLink to="/">Dashboard</RouterLink>
        <RouterLink to="/projects">Projects</RouterLink>
        <RouterLink to="/studio">Writer Studio</RouterLink>
        <RouterLink to="/ai">AI Assistant</RouterLink>
        <RouterLink to="/ai/workspace">AI Workspace</RouterLink>
        <RouterLink to="/login">Login</RouterLink>
      </nav>
    </aside>

    <main class="content">
      <RouterView />
    </main>
  </div>
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<style scoped>
:global(html, body, #app) {
  margin: 0;
  width: 100%;
  min-height: 100%;
  height: 100%;
  font-family: Inter, 'Segoe UI', sans-serif;
  background: #07192d;
  color: #e2e8f0;
}

:global(body) {
  min-height: 100vh;
}

* { box-sizing: border-box; }

.app-shell {
  display: flex;
  min-height: 100vh;
  width: 100%;
  background: #07192d;
}

.sidebar {
  width: 280px;
  min-width: 280px;
  background: linear-gradient(180deg, #07192d, #091d34);
  border-right: 1px solid rgba(148, 163, 184, 0.18);
  padding: 24px 20px 32px;
}

.brand {
  font-size: 1.9rem;
  font-weight: 700;
  margin-bottom: 28px;
  letter-spacing: -0.04em;
  color: #f8fafc;
}

nav {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

nav a {
  color: #dfeaf8;
  text-decoration: none;
  padding: 12px 14px;
  border-radius: 12px;
  background: rgba(148, 163, 184, 0.06);
  border: 1px solid transparent;
  transition: all 0.2s ease;
  font-size: 1.08rem;
}

nav a:hover,
nav a.router-link-active {
  background: rgba(37, 99, 235, 0.08);
  border-color: rgba(96, 165, 250, 0.35);
  color: #e5f3ff;
}

.content {
  flex: 1;
  min-width: 0;
  padding: 24px 28px 32px;
  background: #07192d;
}
</style>
