export const API_CONFIG = {
  BASE_URL: 'http://localhost:8080/api/v1',
  TIMEOUT: 5000,
  VERSION: 'v1'
};

export const API_ENDPOINTS = {
  AUTH: {
    LOGIN: '/auth/login',
    REGISTER: '/auth/register',
    ME: '/auth/me'
  },
  EXAM: {
    LIST: '/exams',
    DETAIL: (id: number) => `/exams/${id}`,
    START: (id: number) => `/exams/${id}/start`
  }
} as const;