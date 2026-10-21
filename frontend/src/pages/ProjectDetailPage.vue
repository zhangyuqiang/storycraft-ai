<template>
  <div>
    <header class="header-row">
      <div>
        <span class="eyebrow">项目详情</span>
        <h1>月落之城</h1>
      </div>
      <button>发布</button>
    </header>

    <section class="summary-grid">
      <article class="panel wide">
        <h3>故事概述</h3>
        <p>
          一座悬浮在海上的城市，守护者藏着失落的记忆，必须在海侵到最后一座居民区之前唤醒被封印的潮汐引擎。
        </p>
      </article>

      <article class="panel small">
        <h3>项目属性</h3>
        <ul>
          <li><span>类型</span><strong>奇幻</strong></li>
          <li><span>状态</span><strong>进行中</strong></li>
          <li><span>章节</span><strong>6</strong></li>
          <li><span>AI 任务</span><strong>12</strong></li>
        </ul>
      </article>
    </section>

    <section class="split-grid">
      <article class="panel">
        <h3>章节列表</h3>
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
        <h3>当前草稿</h3>
        <textarea readonly>
旧天文台仍然亮着灯。潮汐下的城市最后一位预言者转头望向西方，数百年后，人类第一次回应了它的呼唤。
        </textarea>
      </article>
    </section>
  </div>
</template>

<script setup lang="ts">
const chapters = [
  { id: 1, title: '第 01 章 — 淹没之门', summary: '城市开始重新记起自己的原始职责。', status: 'ready' },
  { id: 2, title: '第 02 章 — 海下引擎', summary: '一个隐藏的系统在海底悄然苏醒。', status: 'draft' },
  { id: 3, title: '第 03 章 — 最后的预言者', summary: '守护者终于掌握了洪水背后的真相。', status: 'ready' },
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
