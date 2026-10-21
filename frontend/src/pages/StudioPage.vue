<template>
  <div>
    <header class="header-row">
      <h1>写作工作室</h1>
      <button>保存草稿</button>
    </header>

    <div class="studio-grid">
      <section class="panel big">
        <div class="panel-header">
          <span>当前场景</span>
          <button class="small">保存</button>
        </div>
        <h3>场景 08 — 天文台</h3>
        <textarea value="旧天文台仍在守望风暴线之上，熟悉的身影缓缓步入光中。" />
      </section>

      <aside class="panel">
        <h3>场景列表</h3>
        <ul>
          <li>01 — 觉醒</li>
          <li>02 — 信号</li>
          <li>03 — 天文台</li>
        </ul>
      </aside>
    </div>
  </div>
</template>

<style scoped>
.panel {
  background: rgba(15, 23, 42, 0.75);
  border: 1px solid rgba(148, 163, 184, 0.14);
  border-radius: 16px;
  padding: 20px;
}

.header-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 18px;
}

.studio-grid {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 20px;
  margin-top: 20px;
}

.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

textarea {
  width: 100%;
  min-height: 220px;
  border-radius: 12px;
  border: 1px solid rgba(148, 163, 184, 0.2);
  background: rgba(15, 23, 42, 0.9);
  color: #e2e8f0;
  padding: 16px;
  resize: vertical;
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

.small {
  padding: 8px 14px;
}

ul {
  list-style: none;
  padding: 0;
  margin: 10px 0 0;
  display: grid;
  gap: 12px;
}
</style>
