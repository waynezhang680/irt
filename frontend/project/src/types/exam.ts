export interface Exam {
  id: number;
  title: string;
  description: string;
  duration: number;
  totalQuestions: number;
  status: 'pending' | 'in_progress' | 'completed';
}

export interface ExamResponse {
  exams: Exam[];
  total: number;
}