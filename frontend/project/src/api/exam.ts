import api from './axios';
import { API_ENDPOINTS } from '../config/api';
import type { Exam, ExamResponse } from '../types/exam';

export const examApi = {
  getExams: (page = 1, limit = 10) =>
    api.get<ExamResponse>(API_ENDPOINTS.EXAM.LIST, { params: { page, limit } }),
    
  startExam: (examId: number) =>
    api.post<Exam>(API_ENDPOINTS.EXAM.START(examId)),
    
  getExamDetails: (examId: number) =>
    api.get<Exam>(API_ENDPOINTS.EXAM.DETAIL(examId))
};