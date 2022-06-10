<template>
    <header>
        <div class="l-content">
            <el-button @click="handleMenu" plain icon="el-icon-menu" size="mini"></el-button>
          <h3 style="color:#000000"></h3>
            <el-breadcrumb separator="/">
             <!-- <el-breadcrumb-item v-for="item in tags" :key="item.path" :to="{ path: item.path }">{{item.lable}}
			 </el-breadcrumb-item> -->
           
            </el-breadcrumb>
        </div>
        <div class="r-content">
            <el-dropdown trigger="click" size="mini">
                <span>
                    <img class="user" :src = "userImg">
                </span>
                  <el-dropdown-menu slot="dropdown">	   
						<div class="con_right_btn" @click="showLogoutDialog">
						   退出
						 </div>
             </el-dropdown-menu>
            </el-dropdown>
        </div>	
		<div class="mask" v-if="exitDialogFlag">
		  <div class="dialog">
		    <div class="dialog_top">
		      <div class="dialog_top_title">退出程序</div>
		      <div class="dialog_top_right">
		        <img @click="closeExitDialog" src="../assets/images/close_icon.png" alt="">
		      </div>
		    </div>
		    <div class="dialog_con" style="{height : 1rem}">
		      <div class="sure_btn" @click="exitApp">确定</div>
		    </div>
		  </div>
		</div>
		
    </header>
</template>


<script>
import { mapState } from 'vuex'
import { getStore,removeStore } from '@/controller/utils'
import { logout } from '@/controller/api'

export default {
	 name:'Header',
  data() {
    return {
		dialogFlag: false,
		 exitDialogFlag: false,
	  userImg:require('../assets/images/User.jpg'),
    }
  },
  created(){
    this.username = getStore('username');
  },
  methods: {
      handleMenu(){
         this.$store.commit('CollapseMenu')
     },
	 showLogoutDialog() {
	   this.exitDialogFlag = true;	 
	 },
	 closeExitDialog() {
	   this.exitDialogFlag = false;
	 },
	 exitApp() {
	   logout().then(
	     res => {
	       console.log(res);
	       if (res.data.code == '0') {
	         this.$message('退出成功！');
			 this.exitDialogFlag = false;
	         this.$router.replace('/login');
	         removeStore('username')
	       } else {
	         this.$message.error('退出失败，请重试！');
	       }
	     }
	   ).catch(
	     error => {
	       console.log(error);
	     }
	   );
	 }
  },
  computed:{
          ...mapState({
              tags : state => state.tab.tabsList
          })
      }
}
</script>

<style lang="less" scoped>
 header{
     display: flex;
     height:100%;
     justify-content: space-between;
     align-items: center;
 }
 .l-content{
     display: flex;
     align-items: center;
     .el-button{
         margin-right: 20px;
     }
 }
 .r-content{
     .user{
         width: 40px;
         height: 40px;
         border-radius: 50%;
     }
 }
 .con_right_btn {
   cursor: pointer;
   span {
     font-size: 0.22rem;
     font-weight: bold;
     line-height: 0.6rem;
   }
   float: right;
   width: 0.8rem;
   height: 0.6rem;
   line-height: 0.6rem;
   text-align: center;
   background: #0378ff;
   border-radius: 0.08rem;
   color: #fff;
   font-size: 0.18rem;
 }
 .mask{
   position: fixed;
   width: 100%;
   top: 0;
   right: 0;
   left: 0;
   bottom: 0;
   background:rgba(0,0,0,.5);
   z-index: 999;
   .dialog{
     width: 3.8rem;
     height: 4rem;
     border-radius: 0.08rem;
     position: absolute;
     left: 50%;
     top: 50%;
     transform: translate(-50%, -50%);
     .dialog_top{
       width: 100%;
       height: 0.6rem;
       padding: 0 0.4rem;
       overflow: hidden;
       background: #262729;
       border-radius: 0.08rem 0.08rem 0 0 ;
       .dialog_top_title{
         float: left;
         font-size: 0.18rem;
         color: #fff;
         line-height: 0.6rem;
       }
       .dialog_top_right{
         float: right;
         img{
           display: block;
           width: 0.14rem;
           height: 0.14rem;
           margin-top: 0.23rem;
           cursor: pointer;
         }
       }
     }
     .dialog_con{
       background-color: #fff;
       width: 100%;
       height: 2.4rem;
       border-radius: 0 0 0.08rem 0.08rem;
       padding-top: 0.3rem;
       /deep/.el-form-item{
         margin-bottom: 0.3rem;
         height: 0.55rem;
       }
       /deep/.el-form-item__label{
         font-size: 0.16rem;
         color: #262729;
         padding: 0 0.2rem 0 0.3rem;
         text-align: left;
         line-height: 0.55rem;
       }
       /deep/.el-form-item__content{
         line-height: 0.55rem;
       }
       /deep/.el-input__icon{
         line-height: 0.55rem;
       }
       /deep/.el-input__inner{
         width: 5.16rem;
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
       }
       /deep/.el-select-dropdown__item{
         font-size: 0.16rem;
       }
       .sure_btn{
         cursor: pointer;
         margin: 0.3rem auto 0;
         width: 2.4rem;
         height: 0.6rem;
         line-height: 0.6rem;
         border-radius: 0.08rem;
         background: #0378FF;
         font-size: 0.18rem;
         color: #fff;
         text-align: center;
       }
     }
   }
 }
 
</style>