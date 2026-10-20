<template>
  <div class="app-shell">
    <aside class="sidebar">
      <div class="brand">StoryCraft AI</div>
      <nav>
        <a>Dashboard</a>
        <a>Projects</a>
        <a>Writer Studio</a>
        <a>Storyboard</a>
        <a>AI Assistant</a>
        <a>Assets</a>
      </nav>
    </aside>

    <main class="content">
      <header class="topbar">
        <h1>Creative workspace</h1>
        <button>New project</button>
      </header>

      <section class="cards">
        <div class="card primary">
          <span class="eyebrow">Status</span>
          <h2>Project ready</h2>
          <p>Use this base app as the foundation for chapters, scenes, AI generation, and storyboard modeling.</p>
        </div>

        <div class="card">
          <span class="eyebrow">Backend</span>
          <h3>Health check</h3>
          <p><strong>OK</strong> - /api/v1/health</p>
        </div>

        <div class="card">
          <span class="eyebrow">Frontend</span>
          <h3>Vue 3</h3>
          <p>Ready for route-driven project pages and AI tools.</p>
        </div>
      </section>
    </main>
  </div>
</template>

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

.topbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 32px;
}

.topbar h1 {
  margin: 0;
  font-size: 2rem;
}

button {
  border: 0;
  padding: 10px 18px;
  border-radius: 10px;
  background: linear-gradient(135deg, #8b5cf6, #2563eb);
  color: white;
  font-weight: 600;
  cursor: pointer;
}

.cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 20px;
}

.card {
  padding: 22px;
  border-radius: 16px;
  background: rgba(15, 23, 42, 0.75);
  border: 1px solid rgba(148, 163, 184, 0.14);
  min-height: 180px;
}

.card.primary {
  background: linear-gradient(135deg, rgba(59,130,246,0.25), rgba(139,92,246,0.18));
}

.eyebrow {
  display: block;
  font-size: 0.72rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: #93c5fd;
  margin-bottom: 10px;
}

h2, h3 {
  margin: 0 0 10px;
}

p {
  margin: 0;
  color: #cbd5e1;
  line-height: 1.6;
}
</style>
