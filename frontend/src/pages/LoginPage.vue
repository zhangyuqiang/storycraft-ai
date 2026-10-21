<template>
  <div class="auth-shell">
    <div class="auth-card">
      <h1>登录</h1>
      <input v-model="email" type="email" placeholder="邮箱" />
      <input v-model="password" type="password" placeholder="密码" />
      <button @click="submit">继续</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const email = ref('creator@example.com')
const password = ref('demo123')

function submit() {
  console.log('login submitted', { email: email.value, password: password.value })
}
</script>

<style scoped>
.auth-shell {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 80vh;
}

.auth-card {
  width: min(460px, 100%);
  background: rgba(15, 23, 42, 0.8);
  border: 1px solid rgba(148, 163, 184, 0.14);
  border-radius: 20px;
  padding: 28px;
  display: grid;
  gap: 16px;
}

input {
  width: 100%;
  padding: 12px 14px;
  border-radius: 10px;
  border: 1px solid rgba(148, 163, 184, 0.2);
  background: rgba(15, 23, 42, 0.9);
  color: white;
}

button {
  border: 0;
  padding: 12px 18px;
  border-radius: 10px;
  background: linear-gradient(135deg, #8b5cf6, #2563eb);
  color: white;
  font-weight: 600;
  cursor: pointer;
}
</style>
