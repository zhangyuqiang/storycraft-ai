<template>
  <div>
    <header class="topbar">
      <h1>创意仪表盘</h1>
      <button>新建项目</button>
    </header>

    <section class="cards">
      <div class="card primary">
        <span class="eyebrow">概览</span>
        <h2>StoryCraft AI 工作台已就绪</h2>
        <p>在这里查看故事、章节、角色、AI 生成任务与素材产出流水线。</p>
      </div>

      <div class="card">
        <span class="eyebrow">项目</span>
        <h3>12 个进行中</h3>
        <p>包含剧本草稿、分镜板、已审核故事线与待推进剧情。</p>
      </div>

      <div class="card">
        <span class="eyebrow">AI 生成</span>
        <h3>7 个任务</h3>
        <p>文本生成、图像生成和视频生成任务正在处理或排队中。</p>
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
