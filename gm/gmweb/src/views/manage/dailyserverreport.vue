<template>
    <div class="app-container">
        <div class="filter-container">
            <el-select v-model="listQuery.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>

            <el-select v-model="listQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item">
                <el-option v-for="item in tempPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>

            <div class="filter-item">
                <el-date-picker v-model="listQuery.timeArray" type="datetimerange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期">
                </el-date-picker>
            </div>
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
        </div>

        <el-table
            v-loading="listLoading"
            :key="tableKey"
            :data="list"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;">
            <el-table-column label="服务器Id" align="center" width="80px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.serverId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="服务器名" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.serverName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="合服序号" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.serverChildStr }}</span>
                </template>
            </el-table-column>
            <el-table-column label="第一战力" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.firstPower }}</span>
                </template>
            </el-table-column>
            <el-table-column label="第二战力" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.secondPower }}</span>
                </template>
            </el-table-column>
            <el-table-column label="第三战力" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.thirePower }}</span>
                </template>
            </el-table-column>
            <el-table-column label="项目类型" align="center" width="100px" >
                <template slot-scope="scope">
                    <span>最高在线</span><br />
                    <span>登录人数</span><br />
                    <span>当日金额</span>
                </template>
            </el-table-column>
            <el-table-column  v-for="(item,index) in datelist" :label="item | parseTime"  align="center" width="100px" v-bind:key="index" >
                <template slot-scope="scope">
                    <span>{{scope.row.dailyData[item].maxOnLine}}</span><br />
                    <span>{{scope.row.dailyData[item].loginNum}}</span><br />
                    <span>{{scope.row.dailyData[item].orderAmount}}</span>
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
import { getServerDailyStatic } from "@/api/serverdaily";

export default {
  name: "DailyServerReportList",
  directives: {
    waves
  },
  filters: {
    parseTime: function(value) {
      return parseTime(value, "{m}-{d}");
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
    parseStatus: function(value) {
      console.log(value);
      let item = tradeItemStateList[value];
      console.log(item);
      if (item) {
        return item.name;
      }
      return value;
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
        timeArray: [],
        platformId: undefined,
        channelId: undefined,
        serverId: undefined,
        centerPlatformId: undefined,
        centerServerId: undefined,
        tradeId: undefined,
        playerId: undefined,
        level: undefined,
        state: undefined
      },
      textMap: {
        update: "编辑",
        create: "添加"
      },
      dialogStatus: "",
      dialogPvVisible: false,
      dialogFormVisible: false,
      temp: {},
      list: [],
      datelist:[],
      queryServerList: [],
      channelList: [],
      platformList: [],
      groupList: [],
      tempPlatformList: [],
      serverList: [],
      monitorTemp: {}
    };
  },
  methods: {
    handleFilter: function() {
      if (!this.listQuery.platformId) {
        Message({
          message: "请选择平台",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      if (this.listQuery.timeArray && this.listQuery.timeArray.length == 2) {
        this.listQuery.startTime = this.listQuery.timeArray[0].valueOf();
        this.listQuery.endTime = this.listQuery.timeArray[1].valueOf();
      }
      this.list = [];
      this.getList();
    },

    getList() {
      this.listLoading = true;
      this.listQuery.centerPlatformId = undefined;
      this.listQuery.centerServerId = undefined;
      if (this.listQuery.platformId) {
        let platformInfo = this.findPlatFormItem(this.listQuery.platformId);
        if (platformInfo) {
          this.listQuery.centerPlatformId = platformInfo.centerPlatformId;
        }
      }
      if (this.listQuery.serverId) {
        let serverInfo = this.findServerItem(this.listQuery.serverId);
        if (serverInfo) {
          this.listQuery.centerServerId = serverInfo.serverId;
        }
      }
      getServerDailyStatic(this.listQuery)
        .then(res => {
          this.list = res.itemArray;
          this.datelist = res.dailyArray;
          this.listLoading = false;
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
    handleServerChange: function(e) {
      if (e) {
        this.listQuery.serverId = e;
      } else {
        this.listQuery.serverId = undefined;
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

