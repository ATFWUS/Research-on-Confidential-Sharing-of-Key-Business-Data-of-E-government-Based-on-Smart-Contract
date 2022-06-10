<template>
    <el-container style="height:100%">
  <el-aside width="auto">
      <Aside></Aside>
  </el-aside>
  <el-container>
    <el-header>
      <Header></Header>
    </el-header>
    <el-main>
        <router-view></router-view>
    </el-main>
  </el-container>
</el-container>
</template>

<script>
import Aside from '../components/Aside.vue'
import Header from '../components/Header.vue'

export default{
 
    name:'Home',
    components:{
        Aside,
        Header
    },
    data(){
        return{}
    }

}
</script>
<style lang="less" scoped>
    .el-header{
        background-color: #eaeff5;
    }
    .el-main{
        padding-top: 0;
    }

</style>