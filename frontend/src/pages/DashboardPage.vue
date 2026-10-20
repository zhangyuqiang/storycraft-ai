<template>
  <div>
    <header class="topbar">
      <h1>Creative dashboard</h1>
      <button>New project</button>
    </header>

    <section class="cards">
      <div class="card primary">
        <span class="eyebrow">Overview</span>
        <h2>Storycraft AI workspace is ready</h2>
        <p>Use this as the creative overview for your story, chapter, character, AI generation, and asset pipeline.</p>
      </div>

      <div class="card">
        <span class="eyebrow">Projects</span>
        <h3>12 active</h3>
        <p>Includes script drafts, storyboard pages, and approved story arcs.</p>
      </div>

      <div class="card">
        <span class="eyebrow">AI generation</span>
        <h3>7 jobs</h3>
        <p>Text-based story prompts and image generation sessions are in progress.</p>
      </div>
    </section>
  </div>
</template>

<style scoped>
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
