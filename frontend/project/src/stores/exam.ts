import { defineStore } from 'pinia';
import { ref } from 'vue';
import type { Exam } from '../types/exam';
import { examApi } from '../api/exam';

export const useExamStore = defineStore('exam', () => {
  const currentExam = ref<Exam | null>(null);
  const examList = ref<Exam[]>([]);
  const loading = ref(false);

  const fetchExams = async (page = 1, limit = 10) => {
    loading.value = true;
    try {
      const response = await examApi.getExams(page, limit);
      examList.value = response.data.exams;
      return response.data;
    } finally {
      loading.value = false;
    }
  };

  const startExam = async (examId: number) => {
    loading.value = true;
    try {
      const response = await examApi.startExam(examId);
      currentExam.value = response.data;
      return response.data;
    } finally {
      loading.value = false;
    }
  };

  const getExamDetails = async (examId: number) => {
    loading.value = true;
    try {
      const response = await examApi.getExamDetails(examId);
      currentExam.value = response.data;
      return response.data;
    } finally {
      loading.value = false;
    }
  };

  return {
    currentExam,
    examList,
    loading,
    fetchExams,
    startExam,
    getExamDetails
  };
});