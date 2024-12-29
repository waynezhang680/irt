<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import WelcomeCard from '../components/home/WelcomeCard.vue';
import ExamStats from '../components/home/ExamStats.vue';
import RecentExams from '../components/home/RecentExams.vue';

const router = useRouter();

const handleStartNewExam = () => {
  router.push('/exams');
};
</script>

<template>
  <div class="home-container">
    <WelcomeCard @start-exam="handleStartNewExam" />
    <div class="stats-container">
      <ExamStats />
    </div>
    <div class="recent-exams">
      <RecentExams />
    </div>
  </div>
</template>

<style scoped>
.home-container {
  padding: 20px;
  display: grid;
  gap: 20px;
  grid-template-columns: 1fr;
}

.stats-container {
  display: grid;
  gap: 20px;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
}

@media (min-width: 768px) {
  .home-container {
    grid-template-columns: 2fr 1fr;
  }
  
  .stats-container {
    grid-column: 1 / -1;
  }
}
</style>