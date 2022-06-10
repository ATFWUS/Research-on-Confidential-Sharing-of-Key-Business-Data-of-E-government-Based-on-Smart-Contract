<template>
    <el-row class="home" :gutter="20">
        <el-col :span="12" style="margin-top:20px">
            <el-card shadow="hover">
                <div class="user">
                    <img :src = "userImg" />
                    <div class="userinfo">
                        <p class="name">{{username}}</p>
                    </div>
                </div>
                <div class="login-info">
                    <p>天气:<span>晴</span></p>
                     <p>今天的日期:<span>{{FormatTime(nowTime)}}</span></p>
                    <p>上次登录地点:<span>成都</span></p>
                </div>
            </el-card>  
        </el-col>
        
         
          <el-col :span="10" style="margin-top:20px" >
         <el-card style="margin=top:20px ;height:318px;">
                     <div class="login-info">
                       <p class="b">中国区块链行业市场规模(亿元)</p>
                    <div style = "height:280px" ref="zhuecharts"></div>
                    </div>

         </el-card>
         </el-col>
          <el-col :span="10" style="margin-top:20px">
         <el-card style="margin=top:20px ;height:360px;">
                     <div class="login-info">  
                       <p class="a">信息泄露的网站备案情况分析</p>
                        <div style = "height:280px" ref="videoecharts"></div>
                    </div>

         </el-card>
         </el-col>
          <el-col :span="12" style="margin-top:20px">
         <el-card style="margin=top:20px ;height:360px;">
                     <div class="login-info">  
                       <p class="a">云服务器的发展趋势</p>
                        <div style = "height:280px" ref="zheecharts"></div>
                    </div>

         </el-card>
         </el-col>
                 
    </el-row>
</template>
<script>


import { getStore } from '../../controller/utils'
export default{
    name:'home',
    data(){
        return{
            userImg:require('../../assets/images/User.jpg'),
             timer: undefined,
             nowTime: new Date(),
             latedata:"2020-1-9",
			 username: ''
        }
    },
   
    created() {
      this.timer = setInterval(() => {
      this.nowTime= new Date().toLocaleString();
    });
	 this.username = getStore('username');
  },
    beforeDestroy() {
    if (this.timer) {
      clearInterval(this.timer);
    }
  },
  mounted(){
      this.drawLine()
      this.drawZhu()
      this.drawheng()
  },
    methods: {
    FormatTime() {
      //返回显示的日期时间格式
      var date = new Date();
      var year = this.formatTime(date.getFullYear());
      var month = this.formatTime(date.getMonth() + 1);
      var day = this.formatTime(date.getDate());
      var hour = this.formatTime(date.getHours());
      var minute = this.formatTime(date.getMinutes());
      var second = this.formatTime(date.getSeconds());
      var weekday = date.getDay();
      var weeks = new Array(
          "星期日",
          "星期一",
          "星期二",
          "星期三",
          "星期四",
          "星期五",
          "星期六"
      );
      var week = weeks[weekday];
      return `${year}-${month}-${day} ${hour}:${minute}:${second} ${week}`;
    },
    formatTime(n) {
      //判断是否需要加0
      if (n < 10) {
        return "0" + n;
      } else {
        return n;
      }
    },
    //饼状图
    drawLine(){
     
      // 指定图表的配置项和数据
      var option = {
          tooltip:{
             formatter: " <br/>{b}:  ({d}%)"
          },
        series: [
        {
      type: 'pie',
      data: [
        {
          value: 69.7,
          name: '企业'
        },
        {
          value: 19.5,
          name: '政府机构'
        },
        {
          value: 4,
          name: '事业单位'
        },
        {
          value: 0.8,
          name: '社会团体'
        },
        {
          value: 6,
          name: '未备案'
        }
      ]
    }
  ]
      }
      // 使用刚指定的配置项和数据显示图表。
     const v = this.$echarts.init(this.$refs.videoecharts)
      v.setOption(option);
    },
    //画柱状图
    drawZhu () {
       const z = this.$echarts.init(this.$refs.zhuecharts)
       z.setOption({
         tooltip:{},
        xAxis: {
          name:'年份',
    data: ['2015', '2016', '2017', '2018', '2019', '2020', '2021','2022']
  },
  yAxis: {
    name:'亿元'
  },
  series: [
    {
      type: 'bar',
      data: [0.06, 0.11, 0.32, 0.67, 1.12, 1.73, 2.74,4.59]
    }
  ]   
       })
    },
  //画横折线图图
    drawheng () {
      const wxc = this.$echarts.init(this.$refs.zheecharts)
     wxc.setOption ( {
       tooltip:{
      
       },
  xAxis: {
    name:'年份',
    data: ['2018', '2019', '2020', '2021', '2022']
  },
  yAxis: {
    name:'亿元'
  },
  series: [
    {
      data: [962.8, 1312.5, 1737.4, 2529.7, 2902.9],
      type: 'line',
      smooth: true
    }
  ]
})
    }


    
  }
}
</script>


<style lang="less" scoped>
    .a{
       text-align: center;
    }
    .b{
       text-align: center;
    }
  
</style>
