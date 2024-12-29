<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useExamStore } from '../stores/exam';
import { ElMessage } from 'element-plus';

const route = useRoute();
const router = useRouter();
const examStore = useExamStore();
const loading = ref(true);

const examId = Number(route.params.id);

onMounted(async () => {
  try {
    await examStore.getExamDetails(examId);
  } catch (error) {
    ElMessage.error('获取考试详情失败');
    router.push('/exams');
  } finally {
    loading.value = false;
  }
});
</script>

<template>
  <div class="exam-detail-container" v-loading="loading">
    <el-card v-if="examStore.currentExam">
      <template #header>
        <div class="exam-header">
          <h2>{{ examStore.currentExam.title }}</h2>
          <el-tag :type="examStore.currentExam.status === 'completed' ? 'success' : 'primary'">
            {{ examStore.currentExam.status === 'completed' ? '已完成' : '进行中' }}
          </el-tag>
        </div>
      </template>
      <div class="exam-info">
        <p>{{ examStore.currentExam.description }}</p>
        <div class="exam-meta">
          <div class="meta-item">
            <span class="label">考试时长：</span>
            <span>{{ examStore.currentExam.duration }} 分钟</span>
          </div>
          <div class="meta-item">
            <span class="label">题目数量：</span>
            <span>{{ examStore.currentExam.totalQuestions }}</span>
          </div>
        </div>
      </div>
    </el-card>
  </div>
</template>

<style scoped>
.exam-detail-container {
  padding: 20px;
}

.exam-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.exam-info {
  padding: 20px 0;
}

.exam-meta {
  margin-top: 20px;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.label {
  font-weight: bold;
  color: #666;
}
</style>