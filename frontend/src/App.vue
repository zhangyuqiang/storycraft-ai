<template>
  <div class="app-shell">
    <aside class="sidebar">
      <div class="brand">StoryCraft AI</div>
      <nav>
        <RouterLink to="/">Dashboard</RouterLink>
        <RouterLink to="/projects">Projects</RouterLink>
        <RouterLink to="/studio">Writer Studio</RouterLink>
        <RouterLink to="/ai">AI Assistant</RouterLink>
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
:global(body) {
  margin: 0;
  font-family: Inter, 'Segoe UI', sans-serif;
  background: #0f172a;
  color: #e2e8f0;
}

* {
  box-sizing: border-box;
}

.app-shell {
  display: flex;
  min-height: 100vh;
}

.sidebar {
  width: 260px;
  background: linear-gradient(180deg, #111827, #0f172a);
  border-right: 1px solid rgba(148, 163, 184, 0.2);
  padding: 24px 20px;
}

.brand {
  font-size: 1.4rem;
  font-weight: 700;
  margin-bottom: 32px;
}

nav {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

nav a {
  color: #cbd5e1;
  text-decoration: none;
  padding: 10px 12px;
  border-radius: 10px;
  background: rgba(148, 163, 184, 0.06);
}

.content {
  flex: 1;
  padding: 32px;
}
</style>
