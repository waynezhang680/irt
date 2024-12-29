<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import type { Exam } from '../../types/exam';
import { examApi } from '../../api/exam';

const router = useRouter();
const loading = ref(true);
const recentExams = ref<Exam[]>([]);

const loadRecentExams = async () => {
  try {
    const response = await examApi.getExams(1, 5);
    recentExams.value = response.data.exams;
  } finally {
    loading.value = false;
  }
};

const handleExamClick = (examId: number) => {
  router.push(`/exam/${examId}`);
};

onMounted(() => {
  loadRecentExams();
});
</script>

<template>
  <el-card v-loading="loading">
    <template #header>
      <div class="card-header">
        <h3>最近考试</h3>
      </div>
    </template>
    <div class="recent-exams-list">
      <el-empty v-if="recentExams.length === 0" description="暂无考试记录" />
      <el-timeline v-else>
        <el-timeline-item
          v-for="exam in recentExams"
          :key="exam.id"
          :type="exam.status === 'completed' ? 'success' : 'primary'"
        >
          <div
            class="exam-item"
            @click="handleExamClick(exam.id)"
          >
            <h4>{{ exam.title }}</h4>
            <p>{{ exam.description }}</p>
            <el-tag :type="exam.status === 'completed' ? 'success' : 'info'">
              {{ exam.status === 'completed' ? '已完成' : '进行中' }}
            </el-tag>
          </div>
        </el-timeline-item>
      </el-timeline>
    </div>
  </el-card>
</template>

<style scoped>
.recent-exams-list {
  max-height: 400px;
  overflow-y: auto;
}

.exam-item {
  cursor: pointer;
  padding: 8px;
  border-radius: 4px;
  transition: background-color 0.2s;
}

.exam-item:hover {
  background-color: #f5f7fa;
}

.exam-item h4 {
  margin: 0 0 8px 0;
}

.exam-item p {
  margin: 0 0 8px 0;
  color: #666;
  font-size: 14px;
}
</style>