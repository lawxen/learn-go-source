<template>
  <div class="main-layout card-bg-1">
    <div class="container d-flex flex-column">
      <div class="row no-gutters text-center align-items-center justify-content-center min-vh-100">
        <div class="col-12 col-md-6 col-lg-5 col-xl-4">
          <h1 class="font-weight-bold">用户登录</h1>
          <p class="text-dark mb-3">民主、文明、和谐、自由、平等</p>
          <div class="mb-3">
            <div class="form-group">
              <label for="username" class="sr-only">账号</label>
              <input type="text" class="form-control form-control-md" id="username" placeholder="请输入账号"
                     v-model="username">
            </div>
            <div class="form-group">
              <label for="password" class="sr-only">密码</label>
              <input type="password" class="form-control form-control-md" id="password"
                     placeholder="请输入密码" v-model="password">
            </div>
            <button class="btn btn-primary btn-lg btn-block text-uppercase font-weight-semibold" type="submit"
                    @click="login()">登录
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Signin',
  data () {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    login: function () {
      // 判断是否输入账号
      if (this.username.length > 0 && this.password.length > 0) {
        // 向后端发送POST请求
        let data = new FormData();
        data.append('username',this.username);
        data.append('password',this.password);
        this.axios.post('http://127.0.0.1:8000/', data).then((res)=> {
          // POST请求发送成功则获取响应结果的result
          // 如果result为true，则说明存在此用户
          if (res.data.result) {
            // 将访问路由chat，并设置参数
            this.$router.push({
              path: '/product'
            })
          } else {
            // 当前用户不存在后端的数据库
            window.alert('账号不存在或异常')
            // 清空用户输入的账号和密码
            this.username = ''
            this.password = ''
          }}).catch(function () {
          // PSOT请求发送失败
          window.alert('账号获取失败')
          // 清空用户输入的账号和密码
          this.username = ''
          this.password = ''
        })
      } else {
        // 提示没有输入账号或密码
        window.alert('请输入账号或密码')
      }
    }
  }
}
</script>

<style scoped>
  .text-center {
    text-align: center!important;
  }
  .min-vh-100 {
      min-height: 100vh!important;
  }
  .align-items-center {
      align-items: center!important;
  }
  .justify-content-center {
      justify-content: center!important;
  }
  .no-gutters {
      margin-right: 0;
      margin-left: 0;
  }
  .row {
      display: flex;
      flex-wrap: wrap;
      margin-right: -15px;
      margin-left: -15px;
  }
  *, :after, :before {
      box-sizing: border-box;
  }
</style>
