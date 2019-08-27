<template>
    <div class="app-container">
        <div class="filter-container">
            <el-select v-model="listQuery.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>

            <el-select v-model="listQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handlePlatformChange">
                <el-option v-for="item in tempPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>

            <el-select v-model="listQuery.serverId" collapse-tags placeholder="服务器" clearable style="width: 220px" class="filter-item" @change="handleServerChange" >
              <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>

            <el-input placeholder="最小充值" v-model="listQuery.startMoney" type="number" style="width: 160px;" class="filter-item"/>
            <el-input placeholder="最大充值" v-model="listQuery.endMoney" type="number" style="width: 160px;" class="filter-item"/>
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
            height="800"
            style="width: 100%;margin-top:15px;">
            <el-table-column fixed="left" label="消耗原因" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ goldTypeName(scope.row.reason) }}</span>
                </template>
            </el-table-column>
            <el-table-column  v-for="(item,index) in queryServerList" :label="item.serverName"  align="center" width="300px" v-bind:key="index" >
                <el-table-column label="消耗数量" align="center" width="140px" >
                    <template slot-scope="scope">
                        <span>{{ scope.row.serverMap[item.serverId] ? scope.row.serverMap[item.serverId].changeNum : "" }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="消耗占比" align="center" width="80px" >
                    <template slot-scope="scope">
                        <span>{{ scope.row.serverMap[item.serverId] && item.changeNum>0 ? (scope.row.serverMap[item.serverId].changeNum*100/item.changeNum).toFixed(2) +"%" : "" }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="人数" align="center" width="80px" >
                    <template slot-scope="scope">
                        <span>{{ scope.row.serverMap[item.serverId] ? scope.row.serverMap[item.serverId].playerCount : "" }}</span>
                    </template>
                </el-table-column>      
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
import { getGoldChange, getGoldChangeType } from "@/api/goldchange";

export default {
  name: "GetGoldChangeList",
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
        goldType:2,
        pageIndex: 1,
        allianceName: "",
        ordercol: 1,
        ordertype: 0,
        timeArray:[],
        platformId: undefined,
        channelId: undefined,
        serverId: undefined,
        startMoney: undefined,
        endMoney: undefined
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
      totalMoney: 0,
      queryServerList: [],
      channelList: [],
      platformList: [],
      groupList: [],
      tempPlatformList: [],
      serverList: [],
      monitorTemp: {},
      goldType: {}
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
      this.listQuery.startMoney = parseInt(this.listQuery.startMoney);
      this.listQuery.endMoney = parseInt(this.listQuery.endMoney);
      console.log(this.listQuery);
      this.queryServerList = [];
      this.list = [];
      this.getList();
    },

    getList() {
      this.listLoading = true;
      getGoldChange(this.listQuery)
        .then(res => {
          this.list = res.itemArray;
          this.totalMoney = res.totalChangeNum;
          this.queryServerList = res.serverList;
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
      this.listQuery.timeArray = [startDate, endDate];
      getAllChannel().then(res => {
        this.channelList = res.itemArray;
      });
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
        // this.tempPlatformList = this.platformList;
      });
       let postData = {
        goldType:this.listQuery.goldType
      }
      getGoldChangeType(postData).then(res => {
        this.goldType = res;
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
    goldTypeName(goldType) {
      if (this.goldType[goldType]) {
        return this.goldType[goldType];
      }
      return "";
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

