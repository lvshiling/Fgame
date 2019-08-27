<template>
 <div>
    <div class="filter-container">
        <el-select v-model="listQuery.successFlagSel" placeholder="是否成功" style="width: 160px" class="filter-item" clearable @change="handleYesOrNoChange">
            <el-option v-for="item in yesOrNoArray" :key="item.key" :label="item.name" :value="item.key" />
        </el-select>
        <div class="filter-item">
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
            style="width: 100%;margin-top:15px;">
            <el-table-column label="渠道" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ findChannelName(scope.row.channelId) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="平台" width="160px" align="left">
                <template slot-scope="scope">
                    <span>{{ findPlatformName(scope.row.platformId) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="服务器" width="160px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverName}}</span>
                </template>
            </el-table-column>
            <el-table-column label="发送时间" width="160px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.createTime | parseTimeFilter}}</span>
                </template>
            </el-table-column>
            <el-table-column label="发送内容" min-width="160px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.content}}</span>
                </template>
            </el-table-column>
            <el-table-column label="开始时间" width="160px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.beginTime | parseTimeFilter}}</span>
                </template>
            </el-table-column>
            <el-table-column label="结束时间" width="160px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.endTime | parseTimeFilter}}</span>
                </template>
            </el-table-column>
            <el-table-column label="时间间隔(分钟)" width="120px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.intervalTime}}</span>
                </template>
            </el-table-column>
            <el-table-column label="是否成功" width="100px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.successFlag | yesOrNoFilter}}</span>
                </template>
            </el-table-column>
            <el-table-column label="错误信息" width="100px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.errorMsg}}</span>
                </template>
            </el-table-column>
        </el-table>
    <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>
 </div>    
    
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { noticeList } from "@/api/notice";
import { parseTime } from "@/utils/index";
import { yesOrNoList } from "@/types/public";
import { Message, MessageBox } from "element-ui";
export default {
  name: "NoticeList",
  directives: {
    waves
  },
  filters: {
    parseTimeFilter: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    },
    yesOrNoFilter:function(value){
        if(value ==0){
            return "否"
        }
        return "是"
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
        playerId: undefined,
        successFlagSel: undefined,
        successFlag: -1
      },
      logData: [],
      total: 0,
      listLoading: false,
      tableMetaMap: new Map(),
      metaColumnArray: [],
      channelList: [],
      tempPlatformList: [],
      yesOrNoArray: []
    };
  },
  methods: {
    handleFilter(e) {
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
      this.yesOrNoArray = yesOrNoList;
    },
    handleYesOrNoChange:function(e){
        if(e){
            this.listQuery.successFlag = parseInt(e)
        }else{
            this.listQuery.successFlag = -1
        }
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
        this.listQuery.centerPlatformId = item.centerPlatformId;
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
    findChannelName(channelId) {
        let list = this.channelList.filter(function(item, index) {
            return item.channelId == channelId;
        })
        if(list && list.length > 0){
            return list[0].channelName
        }
        return ""
    },
    findPlatformName(platformId) {
        let item = this.findPlatFormItem(platformId)
        if(item){
            return item.platformName
        }
        return ""
    },
    loadData() {
      this.listLoading = true;
      noticeList(this.listQuery).then(res => {
        this.logData = res.itemArray;
        this.total = res.total;
        this.listLoading = false;
      });
    }
  }
};
</script>

