<template>
  <div>
    <header class="header-row">
      <div>
        <span class="eyebrow">Project detail</span>
        <h1>Moonfall City</h1>
      </div>
      <button>Publish</button>
    </header>

    <section class="summary-grid">
      <article class="panel wide">
        <h3>Story overview</h3>
        <p>
          A suspended city hangs above a drowned world. A guardian with a stolen memory is forced to awaken the hidden tide engine before the sea reaches the last living district.
        </p>
      </article>

      <article class="panel small">
        <h3>Metadata</h3>
        <ul>
          <li><span>Genre</span><strong>Fantasy</strong></li>
          <li><span>Status</span><strong>In progress</strong></li>
          <li><span>Chapters</span><strong>6</strong></li>
          <li><span>AI jobs</span><strong>12</strong></li>
        </ul>
      </article>
    </section>

    <section class="split-grid">
      <article class="panel">
        <h3>Chapter list</h3>
        <div class="chapter-list">
          <div v-for="chapter in chapters" :key="chapter.id" class="chapter-item">
            <div>
              <div class="chapter-title">{{ chapter.title }}</div>
              <div class="chapter-summary">{{ chapter.summary }}</div>
            </div>
            <span class="status-badge" :class="chapter.status">{{ chapter.status }}</span>
          </div>
        </div>
      </article>

      <article class="panel">
        <h3>Current draft</h3>
        <textarea readonly>
The old observatory is still lit. Beneath the tide, the city's last oracle turns toward the west, and for the first time in generations, a human voice answers.
        </textarea>
      </article>
    </section>
  </div>
</template>

<script setup lang="ts">
const chapters = [
  { id: 1, title: 'Chapter 01 — The Drowning Gate', summary: 'The city begins to remember its original purpose.', status: 'ready' },
  { id: 2, title: 'Chapter 02 — The Engine Below', summary: 'A hidden system wakes beneath the sea.', status: 'draft' },
  { id: 3, title: 'Chapter 03 — The Last Oracle', summary: 'The guardian learns the truth behind the flood.', status: 'ready' },
]
</script>

<style scoped>
.header-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.eyebrow {
  display: block;
  color: #93c5fd;
  font-size: 0.72rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  margin-bottom: 8px;
}

h1, h3 {
  margin: 0 0 12px;
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

.summary-grid,
.split-grid {
  display: grid;
  gap: 20px;
  margin-bottom: 20px;
}

.summary-grid {
  grid-template-columns: 2fr 1fr;
}

.split-grid {
  grid-template-columns: 1.1fr 0.9fr;
}

.panel {
  background: rgba(15, 23, 42, 0.75);
  border: 1px solid rgba(148, 163, 184, 0.14);
  border-radius: 16px;
  padding: 20px;
}

ul {
  list-style: none;
  padding: 0;
  margin: 0;
  display: grid;
  gap: 12px;
}

ul li {
  display: flex;
  justify-content: space-between;
  gap: 12px;
  color: #dbeafe;
}

.chapter-list {
  display: grid;
  gap: 14px;
}

.chapter-item {
  display: flex;
  justify-content: space-between;
  gap: 12px;
  padding: 12px 0;
  border-bottom: 1px solid rgba(148, 163, 184, 0.12);
}

.chapter-title {
  font-weight: 600;
  margin-bottom: 6px;
}

.chapter-summary {
  color: #cbd5e1;
  font-size: 0.92rem;
}

.status-badge {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-width: 72px;
  height: 30px;
  border-radius: 999px;
  font-size: 0.72rem;
  font-weight: 700;
  text-transform: uppercase;
}

.status-badge.ready {
  background: rgba(34, 197, 94, 0.12);
  color: #86efac;
}

.status-badge.draft {
  background: rgba(251, 191, 36, 0.12);
  color: #fbbf24;
}

textarea {
  width: 100%;
  min-height: 200px;
  resize: vertical;
  border-radius: 12px;
  border: 1px solid rgba(148, 163, 184, 0.2);
  background: rgba(15, 23, 42, 0.95);
  color: #e2e8f0;
  padding: 16px;
}
</style>
