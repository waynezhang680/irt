<script setup lang="ts">
import { useUserStore } from './stores/user';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';

const userStore = useUserStore();
const router = useRouter();

const handleLogout = () => {
  userStore.logout();
  ElMessage.success('已退出登录');
  router.push('/login');
};
</script>

<template>
  <el-container v-if="userStore.token">
    <el-header>
      <div class="header-content">
        <h1>考试系统</h1>
        <div class="user-info">
          <span v-if="userStore.user">{{ userStore.user.username }}</span>
          <el-button type="text" @click="handleLogout">退出登录</el-button>
        </div>
      </div>
    </el-header>
    <el-main>
      <router-view />
    </el-main>
  </el-container>
  <router-view v-else />
</template>

<style scoped>
.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100%;
  padding: 0 20px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.user-info {
  display: flex;
  align-items: center;
  gap: 16px;
}

.el-main {
  background-color: #f5f7fa;
  min-height: calc(100vh - 60px);
}
</style>