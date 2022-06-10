<template>
    
   <div style="margin-top:20px">
    
	  <el-form ref="form" :model="senddata" >
     <el-input placeholder="请输入接收者名字" v-model="senddata.receivename"  value="" style="margin-bottom:12px ;width:300px; height=30px; " ></el-input>
    <el-input placeholder="请输入文件名称" v-model="senddata.filename" value=""  style="margin-bottom:12px ;width:300px; height=30px;left:15px"></el-input>
      <el-input placeholder="请输入关键词" v-model="senddata.keyword"  value="" style="margin-bottom:12px ;width:300px; height=30px;left:30px"></el-input>
        <!--富文本框-->
        <quill-editor ref="myTextEditor" v-model="senddata.content"  value=""
        :options="editorOption" @change="onEditorChange" style="margin-bottom:55px;height: 300px;" ></quill-editor>
        <div class="login_btn" @click="send(senddata)" >发送</div>
		 </el-form>
    </div>
</template>


<script>
import { send } from '@/controller/api'
import { getStore } from '../../controller/utils'


export default{

    name:'pageOne',
    data(){
        return{
			username:'',
			
			senddata:{
			filename:'',
            content:'',
			receivename:'',
			keyword:''
			},
             editorOption: {
               placeholder: '编辑文章内容'
             }
        }
    },
	created(){
	  this.username = getStore('username');
	},
      methods:{
        onEditorChange({html}) { // text是文本内容
            this.content = html;
        },  
		send(obj){
			if(/^\s*$/g.test(obj.filename)){
			  this.$message({
			    message: '请输入文件名',
			    type: 'error'
			  });
			  return
			}
			if(/^\s*$/g.test(obj.content)){
			  this.$message.error('请输入内容')
			  return
			}
			if(/^\s*$/g.test(obj.keyword)){
			  this.$message.error('请输入关键词')
			  return
			}
			if(/^\s*$/g.test(obj.receivename)){
			  this.$message.error('请输入收件人')
			  return
			}
		  var params
		    params = {
		      "sendname": this.username,
		      "receivename": obj.receivename,
		      "keyword": obj.keyword,
			  "content":obj.content,
			  "filename":obj.filename
			  }
		   console.log(params)
		  send(params).then(res=>{
		     //this.$refs[form].resetFields();
			  console.log(this);
		    if(res.data.code == '0'){
		      this.$message.success('操作成功') ;
			     console.log(this);
		    }else{
		      this.$message.error(res.data.msg);
			   
		    }
			
		  })
		window.location.reload();
		}
      },
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
	
</style>
