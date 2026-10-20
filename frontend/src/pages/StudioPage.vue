<template>
  <div>
    <h1>Writer Studio</h1>

    <div class="studio-grid">
      <section class="panel big">
        <div class="panel-header">
          <span>Current Scene</span>
          <button>Save draft</button>
        </div>
        <h3>Scene 08 — The Observatory</h3>
        <textarea value="The old observatory watches the stormline above the dead sea. A familiar silhouette steps into the light." />
      </section>

      <aside class="panel">
        <h3>Scene list</h3>
        <ul>
          <li>01 — Awakening</li>
          <li>02 — The Signal</li>
          <li>03 — The Observatory</li>
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

ul {
  list-style: none;
  padding: 0;
  margin: 10px 0 0;
  display: grid;
  gap: 12px;
}
</style>
