<template>
  <div>
    <h3>用户列表页面</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search v-model="queryParam.username"
                          placeholder="输入用户名查找"
                          enter-button
                          allowClear
                          @search="getUserList" />
        </a-col>
        <a-col :span="4">
          <a-button type="primary"
                    @click="addUserVisible=true">新增</a-button>
        </a-col>
      </a-row>
      <a-table rowKey="username"
               :columns="columns"
               :pagination="pagination"
               :dataSource="userlist"
               bordered
               @change="handleTableChange">
        <span slot="role"
              slot-scope="role">{{ role === 1 ? '管理员' : '订阅者' }}</span>
        <template slot="action"
                  slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary"
                      icon='edit'
                      style="margin-right:15px"
                      @click="editUser(data.ID)">编辑</a-button>
            <a-button type="danger"
                      icon="delete"
                      style="margin-right:15px"
                      @click="deleteUser(data.ID)">删除</a-button>
            <a-button type="danger"
                      ghost
                      icon="warning"
                      style="margin-right:15px"
                      @click="updatePass(data.ID)">重置密码</a-button>
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 新增用户 -->
    <a-modal closable
             title="新增用户"
             :visible="addUserVisible"
             width="60%"
             @ok="addUserOk"
             @cancel="addUserCancel">
      <!-- :destroyOnClose="true"> -->
      <a-form-model :model="userInfo"
                    :rules="userRules"
                    ref="addUserRef">
        <a-form-model-item label="用户名"
                           prop="username">
          <a-input v-model="userInfo.username">
          </a-input>
        </a-form-model-item>
        <a-form-model-item has-feedback
                           label="密码"
                           prop="password">
          <a-input-password v-model="userInfo.password"></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback
                           label="确认密码"
                           prop="checkpass">
          <a-input-password v-model="userInfo.checkpass"></a-input-password>
        </a-form-model-item>
        <a-form-model-item label="是否为管理员"
                           prop="role">
          <a-select defaultValue="2"
                    style="120px"
                    @change="adminChange">
            <a-select-option key="1"
                             value="1">是</a-select-option>
            <a-select-option key="2"
                             value="2">否</a-select-option>
          </a-select>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 编辑用户 -->
    <a-modal closable
             title="编辑用户"
             :visible="editUserVisible"
             width="60%"
             @ok="editUserOk"
             @cancel="editUserCancel">
      <!-- :destroyOnClose="true"> -->
      <a-form-model :model="userInfo"
                    :rules="userRules"
                    ref="editUserRef">
        <a-form-model-item label="用户名"
                           prop="username">
          <a-input v-model="userInfo.username">
          </a-input>
        </a-form-model-item>
        <a-form-model-item label="是否为管理员"
                           prop="role">
          <a-select defaultValue="2"
                    style="120px"
                    @change="adminChange">
            <a-select-option key="1"
                             value="1">是</a-select-option>
            <a-select-option key="2"
                             value="2">否</a-select-option>
          </a-select>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id'
  },
  {
    title: '用户名',
    dataIndex: 'username',
    width: '20%',
    key: 'username',
    align: 'center'
  },
  {
    title: '角色',
    dataIndex: 'role',
    width: '20%',
    key: 'role',
    align: 'center',
    scopedSlots: { customRender: 'role' }
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' }
  }
]
export default {
  data () {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`
      },
      userlist: [],
      columns,
      queryParam: {
        username: '',
        pagesize: 5,
        pagenum: 1
      },
      userInfo: {
        id: 0,
        username: '',
        password: '',
        checkpass: '',
        role: 2
      },
      visible: false,
      addUserVisible: false,
      editUserVisible: false,
      userRules: {
        username: [{
          validator: (rule, value, callback) => {
            if (this.userInfo.username === '') {
              callback(new Error('请输入用户名'))
            } else if (this.userInfo.username.length < 4 || this.userInfo.username.length > 12) {
              // } else if ([...this.userInfo.username].length < 4 || [...this.userInfo.username].length > 12) {
              callback(new Error('用户名须在4-12位之间'))
            } else {
              callback()
            }
          },
          trigger: 'blur'
        }
        ],
        password: [{
          validator: (rule, value, callback) => {
            if (this.userInfo.password === '') {
              callback(new Error('请输入密码'))
              // } else if ([...this.userInfo.password].length < 6 || [...this.userInfo.password].length > 20) {
            } else if (this.userInfo.password.length < 6 || this.userInfo.password.length > 20) {
              callback(new Error('密码须在6-20位之间'))
            } else {
              callback()
            }
          },
          trigger: 'blur'
        }
        ],
        checkpass: [{
          validator: (rule, value, callback) => {
            if (this.userInfo.checkpass === '') {
              callback(new Error('请输入密码'))
            } else if (this.userInfo.password !== this.userInfo.checkpass) {
              callback(new Error('密码不一致'))
            } else {
              callback()
            }
          },
          trigger: 'blur'
        }
        ]
      }
    }
  },
  created () {
    this.getUserList()
  },
  methods: {
    // 获取用户列表
    async getUserList () {
      const { data: res } = await this.$http.get('users', {
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 0 && res.status !== 200) return this.$message.error(res.message)
      this.userlist = res.data
      this.pagination.total = res.total
    },
    handleTableChange (pagination) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagienum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getUserList()
    },
    // 删除用户
    deleteUser (id) {
      this.$confirm({
        title: '确定要删除该用户吗?',
        content: '一旦删除，无法恢复',
        onOk: async () => {
          const res = await this.$http.delete(`user/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getUserList()
        },
        onCancel: () => {
          this.$message.info('已取消删除操作')
        }
      })
    },
    // 添加用户
    addUserOk () {
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('输入的参数不合法')
        const { data: res } = await this.$http.post('user/add', {
          username: this.userInfo.username,
          password: this.userInfo.password,
          role: this.userInfo.role
        })
        if (res.status !== 200 && res.status !== 0) return this.$message.error(res.message)
        this.addUserVisible = false
        this.$message.success('添加用户成功')
        this.getUserList()
      })
    },
    addUserCancel () {
      this.$refs.addUserRef.resetFields()
      this.addUserVisible = false
      this.$message.info('编辑已取消')
    },
    adminChange (value) {
      this.userInfo.role = Number(value)
    },
    // 编辑用户
    async editUser (id) {
      this.editUserVisible = true
      const { data: res } = await this.$http.get(`user/${id}`)
      console.log(id)
      console.log(res)
      this.userInfo = res.data
      this.userInfo.id = id
      console.log(this.userInfo.id)
      console.log(this.userInfo)
    },
    editUserOk () {
      this.$refs.editUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('输入的参数不合法')
        const { data: res } = await this.$http.put(`user/${this.userInfo.id}`, {
          username: this.userInfo.username,
          role: this.userInfo.role
        })
        if (res.status !== 200 && res.status !== 0) return this.$message.error(res.message)
        this.editUserVisible = false
        this.$message.success('更新用户成功')
        this.getUserList()
      })
    },
    editUserCancel () {
      this.$refs.editUserRef.resetFields()
      this.editUserVisible = false
      this.$message.info('编辑已取消')
    }
  }

}
</script>
