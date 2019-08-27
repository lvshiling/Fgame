<template>
<div class="app-container">
    <el-form ref="dataForm" label-position="left" label-width="140px" style="width: 600px; margin-left:50px;">
        <el-form-item label="交易服务器ip">
            <el-input v-model="defaultInfo.tradeServerIp" autosize type="textarea"/>
        </el-form-item>
        <el-button v-waves class="filter-item" type="primary" @click="handleSubmit">提交</el-button>    
    </el-form>
    <el-dialog :visible.sync="dialogFormVisible" title="是否确定提交">
      <div>是否确定提交</div>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="updateServer">提交</el-button>
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
  getPlatformServerConfig,
  setPlatformServerConfig
} from "@/api/centerset";

export default {
  name: "CenterClientVersionSet",
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
    }
  },
  data() {
    return {
      dialogFormVisible: false,
      defaultInfo:{
          tradeServerIp:""
      }
    };
  },
  created() {
    this.initMetaData();
  },
  methods: {
    handleSubmit:function(e){
        this.dialogFormVisible = true
    },
    updateServer:function(e){
        setPlatformServerConfig(this.defaultInfo).then(res =>{
            this.showSuccess()
            this.dialogFormVisible = false
        })
    },
    initMetaData(){
        getPlatformServerConfig().then(res =>{
            this.defaultInfo = res
        })
    },
    showSuccess() {
      this.$message({
        message: "设置成功",
        type: "success",
        duration: 1000
      });
    }
  }
};
</script>
