<template>
  <div>
    <h1>AI Assistant</h1>

    <div class="ai-grid">
      <section class="panel">
        <h3>Generate story outline</h3>
        <textarea v-model="prompt" placeholder="Describe the story theme, genre, and key emotional tone" />
        <button @click="generateStory">Generate</button>

        <div v-if="result" class="result-box">
          <h4>Result</h4>
          <pre>{{ result }}</pre>
        </div>
      </section>

      <aside class="panel">
        <h3>Prompt templates</h3>
        <ul>
          <li>Fantasy romance</li>
          <li>Post-apocalypse drama</li>
          <li>Urban mystery</li>
          <li>Counter-culture sci-fi</li>
        </ul>
      </aside>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const prompt = ref('A young architect discovers a city beneath the sea and has to choose between love and duty.')
const result = ref('')

async function generateStory() {
  result.value = `Story outline generated from prompt: ${prompt.value}\n\n- Title: Deep Sea Memory\n- Genre: Fantasy Drama\n- Arc: A vanished city, a hidden identity, and a final choice.`
}
</script>

<style scoped>
.ai-grid {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 20px;
  margin-top: 20px;
}

.panel {
  background: rgba(15, 23, 42, 0.75);
  border: 1px solid rgba(148, 163, 184, 0.14);
  border-radius: 16px;
  padding: 20px;
}

textarea {
  width: 100%;
  min-height: 180px;
  border-radius: 12px;
  border: 1px solid rgba(148, 163, 184, 0.2);
  background: rgba(15, 23, 42, 0.9);
  color: #e2e8f0;
  padding: 16px;
  resize: vertical;
  margin: 12px 0;
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

ul {
  list-style: none;
  padding: 0;
  margin: 10px 0 0;
  display: grid;
  gap: 12px;
}

.result-box {
  margin-top: 18px;
  background: rgba(12, 18, 36, 0.8);
  border-radius: 12px;
  padding: 16px;
}

pre {
  white-space: pre-wrap;
  word-break: break-word;
  margin: 0;
  color: #dbeafe;
}
</style>
