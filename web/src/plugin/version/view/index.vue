<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
        show-overflow-tooltip="true"
      >
        <el-table-column align="left" label="ID" prop="id" width="120" />
        <el-table-column align="left" label="版本号" prop="version" width="120" />
        <el-table-column align="left" label="日志" prop="change_log" max-width="120"  />
        <el-table-column align="left" label="新增日期" width="180">
          <template #default="scope">
            <span>{{ formatDate(scope.row.created_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="160">
          <template #default="scope">
            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" @click="deleteCustomer(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button type="primary" link icon="delete" @click="scope.row.visible = true">删除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="日志">
      <el-form :model="form" label-width="80px">
        <el-form-item label="版本号">
          <el-input v-model="form.version" autocomplete="off" />
        </el-form-item>
        <el-form-item label="更新日志">
          <el-input v-model="form.change_log" autocomplete="off" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>

</template>

<script>
export default {
  name: 'Email1',
}
</script>

<script setup>
import { getVersion,deleteVersion,addVersion } from '@/plugin/version/api/version.js'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/format'
import { reactive, ref } from 'vue'
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const dialogFormVisible = ref(false)
const type = ref('')
const form = ref({
  version: '',
  change_log: ''
})
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}
const closeDialog = () => {
  dialogFormVisible.value = false
  form.value = {
    version: '',
    change_log: ''
  }
}
const enterDialog = async() => {
  let res
  console.log(form.value)
  res = await addVersion(form.value)
  if (table.code === 0) {
    closeDialog()
    getTableData() 
  }
}

const getTableData = async() => {
  const table = await getVersion()
  if (table.code === 0) {
    var temp=JSON.parse(table.data)
    console.log(temp)
    if(temp.code===200){
      tableData.value = temp.data
    }
  }
}

getTableData()

const deleteCustomer = async(row) => {
  console.log(row)
  row.visible = false
  const res = await deleteVersion({ ID: row.id })
  console.log(res)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// const sendTestEmail = async() => {
//   const res = await emailTest()
//   if (res.code === 0) {
//     ElMessage.success('发送成功')
//   }
// }

// const sendEmail = async() => {
//   const res = await emailTest()
//   if (res.code === 0) {
//     ElMessage.success('发送成功,请查收')
//   }
// }
</script>

