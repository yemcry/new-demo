<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2 class="title">登录</h2>
      <el-form :model="form" :rules="rules" ref="form" label-width="100px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="form.password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="login()">登录</el-button>
          <el-dialog
            title="提示"
            :visible.sync="dialogVisible"
            width="30%">
            <span>登录成功</span>
            <span slot="footer" class="dialog-footer">
              <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
            </span>
          </el-dialog>
          <el-button @click="goToRegister">注册</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import axios from 'axios';
import { mapActions } from 'vuex';
export default {
  data() {
    return {
      form: {
        username: '',
        password: '',
      },
      dialogVisible: false,
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ],
      },
    };
  },
  methods: {
    ...mapActions(['setUser']),
    async login() {
      if (this.form.username.trim() === '' || this.form.password.trim() === '') {
        this.$notify({
          title: '警告',
          message: '用户名或密码不能为空',
          type: 'warning'
        });
        return;
      }
      try {
        const response = await axios.post('http://192.168.1.5:12345/login', {
          username: this.form.username,
          password: this.form.password,
        });
        // 处理登录成功的逻辑
        if (response.data.code === 0) {
          this.setUser(this.form.username)
          this.$router.push('/home');
          this.dialogVisible = true
          this.$notify({
          title: '登录成功',
          message: '用户'+this.form.username+'登录！！！',
          type: 'success'
          });
        } else {
          this.error = response.data.message; // 显示后端返回的错误信息
        }
      } catch (err) {
        this.$notify({
          title: '警告',
          message: '用户名或密码错误',
          type: 'warning'
        });
      }
    },

    goToRegister() {
      // 跳转到注册页面的逻辑
      this.$router.push('/register')
    },
  },
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.login-card {
  width: 400px;
}
.title {
  text-align: center;
  margin-bottom: 20px;
}
</style>