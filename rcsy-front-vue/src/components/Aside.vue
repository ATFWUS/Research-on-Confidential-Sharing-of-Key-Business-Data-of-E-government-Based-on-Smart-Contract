<template>
	<el-menu default-active="1-4-1" class="el-menu-vertical-demo" background-color="#000000" 
	text-color="#FFFFFF" active-text-color="#FFFFFF" 
	@open="handleOpen" @close="handleClose" :collapse="isCollapse">
	       
	        <h6 style="font-size: 18px;text-align: center;margin-top: 15px;color: #F0F0F0;"> {{isCollapse ? '密文' : '关键密文共享系统'}}</h6>
	     
	 <el-menu-item @click="clickMenu(item)" v-for="item in noChildren" :index="item.path+''" :key="item.path">
	    <i :class="'el-icon-'+item.icon"></i>
	    <span slot="title">{{item.lable}}</span>
	  </el-menu-item>
	 
	  <el-submenu v-for="item in hasChildren"  :index="item.path+''" :key="item.path">
	    <template slot="title">
	      <i :class="'el-icon-'+item.icon"></i>
	      <span slot="title">{{item.lable}}</span>
	    </template>
	    <el-menu-item-group v-for="(subitem,subindex) in item.children" :key="subitem.path" style="margin-top:-40px">
	    <el-menu-item @click="clickMenu(subitem)" :index="subindex+''" :class="'el-icon-'+subitem.icon" >{{subitem.lable}} </el-menu-item>
	
	    </el-menu-item-group>
	     </el-submenu>
	    </el-menu>
</template>

<style lang="less" scoped>
  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
  }
  .el-menu{
      height: 100%;
      border: none;
      h4{
          color: #FFFFFF;
          text-align:center;
          line-height:48px
        
      }
  }
</style>

<script>
  export default {
    data() {
      return {
       // isCollapse: false,
        menu:[
            {
                path:'/',
                name:'home',
                lable:'首页',
                icon:'s-home',
                url:'Home/Home'
            },
            {
                lable:'文件管理',
                icon:'message',
                children:[
                    {
                    path:'/send',
                    name:'page1',
                    lable:'发送文件',
                     icon:'s-promotion',
                    url:'Other/PageOne'
                    },
                    {
                        path:'/search',
                        name:'page2',
                        lable:'搜索文件',
                        icon:'search',
                        url:"Other/PageTwo"
                    }
                ]
            }
        ]
      };
    },
    methods: {
      handleOpen(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath);
      },
      clickMenu(item){
          //页面全局跳转
          this.$router.push({
              name:item.name
          })
          this.$store.commit('selectMenu',item)
      }
    },
    computed:{
        noChildren(){
         return   this.menu.filter(item => !item.children)
        },
        hasChildren(){
         return this.menu.filter(item => item.children)
        },
        isCollapse(){
          return this.$store.state.tab.isCollapse
        }
       
    }
  }
</script>