<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
       </el-form-item>
        <el-form-item label="资源类型:" prop="type">
           <el-select v-model="formData.type" placeholder="请选择资源类型" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in house_typeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="房东id:" prop="landlord_id">
          <el-input v-model.number="formData.landlord_id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="资源位置:" prop="location">
          <el-input v-model="formData.location" :clearable="true"  placeholder="请输入资源位置" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createResource,
  updateResource,
  findResource
} from '@/api/rent/resource'

defineOptions({
    name: 'ResourceForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const house_typeOptions = ref([])
const formData = ref({
            name: '',
            type: '',
            landlord_id: 0,
            location: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               landlord_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               location : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findResource({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reresource
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    house_typeOptions.value = await getDictFunc('house_type')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createResource(formData.value)
               break
             case 'update':
               res = await updateResource(formData.value)
               break
             default:
               res = await createResource(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
