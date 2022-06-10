<template>
        <div class="manage-header" > 
	
        <el-col :span="7" style="margin-bottom:12px">
			<el-form ref="form">
            <el-input placeholder="请输入关键词" v-model="searchData.keyword" style="margin-top:12px;"  icon="el-icon-search"></el-input>
              <div class="login_btn"  style="margin-top:-42px;margin-left: 420px;"   @click="searchfile(searchData)" >搜索</div>
			</el-form>
        </el-col>
        
  <el-table
    :data="tableData"
	
    style="width: 83%">
  <el-table-column

        label="发送人"
        width="180">
		
		<template slot-scope="scope">
		 <span style="margin-left: 10px">{{ scope.row.sendname }}</span>
		 </template>
			
      </el-table-column>
	  
      <el-table-column
        label="哈希值"
        width="180">
		
		<template slot-scope="scope">
		 <span style="margin-left: 10px">{{ scope.row.hash }}</span>
		 </template>
		
      </el-table-column>
      <el-table-column
       
        label="文件名称"
		width="300">
		
		<template slot-scope="scope">
		 <span style="margin-left: 10px">{{ scope.row.filename }}</span>
		 </template>
		 
      </el-table-column>
    
    <el-table-column label="操作" >
      <template slot-scope="scope">
        <el-button
          size="mini"
          @click="handleSee(scope.$index,scope.row)">查看</el-button>
        <el-button
          size="mini"
          type="danger"
          @click="handleYan(scope.$index,scope.row)">验证</el-button>
      </template>
    </el-table-column>
  </el-table>
  
    <el-pagination style="margin-left: 920px;"
      layout="prev, pager, next"
      :total="50">
    </el-pagination>
	
	<el-divider></el-divider>
	
	<el-col :span="20" style="margin-top:2px;margin-left: 0px;" >
	<el-card style="margin=top:20px ;height:350px;">
	            <div class="login-info">
	              <p class="b">{{this.neirong}}</p> 
	           <div style = "height:280px" ></div>
	           </div>
	
	</el-card>
	</el-col>
  
        </div>	  
		
</template>

<script>
import { searchfile,handleYan } from '@/controller/api'
import { getStore } from '../../controller/utils'
export default{
  
    name:'pageTwo',
    data(){
        return{
			//每个人登录时的名字
          searchData:{
			  keyword:'',
			  username:''
		  },
		  tableData:[],
		  content:'',
		  neirong:'',
		  Nei:'',
		  wen:{
			  contenta:''
		  },
		  findcontentByhash:''
        }
    },
	created(){
	  this.searchData.username = getStore('username');
	},
    methods:{
        searchfile(obj){
        	if(/^\s*$/g.test(obj.keyword)){
        	  this.$message({
        	    message: '请输入关键词',
        	    type: 'error'
        	  });
        	  return
        	}   
         	  var params
         	    params = {
         	      "receivename": obj.username,
				  "keyword":obj.keyword
         	  	 }
           
          searchfile(params).then(res=>{
			  
            console.log(res);
			console.log(res.data);   
			console.log(this);
				
			this.tableData = res.data.data;
			
						
          })
    
        },
	
        handleSee(index,row) {
			
		this.findcontentByhash = this.tableData.map(o=>{
			if(o.hash.toString().replace(/,/g, '') == row.hash)
			{
				return o.content
			}
		}) 
		
        this.neirong = this.findcontentByhash.toString().replace(/,/g, '')
		this.neirong = this.neirong.toString().replace(/\<|>/g,'')
		this.neirong = this.neirong.toString().replace(/\p|p/g,'')
		this.neirong = this.neirong.toString().replace(/\/|/g,'')
		
      }, 
	  handleYan(index,row){
		  var params
		  this.findcontentByhash = this.tableData.map(o=>{
		  	if(o.hash.toString().replace(/,/g, '') == row.hash)
		  	{
		  		return o.content
		  	}
		  })  
		  this.Nei = this.findcontentByhash.toString().replace(/,/g, '')
		  params = {  
		    contenta: this.Nei
		  }
		  
		handleYan(params).then(res=>{
			// console.log(res)
			if(res.data.code == '0'){
			this.$message.success('验证成功，数据完整')
		 }else{
		     this.$message.error('数据不完整')
			   }
	     })
       	}
		
	}

}
</script>


<style lang="less" scoped>
	.login_btn{
	  margin: 0.4rem auto 0;
	  width: 1.28rem;
	  height: 0.55rem;
	  line-height: 0.55rem;
	  text-align: center;
	  border-radius: 0.08rem;
	  color: #fff;
	  font-size: 0.2rem;
	  background: #0378FF;
	  cursor: pointer;
	  margin-top: 109px;
	  margin-left: -1px;
	  
	}
	.b{
	   
	   font-size: 20px;
	}
	
</style>