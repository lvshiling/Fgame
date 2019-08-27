<template>
    <div class="app-container">
        <div class="filter-container">
            <el-select v-model="listQuery.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>

            <el-select v-model="listQuery.platformId" placeholder="平台" clearable style="width: 160px" class="filter-item" >
                <el-option v-for="item in tempPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilterExport">导出</el-button>
        </div>

        <el-table
            v-loading="listLoading"
            :key="tableKey"
            :data="list"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;">
            <el-table-column  label="战区服务器序号" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.zhanquServerId }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="战区服务器名称" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.zhanquServerName }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="战区服" min-width="200px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.zhanQuChild}}</span>
                </template>
            </el-table-column>

        </el-table>
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import { parseTime } from "@/utils/index";
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { getZhanQuCenterServerList,getZhanQuCenterServerListExport } from "@/api/centerserver";

export default {
  name: "CenterServerZhanQu",
  directives: {
    waves
  },
  filters: {
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}");
    },
    parseSecond: function(value) {
      let hour = parseInt(value / 60 / 60 / 1000);
      let reseMinute = value % (60 * 60);
      let minute = parseInt(reseMinute / 60);
      let reseSecond = reseMinute % 60;
      return hour + "时" + minute + "分" + reseSecond + "秒";
    },
    parsesex: function(value) {
      if (value == 1) {
        return "男";
      }
      if (value == 2) {
        return "女";
      }
      return value;
    },
    parseYesOrNo:function(value){
        if(value == 1){
            return "是"
        }
        return "否"
    }
  },
  created() {
    this.initMetaData();
    // this.getList();
  },
  data() {
    return {
      listLoading: false,
      tableKey: 0,
      total: 0,
      listQuery: {
        pageIndex: 1,
        centerServerName:"",
        platformId: undefined,
        channelId: undefined,
        serverId: undefined
      },
      textMap: {
        update: "编辑",
        create: "添加"
      },
      dialogStatus: "update",
      dialogFormVisible: false,
      temp: {},
      list: [],
      channelList: [],
      platformList: [],
      groupList: [],
      tempPlatformList: [],
      serverList: [],
      chatForbidTimeArray: [],
      monitorTemp: {}
    };
  },
  methods: {
    handleFilter: function() {
      console.log(this.listQuery);
      this.listQuery.pageIndex = 1;
      this.getList();
    },
    handleFilterExport:function(){
        this.exportList()
    },
    getList() {
      if(this.listQuery.platformId){
          let item = this.findPlatFormItem(this.listQuery.platformId)
          if(item.centerPlatformId){
              this.listQuery.centerPlatformId = item.centerPlatformId
          }
      }else{
          this.listQuery.centerPlatformId = undefined
      }
      if(!this.listQuery.centerPlatformId){
          this.$message({
                message: "请选择平台",
                type: "error",
                duration: 1500
            });
            return
      }
      console.log(this.listQuery)
      this.listLoading = true;
      getZhanQuCenterServerList(this.listQuery)
        .then(res => {
          this.list = res.itemArray;
          this.listLoading = false;
        })
        .catch(() => {
          this.listLoading = false;
        });
    },
    exportList(){
      if(this.listQuery.platformId){
          let item = this.findPlatFormItem(this.listQuery.platformId)
          if(item.centerPlatformId){
              this.listQuery.centerPlatformId = item.centerPlatformId
          }
      }else{
          this.listQuery.centerPlatformId = undefined
      }
      if(!this.listQuery.centerPlatformId){
          this.$message({
                message: "请选择平台",
                type: "error",
                duration: 1500
            });
            return
      }
      console.log(this.listQuery)
      this.listLoading = true;
      getZhanQuCenterServerListExport(this.listQuery)
        .then(res => {
          this.listLoading = false;
          if (!res) {
                return
            }
            let url = window.URL.createObjectURL(new Blob([res.data]))
            let link = document.createElement('a')
            link.style.display = 'none'
            link.href = url
            link.setAttribute('download', '战区列表.xlsx')

            document.body.appendChild(link)
            link.click()
        })
        .catch(() => {
          this.listLoading = false;
        });
    },
    handleCurrentChange(e) {
      console.log(e);
      this.listQuery.pageIndex = e;
      this.getList();
    },
    initMetaData() {
      getAllChannel().then(res => {
        this.channelList = res.itemArray;
      });
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
        // this.tempPlatformList = this.platformList;
      });
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
          });
        }
      }
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

