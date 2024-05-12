<template>
    <el-card>
        <el-button @click="routerTo('rentHome')">去首页</el-button>
        <el-button @click="routerTo('rentTenants')">去我的租房</el-button>
        <el-button @click="routerTo('rentLandlord')">去我的出租</el-button>
        <el-button @click="routerTo('rentAdmin')">去系统管理</el-button>
    </el-card>
</template>

<script setup>
import { useRouter } from 'vue-router';
const router = useRouter()
const routerTo = (name) => {
    router.push({ name })
}
</script>

<style scoped></style>