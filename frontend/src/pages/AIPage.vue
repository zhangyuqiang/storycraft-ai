<template>
  <div>
    <header class="header-row">
      <h1>AI 助手</h1>
      <button @click="generateStory">生成故事</button>
    </header>

    <div class="ai-grid">
      <section class="panel">
        <h3>生成故事大纲</h3>
        <textarea v-model="prompt" placeholder="描述故事主题、类型和情绪基调" />
        <button @click="generateStory" class="secondary">生成</button>

        <div v-if="result" class="result-box">
          <h4>生成结果</h4>
          <pre>{{ result }}</pre>
        </div>
      </section>

      <aside class="panel">
        <h3>提示词模板</h3>
        <ul>
          <li>奇幻爱情</li>
          <li>末日戏剧</li>
          <li>都市悬疑</li>
          <li>反乌托邦科幻</li>
        </ul>
      </aside>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const prompt = ref('一位年轻建筑师发现海底之下沉睡着一座失落城市，他必须在爱与职责之间做出抉择。')
const result = ref('')

async function generateStory() {
  result.value = `故事大纲已生成：\n\n- 标题：深海记忆\n- 类型：奇幻剧情\n- 主线：失落城市、隐藏身份与最终选择。`
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

.secondary {
  margin-top: 8px;
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
