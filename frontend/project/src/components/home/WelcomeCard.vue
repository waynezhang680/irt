<script setup lang="ts">
import { useUserStore } from '../../stores/user';

const userStore = useUserStore();

defineEmits<{
  (e: 'startExam'): void
}>();
</script>

<template>
  <el-card class="welcome-card">
    <template #header>
      <div class="welcome-header">
        <h2>欢迎回来, {{ userStore.user?.username }}</h2>
      </div>
    </template>
    <div class="welcome-content">
      <p>开始新的考试或继续未完成的考试。</p>
      <el-button type="primary" @click="$emit('startExam')">
        开始新考试
      </el-button>
    </div>
  </el-card>
</template>

<style scoped>
.welcome-card {
  margin-bottom: 20px;
}

.welcome-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.welcome-content {
  text-align: center;
  padding: 20px 0;
}
</style>