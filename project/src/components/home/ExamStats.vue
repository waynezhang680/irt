<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useExamStore } from '../../stores/exam';

const examStore = useExamStore();
const loading = ref(true);

const stats = ref({
  total: 0,
  completed: 0,
  inProgress: 0
});

onMounted(async () => {
  try {
    // 这里可以调用API获取实际的统计数据
    stats.value = {
      total: 10,
      completed: 5,
      inProgress: 2
    };
  } finally {
    loading.value = false;
  }
});
</script>

<template>
  <el-card v-loading="loading">
    <template #header>
      <div class="card-header">
        <h3>考试统计</h3>
      </div>
    </template>
    <div class="stats-grid">
      <div class="stat-item">
        <h4>总考试数</h4>
        <div class="stat-value">{{ stats.total }}</div>
      </div>
      <div class="stat-item">
        <h4>已完成</h4>
        <div class="stat-value">{{ stats.completed }}</div>
      </div>
      <div class="stat-item">
        <h4>进行中</h4>
        <div class="stat-value">{{ stats.inProgress }}</div>
      </div>
    </div>
  </el-card>
</template>

<style scoped>
.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

.stat-item {
  text-align: center;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #409EFF;
}
</style>