<template>
 <div class="app-container">
    <div class="filter-container">
        <div class="filter-item">
            <el-select v-model="listQuery.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>

            <el-select v-model="listQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handlePlatformChange">
                <el-option v-for="item in tempPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>
            <el-date-picker v-model="listQuery.startEnd" type="datetimerange" range-separator="至" start-placeholder="开始时间" end-placeholder="结束时间">
            </el-date-picker>
        </div>
        <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
    </div>
    <el-table
            v-loading="listLoading"
            :key="tableKey"
            :data="logData"
            border
            fit
            highlight-current-row
            height="750"
            style="width: 100%;margin-top:15px;">
            <el-table-column label="日期" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.datestr | parseTimeFilter }}</span>
                </template>
            </el-table-column>
            <el-table-column label="时间" width="80px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.minuteindex | parseSingleTimeFilter}}</span>
                </template>
            </el-table-column>
            <el-table-column v-for="(item,index) in metaColumnArray" :label="item.serverName"  min-width="100px" align="left" v-bind:key="index">
                <template slot-scope="scope">
                    <span>{{ scope.row.onlineMap[item.serverId] }}</span>
                </template>
            </el-table-column>
        </el-table>
 </div>    
    
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getOnLineReport } from "@/api/report";
import { parseTime } from "@/utils/index";
import { Message, MessageBox } from "element-ui";
export default {
  name: "OnLineReport",
  directives: {
    waves
  },
  filters: {
    parseTimeFilter: function(value) {
      return parseTime(value, "{y}-{m}-{d}");
    },
    parseSingleTimeFilter: function(value) {
      return parseTime(value, "{h}:{i}:{s}");
    },
    commonFilter: function(value, type) {
      if (type == "datetime") {
        return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
      }
      if (type == "byte") {
        return binaryToStr(value);
      }
      if (type == "normal") {
        return value;
      }
      return value;
    }
  },
  created() {
    this.initMetaData();
  },
  data() {
    return {
      tableKey: 1,
      //基础元数据
      metaLogType: [], //日志列表
      platformList: [], //平台列表
      serverList: [], //服务器列表
      listQuery: {
        tableName: undefined,
        beginTime: undefined,
        endTime: undefined,
        platformId: undefined,
        sid: undefined,
        serverType: -1,
        startEnd: [],
        serverId: undefined,
        pageIndex: 1,
        playerId: undefined
      },
      logData: [],
      totalCount: 0,
      listLoading: false,
      tableMetaMap: new Map(),
      metaColumnArray: [],
      channelList: [],
      tempPlatformList: []
    };
  },
  methods: {
    handleFilter(e) {
      if (!this.listQuery.centerPlatformId) {
        Message({
          message: "请选平台",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      this.listQuery.pageIndex = 1;
      if (this.listQuery.startEnd && this.listQuery.startEnd.length == 2) {
        this.listQuery.beginTime = this.listQuery.startEnd[0].valueOf();
        this.listQuery.endTime = this.listQuery.startEnd[1].valueOf();
      }
      this.loadData();
    },
    handleCurrentChange(e) {
      this.listQuery.pageIndex = e;
      this.loadData();
    },
    initMetaData() {
      let startDate = new Date();
      startDate = new Date(
        startDate.getFullYear(),
        startDate.getMonth(),
        startDate.getDate()
      );
      let endDate = new Date();
      endDate.setDate(endDate.getDate() + 1);
      endDate = new Date(
        endDate.getFullYear(),
        endDate.getMonth(),
        endDate.getDate()
      );
      this.listQuery.startEnd = [startDate, endDate];
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
        this.listQuery.centerPlatformId = undefined;
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
        this.listQuery.centerPlatformId = item.centerPlatformId
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
    loadData() {
      this.listLoading = true;
      getOnLineReport(this.listQuery).then(res => {
        this.logData = res.itemArray;
        this.metaColumnArray = res.serverArray
        this.listLoading = false;
      });
    }
  }
};
</script>

