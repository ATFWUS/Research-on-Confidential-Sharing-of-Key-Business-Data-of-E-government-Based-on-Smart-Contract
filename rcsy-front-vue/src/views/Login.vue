<template>
  <div class="login">
    <div class="login_box" v-if="flag==='login'">
      <div class="login_title">基于智能合约的关键密文共享系统</div>
      <el-form ref="form" :model="loginData">
        <el-input v-model="loginData.username" placeholder="请输入用户账号"></el-input>
        <el-input v-model="loginData.password" placeholder="请输入登录密码" show-password></el-input>
        <div class="login_btn" @click="login(loginData)">登录</div>
        <div class="register_btn" @click="flag='register'">注册</div>
      </el-form>
    </div>
    <div class="register_box" v-else>
      <div class="login_title">基于智能合约的关键密文共享系统</div>
      <el-form :model="registerData">
        <el-input v-model="registerData.username" placeholder="请输入用户账号"></el-input>
        <el-input v-model="registerData.password" placeholder="请输入登录密码" show-password></el-input>
      
        <div class="login_btn" @click="register(registerData)">注册</div>
        <div class="register_btn" @click="flag='login'">登录</div>
      </el-form>
    </div>
  </div>
</template>

<script>
import { login,register } from '@/controller/api'
import { setStore } from '@/controller/utils'
export default {
	name:'login',
  data(){
    return {
      flag: 'login',
      loginData: {
        username: '',
        password: ''
      },
      registerData: {
        username: '',
        password: ''
       
      }
    }
  },
  methods:{
    login(obj){
      if(/^\s*$/g.test(obj.username)){
        this.$message({
          message: '请输入用户账户',
          type: 'error'
        });
        return
      }
      if(/^\s*$/g.test(obj.password)){
        this.$message.error('请输入登录密码')
        return
      }
      var params = {
        username: obj.username,
        password: obj.password
      }
      login(params).then(res=>{
        // console.log(res)
        if(res.data.code == '0'){
          this.$message.success('登录成功')
          setStore('username',obj.username)
          this.$router.push('/')
        }else{
          this.$message.error(res.data.msg)
        }
      })
    },
    register(obj){
      if(/^\s*$/g.test(obj.username)){
        this.$message({
          message: '请输入用户账户',
          type: 'error'
        });
        return
      }
      if(/^\s*$/g.test(obj.password)){
        this.$message.error('请输入登录密码')
        return
      }
      var params
      if(obj.nickname){
        params = {
          username: obj.username,
          password: obj.password
         
        }
      }else{
        params = {
          username: obj.username,
          password: obj.password,
        }
      }
      register(params).then(res=>{
        // console.log(res)
        if(res.data.code == '0'){
          this.$message.success('注册成功，请登录')
        }else{
          this.$message.error(res.data.msg)
        }
      })
    }
  }
}
</script>

<style lang="less" scoped>
.login{
  position: fixed;
  width: 100%;
  top: 0;
  right: 0;
  left: 0;
  bottom: 0;
  background: url(../assets/images/login_bg.jpg) no-repeat;
  background-size: cover;
  .login_box{
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    width: 5.7rem;
    height: 4.8rem;
    background-color: #fff;
    border-radius: 0.08rem;
    padding: 0 0.54rem;
    .login_title{
      font-size: 0.26rem;
      color: #071C35;
      line-height: 0.37rem;
      margin-top: 0.6rem;
      text-align: center;
      font-weight: bold;
    }
    /deep/.el-form{
      margin-top: 0.4rem;
      /deep/.el-input{
        display: block;
      }
      /deep/.el-input__icon{
        height: 0.55rem;
      }
      /deep/.el-input__inner{
        width: 100%;
        height: 0.55rem;
        line-height: 0.55rem;
        background:rgba(243,245,248,.8);
        border-radius: 0.08rem;
        outline: none;
        font-size: 0.16rem;
        color:rgba(7,28,53,.7);
        border: none;
        &::placeholder{
          color:rgba(7,28,53,.3);
          font-size: 0.16rem;
        }
        &:first-of-type{
          margin-bottom: 0.2rem;
        }
      }
      .login_btn{
        margin: 0.4rem auto 0;
        width: 3.28rem;
        height: 0.55rem;
        line-height: 0.55rem;
        text-align: center;
        border-radius: 0.08rem;
        color: #fff;
        font-size: 0.2rem;
        background: #0378FF;
        cursor: pointer;
      }
      .register_btn{
        width: 0.5rem;
        margin: 0.15rem auto 0;
        cursor: pointer;
        font-size: 0.2rem;
        color: #0378FF;
        line-height: 0.28rem;
        text-align: center;
      }
    }
  }
  .register_box{
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    width: 5.7rem;
    height: 5.4rem;
    background-color: #fff;
    border-radius: 0.08rem;
    padding: 0 0.54rem;
    .login_title{
      font-size: 0.26rem;
      color: #071C35;
      line-height: 0.37rem;
      margin-top: 0.6rem;
      text-align: center;
      font-weight: bold;
    }
    /deep/.el-form{
      margin-top: 0.4rem;
      /deep/.el-input{
        display: block;
      }
      /deep/.el-input__icon{
        height: 0.55rem;
      }
      /deep/.el-input__inner{
        width: 100%;
        height: 0.55rem;
        line-height: 0.55rem;
        background:rgba(243,245,248,.8);
        border-radius: 0.08rem;
        outline: none;
        font-size: 0.16rem;
        color:rgba(7,28,53,.7);
        border: none;
        &::placeholder{
          color:rgba(7,28,53,.3);
          font-size: 0.16rem;
        }
        &:first-of-type{
          margin-bottom: 0.2rem;
        }
      }
      .login_btn{
        margin: 0.4rem auto 0;
        width: 3.28rem;
        height: 0.55rem;
        line-height: 0.55rem;
        text-align: center;
        border-radius: 0.08rem;
        color: #fff;
        font-size: 0.2rem;
        background: #0378FF;
        cursor: pointer;
      }
      .register_btn{
        width: 0.5rem;
        margin: 0.15rem auto 0;
        cursor: pointer;
        font-size: 0.2rem;
        color: #0378FF;
        line-height: 0.28rem;
        text-align: center;
      }
    }
  }
}
</style>
