<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>


<style lang="less">
	html,body{
	  margin:0;
	  padding:0;
	}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  height: 100vh;
}
*{
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  list-style: none;
}
/deep/.el-message{
  height: 0.5rem;
}
/deep/.el-message__icon{
  margin-right: 0.15rem!important;
  font-size: 0.3rem!important;
}
/deep/.el-icon-error:before{
  font-size: 0.2rem;
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
}
</style>
