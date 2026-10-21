<template>
  <div>
    <header class="header-row">
      <h1>AI 工作台</h1>
      <button @click="createStoryJob">生成故事</button>
    </header>

    <div class="grid">
      <section class="panel">
        <h3>故事提示词</h3>
        <textarea v-model="storyPrompt" placeholder="描述世界观、核心冲突和情绪基调" />

        <div class="field-row">
          <label>
            类型
            <select v-model="genre">
              <option value="fantasy">奇幻</option>
              <option value="sci-fi">科幻</option>
              <option value="romance">爱情</option>
              <option value="drama">剧情</option>
            </select>
          </label>

          <label>
            氛围
            <select v-model="tone">
              <option value="epic">史诗</option>
              <option value="mysterious">神秘</option>
              <option value="emotional">感伤</option>
              <option value="cinematic">电影感</option>
            </select>
          </label>
        </div>
      </section>

      <aside class="panel">
        <h3>任务状态</h3>
        <ul class="status-list">
          <li v-for="task in tasks" :key="task.id">
            <span>{{ task.name }}</span>
            <strong :class="task.status">{{ task.status }}</strong>
          </li>
        </ul>
      </aside>
    </div>

    <section class="panel result-panel" v-if="resultText">
      <h3>生成结果</h3>
      <pre>{{ resultText }}</pre>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { generateStory } from '../api/client'

const storyPrompt = ref('一座海底之城，沉睡的神谕，以及只有潮汐涨落时才会被唤醒的力量。')
const genre = ref('fantasy')
const tone = ref('epic')
const resultText = ref('')

const tasks = ref([
  { id: 1, name: '故事大纲', status: 'done' },
  { id: 2, name: '分镜生成', status: 'running' },
  { id: 3, name: '视频生成', status: 'pending' },
])

async function createStoryJob() {
  try {
    const res = await generateStory({
      topic: storyPrompt.value,
      genre: genre.value,
      tone: tone.value,
    })

    const data = res?.data ?? res
    resultText.value = JSON.stringify(data, null, 2)
  } catch (err) {
    resultText.value = `错误: ${err instanceof Error ? err.message : '生成失败'}`
  }
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
