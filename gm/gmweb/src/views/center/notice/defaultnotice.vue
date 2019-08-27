<template>
<div>
    <el-form ref="dataForm" label-position="left" label-width="140px" style="width: 600px; margin-left:50px;">
        <el-form-item label="公告内容">
            <el-input v-model="defaultInfo.content" autosize type="textarea"/>
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
  getCenterDefaultNotice,
  updateCenterDefaultNotice
} from "@/api/centernotice";

export default {
  name: "CenterNoticeSet",
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
      activeName: "first",
      listQuery: {
        channelId: undefined,
        platformId: undefined,
        serverId: undefined,
        closeFlag: undefined
      },
      dialogFormVisible: false,
      defaultInfo:{
          content:"",
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
        updateCenterDefaultNotice(this.defaultInfo).then(res =>{
            this.showSuccess()
            this.dialogFormVisible = false
        })
    },
    initMetaData(){
        getCenterDefaultNotice().then(res =>{
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
