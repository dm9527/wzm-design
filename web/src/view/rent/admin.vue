<template>
    <div class="p-4">
        <el-card>
            <div class="text-2xl">系统管理</div>
        </el-card>
        <div class="h-4"></div>
        <el-card>
            <div class="text-2xl mb-8">用户查询</div>
            <!-- form box -->
            <div class="">
                <el-select v-model="formData.select64630" class="full-width-input" clearable>
                    <el-option v-for="(item, index) in select64630Options" :key="index" :label="item.label"
                        :value="item.value" :disabled="item.disabled"></el-option>
                </el-select>
            </div>
            <!-- res box -->
            <div class=""></div>
        </el-card>
    </div>

</template>

<script setup>
import { reactive } from 'vue';
import ResForm from './components/res_Form.vue';
const searchFormData=reactive({
    
})
</script>

<style scoped></style>