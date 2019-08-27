<template>
<div>
    <el-form ref="dataForm" label-position="left" label-width="140px" style="width: 600px; margin-left:50px;">
        <el-form-item label="渠道">
            <el-select v-model="listQuery.channelId" placeholder="渠道" style="width: 180px" class="filter-item" clearable @change="handleChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>
        </el-form-item>
        <el-form-item label="平台">
            <el-select v-model="listQuery.platformId" placeholder="平台" style="width: 180px" class="filter-item" clearable @change="handlePlatformChange">
                <el-option v-for="item in tempPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>
        </el-form-item>
        <el-form-item label="服务器">
            <el-select v-model="listQuery.serverId" collapse-tags placeholder="服务器" style="width: 180px" class="filter-item" clearable @change="handleServerChange">
                <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>
        </el-form-item>
        <el-form-item label="公告内容">
            <el-input v-model="listQuery.content"/>
        </el-form-item>
        <el-form-item label="间隔时间(分钟)">
            <el-input v-model="listQuery.intervalTime" type="number"/>
        </el-form-item>
        <el-form-item label="起止时间">
            <el-date-picker v-model="listQuery.startEnd" type="datetimerange" range-separator="至" start-placeholder="开始时间" end-placeholder="结束时间">
            </el-date-picker>
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
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { yesOrNoList } from "@/types/public";
import { parseTime } from "@/utils/index";
import { notice } from "@/api/notice";
import {
  getServerState,
  registerServerState,
  serverSetloglist
} from "@/api/alliance";

export default {
  name: "NoticeSet",
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
      channelList: [],
      platformList: [],
      tempPlatformList: [],
      serverList: [],
      tableKey: 1,
      listLoading: false,
      list: [],
      total:0
    };
  },
  created() {
    this.initMetaData();
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    },
    handleSubmit(e) {
      // if (!this.listQuery.serverId) {
      //   Message({
      //     message: "请选择参数",
      //     type: "error",
      //     duration: 1.5 * 1000
      //   });
      //   return;
      // }
      this.dialogFormVisible = true;
    },
    handleChannelChange: function(e) {
      if (e) {
        this.listQuery.platformId = undefined;
        this.tempPlatformList = this.findPlatFormList(e);
        if (this.tempPlatFormList && this.tempPlatFormList.length > 0) {
          this.listQuery.platformId = this.tempPlatFormList[0].platformId;
        }
        this.groupList = [];
        this.listQuery.serverId = undefined;
        this.listQuery.closeFlag = undefined;
      }
    },
    handlePlatformChange: function(e) {
      console.log(e);
      if (e) {
        let item = this.findPlatFormItem(e);

        if (item) {
          getCenterServerList(item.centerPlatformId).then(res => {
            this.serverList = res.itemArray;
            this.listQuery.serverId = undefined;
            this.listQuery.closeFlag = undefined;
          });
        }
      }
    },
    handleServerChange: function(e) {
      
    },
    updateServer() {
      if (this.listQuery.startEnd && this.listQuery.startEnd.length == 2) {
        this.listQuery.beginTime = this.listQuery.startEnd[0].valueOf();
        this.listQuery.endTime = this.listQuery.startEnd[1].valueOf();
      }
      var postData = {
        channelId: parseInt(this.listQuery.channelId),
        platformId: parseInt(this.listQuery.platformId),
        serverId:parseInt(this.listQuery.serverId),
        content:this.listQuery.content,
        intervalTime:parseInt(this.listQuery.intervalTime),
        beginTime:parseInt(this.listQuery.beginTime),
        endTime:parseInt(this.listQuery.endTime),
      };
      notice(postData).then(res => {
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    findPlatFormList(channelId) {
      if (!this.platformList || this.platformList.length == 0) {
        return;
      }
      return this.platformList.filter(function(item, index) {
        return item.channelId == channelId;
      });
    },
    findPlatFormItem(platformId) {
      let platform = this.platformList.find(n => {
        return n.platformId == platformId;
      });
      if (platform) {
        return platform;
      }
      return undefined;
    },
    findServerItem(id) {
      let server = this.serverList.find(n => {
        return n.id == id;
      });
      if (server) {
        return server;
      }
      return undefined;
    },
    initMetaData() {
      this.yesOrNoArray = yesOrNoList;
      getAllChannel().then(res => {
        this.channelList = res.itemArray;
      });
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
      });
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
