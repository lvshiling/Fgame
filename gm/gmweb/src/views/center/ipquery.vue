<template>
<div class="app-container">
    <el-form ref="dataForm" label-position="left" label-width="140px" style="width: 600px; margin-left:50px;">
        <el-form-item label="IP">
            <el-input placeholder="请输入ip" v-model="defaultInfo.ip"/>
        </el-form-item>
        <el-form-item label="是否已经封禁">
            <span>{{defaultInfo.forbid | parseIpForbit}}</span>
            <el-button v-if="defaultInfo.forbid" v-waves class="filter-item" type="primary" @click="handleSubmit" >解禁</el-button> 
        </el-form-item>
    </el-form>
       
    <el-button v-waves class="filter-item" type="primary" @click="handleQuery">查询</el-button>
    
    <el-dialog :visible.sync="dialogFormVisible" title="是否确定提交">
      <div>是否确定解禁</div>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="updateServer">解禁</el-button>
      </div>
    </el-dialog>
</div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import { getCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { yesOrNoList } from "@/types/public";
import { parseTime } from "@/utils/index";
import {
  getCenterIpState,
  updateCenterIpUnForbid
} from "@/api/centeruser";

export default {
  name: "CenterIpQuery",
  directives: {
    waves
  },
  filters: {
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}");
    },
    parseYesOrNo:function(value){
      if(value == 1){
        return '是'
      }
      return '否'
    },
    parseIpForbit:function(value){
      if(value){
        return '是'
      }
      return '否'
    }
  },
  data() {
    return {
      dialogFormVisible: false,
      defaultInfo:{
          androidVersion:"",
          iosVersion:"",
          forbid:undefined,
          ip:""
      }
    };
  },
  created() {
  },
  methods: {
    handleSubmit:function(e){
        this.dialogFormVisible = true
    },
    updateServer:function(e){
        updateCenterIpUnForbid(this.defaultInfo).then(res =>{
            this.showSuccess()
            this.dialogFormVisible = false
            this.getIpInfo()
        })
    },
    handleQuery:function(e){
        if(this.defaultInfo.ip === undefined || this.defaultInfo.ip.length == 0){
            this.showErrorMsg('ip不能为空')
            return
        }
        getCenterIpState(this.defaultInfo).then(res =>{
            this.defaultInfo.forbid = res.forbid
            this.showSuccessMsg('获取成功')
            console.log(this.defaultInfo.forbid)
            this.dialogFormVisible = false
        })
    },
    getIpInfo:function(){
        getCenterIpState(this.defaultInfo).then(res =>{
            this.defaultInfo.forbid = res.forbid
            this.dialogFormVisible = false
        })
    },
    showSuccess() {
      this.$message({
        message: "设置成功",
        type: "success",
        duration: 1000
      });
    },
    showSuccessMsg(msg) {
      this.$message({
        message: msg,
        type: "success",
        duration: 1000
      });
    },
    showErrorMsg(msg) {
      this.$message({
        message: msg,
        type: "error",
        duration: 1000
      });
    },
  }
};
</script>
