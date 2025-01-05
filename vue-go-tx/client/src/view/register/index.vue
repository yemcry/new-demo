<template>
  <div class="register-container">
    <el-card class="register-card">
      <h2 class="title">注册</h2>
      <el-form :model="form" :rules="rules" ref="form" label-width="100px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="form.password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input type="password" v-model="form.confirmPassword" placeholder="请再次输入密码"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="register()" :disabled="form.password !== form.confirmPassword">注册</el-button>
          <el-dialog
            title="提示"
            :visible.sync="dialogVisible"
            width="30%">
            <span>注册成功</span>
            <span slot="footer" class="dialog-footer">
              <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
            </span>
          </el-dialog>
          <el-button @click="goToLogin">返回登录</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data() {
    const validateConfirmPassword = (rule, value, callback) => {
      if (value !== this.form.password) {
        callback(new Error('两次输入密码不一致'));
      } else {
        callback();
      }
    };
    return {
      form: {
        username: '',
        password: '',
        confirmPassword: '',
      },
      dialogVisible: false,
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 6, max: 9, message: '用户名长度应在6到9位之间', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 9, message: '密码长度应在6到9位之间', trigger: 'blur' }
        ],
        confirmPassword: [
          { required: true, message: '请再次输入密码', trigger: 'blur' },
          { validator: validateConfirmPassword, trigger: 'blur' }
        ],
      },
    };
  },
  methods: {
    async register() {
      // 如果两次密码不一致,则不执行注册操作
      if (this.form.password !== this.form.confirmPassword) {
        this.$notify({
          title: '警告',
          message: '密码不一致',
          type: 'warning'
        });
        return;
      }

      if (this.form.username.trim() === '' || this.form.password.trim() === '') {
        this.$notify({
          title: '警告',
          message: '用户名或密码不能为空',
          type: 'warning'
        });
        return;
      }

      try {
        const response = await axios.post('http://192.168.1.5:12345/register', {
          username: this.form.username,
          password: this.form.password,
        });
        // 处理注册成功的逻辑
        if (response.data.status === 200) {
          this.$router.push('/');
          this.dialogVisible = true
          console.log("注册成功")
        } else {
          this.error = response.data.message; // 显示后端返回的错误信息
        }
      } catch (err) {
        this.error = '注册失败，请检查输入信息。';
        console.error('注册失败:', err);
      }
    },

    goToLogin() {
      // 跳转到登录页面的逻辑
      this.$router.push('/');
    },
  },
};
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.register-card {
  width: 400px;
}
.title {
  text-align: center;
  margin-bottom: 20px;
}
</style>