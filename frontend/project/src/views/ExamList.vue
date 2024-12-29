<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { examApi } from '../api/exam';
import type { Exam } from '../types/exam';
import { ElMessage } from 'element-plus';

const router = useRouter();
const exams = ref<Exam[]>([]);
const loading = ref(false);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(10);

const loadExams = async () => {
  try {
    loading.value = true;
    const response = await examApi.getExams(currentPage.value, pageSize.value);
    exams.value = response.data.exams;
    total.value = response.data.total;
  } catch (error) {
    ElMessage.error('获取考试列表失败');
  } finally {
    loading.value = false;
  }
};

const handleStartExam = async (examId: number) => {
  try {
    await examApi.startExam(examId);
    router.push(`/exam/${examId}`);
  } catch (error) {
    ElMessage.error('启动考试失败');
  }
};

const handlePageChange = (page: number) => {
  currentPage.value = page;
  loadExams();
};

onMounted(() => {
  loadExams();
});
</script>

<template>
  <div class="exam-list-container">
    <h2>考试列表</h2>
    <el-table
      v-loading="loading"
      :data="exams"
      style="width: 100%"
    >
      <el-table-column prop="title" label="考试名称" />
      <el-table-column prop="description" label="描述" />
      <el-table-column prop="duration" label="时长(分钟)" width="120" />
      <el-table-column prop="totalQuestions" label="题目数量" width="120" />
      <el-table-column prop="status" label="状态" width="120">
        <template #default="{ row }">
          <el-tag :type="row.status === 'completed' ? 'success' : 'info'">
            {{ row.status === 'completed' ? '已完成' : '未开始' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="120">
        <template #default="{ row }">
          <el-button
            type="primary"
            size="small"
            :disabled="row.status === 'completed'"
            @click="handleStartExam(row.id)"
          >
            开始考试
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <div class="pagination">
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :total="total"
        :page-sizes="[10, 20, 30, 50]"
        layout="total, sizes, prev, pager, next"
        @size-change="loadExams"
        @current-change="handlePageChange"
      />
    </div>
  </div>
</template>

<style scoped>
.exam-list-container {
  padding: 20px;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>