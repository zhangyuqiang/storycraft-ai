<template>
  <div>
    <header class="header-row">
      <h1>AI Workspace</h1>
      <button @click="createStoryJob">Generate story</button>
    </header>

    <div class="grid">
      <section class="panel">
        <h3>Story prompt</h3>
        <textarea v-model="storyPrompt" placeholder="Describe the world, core conflict, genre, and emotional tone." />

        <div class="field-row">
          <label>
            Genre
            <select v-model="genre">
              <option value="fantasy">Fantasy</option>
              <option value="sci-fi">Sci-fi</option>
              <option value="romance">Romance</option>
              <option value="drama">Drama</option>
            </select>
          </label>

          <label>
            Tone
            <select v-model="tone">
              <option value="epic">Epic</option>
              <option value="mysterious">Mysterious</option>
              <option value="emotional">Emotional</option>
              <option value="cinematic">Cinematic</option>
            </select>
          </label>
        </div>
      </section>

      <aside class="panel">
        <h3>Task status</h3>
        <ul class="status-list">
          <li v-for="task in tasks" :key="task.id">
            <span>{{ task.name }}</span>
            <strong :class="task.status">{{ task.status }}</strong>
          </li>
        </ul>
      </aside>
    </div>

    <section class="panel result-panel" v-if="resultText">
      <h3>Result</h3>
      <pre>{{ resultText }}</pre>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const storyPrompt = ref('A city beneath the sea, a hidden oracle, and a power that awakens only when the tide turns.')
const genre = ref('fantasy')
const tone = ref('epic')
const resultText = ref('')

const tasks = ref([
  { id: 1, name: 'Story outline', status: 'done' },
  { id: 2, name: 'Shot generation', status: 'running' },
  { id: 3, name: 'Video generation', status: 'pending' },
])

function createStoryJob() {
  resultText.value = `Story outline generated:\n\nTitle: The City Beneath the Tide\nGenre: ${genre.value}\nTone: ${tone.value}\n\nSynopsis: A forgotten oracle wakes beneath the sea and points the realm toward an impossible truth.\nChapters: [The Drowning Gate, The Signal Beneath, The Last Tide, The Awakening].`
}
</script>

<style scoped>
.header-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

h1, h3 {
  margin: 0 0 12px 0;
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

.grid {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 20px;
}

.panel {
  background: rgba(15, 23, 42, 0.75);
  border: 1px solid rgba(148, 163, 184, 0.14);
  border-radius: 16px;
  padding: 20px;
}

textarea,
select {
  width: 100%;
  border-radius: 12px;
  border: 1px solid rgba(148, 163, 184, 0.2);
  background: rgba(15, 23, 42, 0.95);
  color: #e2e8f0;
  padding: 12px 14px;
}

textarea {
  min-height: 180px;
  resize: vertical;
  margin-bottom: 16px;
}

.field-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

label {
  display: grid;
  gap: 8px;
  color: #cbd5e1;
}

.status-list {
  list-style: none;
  padding: 0;
  margin: 12px 0 0;
  display: grid;
  gap: 12px;
}

.status-list li {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
}

.done { color: #86efac; }
.running { color: #fbbf24; }
.pending { color: #93c5fd; }

.result-panel {
  margin-top: 20px;
}

pre {
  white-space: pre-wrap;
  word-break: break-word;
  margin: 0;
  color: #dbeafe;
}
</style>
